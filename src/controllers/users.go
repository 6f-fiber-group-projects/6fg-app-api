package controllers

import (
	renty "6fg-app-api/entities/request_entities"
	repo "6fg-app-api/repositories"
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Note that you need to set the corresponding binding tag on all fields you want to bind. For example, when binding from JSON, set json:"fieldname".
// cited from https://github.com/gin-gonic/gin#bind-form-data-request-with-custom-struct
//
// As part of GoLang, if either the var's or structs name start uncapitalized, it will be private, that's why neither Go or Gin don't let you access it content.
// cited from	https://github.com/gin-gonic/gin/issues/149#issuecomment-63179871

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "It works!",
	})
}

func CreateUser(c *gin.Context) {
	user := renty.UserRequest{}

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	err = repo.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
}
