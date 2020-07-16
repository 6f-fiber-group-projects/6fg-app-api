package repository

import (
	menty "6fg-app-api/entity/model_entity"
	reqenty "6fg-app-api/entity/request_entity"
	model "6fg-app-api/model"
	"fmt"
	"time"
)

func CreateEquipmentHistory(e *reqenty.EquipmentHistoryRequest) (menty.EquipmentHistory, error) {
	equipHistory, err := equipHistoryReqToModel(e)
	fmt.Printf("%#v", equipHistory)
	_, err = model.CreateEquipHistory(&equipHistory)
	if err != nil {
		return menty.EquipmentHistory{}, fmt.Errorf("%s", err)
	}
	return equipHistory, nil
}

func UpdateEquipmentHistory(e *reqenty.EquipmentHistoryRequest) (menty.EquipmentHistory, error) {
	equipHistory, err := equipHistoryUpdateReqToModel(e)
	fmt.Printf("%#v", equipHistory)
	_, err = model.UpdateEquipHistoryByEquipId(&equipHistory)
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

func equipHistoryUpdateReqToModel(e *reqenty.EquipmentHistoryRequest) (menty.EquipmentHistory, error) {
	_, endDate, err := parseTime(e.StartDate, e.EndDate)
	if err != nil {
		endDate = time.Now()
	}
	return menty.EquipmentHistory{
		EquipId:       e.EquipId,
		UserId:        e.UserId,
		ReservationId: e.ReservationId,
		EndDate:       endDate,
		UpdatedAt:     time.Now(),
	}, nil
}
