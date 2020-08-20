package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ycjiafei/go-micro-project/pkg/errcode"
)

func Login(c *gin.Context) {
	phone := c.PostForm("phone")
	code := c.PostForm("code")
	if phone == "" || code == "" {
		FailResp(c, errcode.Bad, errors.New("缺少必要参数"))
		return
	}
	SuccessResp(c, "登录成功")
}
