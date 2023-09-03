//初始化mysql

package core

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvd_server/global"
	"time"
)

//使用GORM库初始化MySQL数据库连接，并返回一个数据库对象
func InitMysql() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		logrus.Warn("未配置mysql，取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface

	//基于全局配置中指定的日志级别，配置了GORM的日志记录器。
	//日志级别确定在数据库操作期间输出哪些类型的日志消息
	var logLevel = logger.Error
	switch global.Config.Mysql.LogLevel {
	case "info" :
		logLevel = logger.Info
	case "warn" :
		logLevel = logger.Warn
	}

	mysqlLogger = logger.Default.LogMode(logLevel)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		logrus.Fatalf(fmt.Sprintf("[%s] mysql 连接失败， err: %s", dsn, err.Error()))
	}
	//从 GORM 数据库对象 db 中获取与之关联的原生数据库连接对象 *sql.DB
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10) //最大空闲连接数
	sqlDB.SetMaxOpenConns(100) //最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) //连接最大复用时间，不能超过mysql的wait_timeout
	return db

}