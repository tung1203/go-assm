package dal

import (
	"quizz/model"
)

// GetCategories get all category
func GetCategories() ([]model.Category, error) {
	GetConnection()

	var categories []model.Category

	// categoriy := model.Category{Name: "dcm", Image: "dcm", Topics: []model.Topic{{Name: "sd", Questions: []model.Question{{Content: "wq", Answers: []model.Answer{{Content: "ds", IsCorrect: true}}}}}}}

	// db.Create(&categoriy)

	db.Limit(10).Find(&categories)

	return categories, nil

}
