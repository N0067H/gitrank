package githubclient

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID    string `json:"id"`
	Login string `json:"login"`
}

func (c GithubClient) GetUsers(org string, perPage, page int) ([]User, error) {
	url := apiRoot + fmt.Sprintf("/orgs/%s/members?filter=all&per_page=%d&page=%d", org, perPage, page)
	body, err := c.doRequest(url)
	if err != nil {
		return nil, ErrGithubAPIError
	}

	var members []User
	json.Unmarshal(body, &members)
	return members, nil
}
