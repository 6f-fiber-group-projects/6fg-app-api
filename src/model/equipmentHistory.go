package model

import (
	menty "6fg-app-api/entity/model_entity"
)

func GetLatestEquipmentHistoryByEquipId(equipId int) (menty.EquipmentHistory, error) {
	db := gormConnect()
	defer db.Close()

	equipHistroy := menty.EquipmentHistory{}
	result := db.Where("equip_id = ?", equipId).Last(&equipHistroy)
	return equipHistroy, result.Error
}

func CreateEquipHistory(equipHistroy *menty.EquipmentHistory) (menty.EquipmentHistory, error) {
	db := gormConnect()
	defer db.Close()

	result := db.Create(equipHistroy)
	return *equipHistroy, result.Error
}

func UpdateEquipHistory(e *menty.EquipmentHistory) (menty.EquipmentHistory, error) {
	db := gormConnect()
	defer db.Close()

	equipHistroy := menty.EquipmentHistory{}
	result := db.Model(&equipHistroy).Where("id = ?", e.Id).Omit("id", "created_at").Updates(e)
	return equipHistroy, result.Error
}
