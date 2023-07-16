package model

import (
	"time"
)

type Invoice_items struct {
	Id         string    `gorm:"primaryKey" json:"id"`
	Invoice_id string    `gorm:"index" json:"invoice_id"`
	Invoice    Invoices  `gorm:"foreignKey:Invoice_id"  json:"-"`
	Product_id string    `gorm:"index" json:"product_id"`
	Product    Products  `gorm:"foreignKey:Product_id" json:"-"`
	Qty        uint      `gorm:"not null" json:"qty"`
	Price      uint      `gorm:"not null" json:"price"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
