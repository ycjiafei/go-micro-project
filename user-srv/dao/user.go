package dao

import (
	"github.com/ycjiafei/go-micro-project/database"
	"github.com/ycjiafei/go-micro-project/database/structs"
	slog "github.com/ycjiafei/go-micro-project/pkg/log"
	"go.uber.org/zap"
)

func GetUserInfoByUid(uid int64) structs.UserInfo {
	db, err := database.NewMysqlConn()
	if err != nil {
		slog.SLog.Bg().Error("连接数据库失败", zap.String("err", err.Error()))
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
		slog.SLog.Bg().Error("连接数据库失败", zap.String("err", err.Error()))
		return structs.UserInfo{}
	}
	defer db.Close()
	user := structs.UserInfo{}
	db.Where("phone = ?", phone).First(&user)
	return user
}