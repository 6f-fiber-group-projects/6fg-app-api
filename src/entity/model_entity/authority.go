package model_entity

import (
	"time"
)

// type name should be single form of table name
type Authority struct {
	Id         int
	Name       string
	Created_at time.Time
	Updated_at time.Time
}
