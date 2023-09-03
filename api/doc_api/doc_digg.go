package doc_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/service/redis_service"
)

// DocDiggView 文档点赞
// @Tags 文档管理
// @Summary 文档点赞
// @Description 文档点赞
// @Param id path int true "id"
// @Router /api/docs/digg/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (DocApi) DocDiggView(c *gin.Context) {
	var cr models.IDRequest
	c.ShouldBindUri(&cr)

	// 在缓存中设置
	redis_service.NewDocDigg().SetById(cr.ID)
	res.OKWithMsg("文档点赞成功", c)
}