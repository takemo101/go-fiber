package model

import "gorm.io/gorm"

// Category is request category
type Category struct {
	gorm.Model
	Name     string `gorm:"type:varchar(191);uniqueIndex;not null"`
	Sort     uint   `gorm:"index;default:1"`
	IsActive bool   `gorm:"index"`
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
