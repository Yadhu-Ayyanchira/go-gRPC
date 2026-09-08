package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Yadhu-Ayyanchira/go-gRPC/proto"
)

func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Streaming started.")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.SayHelloClientStreaming(ctx)
	if err != nil {
		log.Fatalf("Could not greet %v", err)
	}

	for _, name := range names.Names {
		req := &pb.NameList{
			Names: []string{name},
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error sending request: %v", err)
		}
		log.Printf("Sent request with name: %v", name)
		time.Sleep(2 * time.Second)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}
	log.Printf("Greeting: %v", resp)

	log.Printf("Clent Streaming finished.")
	log.Printf("%v", resp.Messages)
}
