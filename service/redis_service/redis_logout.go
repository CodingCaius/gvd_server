package redis_service

import (
  "gvd_server/global"
  "time"
)

const prefix = "logout_"

// Logout 设置一个注销的token  expiration 过期时间
func Logout(token string, expiration time.Duration) error {
  err := global.Redis.Set(prefix+token, "", expiration).Err()
  return err
}

// CheckLogout 判断一个token是否是属于注销的token
func CheckLogout(token string) bool {
  _, err := global.Redis.Get(prefix + token).Result()
  return err == nil
}
