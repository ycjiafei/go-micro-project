package dao

import (
	"github.com/ycjiafei/go-micro-project/database"
	"github.com/ycjiafei/go-micro-project/database/structs"
	"log"
)

func GetUserInfoByUid(uid int64) structs.UserInfo {
	db, err := database.NewMysqlConn()
	if err != nil {
		log.Printf("连接数据库失败, err: %v", err)
		return structs.UserInfo{}
	}
	defer db.Close()
	user := structs.UserInfo{}
	db.Find(&user, uid)
	return user
}


func GetUserInfoByPhone(phone string) structs.UserInfo {
	db, err := database.NewMysqlConn()
	if err != nil {
		log.Printf("连接数据库失败, err: %v", err)
		return structs.UserInfo{}
	}
	defer db.Close()
	user := structs.UserInfo{}
	db.Where("phone = ?", phone).First(&user)
	return user
}