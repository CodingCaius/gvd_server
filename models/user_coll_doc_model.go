package models

//自定义连接表

//用户收藏文档模型
type UserCollDocModel struct {
	Model
	DocID     uint      `gorm:"column:doc_id" json:"docID"`
	DocModel  DocModel  `gorm:"foreignKey:DocID"`
	UserID    uint      `gorm:"column:user_id" json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
}
