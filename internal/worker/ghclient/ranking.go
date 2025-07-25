package ghclient

import (
	"github.com/shurcooL/githubv4"
	"sort"
)

func GetUsers() []User {
	var users []User

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

		users = append(users, User{Login: node.Login, TotalContributions: totalContributions})
	}

	return users
}

func GetRanking(users []User) []User {
	sort.Slice(users, func(i, j int) bool {
		return users[i].TotalContributions > users[j].TotalContributions
	})
	return users
}
