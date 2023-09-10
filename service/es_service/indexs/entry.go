package indexs

import (
	"context"
	"gvd_server/global"
	"gvd_server/models"

	"github.com/sirupsen/logrus"
)

// CreateIndex 创建索引
func CreateIndex(esIndexInterFace models.ESIndexInterFace) {
	CreateIndexByJson(esIndexInterFace.Index(), esIndexInterFace.Mapping())
}

func CreateIndexByJson(index string, mapping string) {
	if global.ESClient == nil {
		logrus.Fatalf("请配置es连接")
	}
	if ExistsIndex(index) {
		DeleteIndex(index)
	}
	createIndex, err := global.ESClient.
		CreateIndex(index).
		BodyString(mapping).Do(context.Background())
	if err != nil {
		logrus.Fatalf("%s err:%s", index, err.Error())
	}
	logrus.Infof("索引 %s 创建成功", createIndex.Index)
}

// ExistsIndex 判断索引是否存在
func ExistsIndex(index string) bool {
	if global.ESClient == nil {
		logrus.Fatalf("请配置es连接")
	}
	exists, _ := global.ESClient.IndexExists(index).Do(context.Background())
	return exists
}
func DeleteIndex(index string) {
	if global.ESClient == nil {
		logrus.Fatalf("请配置es连接")
	}
	_, err := global.ESClient.
		DeleteIndex(index).Do(context.Background())
	if err != nil {
		logrus.Fatalf("%s err:%s", index, err.Error())
	}
	logrus.Infof("索引 %s 删除成功", index)
}
