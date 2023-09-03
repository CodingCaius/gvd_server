//作为router文件夹的入口
//在这里进行 路由分组 以及 路由添加

//但 不同模块的路由添加由其他文件实现


package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
  	gs "github.com/swaggo/gin-swagger"
	"gvd_server/middleware"
)

type RouterGroup struct {
	*gin.RouterGroup
}

// Routers 函数用于创建 Gin 的路由引擎（*gin.Engine）并配置路由
func Routers() *gin.Engine {
	//创建了一个默认的 Gin 路由引擎
	router := gin.Default()

	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	//创建api路由组
	apiGroup := router.Group("api")
	apiGroup.Use(middleware.LogMiddleWare())
	routerGroup := RouterGroup{apiGroup}

	//线上如果有nginx,可以做反向代理，这一步可以省略
	//第一个参数是web的访问别名  第二个参数是内部的映射目录
	router.Static("/uploads", "uploads")

	//定义具体的路由
	routerGroup.UserRouter()
	routerGroup.ImageRouter()
	routerGroup.LogRouter()
	routerGroup.SiteRouter()
	routerGroup.RoleRouter()
	routerGroup.DocRouter()
	routerGroup.RoleDocRouter() // 角色文档路由

	return router
}
