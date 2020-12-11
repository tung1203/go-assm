package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	username string
	password string
	score    int
}
