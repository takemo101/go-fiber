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

func Test_AdminService_OK(t *testing.T) {
	boot.Testing(
		func(
			db pkg.Database,
			service service.AdminService,
		) {
			db.GormDB.AutoMigrate(database.Models...)

			name := "test"
			email := "test@example.com"
			role := string(model.RoleAdmin)

			admin, _ := service.Store(
				object.NewAdminInput(
					name,
					email,
					"test",
					role,
				),
			)

			assert.Equal(t, name, admin.Name)
		},
		func(
			db pkg.Database,
			service service.AdminService,
		) {
			name := "testing"
			email := "test@example.com"
			role := string(model.RoleSystem)

			admin, _ := service.Update(
				1,
				object.NewAdminInput(
					name,
					email,
					"test",
					role,
				),
			)

			assert.Equal(t, name, admin.Name)
		},
		func(
			db pkg.Database,
			service service.AdminService,
		) {
			admins, paginator, _ := service.Search(
				object.NewAdminSearchInput(
					"",
					0,
				),
				20,
			)

			assert.Equal(t, 1, len(admins))
			assert.Equal(t, 1, paginator.TotalCount)
		},
		func(
			db pkg.Database,
			service service.AdminService,
		) {

			err := service.Delete(1)

			assert.Equal(t, nil, err)
		},
	)
}
