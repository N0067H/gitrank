package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	myredis "github.com/n0067h/gitrank/internal/redis"
	"github.com/redis/go-redis/v9"
	"time"
)

func GetRanking(rdb *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cache, err := myredis.GetRanking(rdb)
		if err != nil {
			if err == redis.Nil {
				handled := myredis.TryPublishRankingUpdate(rdb)
				if !handled {
					log.Warn("skipped publishing ranking update: update already in progress")
				}
				return c.SendStatus(fiber.StatusAccepted)
			}

			log.Errorf("failed to get ranking: %v", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		if cache.ExpiresIn.Before(time.Now()) {
			go func() {
				if !myredis.TryPublishRankingUpdate(rdb) {
					log.Warn("skipped publishing ranking update: update already in progress")
				}
			}()
		}

		return c.JSON(cache)
	}
}
