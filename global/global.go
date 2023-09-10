package global

import (
	"gvd_server/config"

	"github.com/cc14514/go-geoip2"
	"github.com/go-redis/redis"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config   *config.Config
	Log      *logrus.Logger
	DB       *gorm.DB
	Redis    *redis.Client
	ESClient *elastic.Client
	AddrDB   *geoip2.DBReader // 解析地址信息
)

// DocSplitSign 文本特殊分隔符
const DocSplitSign = "---===---"
