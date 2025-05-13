# OrganizationRoles
(*OrganizationRoles*)

## Overview

### Available Operations

* [Get](#get) - Get organization role
* [Update](#update) - Update organization role
* [Delete](#delete) - Delete organization role
* [List](#list) - Get organization roles
* [Create](#create) - Create an organization role
* [ListScopes](#listscopes) - Get organization role scopes
* [AssignScope](#assignscope) - Assign organization scopes to organization role
* [ReplaceScopes](#replacescopes) - Replace organization scopes for organization role
* [DeleteScope](#deletescope) - Remove organization scope
* [ListResourceScopes](#listresourcescopes) - Get organization role resource scopes
* [AssignResourceScopes](#assignresourcescopes) - Assign resource scopes to organization role
* [ReplaceResourceScopes](#replaceresourcescopes) - Replace resource scopes for organization role
* [DeleteResourceScope](#deleteresourcescope) - Remove resource scope

## Get

Get organization role details by ID.

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

    res, err := s.OrganizationRoles.Get(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization role.          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetOrganizationRoleResponse](../../models/operations/getorganizationroleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update organization role details by ID with the given data.

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

    res, err := s.OrganizationRoles.Update(ctx, "<id>", operations.UpdateOrganizationRoleRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `id`                                                                                                         | *string*                                                                                                     | :heavy_check_mark:                                                                                           | The unique identifier of the organization role.                                                              |
| `requestBody`                                                                                                | [operations.UpdateOrganizationRoleRequestBody](../../models/operations/updateorganizationrolerequestbody.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.UpdateOrganizationRoleResponse](../../models/operations/updateorganizationroleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete organization role by ID.

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

    res, err := s.OrganizationRoles.Delete(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization role.          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteOrganizationRoleResponse](../../models/operations/deleteorganizationroleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Get organization roles with pagination.

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

    res, err := s.OrganizationRoles.List(ctx, nil, nil, nil)
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
| `q`                                                      | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | Page number (starts from 1).                             |
| `pageSize`                                               | **int64*                                                 | :heavy_minus_sign:                                       | Entries per page.                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListOrganizationRolesResponse](../../models/operations/listorganizationrolesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a new organization role with the given data.

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

    res, err := s.OrganizationRoles.Create(ctx, operations.CreateOrganizationRoleRequest{
        Name: "<value>",
        OrganizationScopeIds: []string{
            "<value>",
            "<value>",
            "<value>",
        },
        ResourceScopeIds: []string{
            "<value>",
            "<value>",
            "<value>",
        },
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateOrganizationRoleRequest](../../models/operations/createorganizationrolerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateOrganizationRoleResponse](../../models/operations/createorganizationroleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListScopes

Get organization scopes that are assigned to the specified organization role with optional pagination.

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

    res, err := s.OrganizationRoles.ListScopes(ctx, "<id>", nil, nil)
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization role.          |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | Page number (starts from 1).                             |
| `pageSize`                                               | **int64*                                                 | :heavy_minus_sign:                                       | Entries per page.                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListOrganizationRoleScopesResponse](../../models/operations/listorganizationrolescopesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AssignScope

Assign organization scopes to the specified organization role

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

    res, err := s.OrganizationRoles.AssignScope(ctx, "<id>", operations.CreateOrganizationRoleScopeRequestBody{
        OrganizationScopeIds: []string{
            "<value>",
            "<value>",
            "<value>",
        },
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `id`                                                                                                                   | *string*                                                                                                               | :heavy_check_mark:                                                                                                     | The unique identifier of the organization role.                                                                        |
| `requestBody`                                                                                                          | [operations.CreateOrganizationRoleScopeRequestBody](../../models/operations/createorganizationrolescoperequestbody.md) | :heavy_check_mark:                                                                                                     | N/A                                                                                                                    |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.CreateOrganizationRoleScopeResponse](../../models/operations/createorganizationrolescoperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ReplaceScopes

Replace all organization scopes that are assigned to the specified organization role with the given organization scopes. This effectively removes all existing organization scope assignments and replaces them with the new ones.

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

    res, err := s.OrganizationRoles.ReplaceScopes(ctx, "<id>", operations.ReplaceOrganizationRoleScopesRequestBody{
        OrganizationScopeIds: []string{
            "<value>",
            "<value>",
            "<value>",
        },
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `id`                                                                                                                       | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The unique identifier of the organization role.                                                                            |
| `requestBody`                                                                                                              | [operations.ReplaceOrganizationRoleScopesRequestBody](../../models/operations/replaceorganizationrolescopesrequestbody.md) | :heavy_check_mark:                                                                                                         | N/A                                                                                                                        |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.ReplaceOrganizationRoleScopesResponse](../../models/operations/replaceorganizationrolescopesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteScope

Remove a organization scope assignment from the specified organization role.

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

    res, err := s.OrganizationRoles.DeleteScope(ctx, "<id>", "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization role.          |
| `organizationScopeID`                                    | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization scope.         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteOrganizationRoleScopeResponse](../../models/operations/deleteorganizationrolescoperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListResourceScopes

Get resource scopes that are assigned to the specified organization role with optional pagination.

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

    res, err := s.OrganizationRoles.ListResourceScopes(ctx, "<id>", nil, nil)
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization role.          |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | Page number (starts from 1).                             |
| `pageSize`                                               | **int64*                                                 | :heavy_minus_sign:                                       | Entries per page.                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListOrganizationRoleResourceScopesResponse](../../models/operations/listorganizationroleresourcescopesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AssignResourceScopes

Assign resource scopes to the specified organization role

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

    res, err := s.OrganizationRoles.AssignResourceScopes(ctx, "<id>", operations.CreateOrganizationRoleResourceScopeRequestBody{
        ScopeIds: []string{
            "<value>",
        },
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

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `id`                                                                                                                                   | *string*                                                                                                                               | :heavy_check_mark:                                                                                                                     | The unique identifier of the organization role.                                                                                        |
| `requestBody`                                                                                                                          | [operations.CreateOrganizationRoleResourceScopeRequestBody](../../models/operations/createorganizationroleresourcescoperequestbody.md) | :heavy_check_mark:                                                                                                                     | N/A                                                                                                                                    |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.CreateOrganizationRoleResourceScopeResponse](../../models/operations/createorganizationroleresourcescoperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ReplaceResourceScopes

Replace all resource scopes that are assigned to the specified organization role with the given resource scopes. This effectively removes all existing organization scope assignments and replaces them with the new ones.

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

    res, err := s.OrganizationRoles.ReplaceResourceScopes(ctx, "<id>", operations.ReplaceOrganizationRoleResourceScopesRequestBody{
        ScopeIds: []string{
            "<value>",
            "<value>",
            "<value>",
        },
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

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `id`                                                                                                                                       | *string*                                                                                                                                   | :heavy_check_mark:                                                                                                                         | The unique identifier of the organization role.                                                                                            |
| `requestBody`                                                                                                                              | [operations.ReplaceOrganizationRoleResourceScopesRequestBody](../../models/operations/replaceorganizationroleresourcescopesrequestbody.md) | :heavy_check_mark:                                                                                                                         | N/A                                                                                                                                        |
| `opts`                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                   | :heavy_minus_sign:                                                                                                                         | The options for this request.                                                                                                              |

### Response

**[*operations.ReplaceOrganizationRoleResourceScopesResponse](../../models/operations/replaceorganizationroleresourcescopesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteResourceScope

Remove a resource scope assignment from the specified organization role.

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

    res, err := s.OrganizationRoles.DeleteResourceScope(ctx, "<id>", "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the organization role.          |
| `scopeID`                                                | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the scope.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteOrganizationRoleResourceScopeResponse](../../models/operations/deleteorganizationroleresourcescoperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |