package repository

import (
	menty "6fg-app-api/entity/model_entity"
	model "6fg-app-api/model"
	"fmt"
)

func GetAllAuthorities() ([]menty.Authority, error) {
	users, err := model.GetAllAuthorities()
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return users, nil
}
