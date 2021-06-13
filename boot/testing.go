package boot

import (
	"github.com/takemo101/go-fiber/app"
	"github.com/takemo101/go-fiber/pkg"
	"go.uber.org/fx"
)

// Testing test func
func Testing(tests ...interface{}) {
	pkg.ConfigPath = "../config.testing.yml"

	fx.New(
		pkg.Module,
		app.Module,
		fx.Invoke(tests...),
	).Done()
}
