package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ycjiafei/go-micro-project/api/handler"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login", handler.Login)
	return r
}
