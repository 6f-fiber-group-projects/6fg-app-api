package model_entity

import (
	"time"
)

type EquipmentReservation struct {
	Id        int
	EquipId   int
	UserId    int
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
