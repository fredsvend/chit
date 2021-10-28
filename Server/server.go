package main

import (
	"net"
	"google.golang.org/grpc"
	"log"
	"google.golang.org/grpc/reflection"
	context "golang.org/x/net/context"
	pb "github.com/fredsvend/Chit-Chat"
)

type server struct{}


const (
	port = ":8080"
)

func main() {

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	// Initializes the gRPC server.
	s := grpc.NewServer()

	// Register the server with gRPC.
	pb.RegisterChatServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}