package main

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0067h/gitrank/internal/config"
	myredis "github.com/n0067h/gitrank/internal/redis"
	"time"
)

func main() {
	config.Load()

	rdb := myredis.Init()
	for myredis.Init(); rdb == nil; myredis.Init() {
		log.Warn("Failed to connect to Redis; Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}
	log.Info("Connected to Redis server")

	go func() {
		for {
			myredis.SetHeartbeat(rdb)
			time.Sleep(10 * time.Second)
		}
	}()

	updateChannel := myredis.Subscribe(rdb, "ranking:update_request")
	myredis.CacheRanking(rdb, updateChannel)
}
