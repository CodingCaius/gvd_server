package full_search_service

import (
	"gvd_server/global"
	"gvd_server/models"
)

// FullSearchUpdate 更新，先删除再添加
func FullSearchUpdate(doc models.DocModel) {
	if global.ESClient == nil {
		return
	}
	// 添加之前先删除之前的
	FullSearchDelete(doc.ID)
	FullSearchCreate(doc)
}