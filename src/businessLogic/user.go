package businessLogic

import (
	"fmt"
	menty "github.com/6f-fiber-group-projects/6fg-app-api/entity/model_entity"
	reqenty "github.com/6f-fiber-group-projects/6fg-app-api/entity/request_entity"
	repo "github.com/6f-fiber-group-projects/6fg-app-api/repository"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) (*[]menty.User, error) {
	if !IsAdmin(c) {
		return nil, fmt.Errorf("unauthorized")
	}

	users, err := repo.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}

	return &users, nil
}

func GetUserById(c *gin.Context, userId int) (*menty.User, error) {
	if !IsAdmin(c) && !IsSameUser(c, userId) {
		return nil, fmt.Errorf("unauthorized")
	}

	user, err := repo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(c *gin.Context, u *reqenty.UserUpdateRequest) (*menty.User, error) {
	if !IsAdmin(c) && !IsSameUser(c, u.Id) {
		return nil, fmt.Errorf("unauthorized")
	}

	user, err := repo.GetUserById(u.Id)
	if err != nil {
		return nil, err
	}

	// not admin users can't change their athority
	if !IsAdmin(c) && u.Authority_id != 0 && user.Authority_id != u.Authority_id {
		return nil, fmt.Errorf("unauthorized")
	}

	user, err = repo.UpdateUser(u)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func DeleteUser(c *gin.Context, userId int) (*menty.User, error) {
	if !IsAdmin(c) && !IsSameUser(c, userId) {
		return nil, fmt.Errorf("unauthorized")
	}

	user, err := repo.DeleteUser(userId)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
