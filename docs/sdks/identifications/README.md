# Identifications
(*Identifications*)

## Overview

### Available Operations

* [IdentifyUser](#identifyuser) - Identify user for the current interaction

## IdentifyUser

This API identifies the user based on the verificationId within the current experience interaction: <br/>- `SignIn` and `ForgotPassword` interactions: Verifies the user's identity using the provided `verificationId`. <br/>- `Register` interaction: Creates a new user account using the profile data from the current interaction. If a verificationId is provided, the profile data will first be updated with the verification record before creating the account. If not, the account is created directly from the stored profile data.

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

    res, err := s.Identifications.IdentifyUser(ctx, operations.IdentifyUserRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.IdentifyUserRequest](../../models/operations/identifyuserrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.IdentifyUserResponse](../../models/operations/identifyuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |