package main

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	pb "github.com/yank1/golang-helloworld/pkg/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Empty test to avoid no-tests warning.
func TestEmpty(t *testing.T) {}

func startServer() {
	ctxServer, cancelServer := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelServer()
	go RunInProcessGateway(ctxServer, ":8080")

	time.Sleep(1 * time.Second)
}

func TestMainGRPC(m *testing.T) {
	startServer()
	// Set up a connection to the server.
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		m.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewHelloworldClient(conn)
	ctx, client := context.WithTimeout(context.Background(), time.Second)
	defer client()
	// Contact the server and print out its response.
	r, err := c.Ping(ctx, &emptypb.Empty{})
	if err != nil {
		m.Fatal(err)
	}
	is := assert.New(m)
	is.Equal(r.GetMessage(), "Hello World")

	log.Printf("Ping Sucess: %s", r.GetMessage())

	time.Sleep(1000 * time.Second)

}
