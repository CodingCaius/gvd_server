package hash

import (
  "crypto/md5"
  "fmt"
  "io"
  "mime/multipart"
)


// Md5 计算md5值
func Md5(byteData []byte) string {
  hash := md5.New()
  hash.Write(byteData)
  hashByteData := hash.Sum(nil)
  return fmt.Sprintf("%x", hashByteData)
}

// FileMd5 计算上传文件的md5值
func FileMd5(file multipart.File) string {
  hash := md5.New()
  // 这里用到了copy方法，千万不要直接读取原file对象
  io.Copy(hash, file)
  hashByteData := hash.Sum(nil)
  //%x 将整数或字节切片转化成十六进制表示的字符串
  return fmt.Sprintf("%x", hashByteData)
}
