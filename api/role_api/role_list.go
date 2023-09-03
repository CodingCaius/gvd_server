package role_api

import (
  "github.com/gin-gonic/gin"
  "gvd_server/models"
  "gvd_server/service/common/list"
  "gvd_server/service/common/res"
)

type RoleListResponse struct {
  models.RoleModel
  DocCount int `json:"docCount"` // 角色拥有的文档数
  UserCount int `json:"userCount"` // 角色下的用户数
}

// RoleListView 角色列表
// @Tags 角色管理
// @Summary 角色列表
// @Description 角色列表
// @Param data query models.Pagination true "参数"
// @Param token header string true "token"
// @Router /api/roles [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[RoleListResponse]}
func (RoleApi) RoleListView(c *gin.Context) {
  var cr models.Pagination
  c.ShouldBindQuery(&cr)
  _list, count, _ := list.QueryList(models.RoleModel{}, list.Option{
    Pagination: cr,
    Likes:      []string{"title"},
    Preload:    []string{"DocsList", "UserList"},
  })
  roleList := make([]RoleListResponse, 0)
  for _, model := range _list {
    roleList = append(roleList, RoleListResponse{
      RoleModel: model,
      DocCount:  len(model.DocsList),
      UserCount:  len(model.UserList),
    })
  }
  res.OKWithList(roleList, count, c)
}
