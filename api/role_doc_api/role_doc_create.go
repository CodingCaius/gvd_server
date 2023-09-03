package role_doc_api

import (
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"

	"github.com/gin-gonic/gin"
)

type RoleDocRequest struct {
	// 绑定query时，对应标签是 form， 绑定form表单数据也是用 form标签
	RoleID uint `json:"roleID" form:"roleID" binding:"required" label:"角色id"`
	DocID  uint `json:"docID" form:"docID" binding:"required" label:"文档id"`
}

// RoleDocCreateView 添加一篇角色文档
// @Tags 角色文档管理
// @Summary 添加一篇角色文档
// @Description 添加一篇角色文档
// @Param token header string true "token"
// @Param data body RoleDocRequest true "参数"
// @Router /api/role_docs [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (RoleDocApi) RoleDocCreateView(c *gin.Context) {
	var cr RoleDocRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		//res.FailWithError(err, c)
		return
	}

	var roleDoc models.RoleDocModel
	err = global.DB.Take(&roleDoc, "role_id = ? and doc_id = ?", cr.RoleID, cr.DocID).Error
	if err == nil {
		res.FailWithMsg("已存在", c)
		return
	}
	global.DB.Create(&models.RoleDocModel{
		RoleID: cr.RoleID,
		DocID:  cr.DocID,
	})

	res.OKWithMsg("添加成功", c)

}
