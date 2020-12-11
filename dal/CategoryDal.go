package dal

import (
	"quizz/model"
)

// GetCategories get all category
func GetCategories() ([]model.Category, error) {
	GetConnection()

	var categories []model.Category

	db.Limit(10).Find(&categories)

	return categories, nil

}
