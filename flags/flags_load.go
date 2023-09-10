package flags

import (
	"github.com/sirupsen/logrus"
	"gvd_server/global"
	"os"
	"strings"
)

func Load(sqlPath string) {
	byteData, err := os.ReadFile(sqlPath)
	if err != nil {
		logrus.Fatalf("%s err: %s", sqlPath, err.Error())
	}

	// 一定要按照\r\n分割
	sqlList := strings.Split(string(byteData), ";\r\n")
	for _, sql := range sqlList {
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		err = global.DB.Exec(sql).Error
		if err != nil {
			logrus.Errorf("%s err:%s", sql, err.Error())
			continue
		}
	}

	logrus.Infof("%s sql导入成功", sqlPath)

}