package query

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/object"
	"github.com/takemo101/go-fiber/pkg"
)

// UserQuery database structure
type UserQuery struct {
	db pkg.Database
}

// NewUserQuery creates a new user query
func NewUserQuery(db pkg.Database) UserQuery {
	return UserQuery{
		db: db,
	}
}

// Search get users limit
func (r UserQuery) Search(object object.UserSearchInput, limit int) (users []model.User, err error) {
	return users, r.db.GormDB.Order("id desc").Limit(limit).Find(&users).Error
}
