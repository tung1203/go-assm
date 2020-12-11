package model

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	Content    string
	IsCorrect  bool
	QuestionID int
}
