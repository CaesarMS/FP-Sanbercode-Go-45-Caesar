package helper

import (
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/api/config"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/internal/service"
	"gorm.io/gorm"
)

func GetProducts(db *gorm.DB) (status int, products []*config.ProductOutput, err error) {
	status, products, err = service.GetProducts(db)
	return
}

func GetProductById(db *gorm.DB, id string) (status int, product *config.ProductOutput, err error) {
	status, product, err = service.GetProductById(db, id)
	return
}

func GetProductByCategory(db *gorm.DB, category string) (status int, products []*config.ProductOutput, err error) {
	var categoryResult *config.CategoryOutput
	status, categoryResult, err = GetCategoryByName(db, category)
	if err != nil {
		return
	}

	status, products, err = service.GetProductByCategory(db, categoryResult.Name)
	return
}

func InsertProduct(db *gorm.DB, name string, price uint, desc string, stock uint, category_id, seller_id string) (status int, newProduct *config.ProductOutput, err error) {
	status, newProduct, err = service.InsertProduct(db, name, price, desc, stock, category_id, seller_id)
	return
}

func UpdateProduct(db *gorm.DB, id string, name string, price uint, desc string, stock uint, category_id string) (status int, updatedProduct *config.ProductOutput, err error) {
	status, updatedProduct, err = service.UpdateProduct(db, id, name, price, desc, stock, category_id)
	return
}

func DeleteProduct(db *gorm.DB, id string) (status int, deletedProduct *config.ProductOutput, err error) {
	status, deletedProduct, err = service.DeleteProduct(db, id)
	return
}
