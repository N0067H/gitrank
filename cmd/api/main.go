package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/n0067h/gitrank/internal/api/route"
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

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.AppConfig.AllowOrigins,
	}))
	route.SetupRoutes(app)

	if err := app.Listen(":" + config.AppConfig.APIPort); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
