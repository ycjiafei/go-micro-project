package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ycjiafei/go-micro-project/api/handler"
	"github.com/ycjiafei/go-micro-project/api/middleware"
)

func InitRoutes() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.JaegerTrace)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/login", handler.Login)
	return r
}
