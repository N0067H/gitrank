package transport

import (
	"context"
	"github.com/gbswhs/gbsw-gitrank/internal/protobuf/rank"
	"github.com/gbswhs/gbsw-gitrank/internal/worker/ghclient"
	"log"
)

type Server struct {
	rank.UnimplementedRankServer
}

func (s *Server) GetRankings(ctx context.Context, req *rank.RankRequest) (*rank.RankReply, error) {
	log.Printf("GetRanks called with %v", req)

	ghclient.GetRanks()

	return &rank.RankReply{
		Users: ghclient.ToProtoUsers(ghclient.Users),
	}, nil
}
