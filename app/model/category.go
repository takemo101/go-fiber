package model

import "gorm.io/gorm"

// Category is menu category
type Category struct {
	gorm.Model
	Name     string `gorm:"type:varchar(191);not null"`
	Sort     uint   `gorm:"index;default:1"`
	IsActive bool   `gorm:"index;default:true"`
}
