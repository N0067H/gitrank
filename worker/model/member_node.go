package model

import "github.com/shurcooL/githubv4"

type MemberNode struct {
	Login     githubv4.String
	Name      githubv4.String
	Email     githubv4.String
	Company   githubv4.String
	CreatedAt githubv4.DateTime
}
