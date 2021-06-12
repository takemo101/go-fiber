package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/takemo101/go-fiber/pkg"
)

// StorageLinkCommand is struct
type StorageLinkCommand struct {
	logger pkg.Logger
	config pkg.Config
	path   pkg.Path
	root   RootCommand
}

// Setup is setup command
func (c StorageLinkCommand) Setup() {
	c.logger.Info("setup storage-link")

	c.root.Cmd.AddCommand(&cobra.Command{
		Use:   "storage:link",
		Short: "static to storage symlink",
		Run: func(cmd *cobra.Command, args []string) {
			publicDir := c.path.Public("")

			// mkedir
			if f, err := os.Stat(publicDir); os.IsNotExist(err) || !f.IsDir() {
				if err := os.MkdirAll(publicDir, 0777); err != nil {
					fmt.Println(err)
				}
			}

			// symlink
			if err := os.Symlink(publicDir, c.path.Static(c.config.File.Public)); err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("success storage symlink")
		},
	})

}

// NewStorageLinkCommand create test command
func NewStorageLinkCommand(
	root RootCommand,
	config pkg.Config,
	logger pkg.Logger,
	path pkg.Path,
) StorageLinkCommand {
	return StorageLinkCommand{
		root:   root,
		config: config,
		logger: logger,
		path:   path,
	}
}
