package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/app/support"
	"github.com/takemo101/go-fiber/pkg"
)

// SuggestController is suggest
type SuggestController struct {
	service service.SuggestService
	value   support.RequestValue
	upload  helper.UploadHelper
	config  pkg.Config
}

// NewSuggestController is create suggest controller
func NewSuggestController(
	service service.SuggestService,
	value support.RequestValue,
	upload helper.UploadHelper,
	config pkg.Config,
) SuggestController {
	return SuggestController{
		service: service,
		value:   value,
		upload:  upload,
		config:  config,
	}
}

// Detail render suggest detail
func (ctl SuggestController) Detail(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	suggest, findErr := ctl.service.FindWithDiscussions(uint(id))
	if findErr != nil {
		return response.Error(findErr)
	}

	return response.View("suggest/detail", helper.DataMap{
		"content_footer": true,
		"suggest":        suggest,
	})
}

// Delete suggest delete process
func (ctl SuggestController) Delete(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	if err := ctl.service.Delete(uint(id)); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrDelete, ToastrDelete.Message(), Messages{})
	return response.Redirect(c, "system/discussion")
}
