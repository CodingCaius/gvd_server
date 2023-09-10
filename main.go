package main

import (
	"gvd_server/core"
	_ "gvd_server/docs"
	"gvd_server/flags"
	"gvd_server/global"
	"gvd_server/routers"
	"gvd_server/service/cron_service"
)

// @title 文档项目api文档
// @version 1.0
// @description API文档
// @host 127.0.0.1:8000
// @BasePath /
func main() {
	//读取配置文件
	global.Config = core.InitConfig()
	//初始化连接
	global.Log = core.InitLogger()
	global.DB = core.InitMysql()
	global.Redis = core.InitRedis(0)
	global.ESClient = core.InitEs()

	option := flags.Parse()
	if option.Run() {
		return
	}

	//定时任务，同步文档数据
	cron_service.CornInit()

	router := routers.Routers()
	addr := global.Config.System.Addr()
	router.Run(addr)
}
