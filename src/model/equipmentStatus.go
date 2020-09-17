package model

import (
	menty "github.com/6f-fiber-group-projects/6fg-app-api/entity/model_entity"
)

// func GetEquipmentStatusByEquipId(equipId int) (menty.EquipmentStatus, error) {
// 	db := gormConnect()
// 	defer db.Close()

// 	equipStatus := menty.EquipmentStatus{}
// 	result := db.Find(&equipStatus, "equip_id=?", equipId)
// 	return equipStatus, result.Error
// }

func CreateEquipStatus(equipStatus *menty.EquipmentStatus) (menty.EquipmentStatus, error) {
	db := gormConnect()
	defer db.Close()

	result := db.Create(equipStatus)
	return *equipStatus, result.Error
}

// func UpdateEquipStatus(e *menty.EquipmentStatus) (menty.EquipmentStatus, error) {
// 	db := gormConnect()
// 	defer db.Close()

// 	equipStatus := menty.EquipmentStatus{}
// 	result := db.Model(&equipStatus).Where("eqip_id=?", e.EquipId).Omit("eqip_id", "created_at").Updates(e)
// 	return equipStatus, result.Error
// }
