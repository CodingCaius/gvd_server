package doc_api

import (
	"context"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/utils/jwts"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

// DocSearchView 全文搜索
// @Tags 文档管理
// @Summary 全文搜索
// @Description 全文搜索
// @Param data query models.Pagination false "参数"
// @Router /api/docs/search [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.FullTextModel]}
func (DocApi) DocSearchView(c *gin.Context) {
	var cr models.Pagination
	_ = c.ShouldBindQuery(&cr)

	token := c.Request.Header.Get("token")
	claims, err := jwts.ParseToken(token)

	var roleID uint = 2 // 默认给一个访客角色

	if err == nil {
		roleID = claims.RoleID
	}

	// 创建一个空的 docIDList 数组，用于存储文档ID
	var docIDList []uint
	// 将用户的角色对应的文档ID存储在 docIDList 中
	global.DB.Model(models.RoleDocModel{}).Where("role_id = ?", roleID).Select("doc_id").Scan(&docIDList)

	if global.ESClient == nil {
		res.FailWithMsg("请配置es连接", c)
		return
	}

	if cr.Limit == 0 || cr.Limit >= 200 {
		cr.Limit = 10
	}
	if cr.Page == 0 {
		cr.Page = 1
	}
	from := (cr.Page - 1) * cr.Limit

	// 创建一个Elasticsearch的Bool查询 query
	var query = elastic.NewBoolQuery()
	// 如果查询关键字 cr.Key 不为空，将创建一个多字段匹配查询，并将其添加到 query 中
	if cr.Key != "" {
		query.Must(elastic.NewMultiMatchQuery(cr.Key, "title", "body"))
	}

	// 建一个空的 ids 数组，用于存储用户具有权限查看的文档ID
	var ids []interface{}
	for _, u := range docIDList {
		ids = append(ids, u)
	}

	// 用户只能搜自己权限才能看的文档
	// 使用 elastic.NewTermsQuery 创建一个Terms查询，限制搜索结果只包括用户具有权限查看的文档，将文档ID列表添加到 query 中
	query.Must(elastic.NewTermsQuery("docID", ids...))

	// 使用Elasticsearch客户端执行查询，并获取搜索结果 result
	result, err := global.ESClient.Search(models.FullTextModel{}.Index()).
		Query(query).
		Highlight(elastic.NewHighlight().Field("body").Field("title")).
		From(from).Size(cr.Limit).Do(context.Background())

	if err != nil {
		logrus.Error(err)
		res.FailWithMsg("查询失败", c)
		return
	}

	// 提取高亮字段（"body"和"title"）和反序列化文档模型

	// 获取搜索结果中总的文档数量
	count := result.Hits.TotalHits.Value
	var list = make([]models.FullTextModel, 0)
	for _, hit := range result.Hits.Hits {
		var model models.FullTextModel
		_ = json.Unmarshal(hit.Source, &model)

		bodyList, ok := hit.Highlight["body"]
		if ok {
			model.Body = bodyList[0]
		}
		titleList, ok := hit.Highlight["title"]
		if ok {
			model.Title = titleList[0]
		}

		list = append(list, model)
	}

	res.OKWithList(list, int(count), c)

}
