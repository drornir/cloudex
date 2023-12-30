package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/drornir/cloudex/pkg/app"
	"github.com/drornir/cloudex/pkg/config"
	"github.com/drornir/cloudex/pkg/db"
	"github.com/drornir/cloudex/pkg/htmlserver"
)

//go:embed css assets
var fileSystem embed.FS

func main() {
	var conf config.Config
	if err := conf.Factor3Load(os.Args[1:]); err != nil {
		log.Fatalf("error: loading config: %s", err)
	}
	db, err := db.New(conf.SQLiteURL)
	if err != nil {
		log.Fatalf("error: connecting to db: %s", err)
	}

	appl := &app.App{
		DB: db,
	}

	s := htmlserver.New(appl)

	listenOn := fmt.Sprintf("127.0.0.1:%s", conf.Port)
	if err := s.Start(listenOn); err != nil {
		s.Logger.Fatalf("server finished with %s", err)
	}
}
