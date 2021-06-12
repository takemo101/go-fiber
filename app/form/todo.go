package form

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Todo create form
type Todo struct {
	Text   string `json:"text" form:"text"`
	Status string `json:"status" form:"status"`
}

// Validate create or edit form validation
func (form Todo) Validate() error {
	fields := []*validation.FieldRules{
		validation.Field(
			&form.Text,
			validation.Required.Error("内容を入力してください"),
			validation.Length(0, 3000).Error("内容は3,000字以内で入力してください"),
		),
		validation.Field(
			&form.Status,
			validation.Required.Error("進捗状況を入力してください"),
		),
	}

	return validation.ValidateStruct(&form, fields...)
}

// TodoSearch search form
type TodoSearch struct {
	Keyword string `json:"keyword" form:"keyword"`
	Page    int    `json:"page" form:"page"`
}
