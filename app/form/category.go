package form

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/takemo101/go-fiber/app/repository"
)

// Category create form
type Category struct {
	Name     string `json:"name" form:"name"`
	IsActive string `json:"is_active" form:"is_active"`
}

// Validate create or edit form validation
func (form Category) Validate(create bool, id uint, repository repository.CategoryRepository) error {
	fields := []*validation.FieldRules{
		validation.Field(
			&form.Name,
			validation.Required.Error("カテゴリ名を入力してください"),
			validation.Length(0, 80).Error("カテゴリ名は80字以内で入力してください"),
			validation.By(func(value interface{}) error {
				email := value.(string)
				var check bool
				if create {
					check, _ = repository.ExistsByName(email)
				} else {
					check, _ = repository.ExistsByIDName(id, email)
				}
				if check {
					return errors.New("カテゴリ名が重複しています")
				}
				return nil
			}),
		),
	}

	return validation.ValidateStruct(&form, fields...)
}

// Category sort input
type CategorySort struct {
	IDs []string `json:"ids[]" form:"ids[]"`
}

// Validate sort input validation
func (form CategorySort) Validate() error {
	fields := []*validation.FieldRules{
		validation.Field(
			&form.IDs,
			validation.Required.Error("カテゴリIDを設定してください"),
			validation.Each(validation.Required.Error("カテゴリIDを設定してください")),
		),
	}

	return validation.ValidateStruct(&form, fields...)
}
