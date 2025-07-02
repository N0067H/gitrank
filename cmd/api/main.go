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
	err := config.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	myredis.Init()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.AppConfig.AllowOrigins,
	}))
	route.SetupRoutes(app)

	log.Fatal(app.Listen(":" + config.AppConfig.APIPort))
}
