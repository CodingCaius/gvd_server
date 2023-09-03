package log_stash

import "gvd_server/global"

//运行日志可以把一段时间内的数据放在一起，所以创建运行日志的时候
//要先查询一下，如果该日志已经存在就不需要再创建

// NewRuntime 创建一个运行日志的log
func NewRuntime(serviceName string) *Action {
	log := &Action{
		serviceName: serviceName,
		logType:     RuntimeType,
	}
	var logModel LogModel
	err := global.DB.Take(&logModel,
		"type = ? and to_days(createdAt) = to_days(now()) and serviceName = ?", RuntimeType, serviceName).Error
	//如果存在的话就不用创建，直接赋值
	if err == nil {
		log.model = &logModel
	}
	return log
}
