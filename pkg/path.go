package pkg

import (
	"fmt"
	"path"
)

// path is path helper
type Path struct {
	config Config
}

func NewPath(
	config Config,
) Path {
	return Path{
		config: config,
	}
}

func (p Path) Static(suffix string) string {
	return p.URL(path.Join(p.config.Static.Prefix, suffix))
}

func (p Path) URL(suffix string, a ...interface{}) string {
	return p.config.App.URL + "/" + fmt.Sprintf(suffix, a...)
}
