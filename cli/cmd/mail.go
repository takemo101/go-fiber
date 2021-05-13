package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/takemo101/go-fiber/pkg"
)

// MailCommand is struct
type MailCommand struct {
	logger  pkg.Logger
	root    RootCommand
	factory pkg.MailFactory
}

// Setup is setup command
func (c MailCommand) Setup() {
	c.logger.Info("setup mail_send-command")

	var subject, to, message string

	cmd := &cobra.Command{
		Use:   "mail:send",
		Short: "mail send",
		Run: func(cmd *cobra.Command, args []string) {
			mail := c.factory.Create()
			mail.Subject(subject)
			mail.To(to)
			mail.TemplateText("mail/test", pkg.BindData{
				"message": message,
			})
			if err := mail.Send(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(fmt.Sprintf("success send mail to %s", to))
			}
		},
	}

	cmd.Flags().StringVarP(&subject, "subject", "s", "subject", "create mail subject")
	cmd.Flags().StringVarP(&to, "to", "t", "xxx@xxx.com", "create mail to address")
	cmd.Flags().StringVarP(&message, "message", "m", "message", "create mail message")

	c.root.Cmd.AddCommand(cmd)
}

// NewMailCommand create new mail send command
func NewMailCommand(
	root RootCommand,
	logger pkg.Logger,
	factory pkg.MailFactory,
) MailCommand {
	return MailCommand{
		root:    root,
		logger:  logger,
		factory: factory,
	}
}
