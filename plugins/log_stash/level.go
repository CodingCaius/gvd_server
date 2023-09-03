package log_stash

import "encoding/json"

type Level int

const (
	Info    Level = 1
	Warning Level = 2
	Error   Level = 3
)

// 转字符串
func (level Level) String() string {
	switch level {
	case Info:
		return "info"
	case Warning:
		return "warning"
	case Error:
		return "error"
	}
	return ""
}

//自定义类型转化为json字节切片
func (level Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(level.String())
}
