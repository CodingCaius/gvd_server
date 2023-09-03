package image_api

import (
  "github.com/gin-gonic/gin"
  "gvd_server/models"
  "gvd_server/service/common/list"
  "gvd_server/service/common/res"
)

type ImageListResponse struct {
  models.ImageModel
  WebPath  string `json:"webPath"`
  NickName string `json:"nickName"`
}

// ImageListView 图片列表
// @Tags 图片管理
// @Summary 图片列表
// @Description 图片列表
// @Param data query models.Pagination true "参数"
// @Param token header string true "token"
// @Router /api/images [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[ImageListResponse]}
func (ImageApi) ImageListView(c *gin.Context) {
  var cr models.Pagination
  c.ShouldBindQuery(&cr)
  _list, count, _ := list.QueryList(models.ImageModel{}, list.Option{
    Pagination: cr,
    Likes:      []string{"fileName"},
    Preload:    []string{"UserModel"},
  })
  var imageList = make([]ImageListResponse, 0)
  for _, model := range _list {
    imageList = append(imageList, ImageListResponse{
      ImageModel: model,
      WebPath:    model.WebPath(),
      NickName:   model.UserModel.NickName,
    })
  }
  res.OKWithList(imageList, count, c)
}