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

func (s *Server) GetRankings(ctx context.Context, req *rank.RankRequest) (*rank.RankReply, error) {
	log.Printf("GetRanks called with %v", req)

	ghclient.GetRanks()
	
	return &rank.RankReply{
		Users: ghclient.ToProtoUsers(ghclient.Users),
	}, nil
}
