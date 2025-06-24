package rpc

import (
	"context"
	"github.com/gbswhs/gbsw-gitrank/api/config"
	rank "github.com/gbswhs/gbsw-gitrank/proto"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type ResponseUser struct {
	Login              string `json:"login"`
	TotalContributions uint32 `json:"totalContributions"`
}

func ConnectToWorker() *grpc.ClientConn {
	conn, err := grpc.NewClient("localhost:"+config.GetConfig().WorkerPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	return conn
}

func GetRankClient(conn *grpc.ClientConn) rank.RankClient {
	return rank.NewRankClient(conn)
}

func ToResponseUsers(users []*rank.User) []*ResponseUser {
	var responseUsers []*ResponseUser
	for _, user := range users {
		responseUsers = append(responseUsers, &ResponseUser{
			Login:              user.Login,
			TotalContributions: user.TotalContributions,
		})
	}
	return responseUsers
}

func GetRank(c rank.RankClient) fiber.Handler {
	return func(cc *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		r, err := c.GetRankings(ctx, &rank.RankRequest{
			Organization: config.GetConfig().OrganizationName,
		})

		if err != nil {
			log.Fatalf("failed to get ranks: %v", err)
		}

		responseUsers := ToResponseUsers(r.Users)
		return cc.JSON(responseUsers)
	}
}
