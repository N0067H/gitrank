package githubclient

import (
	"encoding/json"
	"fmt"
)

type Event struct {
	ID    int    `json:"id"`
	Type  string `json:"type"`
	Actor User   `json:"actor"`
	Repo  Repo   `json:"repo"`
}

type Actor struct {
	ID        int    `json:"id"`
	Login     string `json:"login"`
	URL       string `json:"url"`
	AvatarURL string `json:"avatar_url"`
}

type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (c GithubClient) GetEvents(username string) ([]Event, error) {
	url := apiRoot + fmt.Sprintf("/users/%s/events", username)
	body, err := c.doRequest(url)
	if err != nil {
		return nil, err
	}

	var events []Event
	json.Unmarshal(body, &events)
	return events, nil
}
