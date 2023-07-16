package helper

import (
	"errors"
	"net/http"

	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/api/config"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/internal/model"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/internal/service"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/pkg/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB) (status int, users []*config.UserOutput, err error) {
	status, users, err = service.GetUsers(db)
	return
}

func GetSellers(db *gorm.DB) (status int, users []*config.UserOutput, err error) {
	status, users, err = service.GetSellers(db)
	return
}

func GetUserById(db *gorm.DB, id string) (status int, newUser *config.UserOutput, err error) {
	status, newUser, err = service.GetUserById(db, id)
	return
}

func GetUserByEmail(db *gorm.DB, email string) (status int, newUser *config.UserOutput, err error) {
	// Validate email pattern
	if !util.IsValidEmail(email) {
		status = http.StatusBadRequest
		err = errors.New("invalid email pattern")
		return
	}

	status, newUser, err = service.GetUserByEmail(db, email)
	return
}

func AuthenticateUser(db *gorm.DB, email, password string) (status int, knownUser config.UserOutput, jwt string, err error) {
	// Validate email pattern
	if !util.IsValidEmail(email) {
		status = http.StatusBadRequest
		err = errors.New("invalid email pattern")
		return
	}

	var user model.Users
	status, user, err = service.GetUserData(db, email)
	if err != nil {
		return
	}

	knownUser = config.UserOutput{
		Id:    user.Id,
		Email: user.Email,
	}

	if isNotValid := util.VerifyPassword(password, user.Password); isNotValid != nil && isNotValid == bcrypt.ErrMismatchedHashAndPassword {
		status = http.StatusUnauthorized
		err = isNotValid
		return
	}

	var memberLevel uint = 2 // buyer
	if user.Is_seller {
		memberLevel = 1 // seller
	}

	jwt, err = util.GenerateToken(knownUser.Id, memberLevel)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}

	return
}

func InsertUser(db *gorm.DB, name, email, password, address string) (status int, newUser *config.UserOutput, err error) {
	// Validate email pattern
	if !util.IsValidEmail(email) {
		status = http.StatusBadRequest
		err = errors.New("invalid email pattern")
		return
	}

	// Hash password
	hashedPassword, errData := util.BcryptPassword(password)
	if errData != nil {
		status = http.StatusInternalServerError
		err = errData
		return
	}

	status, newUser, err = service.InsertUser(db, name, email, hashedPassword, address)
	return
}

func UpdateUserPassword(db *gorm.DB, id, password string) (status int, updatedUser *config.UserOutput, err error) {
	// Hash password
	hashedPassword, errData := util.BcryptPassword(password)
	if errData != nil {
		status = http.StatusInternalServerError
		err = errData
		return
	}

	status, updatedUser, err = service.UpdateUserPassword(db, id, hashedPassword)
	return
}

func UpdateUserData(db *gorm.DB, id, name, address string) (status int, updatedUser *config.UserOutput, err error) {
	status, updatedUser, err = service.UpdateUserData(db, id, name, address)
	return
}

func UpdateUserToSeller(db *gorm.DB, id string) (status int, updatedUser *config.UserOutput, err error) {
	status, updatedUser, err = service.UpdateUserToSeller(db, id)
	return
}

func DeleteUser(db *gorm.DB, id string) (status int, updatedUser *config.UserOutput, err error) {
	status, updatedUser, err = service.DeleteUser(db, id)
	return
}
