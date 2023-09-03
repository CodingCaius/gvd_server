package log_api

import (
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/plugins/log_stash"
	"gvd_server/service/common/list"
	"gvd_server/service/common/res"
	"time"

	"github.com/gin-gonic/gin"
)

type LogListRequest struct {
	models.Pagination
	Level  log_stash.Level   `json:"level" form:"level"`   // 日志查询的等级
	Type   log_stash.LogType `json:"type" form:"type"`     // 日志的类型   1 登录日志  2 操作日志  3 运行日志
	IP     string            `json:"ip" form:"ip"`         // 根据ip查询
	UserID uint              `json:"userID" form:"userID"` // 根据用户id查询
	Addr   string            `json:"addr" form:"addr"`     // 感觉地址查询
	Date   string            `json:"date" form:"date"`     // 查某一天，格式是年-月-日
}

// LogListView 日志列表 可根据type level 搜索日志
// @Tags 日志管理
// @Summary 日志列表
// @Description 日志列表
// @Param data query LogListRequest true "参数"
// @Param token header string true "token"
// @Router /api/logs [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[log_stash.LogModel]}
func (LogApi) LogListView(c *gin.Context) {
	var cr LogListRequest
	c.ShouldBindQuery(&cr)

	var query = global.DB.Where("")
	if cr.Date != "" {
		_, dateTimeErr := time.Parse("2006-01-02", cr.Date)
		if dateTimeErr != nil {
			res.FailWithMsg("时间格式错误", c)
			return
		}
		query.Where("date(createdAt) = ?", cr.Date)
	}

	_list, count, _ := list.QueryList(log_stash.LogModel{
		Type:   cr.Type,
		Level:  cr.Level,
		IP:     cr.IP,
		UserID: cr.UserID,
		Addr:   cr.Addr,
	}, list.Option{
		Pagination: cr.Pagination,
		Where: query,
		Likes: []string{"title", "userName"},
	})
	res.OKWithList(_list, count, c)
}
