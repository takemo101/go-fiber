package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/takemo101/go-fiber/app/object"
	"github.com/takemo101/go-fiber/app/service"
	"github.com/takemo101/go-fiber/boot"
	"github.com/takemo101/go-fiber/database"
	"github.com/takemo101/go-fiber/pkg"
)

func Test_UserService_OK(t *testing.T) {
	boot.Testing(
		func(
			db pkg.Database,
			service service.UserService,
		) {
			db.GormDB.AutoMigrate(database.Models...)

			name := "test"
			email := "test@example.com"

			user, _ := service.Store(
				object.NewUserInput(
					name,
					email,
					"test",
				),
			)

			assert.Equal(t, name, user.Name)
		},
		func(
			db pkg.Database,
			service service.UserService,
		) {
			name := "testing"
			email := "test@example.com"

			user, _ := service.Update(
				1,
				object.NewUserInput(
					name,
					email,
					"test",
				),
			)

			assert.Equal(t, name, user.Name)
		},
		func(
			db pkg.Database,
			service service.UserService,
		) {
			users, paginator, _ := service.Search(
				object.NewUserSearchInput(
					"",
					0,
				),
				20,
			)

			assert.Equal(t, 1, len(users))
			assert.Equal(t, 1, paginator.TotalCount)
		},
		func(
			db pkg.Database,
			service service.UserService,
		) {

			err := service.Delete(1)

			assert.Equal(t, nil, err)
		},
	)
}
