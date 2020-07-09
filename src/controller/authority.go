package controller

import (
	menty "6fg-app-api/entity/model_entity"
	resenty "6fg-app-api/entity/response_entity"
	repo "6fg-app-api/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

// /authority
func GetAuthorities(c *gin.Context) {
	auths, err := repo.GetAllAuthorities()
	if err != nil {
		ResponseErrorMessage(c, "#QHPNLFVA", "No authority")
		return
	}

	formatedAuthorities := []resenty.AuthorityResponse{}
	for _, auth := range auths {
		formatedAuthorities = append(formatedAuthorities, formatAuthorityResponse(auth))
	}

	c.JSON(http.StatusOK, gin.H{"message": formatedAuthorities})
}

func formatAuthorityResponse(a menty.Authority) resenty.AuthorityResponse {
	return resenty.AuthorityResponse{
		Id:   a.Id,
		Name: a.Name,
	}
}
