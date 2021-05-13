package pkg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"runtime"
	"time"

	"github.com/jinzhu/configor"
)

const (
	Production  = "production"
	Development = "development"
	Local       = "local"
)

// App config
type App struct {
	Name    string
	Host    string
	Port    int
	URL     string
	Version string
	Env     string
	Secret  string
	Debug   bool
	Config  string
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

// SMTP config
type SMTP struct {
	Host       string
	Port       int
	Identity   string
	User       string
	Pass       string
	Encryption string
	From       struct {
		Address string
		Name    string
	}
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

// JWT is config
type JWT struct {
	Signing struct {
		Key    []byte
		Method string
	}
	Context struct {
		Key string
	}
	Lookup string
	Scheme string
}

// Config full config
type Config struct {
	App
	DB
	Server
	Log
	SMTP
	Static
	Template
	Cache
	Session
	Cors
	GoVersion     string
	ConfigMapData map[string]interface{}
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

	Conf.GoVersion = runtime.Version()

	if Conf.App.Env == "" {
		Conf.App.Env = Local
	}

	Conf.ConfigMapData = make(map[string]interface{})
	return Conf
}

// Load config json data
func (c *Config) Load(key string) (map[string]interface{}, error) {
	if mapValue, ok := c.ConfigMapData[key]; ok {
		return mapValue.(map[string]interface{}), nil
	}

	path := path.Join(c.App.Config, key+".json")
	jsonString, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var v interface{}
	err = json.Unmarshal(jsonString, &v)
	if err != nil {
		return nil, err
	}
	c.ConfigMapData[key] = v

	return v.(map[string]interface{}), nil
}
