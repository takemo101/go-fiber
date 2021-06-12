package query

import (
	"github.com/takemo101/go-fiber/app/model"
	"github.com/takemo101/go-fiber/pkg"
)

// SuggestQuery database structure
type SuggestQuery struct {
	db pkg.Database
}

// NewSuggestQuery creates a new suggest query
func NewSuggestQuery(db pkg.Database) SuggestQuery {
	return SuggestQuery{
		db: db,
	}
}

// GetUpdateSuggests get suggests order by update_at
func (r SuggestQuery) GetUpdateSuggests(limit int) (discussions []model.Suggest, err error) {
	return discussions, r.db.GormDB.
		Preload("Suggester").
		Preload("Request.User").
		Order("updated_at desc").
		Limit(limit).
		Find(&discussions).Error
}
