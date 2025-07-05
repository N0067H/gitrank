package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0067h/gitrank/internal/worker/ghclient"
	"github.com/redis/go-redis/v9"
	"time"
)

func GetRanking() (string, error) {
	val, err := Rdb.Get(context.TODO(), "ranking").Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func SetRanking(cache *Cache) error {
	data, err := json.Marshal(cache)
	if err != nil {
		log.Fatal(err)
	}

	return Rdb.Set(context.TODO(), "ranking", string(data), 0).Err()
}

func CacheRanking(pubsub *redis.PubSub) error {
	for {
		_, err := ReceiveMessage(pubsub)
		if err != nil {
			return fmt.Errorf("failed to receive message: %w", err)
		}
		log.Info("Received update request, fetching new ranking")

		users := ghclient.GetRanking()
		cache := &Cache{
			Users:     users,
			ExpiresIn: time.Now().Add(3 * time.Minute),
		}

		err = SetRanking(cache)
		if err != nil {
			return fmt.Errorf("failed to set ranking: %w", err)
		}
		log.Info("New ranking cached")
	}
}
