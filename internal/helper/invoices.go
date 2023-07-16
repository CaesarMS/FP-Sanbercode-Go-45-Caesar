package helper

import (
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/api/config"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/internal/service"
	"gorm.io/gorm"
)

func GetInvoices(db *gorm.DB, buyerId string) (status int, invoices []*config.InvoiceOutput, err error) {
	status, invoices, err = service.GetInvoices(db, buyerId)
	return
}

func GetInvoiceById(db *gorm.DB, buyerId, id string) (status int, invoice *config.InvoiceOutput, err error) {
	status, invoice, err = service.GetInvoiceById(db, buyerId, id)
	return
}

func InsertInvoice(db *gorm.DB, buyerId string, items []config.InvoiceItemInput) (status int, invoice *config.InvoiceOutput, err error) {
	status, invoice, err = service.InsertInvoice(db, buyerId, items)
	return
}
