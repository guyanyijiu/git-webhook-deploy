package router

import (
	"git-webhook-deploy/handler"
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(200, "git-webhook-deploy is running")
		return
	})

	r.POST("/github", handler.Github)
}
