package main

import (
	"io"
	"log"

	pb "github.com/Yadhu-Ayyanchira/go-gRPC/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %v", req.Name)
		if err := stream.Send(&pb.HelloResponse{
			Message: "Hello " + req.Name,
		}); err != nil {
			return err
		}
	}
}
