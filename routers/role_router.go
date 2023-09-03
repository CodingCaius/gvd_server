package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) RoleRouter() {
	app := api.App.RoleApi

	//api路由组下在创建一个roles路由组，同意添加中间件
	r := router.Group("roles").Use(middleware.JwtAdmin())

	r.GET("", app.RoleListView)      // 角色列表
	r.POST("", app.RoleCreateView)   // 角色添加
	r.PUT("", app.RoleUpdateView)    // 角色更新
	r.DELETE("", app.RoleRemoveView) // 角色删除
}