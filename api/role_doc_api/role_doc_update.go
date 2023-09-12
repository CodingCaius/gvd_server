package role_doc_api

import (
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/service/redis_service"
	"gvd_server/utils/set"

	"github.com/gin-gonic/gin"
)

type DocItem struct {
	DocID uint `json:"docID"`
	Sort  int  `json:"sort"`
}

type RoleDocUpdateRequest struct {
	RoleID  uint      `json:"roleID" binding:"required"`
	DocList []DocItem `json:"docList"`
}

// RoleDocUpdateView 角色文档 更新
// @Tags 角色文档管理
// @Summary 角色文档 更新
// @Description 角色文档 更新
// @Param token header string true "token"
// @Param data body RoleDocUpdateRequest true "参数"
// @Router /api/role_docs [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (RoleDocApi) RoleDocUpdateView(c *gin.Context) {
	var cr RoleDocUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		//res.FailWithError(err, c)
		return
	}

	var role models.RoleModel
	err = global.DB.Take(&role, cr.RoleID).Error
	if err != nil {
		res.FailWithMsg("角色不存在", c)
		return
	}

	var set1 = make([]uint, 0) // 前端传递来的id列表
	var set2 = make([]uint, 0) // 后端本来就有的id列表
	var docMap = make(map[uint]DocItem)
	for _, s := range cr.DocList {
		set1 = append(set1, s.DocID)
		docMap[s.DocID] = s
	}

	global.DB.Model(models.RoleDocModel{}).
		Where("role_id = ?", cr.RoleID).Select("doc_id").Scan(&set2)

	delSet, addSet := set.SetSub2(set1, set2)
	global.Log.Infof("del: %v, add: %v", delSet, addSet)
	if len(delSet) > 0 {
		global.DB.Where("role_id = ? and doc_id in ?", cr.RoleID, delSet).Delete(&models.RoleDocModel{})
	}
	if len(addSet) > 0 {
		var roleDocList []models.RoleDocModel
		for _, t := range addSet {
			sort := docMap[t].Sort
			roleDocList = append(roleDocList, models.RoleDocModel{
				RoleID: cr.RoleID,
				DocID:  t,
				Sort:   sort,
			})
		}
		err = global.DB.Create(&roleDocList).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMsg("角色-文档更新失败", c)
			return
		}
	}

	// 更新一些参数
	// 先查一下全部
	var roleDocList []models.RoleDocModel
	global.DB.Find(&roleDocList, "role_id = ? and doc_id in ?", cr.RoleID, set1)
	for _, model := range roleDocList {
		// 挨个判断 哪些值有变化
		doc := docMap[model.DocID]
		if model.Sort == doc.Sort {
			continue
		}
		// 不同的
		global.DB.Model(&model).Updates(map[string]any{
			"sort": doc.Sort,
		})
	}

	redis_service.ClearDocDocTree()
	redis_service.ClearDocContent()

	res.OKWithMsg("角色文档更新成功", c)

}
