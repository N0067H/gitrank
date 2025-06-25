package route

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// app.Get("/rank", transport.GetRank(c))
}
