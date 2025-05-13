<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"github.com/bllli/logto-management-api/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := logtomanagementapi.New(
		logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
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