package ghclient

import (
	"context"
	"github.com/gbswhs/gbsw-gitrank/worker/model"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"os"
)

func GetClient() *githubv4.Client {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: os.Getenv("GITHUB_TOKEN"),
		})

	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)
	return client
}

func QueryOrganization(client *githubv4.Client, query *model.OrganizationQuery, cursor *githubv4.String) error {
	return client.Query(context.Background(),
		&query,
		map[string]interface{}{
			"login":  githubv4.String(os.Getenv("ORGANIZATION_NAME")),
			"cursor": cursor,
		})
}

func QueryContributions(client *githubv4.Client, query *model.ContributionQuery, login githubv4.String) error {
	return client.Query(
		context.Background(),
		&query,
		map[string]interface{}{
			"login": githubv4.String(login),
		})
}
