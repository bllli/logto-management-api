# CaptchaProvider
(*CaptchaProvider*)

## Overview

### Available Operations

* [Get](#get) - Get captcha provider
* [Update](#update) - Update captcha provider

## Get

Get the captcha provider, you can only have one captcha provider.

### Example Usage

```go
package main

import(
	"context"
	"os"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
    )

    res, err := s.CaptchaProvider.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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

**[*operations.GetCaptchaProviderResponse](../../models/operations/getcaptchaproviderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the captcha provider with the provided settings.

### Example Usage

```go
package main

import(
	"context"
	"os"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"github.com/bllli/logto-management-api/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
    )

    res, err := s.CaptchaProvider.Update(ctx, operations.UpdateCaptchaProviderRequest{
        Config: operations.CreateConfigRequestUnionUpdateCaptchaProviderConfigRequest1(
            operations.UpdateCaptchaProviderConfigRequest1{
                Type: "<value>",
                SiteKey: "<value>",
                SecretKey: "<value>",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateCaptchaProviderRequest](../../models/operations/updatecaptchaproviderrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateCaptchaProviderResponse](../../models/operations/updatecaptchaproviderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |