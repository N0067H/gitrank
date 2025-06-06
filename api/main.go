package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

var apiKey string

func main() {
	client := api.New(apiKey)

	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong!")
	})

	app.Get("/members", func(c *fiber.Ctx) error {
		members, err := client.GetMembers("gbswhs", 100, 1)
		if err != nil {

		}

		return c.JSON(members)
	})

	app.Get("/events", func(c *fiber.Ctx) error {
		members, err := client.GetMembers("gbswhs", 10, 1)
		if err != nil {

		}

		type Resp struct {
			Login          string `json:"login"`
			NumberOfEvents int    `json:"events"`
		}

		var resp []Resp

		for _, e := range members {
			cnt, err := client.CountEvents(e.Login)
			if err != nil {

			}

			fmt.Println(cnt)
			resp = append(resp, Resp{Login: e.Login, NumberOfEvents: cnt})
		}

		return c.JSON(resp)
	})

	log.Fatal(app.Listen("localhost:3000"))

	// for _, e := range users {
	// 	events, err := client.CountEvents(e.Login, 100, 1)
	// 	if err != nil {

	// 	}

	// 	fmt.Printf("%s's events: %d\n", e.Login, events)
	// }

}
