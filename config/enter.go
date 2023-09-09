//该文件夹用于存放解析到的配置
//全部是系统的配置信息

package config

type Config struct {
	System System `yaml:"system"`
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	ES     Es     `yaml:"es"`
	Jwt    Jwt    `yaml:"jwt"`
	Site   Site   `yaml:"site"`
}
