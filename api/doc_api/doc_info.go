package doc_api

import (
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/service/redis_service"

	"github.com/gin-gonic/gin"
)

type DocInfoResponse struct {
	models.Model
	Title         string `json:"title"`
	ContentLength int    `json:"contentLength"` // 正文内容
	DiggCount     int    `json:"diggCount"`
	LookCount     int    `json:"lookCount"`
	Key           string `json:"key"`
}

// DocInfoView 文档信息
// @Tags 文档管理
// @Summary 文档基础信息
// @Description 文档基础信息
// @Param id path int true "id"
// @Param token header string true "token"
// @Router /api/docs/info/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{data=DocInfoResponse}
func (DocApi) DocInfoView(c *gin.Context) {
	var cr models.IDRequest
	// ShouldBindUri 用来捕获 URL 路径中的参数
	// 例如：/api/docs/info/1
	// ShouldBindHeader 用来捕获请求头中的参数
	c.ShouldBindUri(&cr)

	var doc models.DocModel
	err := global.DB.Take(&doc, cr.ID).Error
	if err != nil {
		res.FailWithMsg("文档不存在", c)
		return
	}

	docDigg := redis_service.NewDocDigg().GetById(doc.ID)
	docLook := redis_service.NewDocLook().GetById(doc.ID)


	var docInfo = DocInfoResponse{
		Model:         doc.Model,
		Title:         doc.Title,
		ContentLength: len(doc.Content),
		DiggCount:     docDigg + doc.DiggCount,
		LookCount:     docLook + doc.LookCount,
		Key:           doc.Key,
	}

	res.OKWithData(docInfo, c)

}