package dao

import (
	"github.com/ycjiafei/go-micro-project/database"
	"github.com/ycjiafei/go-micro-project/database/structs"
)

func GetUserInfoByUid(uid int64) structs.UserInfo {
	db, err := database.NewMysqlConn()
	if err != nil {

	}
}
