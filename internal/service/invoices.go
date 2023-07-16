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

func GetInvoices(db *gorm.DB, buyerId string) (status int, invoicesLists []*config.InvoiceOutput, err error) {
	db.Model(&model.Invoices{}).Select("invoices.id as id, users.name as buyer_name, invoices.total_price, invoices.created_at as bought_at").Joins("join users on users.id=invoices.user_id").Where("users.id = ?", buyerId).Find(&invoicesLists)

	if len(invoicesLists) == 0 {
		status = http.StatusBadRequest
		err = errors.New("data not found")
		return
	}

	for i, invoice := range invoicesLists {
		var invoiceItemsLists []*config.InvoiceItemOutput
		status, invoiceItemsLists, err = getInvoiceItems(db, invoice.Id)
		if err != nil {
			return
		}

		for _, invoiceItem := range invoiceItemsLists {
			invoicesLists[i].Items = append(invoicesLists[i].Items, *invoiceItem)
		}
	}

	status = http.StatusOK
	err = nil
	return
}

func GetInvoiceById(db *gorm.DB, buyerId, id string) (status int, findInvoice *config.InvoiceOutput, err error) { // Get model if exist
	if errData := db.Model(&model.Invoices{}).Select("invoices.id as id, users.name as buyer_name, invoices.total_price, invoices.created_at as bought_at").Joins("join users on users.id=invoices.user_id").Where("invoices.id = ? AND users.id = ?", id, buyerId).First(&findInvoice).Error; errData != nil {
		status = http.StatusInternalServerError
		err = errors.New("data not found")
		return
	}

	var invoiceItemsLists []*config.InvoiceItemOutput
	status, invoiceItemsLists, err = getInvoiceItems(db, findInvoice.Id)
	if err != nil {
		return
	}

	for _, invoiceItem := range invoiceItemsLists {
		findInvoice.Items = append(findInvoice.Items, *invoiceItem)
	}

	status = http.StatusOK
	err = nil

	return
}

func InsertInvoice(
	db *gorm.DB,
	buyerId string,
	items []config.InvoiceItemInput,
) (status int, newInvoice *config.InvoiceOutput, err error) {
	var calc_price uint = 0
	for _, item := range items {
		// Check product existance
		if status, _, err = GetProductById(db, item.Product_id); err != nil {
			return
		}

		qty := item.Qty
		product_id := item.Product_id
		_, product, _ := GetProductData(db, product_id)

		if product.Stock < qty {
			status = http.StatusBadRequest
			err = errors.New(product.Name + " is out of stock")
			return
		}

		new_stock := product.Stock - qty
		_, _, err = UpdateProduct(db, product_id, product.Name, product.Price, product.Description, new_stock, product.Category_id)
		if err != nil {
			status = http.StatusInternalServerError
			return
		}

		calc_price += (qty * product.Price)
	}

	// Insert Invoice
	invoiceId := util.GenerateUUID()
	invoice := model.Invoices{
		Id:          invoiceId,
		User_id:     buyerId,
		Total_price: calc_price,
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
	}
	db.Create(&invoice)

	// Insert Invoice Items
	_, err = insertInvoiceItems(db, invoiceId, items)
	if err != nil {
		status = http.StatusInternalServerError

		// remove invoice
		status, err = deleteInvoice(db, invoiceId)
		if err != nil {
			return
		}

		return
	}

	status, newInvoice, err = GetInvoiceById(db, buyerId, invoiceId)
	if err != nil {
		return
	}

	status = http.StatusOK
	err = nil
	return
}

func deleteInvoice(db *gorm.DB, id string) (status int, err error) {
	var invoice model.Invoices
	if errData := db.Where("id = ?", id).First(&invoice).Error; errData != nil {
		status = http.StatusBadRequest
		err = errors.New("data not found")
		return
	}

	db.Delete(&invoice)

	status = http.StatusOK
	err = nil
	return
}
