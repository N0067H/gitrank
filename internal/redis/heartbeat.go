package redis

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/redis/go-redis/v9"
	"time"
)

var workerStatus bool = false

func CheckHeartbeat(rdb *redis.Client) {
	heartbeat, err := GetHeartbeat(rdb)
	if err != nil {
		log.Error(err)
		if workerStatus {
			log.Warn("Worker status unknown; Retrying in 20 seconds.")
			workerStatus = false
		}
	} else if !heartbeat {
		workerStatus = false
		log.Warn("Worker is not running; Retrying in 20 seconds.")
	} else if !workerStatus {
		workerStatus = true
		log.Info("Worker is running.")
	}
}

func GetHeartbeat(rdb *redis.Client) (bool, error) {
	_, err := rdb.Get(context.TODO(), "heartbeat").Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf("failed to get heartbeat: %w", err)
	}

	return true, nil
}

func SetHeartbeat(rdb *redis.Client) {
	err := rdb.Set(context.TODO(), "heartbeat", "", 10*time.Second).Err()
	if err != nil {
		log.Errorf("Failed to set heartbeat: %v", err)
	}
}
