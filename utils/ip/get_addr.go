package ip

import (
	"fmt"
	"gvd_server/global"
	"net"
)

// ExternalIp 判断是否是外网地址
func ExternalIp(ip string) (ok bool) {
	IP := net.ParseIP(ip)
	if IP == nil {
		return false
	}

	ip4 := IP.To4()
	if ip4 == nil {
		return false
	}
	if !IP.IsPrivate() && !IP.IsLoopback() {
		return true
	}
	return false

}

func GetAddr(ip string) (addr string) {
	if !ExternalIp(ip) {
		return "内网地址"
	}
	citys, err := global.AddrDB.City(net.ParseIP(ip))
	if err != nil {
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
	if city != "" {
		return fmt.Sprintf("%s-%s", country, city)
	}
	return "未知地址"
}
