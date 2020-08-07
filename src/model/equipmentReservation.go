package model

import (
	menty "github.com/6f-fiber-group-projects/6fg-app-api/entity/model_entity"
)

func GetEquipmentReservationById(rsvnId int) (menty.EquipmentReservation, error) {
	db := gormConnect()
	defer db.Close()

	rsvn := menty.EquipmentReservation{}
	result := db.Find(&rsvn, "id=?", rsvnId)
	return rsvn, result.Error
}

func GetEquipmentReservationByEquipId(equipId int) ([]menty.EquipmentReservation, error) {
	db := gormConnect()
	defer db.Close()

	rsvn := []menty.EquipmentReservation{}
	result := db.Order("id desc").Find(&rsvn, "equip_id=?", equipId).Limit(100)
	return rsvn, result.Error
}

func CreateEquipmentReservation(rsvn *menty.EquipmentReservation) (menty.EquipmentReservation, error) {
	db := gormConnect()
	defer db.Close()

	result := db.Create(rsvn)
	return *rsvn, result.Error
}

func UpdateEquipmentReservation(r *menty.EquipmentReservation) (menty.EquipmentReservation, error) {
	db := gormConnect()
	defer db.Close()

	rsvn := menty.EquipmentReservation{}
	result := db.Model(&rsvn).Where("id=?", r.Id).Omit("id", "created_at").Updates(r)
	return rsvn, result.Error
}

func DeleteEquipmentReservation(rsvn *menty.EquipmentReservation) (menty.EquipmentReservation, error) {
	db := gormConnect()
	defer db.Close()

	result := db.Delete(&rsvn)
	return *rsvn, result.Error
}
