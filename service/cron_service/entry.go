// 定时任务
// 同步文档的点赞量和浏览量

package cron_service

import (
	"time"

	"github.com/robfig/cron/v3"
)

func CornInit() {
	timezone, _ := time.LoadLocation("Asia/Shanghai")

	Cron := cron.New(cron.WithSeconds(), cron.WithLocation(timezone))

	// 每天的2点去同步 redis 中的数据
	Cron.AddFunc("0 0 2 * * ?", SyncDocData)

	// 每天的3点去同步文档的浏览数据
	Cron.AddFunc("0 0 3 * * ?", SyncDocDataDate)

	// 会在后台运行，不会阻塞主程序的执行。
	Cron.Start()
}
