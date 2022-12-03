package main

import (
	"log"

	server "github.com/kalpit-sharma-dev/url-shortner/src/server"
)

func main() {

	log.Print("starting the application ....")
	server.LoadRoute()

}
