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

func GetProducts(db *gorm.DB) (status int, productsLists []*config.ProductOutput, err error) {
	db.Model(&model.Products{}).Select("products.id, products.name, products.price, products.description, products.stock, categories.name as category, users.name as seller").Joins("join categories on categories.id=products.category_id").Joins("join users on users.id=products.user_id").Find(&productsLists)

	if len(productsLists) == 0 {
		status = http.StatusBadRequest
		err = errors.New("data not found")
		return
	}

	status = http.StatusOK
	err = nil
	return
}

func GetProductData(db *gorm.DB, id string) (status int, findProduct model.Products, err error) {
	if errData := db.Where("id = ?", id).First(&findProduct).Error; errData != nil {
		status = http.StatusInternalServerError
		err = errors.New("data not found")
		return
	}

	status = http.StatusOK
	err = nil
	return
}

func GetProductById(db *gorm.DB, id string) (status int, findProduct *config.ProductOutput, err error) {
	if errData := db.Model(&model.Products{}).Select("products.id, products.name, products.price, products.description, products.stock, categories.name as category, users.name as seller").Joins("join categories on categories.id=products.category_id").Joins("join users on users.id=products.user_id").Where("products.id = ?", id).First(&findProduct).Error; errData != nil {
		status = http.StatusInternalServerError
		err = errors.New("data not found")
		return
	}

	status = http.StatusOK
	err = nil
	return
}

func GetProductByCategory(db *gorm.DB, category string) (status int, productsLists []*config.ProductOutput, err error) { // Get model if exist
	if errData := db.Model(&model.Products{}).Select("products.id, products.name, products.price, products.description, products.stock, categories.name as category, users.name as seller").Joins("join categories on categories.id=products.category_id").Joins("join users on users.id=products.user_id").Where("categories.name = ?", category).Find(&productsLists).Error; errData != nil {
		status = http.StatusInternalServerError
		err = errors.New("data not found")
		return
	}

	status = http.StatusOK
	err = nil
	return
}

func InsertProduct(
	db *gorm.DB,
	name string, price uint, desc string, stock uint, category_id, seller_id string,
) (status int, newProduct *config.ProductOutput, err error) {
	if status, _, err = GetCategoryById(db, category_id); err != nil {
		return
	}

	product := model.Products{
		Id:          util.GenerateUUID(),
		Name:        name,
		Price:       price,
		Description: desc,
		Stock:       stock,
		Category_id: category_id,
		User_id:     seller_id,
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
	}
	db.Create(&product)

	status, newProduct, err = GetProductById(db, product.Id)
	if err != nil {
		return
	}

	status = http.StatusOK
	err = nil
	return
}

func UpdateProduct(
	db *gorm.DB,
	id, name string,
	price uint,
	desc string,
	stock uint,
	category_id string,
) (status int, updatedProduct *config.ProductOutput, err error) {
	if status, _, err = GetCategoryById(db, category_id); err != nil {
		return
	}

	var product model.Products
	if errData := db.Where("id = ?", id).First(&product).Error; errData != nil {
		status = http.StatusBadRequest
		err = errors.New("data not found")
		return
	}

	var updatedInput model.Products
	updatedInput.Name = name
	updatedInput.Price = price
	updatedInput.Description = desc
	updatedInput.Stock = stock
	updatedInput.Category_id = category_id
	updatedInput.Updated_at = time.Now()

	db.Model(&product).Updates(updatedInput)

	status, updatedProduct, err = GetProductById(db, product.Id)
	if err != nil {
		return
	}

	status = http.StatusOK
	err = nil
	return
}

func DeleteProduct(db *gorm.DB, id string) (status int, deletedProduct *config.ProductOutput, err error) {
	var product model.Products
	if errData := db.Where("id = ?", id).First(&product).Error; errData != nil {
		status = http.StatusBadRequest
		err = errors.New("data not found")
		return
	}

	db.Delete(&product)

	deletedProduct = &config.ProductOutput{
		Id:          product.Id,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Stock:       product.Stock,
	}

	status = http.StatusOK
	err = nil
	return
}
