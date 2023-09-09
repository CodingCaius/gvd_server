
// 将文档数据添加到 ES 搜索引擎中
// 以便进行全文搜索操作
// 在创建文档时调用该方法


package full_search_service

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvd_server/global"
	"gvd_server/models"
)

// FullSearchCreate 添加
func FullSearchCreate(doc models.DocModel) {
	// 如果ES客户端为空，则返回
	if global.ESClient == nil {
		return
	}
	// 解析markdown文档
	searchDataList := MarkdownParse(doc.ID, doc.Title, doc.Content)
	// 初始化bulk请求
	bulk := global.ESClient.Bulk().Index(models.FullTextModel{}.Index()).Refresh("true")
	// 遍历文档解析后的数据，批量添加到es中
	for _, model := range searchDataList {
		// 创建一个Elasticsearch的批量创建请求对象
		req := elastic.NewBulkCreateRequest().Doc(models.FullTextModel{
			DocID: doc.ID,
			Title: model.Title,
			Body:  model.Body,
			Slug:  model.Slug,
		})
		// 添加到bulk请求中
		bulk.Add(req)
	}
	// 执行bulk请求
	res, err := bulk.Do(context.Background())
	if err!= nil {
		logrus.Errorf("%#v 数据添加失败 err:%s", doc, err.Error())
		return
	}
	logrus.Infof("添加全文搜索记录 %d 条", len(res.Succeeded()))
}