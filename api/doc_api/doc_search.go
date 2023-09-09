package doc_api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/utils/jwts"
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

	var docIDList []uint
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
	var query = elastic.NewBoolQuery()

	if cr.Key != "" {
		query.Must(elastic.NewMultiMatchQuery(cr.Key, "title", "body"))
	}

	var ids []interface{}
	for _, u := range docIDList {
		ids = append(ids, u)
	}

	// 用户只能搜自己权限才能看的文档
	query.Must(elastic.NewTermsQuery("docID", ids...))

	result, err := global.ESClient.Search(models.FullTextModel{}.Index()).
		Query(query).
		Highlight(elastic.NewHighlight().Field("body").Field("title")).
		From(from).Size(cr.Limit).Do(context.Background())

	if err != nil {
		logrus.Error(err)
		res.FailWithMsg("查询失败", c)
		return
	}

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