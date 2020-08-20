package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "github.com/ycjiafei/go-micro-project/user-srv/proto"
)

const (
	port = ":50050"
)

type server struct {
	pb.UnimplementedUserServer
}

func (s *server) GetUserByID(ctx context.Context, req *pb.UidReq) (*pb.UserInfoReply, error) {
	return &pb.UserInfoReply{
		Id: req.Uid,
		Name: "邱佳飞",
		Phone: 13100635547,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen grpc server err: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to user service err: %v", err)
	}
}
