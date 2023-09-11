// 同步每天的文档浏览量点赞量收藏量

package cron_service

import (
	"github.com/sirupsen/logrus"
	"gvd_server/global"
	"gvd_server/models"
)

// SyncDocDataDate 同步文档时间数据
func SyncDocDataDate() {
	logrus.Infof("开始同步文档时间数据")
	var docList []models.DocModel
	global.DB.Preload("UserCollDocList").Find(&docList)
	var docDataList []models.DocDataModel
	for _, model := range docList {
		docDataList = append(docDataList, models.DocDataModel{
			DocID:     model.ID,
			DocTitle:  model.Title,
			LookCount: model.LookCount,
			DiggCount: model.DiggCount,
			CollCount: len(model.UserCollDocList),
		})
	}

	err := global.DB.Create(&docDataList).Error
	if err != nil {
		logrus.Errorln(err)
		return
	}
	logrus.Infof("同步文档时间数据成功 共%d条", len(docDataList))

}