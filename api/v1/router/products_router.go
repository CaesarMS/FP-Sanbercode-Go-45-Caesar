package router

import (
	"net/http"
	"strings"

	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/api/config"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/internal/helper"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/internal/service"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/pkg/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Get product list godoc
// @Summary Get product list
// @Description Get all product, by id, or by category name (select one)
// @Tags Product
// @Param id query string false "insert product id in UUID format"
// @Param category query string false "insert product category name"
// @Produce json
// @Success 200 {object} config.ProductOutput
// @Router /product [get]
func GetProducts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var status int
	var err error

	if id := c.Query("id"); id != "" {
		var product *config.ProductOutput
		status, product, err = helper.GetProductById(db, id)
		if err != nil {
			c.JSON(status, config.Response{
				Error: err.Error(),
			})
			return
		}

		c.JSON(status, config.Response{
			Data: product,
		})
		return
	}

	if category := c.Query("category"); category != "" {
		var products []*config.ProductOutput
		status, products, err = helper.GetProductByCategory(db, category)
		if err != nil {
			c.JSON(status, config.Response{
				Error: err.Error(),
			})
			return
		}

		c.JSON(status, config.Response{
			Data: products,
		})
		return
	}

	var products []*config.ProductOutput
	status, products, err = helper.GetProducts(db)
	if err != nil {
		c.JSON(status, config.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(status, config.Response{
		Data: products,
	})
}

// Create a new product godoc
// @Summary Create a new product by seller
// @Description Create a new product as a seller
// @Tags Product
// @Param Body body config.ProductInput true "the body to create a new product"
// @Param Authorization header string true "Seller Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} config.ProductOutput
// @Router /product [post]
func InsertProduct(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input config.ProductInput
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
	tokenUserId := util.ExtractTokenId(token)

	status, newProduct, err := helper.InsertProduct(db, input.Name, input.Price, input.Description, input.Stock, input.Category_id, tokenUserId)
	if err != nil {
		c.JSON(status, config.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(status, config.Response{
		Message: "data inserted",
		Data:    newProduct,
	})
}

// Update product data godoc
// @Summary Update product data
// @Description Update product data
// @Tags Product
// @Param id path string true "Product id in UUID format"
// @Param Body body config.ProductInput true "the body to update a product"
// @Param Authorization header string true "Seller Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} config.UserOutput
// @Router /product/:id [patch]
func UpdateProduct(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input config.ProductInput
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
	tokenUserId := util.ExtractTokenId(token)

	// get user id by product
	productId := c.Param("id")
	_, product, err := service.GetProductData(db, productId)
	if err != nil {
		return
	}

	if tokenUserId != product.User_id {
		c.JSON(http.StatusBadRequest, config.Response{
			Error: "invalid product id",
		})
		return
	}

	status, updatedProduct, err := helper.UpdateProduct(db, productId, input.Name, input.Price, input.Description, input.Stock, input.Category_id)
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
			Data:    updatedProduct,
		},
	)
}

// Delete product godoc
// @Summary Delete product
// @Description Delete product
// @Tags Product
// @Param id path string true "Product id in UUID format"
// @Param Authorization header string true "Seller Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} config.ProductOutput
// @Router /product/:id [delete]
func DeleteProduct(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// Get token
	token := c.Request.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		token = strings.Split(token, " ")[1]
	}

	// extract user id
	tokenUserId := util.ExtractTokenId(token)

	// get user id by product
	productId := c.Param("id")
	_, product, err := service.GetProductData(db, productId)
	if err != nil {
		return
	}

	if tokenUserId != product.User_id {
		c.JSON(http.StatusBadRequest, config.Response{
			Error: "invalid product id",
		})
		return
	}

	status, deletedProduct, err := helper.DeleteProduct(db, productId)
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
			Data:    deletedProduct,
		},
	)
}
