package role_doc_api

import (
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"

	"github.com/gin-gonic/gin"
)

// RoleDocRemoveView 删除一篇角色文档
// @Tags 角色文档管理
// @Summary 删除一篇角色文档
// @Description 删除一篇角色文档
// @Param token header string true "token"
// @Param data body RoleDocRequest true "参数"
// @Router /api/role_docs [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (RoleDocApi) RoleDocRemoveView(c *gin.Context) {
	var cr RoleDocRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		//res.FailWithError(err, c)
		return
	}

	var roleDoc models.RoleDocModel
	err = global.DB.Take(&roleDoc, "role_id = ? and doc_id = ?", cr.RoleID, cr.DocID).Error
	if err != nil {
		res.FailWithMsg("不存在的文档", c)
		return
	}
	global.DB.Delete(&roleDoc)
	res.OKWithMsg("删除成功", c)

}
