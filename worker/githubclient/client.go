package githubclient

import (
	"io"
	"net/http"
)

const apiRoot = "https://api.github.com"

type GithubClient struct {
	apiKey string
}

func New(apiKey string) GithubClient {
	return GithubClient{
		apiKey,
	}
}

func (c GithubClient) setHeaders(req *http.Request) {
	req.Header.Set("Authorization", "bearer "+c.apiKey)
	req.Header.Set("Accept", "application/vnd.github+json")
}

func (c GithubClient) doRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, ErrCreatingRequestFailed
	}

	c.setHeaders(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, ErrInternalServerError
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrGithubAPIError
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrReadingResponseFailed
	}

	return body, nil
}
