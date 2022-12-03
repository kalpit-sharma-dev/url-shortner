package main

import (
	"log"

	"github.com/kalpit-sharma-dev/url-shortner/src/config"
	server "github.com/kalpit-sharma-dev/url-shortner/src/server"
)

func main() {

	log.Print("INFO : starting the application ....")
	err := config.LoadDBConfig()
	if err != nil {
		log.Fatal("Unable to load file from env", err)
	}
	server.LoadRoute()

}
