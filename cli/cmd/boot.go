package cmd

import (
	"github.com/takemo101/go-fiber/pkg/contract"
	"go.uber.org/fx"
)

// Module export
var Module = fx.Options(
	fx.Provide(NewMigrateRollbackCommand),
	fx.Provide(NewAutoMigrateCommand),
	fx.Provide(NewAdminCreateCommand),
	fx.Provide(NewMailCommand),
	fx.Provide(NewCommandRoot),
	fx.Provide(NewCommand),
)

// Commands is slice
type Commands []contract.Command

// NewCommand is setup routes
func NewCommand(
	migrateRollbackCommand MigrateRollbackCommand,
	autoMigrateCommand AutoMigrateCommand,
	adminCreateCommand AdminCreateCommand,
	mailCommand MailCommand,
) Commands {
	return Commands{
		migrateRollbackCommand,
		autoMigrateCommand,
		adminCreateCommand,
		mailCommand,
	}
}

// Setup all the command
func (commands Commands) Setup() {
	for _, cmd := range commands {
		cmd.Setup()
	}
}

// CLIBoot all command setup
func (commands Commands) CLIBoot() {
	commands.Setup()
}
