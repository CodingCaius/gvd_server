package role_doc_api

import (
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"

	"github.com/gin-gonic/gin"
)

type RoleDocInfoUpdateRequest struct {
	RoleDocRequest
	IsPwd       bool   `json:"isPwd"`       // 是否开启密码
	RoleDocPwd  string `json:"roleDocPwd"`  // 角色文档的密码
	IsSee       bool   `json:"isSee"`       // 是否开启了试看
	FreeContent string `json:"freeContent"` // 文档的试看内容
}

// RoleDocInfoUpdateView 角色文档信息 更新
// @Tags 角色文档管理
// @Summary 角色文档信息 更新
// @Description 角色文档信息 更新
// @Param token header string true "token"
// @Param data body RoleDocInfoUpdateRequest true "参数"
// @Router /api/role_docs/info [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (RoleDocApi) RoleDocInfoUpdateView(c *gin.Context) {
	var cr RoleDocInfoUpdateRequest
	err := c.ShouldBindJSON(&cr)
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

	// 判断密码有没有修改，修改之后需求把文档-密码表 对应文档的数据清空
	if !(roleDoc.Pwd != nil && *roleDoc.Pwd == cr.RoleDocPwd) {
		// 不同
		var userPwdDocs []models.UserPwdDocModel
		global.DB.Find(&userPwdDocs, "doc_id = ?", cr.DocID).Delete(&userPwdDocs)
	}

	/*
		IsPwd true
		RoleDocPwd  123
	*/

	// 用 map 来更新，防止传入的结构体是空值，更新就会失效
	// 不用map 的话，如果本来有密码，改成没有密码，更新会失败
	var roleDocInfo = map[string]any{
		"pwd":         nil,
		"freeContent": nil,
	}
	if cr.IsPwd {
		roleDocInfo["pwd"] = &cr.RoleDocPwd
	}
	if cr.IsSee {
		roleDocInfo["freeContent"] = &cr.FreeContent
	}

	// map空值的字段也会更新
	global.DB.Model(&roleDoc).Updates(roleDocInfo)

	res.OKWithMsg("文档更新成功", c)

}
