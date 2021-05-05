package query

import "go.uber.org/fx"

// Module export
var Module = fx.Options(
	fx.Provide(NewAdminQuery),
)
