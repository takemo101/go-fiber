package form

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Discussion create form
type Discussion struct {
	Message string `json:"message" form:"message"`
}

// Validate create form validation
func (form Discussion) Validate(requestID uint) error {
	fields := []*validation.FieldRules{
		validation.Field(
			&form.Message,
			validation.Required.Error("メッセージを入力してください"),
			validation.Length(0, 3000).Error("メッセージは3,000字以内で入力してください"),
		),
	}

	return validation.ValidateStruct(&form, fields...)
}
