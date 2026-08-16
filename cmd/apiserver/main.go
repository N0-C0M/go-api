package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/N0-C0M/go-api/internal/app/apiserver"
)

var (
	confgPath string
)

func init() {
	flag.StringVar(&confgPath, "config-path", "configs/apiserver.toml", "Path to configuration file")
}
func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(confgPath, config)
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
