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

func CacheRanking(rdb *redis.Client, updateChannel *redis.PubSub) {
	for {
		ReceiveMessage(updateChannel)

		locked, err := isRankingLocked(rdb)
		if err != nil {
			log.Error(err)
			continue
		}
		if locked {
			continue
		}

		users := ghclient.GetRanking()
		cache := &Cache{
			Users:     users,
			ExpiresIn: time.Now().Add(3 * time.Minute),
		}

		if err = saveRanking(rdb, cache); err != nil {
			log.Errorf("failed to cache ranking: %v", err)
			continue
		}

		if err = lockRanking(rdb); err != nil {
			log.Errorf("failed to cache ranking: %v", err)
		}

		time.Sleep(3 * time.Minute)
	}
}

func saveRanking(rdb *redis.Client, cache *Cache) error {
	val, err := json.Marshal(cache)
	if err != nil {
		return fmt.Errorf("failed to marshal cache: %w", err)
	}

	err = rdb.Set(context.TODO(), "ranking", string(val), 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set ranking: %w", err)
	}

	return nil
}

func isRankingLocked(rdb *redis.Client) (bool, error) {
	_, err := rdb.Get(context.TODO(), "ranking:proceed").Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf("failed to get ranking proceed: %w", err)
	}

	return true, nil
}

func lockRanking(rdb *redis.Client) error {
	err := rdb.Set(context.TODO(), "ranking:proceed", "", 3*time.Minute).Err()
	if err != nil {
		return fmt.Errorf("failed to save ranking: %w", err)
	}

	return nil
}

func FetchRankingCache(rdb *redis.Client) (*Cache, error) {
	val, err := rdb.Get(context.Background(), "ranking").Result()
	if err == redis.Nil {
		return nil, redis.Nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get ranking: %w", err)
	}

	var cache Cache
	if err := json.Unmarshal([]byte(val), &cache); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ranking cache: %w", err)
	}

	return &cache, nil
}

func TryPublishRankingUpdate(rdb *redis.Client) bool {
	locked, err := isRankingLocked(rdb)
	if err != nil {
		log.Errorf("failed to check ranking lock: %v", err)
		return false
	}

	if locked {
		return false
	}

	if err := Publish(rdb, "ranking:update_request", "update"); err != nil {
		log.Errorf("failed to publish ranking update request: %v", err)
		return false
	}

	return true
}
