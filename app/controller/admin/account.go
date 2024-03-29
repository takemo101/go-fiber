package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/form"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/app/object"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/app/support"
)

// AccountController is account
type AccountController struct {
	service service.AdminService
	value   support.RequestValue
}

// NewAccountController is create admin account controller
func NewAccountController(
	service service.AdminService,
	value support.RequestValue,
) AccountController {
	return AccountController{
		service: service,
		value:   value,
	}
}

// Edit render admin edit form
func (ctl AccountController) Edit(c *fiber.Ctx) error {
	auth := ctl.value.GetSessionAdminAuth(c)
	response := ctl.value.GetResponseHelper(c)
	return response.View("account/edit", helper.DataMap{
		"admin":          auth.Admin(),
		"content_footer": true,
	})
}

// Update admin update process
func (ctl AccountController) Update(c *fiber.Ctx) error {
	var form form.Admin
	response := ctl.value.GetResponseHelper(c)

	if err := c.BodyParser(&form); err != nil {
		return response.Error(err)
	}

	auth := ctl.value.GetSessionAdminAuth(c)

	if err := form.AccountValidate(auth.ID(), ctl.service.Repository); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message(), Messages{})
		return response.Back(c)
	}

	if _, err := ctl.service.Update(auth.ID(), object.NewAdminInput(
		form.Name,
		form.Email,
		form.Password,
		form.Role,
	)); err != nil {
		return response.Error(err)
	}

	SetToastr(c, ToastrUpdate, ToastrUpdate.Message(), Messages{})
	return response.Back(c)
}
