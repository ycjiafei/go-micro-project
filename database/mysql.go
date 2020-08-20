package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewMysqlConn() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "test.db")
	return db, err
}