package cmd

import (
	"fmt"
	"strings"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/cobra"
	"github.com/takemo101/go-fiber/database"
	"github.com/takemo101/go-fiber/pkg"
)

// MigrateRollbackCommand is struct
type MigrateRollbackCommand struct {
	logger     pkg.Logger
	root       RootCommand
	db         pkg.Database
	migrations []*gormigrate.Migration
}

// Setup is setup route
func (c MigrateRollbackCommand) Setup() {
	c.logger.Info("setup migrate-command")

	var rollback string

	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "auto migrate",
		Run: func(cmd *cobra.Command, args []string) {

			m := gormigrate.New(c.db.GormDB, gormigrate.DefaultOptions, c.migrations)

			switch n := strings.ToLower(rollback); n {
			case "step":
				if err := m.RollbackLast(); err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("finish rollback step")
				}
			case "all":
				for {
					if err := m.RollbackLast(); err != nil {
						break
					}
				}
				fmt.Println("finish rollback all")
			default:
				if len(n) > 0 {
					if err := m.RollbackTo(n); err != nil {
						fmt.Println(err)
					} else {
						fmt.Println("finish rollback to id:" + n)
					}
				} else {
					if err := m.Migrate(); err != nil {
						fmt.Println(err)
					} else {
						fmt.Println("finish migrate")
					}
				}
			}
		},
	}

	cmd.Flags().StringVarP(&rollback, "rollback", "r", "", "migrate process name migrate or rollback")

	c.root.Cmd.AddCommand(cmd)
}

// NewMigrateRollbackCommand create migrate command
func NewMigrateRollbackCommand(
	root RootCommand,
	logger pkg.Logger,
	db pkg.Database,
) MigrateRollbackCommand {
	return MigrateRollbackCommand{
		root:   root,
		logger: logger,
		db:     db,
		// add migrations ...
		migrations: database.Migrations,
	}
}
