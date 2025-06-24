package main

import (
	"github.com/gbswhs/gbsw-gitrank/internal/api/config"
	"github.com/gbswhs/gbsw-gitrank/internal/api/route"
	"github.com/gbswhs/gbsw-gitrank/internal/api/transport"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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
