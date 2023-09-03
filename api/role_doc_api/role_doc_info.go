package role_doc_api

import (
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"

	"github.com/gin-gonic/gin"
)

type RoleDocInfoResponse struct {
	IsPwd       bool   `json:"isPwd"`       // 是否开启密码
	RoleDocPwd  string `json:"roleDocPwd"`  // 角色文档的密码
	RolePwd     string `json:"rolePwd"`     // 角色的密码
	IsSee       bool   `json:"isSee"`       // 是否开启了试看
	FreeContent string `json:"freeContent"` // 文档的试看内容
}

// RoleDocInfoView 角色文档信息
// @Tags 角色文档管理
// @Summary 角色文档信息
// @Description 角色文档信息
// @Param token header string true "token"
// @Param data query RoleDocRequest true "参数"
// @Router /api/role_docs/info [get]
// @Produce json
// @Success 200 {object} res.Response{data=RoleDocInfoResponse}
func (RoleDocApi) RoleDocInfoView(c *gin.Context) {
	var cr RoleDocRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		//res.FailWithError(err, c)
		return
	}

	var roleDoc models.RoleDocModel
	err = global.DB.Preload("RoleModel").Take(&roleDoc, "role_id = ? and doc_id = ?", cr.RoleID, cr.DocID).Error
	if err != nil {
		res.FailWithMsg("文档不存在", c)
		return
	}

	// 角色密码
	response := RoleDocInfoResponse{
		RolePwd: roleDoc.RoleModel.Pwd,
	}
	// 角色文档密码
	if roleDoc.Pwd != nil {
		response.IsPwd = true
		response.RoleDocPwd = *roleDoc.Pwd
	}
	if roleDoc.FreeContent != nil {
		response.IsSee = true
		response.FreeContent = *roleDoc.FreeContent
	}

	res.OKWithData(response, c)

}
