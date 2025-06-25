package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/n0067h/gitrank/internal/api/transport"
	"github.com/n0067h/gitrank/internal/protobuf/rank"
)

func SetupRoute(app *fiber.App, c rank.RankClient) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Get("/rank", transport.GetRank(c))
}
