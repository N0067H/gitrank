package redis

import (
	"github.com/n0067h/gitrank/internal/config"
	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func Init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.RedisHost + ":" + config.AppConfig.RedisPort,
		Password: config.AppConfig.RedisPassword,
		DB:       0,
	})
}
