package user_api

import (
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/plugins/log_stash"
	"gvd_server/service/common/res"
	"gvd_server/utils/jwts"
	"gvd_server/utils/pwd"
	"time"

	"github.com/gin-gonic/gin"
)

type UserLoginRequest struct {
	UserName string `json:"userName" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required" label:"密码"`
}

// UserLoginView 用户登录
// @Tags 用户管理
// @Summary 用户登录
// @Description 用户登录
// @Param data body UserLoginRequest true "参数"
// @Router /api/login [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserLoginView(c *gin.Context) {
	var ulr UserLoginRequest

	log := log_stash.NewAction(c)
	//设置请求日志
	log.SetRequest(c)

	err := c.ShouldBindJSON(&ulr)
	if err != nil {
		return
	}
	// if err != nil {
	// 	res.FailWithVaildError(err, &ulr, c)
	// 	return
	// }

	var user models.UserModel
	err = global.DB.Take(&user, "userName = ?", ulr.UserName).Error
	if err != nil {
		global.Log.Warn("用户名不存在", ulr.UserName)
		log_stash.NewFailLogin("用户名不存在", ulr.UserName, ulr.Password, c)
		res.FailWithMsg("用户名或密码错误", c)
		return
	}
	if !pwd.CheckPwd(user.Password, ulr.Password) {
		global.Log.Warn("用户密码错误", ulr.UserName, ulr.Password)
		log_stash.NewFailLogin("用户密码错误", ulr.UserName, ulr.Password, c)
		res.FailWithMsg("用户名或密码错误", c)
		return
	}

	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserName: user.UserName,
		NickName: user.NickName,
		RoleID:   user.RoleID,
		UserID:   user.ID,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("生成Token失败", c)
		return
	}

	c.Request.Header.Set("token", token)
	log_stash.NewSuccessLogin(c)


	log.SetItem("用户登录", "成功")
	log.SetItem("token", token)
	log.SetItem("数组", []string{"1", "oahnf"})
	log.SetItem("对象", user)
	log.SetItem("对象", map[string]any{"name": user.UserName, "age": 18})

	global.DB.Model(&user).Update("lastLogin", time.Now())

	log.Info("用户登录成功")

	res.OKWithData(token, c)

	//设置响应日志
	log.SetResponse(c)
}
