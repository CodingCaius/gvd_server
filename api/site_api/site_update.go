package site_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/config"
	"gvd_server/core"
	"gvd_server/global"
	"gvd_server/service/common/res"
	"reflect"
)

// SiteUpdateView 站点配置更新
// @Tags 站点配
// @Summary 站点配置更新
// @Description 站点配置更新
// @Router /api/site [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (SiteApi) SiteUpdateView(c *gin.Context) {
	var cr config.Site
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}
	updateStructValue(cr, reflect.ValueOf(&global.Config.Site))

	//更新配置文件
	core.SetYaml()

	res.OKWithMsg("更新成功", c)
}

//更新结构体的值
func updateStructValue(data any, oldValue reflect.Value) {
	v := reflect.ValueOf(data)
	var updateIndexSlice []int
	for i := 0; i < v.NumField(); i++ {
		if !v.Field(i).IsZero() {
			// 如果不为空，就加入到要更新的列表中
			updateIndexSlice = append(updateIndexSlice, i)
		}
	}
	t := reflect.TypeOf(data)
	for _, updateIndex := range updateIndexSlice {
		// 拿字段名
		name := t.Field(updateIndex).Name
		// 拿到字段的value
		field := v.Field(updateIndex)
		// 动态修改
		oldValue.Elem().FieldByName(name).Set(field)
	}

}