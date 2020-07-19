package repository

import (
	menty "6fg-app-api/entity/model_entity"
	reqenty "6fg-app-api/entity/request_entity"
	model "6fg-app-api/model"
	"fmt"
	"time"
)

func GetAllEquipments() ([]menty.Equipment, error) {
	equips, err := model.GetAllEquips()
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	for idx, equip := range equips {
		if result := equip.StatusNilTozero(); result != nil {
			equips[idx] = *result
		}
	}
	return equips, nil
}

func GetEquipmentById(equipId int) (menty.Equipment, error) {
	equip, err := model.GetEquipmentById(equipId)
	if err != nil {
		return menty.Equipment{}, fmt.Errorf("%s", err)
	}

	if result := equip.StatusNilTozero(); result != nil {
		equip = *result
	}

	return equip, nil
}

func CreateEquipment(e *reqenty.EquipmentRequest) (menty.Equipment, error) {
	equip := equipReqToModel(e)
	_, err := model.CreateEquip(&equip)
	if err != nil {
		return menty.Equipment{}, fmt.Errorf("%s", err)
	}
	return menty.Equipment{}, nil
}

func UpdateEquipment(e *reqenty.EquipmentUpdateRequest) (menty.Equipment, error) {
	equip := equipUpdateReqToModel(e)

	_, err := model.UpdateEquip(&equip)
	if err != nil {
		return menty.Equipment{}, fmt.Errorf("%s", err)
	}
	return menty.Equipment{}, nil
}

func DeleteEquipment(equipId int) (menty.Equipment, error) {
	equip, err := model.GetEquipmentById(equipId)
	if err != nil {
		return menty.Equipment{}, fmt.Errorf("%s", err)
	}
	_, err = model.DeleteEquip(&equip)
	return menty.Equipment{}, nil
}

func equipReqToModel(e *reqenty.EquipmentRequest) menty.Equipment {
	zero := 0
	return menty.Equipment{
		Name:      e.Name,
		Status:    &zero,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func equipUpdateReqToModel(e *reqenty.EquipmentUpdateRequest) menty.Equipment {
	return menty.Equipment{
		Id:        e.Id,
		Name:      e.Name,
		Status:    &e.Status,
		UpdatedAt: time.Now(),
	}
}
