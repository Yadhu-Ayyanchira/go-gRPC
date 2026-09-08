package main

import (
	"context"

	pb "github.com/Yadhu-Ayyanchira/go-gRPC/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParams) (*pb.HelloResponse, error) {

	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
