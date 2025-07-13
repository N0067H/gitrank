package redis

import (
	"github.com/n0067h/gitrank/internal/worker/ghclient"
	"time"
)

type Cache struct {
	Ranking   []ghclient.User `json:"ranking"`
	ExpiresIn time.Time       `json:"expires_in"`
}
