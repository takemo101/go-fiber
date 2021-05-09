package model

import (
	"gorm.io/gorm"
)

// User is auth user
type Admin struct {
	gorm.Model
	Name  string `gorm:"type:varchar(191);index;not null"`
	Email string `gorm:"type:varchar(191);uniqueIndex;not null"`
	Pass  []byte
	Todos []Todo
}
