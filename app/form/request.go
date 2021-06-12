package form

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/repository"
)

// Request create form
type Request struct {
	Title      string `json:"title" form:"title"`
	Content    string `json:"content" form:"content"`
	Thumbnail  string `json:"thumbnail" form:"thumbnail"`
	Status     string `json:"status" form:"status"`
	TagIDs     []uint `json:"tag_ids" form:"tag_ids"`
	CategoryID uint   `json:"category_id" form:"category_id"`
}

// Validate create or edit form validation
func (form Request) Validate(
	c *fiber.Ctx,
	categoryRepository repository.CategoryRepository,
	tagRepository repository.TagRepository,
) error {
	categoryIDs := categoryRepository.GetAllStringIDs()
	tagIDs := tagRepository.GetAllStringIDs()

	fields := []*validation.FieldRules{
		validation.Field(
			&form.Title,
			validation.Required.Error("タイトルを入力してください"),
			validation.Length(0, 180).Error("タイトルは180字以内で入力してください"),
		),
		validation.Field(
			&form.Content,
			validation.Required.Error("内容を入力してください"),
		),
		validation.Field(
			&form.Thumbnail,
			ImageRule(c, "thumbnail", "画像ファイルを選択してください"),
		),
		validation.Field(
			&form.Status,
			validation.Required.Error("投稿状況を選択してください"),
			validation.NotIn(
				model.RequestStatusDraft,
				model.RequestStatusApply,
				model.RequestStatusRemand,
				model.RequestStatusRelease,
				model.RequestStatusCancel,
			).Error("投稿状況に正しい値を設定してください"),
		),
		validation.Field(
			&form.TagIDs,
			validation.Required.Error("タグを選択してください"),
			validation.Each(
				validation.NotIn(tagIDs).Error("タグの値を正しく設定してください"),
			),
		),
		validation.Field(
			&form.CategoryID,
			validation.Required.Error("カテゴリを選択してください"),
			validation.NotIn(categoryIDs).Error("カテゴリの値を正しく設定してください"),
		),
	}

	return validation.ValidateStruct(&form, fields...)
}

// RequestSearch search form
type RequestSearch struct {
	Keyword string `json:"keyword" form:"keyword"`
	Page    int    `json:"page" form:"page"`
}
