package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func Task1() {
	fmt.Println("task1", time.Now())
}

func Task2() {
	fmt.Println("task2", time.Now())
}

func Task3(name string) func() {
	return func() {
		fmt.Println(name, "task3", time.Now())
	}
}

type Job struct {
	Name string
}

func (j Job) Run() {
	fmt.Println(j.Name, "job", time.Now())
}

func main() {
	// 创建了一个新的 cron 调度器对象 cronTab。通过 cron.WithSeconds() 选项，指定了调度器应该精确到秒级别的精度。可以在 cron 表达式中指定秒数来触发任务。
	cronTab := cron.New(cron.WithSeconds())

	//cronTab.AddFunc("1-20,40-50 * * * *?", Task1)
	//_, err := cronTab.AddFunc("20-40,50-59 * * * *?", Task2)
	//fmt.Println(err)

	//cronTab.AddJob("1-59 * * * * *", Job{Name: "caius"})

	// 这样写可以传参数
	cronTab.AddFunc("1-59 * * * *?", Task3("caius"))

	// 启动 cron 调度器，使其开始运行定时任务
	cronTab.Start()

	// 使程序保持运行状态
	//  cronTab.Start() 启动后会在后台运行，程序如果直接退出，那么定时任务可能不会触发
	select {}

}
