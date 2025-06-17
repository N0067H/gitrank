package model

type OrganizationQuery struct {
	Organization Organization `graphql:"organization(login: $login)"`
}
