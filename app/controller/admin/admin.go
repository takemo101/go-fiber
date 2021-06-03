package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/form"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/app/support"
	"github.com/takemo101/go-fiber/pkg"
)

// AdminController is admin
type AdminController struct {
	logger  pkg.Logger
	value   support.RequestValue
	service service.AdminService
}

// NewAdminController is create admin controller
func NewAdminController(
	logger pkg.Logger,
	value support.RequestValue,
	service service.AdminService,
) AdminController {
	return AdminController{
		logger:  logger,
		value:   value,
		service: service,
	}
}

// Index render admin list
func (ctl AdminController) Index(c *fiber.Ctx) error {
	var form form.AdminSearch
	response := ctl.value.GetResponseHelper(c)

	if err := c.QueryParser(&form); err != nil {
		return response.Error(err)
	}

	admins, err := ctl.service.Search(form, 20)
	if err != nil {
		return response.Error(err)
	}
	return response.View("admin/index", helper.DataMap{
		"admins": admins,
	})
}

// Create render admin create form
func (ctl AdminController) Create(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	return response.View("admin/create", helper.DataMap{
		"content_footer": true,
		"roles":          model.ToRoleArray(),
	})
}

// Store admin store process
func (ctl AdminController) Store(c *fiber.Ctx) error {
	var form form.Admin
	response := ctl.value.GetResponseHelper(c)

	if err := c.BodyParser(&form); err != nil {
		return response.Error(err)
	}

	if err := form.Validate(true, 0, ctl.service.Repository); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message())
		return response.Back(c)
	}

	if _, err := ctl.service.Store(form); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrStore, ToastrStore.Message())
	return response.Redirect(c, "system/admin")
}

// Edit render admin edit form
func (ctl AdminController) Edit(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	id, convErr := strconv.Atoi(c.Params("id"))

	if convErr != nil {
		return response.Error(convErr)
	}

	admin, findErr := ctl.service.Find(uint(id))
	if findErr != nil {
		return response.Error(findErr)
	}

	return response.View("admin/edit", helper.DataMap{
		"admin":          admin,
		"roles":          model.ToRoleArray(),
		"content_footer": true,
	})
}

// Update admin update process
func (ctl AdminController) Update(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	var form form.Admin

	if err := c.BodyParser(&form); err != nil {
		return response.Error(err)
	}

	if err := form.Validate(false, uint(id), ctl.service.Repository); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message())
		return response.Back(c)
	}

	if _, err := ctl.service.Update(uint(id), form); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrUpdate, ToastrUpdate.Message())
	return response.Back(c)
}

// Delete admin delete process
func (ctl AdminController) Delete(c *fiber.Ctx) error {
	response := ctl.value.GetResponseHelper(c)
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return response.Error(convErr)
	}

	if err := ctl.service.Delete(uint(id)); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrDelete, ToastrDelete.Message())
	return response.Redirect(c, "system/admin")
}
