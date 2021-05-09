package repository

import "go.uber.org/fx"

// Module export
var Module = fx.Options(
	fx.Provide(NewUserRepository),
	fx.Provide(NewAdminRepository),
	fx.Provide(NewTodoRepository),
)
