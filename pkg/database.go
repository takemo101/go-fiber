package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database is model
type Database struct {
	GormDB *gorm.DB
}

// DatabaseCreator is databse create
type DatabaseCreator struct {
	config DB
}

// CreateDialector is create gorm dialector
func (creator *DatabaseCreator) CreateDialector() gorm.Dialector {
	var dialector gorm.Dialector
	dbtype := strings.ToLower(creator.config.Type)

	switch dbtype {
	case "mysql":
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=%s",
			creator.config.User,
			creator.config.Pass,
			creator.config.Host,
			creator.config.Port,
			creator.config.Name,
			creator.config.Charset,
		)
		dialector = mysql.Open(dsn)
	default:
		dialector = sqlite.Open(creator.config.Name)
	}

	return dialector
}

// CreateDatabase is create gorm database
func (creator *DatabaseCreator) CreateDatabase(config *gorm.Config) (*gorm.DB, error) {
	return gorm.Open(creator.CreateDialector(), config)

}

// DB is create database/sql.DB
func (db *Database) DB() (*sql.DB, error) {
	return db.GormDB.DB()

}

// NewDatabase is create database
func NewDatabase(config Config) Database {
	creator := DatabaseCreator{
		config: config.DB,
	}

	db, err := creator.CreateDatabase(&gorm.Config{})

	if err != nil {
		log.Fatalf("database connection failed : %v", err)
	}

	return Database{
		GormDB: db,
	}
}
