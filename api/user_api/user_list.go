package user_api

import (
	"gvd_server/models"
	"gvd_server/service/common/list"
	"gvd_server/service/common/res"

	"github.com/gin-gonic/gin"
)

type UserListRequest struct {
	models.Pagination
	RoleID uint `json:"roleID" form:"roleID"`
}

// UserListView 用户列表
// @Tags 用户管理
// @Summary 用户列表
// @Description 用户列表
// @Param data query UserListRequest true "参数"
// @Param token header string true "token"
// @Router /api/users [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.UserModel]}
func (UserApi) UserListView(c *gin.Context) {
	var cr UserListRequest
	c.ShouldBindQuery(&cr)
	_list, count, _ := list.QueryList(models.UserModel{RoleID: cr.RoleID}, list.Option{
		Pagination: cr.Pagination,
		Likes:      []string{"nickName", "userName"},
		Preload:    []string{"RoleModel"},
	})
	res.OKWithList(_list, count, c)
}
