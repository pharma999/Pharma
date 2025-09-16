package routes

import (
	"demo/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/users", controller.CreateUser)
		api.GET("/user", controller.GetUsers)
		api.POST("/login", controller.LoginUser)
		api.POST("/verify", controller.VerifyOTP)

	}
}
