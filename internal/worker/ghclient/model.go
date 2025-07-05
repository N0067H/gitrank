package ghclient

import "github.com/shurcooL/githubv4"

type User struct {
	Login              githubv4.String `json:"login"`
	TotalContributions githubv4.Int    `json:"totalContributions"`
}
