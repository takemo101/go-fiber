package pkg

import "go.uber.org/fx"

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewConfig),
	fx.Provide(NewLogger),
	fx.Provide(NewDatabase),
	fx.Provide(NewApplication),
)
