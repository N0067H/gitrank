package rpc

import (
	"context"
	rank "github.com/gbswhs/gbsw-gitrank/proto"
	"github.com/gbswhs/gbsw-gitrank/worker/ghclient"
	"log"
)

type Server struct {
	rank.UnimplementedRankServer
}

func ToProtoUsers(users []ghclient.User) []*rank.User {
	var protoUsers []*rank.User

	for _, user := range users {
		protoUsers = append(protoUsers, &rank.User{
			Login:              string(user.Login),
			TotalContributions: uint32(user.TotalContributions),
		})
	}

	return protoUsers
}

func (s *Server) GetRankings(ctx context.Context, req *rank.RankRequest) (*rank.RankReply, error) {
	log.Printf("GetRanks called with %v", req)
	return &rank.RankReply{
		Users: ToProtoUsers(ghclient.Users),
	}, nil
}
