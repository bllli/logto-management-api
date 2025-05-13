# Sso
(*Interaction.Sso*)

## Overview

### Available Operations

* [PostAuthorizationURL](#postauthorizationurl)

## PostAuthorizationURL

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

    res, err := s.Interaction.Sso.PostAuthorizationURL(ctx, "<id>", operations.PostAPIInteractionSingleSignOnConnectorIDAuthorizationURLRequestBody{
        State: "Kentucky",
        RedirectURI: operations.PostAPIInteractionSingleSignOnConnectorIDAuthorizationURLRedirectURI{},
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

| Parameter                                                                                                                                                                          | Type                                                                                                                                                                               | Required                                                                                                                                                                           | Description                                                                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                              | :heavy_check_mark:                                                                                                                                                                 | The context to use for the request.                                                                                                                                                |
| `connectorID`                                                                                                                                                                      | *string*                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                 | The unique identifier of the connector.                                                                                                                                            |
| `requestBody`                                                                                                                                                                      | [operations.PostAPIInteractionSingleSignOnConnectorIDAuthorizationURLRequestBody](../../models/operations/postapiinteractionsinglesignonconnectoridauthorizationurlrequestbody.md) | :heavy_check_mark:                                                                                                                                                                 | N/A                                                                                                                                                                                |
| `opts`                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                 | The options for this request.                                                                                                                                                      |

### Response

**[*operations.PostAPIInteractionSingleSignOnConnectorIDAuthorizationURLResponse](../../models/operations/postapiinteractionsinglesignonconnectoridauthorizationurlresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |