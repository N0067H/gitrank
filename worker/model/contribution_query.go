package model

type ContributionQuery struct {
	User User `graphql:"user(login: $login)"`
}
