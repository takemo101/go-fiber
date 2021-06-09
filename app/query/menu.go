package query

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/object"
	"github.com/takemo101/go-fiber/pkg"
)

// MenuQuery database structure
type MenuQuery struct {
	db pkg.Database
}

// NewMenuQuery creates a new menu query
func NewMenuQuery(db pkg.Database) MenuQuery {
	return MenuQuery{
		db: db,
	}
}

// Search get menus
func (r MenuQuery) Search(object object.MenuSearchInput, limit int) (menus []model.Menu, err error) {
	return menus, r.db.GormDB.
		Preload("Tags").
		Preload("Category").
		Preload("User").
		Order("id desc").
		Limit(limit).
		Find(&menus).Error
}

// GetUpdateMenus get menus order by update_at
func (r MenuQuery) GetUpdateMenus(limit int) (menus []model.Menu, err error) {
	return menus, r.db.GormDB.Order("updated_at desc").Limit(limit).Find(&menus).Error
}
