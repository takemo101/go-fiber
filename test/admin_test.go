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

func Test_AdminService_OK(t *testing.T) {
	boot.Testing(
		t,
		func(
			db pkg.Database,
			service service.AdminService,
		) {
			db.GormDB.AutoMigrate(database.Models...)

			t.Run("admin store", func(t *testing.T) {
				name := "test"
				email := "test@example.com"
				role := string(model.RoleAdmin)

				for i := 0; i < 10; i++ {
					service.Store(
						object.NewAdminInput(
							name+strconv.Itoa(i),
							email+strconv.Itoa(i),
							"test",
							role,
						),
					)
				}

				admin, err := service.Find(1)

				assert.Equal(t, nil, err)
				assert.Equal(t, uint(1), admin.ID)
			})
			t.Run("admin update", func(t *testing.T) {
				name := "testing"
				email := "test@example.com"
				role := string(model.RoleSystem)

				service.Update(
					1,
					object.NewAdminInput(
						name,
						email,
						"test",
						role,
					),
				)

				findAdmin, _ := service.Find(1)

				assert.Equal(t, name, findAdmin.Name)
			})
			t.Run("admin index", func(t *testing.T) {
				admins, paginator, _ := service.Search(
					object.NewAdminSearchInput(
						"",
						0,
					),
					20,
				)

				assert.Equal(t, 10, len(admins))
				assert.Equal(t, 10, paginator.TotalCount)
			})
			t.Run("admin delete", func(t *testing.T) {
				deleteErr := service.Delete(1)
				_, findErr := service.Find(1)
				assert.Equal(t, nil, deleteErr)
				assert.NotEqual(t, nil, findErr)
			})
		},
	)
}
