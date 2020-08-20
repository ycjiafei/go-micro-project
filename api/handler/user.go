package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ycjiafei/go-micro-project/api/service"
	"github.com/ycjiafei/go-micro-project/pkg/errcode"
)

func Login(c *gin.Context) {
	type login struct {
		Phone int64 `json:"phone" bind:"require"`
		Code int `json:"code" bind:"require"`
	}
	form := login{}
	if err := c.ShouldBindJSON(&form); err != nil {
		FailResp(c, errcode.MissArgument, err)
		return
	}
	srv, err := service.NewUserService(c)
	if err != nil {
		FailResp(c, errcode.NewServiceFail, err)
	}
	SuccessResp(c, srv.GetUserInfoByID(1))
}
