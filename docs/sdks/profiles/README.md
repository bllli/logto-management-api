# Profiles
(*Profiles*)

## Overview

### Available Operations

* [Add](#add) - Add user profile

## Add

Adds user profile data to the current experience interaction. <br/>- For `Register`: The profile data provided before the identification request will be used to create a new user account. <br/>- For `SignIn` and `Register`: The profile data provided after the user is identified will be used to update the user's profile when the interaction is submitted. <br/>- `ForgotPassword`: Not supported.

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

    res, err := s.Profiles.Add(ctx, operations.CreateAddUserProfileRequestAddUserProfileRequestBody3(
        operations.AddUserProfileRequestBody3{
            Type: "<value>",
            VerificationID: "<id>",
        },
    ))
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
| `request`                                                                            | [operations.AddUserProfileRequest](../../models/operations/adduserprofilerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.AddUserProfileResponse](../../models/operations/adduserprofileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |