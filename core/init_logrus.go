//初始化日志
package core

import (
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

type LogRequest struct {
	LogPath string //日志的目录
	AppName string //app的名字
	NoDate  bool   //是否需要按照时间分割
	NoErr   bool   //是否单独存放err日志
	NoGlobal bool //是否替换全局logrus
}

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s\n", timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

// 按照时间分隔写入日志文件的hook
type DateHook struct {
	file     *os.File
	fileDate string
}

func (hook *DateHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *DateHook) Fire(entry *logrus.Entry) error {
	timer := entry.Time.Format("2006-01-02")
	line, _ := entry.String()
	if hook.fileDate == timer {
		hook.file.Write([]byte(line))
		return nil
	}
	// 时间不等
	hook.file.Close()
	os.MkdirAll(path.Join("logs", timer), os.ModePerm)
	filename := path.Join("logs", timer, "gvd.log")

	hook.file, _ = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	hook.fileDate = timer
	hook.file.Write([]byte(line))
	return nil
}

// 将error级别的日志写入到具体的文件中
type ErrorHook struct {
	file     *os.File
	fileDate string
}

func (hook *ErrorHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (hook *ErrorHook) Fire(entry *logrus.Entry) error {
	timer := entry.Time.Format("2006-01-02")
	line, _ := entry.String()
	if hook.fileDate == timer {
		hook.file.Write([]byte(line))
		return nil
	}
	// 时间不等的话就关掉文件句柄，重新打开一个
	hook.file.Close()
	os.MkdirAll(path.Join("logs", timer), os.ModePerm)
	filename := path.Join("logs", timer, "err.log")

	hook.file, _ = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	hook.fileDate = timer
	hook.file.Write([]byte(line))
	return nil
}

func InitLogger(requestList ...LogRequest) *logrus.Logger {
	var request LogRequest
	if len(requestList) > 0 {
		request = requestList[0]
	}
	mLog := logrus.New()               //新建一个实例
	mLog.SetOutput(os.Stdout)          //设置输出类型
	mLog.SetReportCaller(true)         //开启返回函数名和行号
	mLog.SetFormatter(&LogFormatter{}) //设置自己定义的Formatter
	mLog.SetLevel(logrus.DebugLevel)   //设置最低的Level
	if !request.NoDate {
		mLog.AddHook(&DateHook{})
	}
	if !request.NoErr {
		mLog.AddHook(&ErrorHook{})
	}
	if !request.NoGlobal {
		InitDefaultLogger()
	}
	return mLog
}

// 初始化全局的logrus日志对象，而不是创建一个新的实例
func InitDefaultLogger() {
	//全局log
	//全局的 logrus.Logger 是通过 github.com/sirupsen/logrus 包的包级别函数直接创建的
	logrus.SetOutput(os.Stdout)          //设置输出类型
	logrus.SetReportCaller(true)         //开启返回函数名和行号
	logrus.SetFormatter(&LogFormatter{}) //设置自己定义的Formatter
	logrus.SetLevel(logrus.DebugLevel)   //设置最低的level
}
