package main

import (
	"context"
	"encoding/json"
	"github.com/n0067h/gitrank/internal/config"
	myredis "github.com/n0067h/gitrank/internal/redis"
	"github.com/n0067h/gitrank/internal/worker/ghclient"
	"log"
	"time"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}
	log.Println("Load .env")

	myredis.Init()
	log.Println("Init cache")

	pubsub := myredis.Subscribe("update_request")
	for {
		_, err := pubsub.ReceiveMessage(context.Background())
		if err != nil {
			panic(err)
		}

		users := ghclient.GetRanks()
		data, err := json.Marshal(users)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Cached: ", string(data)[:100])
		if err := myredis.Rdb.Set(context.TODO(), "ranking", string(data), time.Minute*30).Err(); err != nil {
			panic(err)
		}
	}
}
