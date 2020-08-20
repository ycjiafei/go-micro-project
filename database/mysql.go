package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

func NewMysqlConn() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("GO_MICRO_DB_USER"),
		os.Getenv("GO_MICRO_DB_PASSWORD"),
		os.Getenv("GO_MICRO_DB_HOST"),
		os.Getenv("GO_MICRO_DB_NAME"),
	))
	return db, err
}