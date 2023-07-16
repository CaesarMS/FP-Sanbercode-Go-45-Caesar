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

// Get invoice list godoc
// @Summary Get invoice list by buyer
// @Description Get all invoice or by id (select one)
// @Tags Invoice
// @Param id query string false "insert invoice id in UUID format"
// @Param Authorization header string true "User Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} config.InvoiceOutput
// @Router /invoice [get]
func GetInvoices(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var status int
	var err error

	// Get token
	token := c.Request.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		token = strings.Split(token, " ")[1]
	}

	// extract user id
	tokenUserId := util.ExtractTokenId(token)

	if id := c.Query("id"); id != "" {
		var invoice *config.InvoiceOutput
		status, invoice, err = helper.GetInvoiceById(db, tokenUserId, id)
		if err != nil {
			c.JSON(status, config.Response{
				Error: err.Error(),
			})
			return
		}

		c.JSON(status, config.Response{
			Data: invoice,
		})
		return
	}

	var invoices []*config.InvoiceOutput
	status, invoices, err = helper.GetInvoices(db, tokenUserId)
	if err != nil {
		c.JSON(status, config.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(status, config.Response{
		Data: invoices,
	})
}

// Create a new invoice godoc
// @Summary Create a new invoice by buyer
// @Description Create a new invoice as a buyer
// @Tags Invoice
// @Param Body body config.InvoiceInput true "the body to create a new invoice"
// @Param Authorization header string true "User Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} config.InvoiceOutput
// @Router /invoice [post]
func InsertInvoice(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input config.InvoiceInput
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

	status, invoice, err := helper.InsertInvoice(
		db,
		tokenUserId,
		input.Items,
	)

	if err != nil {
		c.JSON(status, config.Response{
			Error: err.Error(),
		})
		return
	}

	c.JSON(status, config.Response{
		Message: "data inserted",
		Data:    invoice,
	})
}
