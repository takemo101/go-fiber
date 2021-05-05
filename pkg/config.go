package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"time"

	"github.com/jinzhu/configor"
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
	MapData map[string]interface{}
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
	Conf.MapData = make(map[string]interface{})
	return Conf
}

func (c *Config) Load(key string) (map[string]interface{}, error) {
	if mapValue, ok := c.MapData[key]; ok {
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
	c.MapData[key] = v

	return v.(map[string]interface{}), nil
}
