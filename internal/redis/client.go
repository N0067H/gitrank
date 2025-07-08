package redis

import (
	"context"
	"github.com/n0067h/gitrank/internal/config"
	"github.com/redis/go-redis/v9"
)

func Init() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.RedisHost + ":" + config.AppConfig.RedisPort,
		Password: config.AppConfig.RedisPassword,
		DB:       0,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil
	}

	return rdb
}
