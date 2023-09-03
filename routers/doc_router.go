package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) DocRouter() {
	app := api.App.DocApi
	router.POST("docs", middleware.JwtAdmin(), app.DocCreateView)       // 创建文档
	router.GET("docs/info/:id", middleware.JwtAdmin(), app.DocInfoView) // 文档信息
	router.GET("docs/:id", app.DocContentView)                          // 文档内容
	router.GET("docs/edit/:id", middleware.JwtAdmin(), app.DocEditContentView) // 文档的完整内容
	router.GET("docs/digg/:id", app.DocDiggView) // 文档点赞
	router.POST("docs/pwd", app.DocPwdView)                             // 输入密码，查看文档
	router.DELETE("docs/:id", middleware.JwtAdmin(), app.DocRemoveView) // 删除文档
}

