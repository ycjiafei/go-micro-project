package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ycjiafei/go-micro-project/database/structs"
	"os"
)

/**
搞成单例更好用
 */
func NewMysqlConn() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("GO_MICRO_DB_USER"),
		os.Getenv("GO_MICRO_DB_PASSWORD"),
		os.Getenv("GO_MICRO_DB_HOST"),
		os.Getenv("GO_MICRO_DB_NAME"),
	))
	autoMigrateTables(db)
	return db, err
}

/**
连接时候自动建表
 */
func autoMigrateTables(db *gorm.DB) {
	tables := []interface{}{
		&structs.UserInfo{},
	}
	db.AutoMigrate(tables...)
}