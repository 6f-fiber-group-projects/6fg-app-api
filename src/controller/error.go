package controller

import (
	"github.com/6f-fiber-group-projects/6fg-app-api/lib"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseServerErrorMessage(c *gin.Context, hash, msg string) {
	fe := lib.FormatError(hash, msg)
	c.JSON(http.StatusInternalServerError, gin.H{"error": fe.Error()})
}

func ResponseUnauthorizedMessage(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
}
