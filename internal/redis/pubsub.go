package redis

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"github.com/redis/go-redis/v9"
)

func Publish(rdb *redis.Client, channel, message string) error {
	return rdb.Publish(context.TODO(), channel, message).Err()
}

func Subscribe(rdb *redis.Client, channel string) *redis.PubSub {
	pubsub := rdb.Subscribe(context.TODO(), channel)
	if pubsub == nil {
		log.Fatal("Failed to subscribe to channel " + channel)
	}

	return pubsub
}

func ReceiveMessage(pubsub *redis.PubSub) {
	_, err := pubsub.ReceiveMessage(context.TODO())
	if err != nil {
		log.Errorf("Failed to receive message: %v", err)
	}
}
