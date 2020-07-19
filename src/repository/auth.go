package repository

import (
	menty "github.com/6f-fiber-group-projects/6fg-app-api/entity/model_entity"
	reqenty "github.com/6f-fiber-group-projects/6fg-app-api/entity/request_entity"
	"github.com/6f-fiber-group-projects/6fg-app-api/model"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func BasicAuthenticate(req *reqenty.AuthRequest) (menty.User, error) {
	user, err := model.GetUserByEmail(req.Email)
	if err != nil {
		return menty.User{}, fmt.Errorf("%s", err)
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password))
	if err != nil {
		return menty.User{}, fmt.Errorf("Password is incorrect")
	}

	return user, nil
}
