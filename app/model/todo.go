package model

import "gorm.io/gorm"

// Todo is todo list
type Todo struct {
	gorm.Model
	Text   string
	Status uint64
	User   string
}
