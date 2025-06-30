package redis

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func GetRanking() (string, error) {
	val, err := Rdb.Get(context.TODO(), "ranking").Result()
	if errors.Is(err, redis.Nil) {
		log.Println("Cache miss: Ranking")
		return "", err
	} else if err != nil {
		return "", err
	}

	return val, nil
}

func SetRanking(data string) error {
	return Rdb.Set(context.TODO(), "ranking", data, time.Minute*30).Err()
}
