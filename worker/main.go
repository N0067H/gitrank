package main

import (
	"fmt"
	"github.com/gbswhs/gbsw-gitrank/worker/ghclient"
	"github.com/gbswhs/gbsw-gitrank/worker/model"
	"github.com/shurcooL/githubv4"
)

func main() {
	cursor := (*githubv4.String)(nil)
	client := ghclient.GetClient()

	orgQuery := new(model.OrganizationQuery)
	err := ghclient.QueryOrganization(client, orgQuery, cursor)
	if err != nil {
		panic(err)
	}

	org := orgQuery.Organization
	nodes := org.MembersWithRole.Nodes
	
	for org.MembersWithRole.PageInfo.HasNextPage {
		err := ghclient.QueryOrganization(client, orgQuery, cursor)
		if err != nil {
			panic(err)
		}

		newOrg := orgQuery.Organization
		newNodes := newOrg.MembersWithRole.Nodes
		nodes = append(nodes, newNodes...)
	}

	for _, node := range nodes {
		contribQuery := new(model.ContributionQuery)
		err := ghclient.QueryContributions(client, contribQuery, node.Login)
		if err != nil {
			panic(err)
		}

		contributionCollection := contribQuery.User.ContributionsCollection
		contributionCalendar := contributionCollection.ContributionCalendar
		totalContributions := contributionCalendar.TotalContributions
		fmt.Printf("Name: %s / Total: %d\n", node.Login, totalContributions)
	}
}
