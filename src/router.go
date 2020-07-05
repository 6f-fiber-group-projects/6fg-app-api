package main

import (
	"6fg-app-api/controllers"
	"github.com/gin-gonic/gin"
)

func DefineRoutes() *gin.Engine {
	router := gin.Default()

	// user
	router.GET("users", controllers.GetUsers)
	router.POST("users", controllers.CreateUser)

	return router
}
