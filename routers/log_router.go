package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) LogRouter() {
	app := api.App.LogApi
	r := router.Group("logs").Use(middleware.JwtAdmin())
	r.GET("", app.LogListView)     // 日志列表
	r.GET("read", app.LogReadView) // 日志读取
	r.DELETE("", app.LogRemoveView) // 日志删除
}