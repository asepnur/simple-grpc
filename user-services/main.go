package main

import (
	"log"
	"net"

	pb "github.com/asepnur/simple-grpc/user-services/grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

func main() {
	// Set up a connection to the server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Ping(ctx context.Context, in *pb.Hello) (*pb.Hello, error) {
	return &pb.Hello{
		Greeting: "Connected to user services",
	}, nil
}
