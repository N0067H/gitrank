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

func CacheRanking(pubsub *redis.PubSub) error {
	for {
		_, err := ReceiveMessage(pubsub)
		if err != nil {
			return err
		}

		proceed, err := GetRankingProceed()
		if err != nil {
			log.Error(err)
			continue
		}
		if proceed {
			continue
		}

		users := ghclient.GetRanking()
		cache := &Cache{
			Users:     users,
			ExpiresIn: time.Now().Add(3 * time.Minute),
		}

		err = SetRanking(cache)
		if err != nil {
			return err
		}

		err = SetRankingProceed()
		if err != nil {
			return err
		}
	}
}

func GetRanking() (string, error) {
	val, err := Rdb.Get(context.TODO(), "ranking").Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func SetRanking(cache *Cache) error {
	val, err := json.Marshal(cache)
	if err != nil {
		return fmt.Errorf("failed to marshal cache: %w", err)
	}

	err = Rdb.Set(context.TODO(), "ranking", string(val), 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set ranking: %w", err)
	}

	return nil
}

func GetRankingProceed() (bool, error) {
	_, err := Rdb.Get(context.TODO(), "ranking:proceed").Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf("failed to get ranking proceed: %w", err)
	}

	return true, nil
}

func SetRankingProceed() error {
	err := Rdb.Set(context.TODO(), "ranking:proceed", "", 3*time.Minute).Err()
	if err != nil {
		return fmt.Errorf("failed to set ranking proceed: %w", err)
	}

	return nil
}
