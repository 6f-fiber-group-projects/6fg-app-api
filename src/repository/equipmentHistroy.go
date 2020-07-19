package repository

import (
	menty "6fg-app-api/entity/model_entity"
	reqenty "6fg-app-api/entity/request_entity"
	model "6fg-app-api/model"
	"fmt"
	"time"
)

func GetLatestEquipmentHistoryByEquipId(equipId int) (menty.EquipmentHistory, error) {
	equipHistory, err := model.GetLatestEquipmentHistoryByEquipId(equipId)
	if err != nil {
		return menty.EquipmentHistory{}, fmt.Errorf("%s", err)
	}
	return equipHistory, nil
}

func CreateEquipmentHistory(e *reqenty.EquipmentHistoryRequest) (menty.EquipmentHistory, error) {
	equipHistory, err := equipHistoryReqToModel(e)
	_, err = model.CreateEquipHistory(&equipHistory)
	if err != nil {
		return menty.EquipmentHistory{}, fmt.Errorf("%s", err)
	}
	return equipHistory, nil
}

func UpdateEquipmentHistory(e *reqenty.EquipmentHistoryUpdateRequest) (menty.EquipmentHistory, error) {
	equipHistory, err := equipHistoryUpdateReqToModel(e)
	_, err = model.UpdateEquipHistory(&equipHistory)
	if err != nil {
		return menty.EquipmentHistory{}, fmt.Errorf("%s", err)
	}
	return equipHistory, nil
}

func equipHistoryReqToModel(e *reqenty.EquipmentHistoryRequest) (menty.EquipmentHistory, error) {
	return menty.EquipmentHistory{
		EquipId:       e.EquipId,
		UserId:        e.UserId,
		ReservationId: e.ReservationId,
		StartDate:     time.Now(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}, nil
}

func equipHistoryUpdateReqToModel(e *reqenty.EquipmentHistoryUpdateRequest) (menty.EquipmentHistory, error) {
	return menty.EquipmentHistory{
		Id:            e.Id,
		EquipId:       e.EquipId,
		UserId:        e.UserId,
		ReservationId: e.ReservationId,
		EndDate:       time.Now(),
		UpdatedAt:     time.Now(),
	}, nil
}
