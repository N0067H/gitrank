package main

import (
	"github.com/gbswhs/gbsw-gitrank/internal/worker/config"
	"github.com/gbswhs/gbsw-gitrank/internal/worker/transport"
	"log"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	if err := transport.Run(); err != nil {
		log.Fatalf("failed to run server: %s", err)
	}
}
