package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/takemo101/go-fiber/app/helper"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/pkg"
)

// UserController is index
type UserController struct {
	logger  pkg.Logger
	path    pkg.Path
	render  *helper.ViewRender
	service service.UserService
}

// NewUserController is create user controller
func NewUserController(
	logger pkg.Logger,
	path pkg.Path,
	render *helper.ViewRender,
	service service.UserService,
) UserController {
	return UserController{
		logger:  logger,
		path:    path,
		render:  render,
		service: service,
	}
}

// Index render user list
func (ctl UserController) Index(c *fiber.Ctx) error {
	users, err := ctl.service.Search()
	if err != nil {
		return ctl.render.Error(err)
	}
	return ctl.render.Render("user/index", helper.DataMap{
		"users": users,
	})
}
