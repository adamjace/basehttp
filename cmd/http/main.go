package main

import (
	"log"

	"github.com/adamjace/basehttp"
)

func main() {

	config := basehttp.NewConfig()

	log.Printf("Starting HTTP service on localhost%s", config.Bind)

	server := basehttp.NewServer(config)

	if err := server.Start(); err != nil {
		log.Fatalf("server.Start error: %v", err)
	}
}
