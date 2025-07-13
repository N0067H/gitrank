package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/n0067h/gitrank/internal/config"
	"github.com/n0067h/gitrank/internal/worker/ghclient"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func saveUsers(rdb *redis.Client, users []ghclient.User) error {
	val, err := json.Marshal(users)
	if err != nil {
		return fmt.Errorf("failed to marshal users: %w", err)
	}

	ttl, err := time.ParseDuration(config.AppConfig.UserCacheTTL)
	if err != nil {
		log.Fatalf("Failed to parse user cache TTL: %s", err)
	}
	rdb.Set(context.TODO(), "users", val, ttl)

	return nil
}

func getUsers(rdb *redis.Client) ([]ghclient.User, error) {
	val, err := rdb.Get(context.TODO(), "users").Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	var users []ghclient.User
	err = json.Unmarshal([]byte(val), &users)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal users: %w", err)
	}

	return users, nil
}
