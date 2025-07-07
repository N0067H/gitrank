package redis

import (
	"context"
	"fmt"
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
	ps, err := pubsub.ReceiveMessage(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to receive message: %w", err)
	}

	return ps, nil
}
