package main

import (
	"context"
	"fmt"

	logtomanagementapi "github.com/bllli/logto-management-api"
	"github.com/bllli/logto-management-api/models/components"
)

func main() {
	clientId := "your-m2m-client-id"
	clientSecret := "your-m2m-client-secret"
	// Do not forget to grant the m2m client the LogtoManagementAPI scope
	resource := "https://default.logto.app/api"
	scope := "all"

	client := logtomanagementapi.New(
		logtomanagementapi.WithServerURL("https://<your-logto-instance>"),
		logtomanagementapi.WithSecurity(components.Security{
			ClientID:     &clientId,
			ClientSecret: &clientSecret,
			Resource:     &resource,
			Scope:        &scope,
		}),
	)
	ctx := context.Background()
	orgs, err := client.Organizations.List(ctx, nil, nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(orgs)

	if len(orgs.ResponseBodies) == 0 {
		panic("no organizations found")
	}

	orgId := orgs.ResponseBodies[0].ID

	fmt.Println(orgId)

	if orgId == "" {
		panic("orgId is empty")
	}

	org, err := client.Organizations.Get(ctx, orgId)
	if err != nil {
		panic(err)
	}
	fmt.Println(org)
}
