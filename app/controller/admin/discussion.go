package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/form"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/object"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/app/support"
	"github.com/takemo101/go-fiber/pkg"
)

// DiscussionController is discussion
type DiscussionController struct {
	logger  pkg.Logger
	service service.DiscussionService
	value   support.RequestValue
}

// NewDiscussionController is create discussion controller
func NewDiscussionController(
	logger pkg.Logger,
	service service.DiscussionService,
	value support.RequestValue,
) DiscussionController {
	return DiscussionController{
		logger:  logger,
		service: service,
		value:   value,
	}
}

// Index render discussion list
func (ctl DiscussionController) Index(c *fiber.Ctx) error {
	var form form.DiscussionSearch
	response := ctl.value.GetResponseHelper(c)

	if err := c.QueryParser(&form); err != nil {
		return response.Error(err)
	}

	discussions, err := ctl.service.Search(object.NewDiscussionSearchInput(
		form.Keyword,
		form.Page,
	), 20)
	if err != nil {
		return response.Error(err)
	}

	return response.View("discussion/index", helper.DataMap{
		"discussions": discussions,
	})
}

// Delete discussion delete process
func (ctl DiscussionController) Delete(c *fiber.Ctx) error {
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
