package database

import "github.com/takemo101/go-fiber/app/model"

// Models is slice interface{}
var Models = []interface{}{
	&model.Admin{},
	&model.User{},
	&model.Todo{},
	&model.Category{},
	&model.Tag{},
	&model.Menu{},
	&model.MenuTag{},
}
