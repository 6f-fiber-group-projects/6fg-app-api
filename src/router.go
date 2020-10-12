package main

import (
	"github.com/6f-fiber-group-projects/6fg-app-api/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"os"
)

func DefineRoutes() *gin.Engine {
	router := gin.Default()

	// middle ware
	// CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("ALLOW_ORIGIN")}
	config.AllowCredentials = true
	// config.AllowHeaders = []string{"Set-Cookie"}
	router.Use(cors.New(config))

	// session
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))
	router.Use(sessionCheck())

	// http://shinriyo.hateblo.jp/entry/2017/09/18/gin%E3%81%A7Group%E5%8C%96%E3%81%97%E3%81%9F%E3%83%AB%E3%83%BC%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0
	user := router.Group("user")
	{
		user.GET("", controller.GetUsers)
		user.POST("", controller.CreateUser)
		user.GET("/:userId", controller.GetUserById)
		user.PUT("/:userId", controller.UpdateUser)
		user.DELETE("/:userId", controller.DeleteUser)
	}

	authority := router.Group("authority")
	{
		authority.GET("", controller.GetAuthorities)
	}

	auth := router.Group("auth")
	{
		auth.POST("", controller.BasicAuthenticate)
		auth.GET("/logout", controller.Logout)
	}

	equip := router.Group("equipment")
	{
		equip.GET("", controller.GetEquipments)
		equip.POST("", controller.CreateEquipment)
		equip.GET("/:equipId", controller.GetEquipmentById)
		equip.PUT("/:equipId", controller.UpdateEquipment)
		equip.DELETE("/:equipId", controller.DeleteEquipment)
		equip.GET("/:equipId/qrcode", controller.GetEquipmentQRcode)
		equip.POST("/:equipId/status", controller.UpdateEquipmentStatus)
	}

	rsvn := router.Group("reservation")
	{
		rsvn.GET("equipment", controller.GetEquipmentReservationByEquipId)
		rsvn.POST("equipment", controller.CreateEquipmentReservation)
		rsvn.GET("equipment/:rsvnId", controller.GetEquipmentReservationById)
		rsvn.PUT("equipment/:rsvnId", controller.UpdateEquipmentReservation)
		rsvn.DELETE("equipment/:rsvnId", controller.DeleteEquipmentReservation)
	}

	return router
}

func sessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		isLogin := session.Get("isLogin")
		if isLogin != true && c.FullPath() != "/auth" {
			controller.ResponseUnauthorizedMessage(c)
			c.Abort()
		}
		c.Next()
	}
}
