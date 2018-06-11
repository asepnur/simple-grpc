package user

import (
	"context"

	pb "github.com/asepnur/simple-grpc/user-services/grpc"
	ctr "github.com/asepnur/simple-grpc/user-services/webserver/handler"
)

type UserServer struct{}

// Ping function
func (s *UserServer) Ping(ctx context.Context, in *pb.Hello) (*pb.Hello, error) {
	return &pb.Hello{
		Greeting: "Connected to user services",
	}, nil
}

// SelectUserInfo function
func (s *UserServer) SelectUserInfo(ctx context.Context, in *pb.UserResponse) (*pb.UserResponse, error) {
	var response *pb.UserResponse
	params := ctr.UserParams{
		Q: "",
	}
	resp, _, err := params.SelectUserController()
	if err != nil {
		return response, err
	}
	var tmpUser []*pb.UserInfo
	for _, val := range resp {
		tmpUser = append(tmpUser, &pb.UserInfo{
			ID:       val.UserID,
			FullName: val.FullName,
			Email:    val.UserEmail,
			MSIDN:    val.MSISDN,
		})
	}
	return &pb.UserResponse{
		User: tmpUser,
	}, nil
}
