package main

import (
	"context"
	"log"
	"time"

	pb "github.com/asepnur/simple-grpc/user-services/grpc"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:9001"
	defaultName = "asepnur"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Ping(ctx, &pb.Hello{})
	if err != nil {
		log.Println("Err", err)
	}
	log.Println(r.Greeting)
	resp, err := c.SelectUserInfo(ctx, &pb.UserResponse{})
	if err != nil {
		log.Println("Err", err)
	}
	log.Printf("ID\t\tName\t\t Email \t MSISDN\n")
	for _, val := range resp.User {
		log.Printf("%v\t%v\t\t %v \t %v",
			val.ID, val.FullName, val.Email, val.MSIDN,
		)
	}
}
