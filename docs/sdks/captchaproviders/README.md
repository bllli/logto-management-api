# CaptchaProviders
(*CaptchaProviders*)

## Overview

### Available Operations

* [Delete](#delete) - Delete captcha provider

## Delete

Delete the captcha provider.

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

    res, err := s.CaptchaProviders.Delete(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
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

**[*operations.DeleteCaptchaProviderResponse](../../models/operations/deletecaptchaproviderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |