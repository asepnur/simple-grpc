package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/asepnur/galahad/src/util/jsonconfig"
	pb "github.com/asepnur/simple-grpc/user-services/grpc"
	"github.com/asepnur/simple-grpc/user-services/util/env"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/asepnur/iskandar/src/util/conn"
	"github.com/asepnur/iskandar/src/webserver"
)

const (
	port = ":50051"
)

type configuration struct {
	Database  conn.DatabaseConfig `json:"database"`
	Redis     conn.RedisConfig    `json:"redis"`
	Webserver webserver.Config    `json:"webserver"`
}

// server is used to implement helloworld.GreeterServer.
type server struct{}

func main() {
	flag.Parse()

	//Load configuration
	cfgenv := env.Get()
	config := &configuration{}
	fmt.Println("./files/etc", cfgenv)
	isLoaded := jsonconfig.Load(&config, "/etc", cfgenv) || jsonconfig.Load(&config, "./files/etc", cfgenv)
	if !isLoaded {
		log.Fatal("Failed to load configuration")
	}
	conn.InitDB(config.Database)
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
