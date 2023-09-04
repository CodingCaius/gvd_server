// 编写文档树步骤
// 1,先找这个角色拥有的文档id列表
// 2,按照key里面点的个数排序，并且把最小的点数返回
// 3,循环上一步排序好的文档列表，如果它的key的点数和最小点数一样,那么它就是根文档
// 4,把根文档的那个树一维化便于循环，
// 然后通过最大字符前缀匹配，将除根文档外的其他文档插入到对应的父文档中
// 5,最后将角色文档树转化为特定类型 RoleDocTreeResponse，便于响应


package role_doc_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/utils"
	"gvd_server/utils/jwts"
)

type RoleDocTree struct {
	ID       uint          `json:"key"`
	Title    string        `json:"title"`
	Children []RoleDocTree `json:"children"`
	IsPwd    bool          `json:"isPwd"`  // 是否需要密码
	Unlock   bool          `json:"unlock"` // 是否解锁
	IsColl   bool          `json:"isColl"` // 是否收藏
	IsSee    bool          `json:"isSee"`  // 是否试看
}

type RoleDocTreeResponse struct {
	List []RoleDocTree `json:"list"`
}

// RoleDocTreeView 角色文档树
// @Tags 角色文档管理
// @Summary 角色文档树
// @Description 角色文档树
// @Param token header string true "token"
// @Router /api/role_docs [get]
// @Produce json
// @Success 200 {object} res.Response{data=RoleDocTreeResponse}
func (RoleDocApi) RoleDocTreeView(c *gin.Context) {
	token := c.Request.Header.Get("token")
	claims, err := jwts.ParseToken(token)

	var roleID uint = 2 // 默认给一个访客角色

	if err == nil {
		roleID = claims.RoleID
	}

	var response = RoleDocTreeResponse{
		List: make([]RoleDocTree, 0),
	}

	var docIDList []uint
	var roleDocList []models.RoleDocModel
	global.DB.
		Preload("RoleModel").
		Preload("DocModel").
		Find(&roleDocList, "role_id = ?", roleID).Select("doc_id").Scan(&docIDList)
	if len(roleDocList) == 0 {
		res.OKWithData(response, c)
		return
	}

	// 查文档列表
	var docList []*models.DocModel
	global.DB.Find(&docList, docIDList)
	if len(docList) == 0 {
		res.OKWithData(response, c)
		return
	}
	// 按照key里面点的个数排序，并且把最小的点数返回
	minCount := models.SortDocByPotCount(docList)
	// 构造一个新的docList
	// docListPointer 是指针
	var docListPointer = new([]*models.DocModel)
	// 循环排序好之后的 docList
	for _, model := range docList {
		// 判断，它们的key的点数是不是和最小的那个点数一样，一样就放入根列表
		if models.GetByPotCount(model) == minCount {
			// 根文档
			*docListPointer = append(*docListPointer, model)
			continue
		}
		// 子文档，需要找他们的父文档
		insertDoc(docListPointer, model)
	}
	// 文档树转换
	var docPwdMap = map[uint]bool{}
	var docSeeMap = map[uint]bool{}
	var docCollMap = map[uint]bool{}
	var unLockMap = map[uint]bool{}

	for _, model := range roleDocList {
		// 判断有密码
		// 角色文档的 密码 优先级更大
	    // 如果文档开启了密码，角色文档本身没有，那就用角色的密码
		// model.Pwd不为nil，说明密码按键已打开，在角色文档表中，密码可能为空，也可能填写有密码
		if model.Pwd != nil && (*model.Pwd != "" || model.RoleModel.Pwd != "") {
			docPwdMap[model.DocID] = true
		}
		// 判断试看
		if model.FreeContent != nil {
			docSeeMap[model.DocID] = true
		}
	}

	// 判断这个人
	if claims != nil && claims.UserID != 0 {
		// 判断是否收藏了
		// 判断是否解锁了
		var docCollIDList []uint // 用户收藏了哪些文档
		global.DB.Model(models.UserCollDocModel{}).Where("user_id = ?", claims.UserID).
			Select("doc_id").Scan(&docCollIDList)
		var userPwdDocIDList []uint // 用户解锁了哪些文档
		global.DB.Model(models.UserPwdDocModel{}).Where("user_id = ?", claims.UserID).
			Select("doc_id").Scan(&userPwdDocIDList)
		for _, id := range docCollIDList {
			docCollMap[id] = true
		}
		for _, id := range userPwdDocIDList {
			unLockMap[id] = true
		}
	}

	// 角色文档树转换
	// 返回给前端的 list
	list := RoleDocTreeTransition(*docListPointer, docPwdMap, docSeeMap, docCollMap, unLockMap)
	response.List = list
	res.OKWithData(response, c)

}

// 将文档模型的扁平列表转换为一个有层次结构的树，其中每个文档模型可以包含子文档，以及其他与文档相关的属性
// RoleDocTreeTransition 角色文档树 转换为特定类型
func RoleDocTreeTransition(docList []*models.DocModel, docPwdMap, docSeeMap, docCollMap, unLockMap map[uint]bool) (list []RoleDocTree) {
	for _, model := range docList {
		// 递归调用 RoleDocTreeTransition 函数，以处理当前文档模型的子文档，并将结果存储在 children 中
		children := RoleDocTreeTransition(model.Child, docPwdMap, docSeeMap, docCollMap, unLockMap)
		if children == nil {
			children = make([]RoleDocTree, 0)
		}
		docTree := RoleDocTree{
			ID:       model.ID,
			Title:    model.Title,
			Children: children,
			IsPwd:    docPwdMap[model.ID],
			Unlock:   unLockMap[model.ID],
			IsColl:   docCollMap[model.ID],
			IsSee:    docSeeMap[model.ID],
		}
		list = append(list, docTree)
	}
	// 最终的 list切片 包含了文档树的层次结构
	return
}

// 找符合自己的父文档，并且插入进去
func insertDoc(docList *[]*models.DocModel, doc *models.DocModel) {
	// 把根文档的那个树一维化，通过最大字符前缀匹配，找到后面的key，最有可能匹配的key
	// 一维化
	oneDimensionalDocList := models.TreeByOneDimensional(*docList)
	// 通过最大前缀匹配找到这个model的key，对应应该放在哪一个对象上
	var keys []string
	for _, model := range oneDimensionalDocList {
		keys = append(keys, model.Key)
	}
	_, index := utils.FindMaxPrefix(doc.Key, keys)
	if index == -1 {
		// 没有满足的，那么就只能把它放到根文档上去了
		*docList = append(*docList, doc)
	} else {
		// 匹配到父文档，添加
		oneDimensionalDocList[index].Child = append(oneDimensionalDocList[index].Child, doc)
	}

}