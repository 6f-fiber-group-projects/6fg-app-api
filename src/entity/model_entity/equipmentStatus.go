package model_entity

import (
	"time"
)

type EquipmentStatus struct {
	EquipId   int
	UserId    *int
	Status    *int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (EquipmentStatus) TableName() string {
	return "equipments_status"
}

func (e *EquipmentStatus) NilTozero() *EquipmentStatus {
	zero := 0
	status := e.Status
	userId := e.UserId
	if e.Status == nil {
		status = &zero
	}
	if e.UserId == nil {
		userId = &zero
	}
	return &EquipmentStatus{
		EquipId:   e.EquipId,
		UserId:    userId,
		Status:    status,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
