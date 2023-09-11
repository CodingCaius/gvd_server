package data_api

import (
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"

	"github.com/gin-gonic/gin"
)

type DataSumResponse struct {
	UserCount int `json:"userCount"` // 总用户数
	DocCount  int `json:"docCount"`  // 总文档树
	DiggCount int `json:"diggCount"` // 总点赞数
	LookCount int `json:"lookCount"` // 总浏览量
}

// DataSumApiView 首页的求和数据
// @Tags 数据统计
// @Summary 首页的求和数据
// @Description 首页的求和数据
// @Router /api/data/sum [get]
// @Produce json
// @Success 200 {object} res.Response{data=DataSumResponse}
func (DataApi) DataSumApiView(c *gin.Context) {

	var response DataSumResponse
	global.DB.Model(models.UserModel{}).Select("count(id)").Scan(&response.UserCount)
	global.DB.Model(models.DocModel{}).Select("count(id)").Scan(&response.DocCount)
	global.DB.Model(models.DocModel{}).Select("sum(lookCount)").Scan(&response.LookCount)
	global.DB.Model(models.DocModel{}).Select("sum(diggCount)").Scan(&response.DiggCount)

	res.OKWithData(response, c)
}
