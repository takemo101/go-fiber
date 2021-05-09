package controller

import (
	admin_controller "github.com/takemo101/go-fiber/app/controller/admin"
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(admin_controller.NewDashboardController),
	fx.Provide(admin_controller.NewAdminController),
	fx.Provide(admin_controller.NewUserController),
	fx.Provide(admin_controller.NewTodoController),
	fx.Provide(admin_controller.NewAccountController),
	fx.Provide(admin_controller.NewSessionAuthController),
)
