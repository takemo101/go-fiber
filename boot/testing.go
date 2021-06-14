package boot

import (
	"testing"

	"github.com/takemo101/go-fiber/app"
	"github.com/takemo101/go-fiber/pkg"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

// Testing test func
func Testing(t *testing.T, tests ...interface{}) {
	pkg.ConfigPath = "../config.testing.yml"

	fxtest.New(
		t,
		pkg.Module,
		app.Module,
		fx.Invoke(tests...),
	).Done()
}
