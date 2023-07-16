package model

import (
	"time"
)

type Categories struct {
	Id         string    `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"not null;unique" json:"name"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
