package cmd

import "go.uber.org/fx"

// Module export
var Module = fx.Options(
	fx.Provide(NewMigrateCommand),
	fx.Provide(NewAdminCreateCommand),
	fx.Provide(NewCommandRoot),
	fx.Provide(NewCommand),
)

// Commands is slice
type Commands []Command

// Command is interface
type Command interface {
	Setup()
}

// NewCommand is setup routes
func NewCommand(
	migrateCommand MigrateCommand,
	adminCreateCommand AdminCreateCommand,
) Commands {
	return Commands{
		migrateCommand,
		adminCreateCommand,
	}
}

// Setup all the command
func (commands Commands) Setup() {
	for _, cmd := range commands {
		cmd.Setup()
	}
}
