package controller

import (
	admin "github.com/takemo101/go-fiber/app/controller/admin"
	api "github.com/takemo101/go-fiber/app/controller/api"
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	// admin controller
	fx.Provide(admin.NewDashboardController),
	fx.Provide(admin.NewAdminController),
	fx.Provide(admin.NewUserController),
	fx.Provide(admin.NewTodoController),
	fx.Provide(admin.NewTagController),
	fx.Provide(admin.NewCategoryController),
	fx.Provide(admin.NewRequestController),
	fx.Provide(admin.NewAccountController),
	fx.Provide(admin.NewSessionAuthController),

	// api controller
	fx.Provide(api.NewJWTAuthController),
)
