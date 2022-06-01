package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/Heroin-lab/diplom.git/internal/app/server"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	serv := server.New(config)

	if err := serv.Start(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ebash ih blyat")
}
