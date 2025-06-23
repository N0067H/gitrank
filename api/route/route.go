package route

import (
	"github.com/gbswhs/gbsw-gitrank/api/rpc"
	rank "github.com/gbswhs/gbsw-gitrank/proto"
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App, c rank.RankClient) {
	app.Get("/rank", rpc.GetRank(c))
}
