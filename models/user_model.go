package models //表结构

import "time" /*
JSON 标签的作用是指定在将结构体实例转换为 JSON 字符串（序列化）时，使用指定的字段名称来表示结构体字段。同样，在从 JSON 字符串中解析数据（反序列化）时，也会根据 JSON 标签来找到对应的字段并赋值
*/

type UserModel struct {
	Model
	UserName  string    `gorm:"column:userName;size:36;unique;not null;comment:用户名" json:"userName"` // 用户名
	Password  string    `gorm:"column:password;size:128;comment:密码"  json:"-"`                       // 密码
	Avatar    string    `gorm:"column:avatar;size:256;comment:头像"  json:"avatar"`                    // 头像
	NickName  string    `gorm:"column:nickName;size:36;comment:昵称"  json:"nickName"`                 // 昵称
	Email     string    `gorm:"column:email;size:128;comment:邮箱"  json:"email"`                      // 邮箱
	Token     string    `gorm:"column:token;size:64;comment:其他平台的唯一id"  json:"-"`                    // 其他平台的唯一id
	IP        string    `gorm:"column:ip;size:16;comment:ip地址"  json:"ip"`                           // ip
	Addr      string    `gorm:"column:addr;size:64;comment:地址"  json:"addr"`                         // 地址
	RoleID    uint      `gorm:"column:roleID;comment:用户对应的角色" json:"roleID"`                         // 用户对应的角色
	LastLogin time.Time `gorm:"column:lastLogin" json:"lastLogin"`                                   // 用户最后登录时间
	RoleModel RoleModel `gorm:"foreignKey:RoleID" json:"roleModel"`                                  // 用户角色信息
}

//指明两个表的外键关系
//GORM 将会在查询 UserModel 表时自动进行联接（JOIN），并根据 RoleID 字段的值来获取关联的角色信息
