package log_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/plugins/log_stash"
	"gvd_server/service/common/res"
)

// LogReadView 日志读取 前端点击之后 后端就把日志的状态改为已读
// @Tags 日志管理
// @Summary 日志列表
// @Description 日志列表
// @Param data query models.IDRequest true "参数"
// @Param token header string true "token"
// @Router /api/logs/read [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (LogApi) LogReadView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}
	var log log_stash.LogModel
	err = global.DB.Take(&log, cr.ID).Error
	if err != nil {
		res.FailWithMsg("日志不存在", c)
		return
	}
	if log.ReadStatus {
		res.OKWithMsg("日志读取成功", c)
		return
	}

	//更新状态
	global.DB.Model(&log).Update("readStatus", true)
	res.OKWithMsg("日志读取成功", c)
}