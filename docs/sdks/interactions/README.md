# Interactions
(*Interactions*)

## Overview

### Available Operations

* [VerifySocialAuthorizationURI](#verifysocialauthorizationuri)

## VerifySocialAuthorizationURI

### Example Usage

```go
package main

import(
	"context"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"github.com/bllli/logto-management-api/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New()

    res, err := s.Interactions.VerifySocialAuthorizationURI(ctx, operations.PostAPIInteractionVerificationSocialAuthorizationURIRequest{
        ConnectorID: "<id>",
        State: "Utah",
        RedirectURI: operations.PostAPIInteractionVerificationSocialAuthorizationURIRedirectURI{},
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

| Parameter                                                                                                                                                        | Type                                                                                                                                                             | Required                                                                                                                                                         | Description                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                            | :heavy_check_mark:                                                                                                                                               | The context to use for the request.                                                                                                                              |
| `request`                                                                                                                                                        | [operations.PostAPIInteractionVerificationSocialAuthorizationURIRequest](../../models/operations/postapiinteractionverificationsocialauthorizationurirequest.md) | :heavy_check_mark:                                                                                                                                               | The request object to use for the request.                                                                                                                       |
| `opts`                                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                                         | :heavy_minus_sign:                                                                                                                                               | The options for this request.                                                                                                                                    |

### Response

**[*operations.PostAPIInteractionVerificationSocialAuthorizationURIResponse](../../models/operations/postapiinteractionverificationsocialauthorizationuriresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |