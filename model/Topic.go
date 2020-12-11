package model

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	Name       string
	CategoryID int
	Questions  []Question
}
