package query

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/object"
	"github.com/takemo101/go-fiber/pkg"
)

// AdminQuery database structure
type AdminQuery struct {
	db pkg.Database
}

// NewAdminQuery creates a new admin query
func NewAdminQuery(db pkg.Database) AdminQuery {
	return AdminQuery{
		db: db,
	}
}

// Search get admins
func (r AdminQuery) Search(object object.AdminSearchInput, limit int) (admins []model.Admin, paginator Paginator, err error) {
	err = Paging(&PagingParameter{
		DB:      r.db.GormDB,
		Page:    object.GetPage(),
		Limit:   limit,
		OrderBy: []string{"id desc"},
	}, &admins, &paginator)

	return admins, paginator, err
}
