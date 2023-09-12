package redis_service

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"gvd_server/global"
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

const roleDocTreeKey = "role_doc_tree_%d"

// GetRoleDocTree 获取一个角色 + 用户 的文档树
func GetRoleDocTree(roleID uint, userID uint) (roleDoc []RoleDocTree, ok bool) {
	logrus.Infof("获取一个角色的文档树 角色id: %d 用户id：%d", roleID, userID)
	byteData, err := global.Redis.HGet(fmt.Sprintf(roleDocTreeKey, roleID), fmt.Sprintf("%d", userID)).Bytes()
	if err != nil {
		return
	}
	roleDoc = make([]RoleDocTree, 0)
	json.Unmarshal(byteData, &roleDoc)
	ok = true
	return
}

// SetRoleDocTree 设置一个角色的文档树
func SetRoleDocTree(roleID uint, userID uint, docTree []RoleDocTree) {
	logrus.Infof("设置一个角色的文档树 角色id: %d 用户id：%d", roleID, userID)
	byteData, _ := json.Marshal(docTree)
	global.Redis.HSet(fmt.Sprintf(roleDocTreeKey, roleID), fmt.Sprintf("%d", userID), string(byteData)).Result()
}

func ClearDocDocTree() {
	keysToDelete := []string{}
	// 获取所有匹配的键
	iter := global.Redis.Scan(0, "role_doc_tree_*", 0).Iterator()
	for iter.Next() {
		keysToDelete = append(keysToDelete, iter.Val())
	}

	// 删除匹配的键
	global.Redis.Del(keysToDelete...)
}