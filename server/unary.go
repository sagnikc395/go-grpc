package main

import (
	"context"

	pb "github.com/sagnikc395/go-grpc/proto"
)

func (hs *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}


