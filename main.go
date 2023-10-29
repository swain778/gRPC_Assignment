package main

import (
	pb "gRPC_Assignment/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8000"
)

type helloServer struct {
	pb.UserServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, &helloServer{})

	log.Printf("server started at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
