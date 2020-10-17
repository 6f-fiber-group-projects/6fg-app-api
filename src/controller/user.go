package controller

import (
	bl "github.com/6f-fiber-group-projects/6fg-app-api/businessLogic"
	reqenty "github.com/6f-fiber-group-projects/6fg-app-api/entity/request_entity"
	resenty "github.com/6f-fiber-group-projects/6fg-app-api/entity/response_entity"
	repo "github.com/6f-fiber-group-projects/6fg-app-api/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

// Note that you need to set the corresponding binding tag on all fields you want to bind. For example, when binding from JSON, set json:"fieldname".
// cited from https://github.com/gin-gonic/gin#bind-form-data-request-with-custom-struct
//
// As part of GoLang, if either the var's or structs name start uncapitalized, it will be private, that's why neither Go or Gin don't let you access it content.
// cited from	https://github.com/gin-gonic/gin/issues/149#issuecomment-63179871

// /user
func GetUsers(c *gin.Context) {
	users, err := bl.GetAllUsers(c)
	if err != nil {
		ResponseServerErrorMessage(c, "#S3AB7J44", err.Error())
		return
	}

	formatedUsers := []resenty.UserResponse{}
	for _, user := range *users {
		formatedUsers = append(formatedUsers, formatUserResponse(user))
	}

	c.JSON(http.StatusOK, gin.H{"message": formatedUsers})
}

func CreateUser(c *gin.Context) {
	user := reqenty.UserRequest{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		ResponseServerErrorMessage(c, "#A3H884KU", "Bad request")
		return
	}
	user.Authority_id, _ = strconv.Atoi(os.Getenv("VIEWER_ID"))

	_, err = repo.CreateUser(&user)
	if err != nil {
		ResponseServerErrorMessage(c, "#G9GSLGOH", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
}

// user/:id
func GetUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		ResponseServerErrorMessage(c, "#HQFJFHKK", "User id should be integer")
		return
	}

	user, err := bl.GetUserById(c, userId)
	if err != nil {
		ResponseServerErrorMessage(c, "#COIO4KWD", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": formatUserResponse(*user)})
}

func UpdateUser(c *gin.Context) {
	user := reqenty.UserUpdateRequest{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		ResponseServerErrorMessage(c, "#K3C1P0FT", "Bad request")
		return
	}

	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		ResponseServerErrorMessage(c, "#257X17LY", "User id should be integer")
		return
	}
	user.Id = userId

	_, err = bl.UpdateUser(c, &user)
	if err != nil {
		ResponseServerErrorMessage(c, "#COIO4KWD", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": user})
}

func DeleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		ResponseServerErrorMessage(c, "#YFJ4Z6V9", "User id should be integer")
		return
	}

	_, err = bl.DeleteUser(c, userId)
	if err != nil {
		ResponseServerErrorMessage(c, "#6L2AO4MR", err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{})
}
