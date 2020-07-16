package model

import (
	menty "6fg-app-api/entity/model_entity"
)

func GetAllEquips() ([]menty.Equipment, error) {
	db := gormConnect()
	defer db.Close()

	equips := []menty.Equipment{}
	result := db.Find(&equips)
	return equips, result.Error
}

func GetEquipmentById(equipId int) (menty.Equipment, error) {
	db := gormConnect()
	defer db.Close()

	equip := menty.Equipment{}
	result := db.Find(&equip, "id=?", equipId)
	return equip, result.Error
}

func CreateEquip(equip *menty.Equipment) (menty.Equipment, error) {
	db := gormConnect()
	defer db.Close()

	result := db.Create(equip)
	return *equip, result.Error
}

func UpdateEquip(e *menty.Equipment) (menty.Equipment, error) {
	db := gormConnect()
	defer db.Close()

	equip := menty.Equipment{}
	result := db.Model(&equip).Where("id=?", e.Id).Omit("id", "created_at").Updates(e)
	return equip, result.Error
}

func DeleteEquip(equip *menty.Equipment) (menty.Equipment, error) {
	db := gormConnect()
	defer db.Close()

	result := db.Delete(equip)
	return *equip, result.Error
}
