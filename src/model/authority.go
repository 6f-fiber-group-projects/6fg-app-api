package model

import (
	menty "6fg-app-api/entity/model_entity"
	// "fmt"
)

func GetAllAuthorities() ([]menty.Authority, error) {
	db := gormConnect()
	defer db.Close()

	auths := []menty.Authority{}
	result := db.Find(&auths)
	return auths, result.Error
}
