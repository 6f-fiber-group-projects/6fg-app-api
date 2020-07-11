package model_entity

import (
	"time"
)

// type name should be single form of table name
type User struct {
	Id           int
	Authority_id int
	Google_id    int
	Name         string
	Email        string
	Password     []byte
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
