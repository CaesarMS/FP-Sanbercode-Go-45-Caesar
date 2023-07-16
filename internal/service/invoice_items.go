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

func getInvoiceItems(db *gorm.DB, invoiceId string) (status int, invoiceItemsLists []*config.InvoiceItemOutput, err error) {
	db.Model(&model.Invoice_items{}).Select("invoice_items.id, invoice_items.qty, invoice_items.price, products.name as product_name").Joins("join products on products.id=invoice_items.product_id").Where("invoice_items.invoice_id = ?", invoiceId).Find(&invoiceItemsLists)

	if len(invoiceItemsLists) == 0 {
		status = http.StatusBadRequest
		err = errors.New("data not found")
		return
	}

	status = http.StatusOK
	err = nil
	return
}

func insertInvoiceItems(
	db *gorm.DB,
	invoice_id string,
	items []config.InvoiceItemInput,
) (status int, err error) {
	for _, item := range items {
		qty := item.Qty
		product_id := item.Product_id
		_, product, _ := GetProductData(db, product_id)

		// Insert Invoice Items
		invoice_item_id := util.GenerateUUID()
		invoice_item := model.Invoice_items{
			Id:         invoice_item_id,
			Invoice_id: invoice_id,
			Product_id: product_id,
			Qty:        qty,
			Price:      qty * product.Price,
			Created_at: time.Now(),
			Updated_at: time.Now(),
		}
		db.Create(&invoice_item)
	}

	status = http.StatusOK
	err = nil
	return
}
