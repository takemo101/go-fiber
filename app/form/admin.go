package form

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/repository"
)

// Admin create form
type Admin struct {
	Name            string `json:"name" form:"name"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	PasswordConfirm string `json:"password_confirm" form:"password_confirm"`
	Role            string `json:"role" form:"role"`
}

func (form *Admin) createFieldsRules(create bool, id uint, repository repository.AdminRepository) []*validation.FieldRules {
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
			validation.By(func(value interface{}) error {
				email := value.(string)
				var check bool
				if create {
					check, _ = repository.ExistsByEmail(email)
				} else {
					check, _ = repository.ExistsByIDEmail(id, email)
				}
				if check {
					return errors.New("メールアドレスが重複しています")
				}
				return nil
			}),
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

	return fields
}

// Validate create or edit form validation
func (form Admin) Validate(create bool, id uint, repository repository.AdminRepository) error {
	fields := form.createFieldsRules(create, id, repository)

	fields = append(
		fields,
		validation.Field(
			&form.Role,
			validation.Required.Error("権限を選択してください"),
			validation.NotIn(
				model.RoleSystem,
				model.RoleAdmin,
			).Error("権限に正しい値を設定してください"),
		),
	)

	return validation.ValidateStruct(&form, fields...)
}

// Validate edit account validation
func (form Admin) AccountValidate(id uint, repository repository.AdminRepository) error {
	fields := form.createFieldsRules(false, id, repository)
	return validation.ValidateStruct(&form, fields...)
}

// AdminSearch search form
type AdminSearch struct {
	Keyword string `json:"keyword" form:"keyword"`
	Page    int    `json:"page" form:"page"`
}
