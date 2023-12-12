package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/drornir/cloudex/pkg/config"
	"github.com/drornir/cloudex/pkg/server"
)

//go:embed html css assets
var fileSystem embed.FS

func main() {
	var conf config.Config
	if err := conf.Factor3Load(os.Args[1:]); err != nil {
		log.Fatalf("error: loading config: %s", err)
	}

	logger := newLogger(conf)

	s, err := server.New(logger, fileSystem)
	if err != nil {
		log.Fatalf("error: initializing server: %s", err)
	}

	listenOn := fmt.Sprintf("127.0.0.1:%s", conf.Port)
	logger.Info(fmt.Sprintf("server listening on http://%s ", listenOn))
	_ = http.ListenAndServe(listenOn, s.HTTPHandler())
}
