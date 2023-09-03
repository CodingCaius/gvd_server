package file

import "fmt"

// FormatBytes 格式化输出字节单位
func FormatBytes(size int64) string {
	_size := float64(size)
	uints := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	// 1
	// 1025 1.0KB
	//
	var i int = 0
	for _size >= 1024 && i < len(uints)-1 {
		_size /= 1024
		i++
	}
	return fmt.Sprintf("%.2f %s", _size, uints[i])

}