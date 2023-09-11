package data_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"time"
)

// DataLoginDateView 用户登录数据
// @Tags 数据统计
// @Summary 用户登录数据
// @Description 用户登录数据
// @Param data query DataCountRequest true "参数"
// @Router /api/data/login_date [get]
// @Produce json
// @Success 200 {object} res.Response{data=DataCountResponse}
func (DataApi) DataLoginDateView(c *gin.Context) {
	var cr DataCountRequest
	_ = c.ShouldBindQuery(&cr)

	var query = global.DB.Where("")
	now := time.Now()

	var response DataCountResponse
	var dateTypeNum int

	switch cr.Type {
	case 1:
		// 一个月
		dateTypeNum = 30
	case 2:
		// 一年
		dateTypeNum = 365
	default:
		// 7天内
		dateTypeNum = 7
	}
	query.Where(fmt.Sprintf("date_sub(curdate(), interval %d day) <= createdAt", dateTypeNum))
	aDay := now.AddDate(0, 0, -dateTypeNum)
	for i := 1; i <= dateTypeNum; i++ {
		response.DateList = append(response.DateList, aDay.AddDate(0, 0, i).Format("2006-01-02"))
	}

	type dateCountType struct {
		Date  string `json:"date"`
		Count int    `json:"count"`
	}

	var dateCountList []dateCountType
	global.DB.Model(models.LoginModel{}).Where(query).
		Select(
			"date_format(createdAt, '%Y-%m-%d') as date",
			"count(id) as count").
		Group("date").Scan(&dateCountList)
	var dateCountMap = map[string]int{}
	for _, countType := range dateCountList {
		dateCountMap[countType.Date] = countType.Count
	}
	for _, s := range response.DateList {
		count, _ := dateCountMap[s]
		response.CountList = append(response.CountList, count)
	}

	res.OKWithData(response, c)

}