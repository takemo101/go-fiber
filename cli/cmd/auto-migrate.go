package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/pkg"
)

// AutoMigrateCommand is struct
type AutoMigrateCommand struct {
	logger pkg.Logger
	root   RootCommand
	db     pkg.Database
	models Models
}

// Models is slice interface{}
type Models []interface{}

// Setup is setup route
func (c AutoMigrateCommand) Setup() {
	c.logger.Info("setup auto_migrate-command")

	c.root.Cmd.AddCommand(&cobra.Command{
		Use:   "migrate",
		Short: "auto migrate",
		Run: func(cmd *cobra.Command, args []string) {

			c.db.GormDB.AutoMigrate(c.models...)

			fmt.Println("finish migrate")
		},
	})
}

// NewAutoMigrateCommand create migrate command
func NewAutoMigrateCommand(
	root RootCommand,
	logger pkg.Logger,
	db pkg.Database,
) AutoMigrateCommand {
	return AutoMigrateCommand{
		root:   root,
		logger: logger,
		db:     db,
		// add migrate model ...
		models: Models{
			&model.Todo{},
			&model.User{},
			&model.Admin{},
		},
	}
}
