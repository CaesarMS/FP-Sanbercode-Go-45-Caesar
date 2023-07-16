package model

import (
	"time"
)

type Admins struct {
	Id         string    `gorm:"primaryKey" json:"id"`
	Username   string    `gorm:"not null;unique" json:"username"`
	Password   string    `gorm:"not null" json:"password"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
