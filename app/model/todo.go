package model

import (
	"gorm.io/gorm"
)

// TodoStatus for admin
type TodoStatus string

const (
	TodoStatusThing    TodoStatus = "thing"
	TodoStatusWork     TodoStatus = "work"
	TodoStatusCheck    TodoStatus = "check"
	TodoStatusComplete TodoStatus = "complete"
)

func (r TodoStatus) String() string {
	return string(r)
}

func (r TodoStatus) Name() string {
	switch r {
	case TodoStatusThing:
		return "やること"
	case TodoStatusWork:
		return "作業中"
	case TodoStatusCheck:
		return "確認中"
	}
	return "完了"
}

func ToTodoStatusArray() []KeyName {
	return []KeyName{
		{
			Key:  string(TodoStatusThing),
			Name: TodoStatusThing.Name(),
		},
		{
			Key:  string(TodoStatusWork),
			Name: TodoStatusWork.Name(),
		},
		{
			Key:  string(TodoStatusCheck),
			Name: TodoStatusCheck.Name(),
		},
		{
			Key:  string(TodoStatusComplete),
			Name: TodoStatusComplete.Name(),
		},
	}
}

// Todo is todo list
type Todo struct {
	gorm.Model
	Text    string     `gorm:"type:text;not null"`
	Status  TodoStatus `gorm:"type:varchar(191);index;not null;default:up"`
	AdminID uint       `gorm:"index"`
	Admin   Admin      `gorm:"constraint:OnDelete:CASCADE;"`
}
