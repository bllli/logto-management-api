package main

import (
	"context"
	"fmt"

	logtomanagementapi "github.com/bllli/logto-management-api"
	"github.com/bllli/logto-management-api/tokenmanager"
)

func main() {
	var config = &tokenmanager.LogtoM2MConfig{
		ServerURL: "https://your-logto-instance.com",
		AppID:     "your-m2m-client-id",
		AppSecret: "your-m2m-client-secret",
		Resource:  "https://default.logto.app/api",
		Scope:     "all", // Do not forget to grant the m2m client the LogtoManagementAPI scope
	}

	client := logtomanagementapi.NewLogtoClientWithTokenManager(config)
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
