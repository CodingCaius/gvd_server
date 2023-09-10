// 导出mysql数据库

package flags

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"gvd_server/global"
	"os/exec"
	"time"
)

func Dump() {
	// 获取mysql配置信息
	mysql := global.Config.Mysql

	// 获取当前时间，并设置格式
	timer := time.Now().Format("20060102")

	// 设置sql文件的路径
	sqlPath := fmt.Sprintf("%s_%s.sql", mysql.DB, timer)

	// 调用系统命令， 执行mysqldump进行数据库导出
	cmder := fmt.Sprintf("mysqldump -u%s -p%s %s > %s", mysql.Username, mysql.Password, mysql.DB, sqlPath)
	// 创建一个用于执行 shell 命令的 exec.Command 对象
	cmd := exec.Command("sh", "-c", cmder)

	// 创建两个字节缓冲区，分别用于存储命令的标准输出和标准错误
	var out, stderr bytes.Buffer
	// 将命令的标准输出和标准错误重定向到之前创建的缓冲区中，以便捕获命令执行的结果和错误信息
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// 执行命令
	err := cmd.Run()

	// 如果出现错误，则打印错误信息
	if err!= nil {
		logrus.Errorln(err.Error(), stderr.String())
		return
	}
	// 打印sql文件的路径
	logrus.Infof("sql文件 %s 导出成功", sqlPath)
}