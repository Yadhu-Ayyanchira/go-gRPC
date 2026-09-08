package main

import (
	"io"
	"log"

	pb "github.com/Yadhu-Ayyanchira/go-gRPC/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{
				Messages: messages,
			})
		}
		if err != nil {
			return err
		}

		log.Printf("got request with names : %v", req.Names[0])
		// for _, name := range req.Names {
		// 	messages = append(messages, "Hello "+name)
		// }
		var key string
		if req.Names[0] == "Amal" {
			key = "Devops Engineer"
		} else if req.Names[0] == "Naveen" {
			key = "Software Engineer"
		} else {
			key = "Tester"
		}
		log.Printf("Hello " + req.Names[0] + " your designation is " + key)
		messages = append(messages, "Hello "+req.Names[0]+" your designation is "+key)

	}

}
