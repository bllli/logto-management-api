# SsoConnectorProviders
(*SsoConnectorProviders*)

## Overview

### Available Operations

* [List](#list) - List all the supported SSO connector provider details

## List

Get a complete list of supported SSO connector providers.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/bllli/logto-management-api/models/components"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(components.Security{
            ClientID: logtomanagementapi.String(os.Getenv("LOGTOMANAGEMENTAPI_CLIENT_ID")),
            ClientSecret: logtomanagementapi.String(os.Getenv("LOGTOMANAGEMENTAPI_CLIENT_SECRET")),
        }),
    )

    res, err := s.SsoConnectorProviders.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListSsoConnectorProvidersResponse](../../models/operations/listssoconnectorprovidersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |