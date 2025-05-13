<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"github.com/bllli/logto-management-api/models/components"
	"github.com/bllli/logto-management-api/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := logtomanagementapi.New(
		logtomanagementapi.WithSecurity(components.Security{
			ClientID:     logtomanagementapi.String(os.Getenv("LOGTOMANAGEMENTAPI_CLIENT_ID")),
			ClientSecret: logtomanagementapi.String(os.Getenv("LOGTOMANAGEMENTAPI_CLIENT_SECRET")),
		}),
	)

	res, err := s.Applications.Get(ctx, operations.ListApplicationsRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.ResponseBodies != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->