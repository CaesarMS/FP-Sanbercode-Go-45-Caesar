package router

import (
	"net/http"
	"strings"

	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/api/config"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/internal/helper"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/pkg/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Get user list godoc
// @Summary Get user list by admin
// @Description Get all user, by id, or by email (select one)
// @Tags Admin
// @Param id query string false "insert user id in UUID format"
// @Param email query string false "insert user email"
// @Param Authorization header string true "Admin Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} config.UserOutput
// @Router /admin/user [get]
func GetUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var status int
	var err error

	if id := c.Query("id"); id != "" {
		var user *config.UserOutput
		status, user, err = helper.GetUserById(db, id)
		if err != nil {
			c.JSON(status, config.Response{
				Error: err.Error(),
			})
			return
		}

		c.JSON(status, config.Response{
			Data: user,
		})
		return
	}

	if email := c.Query("email"); email != "" {
		var user *config.UserOutput
		status, user, err = helper.GetUserByEmail(db, email)
		if err != nil {
			c.JSON(status, config.Response{
				Error: err.Error(),
			})
			return
		}

		c.JSON(status, config.Response{
			Data: *user,
		})
		return
	}

	var users []*config.UserOutput
	status, users, err = helper.GetUsers(db)
	if err != nil {
		c.JSON(status, config.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(status, config.Response{
		Data: users,
	})
}

// Get seller list godoc
// @Summary Get seller list by admin
// @Description Get all seller
// @Tags Admin
// @Param Authorization header string true "Admin Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} config.UserOutput
// @Router /admin/seller [get]
func GetSeller(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var status int
	var err error

	var users []*config.UserOutput
	status, users, err = helper.GetSellers(db)
	if err != nil {
		c.JSON(status, config.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(status, config.Response{
		Data: users,
	})
}

// Register a new user godoc
// @Summary Register a new user
// @Description Register as a new user
// @Tags User
// @Param Body body config.UserCreate true "the body to register as a user"
// @Produce json
// @Success 200 {object} config.UserOutput
// @Router /user/register [post]
func InsertUser(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input config.UserCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, config.Response{
			Error: err.Error(),
		})
		return
	}

	status, user, err := helper.InsertUser(db, input.Name, input.Email, input.Password, input.Address)
	if err != nil {
		c.JSON(status, config.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(status, config.Response{
		Message: "data inserted",
		Data:    user,
	})
}

// Login authenticate user godoc
// @Summary Login user authenticate
// @Description Authenticate user email & password
// @Tags User
// @Param Body body config.UserAuth true "the body to login as user"
// @Produce json
// @Success 200 {object} config.UserOutput
// @Router /user/login [post]
func AuthenticateUser(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input config.UserAuth
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, config.Response{
			Error: err.Error(),
		})
		return
	}

	status, user, jwt, err := helper.AuthenticateUser(db, input.Email, input.Password)
	if err != nil {
		c.JSON(status, config.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(status, config.Response{
		Data: map[string]string{
			"id":           user.Id,
			"email":        user.Email,
			"access_token": jwt,
		},
	})
}

// Update user password godoc
// @Summary Update user password
// @Description Update user password
// @Tags User
// @Param Body body config.UserUpdatePassword true "the body to login as user"
// @Param Authorization header string true "User Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} config.UserOutput
// @Router /user/password [patch]
func UpdateUserPassword(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input config.UserUpdatePassword
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

	// extract user id
	id := util.ExtractTokenId(token)

	status, user, err := helper.UpdateUserPassword(db, id, input.Password)
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
			Data:    user,
		},
	)
}

// Update user data godoc
// @Summary Update user data
// @Description Update user data
// @Tags User
// @Param Body body config.UserUpdateData true "the body to login as user"
// @Param Authorization header string true "User Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} config.UserOutput
// @Router /user/data [patch]
func UpdateUserData(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input config.UserUpdateData
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

	// extract id user
	id := util.ExtractTokenId(token)
	status, user, err := helper.UpdateUserData(db, id, input.Name, input.Address)
	if err != nil {
		c.JSON(status, config.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(
		status,
		config.Response{
			Message: "data updated",
			Data:    user,
		},
	)
}

// Update user to seller godoc
// @Summary Update user to seller
// @Description Update user to seller. Must re-login to get a new access token
// @Tags User
// @Param Authorization header string true "User Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} config.UserOutput
// @Router /user/seller [patch]
func UpdateUserToSeller(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// Get token
	token := c.Request.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		token = strings.Split(token, " ")[1]
	}

	// extract user id
	id := util.ExtractTokenId(token)
	status, user, err := helper.UpdateUserToSeller(db, id)
	if err != nil {
		c.JSON(status, config.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(
		status,
		config.Response{
			Message: "updated to seller",
			Data:    user,
		},
	)
}

// Delete user godoc
// @Summary Delete user
// @Description Delete user
// @Tags User
// @Param Authorization header string true "User Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} config.AdminOutput
// @Router /user/ [delete]
func DeleteUser(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// Get token
	token := c.Request.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		token = strings.Split(token, " ")[1]
	}

	// extract id user
	id := util.ExtractTokenId(token)
	status, user, err := helper.DeleteUser(db, id)
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
			Data:    user,
		},
	)
}
