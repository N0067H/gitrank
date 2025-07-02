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

func SetRanking(users []ghclient.User) error {
	data, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	return Rdb.Set(context.TODO(), "ranking", string(data), time.Minute*30).Err()
}

func CacheRanking(pubsub *redis.PubSub) error {
	for {
		_, err := ReceiveMessage(pubsub)
		if err != nil {
			return fmt.Errorf("failed to receive message: %w", err)
		}
		log.Info("Received update request, fetching new ranking")

		users := ghclient.GetRanking()
		err = SetRanking(users)
		if err != nil {
			return fmt.Errorf("failed to set ranking: %w", err)
		}
		log.Info("New ranking cached")
	}
}
