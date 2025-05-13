# PasswordIdentityVerifications
(*PasswordIdentityVerifications*)

## Overview

### Available Operations

* [Create](#create) - Create new password identity verification

## Create

Create a NewPasswordIdentity verification record for the new user registration use. The verification record includes a unique user identifier and a password that can be used to create a new user account.

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

    res, err := s.PasswordIdentityVerifications.Create(ctx, operations.CreateNewPasswordIdentityVerificationRequest{
        Identifier: operations.CreateNewPasswordIdentityVerificationIdentifier{
            Type: "<value>",
            Value: "<value>",
        },
        Password: "ACQ5nN2Mca5wl8R",
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

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `request`                                                                                                                          | [operations.CreateNewPasswordIdentityVerificationRequest](../../models/operations/createnewpasswordidentityverificationrequest.md) | :heavy_check_mark:                                                                                                                 | The request object to use for the request.                                                                                         |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.CreateNewPasswordIdentityVerificationResponse](../../models/operations/createnewpasswordidentityverificationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |