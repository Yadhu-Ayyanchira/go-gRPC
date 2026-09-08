package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Yadhu-Ayyanchira/go-gRPC/proto"
)

func callHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Bi-directional streaming started.")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error receiving: %v", err)
			}
			log.Printf("Got %v", message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error sending: %v", err)
		}
		log.Printf("Sent %v", name)
		time.Sleep(2 * time.Second)
	}

	if err := stream.CloseSend(); err != nil {
		log.Fatalf("error closing: %v", err)
	}

	<-waitc
	log.Printf("Bi directional streaming end.")
}
