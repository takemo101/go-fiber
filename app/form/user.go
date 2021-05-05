package form

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	Name            string `json:"name" form:"name"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	PasswordConfirm string `json:"password_confirm" form:"password_confirm"`
}

// Validate admin form validation
func (form User) Validate(create bool) error {
	baseFields := []*validation.FieldRules{
		validation.Field(
			&form.Name,
			validation.Required.Error("名前を入力してください"),
			validation.Length(0, 80).Error("名前は80字以内で入力してください"),
		),
		validation.Field(
			&form.Email,
			validation.Required.Error("メールアドレスを入力してください"),
			is.Email.Error("メールアドレスは[xxx@xxx.com]のような形式で入力してください"),
		),
	}

	var fields []*validation.FieldRules

	if create || form.Password != "" {
		fields = append(
			baseFields,
			validation.Field(
				&form.Password,
				validation.Required.Error("パスワードを入力してください"),
				validation.Length(5, 50).Error("パスワードは5から50文字以内で入力してください"),
				validation.By(func(value interface{}) error {
					if form.PasswordConfirm != value.(string) {
						return errors.New("パスワードが一致しません")
					}
					return nil
				}),
			),
			validation.Field(
				&form.PasswordConfirm,
				validation.Required.Error("パスワード（確認）を入力してください"),
			),
		)
	} else {
		fields = baseFields
	}

	return validation.ValidateStruct(&form, fields...)
}
