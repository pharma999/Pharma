package routes

import (
	"demo/controller"
	"demo/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/users", controller.CreateUser)
		api.GET("/user", controller.GetUsers)
		api.POST("/login", controller.LoginUser)
		api.POST("/verify", controller.VerifyOTP)
		// api.POST("/userdetail", controller.CreateUserDetail)
		api.POST("/userdetailmongo", controller.CreateUserDetailMongo)
		api.POST("/mongouser", controller.CreateUserMongo)

	}

	secured := api.Group("/secured")
	secured.Use(middleware.AuthMiddleware())
	{
		secured.GET("/user", controller.GetUsers)

	}
}
