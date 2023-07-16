package router

import (
	"net/http"
	"strings"

	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/api/config"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/internal/helper"
	_ "github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/internal/model"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/pkg/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Get admins list godoc
// @Summary Get admins list
// @Description Get all admins, by id, or by username (select one)
// @Tags Admin
// @Param id query string false "insert admin id in UUID format"
// @Param username query string false "insert admin username"
// @Param Authorization header string true "Admin Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} config.AdminOutput
// @Router /admin [get]
func GetAdmins(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var status int
	var err error

	if id := c.Query("id"); id != "" {
		var admin *config.AdminOutput
		status, admin, err = helper.GetAdminById(db, id)
		if err != nil {
			c.JSON(status, config.Response{
				Error: err.Error(),
			})
			return
		}

		c.JSON(status, config.Response{
			Data: admin,
		})
		return
	}

	if username := c.Query("username"); username != "" {
		var admin *config.AdminOutput
		status, admin, err = helper.GetAdminByUsername(db, username)
		if err != nil {
			c.JSON(status, config.Response{
				Error: err.Error(),
			})
			return
		}

		c.JSON(status, config.Response{
			Data: *admin,
		})
		return
	}

	var admins []*config.AdminOutput
	status, admins, err = helper.GetAdmins(db)
	if err != nil {
		c.JSON(status, config.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(status, config.Response{
		Data: admins,
	})
}

// Login authenticate admin godoc
// @Summary Login admin authenticate
// @Description Authenticate admin username & password
// @Tags Admin
// @Param Body body config.AdminAuth true "the body to login as admin"
// @Produce json
// @Success 200 {object} config.AdminOutput
// @Router /admin/login [post]
func AuthenticateAdmin(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input config.AdminAuth
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, config.Response{
			Error: err.Error(),
		})
		return
	}

	status, admin, jwt, err := helper.AuthenticateAdmin(db, input.Username, input.Password)
	if err != nil {
		c.JSON(status, config.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(status, config.Response{
		Data: map[string]string{
			"id":           admin.Id,
			"username":     admin.Username,
			"access_token": jwt,
		},
	})
}

// Register a new admin godoc
// @Summary Register a new admin
// @Description Register by username & password
// @Tags Admin
// @Param Body body config.AdminAuth true "the body to login as admin"
// @Param x-api-key header string true "API Key: 9c6f9769-6d5b-493d-ae2e-4fad70711564"
// @Security x-api-key
// @Produce json
// @Success 200 {object} config.AdminOutput
// @Router /admin/register [post]
func InsertAdmin(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input config.AdminAuth
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, config.Response{
			Error: err.Error(),
		})
		return
	}

	status, admin, err := helper.InsertAdmin(db, input.Username, input.Password)
	if err != nil {
		c.JSON(status, config.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(status, config.Response{
		Message: "data inserted",
		Data:    admin,
	})
}

// Update admin password godoc
// @Summary Update admin password
// @Description Update admin password
// @Tags Admin
// @Param Body body config.AdminUpdate true "the body to login as admin"
// @Param Authorization header string true "Admin Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} config.AdminOutput
// @Router /admin/password [patch]
func UpdateAdminPassword(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input config.AdminUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, config.Response{
			Error: err.Error(),
		})
		return
	}

	// Get token
	token := c.Request.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		token = strings.Split(token, " ")[1]
	}

	// extract admin id
	id := util.ExtractTokenId(token)
	status, admin, err := helper.UpdateAdminPassword(db, id, input.Password)
	if err != nil {
		c.JSON(status, config.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(
		status,
		config.Response{
			Message: "password updated",
			Data:    admin,
		},
	)
}

// Delete admin godoc
// @Summary Delete admin
// @Description Delete admin
// @Tags Admin
// @Param Authorization header string true "Admin Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} config.AdminOutput
// @Router /admin/ [delete]
func DeleteAdmin(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// Get token
	token := c.Request.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		token = strings.Split(token, " ")[1]
	}

	// extract admin id
	id := util.ExtractTokenId(token)
	status, admin, err := helper.DeleteAdmin(db, id)
	if err != nil {
		c.JSON(status, config.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(
		status,
		config.Response{
			Message: "data deleted",
			Data:    admin,
		},
	)
}
