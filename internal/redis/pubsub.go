package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

func Publish(channel, message string) error {
	return Rdb.Publish(context.TODO(), channel, message).Err()
}

func Subscribe(channel string) *redis.PubSub {
	pubsub := Rdb.Subscribe(context.TODO(), channel)
	if pubsub == nil {
		log.Fatal("failed to subscribe to channel 'update_request'")
	}

	return pubsub
}

func ReceiveMessage(pubsub *redis.PubSub) (*redis.Message, error) {
	return pubsub.ReceiveMessage(context.TODO())
}
