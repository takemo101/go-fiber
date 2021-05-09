package module

import "go.uber.org/fx"

// Module export
var Module = fx.Options(
	fx.Provide(NewModulePart),
)

// ModuleParts is slice
type ModuleParts []ModulePart

// Module is interface
type ModulePart interface {
	Setup()
}

// NewModule is setup modules
func NewModulePart() ModuleParts {
	return ModuleParts{}
}

// Boot all the module setup
func (modules ModuleParts) Boot() {
	for _, module := range modules {
		module.Setup()
	}
}
