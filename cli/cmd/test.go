package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/pkg"
)

// TestCommand is struct
type TestCommand struct {
	logger          pkg.Logger
	root            RootCommand
	categoryService service.CategoryService
	tagService      service.TagService
	requestService  service.RequestService
	suggestService  service.SuggestService
}

// Setup is setup command
func (c TestCommand) Setup() {
	c.logger.Info("setup test")

	c.root.Cmd.AddCommand(&cobra.Command{
		Use:   "test",
		Short: "test",
		Run: func(cmd *cobra.Command, args []string) {
			//
			/*
				category, categoryErr := c.categoryService.Find(1)
				if categoryErr != nil {
					fmt.Println(categoryErr.Error())
					return
				}

				tag, tagErr := c.tagService.Find(1)
				if tagErr != nil {
					fmt.Println(tagErr.Error())
					return
				}

				request, requestErr := c.requestService.Store(object.NewRequestInput(
					"aaaa",
					"aaaa",
					"release",
					[]string{
						string(tag.ID),
					},
					string(category.ID),
				), 1)
				if requestErr != nil {
					fmt.Println(requestErr.Error())
					return
				}
			*/
			/*
				request, requestErr := c.requestService.Find(1)
				if requestErr != nil {
					fmt.Println(requestErr.Error())
					return
				}

				suggest, suggestErr := c.suggestService.SendSuggestMessage(
					request.ID,
					2,
					"hello",
				)
				if suggestErr != nil {
					fmt.Println(suggestErr.Error())
					return
				}
			*/
			/*
					suggest, suggestErr := c.suggestService.Find(5)
					if suggestErr != nil {
						fmt.Println(suggestErr.Error())
						return
					}

					_, discussionErr := c.discussionService.SendDeclineMessage(
						suggest.ID,
						1,
						"facu",
					)
				if discussionErr != nil {
					fmt.Println(discussionErr.Error())
					return
				}
			*/
			/*
				_, suggest2Err := c.suggestService.ReplyDecline(suggest.ID, 2, true)
				if suggest2Err != nil {
					fmt.Println(suggest2Err.Error())
					return
				}
			*/
			fmt.Println("success")
		},
	})

}

// NewTestCommand create test command
func NewTestCommand(
	root RootCommand,
	logger pkg.Logger,
	categoryService service.CategoryService,
	tagService service.TagService,
	requestService service.RequestService,
	suggestService service.SuggestService,
) TestCommand {
	return TestCommand{
		root:            root,
		logger:          logger,
		categoryService: categoryService,
		tagService:      tagService,
		requestService:  requestService,
		suggestService:  suggestService,
	}
}
