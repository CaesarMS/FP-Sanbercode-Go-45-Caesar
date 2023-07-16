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

func GetUsers(db *gorm.DB) (status int, userLists []*config.UserOutput, err error) {
	var users []model.Users
	db.Find(&users)

	if len(users) == 0 {
		status = http.StatusBadRequest
		err = errors.New("data not found")
		return
	}

	for _, user := range users {
		userLists = append(userLists, &config.UserOutput{
			Id:      user.Id,
			Name:    user.Name,
			Email:   user.Email,
			Address: user.Address,
		})
	}

	status = http.StatusOK
	err = nil
	return
}

func GetSellers(db *gorm.DB) (status int, userLists []*config.UserOutput, err error) {
	var users []model.Users
	db.Where("is_seller = ?", true).Find(&users)

	if len(users) == 0 {
		status = http.StatusBadRequest
		err = errors.New("data not found")
		return
	}

	for _, user := range users {
		userLists = append(userLists, &config.UserOutput{
			Id:      user.Id,
			Name:    user.Name,
			Email:   user.Email,
			Address: user.Address,
		})
	}

	status = http.StatusOK
	err = nil
	return
}

func GetUserById(db *gorm.DB, id string) (status int, findUser *config.UserOutput, err error) { // Get model if exist
	var user model.Users

	if errData := db.Where("id = ?", id).First(&user).Error; errData != nil {
		status = http.StatusInternalServerError
		err = errors.New("data not found")
		return
	}

	status = http.StatusOK
	findUser = &config.UserOutput{
		Id:      user.Id,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
	}
	err = nil

	return
}

func GetUserByEmail(db *gorm.DB, email string) (status int, findUser *config.UserOutput, err error) {
	var user model.Users

	if errData := db.Where("email = ?", email).First(&user).Error; errData != nil {
		status = http.StatusInternalServerError
		err = errors.New("data not found")
		return
	}

	status = http.StatusOK
	findUser = &config.UserOutput{
		Id:      user.Id,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
	}
	err = nil
	return
}

func GetUserData(db *gorm.DB, email string) (status int, findUser model.Users, err error) {
	if errData := db.Where("email = ?", email).First(&findUser).Error; errData != nil {
		status = http.StatusInternalServerError
		err = errors.New("data not found")
		return
	}

	status = http.StatusOK
	err = nil
	return
}

func InsertUser(
	db *gorm.DB,
	name, email, password, address string,
) (status int, newUser *config.UserOutput, err error) {
	user := model.Users{
		Id:         util.GenerateUUID(),
		Name:       name,
		Email:      email,
		Password:   password,
		Address:    address,
		Is_seller:  false,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}
	db.Create(&user)

	status, newUser, err = GetUserById(db, user.Id)
	if err != nil {
		return
	}

	status = http.StatusOK
	err = nil
	return
}

func UpdateUserPassword(db *gorm.DB, id, password string) (status int, updatedUser *config.UserOutput, err error) {
	var user model.Users
	if errData := db.Where("id = ?", id).First(&user).Error; errData != nil {
		status = http.StatusBadRequest
		err = errors.New("data not found")
		return
	}

	var updatedInput model.Users
	updatedInput.Password = password
	updatedInput.Updated_at = time.Now()
	db.Model(&user).Updates(updatedInput)

	status, updatedUser, err = GetUserById(db, user.Id)
	if err != nil {
		return
	}

	status = http.StatusOK
	err = nil
	return
}

func UpdateUserData(db *gorm.DB, id, name, address string) (status int, updatedUser *config.UserOutput, err error) {
	var user model.Users
	if errData := db.Where("id = ?", id).First(&user).Error; errData != nil {
		status = http.StatusBadRequest
		err = errors.New("data not found")
		return
	}

	var updatedInput model.Users
	updatedInput.Name = name
	updatedInput.Address = address
	updatedInput.Updated_at = time.Now()

	db.Model(&user).Updates(updatedInput)

	status, updatedUser, err = GetUserById(db, user.Id)
	if err != nil {
		return
	}

	status = http.StatusOK
	err = nil
	return
}

func UpdateUserToSeller(db *gorm.DB, id string) (status int, updatedUser *config.UserOutput, err error) {
	var user model.Users
	if errData := db.Where("id = ?", id).First(&user).Error; errData != nil {
		status = http.StatusBadRequest
		err = errors.New("data not found")
		return
	}

	var updatedInput model.Users
	updatedInput.Is_seller = true
	updatedInput.Updated_at = time.Now()

	db.Model(&user).Updates(updatedInput)

	status, updatedUser, err = GetUserById(db, user.Id)
	if err != nil {
		return
	}

	status = http.StatusOK
	err = nil
	return
}

func DeleteUser(db *gorm.DB, id string) (status int, deletedUser *config.UserOutput, err error) {
	var user model.Users
	if errData := db.Where("id = ?", id).First(&user).Error; errData != nil {
		status = http.StatusBadRequest
		err = errors.New("data not found")
		return
	}

	db.Delete(&user)

	deletedUser = &config.UserOutput{
		Id:      user.Id,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
	}

	status = http.StatusOK
	err = nil
	return
}
