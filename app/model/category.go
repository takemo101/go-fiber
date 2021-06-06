package model

import "gorm.io/gorm"

// Category is menu category
type Category struct {
	gorm.Model
	Name     string `gorm:"type:varchar(191);uniqueIndex;not null"`
	Sort     uint   `gorm:"uniqueIndex;default:1"`
	IsActive bool   `gorm:"index;default:true"`
}

func CategoriesToArray(categories []Category) []KeyName {
	var result = make([]KeyName, len(categories))
	for index, category := range categories {
		result[index] = KeyName{
			Key:  category.ID,
			Name: category.Name,
		}
	}
	return result
}
