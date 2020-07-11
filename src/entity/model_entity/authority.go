package model_entity

import (
	"time"
)

// type name should be single form of table name
type Authority struct {
	Id         int
	Name       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
