package list

import (
	"fmt"
	"gvd_server/global"
	"gvd_server/models"

	"gorm.io/gorm"
)

type Option struct {
	models.Pagination
	Likes   []string //模糊搜索的字段名列表
	Debug   bool     //是否开启调试模式
	Where   *gorm.DB //自定义的 WHERE 条件
	Preload []string //需要预加载的关联模型
}

// 执行查询操作，接受两个参数，一个是模型对象 model，另一个是查询选项 option，返回满足条件的数据列表
func QueryList[T any](model T, option Option) (list []T, count int, err error) {
	//这里的 Where 方法并没有实际执行查询操作，它只是构建了查询的条件部分。实际的查询操作通常在后续的代码中完成，例如调用 .Find() 或 .First() 等方法。
	query := global.DB.Where(model) //可以基于这个查询对象继续链式调用其他查询方法
	if option.Debug {
		query = query.Debug()
	}

	//默认是按照创建 时间 降序
	if option.Sort == "" {
		option.Sort = "createdAt desc"
	}

	if option.Limit == 0 {
		option.Limit = 10
	}

	if option.Where != nil {
		query.Where(option.Where)
	}

	//    ? 链式方法

	if option.Key != "" {
		//创建一个基础查询对象
		likeQuery := global.DB.Where("")
		//根据传入的查询选项 option，构建一个包含多个模糊搜索条件的查询，并将这些条件添加到原始的查询对象 query 中
		for index, column := range option.Likes {
			if index == 0 {
				likeQuery.Where(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
			} else {
				likeQuery.Or(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
			}
		}
		//需要赋值操作来将新构建的查询对象添加到原始的查询对象中
		query = query.Where(likeQuery)
	}

	//获取查询所影响的行数
	count = int(query.Find(&list).RowsAffected)

	for _, preload := range option.Preload {
		query = query.Preload(preload)
	}

	offset := (option.Page - 1) * option.Limit

	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return

}
