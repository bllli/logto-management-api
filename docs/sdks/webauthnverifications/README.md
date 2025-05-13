# WebAuthnVerifications
(*WebAuthnVerifications*)

## Overview

### Available Operations

* [CreateAuthentication](#createauthentication) - Create WebAuthn authentication verification

## CreateAuthentication

Create a new WebAuthn authentication verification record based on the user's existing WebAuthn credential. This verification record can be used to verify the user's WebAuthn credential.

### Example Usage

```go
package main

import(
	"context"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New()

    res, err := s.WebAuthnVerifications.CreateAuthentication(ctx)
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

**[*operations.CreateWebAuthnAuthenticationVerificationResponse](../../models/operations/createwebauthnauthenticationverificationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |