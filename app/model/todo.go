package model

import "gorm.io/gorm"

// Todo is todo list
type Todo struct {
	gorm.Model
	Text    string `gorm:"type:text;not null"`
	Status  string `gorm:"type:varchar(191);index;not null"`
	AdminID uint   `gorm:"index"`
	Admin   Admin  `gorm:"constraint:OnDelete:CASCADE;"`
}
