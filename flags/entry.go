package flags

import "flag"

type Option struct {
	DB     bool   //初始化数据库
	Port   int    //改端口
	Load   string //导入数据库文件
	Dump   bool   // 导出数据库
	Es     bool   // 创建索引
	ESDump bool   // 导出es索引
	ESLoad string // 导入es索引
}

func Parse() (option *Option) {
	option = new(Option)
	flag.BoolVar(&option.DB, "db", false, "初始化数据库")
	flag.BoolVar(&option.Es, "es", false, "创建索引")
	flag.IntVar(&option.Port, "port", 0, "程序运行的端口")
	flag.StringVar(&option.Load, "load", "", "导入sql数据库")
	flag.BoolVar(&option.Dump, "dump", false, "导出sql数据库")
	flag.BoolVar(&option.ESDump, "esdump", false, "导出es索引")
	flag.StringVar(&option.ESLoad, "esload", "", "导入es索引")
	flag.Parse()
	return option
}

// 根据里面的参数运行不同的脚本
func (option Option) Run() bool {
	if option.DB {
		DB()
		return true
	}
	if option.Port != 0 {
		Port(option.Port)
		return false
	}
	if option.Load != "" {
		Load(option.Load)
		return true
	}
	if option.Es {
		ESIndex()
		return true
	}
	if option.Dump {
		Dump()
		return true
	}
	if option.ESDump {
		ESDump()
		return true
	}
	if option.ESLoad != "" {
		ESLoad(option.ESLoad)
		return true
	}

	return false
}
