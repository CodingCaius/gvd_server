package core

import (
	"github.com/cc14514/go-geoip2"
	geoip2db "github.com/cc14514/go-geoip2-db"
	"github.com/sirupsen/logrus"
)

func InitAddrDB() *geoip2.DBReader {
	// 创建一个geoip2.DBReader对象
	db, err := geoip2db.NewGeoipDbByStatik()
	// 如果加载失败，则输出错误信息
	if err!= nil {
		logrus.Fatal("ip地址数据库加载失败", err)
	}
	// 返回db
	return db
}