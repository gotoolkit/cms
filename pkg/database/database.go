package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	engine *xorm.Engine
	err    error
)

// Setup creates a connection to mysql database and migrates any new Models
func Setup(source string) error {
	engine, err = xorm.NewEngine("mysql", source)
	return err
}

// GetDB ...
func GetDB() *xorm.Engine {
	return engine
}

// CloseDB ...
func CloseDB() error {
	return engine.Close()
}
