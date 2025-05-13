# SignInExperiences
(*SignInExperiences*)

## Overview

### Available Operations

* [CheckPassword](#checkpassword) - Check if a password meets the password policy

## CheckPassword

Check if a password meets the password policy in the sign-in experience settings.

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

    res, err := s.SignInExperiences.CheckPassword(ctx, operations.CheckPasswordWithDefaultSignInExperienceRequest{
        Password: "LZauJuz6UKDWvhS",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `request`                                                                                                                                | [operations.CheckPasswordWithDefaultSignInExperienceRequest](../../models/operations/checkpasswordwithdefaultsigninexperiencerequest.md) | :heavy_check_mark:                                                                                                                       | The request object to use for the request.                                                                                               |
| `opts`                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.CheckPasswordWithDefaultSignInExperienceResponse](../../models/operations/checkpasswordwithdefaultsigninexperienceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |