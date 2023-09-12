package role_api

import (
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"

	"github.com/gin-gonic/gin"
)

// RoleIDListView 角色id列表
// @Tags 角色管理
// @Summary 角色id列表
// @Description 角色id列表
// @Param token header string true "token"
// @Router /api/roles/id [get]
// @Produce json
// @Success 200 {object} res.Response{data=[]models.Options}
func (RoleApi) RoleIDListView(c *gin.Context) {

	var list = make([]models.Options, 0)
	global.DB.Model(models.RoleModel{}).Select("id as value", "title as label").Scan(&list)
	res.OKWithData(list, c)
}
