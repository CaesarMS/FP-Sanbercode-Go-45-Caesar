package service

import (
	"errors"
	"net/http"
	"time"

	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/api/config"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/internal/model"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/pkg/util"

	"gorm.io/gorm"
)

func GetAdmins(db *gorm.DB) (status int, adminLists []*config.AdminOutput, err error) {
	var admins []model.Admins
	db.Find(&admins)

	if len(admins) == 0 {
		status = http.StatusBadRequest
		err = errors.New("data not found")
		return
	}

	// take out password
	for _, admin := range admins {
		adminLists = append(adminLists, &config.AdminOutput{
			Id:       admin.Id,
			Username: admin.Username,
		})
	}

	status = http.StatusOK
	err = nil
	return
}

func GetAdminById(db *gorm.DB, id string) (status int, findAdmin *config.AdminOutput, err error) { // Get model if exist
	var admin model.Admins

	if errData := db.Where("id = ?", id).First(&admin).Error; errData != nil {
		status = http.StatusInternalServerError
		err = errors.New("data not found")
		return
	}

	status = http.StatusOK
	findAdmin = &config.AdminOutput{
		Id:       admin.Id,
		Username: admin.Username,
	}
	err = nil

	return
}

func GetAdminByUsername(db *gorm.DB, username string) (status int, findAdmin *config.AdminOutput, err error) {
	var admin model.Admins

	if errData := db.Where("username = ?", username).First(&admin).Error; errData != nil {
		status = http.StatusInternalServerError
		err = errors.New("data not found")
		return
	}

	status = http.StatusOK
	findAdmin = &config.AdminOutput{
		Id:       admin.Id,
		Username: admin.Username,
	}
	err = nil
	return
}

func GetAdminData(db *gorm.DB, username string) (status int, findAdmin model.Admins, err error) {
	if errData := db.Where("username = ?", username).First(&findAdmin).Error; errData != nil {
		status = http.StatusInternalServerError
		err = errors.New("data not found")
		return
	}

	status = http.StatusOK
	err = nil
	return
}

func InsertAdmin(
	db *gorm.DB,
	username, password string,
) (status int, newAdmin *config.AdminOutput, err error) {
	// Insert Admin
	admin := model.Admins{Id: util.GenerateUUID(), Username: username, Password: password, Created_at: time.Now(), Updated_at: time.Now()}
	db.Create(&admin)

	status, newAdmin, err = GetAdminById(db, admin.Id)
	if err != nil {
		return
	}

	status = http.StatusOK
	err = nil
	return
}

func UpdateAdminPassword(db *gorm.DB, id string, password string) (status int, updatedAdmin *config.AdminOutput, err error) {
	var admin model.Admins
	if errData := db.Where("id = ?", id).First(&admin).Error; errData != nil {
		status = http.StatusBadRequest
		err = errors.New("data not found")
		return
	}

	var updatedInput model.Admins
	updatedInput.Password = password
	updatedInput.Updated_at = time.Now()

	db.Model(&admin).Updates(updatedInput)

	status, updatedAdmin, err = GetAdminById(db, admin.Id)
	if err != nil {
		return
	}

	status = http.StatusOK
	err = nil
	return
}

func DeleteAdmin(db *gorm.DB, id string) (status int, deletedAdmin *config.AdminOutput, err error) {
	var admin model.Admins
	if errData := db.Where("id = ?", id).First(&admin).Error; errData != nil {
		status = http.StatusBadRequest
		err = errors.New("data not found")
		return
	}

	db.Delete(&admin)

	deletedAdmin = &config.AdminOutput{
		Id:       admin.Id,
		Username: admin.Username,
	}

	status = http.StatusOK
	err = nil
	return
}
