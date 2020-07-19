package controller

import (
	menty "github.com/6f-fiber-group-projects/6fg-app-api/entity/model_entity"
	resenty "github.com/6f-fiber-group-projects/6fg-app-api/entity/response_entity"
	"fmt"
)

// user
func formatUserResponse(u menty.User) resenty.UserResponse {
	return resenty.UserResponse{
		Id:           u.Id,
		Authority_id: u.Authority_id,
		Google_id:    u.Google_id,
		Name:         u.Name,
		Email:        u.Email,
	}
}

// equipment
func formatEquipmentResponse(e menty.Equipment) resenty.EquipmentResponse {
	fmt.Printf("%#v", e.Status)
	return resenty.EquipmentResponse{
		Id:     e.Id,
		Name:   e.Name,
		Status: *e.Status,
	}
}

//authority
func formatAuthorityResponse(a menty.Authority) resenty.AuthorityResponse {
	return resenty.AuthorityResponse{
		Id:   a.Id,
		Name: a.Name,
	}
}
