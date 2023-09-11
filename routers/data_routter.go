package routers

import (
	"gvd_server/api"
)

func (router RouterGroup) DataRouter() {
	app := api.App.DataApi
	router.GET("data/sum", app.DataSumApiView)         // 求和数据
	router.GET("data/look_date", app.DataLookDateView) // 浏览量时间统计
}
