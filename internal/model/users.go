package model

import (
	"time"
)

type Users struct {
	Id         string    `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"not null" json:"name"`
	Email      string    `gorm:"not null;unique" json:"email"`
	Password   string    `gorm:"not null" json:"password"`
	Address    string    `gorm:"not null" json:"address"`
	Is_seller  bool      `json:"is_seller"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
