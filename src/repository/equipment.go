package repository

import (
	"fmt"
	menty "github.com/6f-fiber-group-projects/6fg-app-api/entity/model_entity"
	reqenty "github.com/6f-fiber-group-projects/6fg-app-api/entity/request_entity"
	model "github.com/6f-fiber-group-projects/6fg-app-api/model"
	"time"
)

func GetAllEquipments() ([]menty.Equipment, error) {
	equips, err := model.GetAllEquips()
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return equips, nil
}

func GetEquipmentById(equipId int) (menty.Equipment, error) {
	equip, err := model.GetEquipmentById(equipId)
	if err != nil {
		return menty.Equipment{}, fmt.Errorf("%s", err)
	}
	return equip, nil
}

func CreateEquipment(e *reqenty.EquipmentRequest) (menty.Equipment, error) {
	equip := equipReqToModel(e)
	_, err := model.CreateEquip(&equip)
	if err != nil {
		return menty.Equipment{}, fmt.Errorf("%s", err)
	}
	return equip, nil
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
	return menty.Equipment{
		Name:      e.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func equipUpdateReqToModel(e *reqenty.EquipmentUpdateRequest) menty.Equipment {
	return menty.Equipment{
		Id:        e.Id,
		Name:      e.Name,
		UpdatedAt: time.Now(),
	}
}
