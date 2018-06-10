package user

import (
	"context"
	"fmt"

	pb "github.com/asepnur/simple-grpc/user-services/grpc"
)

type UserServer struct{}

func (s *UserServer) Ping(ctx context.Context, in *pb.Hello) (*pb.Hello, error) {
	return &pb.Hello{
		Greeting: "Connected to user services",
	}, nil
}

func (s *UserServer) SelectUserInfo(ctx context.Context, in *pb.UserInfo) (*pb.UserInfo, error) {
	fmt.Println("contex: ", ctx)
	return &pb.UserInfo{}, nil
}
