package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0067h/gitrank/internal/api/config"
	"github.com/n0067h/gitrank/internal/api/route"
	"github.com/n0067h/gitrank/internal/api/transport"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	conn := transport.Connect()
	defer conn.Close()
	c := transport.GetRankClient(conn)

	app := fiber.New()
	route.SetupRoute(app, c)

	log.Fatal(app.Listen(":" + config.GetConfig().APIPort))
}
