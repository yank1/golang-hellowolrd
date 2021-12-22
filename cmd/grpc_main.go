package main

import (
	"log"
	"net"

	"github.com/yank1/golang-helloworld/pkg/helloworld"
	// pb "github.com/yank1/golang-helloworld/pkg/helloworld"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) Ping(context.Context, *emptypb.Empty) (*helloworld.PongResponse, error) {
	return &helloworld.PongResponse{Message: "Hello World"}, nil
}

// grpcurl  -plaintext 127.0.0.1:50051 helloworld.Helloworld.Ping
func _main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// helloworld
	helloworld.RegisterHelloworldServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
