package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/n0067h/gitrank/internal/api/handler"
	"github.com/n0067h/gitrank/internal/config"
	myredis "github.com/n0067h/gitrank/internal/redis"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	config.Load()

	var rdb *redis.Client
	for {
		rdb = myredis.Init()
		if rdb != nil {
			break
		}
		log.Warn("Failed to connect to Redis; Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}
	log.Info("Connected to Redis server")

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.AppConfig.AllowOrigins,
	}))
	setupRoutes(app, rdb)

	go func() {
		for {
			myredis.CheckHeartbeat(rdb)
			time.Sleep(20 * time.Second)
		}
	}()

	if err := app.Listen(":" + config.AppConfig.APIPort); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func setupRoutes(app *fiber.App, rdb *redis.Client) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	app.Get("/rank", handler.GetRanking(rdb))
}
