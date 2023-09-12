package redis_service

import (
	"encoding/json"
	"fmt"
	"gvd_server/global"
	"time"
)

type DocContentResponse struct {
	Content   string    `json:"content"`
	IsSee     bool      `json:"isSee"`     // 是否试看
	IsPwd     bool      `json:"isPwd"`     // 是否需要密码
	IsColl    bool      `json:"isColl"`    // 用户是否收藏
	LookCount int       `json:"lookCount"` // 浏览量
	DiggCount int       `json:"diggCount"` // 点赞量
	CollCount int       `json:"collCount"` // 收藏量
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
}

const docContentKey = "role_doc_content_%d_%d"

// GetDocContent 获取一个角色 + 用户 的文档树
func GetDocContent(roleID, docID, userID uint) (doc DocContentResponse, ok bool) {
	byteData, err := global.Redis.HGet(fmt.Sprintf(docContentKey, roleID, docID), fmt.Sprintf("%d", userID)).Bytes()
	if err != nil {
		return
	}
	json.Unmarshal(byteData, &doc)
	ok = true
	return
}

// SetDocContent 设置一个角色的文档树
func SetDocContent(roleID, docID, userID uint, res DocContentResponse) {
	byteData, _ := json.Marshal(res)
	global.Redis.HSet(fmt.Sprintf(docContentKey, roleID, docID), fmt.Sprintf("%d", userID), string(byteData)).Result()
}

func ClearDocContent() {
	keysToDelete := []string{}
	// 获取所有匹配的键
	iter := global.Redis.Scan(0, "role_doc_content_*", 0).Iterator()
	for iter.Next() {
		keysToDelete = append(keysToDelete, iter.Val())
	}

	// 删除匹配的键
	global.Redis.Del(keysToDelete...)
}