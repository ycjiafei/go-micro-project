package main

import (
	slog "github.com/ycjiafei/go-micro-project/pkg/log"
	"github.com/ycjiafei/go-micro-project/pkg/trace"
	pb "github.com/ycjiafei/go-micro-project/user-srv/proto"
	"github.com/ycjiafei/go-micro-project/user-srv/rpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

const (
	port = ":50050"
)

func main() {
	slog.InitDefault()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		slog.SLog.Bg().Fatal("listen grpc server err", zap.String("err", err.Error()))
	}
	tracer, closer, err := trace.NewJaeger("USER_GRPC")
	if err != nil {
		slog.SLog.Bg().Error("new user_grpc_tracer err", zap.String("err", err.Error()))
	}
	defer closer.Close()
	s := grpc.NewServer(grpc.UnaryInterceptor(trace.ServerInterceptor(tracer)))
	pb.RegisterUserServer(s, &rpc.UserRPC{})

	if err := s.Serve(lis); err != nil {
		slog.SLog.Bg().Error("failed to user service err", zap.String("err", err.Error()))
	}
}
