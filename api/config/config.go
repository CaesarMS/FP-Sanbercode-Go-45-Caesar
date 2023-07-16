package config

import (
	"fmt"

	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/internal/model"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/pkg/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	db_username, err := util.GetEnv("DB_USERNAME")
	if err != nil {
		panic(err.Error())
	}
	username := db_username

	db_password, err := util.GetEnv("DB_PASSWORD")
	if err != nil {
		panic(err.Error())
	}
	password := db_password

	db_host, err := util.GetEnv("DB_HOST")
	if err != nil {
		panic(err.Error())
	}
	host := "tcp(" + db_host + ")"

	db_name, err := util.GetEnv("DB_NAME")
	if err != nil {
		panic(err.Error())
	}
	database := db_name

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

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
