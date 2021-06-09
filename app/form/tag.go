package form

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/takemo101/go-fiber/app/repository"
)

// Tag create form
type Tag struct {
	Name string `json:"name" form:"name"`
}

// Validate create or edit form validation
func (form Tag) Validate(create bool, id uint, repository repository.TagRepository) error {
	fields := []*validation.FieldRules{
		validation.Field(
			&form.Name,
			validation.Required.Error("タグ名を入力してください"),
			validation.Length(0, 80).Error("タグ名は80字以内で入力してください"),
			validation.By(func(value interface{}) error {
				email := value.(string)
				var check bool
				if create {
					check, _ = repository.ExistsByName(email)
				} else {
					check, _ = repository.ExistsByIDName(id, email)
				}
				if check {
					return errors.New("タグ名が重複しています")
				}
				return nil
			}),
		),
	}

	return validation.ValidateStruct(&form, fields...)
}

// Tag sort input
type TagSort struct {
	IDs []string `json:"ids" form:"ids"`
}

// Validate sort input validation
func (form TagSort) Validate() error {
	fields := []*validation.FieldRules{
		validation.Field(
			&form.IDs,
			validation.Required.Error("タグIDを設定してください"),
			validation.Each(validation.Required.Error("タグIDを設定してください")),
		),
	}

	return validation.ValidateStruct(&form, fields...)
}
