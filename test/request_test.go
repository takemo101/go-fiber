package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/object"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/boot"
	"github.com/takemo101/go-fiber/database"
	"github.com/takemo101/go-fiber/pkg"
)

func Test_RequestSuggestDisscussionService_OK(t *testing.T) {
	boot.Testing(
		t,
		func(
			db pkg.Database,
			userService service.UserService,
			tagService service.TagService,
			categoryService service.CategoryService,
			requestService service.RequestService,
			suggestService service.SuggestService,
			discussionService service.DiscussionService,
		) {
			db.GormDB.AutoMigrate(database.Models...)

			// create request tags
			tag, tagErr := tagService.Store(
				object.NewTagInput(
					"request-tag",
				),
			)
			assert.Equal(t, nil, tagErr)

			// create request category
			category, categoryErr := categoryService.Store(
				object.NewCategoryInput(
					"request-category",
					"true",
				),
			)
			assert.Equal(t, nil, categoryErr)

			// create request test user
			requestUser, requestUserErr := userService.Store(
				object.NewUserInput(
					"request-user",
					"user@request.com",
					"test",
				),
			)
			assert.Equal(t, nil, requestUserErr)

			// create request test user
			suggestUser, suggestUserErr := userService.Store(
				object.NewUserInput(
					"suggest-user",
					"user@suggest.com",
					"test",
				),
			)
			assert.Equal(t, nil, suggestUserErr)

			t.Run("request store", func(t *testing.T) {
				title := "request-test"
				content := "request-content"
				thumbnail := "test/image.jpg"
				status := string(model.RequestStatusRelease)
				tagIDs := []uint{tag.ID}
				categoryID := category.ID

				for i := 0; i < 10; i++ {
					requestService.Store(
						object.NewRequestInput(
							title,
							content,
							thumbnail,
							status,
							tagIDs,
							categoryID,
						),
						requestUser.ID,
					)
				}

				request, err := requestService.Find(1)

				assert.Equal(t, nil, err)
				assert.Equal(t, uint(1), request.ID)
			})
			t.Run("request update", func(t *testing.T) {
				title := "request-testing"
				content := "request-content"
				thumbnail := "test/image.jpg"
				status := string(model.RequestStatusRelease)
				tagIDs := []uint{tag.ID}
				categoryID := category.ID

				requestService.Update(
					1,
					object.NewRequestInput(
						title,
						content,
						thumbnail,
						status,
						tagIDs,
						categoryID,
					),
				)

				findRequest, _ := requestService.Find(1)

				assert.Equal(t, title, findRequest.Title)
			})
			t.Run("request index", func(t *testing.T) {
				requests, paginator, _ := requestService.Search(
					object.NewRequestSearchInput(
						"",
						0,
					),
					5,
				)

				assert.Equal(t, 5, len(requests))
				assert.Equal(t, 5, paginator.LastCount)
			})
			t.Run("suggest send start message", func(t *testing.T) {
				message := "suggest start"
				findRequest, _ := requestService.Find(1)

				suggest, _ := suggestService.SendStartMessage(
					findRequest.ID,
					suggestUser.ID,
					message,
				)

				findSuggest, err := suggestService.Find(suggest.ID)

				assert.Equal(t, nil, err)
				assert.Equal(t, suggest.ID, findSuggest.ID)
				assert.Equal(t, model.SuggestStatusStart, findSuggest.Status)
			})
		},
	)
}
