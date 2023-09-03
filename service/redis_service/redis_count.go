//利用redis缓存文档的浏览量

package redis_service

import (
	"fmt"
	"gvd_server/global"
	"strconv"
)

// 定义了两个常量 docDiggIndex 和 docLookIndex，
// 分别表示文档点赞数和浏览数在 Redis 中的存储索引
const (
	docDiggIndex = "docDiggIndex"
	docLookIndex = "docLookIndex"
)

// 表示  存储索引
type CountDB struct {
	Index string // 哈希类型中的 key
}


// 创建不同的 CountDB 对象，每个对象与特定类型的计数（点赞或浏览）相关联

func NewDocDigg() CountDB {
	return CountDB{
		Index: docDiggIndex,
	}
}

func NewDocLook() CountDB {
	return CountDB{
		Index: docLookIndex,
	}
}



// 通过 ID 设置浏览数/点赞数
func (c CountDB) SetById(id uint) error {
	return c.Set(fmt.Sprintf("%d", id))
}

// Set 给某一个 field 设置数字 调用一次+1
func (c CountDB) Set(field string) error {
	return c.SetCount(field, 1)
}

// SetCount 给某一个 field 设置数字,调用一次 + num
func (c CountDB) SetCount(field string, num int) error {
	oldNum, _ := global.Redis.HGet(c.Index, field).Int()
	newNum := oldNum + num
	err := global.Redis.HSet(c.Index, field, newNum).Err()
	return err
}

// Get 返回某一个 field 对应的值
func (c CountDB) Get(field string) int {
	num, _ := global.Redis.HGet(c.Index, field).Int()
	return num
}

// GetById 返回某一个 field 对应的值
func (c CountDB) GetById(id uint) int {
	return c.Get(fmt.Sprintf("%d", id))
}

// GetAll 返回这个索引下的所有 字段和对应的整数值
func (c CountDB) GetAll() map[string]int {
	var countMap = map[string]int{}
	maps := global.Redis.HGetAll(c.Index).Val()
	for field, val := range maps {
		num, _ := strconv.Atoi(val)
		countMap[field] = num
	}
	return countMap
}

// Clear 清空这个索引里面的值
func (c CountDB) Clear() {
	global.Redis.Del(c.Index)
}
