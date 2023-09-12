
// 用于文档密码校验
// 输入密码 查看文档

package doc_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/utils/jwts"
	"strings"
)

type DocPwdRequest struct {
	Pwd   string `json:"pwd"`   // 密码
	DocID uint   `json:"docID"` // 文档id
}

// DocPwdView 输入密码，访问文档
// @Tags 文档管理
// @Summary 输入密码，访问文档
// @Description 输入密码，访问文档
// @Param data body DocPwdRequest true "参数"
// @Router /api/docs/pwd [post]
// @Produce json
// @Success 200 {object} res.Response{data=DocContentResponse}
func (DocApi) DocPwdView(c *gin.Context) {
	var cr DocPwdRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}
	// 因为这个接口，不登录也能访问，所以需要在视图里面解析token
	token := c.Request.Header.Get("token")
	claims, err := jwts.ParseToken(token)
	var roleID uint = 2 // 访客
	if err == nil {
		// 说明登录了
		roleID = claims.RoleID
	}

	// 判断角色是否有这个文档的访问权限
	var roleDoc models.RoleDocModel
	err = global.DB.
		Preload("RoleModel").
		Preload("DocModel").
		Take(&roleDoc, "role_id = ? and doc_id = ?", roleID, cr.DocID).Error
	if err != nil {
		// 这个角色没有这个文档的权限
		res.FailWithMsg("文档鉴权失败", c)
		return
	}

	// 按照优先级，获取这个文档的密码
	if roleDoc.Pwd == nil {
		res.FailWithMsg("无密码文档", c)
		return
	}


	// 先将密码设置为角色密码
	// 如果角色文档表有密码的话，赋值为角色文档表的密码
	pwd := roleDoc.RoleModel.Pwd
	if roleDoc.Pwd != nil && *roleDoc.Pwd != "" {
		pwd = *roleDoc.Pwd
	}

	if cr.Pwd != pwd {
		res.FailWithMsg("密码错误", c)
		return
	}
	var content = strings.ReplaceAll(roleDoc.DocModel.Content, global.DocSplitSign, "") // 实际正文

	
	if claims != nil {
		// 加入到用户密码文档表
		global.DB.Create(models.UserPwdDocModel{
			UserID: claims.UserID,
			DocID:  cr.DocID,
		})
	}

	var response = DocContentResponse{
		Content: content,
	}

	res.OKWithData(response, c)

}