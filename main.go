package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/drornir/cloudex/pkg/config"
)

//go:embed html
var fileSystem embed.FS

func main() {
	var conf config.Config
	if err := conf.Factor3Load(os.Args[1:]); err != nil {
		log.Fatalf("error: loading config: %s", err)
	}
	fmt.Println(conf)
}
