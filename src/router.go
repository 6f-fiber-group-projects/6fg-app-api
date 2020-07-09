package main

import (
	"6fg-app-api/controller"
	"github.com/gin-gonic/gin"
)

func DefineRoutes() *gin.Engine {
	router := gin.Default()

	user := router.Group("user")
	{
		user.GET("/", controller.GetUsers)
		user.POST("/", controller.CreateUser)
		user.GET("/:userId", controller.GetUserById)
		user.PUT("/:userId", controller.UpdateUser)
		user.DELETE("/:userId", controller.DeleteUser)
	}

	// authority := router.Group("authority")
	// {
	// 	authority.GET("/", controller.GetAuthorities)
	// 	authority.POST("/", controller.CreateAuthority)
	// }

	return router
}
