package controller

import (
	bl "github.com/6f-fiber-group-projects/6fg-app-api/businessLogic"
	reqenty "github.com/6f-fiber-group-projects/6fg-app-api/entity/request_entity"
	repo "github.com/6f-fiber-group-projects/6fg-app-api/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BasicAuthenticate(c *gin.Context) {
	auth := reqenty.AuthRequest{}
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		ResponseServerErrorMessage(c, "#A3H884KU", "Bad request")
		return
	}

	user, err := repo.BasicAuthenticate(&auth)
	if err != nil {
		ResponseServerErrorMessage(c, "#G9GSLGOH", err.Error())
		return
	}

	bl.CreateSession(c, &user)

	c.JSON(http.StatusAccepted, gin.H{"message": "User logined successfully"})
}

func Logout(c *gin.Context) {
	bl.KillSession(c)
	c.JSON(http.StatusAccepted, gin.H{"message": "session was cleared"})
}
