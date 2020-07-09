package controller

import (
	reqenty "6fg-app-api/entity/request_entity"
	repo "6fg-app-api/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BasicAuthenticate(c *gin.Context) {
	auth := reqenty.AuthRequest{}
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		ResponseErrorMessage(c, "#A3H884KU", "Bad request")
		return
	}

	_, err = repo.BasicAuthenticate(&auth)
	if err != nil {
		ResponseErrorMessage(c, "#G9GSLGOH", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "User logined successfully"})
}
