package rpc

import (
	"context"
	"github.com/ycjiafei/go-micro-project/user-srv/dao"
	pb "github.com/ycjiafei/go-micro-project/user-srv/proto"
)

type UserRPC struct {
	pb.UnimplementedUserServer
}

func (s *UserRPC) GetUserByID(ctx context.Context, req *pb.UidReq) (*pb.UserInfoReply, error) {
	info := dao.GetUserInfoByUid(req.Uid)
	return &pb.UserInfoReply{
		Id: info.ID,
		Name: info.Name,
		Phone: info.Phone,
	}, nil
}

func (s *UserRPC) GetUserByPhone(ctx context.Context, req *pb.PhoneReq) (*pb.UserInfoReply, error) {
	info := dao.GetUserInfoByPhone(req.Phone)
	return &pb.UserInfoReply{
		Id: info.ID,
		Name: info.Name,
		Phone: info.Phone,
	}, nil
}
