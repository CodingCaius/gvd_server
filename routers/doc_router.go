package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) DocRouter() {
	app := api.App.DocApi

	r := router.Group("docs")

	r.POST("", middleware.JwtAdmin(), app.DocCreateView)             // 创建文档
	r.PUT(":id", middleware.JwtAdmin(), app.DocUpdateView)           // 更新文档
	r.GET("info/:id", middleware.JwtAdmin(), app.DocInfoView)        // 文档信息
	r.GET(":id", app.DocContentView)                                 // 文档内容
	r.GET("edit/:id", middleware.JwtAdmin(), app.DocEditContentView) // 文档的完整内容
	r.POST("pwd", app.DocPwdView)                                    // 输入密码，查看文档
	r.GET("digg/:id", app.DocDiggView)                               // 文档点赞
	r.DELETE(":id", middleware.JwtAdmin(), app.DocRemoveView)        // 删除文档
	r.GET("search", app.DocSearchView)                               // 全文搜索
}
