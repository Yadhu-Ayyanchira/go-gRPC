package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Yadhu-Ayyanchira/go-gRPC/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cance := context.WithTimeout(context.Background(), 10*time.Second)
	defer cance()

	res, err := client.SayHello(ctx, &pb.NoParams{})
	if err != nil {
		log.Fatalf("Could not greet %v", err)
	}
	log.Printf("Greeting: %v", res.Message)
}
