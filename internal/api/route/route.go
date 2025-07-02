package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n0067h/gitrank/internal/api/handler"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/ping", pong)
	app.Get("/rank", handler.GetRanking)
}

func pong(c *fiber.Ctx) error {
	return c.SendString("pong")
}
