// 用户浏览的数据统计

package data_api

import (
	"fmt"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"time"

	"github.com/gin-gonic/gin"
)

type DataCountResponse struct {
	DateList  []string `json:"dateList"`
	CountList []int    `json:"countList"`
}


type DataCountRequest struct {
	Type int `json:"type" form:"type"` // 0 七天内  1 一个月   2 一年
}

type DataLookDateRequest struct {
	DataCountRequest
	DocID uint `json:"docID" form:"docID"` // 文档id，不传就是查全部的
}

// DataLookDateView 文档浏览量数据
// @Tags 数据统计
// @Summary 文档浏览量数据
// @Description 文档浏览量数据
// @Param data query DataLookDateRequest true "参数"
// @Router /api/data/look_date [get]
// @Produce json
// @Success 200 {object} res.Response{data=DataCountResponse}
func (DataApi) DataLookDateView(c *gin.Context) {
	// 声明一个DataLookDateRequest类型的变量cr
	var cr DataLookDateRequest
	// 获取请求体中的参数
	_ = c.ShouldBindQuery(&cr)

	// 初始化一个query变量
	var query = global.DB.Where("")
	// 获取当前时间
	now := time.Now()

	// 如果cr.DocID不为0
	if cr.DocID!= 0 {
		// 根据cr.DocID查询
		query.Where("docID =?", cr.DocID)
	}

	// 声明一个response变量
	var response DataCountResponse
	// 声明一个dateTypeNum变量
	var dateTypeNum int

	switch cr.Type {
	// 如果cr.Type等于1
	case 1:
		// 一个月
		dateTypeNum = 30
	// 如果cr.Type等于2
	case 2:
		// 一年
		dateTypeNum = 365
	// 其他情况
	default:
		// 7天内
		dateTypeNum = 7
	}
	// 根据dateTypeNum查询
	query.Where(fmt.Sprintf("date_sub(curdate(), interval %d day) <= createdAt", dateTypeNum))
	// 获取当前日期的前dateTypeNum天
	aDay := now.AddDate(0, 0, -dateTypeNum)
	// 将前dateTypeNum天的日期添加到response.DateList中
	for i := 1; i <= dateTypeNum; i++ {
		response.DateList = append(response.DateList, aDay.AddDate(0, 0, i).Format("2006-01-02"))
	}

	// 声明一个dateCountType类型的变量
	type dateCountType struct {
		Date  string `json:"date"`
		Count int    `json:"count"`
	}

	// 声明一个dateCountList变量
	var dateCountList []dateCountType
	// 查询指定日期的数据
	global.DB.Model(models.DocDataModel{}).Where(query).
		Select(
			"date_format(createdAt, '%Y-%m-%d') as date",
			"sum(lookCount) as count").
		Group("date").Scan(&dateCountList)

	// 初始化一个dateCountMap变量
	var dateCountMap = map[string]int{}

	// 将查询结果添加到dateCountMap中
	for _, countType := range dateCountList {
		dateCountMap[countType.Date] = countType.Count
	}

	// 将response.DateList中的日期添加到response.CountList中
	for _, s := range response.DateList {
		count := dateCountMap[s]
		response.CountList = append(response.CountList, count)
	}

	// 返回结果
	res.OKWithData(response, c)

}
