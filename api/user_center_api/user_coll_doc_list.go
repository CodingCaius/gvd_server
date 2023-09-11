package user_center_api

import (
	"gvd_server/global"
	"gvd_server/models"
	list2 "gvd_server/service/common/list"
	"gvd_server/service/common/res"
	"gvd_server/utils/jwts"
	"time"

	"github.com/gin-gonic/gin"
)

type UserCollDocListResponse struct {
	CreatedAt     time.Time `json:"createdAt"`     // 什么时候收藏文档的
	DocID         uint      `json:"docID"`         // 文档id
	DocUpdateTime time.Time `json:"docUpdateTime"` // 文档的更新时间
	Title         string    `json:"title"`         // 文档标题
	IsPermission  bool      `json:"isPermission"`  // 是否有文档的访问权限
	LookCount     int       `json:"lookCount"`     // 浏览量
	DiggCount     int       `json:"diggCount"`     // 点赞量
}

// UserCollDocListView 收藏文档列表
// @Tags 个人中心
// @Summary 收藏文档列表
// @Description 收藏文档列表
// @Param token header string true "token"
// @Router /api/user_center/user_coll [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserCenterApi) UserCollDocListView(c *gin.Context) {
	token := c.Request.Header.Get("token")
	claims, _ := jwts.ParseToken(token)
	var cr models.Pagination
	c.ShouldBindQuery(&cr)
	//user, err := claims.GetUser()
	//if err != nil {
	//	res.FailWithMsg("用户信息错误", c)
	//	return
	//}
	_list, count, _ := list2.QueryList(models.UserCollDocModel{}, list2.Option{
		Pagination: cr,
		Preload:    []string{"DocModel", "UserModel"},
	})

	// 如何查询这个用户是否有这个文档的访问权限呢
	var docIDList []uint
	global.DB.Model(models.RoleDocModel{}).Where("role_id = ?", claims.RoleID).Select("doc_id").Scan(&docIDList)
	var docIDMap = map[uint]bool{}
	for _, u := range docIDList {
		docIDMap[u] = true
	}

	var list = make([]UserCollDocListResponse, 0)
	for _, model := range _list {
		list = append(list, UserCollDocListResponse{
			CreatedAt:     model.CreatedAt,
			DocID:         model.DocID,
			DocUpdateTime: model.DocModel.UpdatedAt,
			Title:         model.DocModel.Title,
			LookCount:     model.DocModel.LookCount,
			DiggCount:     model.DocModel.DiggCount,
			IsPermission:  docIDMap[model.DocID],
		})
	}
	res.OKWithList(list, count, c)
}
