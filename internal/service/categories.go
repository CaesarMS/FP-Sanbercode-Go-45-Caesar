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

func GetCategories(db *gorm.DB) (status int, categoriesLists []*config.CategoryOutput, err error) {
	var categories []model.Categories
	db.Find(&categories)

	if len(categories) == 0 {
		status = http.StatusBadRequest
		err = errors.New("data not found")
		return
	}

	for _, category := range categories {
		categoriesLists = append(categoriesLists, &config.CategoryOutput{
			Id:   category.Id,
			Name: category.Name,
		})
	}

	status = http.StatusOK
	err = nil
	return
}

func GetCategoryById(db *gorm.DB, id string) (status int, findCategory *config.CategoryOutput, err error) { // Get model if exist
	var category model.Categories

	if errData := db.Where("id = ?", id).First(&category).Error; errData != nil {
		status = http.StatusInternalServerError
		err = errors.New("data not found")
		return
	}

	status = http.StatusOK
	findCategory = &config.CategoryOutput{
		Id:   category.Id,
		Name: category.Name,
	}
	err = nil

	return
}

func GetCategoryByName(db *gorm.DB, name string) (status int, findCategory *config.CategoryOutput, err error) { // Get model if exist
	var category model.Categories

	if errData := db.Where("name = ?", name).First(&category).Error; errData != nil {
		status = http.StatusInternalServerError
		err = errors.New("data not found")
		return
	}

	status = http.StatusOK
	findCategory = &config.CategoryOutput{
		Id:   category.Id,
		Name: category.Name,
	}
	err = nil
	return
}

func InsertCategory(
	db *gorm.DB,
	name string,
) (status int, newCategory *config.CategoryOutput, err error) {
	category := model.Categories{Id: util.GenerateUUID(), Name: name, Created_at: time.Now(), Updated_at: time.Now()}
	db.Create(&category)

	status, newCategory, err = GetCategoryById(db, category.Id)
	if err != nil {
		return
	}

	status = http.StatusOK
	err = nil
	return
}

func UpdateCategory(db *gorm.DB, id string, name string) (status int, updatedCategory *config.CategoryOutput, err error) {
	var category model.Categories
	if errData := db.Where("id = ?", id).First(&category).Error; errData != nil {
		status = http.StatusBadRequest
		err = errors.New("data not found")
		return
	}

	var updatedInput model.Categories
	updatedInput.Name = name
	updatedInput.Updated_at = time.Now()

	db.Model(&category).Updates(updatedInput)

	status, updatedCategory, err = GetCategoryById(db, category.Id)
	if err != nil {
		return
	}

	status = http.StatusOK
	err = nil
	return
}

func DeleteCategory(db *gorm.DB, id string) (status int, deletedCategory *config.CategoryOutput, err error) {
	var category model.Categories
	if errData := db.Where("id = ?", id).First(&category).Error; errData != nil {
		status = http.StatusBadRequest
		err = errors.New("data not found")
		return
	}

	db.Delete(&category)

	deletedCategory = &config.CategoryOutput{
		Id:   category.Id,
		Name: category.Name,
	}

	status = http.StatusOK
	err = nil
	return
}
