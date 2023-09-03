package models

//该表用于免密判断

// 判断用户对该文档是否免密，如果该表中存在记录，则免密

//用户文档密码模型
type UserPwdDocModel struct {
	Model
	UserID uint `grom:"column:user_id" json:"userID"`
	DocID  uint `grom:"column:doc_id" json:"docID"`
}
