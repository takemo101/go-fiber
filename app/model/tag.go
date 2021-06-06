package model

import "gorm.io/gorm"

// Tag is menu tag
type Tag struct {
	gorm.Model
	Name string `gorm:"type:varchar(191);uniqueIndex;not null"`
	Sort uint   `gorm:"uniqueIndex;default:1"`
}

func TagsToArray(tags []Tag) []KeyName {
	var result = make([]KeyName, len(tags))
	for index, tag := range tags {
		result[index] = KeyName{
			Key:  tag.ID,
			Name: tag.Name,
		}
	}
	return result
}
