package query

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/object"
	"github.com/takemo101/go-fiber/pkg"
)

// RequestQuery database structure
type RequestQuery struct {
	db pkg.Database
}

// NewRequestQuery creates a new request query
func NewRequestQuery(db pkg.Database) RequestQuery {
	return RequestQuery{
		db: db,
	}
}

// Search get requests
func (r RequestQuery) Search(object object.RequestSearchInput, limit int) (requests []model.Request, paginator Paginator, err error) {
	db := r.db.GormDB.
		Preload("Tags").
		Preload("Category").
		Preload("User")

	err = Paging(&PagingParameter{
		DB:      db,
		Page:    object.GetPage(),
		Limit:   limit,
		OrderBy: []string{"id desc"},
	}, &requests, &paginator)

	return requests, paginator, err

}

// GetUpdateRequests get requests order by update_at
func (r RequestQuery) GetUpdateRequests(limit int) (requests []model.Request, err error) {
	return requests, r.db.GormDB.Order("updated_at desc").Limit(limit).Find(&requests).Error
}
