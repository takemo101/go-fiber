package query

import "github.com/takemo101/go-fiber/pkg"

// AdminQuery database structure
type AdminQuery struct {
	db     pkg.Database
	logger pkg.Logger
}

// NewAdminQuery creates a new admin query
func NewAdminQuery(db pkg.Database, logger pkg.Logger) AdminQuery {
	return AdminQuery{
		db:     db,
		logger: logger,
	}
}
