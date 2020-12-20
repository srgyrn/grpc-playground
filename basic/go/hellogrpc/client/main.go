package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	"test/playground/rpc/grpc/basic/go/hellogrpc/messages"
)

const (
	address = "localhost:50000"
	defaultName = "world"
)

func main() {
	// set up connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := messages.NewHelloServiceClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &messages.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greetings: %s", r.GetMessage())
}