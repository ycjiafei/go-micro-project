package main

import (
	"github.com/ycjiafei/go-micro-project/pkg/trace"
	pb "github.com/ycjiafei/go-micro-project/user-srv/proto"
	"github.com/ycjiafei/go-micro-project/user-srv/rpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50050"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen grpc server err: %v", err)
	}
	tracer, closer, err := trace.NewJaeger("USER_GRPC")
	if err != nil {
		log.Println("user_grpc_tracer new err : ", err)
	}
	defer closer.Close()
	s := grpc.NewServer(grpc.UnaryInterceptor(trace.ServerInterceptor(tracer)))
	pb.RegisterUserServer(s, &rpc.UserRPC{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to user service err: %v", err)
	}
}
