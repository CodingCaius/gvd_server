package user_api

import (
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/utils/pwd"

	"github.com/gin-gonic/gin"
)

type UserUpdateRequest struct {
	ID uint `jsn:"id" binding:"required" label:"用户id"`
	//bingding字段 定义字段在数据绑定（比如 HTTP 请求参数绑定）时的行为
	Password string `json:"password" binding:"required"`
	NickName string `json:"nickName"`
	RoleID   uint   `json:"roleID" binding:"required"` //角色ID

}

// UserUpdateView 管理员更新用户信息
// @Tags 用户管理
// @Summary 管理员更新用户信息
// @Description 管理员更新用户信息
// @Param token header string true "token"
// @Param data body UserUpdateRequest true "参数"
// @Router /api/users [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserUpdateView(c *gin.Context) {
	var uur UserUpdateRequest
	err := c.ShouldBindJSON(&uur)
	if err != nil {
		//res.FailWithError(err, c)
		return
	}

	var user models.UserModel
	err = global.DB.Take(&user, uur.ID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}

	if uur.Password != "" {
		uur.Password = pwd.HashPwd(uur.Password)
	}
	if uur.RoleID != 0 {
		var role models.RoleModel
		err = global.DB.Take(&role, uur.RoleID).Error
		if err != nil {
			res.FailWithMsg("角色不存在", c)
			return
		}
	}

	err = global.DB.Model(&user).Updates(models.UserModel{
		Password: uur.Password,
		NickName: uur.NickName,
		RoleID: uur.RoleID,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("用户更新失败", c)
		return
	}

	res.OKWithMsg("用户更新成功", c)
}
