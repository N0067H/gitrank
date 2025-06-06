package main

import (
	"fmt"
	"log"

	"github.com/gbswhs/gbsw-gitrank/worker/githubclient"
)

var apiKey string

func main() {
	client := githubclient.New(apiKey)

	users, err := client.GetUsers("gbswhs", 100, 1)
	if err != nil {
		log.Fatalf("WTF: %v", err)
	}

	for i, e := range users {
		fmt.Printf("User %d: %s\n", i, e.Login)
	}

	events, err := client.GetEvents("n0067h")
	if err != nil {
		log.Fatalf("WTF: %v", err)
	}

	score := 0
	for _, e := range events {
		switch e.Type {
		case "PushEvent":
			score += 1
		case "IssueEvent":
			score += 2
		case "PullRequestEvent":
			score += 3
		}

		fmt.Printf("Actor: %s\nRepo: %s\nType:%s\nScore:%d\n\n", e.Actor.Login, e.Repo.Name, e.Type, score)
	}
}
