package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/drornir/cloudex/pkg/config"
)

//go:embed html css
var fileSystem embed.FS

func main() {
	var conf config.Config
	if err := conf.Factor3Load(os.Args[1:]); err != nil {
		log.Fatalf("error: loading config: %s", err)
	}

	logger := newLogger(conf)

	s, err := NewServer(logger, fileSystem)
	if err != nil {
		log.Fatalf("error: initializing server: %s", err)
	}

	logger.Info("server listening", "port", conf.Port)
	_ = http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", conf.Port), s)
}
