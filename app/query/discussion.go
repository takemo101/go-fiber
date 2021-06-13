package query

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/app/object"
	"github.com/takemo101/go-fiber/pkg"
)

// DiscussionQuery database structure
type DiscussionQuery struct {
	db pkg.Database
}

// NewDiscussionQuery creates a new discussion query
func NewDiscussionQuery(db pkg.Database) DiscussionQuery {
	return DiscussionQuery{
		db: db,
	}
}

// Search get discussions
func (r DiscussionQuery) Search(object object.DiscussionSearchInput, limit int) (discussions []model.Discussion, paginator Paginator, err error) {
	db := r.db.GormDB.
		Preload("Suggest.Request").
		Preload("Sender")

	err = Paging(&PagingParameter{
		DB:      db,
		Page:    object.GetPage(),
		Limit:   limit,
		OrderBy: []string{"id desc"},
	}, &discussions, &paginator)

	return discussions, paginator, err
}

// GetUpdateDiscussions get discussions order by update_at
func (r DiscussionQuery) GetUpdateDiscussions(limit int) (discussions []model.Discussion, err error) {
	return discussions, r.db.GormDB.
		Preload("Suggest.Request").
		Preload("Sender").
		Order("updated_at desc").
		Limit(limit).
		Find(&discussions).Error
}
