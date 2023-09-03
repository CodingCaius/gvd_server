//添加用户的路由

package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) UserRouter() {
	app := api.App.UserApi
	router.POST("login", app.UserLoginView)                         //用户登录
	router.POST("users", middleware.JwtAdmin(), app.UserCreateView) //创建用户
	router.PUT("users", middleware.JwtAdmin(), app.UserUpdateView)  //用户更新
	router.GET("users", middleware.JwtAdmin(), app.UserListView)    //用户列表
	router.DELETE("users", middleware.JwtAdmin(), app.UserRemoveView)    //用户删除
	router.GET("logout", middleware.JwtAuth(), app.UserLogoutView)    //用户注销
	router.GET("users_info", middleware.JwtAuth(), app.UserInfoView)    //用户注销
	//用户修改密码
	router.PUT("users_password", middleware.JwtAuth(), app.UserUpdatePasswordView)
	//用户修改信息
	router.PUT("users_info", middleware.JwtAuth(), app.UserUpdateInfoView)


}
