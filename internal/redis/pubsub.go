package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func Publish(channel, message string) error {
	return Rdb.Publish(context.TODO(), channel, message).Err()
}

func Subscribe(channel string) *redis.PubSub {
	return Rdb.Subscribe(context.TODO(), channel)
}
