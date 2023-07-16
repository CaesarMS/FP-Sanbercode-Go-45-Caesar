package config

import (
	"fmt"

	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/internal/model"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/pkg/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	username := util.GetEnv("DB_USERNAME", "root")
	password := util.GetEnv("DB_PASSWORD", "root")
	host := util.GetEnv("DB_HOST", "127.0.0.1")
	port := util.GetEnv("DB_PORT", "3306")
	database := util.GetEnv("DB_NAME", "sanbercode_final-project")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(
		&model.Admins{},
		&model.Categories{},
		&model.Invoice_items{},
		&model.Invoices{},
		&model.Products{},
		&model.Users{},
	)

	return db
}
