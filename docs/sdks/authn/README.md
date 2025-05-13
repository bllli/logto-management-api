# Authn
(*Authn*)

## Overview

Authentication endpoints for third-party integrations and identity providers.

### Available Operations

* [AssertSingleSignOnSaml](#assertsinglesignonsaml) - SAML ACS endpoint (SSO)

## AssertSingleSignOnSaml

The Assertion Consumer Service (ACS) endpoint for Simple Assertion Markup Language (SAML) single sign-on (SSO) connectors.

This endpoint is used to complete the SAML SSO authentication flow. It receives the SAML assertion response from the identity provider (IdP) and redirects the user to complete the authentication flow.

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

    res, err := s.Authn.AssertSingleSignOnSaml(ctx, "<id>", operations.AssertSingleSignOnSamlRequestBody{
        SAMLResponse: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `connectorID`                                                                                                | *string*                                                                                                     | :heavy_check_mark:                                                                                           | The unique identifier of the connector.                                                                      |
| `requestBody`                                                                                                | [operations.AssertSingleSignOnSamlRequestBody](../../models/operations/assertsinglesignonsamlrequestbody.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.AssertSingleSignOnSamlResponse](../../models/operations/assertsinglesignonsamlresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |