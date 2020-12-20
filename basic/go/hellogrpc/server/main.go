package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"test/playground/rpc/grpc/basic/go/hellogrpc/messages"
)

const port = ":50000"

type server struct{
	messages.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *messages.HelloRequest) (*messages.HelloResponse, error) {
	return &messages.HelloResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	messages.RegisterHelloServiceServer(s, &server{})
	s.Serve(lis)
}
