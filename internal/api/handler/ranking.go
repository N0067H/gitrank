package handler

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	myredis "github.com/n0067h/gitrank/internal/redis"
	"github.com/n0067h/gitrank/internal/worker/ghclient"
	"github.com/redis/go-redis/v9"
)

func GetRanking(c *fiber.Ctx) error {
	ranking, err := myredis.GetRanking()
	if errors.Is(err, redis.Nil) {
		if err := myredis.Publish("ranking:update_request", "update"); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendStatus(fiber.StatusAccepted)
	} else if err != nil {
		log.Errorf("redis.GetRanking(): %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var result []ghclient.User
	if err := json.Unmarshal([]byte(ranking), &result); err != nil {
		log.Errorf("Failed to unmarshal ranking data: %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(result)
}
