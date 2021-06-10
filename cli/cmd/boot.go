package cmd

import (
	"github.com/takemo101/go-fiber/pkg/contract"
	"go.uber.org/fx"
)

// Module export
var Module = fx.Options(
	fx.Provide(NewTestCommand),
	fx.Provide(NewMigrateCommand),
	fx.Provide(NewRollbackCommand),
	fx.Provide(NewAutoMigrateCommand),
	fx.Provide(NewAdminCreateCommand),
	fx.Provide(NewUserCreateCommand),
	fx.Provide(NewMailCommand),
	fx.Provide(NewCommandRoot),
	fx.Provide(NewCommand),
)

// Commands is slice
type Commands []contract.Command

// NewCommand is setup command
func NewCommand(
	testCommand TestCommand,
	migrateCommand MigrateCommand,
	rollbackCommand RollbackCommand,
	autoMigrateCommand AutoMigrateCommand,
	adminCreateCommand AdminCreateCommand,
	userCreateCommand UserCreateCommand,
	mailCommand MailCommand,
) Commands {
	return Commands{
		testCommand,
		migrateCommand,
		rollbackCommand,
		autoMigrateCommand,
		adminCreateCommand,
		userCreateCommand,
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
