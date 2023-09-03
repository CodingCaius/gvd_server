package doc_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/plugins/log_stash"
	"gvd_server/service/common/res"
)

type DocUpdateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// DocUpdateView 更新文档
// @Tags 文档管理
// @Summary 更新文档
// @Description 更新文档，更新文档的标题和正文
// @Param data body DocUpdateRequest true "参数"
// @Param token header string true "token"
// @Param id path int true "id"
// @Router /api/docs/{id} [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (DocApi) DocUpdateView(c *gin.Context) {
	var cr DocUpdateRequest
	log := log_stash.NewAction(c)
	log.SetRequest(c)
	log.SetResponse(c)

	log.Info("更新文档")

	var id models.IDRequest
	c.ShouldBindUri(&id)
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		//res.FailWithError(err, c)
		return
	}
	if cr.Title == "" && cr.Content == "" {
		res.OKWithMsg("文档更新成功", c)
		return
	}

	var doc models.DocModel
	err = global.DB.Take(&doc, id.ID).Error
	if err != nil {
		res.FailWithMsg("文档不存在", c)
		return
	}
	err = global.DB.Model(&doc).Updates(models.DocModel{
		Title:   cr.Title,
		Content: cr.Content,
	}).Error
	if err != nil {
		log.SetItemErr("失败原因", err.Error())
		log.Error("文档更新失败")
		res.FailWithMsg("文档更新失败", c)
		return
	}
	res.OKWithMsg("文档更新成功", c)
}