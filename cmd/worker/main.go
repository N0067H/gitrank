package main

import (
	"github.com/n0067h/gitrank/internal/worker/config"
	"log"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("failed to load .env file")
	}
}
