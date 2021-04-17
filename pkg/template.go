package pkg

import (
	"strings"

	"github.com/gofiber/template/django"
)

// TemplateEngine is view engine
type TemplateEngine struct {
	Engine *django.Engine
}

// NewTemplateEngine set up template
func NewTemplateEngine(config Config) TemplateEngine {
	engine := django.New(config.Template.Path, config.Template.Suffix)
	engine.Reload(config.Template.Reload)
	engine.Debug(config.App.Debug)

	engine.AddFunc("nl2br", func(value interface{}) string {
		if str, ok := value.(string); ok {
			return strings.Replace(str, "\n", "<br />", -1)
		}
		return ""
	})

	return TemplateEngine{
		Engine: engine,
	}
}
