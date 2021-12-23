package main

import (
	// "flag"

	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/soheilhy/cmux"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/yank1/golang-helloworld/pkg/helloworld"
)

// RunInProcessGateway starts the invoke in process http gateway.
func RunInProcessGateway(ctx context.Context, addr string, opts ...runtime.ServeMuxOption) error {
	// Create the main listener.
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	// Create a cmux.
	m := cmux.New(l)

	// Match connections in order:
	// First grpc, then HTTP, and otherwise Go RPC/TCP.
	grpcL := m.Match(cmux.HTTP2())
	httpL := m.Match(cmux.HTTP1Fast())

	gin = cmux.HTTP1HeaderFieldPrefix('path','/v1/hahaha')
	// trpcL := m.Match(cmux.Any()) // Any means anything that is not yet matched.

	// Create your protocol servers.
	grpcS := grpc.NewServer()
	pb.RegisterHelloworldServer(grpcS, pb.NewHelloWorldServer())
	// Register reflection service on gRPC server.
	reflection.Register(grpcS)
	// grpchello.RegisterGreeterServer(grpcS, &server{})

	//HTTP
	mux := runtime.NewServeMux(opts...)
	pb.RegisterHelloworldHandlerServer(ctx, mux, pb.NewHelloWorldServer())
	httpS := &http.Server{
		Handler: mux,
	}

	// trpcS := rpc.NewServer()
	// trpcS.Register(&ExampleRPCRcvr{})

	// Use the muxed listeners for your servers.
	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)
	// go trpcS.Accept(trpcL)

	// Start serving!
	if err := m.Serve(); !strings.Contains(err.Error(), "use of closed network connection") {
		panic(err)
	}
	return nil
}

func main() {
	//Init the command-line flags.
	flag.Parse()

	ctx, _ := context.WithCancel(context.Background())
	RunInProcessGateway(ctx, ":8080")

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
