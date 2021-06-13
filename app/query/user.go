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
func (r UserQuery) Search(object object.UserSearchInput, limit int) (users []model.User, paginator Paginator, err error) {
	err = Paging(&PagingParameter{
		DB:      r.db.GormDB,
		Page:    object.GetPage(),
		Limit:   limit,
		OrderBy: []string{"id desc"},
	}, &users, &paginator)

	return users, paginator, err
}
