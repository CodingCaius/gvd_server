package models

//自定义连接表

// 角色文档关联表
type RoleDocModel struct {
	Model
	RoleID      uint      `gorm:"colmun:role_id;comment:角色id" json:"roleID"`
	RoleModel   RoleModel `gorm:"foreignKey:RoleID" json:"-"`
	DocID       uint      `gorm:"colmun:doc_id;comment:文档id" json:"docID"`
	DocModel    DocModel  `gorm:"foreignKey:DocID" json:"-"`
	// 角色文档的 密码 优先级更大
	// 如果文档开启了密码，角色文档本身没有，那就用角色的密码
	Pwd         *string   `gorm:"colmun:pwd;comment:密码配置" json:"pwd"`                 //null "" "有值" 优先级： 角色文档密码 > 角色密码
	FreeContent *string   `gorm:"colmun:freeContent;comment:试看部分" json:"freeContent"` //试看部分 优先级：角色文档试看> 文档试看字段> 文档按照特殊字符分隔的试看
	Sort        int       `gorm:"colmun:sort;comment:排序" json:"sort"`                 //排序
}
