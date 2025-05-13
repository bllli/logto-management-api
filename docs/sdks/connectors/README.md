# Connectors
(*Connectors*)

## Overview

Connectors are the bridge between Logto and other third-party vendors who provide short message service (SMS), email service, or user information on wildly accepted social media.

To learn more about connectors, please see [🪛 Configure connectors](https://docs.logto.io/docs/recipes/configure-connectors/).

### Available Operations

* [Create](#create) - Create connector
* [List](#list) - Get connectors
* [Get](#get) - Get connector
* [Update](#update) - Update connector
* [Delete](#delete) - Delete connector
* [Test](#test) - Test passwordless connector
* [CreateAuthorizationURI](#createauthorizationuri) - Get connector's authorization URI

## Create

Create a connector with the given data.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/bllli/logto-management-api/models/components"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"github.com/bllli/logto-management-api/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(components.Security{
            ClientID: logtomanagementapi.String(os.Getenv("LOGTOMANAGEMENTAPI_CLIENT_ID")),
            ClientSecret: logtomanagementapi.String(os.Getenv("LOGTOMANAGEMENTAPI_CLIENT_SECRET")),
        }),
    )

    res, err := s.Connectors.Create(ctx, operations.CreateConnectorRequest{
        ConnectorID: "<id>",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateConnectorRequest](../../models/operations/createconnectorrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateConnectorResponse](../../models/operations/createconnectorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Get all connectors in the current tenant.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/bllli/logto-management-api/models/components"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(components.Security{
            ClientID: logtomanagementapi.String(os.Getenv("LOGTOMANAGEMENTAPI_CLIENT_ID")),
            ClientSecret: logtomanagementapi.String(os.Getenv("LOGTOMANAGEMENTAPI_CLIENT_SECRET")),
        }),
    )

    res, err := s.Connectors.List(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `target`                                                 | **string*                                                | :heavy_minus_sign:                                       | Filter connectors by target.                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListConnectorsResponse](../../models/operations/listconnectorsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get connector data by ID

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/bllli/logto-management-api/models/components"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(components.Security{
            ClientID: logtomanagementapi.String(os.Getenv("LOGTOMANAGEMENTAPI_CLIENT_ID")),
            ClientSecret: logtomanagementapi.String(os.Getenv("LOGTOMANAGEMENTAPI_CLIENT_SECRET")),
        }),
    )

    res, err := s.Connectors.Get(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the connector.                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetConnectorResponse](../../models/operations/getconnectorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update connector by ID with the given data. This methods performs a partial update.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/bllli/logto-management-api/models/components"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"github.com/bllli/logto-management-api/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(components.Security{
            ClientID: logtomanagementapi.String(os.Getenv("LOGTOMANAGEMENTAPI_CLIENT_ID")),
            ClientSecret: logtomanagementapi.String(os.Getenv("LOGTOMANAGEMENTAPI_CLIENT_SECRET")),
        }),
    )

    res, err := s.Connectors.Update(ctx, "<id>", operations.UpdateConnectorRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `id`                                                                                           | *string*                                                                                       | :heavy_check_mark:                                                                             | The unique identifier of the connector.                                                        |
| `requestBody`                                                                                  | [operations.UpdateConnectorRequestBody](../../models/operations/updateconnectorrequestbody.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateConnectorResponse](../../models/operations/updateconnectorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete connector by ID.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/bllli/logto-management-api/models/components"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(components.Security{
            ClientID: logtomanagementapi.String(os.Getenv("LOGTOMANAGEMENTAPI_CLIENT_ID")),
            ClientSecret: logtomanagementapi.String(os.Getenv("LOGTOMANAGEMENTAPI_CLIENT_SECRET")),
        }),
    )

    res, err := s.Connectors.Delete(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the connector.                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteConnectorResponse](../../models/operations/deleteconnectorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Test

Test a passwordless (email or SMS) connector by sending a test message to the given phone number or email address.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/bllli/logto-management-api/models/components"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"github.com/bllli/logto-management-api/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(components.Security{
            ClientID: logtomanagementapi.String(os.Getenv("LOGTOMANAGEMENTAPI_CLIENT_ID")),
            ClientSecret: logtomanagementapi.String(os.Getenv("LOGTOMANAGEMENTAPI_CLIENT_SECRET")),
        }),
    )

    res, err := s.Connectors.Test(ctx, "<id>", operations.CreateConnectorTestRequestBody{
        Config: operations.CreateConnectorTestConfig{},
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `factoryID`                                                                                            | *string*                                                                                               | :heavy_check_mark:                                                                                     | The unique identifier of the factory.                                                                  |
| `requestBody`                                                                                          | [operations.CreateConnectorTestRequestBody](../../models/operations/createconnectortestrequestbody.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateConnectorTestResponse](../../models/operations/createconnectortestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateAuthorizationURI

Get authorization URI for specified connector by providing redirect URI and randomly generated state.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/bllli/logto-management-api/models/components"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"github.com/bllli/logto-management-api/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(components.Security{
            ClientID: logtomanagementapi.String(os.Getenv("LOGTOMANAGEMENTAPI_CLIENT_ID")),
            ClientSecret: logtomanagementapi.String(os.Getenv("LOGTOMANAGEMENTAPI_CLIENT_SECRET")),
        }),
    )

    res, err := s.Connectors.CreateAuthorizationURI(ctx, "<id>", operations.CreateConnectorAuthorizationURIRequestBody{
        State: "South Carolina",
        RedirectURI: "https://swift-behest.biz",
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

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `connectorID`                                                                                                                  | *string*                                                                                                                       | :heavy_check_mark:                                                                                                             | The unique identifier of the connector.                                                                                        |
| `requestBody`                                                                                                                  | [operations.CreateConnectorAuthorizationURIRequestBody](../../models/operations/createconnectorauthorizationurirequestbody.md) | :heavy_check_mark:                                                                                                             | N/A                                                                                                                            |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.CreateConnectorAuthorizationURIResponse](../../models/operations/createconnectorauthorizationuriresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |