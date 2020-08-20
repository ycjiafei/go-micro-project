package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ycjiafei/go-micro-project/pkg/errcode"
	"net/http"
)

func SuccessResp(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": errcode.OK,
		"msg": errcode.ErrMsg[errcode.OK],
		"data": data,
	})
}

func FailResp(c *gin.Context, ecode int, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": ecode,
		"msg": errcode.ErrMsg[ecode],
		"data": err,
	})
}
