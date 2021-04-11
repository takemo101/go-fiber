package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/takemo101/go-fiber/pkg"
)

// RootCommand is struct
type RootCommand struct {
	Cmd *cobra.Command
}

// NewCommandRoot is create command
func NewCommandRoot(config pkg.Config) RootCommand {
	var cmd = &cobra.Command{
		Use:   "version",
		Short: "display version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("-----------------")
			fmt.Println(fmt.Sprintf("app version is %s", config.App.Version))
			fmt.Println("-----------------")
		},
	}

	return RootCommand{
		Cmd: cmd,
	}
}
