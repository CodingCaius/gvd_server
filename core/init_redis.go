//初始化redis

package core

import (
	"context"
	"gvd_server/global"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

//初始化 Redis 客户端连接
func InitRedis(db int) *redis.Client {
	redisConf := global.Config.Redis

	//配置 Redis 客户端，并且在内部初始化了连接池等资源，但并不执行实际的连接操作
	client := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		Password: redisConf.Password,
		DB:       db,
		PoolSize: redisConf.PoolSize,
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	// Redis 的 PING 命令，尝试与 Redis 服务器建立连接
	_, err := client.Ping().Result()
	if err != nil {
		logrus.Fatalf("%s redis连接失败 err: %s", redisConf.Addr(), err.Error())
	}
	return client

}
