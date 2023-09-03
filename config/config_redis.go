package config

import "fmt"

type Redis struct {
	Ip string `yaml:"ip"`
	Port int `yaml:"port"`
	Password string `yaml:"password"`
	PoolSize int `yaml:"poolsize"`
}

func (redis Redis) Addr() string {
	return fmt.Sprintf("%s:%d", redis.Ip, redis.Port)
}