package main

import (
	"log"
	"net"

	pb "github.com/yank1/golang-helloworld/pkg/github.com/yank1/golang-helloworld/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) Ping(context.Context, *emptypb.Empty) (*pb.PongResponse, error) {
	return &pb.PongResponse{Message: "Hello World"}, nil
}

// grpcurl  -plaintext 127.0.0.1:50051 helloworld.Ping/Ping
func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// helloworld
	pb.RegisterPingServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
