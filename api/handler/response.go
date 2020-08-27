package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ycjiafei/go-micro-project/api/errcode"
	"github.com/ycjiafei/go-micro-project/api/middleware"
	slog "github.com/ycjiafei/go-micro-project/pkg/log"
	"go.uber.org/zap"
	"net/http"
)

func SuccessResp(c *gin.Context, data interface{}) {
	slog.SLog.For(c.MustGet(middleware.TracerCtx).(context.Context)).Info("请求成功, 返回数据", zap.Any("data", data))
	c.JSON(http.StatusOK, gin.H{
		"code": errcode.OK,
		"msg": errcode.ErrMsg[errcode.OK],
		"data": data,
	})
}

func FailResp(c *gin.Context, ecode int, err string) {
	slog.SLog.For(c.MustGet(middleware.TracerCtx).(context.Context)).Error("请求失败", zap.Any(errcode.ErrMsg[ecode], err))
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": ecode,
		"msg": errcode.ErrMsg[ecode],
		"data": err,
	})
}
