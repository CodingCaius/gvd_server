package role_api

import (
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"

	"github.com/gin-gonic/gin"
)

// RoleRemoveView 删除角色
// @Tags 角色管理
// @Summary 删除角色
// @Description 删除角色
// @Param data body models.IDRequest true "参数"
// @Param token header string true "token"
// @Router /api/roles [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (RoleApi) RoleRemoveView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		//res.FailWithError(err, c)
		return
	}

	var role models.RoleModel
	err = global.DB.Preload("DocsList").Preload("UserList").Take(&role, cr.ID).Error
	if err != nil {
		res.FailWithMsg("不存在的角色", c)
		return
	}

	if role.IsSystem {
		res.FailWithMsg("系统角色，不可删除", c)
		return
	}
	global.Log.Infof("删除角色 %s，关联用户 %d 个，删除关联文档 %d 个", role.Title, len(role.UserList), len(role.DocsList))
	if len(role.UserList) > 0 {
		// 统一修改用户的角色id为2
		global.DB.Model(&role.UserList).Update("roleID", "2")
	}
	if len(role.DocsList) > 0 {
		global.DB.Model(&role).Association("DocsList").Delete(role.DocsList)
	}

	global.DB.Delete(&role)
	res.OKWithMsg("删除角色成功", c)
}
