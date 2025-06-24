package main

import (
	rank "github.com/gbswhs/gbsw-gitrank/proto"
	"github.com/gbswhs/gbsw-gitrank/worker/config"
	"github.com/gbswhs/gbsw-gitrank/worker/rpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	lis, err := net.Listen("tcp", ":"+config.GetConfig().WorkerPort)
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	s := grpc.NewServer()
	rank.RegisterRankServer(s, &rpc.Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	log.Println("Server is running on port " + config.GetConfig().WorkerPort)
}
