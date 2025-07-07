package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0067h/gitrank/internal/config"
	myredis "github.com/n0067h/gitrank/internal/redis"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}

func run() error {
	err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load .env file")
	}

	myredis.Init()
	if myredis.Rdb == nil {
		return fmt.Errorf("failed to initialize Redis client")
	}

	pubsub := myredis.Subscribe("ranking:update_request")
	if pubsub == nil {
		return fmt.Errorf("failed to subscribe to channel 'ranking:update_request'")
	}

	log.Info("Connected to Redis server")

	return myredis.CacheRanking(pubsub)
}
