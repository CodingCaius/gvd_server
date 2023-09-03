package image_api

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "gvd_server/global"
  "gvd_server/models"
  "gvd_server/service/common/res"
  "os"
)

// ImageRemoveView 删除图片
// @Tags 图片管理
// @Summary 删除图片
// @Description 删除图片
// @Param token header string true "token"
// @Router /api/images [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (ImageApi) ImageRemoveView(c *gin.Context) {
  var cr models.IDListRequest
  err := c.ShouldBindJSON(&cr)
  if err != nil {
    res.FailWithMsg("参数错误", c)
    return
  }

  //先进行一致性校验，传过来的数据是否都存在
  var imageList []models.ImageModel
  global.DB.Find(&imageList, cr.IDList)

  if len(cr.IDList) != len(imageList) {
    res.FailWithMsg("数据一致性校验不通过", c)
    return
  }

  for _, model := range imageList {
    imageRemove(model)
  }
  res.OKWithMsg(fmt.Sprintf("批量删除成功，共删除%d张图片", len(cr.IDList)), c)
}

// 删除图片的时候，发现有多个相同的hash，那就只删除记录
func imageRemove(image models.ImageModel) {
  var count int64
  global.DB.Model(models.ImageModel{}).
    Where("hash = ?", image.Hash).Count(&count)
  // count的值肯定是大于等于1的
  // 大于等于2 那就只删记录
  // 等于1 那就删记录，并且删文件
  if count == 1 {
	//删文件
    err := os.Remove(image.Path)
    if err != nil {
      global.Log.Errorf("删除文件 %s 错误 %s", image.Path, err.Error())
    } else {
      global.Log.Infof("删除文件 %s 成功", image.Path)
    }
  }
  //删记录
  global.DB.Delete(&image)
}
