package models

// 用户登录数据
type LoginModel struct {
	Model
	UserID    uint      `gorm:"column:userID" json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	IP        string    `gorm:"size:20" json:"ip"`
	NickName  string    `gorm:"size:42;column:userID" json:"nikeName"`
	UA        string    `gorm:"size:256" json:"ua"`
	Token     string    `gorm:"size:256" json:"token"`
	Device    string    `gorm:"size:256" json:"device"`
	Addr      string    `gorm:"size:64" json:"addr"`
}
