package flags

import (
	"context"
	"encoding/json"
	"fmt"
	"gvd_server/global"
	"gvd_server/models"
	"os"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

// 索引的数据
type ESRawMessage struct {
	Row json.RawMessage `json:"row"`
	ID  string          `json:"id"`
}

// ESIndexResponse 索引响应
type ESIndexResponse struct {
	Data    []ESRawMessage `json:"data"`
	Mapping string         `json:"mapping"`
	Index   string         `json:"index"`
}

// ESDump 执行 es 查询，获取指定索引的数据，然后将数据以 JSON 格式保存到文件中
func ESDump() {
	// 获取索引
	index := models.FullTextModel{}.Index()
	// 获取映射
	mapping := models.FullTextModel{}.Mapping()

	// 搜索索引
	res, err := global.ESClient.Search(index).
		// 搜索全文
		Query(elastic.NewMatchAllQuery()).
		// 设置每页显示数量
		Size(10000).Do(context.Background())

	if err != nil {
		logrus.Fatalf("%s err: %s", index, err.Error())
	}

	// 将搜索结果转换为ESRawMessage类型
	var dataList []ESRawMessage
	for _, hit := range res.Hits.Hits {
		dataList = append(dataList, ESRawMessage{
			Row: hit.Source,
			ID:  hit.Id,
		})
	}
	// 将ESRawMessage类型转换为ESIndexResponse类型
	response := ESIndexResponse{
		Mapping: mapping,
		Index:   index,
		Data:    dataList,
	}

	// 生成导出文件的文件名
	fileName := fmt.Sprintf("%s_%s.json", index, time.Now().Format("20060102"))
	// 创建文件并打开它以供写入数据
	file, _ := os.Create(fileName)

	// 将ESIndexResponse类型转换为json格式
	byteData, _ := json.Marshal(response)
	// 将json格式写入文件
	file.Write(byteData)
	file.Close()

	// 打印索引和文件名
	logrus.Infof("索引 %s 导出成功  %s", index, fileName)

}
