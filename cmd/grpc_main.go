package main

import (
	"log"
	"net"

	"github.com/yank1/golang-helloworld/pkg/helloworld"
	// pb "github.com/yank1/golang-helloworld/pkg/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

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
