package routers

import (
	"gvd_server/api"
)

func (router RouterGroup) SiteRouter() {
	app := api.App.SiteApi
	r := router.Group("site")
	r.GET("", app.SiteDetailView) // 站点配置的信息
	r.PUT("", app.SiteUpdateView) // 站点配置的更新
}