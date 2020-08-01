package controller

import (
	menty "github.com/6f-fiber-group-projects/6fg-app-api/entity/model_entity"
	reqenty "github.com/6f-fiber-group-projects/6fg-app-api/entity/request_entity"
	repo "github.com/6f-fiber-group-projects/6fg-app-api/repository"
	"github.com/gin-contrib/sessions"
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

	user, err := repo.BasicAuthenticate(&auth)
	if err != nil {
		ResponseErrorMessage(c, "#G9GSLGOH", err.Error())
		return
	}

	createSession(c, &user)

	c.JSON(http.StatusAccepted, gin.H{"message": "User logined successfully"})
}

func Logout(c *gin.Context) {
	killSession(c)
	c.JSON(http.StatusAccepted, gin.H{"message": "session was cleared"})
}

func createSession(c *gin.Context, u *menty.User) {
	session := sessions.Default(c)
	session.Set("userId", u.Id)
	session.Set("auth", u.Authority_id)
	session.Set("isLogin", true)

	options := sessions.Options{
		MaxAge: 60 * 60 * 24,
		Path:   "/",
	}
	session.Options(options)

	session.Save()
}

func killSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()

	options := sessions.Options{
		MaxAge: -1,
		Path:   "/",
	}
	session.Options(options)
	session.Save()
}
