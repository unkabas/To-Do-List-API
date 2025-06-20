package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title   string
	Content string
	UserID  string
}
