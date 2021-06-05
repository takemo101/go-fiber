package model

import "gorm.io/gorm"

// Tag is menu tag
type Tag struct {
	gorm.Model
	Name string `gorm:"type:varchar(191);not null"`
}
