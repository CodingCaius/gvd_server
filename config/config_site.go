package config

type Site struct {
	Title    string `json:"title" yaml:"title"`       // 网站名称
	Icon     string `json:"icon" yaml:"icon"`         // 首页的图标
	Abstract string `json:"abstract" yaml:"abstract"` // 网站简介
	IconHref string `json:"iconHref" yaml:"iconHref"` // 图标链接
	Href     string `json:"href" yaml:"href"`         // 点击go的跳转链接
	Footer   string `json:"footer" yaml:"footer"`     // 尾部信息
	Content  string `json:"content" yaml:"content"`   // 内容
}
