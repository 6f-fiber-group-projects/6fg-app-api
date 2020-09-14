package model_entity

import (
	"time"
)

// type name should be single form of table name
type Equipment struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// set the table name
func (Equipment) TableName() string {
	return "equipments"
}
