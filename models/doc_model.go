package models

import (
	"gvd_server/global"
	"sort"
	"strings"
)

type DocModel struct {
	Model
	Title     string `gorm:"comment:文档标题;size:32" json:"title"`
	Content   string `gorm:"comment:文档内容" json:"-"`
	DiggCount int    `gorm:"comment:点赞量" json:"diggcount"`
	LookCount int    `gorm:"comment:浏览量" json:"lookCount"`
	Key       string `gorm:"comment:key;not null;unique" json:"key"`
	ParentID  *uint  `gorm:"comment:父文档id column:parent_id" json:"parentID"`
	//通过 ParentID 字段与父文档DocModel 关联起来的
	ParentModel *DocModel   `gorm:"foreignKey:ParentID" json:"-"` //父文档
	Child       []*DocModel `gorm:"foreignKey:ParentID" json:"child"` //子孙文档
	FreeContent string      `gorm:"comment:预览部分;column:freeContent" json:"freeContent"`

	//收藏该文档的用户列表
	/*连接表 user_coll_doc_models：
	DocID: 外键，关联到文档表的 DocID 字段。
	UserID: 外键，关联到用户表的 UserID 字段。*/
	UserCollDocList []UserModel `gorm:"many2many:user_coll_doc_models;joinForeignKey:DocID;JoinReferences:UserID" json:"-"`
}

// FindAllParentDocList 找一个文档的所有父文档
func FindAllParentDocList(doc DocModel, docList *[]DocModel) {
	// 不管谁来，先把自己放进去
	*docList = append(*docList, doc)
	if doc.ParentID != nil {
		// 说明有父文档
		var parentDoc DocModel
		global.DB.Take(&parentDoc, *doc.ParentID)
		FindAllParentDocList(parentDoc, docList)
	}
}

// FindAllSubDocList 找一个文档的所有子文档
func FindAllSubDocList(doc DocModel) (docList []DocModel) {
	global.DB.Preload("Child").Take(&doc)
	for _, model := range doc.Child {
		docList = append(docList, *model)
		docList = append(docList, FindAllSubDocList(*model)...)
	}
	return
}

// DocTree 返回文档树 根据父亲id来查找子id，刚开始传入 nil
func DocTree(parentID *uint) (docList []*DocModel) {
	var query = global.DB.Where("") // 创建一个查询对象

	if parentID == nil {
		// 如果 parentID 为 nil，表示要找根文档，因此设置查询条件为 parent_id is null
		query.Where("parent_id is null")
	} else {
		// 如果 parentID 不为 nil，表示要找具有指定 parentID 的子文档
		query.Where("parent_id = ?", *parentID)
	}

	// 在数据库中执行查询，并将结果加载到 docList 切片中
	global.DB.Preload("Child").Where(query).Find(&docList)

	// 遍历查询到的每个文档模型
	for _, model := range docList {
		// 递归调用 DocTree 函数，以获取当前文档的子文档，并将子文档列表赋值给当前文档的 Child 字段
		subDocs := DocTree(&model.ID)
		model.Child = subDocs
	}

	return // 返回构建好的文档树
}

// SortDocByPotCount 按照点的个数进行排序  返回最小的那个元素点的个数
func SortDocByPotCount(docList []*DocModel) (minCount int) {
	if len(docList) == 0 {
		return
	}
	sort.Slice(docList, func(i, j int) bool {
		count1 := GetByPotCount(docList[i])
		count2 := GetByPotCount(docList[j])
		if count1 == count2 {
			// 点的个数相同，按照id，小的放在前面，升序排列
			return docList[i].ID < docList[j].ID
		}
		return count1 < count2
	})
	return GetByPotCount(docList[0])
}

// GetByPotCount 获取文档点的个数
func GetByPotCount(doc *DocModel) int {
	return strings.Count(doc.Key, ".")
}

// TreeByOneDimensional 树的一维化， 将文档树形式转换为一维的
func TreeByOneDimensional(docList []*DocModel) (list []*DocModel) {
	for _, model := range docList {
		list = append(list, model)
		//递归调用，将子文档也加入
		list = append(list, TreeByOneDimensional(model.Child)...)
	}
	return
}
