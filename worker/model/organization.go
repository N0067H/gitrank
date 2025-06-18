package model

import "github.com/shurcooL/githubv4"

type organization struct {
	Name            githubv4.String
	MembersWithRole membersWithRole `graphql:"membersWithRole(first: 100, after: $cursor)"`
}
