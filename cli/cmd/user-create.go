package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/pkg"
)

// UserCreateCommand is struct
type UserCreateCommand struct {
	logger  pkg.Logger
	root    RootCommand
	service service.UserService
}

// Setup is setup route
func (c UserCreateCommand) Setup() {
	c.logger.Info("setup admin_create-command")

	var name, email, pass string

	cmd := &cobra.Command{
		Use:   "user:create",
		Short: "create user",
		Run: func(cmd *cobra.Command, args []string) {
			user := model.User{
				Name:  name,
				Email: email,
				Pass:  []byte(pass),
			}
			newUser, err := c.service.StoreByModel(user)
			if err != nil {
				c.logger.Error(err)
				fmt.Println(err)
				return
			}
			fmt.Println(fmt.Sprintf("success create user is ID[%d]", newUser.ID))
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "user", "create user name")
	cmd.Flags().StringVarP(&email, "email", "e", "user@example.com", "create user email")
	cmd.Flags().StringVarP(&pass, "pass", "p", "user", "create user pass")

	c.root.Cmd.AddCommand(cmd)
}

// NewUserCreateCommand create new admin create command
func NewUserCreateCommand(
	root RootCommand,
	logger pkg.Logger,
	service service.UserService,
) UserCreateCommand {
	return UserCreateCommand{
		root:    root,
		logger:  logger,
		service: service,
	}
}
