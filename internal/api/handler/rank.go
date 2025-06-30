package handler

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	myredis "github.com/n0067h/gitrank/internal/redis"
	"github.com/redis/go-redis/v9"
)

type Contributer struct {
	Login              string `json:"Login"`
	TotalContributions int    `json:"TotalContributions"`
}

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

	var result []Contributer
	if err := json.Unmarshal([]byte(ranking), &result); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(result)
}
