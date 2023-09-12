
// 文档更新时， 需要先获取完整的文章

package doc_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
)

type DocItemResponse struct {
	Content string `json:"content"`
	Title   string `json:"title"`
}


// DocEditContentView 获取完整的正文
// @Tags 文档管理
// @Summary 获取完整的正文
// @Description 获取完整的正文
// @Param id path int true "id"
// @Router /api/docs/edit/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{data=DocItemResponse}
func (DocApi) DocEditContentView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}
	var doc models.DocModel
	err = global.DB.Take(&doc, cr.ID).Error
	if err != nil {
		res.FailWithMsg("文档不存在", c)
		return
	}

	
	res.OKWithData(DocItemResponse{
		Title:   doc.Title,
		Content: doc.Content,
	}, c)
}