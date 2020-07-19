package repository

import (
	menty "github.com/6f-fiber-group-projects/6fg-app-api/entity/model_entity"
	model "github.com/6f-fiber-group-projects/6fg-app-api/model"
	"fmt"
)

func GetAllAuthorities() ([]menty.Authority, error) {
	users, err := model.GetAllAuthorities()
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return users, nil
}
