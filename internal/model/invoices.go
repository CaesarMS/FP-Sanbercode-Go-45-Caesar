package model

import (
	"time"
)

type Invoices struct {
	Id          string    `gorm:"primaryKey" json:"id"`
	User_id     string    `gorm:"index" json:"user_id"`
	User        Users     `gorm:"foreignKey:User_id" json:"-"`
	Total_price uint      `gorm:"not null" json:"total_price"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
