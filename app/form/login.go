package form

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Login form
type Login struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// Validate login form validation
func (form Login) Validate(check func(string, string) bool) error {
	fields := []*validation.FieldRules{
		validation.Field(
			&form.Email,
			validation.Required.Error("メールアドレスを入力してください"),
			validation.By(func(value interface{}) error {
				if form.Password != "" && !check(value.(string), form.Password) {
					return errors.New("アカウント情報をお確かめください")
				}
				return nil
			}),
		),
		validation.Field(
			&form.Password,
			validation.Required.Error("パスワードを入力してください"),
		),
	}

	return validation.ValidateStruct(&form, fields...)
}
