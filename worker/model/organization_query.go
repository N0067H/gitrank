package model

type OrganizationQuery struct {
	Organization organization `graphql:"organization(login: $login)"`
}
