package main

import (
	"fmt"
	"github.com/cc14514/go-geoip2"
	geoip2db "github.com/cc14514/go-geoip2-db"
	"github.com/sirupsen/logrus"
	"net"
)

var addrDB *geoip2.DBReader

func main() {
	db, err := geoip2db.NewGeoipDbByStatik()
	if err != nil {
		logrus.Fatal("ip地址数据库加载失败", err)
	}
	addrDB = db

	fmt.Println(GetAddr("143.47.226.4"))

}

// ExternalIp 判断是否是公网地址
func ExternalIp(ip string) (ok bool) {
	IP := net.ParseIP(ip)
	if IP == nil {
		return false
	}

	// 将 IP 转换为 IPv4 地址对象 ip4。如果 IP 不是 IPv4 地址，ip4 也将为 nil
	ip4 := IP.To4()
	if ip4 == nil {
		return false
	}
	// 检查 IP 地址是否既不是私有地址也不是回环地址
	if !IP.IsPrivate() && !IP.IsLoopback() {
		return true
	}
	return false

}

// GetAddr 根据给定的 IP 地址获取其所属的地址信息，包括国家、城市、省份
func GetAddr(ip string) (addr string) {
	// 判断ip是否为公网地址
	if!ExternalIp(ip) {
		// 如果不是，返回内网地址
		return "内网地址"
	}
	// 获取ip地址所属的城市
	citys, err := addrDB.City(net.ParseIP(ip))
	if err!= nil {
		fmt.Println(err)
		return
	}
	// 国家
	country := citys.Country.Names["zh-CN"]
	// 城市
	city := citys.City.Names["zh-CN"]
	// 省份
	var subdivisions string
	if len(citys.Subdivisions) > 0 {
		subdivisions = citys.Subdivisions[0].Names["zh-CN"]
		return fmt.Sprintf("%s-%s", subdivisions, city)
	}
	if city!= "" {
		return fmt.Sprintf("%s-%s", country, city)
	}
	return "未知地址"
}