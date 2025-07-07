package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0067h/gitrank/internal/config"
	myredis "github.com/n0067h/gitrank/internal/redis"
	"time"
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

	updateChannel := myredis.Subscribe("ranking:update_request")
	if updateChannel == nil {
		return fmt.Errorf("failed to subscribe to channel 'ranking:update_request'")
	}
	log.Info("Connected to Redis server")

	go func() {
		for {
			myredis.SetHeartbeat()
			time.Sleep(10 * time.Second)
		}
	}()

	return myredis.CacheRanking(updateChannel)
}
