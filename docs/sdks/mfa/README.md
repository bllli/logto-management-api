# Mfa
(*Mfa*)

## Overview

### Available Operations

* [SkipBinding](#skipbinding) - Skip MFA binding flow

## SkipBinding

Skip MFA verification binding flow. If the MFA is enabled in the sign-in experience settings and marked as `UserControlled`, the user can skip the MFA verification binding flow by calling this API.

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

    res, err := s.Mfa.SkipBinding(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
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

**[*operations.SkipMfaBindingFlowResponse](../../models/operations/skipmfabindingflowresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |