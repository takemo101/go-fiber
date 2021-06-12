package form

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
)

func CheckImageContentType(contentType string) bool {
	return contentType == "image/png" || contentType == "image/jpeg" || contentType == "image/gif"
}

func ImageRule(c *fiber.Ctx, field string, message string) validation.Rule {
	return validation.By(func(value interface{}) error {
		info, err := c.FormFile(field)
		if err != nil {
			return nil
		}

		contentType := info.Header.Get("Content-Type")

		if !CheckImageContentType(contentType) {
			return errors.New(message)
		}

		return nil
	})
}
