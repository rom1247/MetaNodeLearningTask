package database

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(cfg *Config) *gorm.DB {

	var db *gorm.DB
	var err error

	switch cfg.DBType {
	case SQLite:
		db, err = gorm.Open(sqlite.Open(cfg.SqlitePath), &gorm.Config{})
	case MySQL:
		db, err = gorm.Open(mysql.Open(cfg.MySqlDsn), &gorm.Config{})

	default:
		db, err = nil, fmt.Errorf("unknown db type: %s", cfg.DBType)
	}

	if err != nil {
		panic(err)
	}

	return db

}
