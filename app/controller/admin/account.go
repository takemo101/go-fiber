package admin_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/form"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/pkg"
)

// AccountController is index
type AccountController struct {
	logger  pkg.Logger
	path    pkg.Path
	render  *helper.ViewRender
	service service.AdminService
	auth    middleware.SessionAdminAuth
}

// NewAccountController is create admin account controller
func NewAccountController(
	logger pkg.Logger,
	path pkg.Path,
	render *helper.ViewRender,
	service service.AdminService,
	auth middleware.SessionAdminAuth,
) AccountController {
	return AccountController{
		logger:  logger,
		path:    path,
		render:  render,
		service: service,
		auth:    auth,
	}
}

// Edit render admin edit form
func (ctl AccountController) Edit(c *fiber.Ctx) error {
	return ctl.render.Render("account/edit", helper.DataMap{
		"admin":          ctl.auth.Auth.Admin(),
		"content_footer": true,
	})
}

// Update admin update process
func (ctl AccountController) Update(c *fiber.Ctx) error {
	var form form.Admin

	if err := c.BodyParser(&form); err != nil {
		return ctl.render.Error(err)
	}

	if err := form.AccountValidate(ctl.auth.Auth.ID(), ctl.service.Repository); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message())
		return c.Redirect(ctl.path.URL("system/account/edit"))
	}

	if _, err := ctl.service.Update(ctl.auth.Auth.ID(), form); err != nil {
		return ctl.render.Error(err)
	}

	SetToastr(c, ToastrUpdate, ToastrUpdate.Message())
	return c.Redirect(ctl.path.URL("system/account/edit"))
}
