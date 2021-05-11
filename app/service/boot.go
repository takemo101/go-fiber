package service

import "go.uber.org/fx"

// Module export
var Module = fx.Options(
	fx.Provide(NewUserService),
	fx.Provide(NewAdminService),
	fx.Provide(NewTodoService),
)
