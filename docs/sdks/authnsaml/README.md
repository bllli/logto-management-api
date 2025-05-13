# AuthnSaml
(*AuthnSaml*)

## Overview

### Available Operations

* [~~Assert~~](#assert) - SAML ACS endpoint (social) :warning: **Deprecated**

## ~~Assert~~

The Assertion Consumer Service (ACS) endpoint for Simple Assertion Markup Language (SAML) social connectors.

SAML social connectors are deprecated. Use the SSO SAML connector instead.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

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

    res, err := s.AuthnSaml.Assert(ctx, "<id>", operations.AssertSamlRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `connectorID`                                                                        | *string*                                                                             | :heavy_check_mark:                                                                   | The unique identifier of the connector.                                              |
| `requestBody`                                                                        | [operations.AssertSamlRequestBody](../../models/operations/assertsamlrequestbody.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.AssertSamlResponse](../../models/operations/assertsamlresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |