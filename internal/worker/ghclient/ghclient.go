package ghclient

import (
	"context"
	"github.com/gbswhs/gbsw-gitrank/internal/protobuf/rank"
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

func QueryOrganization(client *githubv4.Client, query *OrganizationQuery, cursor *githubv4.String) error {
	return client.Query(context.Background(),
		&query,
		map[string]interface{}{
			"login":  githubv4.String(os.Getenv("ORGANIZATION_NAME")),
			"cursor": cursor,
		})
}

func QueryContributions(client *githubv4.Client, query *ContributionQuery, login githubv4.String) error {
	return client.Query(
		context.Background(),
		&query,
		map[string]interface{}{
			"login": githubv4.String(login),
		})
}

func ToProtoUsers(users []User) []*rank.User {
	var protoUsers []*rank.User

	for _, user := range users {
		protoUsers = append(protoUsers, &rank.User{
			Login:              string(user.Login),
			TotalContributions: uint32(user.TotalContributions),
		})
	}

	return protoUsers
}
