package helper

import (
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/api/config"
	"github.com/CaesarMS/FP-Sanbercode-Go-45-Caesar/internal/service"
	"gorm.io/gorm"
)

func GetCategories(db *gorm.DB) (status int, categories []*config.CategoryOutput, err error) {
	status, categories, err = service.GetCategories(db)
	return
}

func GetCategoryById(db *gorm.DB, id string) (status int, category *config.CategoryOutput, err error) {
	status, category, err = service.GetCategoryById(db, id)
	return
}

func GetCategoryByName(db *gorm.DB, name string) (status int, category *config.CategoryOutput, err error) {
	status, category, err = service.GetCategoryByName(db, name)
	return
}

func InsertCategory(db *gorm.DB, name string) (status int, newCategory *config.CategoryOutput, err error) {
	status, newCategory, err = service.InsertCategory(db, name)
	return
}

func UpdateCategory(db *gorm.DB, id string, name string) (status int, updatedCategory *config.CategoryOutput, err error) {
	status, updatedCategory, err = service.UpdateCategory(db, id, name)
	return
}

func DeleteCategory(db *gorm.DB, id string) (status int, deletedCategory *config.CategoryOutput, err error) {
	status, deletedCategory, err = service.DeleteCategory(db, id)
	return
}
