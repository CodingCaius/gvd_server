package utils

import "strings"

// FindMaxPrefix 最大前缀匹配
func FindMaxPrefix(target string, list []string) (maxPrefix string, index int) {
	index = -1
	if len(list) == 0 {
		return
	}
	for i, s := range list {
		// 1.2
		// 1.3
		// 1.3.4
		// 6.7

		// 1.3.4.5
		if strings.HasPrefix(target, s) && len(s) > len(maxPrefix) {
			maxPrefix = s
			index = i
		}
	}
	return
}