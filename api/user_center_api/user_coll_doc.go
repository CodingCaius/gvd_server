package user_center_api

import (
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/utils/jwts"

	"github.com/gin-gonic/gin"
)

// UserCollDocView 收藏文档或取消收藏
// @Tags 个人中心
// @Summary 收藏文档或取消收藏
// @Description 收藏文档或取消收藏
// @Param token header string true "token"
// @Router /api/user_center/user_coll [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserCenterApi) UserCollDocView(c *gin.Context) {
	var cr models.IDRequest

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}
	token := c.Request.Header.Get("token")
	claims, _ := jwts.ParseToken(token)
	user, err := claims.GetUser()
	if err != nil {
		res.FailWithMsg("用户信息错误", c)
		return
	}

	var userColl models.UserCollDocModel
	err = global.DB.Take(&userColl, "doc_id = ? and user_id = ?", cr.ID, user.ID).Error
	if err != nil {
		// 我要收藏
		global.DB.Create(&models.UserCollDocModel{
			DocID:  cr.ID,
			UserID: user.ID,
		})
		res.OKWithMsg("收藏成功", c)
		return
	}

	// 取消收藏
	global.DB.Delete(&userColl)
	res.OKWithMsg("取消收藏成功", c)
}
