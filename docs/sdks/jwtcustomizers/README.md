# JwtCustomizers
(*JwtCustomizers*)

## Overview

### Available Operations

* [Upsert](#upsert) - Create or update JWT customizer
* [Update](#update) - Update JWT customizer

## Upsert

Create or update a JWT customizer for the given token type.

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

    res, err := s.JwtCustomizers.Upsert(ctx, operations.UpsertJwtCustomizerTokenTypePathClientCredentials, operations.UpsertJwtCustomizerRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredApplicationJSONOneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                | Example                                                                                                    |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |                                                                                                            |
| `tokenTypePath`                                                                                            | [operations.UpsertJwtCustomizerTokenTypePath](../../models/operations/upsertjwtcustomizertokentypepath.md) | :heavy_check_mark:                                                                                         | The token type to create a JWT customizer for.                                                             |                                                                                                            |
| `requestBody`                                                                                              | [operations.UpsertJwtCustomizerRequestBody](../../models/operations/upsertjwtcustomizerrequestbody.md)     | :heavy_check_mark:                                                                                         | N/A                                                                                                        | {}                                                                                                         |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |                                                                                                            |

### Response

**[*operations.UpsertJwtCustomizerResponse](../../models/operations/upsertjwtcustomizerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the JWT customizer for the given token type.

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

    res, err := s.JwtCustomizers.Update(ctx, operations.UpdateJwtCustomizerTokenTypePathClientCredentials, operations.UpdateJwtCustomizerRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                | Example                                                                                                    |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |                                                                                                            |
| `tokenTypePath`                                                                                            | [operations.UpdateJwtCustomizerTokenTypePath](../../models/operations/updatejwtcustomizertokentypepath.md) | :heavy_check_mark:                                                                                         | The token type to update a JWT customizer for.                                                             |                                                                                                            |
| `requestBody`                                                                                              | [operations.UpdateJwtCustomizerRequestBody](../../models/operations/updatejwtcustomizerrequestbody.md)     | :heavy_check_mark:                                                                                         | N/A                                                                                                        | {}                                                                                                         |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |                                                                                                            |

### Response

**[*operations.UpdateJwtCustomizerResponse](../../models/operations/updatejwtcustomizerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |