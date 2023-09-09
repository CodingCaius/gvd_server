package core

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvd_server/global"
)

func InitEs() (client *elastic.Client) {

	es := global.Config.ES
	addr := es.Addr
	if addr == "" {
		return
	}
	client, err := elastic.NewClient(
		elastic.SetURL(addr),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(es.User, es.Password),
	)
	if err != nil {
		logrus.Fatalf(fmt.Sprintf("[%s] es 连接失败, err:%s", addr, err.Error()))
	}
	return client
}