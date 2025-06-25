package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0067h/gitrank/internal/api/route"
	"github.com/n0067h/gitrank/internal/config"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	app := fiber.New()
	route.SetupRoute(app)

	log.Fatal(app.Listen(":" + config.AppConfig.APIPort))
}
