package model

import (
	"time"
)

type Products struct {
	Id          string     `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"not null" json:"name"`
	Price       uint       `gorm:"not null" json:"price"`
	Description string     `gorm:"not null" json:"description"`
	Stock       uint       `gorm:"not null" json:"stock"`
	Category_id string     `gorm:"index" json:"category_id"`
	Category    Categories `gorm:"foreignKey:Category_id" json:"-"`
	User_id     string     `gorm:"index" json:"user_id"`
	User        Users      `gorm:"foreignKey:User_id" json:"-"`
	Created_at  time.Time  `json:"created_at"`
	Updated_at  time.Time  `json:"updated_at"`
}
