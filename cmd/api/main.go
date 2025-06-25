package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0067h/gitrank/internal/api/config"
	"github.com/n0067h/gitrank/internal/api/route"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	app := fiber.New()
	route.SetupRoute(app)

	log.Fatal(app.Listen(":" + config.GetConfig().APIPort))
}
