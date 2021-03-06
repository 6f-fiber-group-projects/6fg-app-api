package repository

import (
	"fmt"
	menty "github.com/6f-fiber-group-projects/6fg-app-api/entity/model_entity"
	reqenty "github.com/6f-fiber-group-projects/6fg-app-api/entity/request_entity"
	"github.com/6f-fiber-group-projects/6fg-app-api/model"
	"time"
)

const TIME_LAYOUT = time.RFC3339Nano

func GetEquipmentReservation() ([]menty.EquipmentReservation, error) {
	rsvns, err := model.GetEquipmentReservation()
	if err != nil {
		return nil, err
	}
	return rsvns, nil
}

func GetEquipmentReservationById(id int) (menty.EquipmentReservation, error) {
	rsvn, err := model.GetEquipmentReservationById(id)
	if err != nil {
		return menty.EquipmentReservation{}, err
	}
	return rsvn, nil
}

func GetEquipmentReservationByEquipId(equipId int) ([]menty.EquipmentReservation, error) {
	rsvns, err := model.GetEquipmentReservationByEquipId(equipId)
	if err != nil {
		return nil, err
	}
	return rsvns, nil
}

func CreateEquipmentReservation(r *reqenty.EquipmentReservationRequest) (menty.EquipmentReservation, error) {
	rsvn, err := equipRsvnReqToModel(r)
	if err != nil {
		return menty.EquipmentReservation{}, err
	}

	rsvn, err = model.CreateEquipmentReservation(&rsvn)
	if err != nil {
		return menty.EquipmentReservation{}, err
	}

	return rsvn, nil
}

func UpdateEquipmentReservation(r *reqenty.EquipmentReservationUpdateRequest) (menty.EquipmentReservation, error) {
	rsvn, err := equipRsvnUpdateReqToModel(r)
	if err != nil {
		return menty.EquipmentReservation{}, err
	}

	rsvn, err = model.UpdateEquipmentReservation(&rsvn)
	if err != nil {
		return menty.EquipmentReservation{}, err
	}

	return rsvn, nil
}

func DeleteEquipmentReservation(id int) (menty.EquipmentReservation, error) {
	rsvn, err := model.GetEquipmentReservationById(id)
	if err != nil {
		return menty.EquipmentReservation{}, fmt.Errorf("%s", err)
	}

	_, err = model.DeleteEquipmentReservation(&rsvn)
	if err != nil {
		return menty.EquipmentReservation{}, err
	}

	return rsvn, nil
}

func equipRsvnReqToModel(r *reqenty.EquipmentReservationRequest) (menty.EquipmentReservation, error) {
	startDate, endDate, err := parseTime(r.StartDate, r.EndDate)
	if err != nil {
		return menty.EquipmentReservation{}, err
	}
	return menty.EquipmentReservation{
		EquipId:   r.EquipId,
		UserId:    r.UserId,
		GroupId:   r.GroupId,
		StartDate: startDate,
		EndDate:   endDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func equipRsvnUpdateReqToModel(r *reqenty.EquipmentReservationUpdateRequest) (menty.EquipmentReservation, error) {
	startDate, endDate, err := parseTime(r.StartDate, r.EndDate)
	if err != nil {
		return menty.EquipmentReservation{}, err
	}
	return menty.EquipmentReservation{
		Id:        r.Id,
		EquipId:   r.EquipId,
		UserId:    r.UserId,
		GroupId:   r.GroupId,
		StartDate: startDate,
		EndDate:   endDate,
		UpdatedAt: time.Now(),
	}, nil
}

func parseTime(st, ed string) (time.Time, time.Time, error) {
	parsedSt, err := time.Parse(TIME_LAYOUT, st)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("%s", err)
	}
	parsedEd, err := time.Parse(TIME_LAYOUT, ed)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("%s", err)
	}
	return parsedSt, parsedEd, nil
}
