package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ycjiafei/go-micro-project/api/errcode"
	"github.com/ycjiafei/go-micro-project/api/service"
)

func Login(c *gin.Context) {
	type login struct {
		Phone int64 `json:"phone" binding:"required"`
		Code int `json:"code" binding:"required"`
	}
	form := login{}
	if err := c.ShouldBindJSON(&form); err != nil {
		fmt.Println(err)
		FailResp(c, errcode.MissArgument, err.Error())
		return
	}
	fmt.Println(form)
	srv, err := service.NewUserService(c)
	if err != nil {
		FailResp(c, errcode.NewServiceFail, err.Error())
		return
	}
	SuccessResp(c, srv.GetUserInfoByID(1))
}
