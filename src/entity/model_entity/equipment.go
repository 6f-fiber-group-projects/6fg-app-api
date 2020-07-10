package model_entity

import (
	"time"
)

// type name should be single form of table name
type Equipment struct {
	Id         int
	Name       string
	Status     int
	Created_at time.Time
	Updated_at time.Time
}

// set the table name
func (Equipment) TableName() string {
	return "equipments"
}
