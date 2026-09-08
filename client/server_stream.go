package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Yadhu-Ayyanchira/go-gRPC/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Streaming started.")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.SayHelloServerStreaming(ctx, names)
	if err != nil {
		log.Fatalf("Could not greet %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving response: %v", err)
		}
		log.Printf("Greeting: %v", message)
	}
	log.Printf("Streaming finished.")
}
