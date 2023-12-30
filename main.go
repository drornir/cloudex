package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/drornir/cloudex/pkg/app"
	"github.com/drornir/cloudex/pkg/config"
	"github.com/drornir/cloudex/pkg/db"
	"github.com/drornir/cloudex/pkg/htmlserver"
)

//go:embed assets
var rootFS embed.FS

func main() {
	var conf config.Config
	if err := conf.Factor3Load(os.Args[1:]); err != nil {
		log.Fatalf("error: loading config: %s", err)
	}
	db, err := db.OpenLibSQL(conf.LibsqlURL)
	if err != nil {
		log.Fatalf("error: connecting to db at %q: %s", conf.LibsqlURL, err)
	}

	assetsFS, err := fs.Sub(rootFS, "assets")
	if err != nil {
		log.Fatalf("error creating sub FS for assets: %s", err)
	}

	appl := &app.App{
		DB:       db,
		AssetsFS: assetsFS,
	}

	s := htmlserver.New(appl)

	listenOn := fmt.Sprintf("127.0.0.1:%s", conf.Port)
	if err := s.Start(listenOn); err != nil {
		s.Logger.Fatalf("server finished with %s", err)
	}
}
