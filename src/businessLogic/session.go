package businessLogic

import (
	menty "github.com/6f-fiber-group-projects/6fg-app-api/entity/model_entity"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreateSession(c *gin.Context, u *menty.User) {
	session := sessions.Default(c)
	session.Set("userId", u.Id)
	session.Set("authId", u.Authority_id)
	session.Set("isLogin", true)

	options := sessions.Options{
		MaxAge: 60 * 60 * 24,
		Path:   "/",
	}
	session.Options(options)

	session.Save()
}

func KillSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()

	options := sessions.Options{
		MaxAge: -1,
		Path:   "/",
	}
	session.Options(options)
	session.Save()
}
