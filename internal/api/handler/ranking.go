package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	myredis "github.com/n0067h/gitrank/internal/redis"
	"github.com/redis/go-redis/v9"
	"time"
)

func GetRanking(c *fiber.Ctx) error {
	ranking, err := myredis.GetRanking()

	if err == redis.Nil {
		proceed, err := myredis.GetRankingProceed()
		if err != nil {
			log.Errorf("Failed to get ranking proceed(): %v", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		if !proceed {
			if err := myredis.Publish("ranking:update_request", "update"); err != nil {
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		return c.SendStatus(fiber.StatusAccepted)
	} else if err != nil {
		log.Errorf("Failed to get ranking(): %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var cache myredis.Cache
	if err := json.Unmarshal([]byte(ranking), &cache); err != nil {
		log.Errorf("Failed to unmarshal ranking data: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if cache.ExpiresIn.Before(time.Now()) {
		if err := myredis.Publish("ranking:update_request", "update"); err != nil {
			log.Errorf("Failed to publish update request: %v", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}

	return c.JSON(cache)
}
