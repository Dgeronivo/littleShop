package main

import (
	apiserver "littleShop/project/src/apiServer"
	"log"

	"github.com/BurntSushi/toml"
)

const confPath = "project/configs/env.toml"

func main() {
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(confPath, config)
	if err != nil {
		log.Fatal(err)
	}

	server := apiserver.New(config)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
