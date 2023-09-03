//core文件夹用于初始化连接，读取配置文件

//该文件用于读取配置文件setting.yaml
//并将读取到的信息存储到config文件夹的Config结构体中

//其余文件用于初始化系统连接

// 逻辑上也相当于初始化Config
// 这个文件可以放在config文件夹下，core文件夹全部放置初始化连接的文件
package core

import (
	"gvd_server/config"
	"gvd_server/global"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const yamlPath = "setting.yaml"

func InitConfig() (c *config.Config) {
	byteData, err := os.ReadFile(yamlPath)
	if err != nil {
		logrus.Fatalln("read yaml err: ", err.Error())
	}

	c = new(config.Config)
	err = yaml.Unmarshal(byteData, c)
	if err != nil {
		logrus.Fatalln("解析 yaml err: ", err.Error())
	}

	return c
}

func SetYaml() {
	//序列化，将配置信息写入yaml文件
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		global.Log.Error(err)
		return
	}
	err = os.WriteFile(yamlPath, byteData, 066)
	if err != nil {
		global.Log.Error(err)
		return
	}
}
