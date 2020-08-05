package main

import (
	"github.com/6f-fiber-group-projects/6fg-app-api/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func DefineRoutes() *gin.Engine {
	router := gin.Default()

	// middle ware
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))
	router.Use(sessionCheck())

	user := router.Group("user")
	{
		user.GET("/", controller.GetUsers)
		user.POST("/", controller.CreateUser)
		user.GET("/:userId", controller.GetUserById)
		user.PUT("/:userId", controller.UpdateUser)
		user.DELETE("/:userId", controller.DeleteUser)
	}

	authority := router.Group("authority")
	{
		authority.GET("/", controller.GetAuthorities)
	}

	auth := router.Group("auth")
	{
		auth.POST("/", controller.BasicAuthenticate)
		auth.GET("/logout", controller.Logout)
	}

	equip := router.Group("equipment")
	{
		equip.GET("/", controller.GetEquipments)
		equip.POST("/", controller.CreateEquipment)
		equip.GET("/:equipId", controller.GetEquipmentById)
		equip.PUT("/:equipId", controller.UpdateEquipment)
		equip.DELETE("/:equipId", controller.DeleteEquipment)
		equip.GET("/:equipId/qrcode", controller.GetEquipmentQRcode)
		equip.POST("/:equipId/status", controller.UpdateEquipmentStatus)
	}

	rsvn := router.Group("reservation")
	{
		rsvn.POST("equipment", controller.CreateEquipmentReservation)
		rsvn.GET("equipment/:rsvnId", controller.GetEquipmentReservation)
		rsvn.PUT("equipment/:rsvnId", controller.UpdateEquipmentReservation)
		rsvn.DELETE("equipment/:rsvnId", controller.DeleteEquipmentReservation)
	}

	return router
}

func sessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		isLogin := session.Get("isLogin")
		if isLogin != true && c.FullPath() != "/auth/" {
			controller.ResponseUnauthorizedMessage(c)
			c.Abort()
		}
		c.Next()
	}
}
