package query

import (
	"github.com/takemo101/go-fiber/app/form"
	"github.com/takemo101/go-fiber/app/model"
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

// Search gets limit users
func (r UserQuery) Search(form form.UserSearch, limit int) (users []model.User, err error) {
	return users, r.db.GormDB.Order("id desc").Limit(limit).Find(&users).Error
}
