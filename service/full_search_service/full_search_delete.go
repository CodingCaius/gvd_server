package full_search_service

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvd_server/global"
	"gvd_server/models"
)

// FullSearchDelete 删除
func FullSearchDelete(docID uint) {
	if global.ESClient == nil {
		return
	}
	
	// 构建条件，根据docID删除
	var query = elastic.NewTermQuery("docID", docID)

	res, err := global.ESClient.DeleteByQuery(models.FullTextModel{}.Index()).
		Query(query).Refresh("true").Do(context.Background())
	if err != nil {
		logrus.Errorf("%d 数据删除失败 err:%s", docID, err.Error())
		return
	}
	logrus.Infof("删除全文搜索记录 %d 条", res.Deleted)
}