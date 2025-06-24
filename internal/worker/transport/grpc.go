package transport

import (
	"context"
	"fmt"
	"github.com/gbswhs/gbsw-gitrank/internal/protobuf/rank"
	"github.com/gbswhs/gbsw-gitrank/internal/worker/config"
	"github.com/gbswhs/gbsw-gitrank/internal/worker/ghclient"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	rank.UnimplementedRankServer
}

func Run() error {
	addr := ":" + config.GetConfig().WorkerPort
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", addr, err)
	}

	s := grpc.NewServer()
	rank.RegisterRankServer(s, &Server{})

	log.Printf("gRPC server starting on %s", addr)
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

func (s *Server) GetRankings(ctx context.Context, req *rank.RankRequest) (*rank.RankReply, error) {
	log.Printf("GetRanks called with %v", req)

	ghclient.GetRanks()

	return &rank.RankReply{
		Users: ghclient.ToProtoUsers(ghclient.Users),
	}, nil
}
