package cron_service

import (
	"fmt"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/redis_service"

	"github.com/sirupsen/logrus"
)

// SyncDocData 同步文档数据
func SyncDocData() {
	diggAll := redis_service.NewDocDigg().GetAll()
	lookAll := redis_service.NewDocLook().GetAll()

	// 获取所有文档
	var docList []models.DocModel
	global.DB.Find(&docList)
	// 遍历文档列表
	for _, model := range docList {
		// 格式化ID
		sID := fmt.Sprintf("%d", model.ID)
		// 获取 redis 中文档的点赞数和查看数
		digg := diggAll[sID]
		look := lookAll[sID]
		// 如果点赞数和查看数都为0，则跳过
		if digg == 0 && look == 0 {
			logrus.Infof("%s 无变化", model.Title)
			continue
		}

		// 更新文档的点赞数和查看数
		newDigg := digg + model.DiggCount
		newLook := look + model.LookCount
		global.DB.Model(&model).Updates(models.DocModel{
			DiggCount: newDigg,
			LookCount: newLook,
		})
		logrus.Infof("%s 更新成功 digg + %d， look + %d", model.Title, digg, look)
	}

	// 清空缓存中的点赞数和查看数
	redis_service.NewDocDigg().Clear()
	redis_service.NewDocLook().Clear()

}
