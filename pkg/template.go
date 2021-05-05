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
func NewTemplateEngine(
	config Config,
	p Path,
) TemplateEngine {
	engine := django.New(config.Template.Path, config.Template.Suffix)
	engine.Reload(config.Template.Reload)
	engine.Debug(config.App.Debug)

	engine.AddFunc("nl2br", func(value interface{}) string {
		if str, ok := value.(string); ok {
			return strings.Replace(str, "\n", "<br />", -1)
		}
		return ""
	})

	engine.AddFunc("static", func(value interface{}) string {
		if str, ok := value.(string); ok {
			for _, prefix := range []string{"https://", "http://"} {
				if strings.Contains(str, prefix) {
					return str
				}
			}
			return p.Static(str)
		}
		return ""
	})

	engine.AddFunc("url", func(value interface{}, a ...interface{}) string {
		if str, ok := value.(string); ok {
			return p.URL(str, a...)
		}
		return ""
	})

	return TemplateEngine{
		Engine: engine,
	}
}
