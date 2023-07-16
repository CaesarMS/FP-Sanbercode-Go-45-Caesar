package helper

import (
	"html"
	"net/http"
	"strings"

	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/api/config"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/internal/model"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/internal/service"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/pkg/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GetAdmins(db *gorm.DB) (status int, admins []*config.AdminOutput, err error) {
	status, admins, err = service.GetAdmins(db)
	return
}

func GetAdminById(db *gorm.DB, id string) (status int, admin *config.AdminOutput, err error) {
	status, admin, err = service.GetAdminById(db, id)
	return
}

func GetAdminByUsername(db *gorm.DB, username string) (status int, admin *config.AdminOutput, err error) {
	status, admin, err = service.GetAdminByUsername(db, username)
	return
}

func AuthenticateAdmin(db *gorm.DB, username string, password string) (status int, knownAdmin config.AdminOutput, jwt string, err error) {
	// Trim username
	username = html.EscapeString(strings.TrimSpace(username))
	var admin model.Admins

	status, admin, err = service.GetAdminData(db, username)
	if err != nil {
		return
	}

	knownAdmin = config.AdminOutput{
		Id:       admin.Id,
		Username: admin.Username,
	}

	if isNotValid := util.VerifyPassword(password, admin.Password); isNotValid != nil && isNotValid == bcrypt.ErrMismatchedHashAndPassword {
		status = http.StatusUnauthorized
		err = isNotValid
		return
	}

	jwt, err = util.GenerateToken(knownAdmin.Id, 0)
	if err != nil {
		status = http.StatusInternalServerError
		return
	}

	return
}

func InsertAdmin(db *gorm.DB, username string, password string) (status int, newAdmin *config.AdminOutput, err error) {
	// Trim username
	username = html.EscapeString(strings.TrimSpace(username))

	// Hash password
	hashedPassword, errData := util.BcryptPassword(password)
	if errData != nil {
		status = http.StatusInternalServerError
		err = errData
		return
	}

	status, newAdmin, err = service.InsertAdmin(db, username, hashedPassword)
	return
}

func UpdateAdminPassword(db *gorm.DB, id string, password string) (status int, updatedAdmin *config.AdminOutput, err error) {
	// Hash password
	hashedPassword, errData := util.BcryptPassword(password)
	if errData != nil {
		status = http.StatusInternalServerError
		err = errData
		return
	}

	status, updatedAdmin, err = service.UpdateAdminPassword(db, id, hashedPassword)
	return
}

func DeleteAdmin(db *gorm.DB, id string) (status int, deletedAdmin *config.AdminOutput, err error) {
	status, deletedAdmin, err = service.DeleteAdmin(db, id)
	return
}
