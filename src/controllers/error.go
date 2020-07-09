package controllers

import (
	lib "6fg-app-api/libs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseErrorMessage(c *gin.Context, hash, msg string) {
	fe := lib.FormatError(hash, msg)
	c.JSON(http.StatusInternalServerError, gin.H{"error": fe.Error()})
}
