package businessLogic

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func IsAdmin(c *gin.Context) bool {
	session := sessions.Default(c)
	authId := session.Get("authId")
	adminId, _ := strconv.Atoi(os.Getenv("ADMIN_ID"))
	if authId == adminId {
		return true
	}
	return false
}

func IsSameUser(c *gin.Context, userId int) bool {
	session := sessions.Default(c)
	if sessionUserId := session.Get("userId"); userId == sessionUserId {
		return true
	}
	return false
}
