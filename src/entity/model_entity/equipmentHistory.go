package model_entity

import (
	"time"
)

type EquipmentHistory struct {
	Id            int
	EquipId       int
	UserId        int
	ReservationId *int
	StartDate     time.Time
	EndDate       time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
