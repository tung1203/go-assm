package model

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Content string
	TopicID int
	Answers []Answer
}
