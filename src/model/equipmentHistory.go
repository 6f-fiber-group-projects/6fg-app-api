package model

import (
	menty "6fg-app-api/entity/model_entity"
)

func CreateEquipHistory(equipHistroy *menty.EquipmentHistory) (menty.EquipmentHistory, error) {
	db := gormConnect()
	defer db.Close()

	result := db.Create(equipHistroy)
	return *equipHistroy, result.Error
}

func UpdateEquipHistoryByEquipId(e *menty.EquipmentHistory) (menty.EquipmentHistory, error) {
	db := gormConnect()
	defer db.Close()

	equipHistroy := menty.EquipmentHistory{}
	result := db.Where("equip_id = ?", e.EquipId).Last(&equipHistroy).Updates(e)
	return equipHistroy, result.Error
}
