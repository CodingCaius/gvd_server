package user_api

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/sirupsen/logrus"
  "gorm.io/gorm"
  "gvd_server/global"
  "gvd_server/models"
  "gvd_server/service/common/res"
)


// UserRemoveView 删除用户
// @Tags 用户管理
// @Summary 删除用户
// @Description 删除用户
// @Param token header string true "token"
// @Router /api/users [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserRemoveView(c *gin.Context) {
  var cr models.IDListRequest
  err := c.ShouldBindJSON(&cr)
  if err != nil {
    //res.FailWithError(err, c)
    return
  }

  var userList []models.UserModel
  global.DB.Find(&userList, cr.IDList)

  if len(cr.IDList) != len(userList) {
    res.FailWithMsg("数据一致性校验不通过", c)
    return
  }
  for _, model := range userList {
    err = UserRemoveService(model)
    if err != nil {
      logrus.Errorf("删除用户 %s 失败 err:%s", model.UserName, err.Error())
    } else {
      logrus.Infof("删除用户 %s 成功", model.UserName)
    }
  }
  res.OKWithMsg(fmt.Sprintf("批量删除成功，共删除%d个用户", len(cr.IDList)), c)

}

func UserRemoveService(user models.UserModel) (err error) {
  err = global.DB.Transaction(func(tx *gorm.DB) error {
    // imageModel 连带删除
    var imageList []models.ImageModel
    tx.Find(&imageList, "userID = ?", user.ID)
    if len(imageList) > 0 {
      err = tx.Delete(&imageList).Error
      if err != nil {
        return err
      }
    }
    // loginModel 不用连带删除

    // UserCollDocModel 连带删除
    var userCollList []models.UserCollDocModel
    tx.Find(&userCollList, "user_id = ?", user.ID)
    if len(userCollList) > 0 {
      err = tx.Delete(&userCollList).Error
      if err != nil {
        return err
      }
    }
    // UserPwdDocModel 连带删除
    var userPwdList []models.UserPwdDocModel
    tx.Find(&userPwdList, "user_id = ?", user.ID)
    if len(userPwdList) > 0 {
      err = tx.Delete(&userPwdList).Error
      if err != nil {
        return err
      }

    }
    err = tx.Delete(&user).Error
    return err
  })
  return err
}
