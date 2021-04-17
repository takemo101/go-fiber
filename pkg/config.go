package pkg

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/configor"
)

// App config
type App struct {
	Name    string
	Host    string
	Port    int
	Version string
	Env     string
	Secret  string
	Debug   bool
}

// DB config
type DB struct {
	Type      string
	Host      string
	Port      int
	Name      string
	User      string
	Pass      string
	Charset   string
	Collation string
}

// Server fiber config
type Server struct {
	Prefork     bool
	Strict      bool
	Case        bool
	Etag        bool
	BodyLimit   int
	Concurrency int
	Timeout     struct {
		Read  time.Duration
		Write time.Duration
		Idel  time.Duration
	}
	Buffer struct {
		Read  int
		Write int
	}
}

// Log config
type Log struct {
	Server string
}

// Static config
type Static struct {
	Prefix string
	Root   string
	Index  string
}

// Template config
type Template struct {
	Path   string
	Suffix string
	Reload bool
}

// Cache config
type Cache struct {
	Expiration time.Duration
	Control    bool
}

// Session is config
type Session struct {
	Expiration time.Duration
	Name       string
	Domain     string
	Path       string
	Secure     bool
	HTTPOnly   bool
}

// Cors is config
type Cors struct {
	Origins []string
	MaxAge  time.Duration
}

// Config full config
type Config struct {
	App
	DB
	Server
	Log
	Static
	Template
	Cache
	Session
	Cors
}

func (app App) SecretKey() []byte {
	return []byte(app.Secret)
}

// Conf is static
var Conf = Config{}

// NewConfig create configure
func NewConfig() Config {
	// config.yml
	err := configor.Load(&Conf, "config.yml")
	if err != nil {
		log.Fatalf("fail to load config.yml : %v", err)
	}

	if Conf.App.Debug {
		fmt.Println(Conf)
	}
	return Conf
}
