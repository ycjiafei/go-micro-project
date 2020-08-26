package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ycjiafei/go-micro-project/api/errcode"
	"github.com/ycjiafei/go-micro-project/api/service"
)

func Login(c *gin.Context) {
	type login struct {
		Phone string `json:"phone" binding:"required"`
		Code int `json:"code" binding:"required"`
	}
	form := login{}
	if err := c.ShouldBindJSON(&form); err != nil {
		FailResp(c, errcode.MissArgument, err.Error())
		return
	}
	srv, err := service.NewUserService(c)
	if err != nil {
		FailResp(c, errcode.NewServiceFail, err.Error())
		return
	}
	info := srv.GetUserInfoByPhone(form.Phone)
	if info.ID == 0 {
		FailResp(c, errcode.UserNotFound, "未找到对应手机的用户")
		return
	}
	SuccessResp(c, info)
}
