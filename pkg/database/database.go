package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gotoolkit/cms/pkg/models"
	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

// Setup creates a connection to mysql database and migrates any new Models
func Setup(source string) error {
	db, err = gorm.Open("mysql", source)
	db.AutoMigrate(&models.User{}, &models.Phone{})
	return err
}

// GetDB ...
func GetDB() *gorm.DB {
	return db
}

// CloseDB ...
func CloseDB() error {
	return db.Close()
}
