package site_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/service/common/res"
)

// SiteDetailView 站点配置查询
// @Tags 站点配置
// @Summary 站点配置查询
// @Description 站点配置查询
// @Router /api/site [get]
// @Produce json
// @Success 200 {object} res.Response{data=config.Site}
func (SiteApi) SiteDetailView(c *gin.Context) {
	res.OKWithData(global.Config.Site, c)
}