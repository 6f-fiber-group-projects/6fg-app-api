package main

import (
	"6fg-app-api/controllers"
	"github.com/gin-gonic/gin"
)

func DefineRoutes() *gin.Engine {
	router := gin.Default()

	users := router.Group("users")
	{
		users.GET("/", controllers.GetUsers)
		users.POST("/", controllers.CreateUser)
		users.GET("/:userId", controllers.GetUserById)
		users.PUT("/:userId", controllers.UpdateUser)
		users.DELETE("/:userId", controllers.DeleteUser)
	}

	return router
}
