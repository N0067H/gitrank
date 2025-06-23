package main

import (
	"github.com/gbswhs/gbsw-gitrank/api/config"
	"github.com/gbswhs/gbsw-gitrank/api/route"
	"github.com/gbswhs/gbsw-gitrank/api/rpc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	conn := rpc.ConnectToWorker()
	defer conn.Close()
	c := rpc.GetRankClient(conn)

	app := fiber.New()
	route.SetupRoute(app, c)

	log.Fatal(app.Listen(":" + config.GetConfig().APIPort))
}
