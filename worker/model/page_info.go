package model

import "github.com/shurcooL/githubv4"

type pageInfo struct {
	EndCursor   githubv4.String
	HasNextPage bool
}
