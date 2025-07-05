package redis

import (
	"github.com/n0067h/gitrank/internal/worker/ghclient"
	"time"
)

type Cache struct {
	Users     []ghclient.User `json:"users"`
	ExpiresIn time.Time       `json:"expires_in"`
}
