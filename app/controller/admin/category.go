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

// CategoryController is category
type CategoryController struct {
	logger  pkg.Logger
	service service.CategoryService
	value   support.RequestValue
}

// NewCategoryController is create category controller
func NewCategoryController(
	logger pkg.Logger,
	service service.CategoryService,
	value support.RequestValue,
) CategoryController {
	return CategoryController{
		logger:  logger,
		service: service,
		value:   value,
	}
}

// Index render category list
func (ctl CategoryController) Index(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)

	categories, err := ctl.service.FindAll()
	if err != nil {
		return response.Error(err)
	}

	return response.View("category/index", helper.DataMap{
		"categories": categories,
	})
}

// Store category store process
func (ctl CategoryController) Store(c *fiber.Ctx) error {
	var form form.Category
	response := ctl.value.GetResponseHelper(c)

	if err := c.BodyParser(&form); err != nil {
		return response.Error(err)
	}

	if err := form.Validate(true, 0, ctl.service.Repository); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message(), Messages{})
		return response.Back(c)
	}

	if _, err := ctl.service.Store(object.NewCategoryInput(
		form.Name,
		form.IsActive,
	)); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrStore, ToastrStore.Message(), Messages{})
	return response.Back(c)
}

// Update category update process
func (ctl CategoryController) Update(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	var form form.Category

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

	if _, err := ctl.service.Update(uID, object.NewCategoryInput(
		form.Name,
		form.IsActive,
	)); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrStore, ToastrStore.Message(), Messages{})
	return response.Back(c)
}

// Delete category delete process
func (ctl CategoryController) Delete(c *fiber.Ctx) error {
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

// Sort category sort process
func (ctl CategoryController) Sort(c *fiber.Ctx) error {
	var form form.CategorySort
	response := ctl.value.GetResponseHelper(c)

	if err := c.BodyParser(&form); err != nil {
		return response.Error(err)
	}

	if err := form.Validate(); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message(), Messages{})
		return response.Back(c)
	}

	if err := ctl.service.Sort(object.NewCategorySortInput(
		form.IDs,
	)); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrUpdate, ToastrUpdate.Message(), Messages{})
	return response.Back(c)
}
