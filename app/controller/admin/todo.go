package admin

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/form"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/middleware"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/pkg"
)

// TodoController is index
type TodoController struct {
	logger  pkg.Logger
	path    pkg.Path
	render  *helper.ViewRender
	service service.TodoService
	auth    middleware.SessionAdminAuth
}

// NewTodoController is create todo controller
func NewTodoController(
	logger pkg.Logger,
	path pkg.Path,
	render *helper.ViewRender,
	service service.TodoService,
	auth middleware.SessionAdminAuth,
) TodoController {
	return TodoController{
		logger:  logger,
		path:    path,
		render:  render,
		service: service,
		auth:    auth,
	}
}

// Index render todo list
func (ctl TodoController) Index(c *fiber.Ctx) error {
	var form form.TodoSearch

	if err := c.QueryParser(&form); err != nil {
		return ctl.render.Error(err)
	}

	todos, err := ctl.service.Search(form, 20)
	if err != nil {
		return ctl.render.Error(err)
	}

	ctl.render.SetJS(helper.DataMap{
		"statuses": model.ToTodoStatusArray(),
	})
	return ctl.render.Render("todo/index", helper.DataMap{
		"todos":    todos,
		"statuses": model.ToTodoStatusArray(),
	})
}

// Store todo store process
func (ctl TodoController) Store(c *fiber.Ctx) error {
	var form form.Todo

	if err := c.BodyParser(&form); err != nil {
		return ctl.render.Error(err)
	}

	if err := form.Validate(); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message())
		return c.Redirect(ctl.path.URL("system/todo"))
	}

	if _, err := ctl.service.Store(form, ctl.auth.Auth.ID()); err != nil {
		return ctl.render.Error(err)
	}

	SetToastr(c, ToastrStore, ToastrStore.Message())
	return c.Redirect(ctl.path.URL("system/todo"))
}

// Update todo update process
func (ctl TodoController) Update(c *fiber.Ctx) error {
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return ctl.render.Error(convErr)
	}

	var form form.Todo

	if err := c.BodyParser(&form); err != nil {
		return ctl.render.Error(err)
	}

	if err := form.Validate(); err != nil {
		middleware.SetErrorResource(c, helper.ErrorsToMap(err), helper.StructToFormMap(&form))
		SetToastr(c, ToastrError, ToastrError.Message())
		return c.Redirect(ctl.path.URL("system/todo"))
	}

	if _, err := ctl.service.Update(uint(id), form); err != nil {
		return ctl.render.Error(err)
	}

	SetToastr(c, ToastrUpdate, ToastrUpdate.Message())
	return c.Redirect(ctl.path.URL("system/todo", c.Params("id")))
}

// ChangeStatus todo update process
func (ctl TodoController) ChangeStatus(c *fiber.Ctx) error {
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return c.JSON(&fiber.Map{
			"success": false,
			"error":   convErr,
		})
	}

	var form form.Todo

	if err := c.BodyParser(&form); err != nil {
		return c.JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	if _, err := ctl.service.ChangeStatus(uint(id), form.Status); err != nil {
		return c.JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return c.JSON(&fiber.Map{
		"success": true,
		"message": "change stauts successfully",
	})
}

// Delete todo delete process
func (ctl TodoController) Delete(c *fiber.Ctx) error {
	id, convErr := strconv.Atoi(c.Params("id"))
	if convErr != nil {
		return ctl.render.Error(convErr)
	}

	if err := ctl.service.Delete(uint(id)); err != nil {
		return ctl.render.Error(err)
	}

	SetToastr(c, ToastrDelete, ToastrDelete.Message())
	return c.Redirect(ctl.path.URL("system/todo"))
}
