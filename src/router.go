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

	authority := router.Group("authority")
	{
		authority.GET("/", controller.GetAuthorities)
	}

	auth := router.Group("auth")
	{
		auth.POST("/", controller.BasicAuthenticate)
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
