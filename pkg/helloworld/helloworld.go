package helloworld

import (
	// "github.com/yank1/golang-helloworld/pkg/helloworld"
	// pb "github.com/yank1/golang-helloworld/pkg/helloworld"
	"github.com/golang/glog"

	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

func NewHelloWorldServer() HelloworldServer {
	return new(server)
}

// SayHello implements helloworld.GreeterServer
func (s *server) Ping(context.Context, *emptypb.Empty) (*PongResponse, error) {
	glog.Infof("Ping")
	return &PongResponse{Message: "Hello World"}, nil
}
