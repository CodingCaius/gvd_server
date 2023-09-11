package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) UserCenterRouter() {
	app := api.App.UserCenterApi
	r := router.Group("user_center").Use(middleware.JwtAdmin())
	r.POST("user_coll", app.UserCollDocView)    // 收藏文档， 取消收藏
	//r.GET("user_coll", app.UserCollDocListView) // 收藏文档列表
}
