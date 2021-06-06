package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/form"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/app/object"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/app/support"
	"github.com/takemo101/go-fiber/pkg"
)

// TagController is tag
type TagController struct {
	logger  pkg.Logger
	service service.TagService
	value   support.RequestValue
}

// NewTagController is create tag controller
func NewTagController(
	logger pkg.Logger,
	service service.TagService,
	value support.RequestValue,
) TagController {
	return TagController{
		logger:  logger,
		service: service,
		value:   value,
	}
}

// Index render tag list
func (ctl TagController) Index(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)

	tags, err := ctl.service.FindAll()
	if err != nil {
		return response.Error(err)
	}

	return response.View("tag/index", helper.DataMap{
		"tags": tags,
	})
}

// Store tag store process
func (ctl TagController) Store(c *fiber.Ctx) error {
	var form form.Tag
	response := ctl.value.GetResponseHelper(c)

	if err := c.BodyParser(&form); err != nil {
		return response.Error(err)
	}

	if err := form.Validate(true, 0, ctl.service.Repository); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message(), Messages{})
		return response.Back(c)
	}

	if _, err := ctl.service.Store(object.NewTagInput(
		form.Name,
	)); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrStore, ToastrStore.Message(), Messages{})
	return response.Back(c)
}

// Update tag update process
func (ctl TagController) Update(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	var form form.Tag

	if err := c.BodyParser(&form); err != nil {
		return response.Error(err)
	}

	uID := uint(id)

	if err := form.Validate(false, uID, ctl.service.Repository); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message(), Messages{
			"edit": true,
		})
		return response.Back(c)
	}

	if _, err := ctl.service.Update(uID, object.NewTagInput(
		form.Name,
	)); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrStore, ToastrStore.Message(), Messages{})
	return response.Back(c)
}

// Delete tag delete process
func (ctl TagController) Delete(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	if err := ctl.service.Delete(uint(id)); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrDelete, ToastrDelete.Message(), Messages{})
	return response.Back(c)
}

// Sort tag sort process
func (ctl TagController) Sort(c *fiber.Ctx) error {
	var form form.TagSort
	response := ctl.value.GetResponseHelper(c)

	if err := c.BodyParser(&form); err != nil {
		return response.Error(err)
	}

	if err := form.Validate(); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message(), Messages{})
		return response.Back(c)
	}

	if err := ctl.service.Sort(object.NewTagSortInput(
		form.IDs,
	)); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrUpdate, ToastrUpdate.Message(), Messages{})
	return response.Back(c)
}
