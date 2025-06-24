package ghclient

import (
	"github.com/shurcooL/githubv4"
	"sort"
)

// Temporary cache store
var Users []User

type User struct {
	Login              githubv4.String
	TotalContributions githubv4.Int
}

func GetRanks() {
	// Cache check
	if len(Users) > 0 {
		return
	}

	client := GetClient()
	cursor := (*githubv4.String)(nil)

	var orgQuery OrganizationQuery
	err := QueryOrganization(client, &orgQuery, cursor)
	if err != nil {
		panic(err)
	}

	org := orgQuery.Organization
	nodes := org.MembersWithRole.Nodes

	for org.MembersWithRole.PageInfo.HasNextPage {
		err := QueryOrganization(client, &orgQuery, cursor)
		if err != nil {
			panic(err)
		}

		newOrg := orgQuery.Organization
		newNodes := newOrg.MembersWithRole.Nodes

		nodes = append(nodes, newNodes...)
	}

	for _, node := range nodes {
		contribQuery := new(ContributionQuery)
		err := QueryContributions(client, contribQuery, node.Login)
		if err != nil {
			panic(err)
		}

		contributionCollection := contribQuery.User.ContributionsCollection
		contributionCalendar := contributionCollection.ContributionCalendar
		totalContributions := contributionCalendar.TotalContributions

		Users = append(Users, User{Login: node.Login, TotalContributions: totalContributions})
	}

	sort.Slice(Users, func(i, j int) bool {
		return Users[i].TotalContributions > Users[j].TotalContributions
	})
}
