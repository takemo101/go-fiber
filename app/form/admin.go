package form

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
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

func (form Admin) createFieldsRules(adminForm *Admin, create bool, id uint, repository repository.AdminRepository) []*validation.FieldRules {
	baseFields := []*validation.FieldRules{
		validation.Field(
			&adminForm.Name,
			validation.Required.Error("名前を入力してください"),
			validation.Length(0, 80).Error("名前は80字以内で入力してください"),
		),
		validation.Field(
			&adminForm.Email,
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

	if create || adminForm.Password != "" {
		fields = append(
			baseFields,
			validation.Field(
				&adminForm.Password,
				validation.Required.Error("パスワードを入力してください"),
				validation.Length(5, 50).Error("パスワードは5から50文字以内で入力してください"),
				validation.By(func(value interface{}) error {
					if adminForm.PasswordConfirm != value.(string) {
						return errors.New("パスワードが一致しません")
					}
					return nil
				}),
			),
			validation.Field(
				&adminForm.PasswordConfirm,
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
	fields := form.createFieldsRules(&form, create, id, repository)

	fields = append(
		fields,
		validation.Field(
			&form.Role,
			validation.Required.Error("権限を選択してください"),
		),
	)

	return validation.ValidateStruct(&form, fields...)
}

// Validate edit account validation
func (form Admin) AccountValidate(id uint, repository repository.AdminRepository) error {
	fields := form.createFieldsRules(&form, false, id, repository)
	return validation.ValidateStruct(&form, fields...)
}

// AdminSearch search form
type AdminSearch struct {
	Keyword string `json:"keyword" form:"keyword"`
	Page    string `json:"page" form:"page"`
}
