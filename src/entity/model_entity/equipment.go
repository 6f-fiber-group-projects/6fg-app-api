package model_entity

import (
	"time"
)

// type name should be single form of table name
type Equipment struct {
	Id        int
	Name      string
	Status    *int // for gorm. 0 will be ignored.
	CreatedAt time.Time
	UpdatedAt time.Time
}

// set the table name
func (Equipment) TableName() string {
	return "equipments"
}
