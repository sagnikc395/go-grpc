package main

import (
	"log"

	pb "github.com/sagnikc395/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	PORT = ":8080"
)

func main() {
	//client start
	conn, err := grpc.NewClient("localhost"+PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"sagnikc", "alice", "bob"},
	}

	callSayHello(client)
}
