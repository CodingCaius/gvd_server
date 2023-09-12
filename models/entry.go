package models //表结构

// 该项目的四大核心表
// 角色表 角色文档表 用户表 文档表

import "time"

/*
JSON 标签的作用是指定在将结构体实例转换为 JSON 字符串（序列化）时，使用指定的字段名称来表示结构体字段。同样，在从 JSON 字符串中解析数据（反序列化）时，也会根据 JSON 标签来找到对应的字段并赋值
*/

type Model struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt" json:"updatedAt"`
}

type Pagination struct {
	Page  int    `json:"page" form:"page"`
	Limit int    `json:"limit" form:"limit"`
	Key   string `json:"key" form:"key"` //模糊匹配的关键字
	Sort  string `json:"sort" form:"sort"`
}

type IDListRequest struct {
	IDList []uint `json:"idlist" form:"idList" binding:"required" label:"id列表"`
}

type IDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

type Options struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}
