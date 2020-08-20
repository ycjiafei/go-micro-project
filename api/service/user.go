package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ycjiafei/go-micro-project/database/structs"
	pb "github.com/ycjiafei/go-micro-project/user-srv/proto"
	"google.golang.org/appengine/log"
	"google.golang.org/grpc"
	"time"
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


func (us userService) GetUserInfoByID(uid int64) structs.UserInfo {
	cli := pb.NewUserClient(us.conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := cli.GetUserByID(ctx, &pb.UidReq{Uid: uid})
	if err != nil {
		log.Errorf(ctx,"GRPC 请求错误, 错误方法 GetUserByID, err: %v", err)
	}
	return structs.UserInfo{
		ID:response.Id,
		Name: response.Name,
		Phone: response.Phone,
	}
}
