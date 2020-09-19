package repository

import (
	"fmt"
	menty "github.com/6f-fiber-group-projects/6fg-app-api/entity/model_entity"
	reqenty "github.com/6f-fiber-group-projects/6fg-app-api/entity/request_entity"
	model "github.com/6f-fiber-group-projects/6fg-app-api/model"
	"time"
)

func GetEquipmentStatusByEquipId(equipId int) (menty.EquipmentStatus, error) {
	equipStauts, err := model.GetEquipmentStatusByEquipId(equipId)
	if err != nil {
		return menty.EquipmentStatus{}, fmt.Errorf("%s", err)
	}

	if result := equipStauts.NilTozero(); result != nil {
		equipStauts = *result
	}

	return equipStauts, nil
}

func CreateEquipmentStatus(equipId int) (menty.EquipmentStatus, error) {
	zero := 0
	equipStatus := menty.EquipmentStatus{
		EquipId:   equipId,
		Status:    &zero,
		UserId:    &zero,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	equipStatus, err := model.CreateEquipStatus(&equipStatus)
	if err != nil {
		return menty.EquipmentStatus{}, fmt.Errorf("%s", err)
	}
	return equipStatus, nil
}

func UpdateEquipmentStatus(e *reqenty.EquipmentStatusUpdateRequest) (menty.EquipmentStatus, error) {
	equipStatus := menty.EquipmentStatus{
		EquipId:   e.EquipId,
		Status:    e.Status,
		UserId:    e.UserId,
		UpdatedAt: time.Now(),
	}
	equipStatus, err := model.UpdateEquipStatus(&equipStatus)
	if err != nil {
		return menty.EquipmentStatus{}, fmt.Errorf("%s", err)
	}
	return equipStatus, nil
}

// func equipReqToModel(e *reqenty.EquipmentRequest) menty.Equipment {
// 	zero := 0
// 	return menty.Equipment{
// 		Name:      e.Name,
// 		Status:    &zero,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}
// }

// func equipUpdateReqToModel(e *reqenty.EquipmentUpdateRequest) menty.Equipment {
// 	return menty.Equipment{
// 		Id:        e.Id,
// 		Name:      e.Name,
// 		Status:    &e.Status,
// 		UpdatedAt: time.Now(),
// 	}
// }
