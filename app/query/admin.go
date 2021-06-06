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

// Search gets limit admins
func (r AdminQuery) Search(object object.AdminSearchInput, limit int) (admins []model.Admin, err error) {
	return admins, r.db.GormDB.Order("id desc").Limit(limit).Find(&admins).Error
}
