# Configs
(*Configs*)

## Overview

Endpoints for managing Logto global configurations for the tenant, such as admin console config and OIDC signing keys.

See [🔑 Signing keys](https://docs.logto.io/docs/recipes/signing-keys-rotation/) to learn more about signing keys and key rotation.

### Available Operations

* [GetAdminConsole](#getadminconsole) - Get admin console config
* [UpdateAdminConsole](#updateadminconsole) - Update admin console config
* [GetOidcKeys](#getoidckeys) - Get OIDC keys
* [DeleteOidcKey](#deleteoidckey) - Delete OIDC key
* [RotateOidcKeys](#rotateoidckeys) - Rotate OIDC keys
* [GetJwtCustomizer](#getjwtcustomizer) - Get JWT customizer
* [DeleteJwtCustomizer](#deletejwtcustomizer) - Delete JWT customizer
* [ListJwtCustomizers](#listjwtcustomizers) - Get all JWT customizers
* [TestJwtCustomizer](#testjwtcustomizer) - Test JWT customizer

## GetAdminConsole

Get the global configuration object for Logto Console.

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

    res, err := s.Configs.GetAdminConsole(ctx)
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

**[*operations.GetAdminConsoleConfigResponse](../../models/operations/getadminconsoleconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateAdminConsole

Update the global configuration object for Logto Console. This method performs a partial update.

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

    res, err := s.Configs.UpdateAdminConsole(ctx, operations.UpdateAdminConsoleConfigRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpdateAdminConsoleConfigRequest](../../models/operations/updateadminconsoleconfigrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UpdateAdminConsoleConfigResponse](../../models/operations/updateadminconsoleconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetOidcKeys

Get OIDC signing keys by key type. The actual key will be redacted from the result.

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

    res, err := s.Configs.GetOidcKeys(ctx, operations.GetOidcKeysKeyTypePrivateKeys)
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                            | Type                                                                                                                                                                                                                 | Required                                                                                                                                                                                                             | Description                                                                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                                  |
| `keyType`                                                                                                                                                                                                            | [operations.GetOidcKeysKeyType](../../models/operations/getoidckeyskeytype.md)                                                                                                                                       | :heavy_check_mark:                                                                                                                                                                                                   | Private keys are used to sign OIDC JWTs. Cookie keys are used to sign OIDC cookies. For clients, they do not need to know private keys to verify OIDC JWTs; they can use public keys from the JWKS endpoint instead. |
| `opts`                                                                                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                   | The options for this request.                                                                                                                                                                                        |

### Response

**[*operations.GetOidcKeysResponse](../../models/operations/getoidckeysresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteOidcKey

Delete an OIDC signing key by key type and key ID.

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

    res, err := s.Configs.DeleteOidcKey(ctx, operations.DeleteOidcKeyKeyTypePrivateKeys, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                            | Type                                                                                                                                                                                                                 | Required                                                                                                                                                                                                             | Description                                                                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                                  |
| `keyType`                                                                                                                                                                                                            | [operations.DeleteOidcKeyKeyType](../../models/operations/deleteoidckeykeytype.md)                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                   | Private keys are used to sign OIDC JWTs. Cookie keys are used to sign OIDC cookies. For clients, they do not need to know private keys to verify OIDC JWTs; they can use public keys from the JWKS endpoint instead. |
| `keyID`                                                                                                                                                                                                              | *string*                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                   | The unique identifier of the key.                                                                                                                                                                                    |
| `opts`                                                                                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                   | The options for this request.                                                                                                                                                                                        |

### Response

**[*operations.DeleteOidcKeyResponse](../../models/operations/deleteoidckeyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RotateOidcKeys

A new key will be generated and prepend to the list of keys.

Only two recent keys will be kept. The oldest key will be automatically removed if there are more than two keys.

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

    res, err := s.Configs.RotateOidcKeys(ctx, operations.RotateOidcKeysKeyTypePrivateKeys, operations.RotateOidcKeysRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseBodies != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                            | Type                                                                                                                                                                                                                 | Required                                                                                                                                                                                                             | Description                                                                                                                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                                   | The context to use for the request.                                                                                                                                                                                  |
| `keyType`                                                                                                                                                                                                            | [operations.RotateOidcKeysKeyType](../../models/operations/rotateoidckeyskeytype.md)                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                   | Private keys are used to sign OIDC JWTs. Cookie keys are used to sign OIDC cookies. For clients, they do not need to know private keys to verify OIDC JWTs; they can use public keys from the JWKS endpoint instead. |
| `requestBody`                                                                                                                                                                                                        | [operations.RotateOidcKeysRequestBody](../../models/operations/rotateoidckeysrequestbody.md)                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                                   | N/A                                                                                                                                                                                                                  |
| `opts`                                                                                                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                   | The options for this request.                                                                                                                                                                                        |

### Response

**[*operations.RotateOidcKeysResponse](../../models/operations/rotateoidckeysresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetJwtCustomizer

Get the JWT customizer for the given token type.

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

    res, err := s.Configs.GetJwtCustomizer(ctx, operations.GetJwtCustomizerTokenTypePathClientCredentials)
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `tokenTypePath`                                                                                      | [operations.GetJwtCustomizerTokenTypePath](../../models/operations/getjwtcustomizertokentypepath.md) | :heavy_check_mark:                                                                                   | The token type to get the JWT customizer for.                                                        |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetJwtCustomizerResponse](../../models/operations/getjwtcustomizerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteJwtCustomizer

Delete the JWT customizer for the given token type.

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

    res, err := s.Configs.DeleteJwtCustomizer(ctx, operations.DeleteJwtCustomizerTokenTypePathAccessToken)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `tokenTypePath`                                                                                            | [operations.DeleteJwtCustomizerTokenTypePath](../../models/operations/deletejwtcustomizertokentypepath.md) | :heavy_check_mark:                                                                                         | The token type path to delete the JWT customizer for.                                                      |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.DeleteJwtCustomizerResponse](../../models/operations/deletejwtcustomizerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListJwtCustomizers

Get all JWT customizers for the tenant.

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

    res, err := s.Configs.ListJwtCustomizers(ctx)
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListJwtCustomizersResponse](../../models/operations/listjwtcustomizersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## TestJwtCustomizer

Test the JWT customizer script with the given sample context and sample token payload.

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

    res, err := s.Configs.TestJwtCustomizer(ctx, operations.CreateTestJwtCustomizerRequestTestJwtCustomizerRequestBody1(
        operations.TestJwtCustomizerRequestBody1{
            TokenType: "<value>",
            Script: "<value>",
            Token: operations.Token1{},
            Context: operations.TestJwtCustomizerContext{
                User: operations.TestJwtCustomizerUser{},
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.TestJwtCustomizerRequest](../../models/operations/testjwtcustomizerrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.TestJwtCustomizerResponse](../../models/operations/testjwtcustomizerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |