package service

import (
	"github.com/gin-gonic/gin"
	"github.com/ycjiafei/go-micro-project/api/model"
	pb "github.com/ycjiafei/go-micro-project/user-srv/proto"
	"google.golang.org/grpc"
)

type userService struct {
	c *gin.Context
	conn *grpc.ClientConn
}

func NewUserService(c *gin.Context) (*userService, error) {
	conn, err := grpc.Dial(
		"127.0.0.1:50050",
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	return &userService{
		c: c,
		conn: conn,
	}, err
}


func (us userService) GetUserInfoByID(uid int64) model.UserInfo {
}
