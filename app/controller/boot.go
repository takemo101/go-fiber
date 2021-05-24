package controller

import (
	admin "github.com/takemo101/go-fiber/app/controller/admin"
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(admin.NewDashboardController),
	fx.Provide(admin.NewAdminController),
	fx.Provide(admin.NewUserController),
	fx.Provide(admin.NewTodoController),
	fx.Provide(admin.NewAccountController),
	fx.Provide(admin.NewSessionAuthController),
)
