package route

import (
	"github.com/gbswhs/gbsw-gitrank/internal/api/transport"
	"github.com/gbswhs/gbsw-gitrank/internal/protobuf/rank"
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App, c rank.RankClient) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Get("/rank", transport.GetRank(c))
}
