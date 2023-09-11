package doc_api

import (
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/service/redis_service"
	"gvd_server/utils/jwts"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type DocContentResponse struct {
	Content   string    `json:"content"`
	IsSee     bool      `json:"isSee"`     // 是否试看
	IsPwd     bool      `json:"isPwd"`     // 是否需要密码
	IsColl    bool      `json:"isColl"`    // 用户是否收藏
	LookCount int       `json:"lookCount"` // 浏览量
	DiggCount int       `json:"diggCount"` // 点赞量
	CollCount int       `json:"collCount"` // 收藏量
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
}

// DocContentView 文档内容
// @Tags 文档管理
// @Summary 文档内容
// @Description 文档内容
// @Param id path int true "id"
// @Router /api/docs/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{data=DocContentResponse}
func (DocApi) DocContentView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindUri(&cr)
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
		Preload("DocModel.UserCollDocList").
		Preload("RoleModel").
		Take(&roleDoc, "role_id = ? and doc_id = ?", roleID, cr.ID).Error
	if err != nil {
		// 这个角色没有这个文档的权限
		res.FailWithMsg("文档鉴权失败", c)
		return
	}

	// 设置浏览量
	redis_service.NewDocLook().SetById(cr.ID)

	// 在角色文档表中拿到 文档
	doc := roleDoc.DocModel

	docDigg := redis_service.NewDocDigg().GetById(doc.ID)
	docLook := redis_service.NewDocLook().GetById(doc.ID)

	var response = DocContentResponse{
		DiggCount: docDigg + doc.DiggCount,
		LookCount: docLook + doc.LookCount,
		CollCount: len(doc.UserCollDocList),
		Title:     doc.Title,
		CreatedAt: doc.CreatedAt,
	}
	// IsSee 这个角色是不是对这个文档有试看
	// 正文分隔符
	// 文档里面的试看内容
	// 角色-文档的试看内容
	// 试看部分 优先级：角色文档试看  > 文档试看字段 > 文档按照特殊字符分隔的试看

	// 判断正文里面是不是有特殊分隔符
	isDocFree := strings.Contains(doc.Content, global.DocSplitSign)

	//如果角色有指定的试看内容（FreeContent），则根据来源的优先级更新试看内容：角色的试看内容、文档的试看内容和基于分隔符的试看内容。

	var freeContent string                                                 // 试看正文
	var content = strings.ReplaceAll(doc.Content, global.DocSplitSign, "") // 实际正文

	//如果正文里面有特殊分隔符
	if isDocFree {
		_list := strings.Split(doc.Content, global.DocSplitSign)
		freeContent = _list[0]
	}
	// 通过判断角色文档的 FreeContent是不是nil，如果不是，那么就开启了试看
	if roleDoc.FreeContent != nil {
		// 如果 FreeContent 为空，对应的优先级也都为空，也算它开启试看，试看内容 空
		// 在前端设置试看的时候，判断一下，有没有对应的试看内容，没有就要提示给用户
		response.IsSee = true

		// 按照优先级去设置试看
		if doc.FreeContent != "" {
			freeContent = doc.FreeContent
		}
		if *roleDoc.FreeContent != "" {
			freeContent = *roleDoc.FreeContent
		}

	}

	// 检查是否需要密码来访问文档。如果角色或文档指定了密码，响应中的 IsPwd 字段将设置为 true
	if roleDoc.Pwd != nil && (*roleDoc.Pwd != "" || roleDoc.RoleModel.Pwd != "") {
		response.IsPwd = true
	}

	// 如果用户不是访客（角色ID不是2），代码会检查用户是否已收藏文档（IsColl）。此外，还会检查用户是否选择免密码访问文档。
	if roleID != 2 {
		// 查用户是否收藏了文档
		var userDoc models.UserCollDocModel
		err = global.DB.Take(&userDoc, "doc_id = ? and user_id = ?", cr.ID, claims.UserID).Error
		if err == nil {
			response.IsColl = true
		}
		// 用户是否对这个文档免密
		var usePwd models.UserPwdDocModel
		err = global.DB.Take(&usePwd, "doc_id = ? and user_id = ?", cr.ID, claims.UserID).Error
		if err == nil {
			response.IsPwd = false
		}
	}

	// Content
	// 有密码。有试看  返回试看内容
	// 无密码，有试看  返回试看内容
	if response.IsSee {
		response.Content = freeContent
	}

	// 有密码，无试看  返回空
	if response.IsPwd && !response.IsSee {
		response.Content = ""
	}

	// 无密码，无试看  返回正文
	if !response.IsPwd && !response.IsSee {
		response.Content = content
	}

	res.OKWithData(response, c)
}
