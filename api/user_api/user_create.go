package user_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/plugins/log_stash"
	"gvd_server/service/common/res"
	"gvd_server/utils/pwd"
	"time"

	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	UserName string `json:"userName" binding:"required" label:"用户名"` // 用户名
	Password string `json:"password" binding:"required"`             // 密码
	NickName string `json:"nickName"`                                // 昵称
	RoleID   uint   `json:"roleID" binding:"required"`               // 角色id
}

// UserCreateView 创建用户
// @Tags 用户管理
// @Summary 创建用户
// @Description 创建用户，只能管理员创建
// @Param data body UserCreateRequest true "参数"
// @Param token header string true "token"
// @Router /api/users [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		//res.FailWithValidError(err, &cr, c)
		return
	}
	log := log_stash.NewAction(c)

	byteData, _ := json.Marshal(cr)
	log.SetItem("创建参数", string(byteData))

	err = createUser(models.UserModel{
		UserName:  cr.UserName,
		Password:  cr.Password,
		NickName:  cr.NickName,
		IP:        c.RemoteIP(),
		RoleID:    cr.RoleID,
		LastLogin: time.Now(),
	}, &log)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	log.SetItem("我的地盘", "我做主")
	log.Warn("用户创建成功")
	res.OKWithMsg("用户创建成功", c)
}


//将log作为参数传入，可以做到操作日志的更新
//根据是否是同一个log，来决定是否更新操作日志
func createUser(user models.UserModel, log *log_stash.Action) (err error) {
	err = global.DB.Take(&user, "userName = ?", user.UserName).Error
	if err == nil {
		log.SetItem("userName", user.UserName)
		log.Warn("创建用户错误，用户名已存在")
		return errors.New("用户名已存在")
	}

	log.SetItem("是否创建昵称", user.NickName == "")
	if user.NickName == "" {
		// 昵称如果不存在，那么就要
		var maxID uint
		global.DB.Model(models.UserModel{}).Select("max(id)").Scan(&maxID)
		user.NickName = fmt.Sprintf("用户_%d", maxID+1)
		log.SetItem("自动生成昵称", user.NickName)

	}
	var role models.RoleModel
	err = global.DB.Take(&role, user.RoleID).Error
	if err != nil {
		log.SetItem("角色id", user.RoleID)
		log.Warn("创建用户错误，角色不存在")
		return errors.New("角色不存在")
	}

	//密码处理
	user.Password = pwd.HashPwd(user.Password)

	err = global.DB.Create(&user).Error
	if err != nil {
		global.Log.Error(err)
		log.SetItem("错误原因", err.Error())
		log.Error("用户创建失败")
		return errors.New("用户创建失败")
	}
	log.Info("用户创建成功123")
	return nil
}
