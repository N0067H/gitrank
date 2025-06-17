package model

import "github.com/shurcooL/githubv4"

type Organization struct {
	Name            githubv4.String
	MembersWithRole MembersWithRole `graphql:"membersWithRole(first: 100, after: $cursor)"`
}
