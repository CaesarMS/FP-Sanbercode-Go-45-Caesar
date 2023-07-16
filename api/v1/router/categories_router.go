package router

import (
	"net/http"

	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/api/config"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/internal/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Get categories list godoc
// @Summary Get categories list
// @Description Get all categories, by id, or by name (select one)
// @Tags Category
// @Param id query string false "insert category id in UUID format"
// @Param name query string false "insert category name"
// @Produce json
// @Success 200 {object} config.CategoryOutput
// @Router /category [get]
func GetCategories(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var status int
	var err error

	if id := c.Query("id"); id != "" {
		var category *config.CategoryOutput
		status, category, err = helper.GetCategoryById(db, id)
		if err != nil {
			c.JSON(status, config.Response{
				Error: err.Error(),
			})
			return
		}

		c.JSON(status, config.Response{
			Data: category,
		})
		return
	}

	if name := c.Query("name"); name != "" {
		var category *config.CategoryOutput
		status, category, err = helper.GetCategoryByName(db, name)
		if err != nil {
			c.JSON(status, config.Response{
				Error: err.Error(),
			})
			return
		}

		c.JSON(status, config.Response{
			Data: *category,
		})
		return
	}

	var categories []*config.CategoryOutput
	status, categories, err = helper.GetCategories(db)
	if err != nil {
		c.JSON(status, config.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(status, config.Response{
		Data: categories,
	})
}

// Insert a new category godoc
// @Summary Insert a new category by admin
// @Description Insert a new category by admin
// @Tags Category
// @Param Body body config.CategoryInput true "the body to insert a new category"
// @Param Authorization header string true "Admin Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} config.CategoryOutput
// @Router /category/ [post]
func InsertCategory(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input config.CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, config.Response{
			Error: err.Error(),
		})
		return
	}

	status, category, err := helper.InsertCategory(db, input.Name)
	if err != nil {
		c.JSON(status, config.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(status, config.Response{
		Message: "data inserted",
		Data:    category,
	})
}

// Update a category godoc
// @Summary Update a category by admin
// @Description Update a category by admin
// @Tags Category
// @Param Body body config.CategoryInput true "the body to update a category"
// @Param id path string true "Category id in UUID format"
// @Param Authorization header string true "Admin Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} config.CategoryOutput
// @Router /category [patch]
func UpdateCategory(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input config.CategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, config.Response{
			Error: err.Error(),
		})
		return
	}

	id := c.Param("id")
	status, category, err := helper.UpdateCategory(db, id, input.Name)
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
			Data:    category,
		},
	)
}

// Delete a category godoc
// @Summary Delete a category by admin
// @Description Delete a category by admin
// @Tags Category
// @Param id path string true "Category id in UUID format"
// @Param Authorization header string true "Admin Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} config.CategoryOutput
// @Router /category/:id [delete]
func DeleteCategory(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	id := c.Param("id")
	status, category, err := helper.DeleteCategory(db, id)
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
			Data:    category,
		},
	)
}
