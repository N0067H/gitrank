package model

type ContributionQuery struct {
	User user `graphql:"user(login: $login)"`
}
