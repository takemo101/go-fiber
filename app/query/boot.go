package query

import "go.uber.org/fx"

// Module export
var Module = fx.Options(
	fx.Provide(NewAdminQuery),
	fx.Provide(NewUserQuery),
	fx.Provide(NewTodoQuery),
	fx.Provide(NewRequestQuery),
)
