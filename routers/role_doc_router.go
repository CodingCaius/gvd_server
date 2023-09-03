package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) RoleDocRouter() {
	app := api.App.RoleDocApi
	r := router.Group("role_docs").Use(middleware.JwtAdmin())
	r.POST("", app.RoleDocCreateView)        // 角色-文档 添加
	r.DELETE("", app.RoleDocRemoveView)      // 角色-文档 删除
	r.PUT("", app.RoleDocUpdateView)         // 角色-文档 批量更新
	r.GET(":id", app.RoleDocListView)        // 角色-文档列表
	r.GET("info", app.RoleDocInfoView)       // 角色-文档 信息
	r.PUT("info", app.RoleDocInfoUpdateView) // 角色-文档 信息更新
}
