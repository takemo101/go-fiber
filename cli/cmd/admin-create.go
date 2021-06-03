package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/pkg"
)

// AdminCreateCommand is struct
type AdminCreateCommand struct {
	logger  pkg.Logger
	root    RootCommand
	service service.AdminService
}

// Setup is setup route
func (c AdminCreateCommand) Setup() {
	c.logger.Info("setup admin_create-command")

	var name, email, pass string

	cmd := &cobra.Command{
		Use:   "admin:create",
		Short: "create admin",
		Run: func(cmd *cobra.Command, args []string) {
			admin := model.Admin{
				Name:  name,
				Email: email,
				Pass:  []byte(pass),
			}
			newUser, err := c.service.StoreByModel(admin)
			if err != nil {
				c.logger.Error(err)
				fmt.Println(err)
				return
			}
			fmt.Println(fmt.Sprintf("success create user is ID[%d]", newUser.ID))
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "admin", "create admin name")
	cmd.Flags().StringVarP(&email, "email", "e", "admin@example.com", "create admin email")
	cmd.Flags().StringVarP(&pass, "pass", "p", "admin", "create admin pass")

	c.root.Cmd.AddCommand(cmd)
}

// NewAdminCreateCommand create new admin create command
func NewAdminCreateCommand(
	root RootCommand,
	logger pkg.Logger,
	service service.AdminService,
) AdminCreateCommand {
	return AdminCreateCommand{
		root:    root,
		logger:  logger,
		service: service,
	}
}
