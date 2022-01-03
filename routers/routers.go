package routers

import "github.com/gin-gonic/gin"

var MainRouter *gin.Engine

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":  "pong",
		"clientIP": c.ClientIP(),
	})
}

func init() {
	MainRouter = gin.Default()
	MainRouter.GET("/ping", pingHandler)
}

func Run() error {
	return MainRouter.Run(":12345")
}
