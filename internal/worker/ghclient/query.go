package ghclient

import "github.com/shurcooL/githubv4"

type ContributionQuery struct {
	User struct {
		ContributionsCollection struct {
			ContributionCalendar struct {
				TotalContributions githubv4.Int
			}
		}
	} `graphql:"user(login: $login)"`
}

type OrganizationQuery struct {
	Organization struct {
		Name            githubv4.String
		MembersWithRole struct {
			Nodes []struct {
				Login     githubv4.String
				Name      githubv4.String
				Email     githubv4.String
				Company   githubv4.String
				CreatedAt githubv4.DateTime
			}
			PageInfo struct {
				EndCursor   githubv4.String
				HasNextPage bool
			}
		} `graphql:"membersWithRole(first: 100, after: $cursor)"`
	} `graphql:"organization(login: $login)"`
}
