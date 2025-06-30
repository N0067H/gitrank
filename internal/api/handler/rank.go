package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	myredis "github.com/n0067h/gitrank/internal/redis"
	"github.com/redis/go-redis/v9"
)

func GetRank(c *fiber.Ctx) error {
	// cache check
	ranking, err := myredis.GetRanking()
	if errors.Is(err, redis.Nil) {
		if err := myredis.Publish("update_request", "update"); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendStatus(fiber.StatusAccepted)
	} else if err != nil {
		log.Fatalf("redis.GetRanking(): %v", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendString(ranking)
}
