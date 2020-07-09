package repository

import (
	menty "6fg-app-api/entity/model_entity"
	// reqenty "6fg-app-api/entity/request_entity"
	model "6fg-app-api/model"
	"fmt"
	// "golang.org/x/crypto/bcrypt"
	// "time"
)

func GetAllAuthorities() ([]menty.Authority, error) {
	users, err := model.GetAllAuthorities()
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return users, nil
}
