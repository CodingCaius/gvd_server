package log_stash

import "time"

type LogModel struct {
	ID          uint      `json:"id" gorm:"primaryKey"`                  // 主键id
	CreatedAt   time.Time `gorm:"column:createdAt" json:"createdAt"`     // 添加时间
	UpdatedAt   time.Time `gorm:"column:updatedAt" json:"updatedAt"`     // 更新时间
	IP          string    `json:"ip"`                                    // ip
	Addr        string    `json:"addr"`                                  // 地址
	Level       Level     `json:"level"`                                 // 等级
	Title       string    `json:"title"`                                 // 标题
	Content     string    `json:"content"`                               // 详情
	UserID      uint      `gorm:"column:userID" json:"userID"`           // 用户id
	UserName    string    `gorm:"column:userName" json:"userName"`       // 用户名
	ServiceName string    `gorm:"column:serviceName" json:"serviceName"` // 服务名称
	Status      bool      `json:"status"`                                // 登录状态
	Type        LogType   `json:"type"`                                  // 日志的类型  1 登录 2 操作 3 运行
	// 阅读状态   true   已读  false  未读  默认是未读状态
	ReadStatus bool `gorm:"column:readStatus;default:false" json:"readStatus"`
}
