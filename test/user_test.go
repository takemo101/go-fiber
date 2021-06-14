package test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/object"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/boot"
	"github.com/takemo101/go-fiber/database"
	"github.com/takemo101/go-fiber/pkg"
)

func Test_UserService_OK(t *testing.T) {
	boot.Testing(
		t,
		func(
			db pkg.Database,
			service service.UserService,
		) {
			db.GormDB.AutoMigrate(database.Models...)

			t.Run("user store", func(t *testing.T) {
				name := "test"
				email := "test@example.com"

				var user model.User
				for i := 0; i < 10; i++ {
					user, _ = service.Store(
						object.NewUserInput(
							name+strconv.Itoa(i),
							email+strconv.Itoa(i),
							"test",
						),
					)
				}

				findUser, err := service.Find(user.ID)

				assert.Equal(t, nil, err)
				assert.Equal(t, user.ID, findUser.ID)
			})
			t.Run("user update", func(t *testing.T) {
				name := "testing"
				email := "test@example.com"

				service.Update(
					1,
					object.NewUserInput(
						name,
						email,
						"test",
					),
				)

				findUser, _ := service.Find(1)

				assert.Equal(t, name, findUser.Name)
			})
			t.Run("user index", func(t *testing.T) {
				users, paginator, _ := service.Search(
					object.NewUserSearchInput(
						"",
						0,
					),
					5,
				)

				assert.Equal(t, 5, len(users))
				assert.Equal(t, 5, paginator.LastCount)
			})
			t.Run("user delete", func(t *testing.T) {
				deleteErr := service.Delete(1)
				_, findErr := service.Find(1)
				assert.Equal(t, nil, deleteErr)
				assert.NotEqual(t, nil, findErr)
			})
		},
	)
}
