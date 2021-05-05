package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/pkg"
)

// MigrateCommand is struct
type MigrateCommand struct {
	logger pkg.Logger
	root   RootCommand
	db     pkg.Database
}

// Models is slice interface{}
type Models []interface{}

// Setup is setup route
func (c MigrateCommand) Setup() {
	c.logger.Info("setup migurate-command")

	c.root.Cmd.AddCommand(&cobra.Command{
		Use:   "migrate",
		Short: "auto migrate",
		Run: func(cmd *cobra.Command, args []string) {

			c.db.GormDB.AutoMigrate(GetModels()...)

			fmt.Println("finish migrate")
		},
	})
}

// NewMigrateCommand create migrate command
func NewMigrateCommand(
	root RootCommand,
	logger pkg.Logger,
	db pkg.Database,
) MigrateCommand {
	return MigrateCommand{
		root:   root,
		logger: logger,
		db:     db,
	}
}

// GetModels is auto migrate model
func GetModels() Models {
	return Models{
		&model.Todo{},
		&model.User{},
		&model.Admin{},
		// add migrate model ...
	}
}
