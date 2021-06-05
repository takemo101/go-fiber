package model

import (
	"gorm.io/gorm"
)

// User is auth user
type User struct {
	gorm.Model
	Name  string `gorm:"type:varchar(191);index;not null"`
	Email string `gorm:"type:varchar(191);uniqueIndex;not null"`
	Pass  []byte
	Menus []Menu
}
