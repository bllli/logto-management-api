# Users
(*Users*)

## Overview

Endpoints for user management. Including creating, updating, deleting, and querying users with flexible filters. In addition to the endpoints, see [🧑‍🚀 Manage users](https://docs.logto.io/docs/recipes/manage-users/) for more insights.

### Available Operations

* [Get](#get) - Get user
* [Update](#update) - Update user
* [Delete](#delete) - Delete user
* [GetCustomData](#getcustomdata) - Get user custom data
* [UpdateCustomData](#updatecustomdata) - Update user custom data
* [UpdateProfile](#updateprofile) - Update user profile
* [Create](#create) - Create user
* [List](#list) - Get users
* [UpdatePassword](#updatepassword) - Update user password
* [VerifyPassword](#verifypassword) - Verify user password
* [HasPassword](#haspassword) - Check if user has password
* [UpdateIsSuspended](#updateissuspended) - Update user suspension status
* [GetRoles](#getroles) - Get roles for user
* [AssignRoles](#assignroles) - Assign roles to user
* [ReplaceRoles](#replaceroles) - Update roles for user
* [DeleteRole](#deleterole) - Remove role from user
* [UpdateIdentity](#updateidentity) - Update social identity of user
* [DeleteIdentity](#deleteidentity) - Delete social identity from user
* [CreateIdentity](#createidentity) - Link social identity to user
* [ListOrganizations](#listorganizations) - Get organizations for a user
* [ListMfaVerifications](#listmfaverifications) - Get user's MFA verifications
* [CreateMfaVerification](#createmfaverification) - Create an MFA verification for a user
* [DeleteMfaVerification](#deletemfaverification) - Delete an MFA verification for a user
* [ListPersonalAccessTokens](#listpersonalaccesstokens) - Get personal access tokens
* [CreatePersonalAccessToken](#createpersonalaccesstoken) - Add personal access token
* [DeletePersonalAccessToken](#deletepersonalaccesstoken) - Delete personal access token
* [UpdatePersonalAccessToken](#updatepersonalaccesstoken) - Update personal access token

## Get

Get user data for the given ID.

### Example Usage

```go
package main

import(
	"context"
	"os"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
    )

    res, err := s.Users.Get(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                  | Type                                                                                                                                                                                       | Required                                                                                                                                                                                   | Description                                                                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                         | The context to use for the request.                                                                                                                                                        |
| `userID`                                                                                                                                                                                   | *string*                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                         | The unique identifier of the user.                                                                                                                                                         |
| `includeSsoIdentities`                                                                                                                                                                     | **string*                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                         | If it's provided with a truthy value (`true`, `1`, `yes`), each user in the response will include a `ssoIdentities` property containing a list of SSO identities associated with the user. |
| `opts`                                                                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                                                                   | :heavy_minus_sign:                                                                                                                                                                         | The options for this request.                                                                                                                                                              |

### Response

**[*operations.GetUserResponse](../../models/operations/getuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update user data for the given ID. This method performs a partial update.

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

    res, err := s.Users.Update(ctx, "<id>", operations.UpdateUserRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `userID`                                                                             | *string*                                                                             | :heavy_check_mark:                                                                   | The unique identifier of the user.                                                   |
| `requestBody`                                                                        | [operations.UpdateUserRequestBody](../../models/operations/updateuserrequestbody.md) | :heavy_check_mark:                                                                   | N/A                                                                                  |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdateUserResponse](../../models/operations/updateuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete user with the given ID. Note all associated data will be deleted cascadingly.

### Example Usage

```go
package main

import(
	"context"
	"os"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
    )

    res, err := s.Users.Delete(ctx, "<id>")
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
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the user.                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteUserResponse](../../models/operations/deleteuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetCustomData

Get custom data for the given user ID.

### Example Usage

```go
package main

import(
	"context"
	"os"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
    )

    res, err := s.Users.GetCustomData(ctx, "<id>")
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
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the user.                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListUserCustomDataResponse](../../models/operations/listusercustomdataresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateCustomData

Update custom data for the given user ID. This method performs a partial update of the custom data object.

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

    res, err := s.Users.UpdateCustomData(ctx, "<id>", operations.UpdateUserCustomDataRequestBody{
        CustomData: operations.UpdateUserCustomDataCustomData{},
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `userID`                                                                                                 | *string*                                                                                                 | :heavy_check_mark:                                                                                       | The unique identifier of the user.                                                                       |
| `requestBody`                                                                                            | [operations.UpdateUserCustomDataRequestBody](../../models/operations/updateusercustomdatarequestbody.md) | :heavy_check_mark:                                                                                       | N/A                                                                                                      |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UpdateUserCustomDataResponse](../../models/operations/updateusercustomdataresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateProfile

Update profile for the given user ID. This method performs a partial update of the profile object.

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

    res, err := s.Users.UpdateProfile(ctx, "<id>", operations.UpdateUserProfileRequestBody{
        Profile: operations.UpdateUserProfileProfile{},
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `userID`                                                                                           | *string*                                                                                           | :heavy_check_mark:                                                                                 | The unique identifier of the user.                                                                 |
| `requestBody`                                                                                      | [operations.UpdateUserProfileRequestBody](../../models/operations/updateuserprofilerequestbody.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateUserProfileResponseResponse](../../models/operations/updateuserprofileresponseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a new user with the given data.

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

    res, err := s.Users.Create(ctx, operations.CreateUserRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.CreateUserRequest](../../models/operations/createuserrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.CreateUserResponse](../../models/operations/createuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Get users with filters and pagination.

Logto provides a very flexible way to query users. You can filter users by almost any fields with multiple modes. To learn more about the query syntax, please refer to [Advanced user search](https://docs.logto.io/docs/recipes/manage-users/advanced-user-search/).

### Example Usage

```go
package main

import(
	"context"
	"os"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
    )

    res, err := s.Users.List(ctx, nil, nil, nil)
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
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | Page number (starts from 1).                             |
| `pageSize`                                               | **int64*                                                 | :heavy_minus_sign:                                       | Entries per page.                                        |
| `searchParams`                                           | map[string]*string*                                      | :heavy_minus_sign:                                       | Search query parameters.                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListUsersResponse](../../models/operations/listusersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdatePassword

Update user password for the given ID.

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

    res, err := s.Users.UpdatePassword(ctx, "<id>", operations.UpdateUserPasswordRequestBody{
        Password: "r4nOwni1EIPkrj4",
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
| `userID`                                                                                             | *string*                                                                                             | :heavy_check_mark:                                                                                   | The unique identifier of the user.                                                                   |
| `requestBody`                                                                                        | [operations.UpdateUserPasswordRequestBody](../../models/operations/updateuserpasswordrequestbody.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.UpdateUserPasswordResponse](../../models/operations/updateuserpasswordresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## VerifyPassword

Test if the given password matches the user's password.

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

    res, err := s.Users.VerifyPassword(ctx, "<id>", operations.VerifyUserPasswordRequestBody{
        Password: "7xNXIJvZ8m6XrIn",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `userID`                                                                                             | *string*                                                                                             | :heavy_check_mark:                                                                                   | The unique identifier of the user.                                                                   |
| `requestBody`                                                                                        | [operations.VerifyUserPasswordRequestBody](../../models/operations/verifyuserpasswordrequestbody.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.VerifyUserPasswordResponse](../../models/operations/verifyuserpasswordresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## HasPassword

Check if the user with the given ID has a password set.

### Example Usage

```go
package main

import(
	"context"
	"os"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
    )

    res, err := s.Users.HasPassword(ctx, "<id>")
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
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the user.                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetUserHasPasswordResponse](../../models/operations/getuserhaspasswordresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateIsSuspended

Update user suspension status for the given ID.

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

    res, err := s.Users.UpdateIsSuspended(ctx, "<id>", operations.UpdateUserIsSuspendedRequestBody{
        IsSuspended: true,
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `userID`                                                                                                   | *string*                                                                                                   | :heavy_check_mark:                                                                                         | The unique identifier of the user.                                                                         |
| `requestBody`                                                                                              | [operations.UpdateUserIsSuspendedRequestBody](../../models/operations/updateuserissuspendedrequestbody.md) | :heavy_check_mark:                                                                                         | N/A                                                                                                        |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpdateUserIsSuspendedResponse](../../models/operations/updateuserissuspendedresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetRoles

Get API resource roles assigned to the user with pagination.

### Example Usage

```go
package main

import(
	"context"
	"os"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
    )

    res, err := s.Users.GetRoles(ctx, "<id>", nil, nil, nil)
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
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the user.                       |
| `page`                                                   | **int64*                                                 | :heavy_minus_sign:                                       | Page number (starts from 1).                             |
| `pageSize`                                               | **int64*                                                 | :heavy_minus_sign:                                       | Entries per page.                                        |
| `searchParams`                                           | map[string]*string*                                      | :heavy_minus_sign:                                       | Search query parameters.                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListUserRolesResponse](../../models/operations/listuserrolesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AssignRoles

Assign API resource roles to the user. The roles will be added to the existing roles.

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

    res, err := s.Users.AssignRoles(ctx, "<id>", operations.AssignUserRolesRequestBody{
        RoleIds: []string{
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `userID`                                                                                       | *string*                                                                                       | :heavy_check_mark:                                                                             | The unique identifier of the user.                                                             |
| `requestBody`                                                                                  | [operations.AssignUserRolesRequestBody](../../models/operations/assignuserrolesrequestbody.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.AssignUserRolesResponse](../../models/operations/assignuserrolesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ReplaceRoles

Update API resource roles assigned to the user. This will replace the existing roles.

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

    res, err := s.Users.ReplaceRoles(ctx, "<id>", operations.ReplaceUserRolesRequestBody{
        RoleIds: []string{},
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `userID`                                                                                         | *string*                                                                                         | :heavy_check_mark:                                                                               | The unique identifier of the user.                                                               |
| `requestBody`                                                                                    | [operations.ReplaceUserRolesRequestBody](../../models/operations/replaceuserrolesrequestbody.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ReplaceUserRolesResponse](../../models/operations/replaceuserrolesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteRole

Remove an API resource role from the user.

### Example Usage

```go
package main

import(
	"context"
	"os"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
    )

    res, err := s.Users.DeleteRole(ctx, "<id>", "<id>")
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
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the user.                       |
| `roleID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the role.                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteUserRoleResponse](../../models/operations/deleteuserroleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateIdentity

Directly update a social identity of the user.

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

    res, err := s.Users.UpdateIdentity(ctx, "<id>", "<value>", operations.ReplaceUserIdentityRequestBody{
        UserID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `userID`                                                                                               | *string*                                                                                               | :heavy_check_mark:                                                                                     | The unique identifier of the user.                                                                     |
| `target`                                                                                               | *string*                                                                                               | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `requestBody`                                                                                          | [operations.ReplaceUserIdentityRequestBody](../../models/operations/replaceuseridentityrequestbody.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.ReplaceUserIdentityResponse](../../models/operations/replaceuseridentityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteIdentity

Delete a social identity from the user.

### Example Usage

```go
package main

import(
	"context"
	"os"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
    )

    res, err := s.Users.DeleteIdentity(ctx, "<id>", "<value>")
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
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the user.                       |
| `target`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteUserIdentityResponse](../../models/operations/deleteuseridentityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateIdentity

Link authenticated user identity from a social platform to a Logto user.

The usage of this API is usually coupled with `POST /connectors/:connectorId/authorization-uri`. With the help of these pair of APIs, you can implement a user profile page with the link social account feature in your application.

Note: Currently due to technical limitations, this API does not support the following connectors that rely on Logto interaction session: `@logto/connector-apple`, `@logto/connector-saml`, `@logto/connector-oidc` and `@logto/connector-oauth`.

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

    res, err := s.Users.CreateIdentity(ctx, "<id>", operations.CreateUserIdentityRequestBody{
        ConnectorID: "<id>",
        ConnectorData: map[string]any{
            "key": map[string]any{

            },
            "key1": map[string]any{

            },
            "key2": map[string]any{

            },
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
| `userID`                                                                                             | *string*                                                                                             | :heavy_check_mark:                                                                                   | The unique identifier of the user.                                                                   |
| `requestBody`                                                                                        | [operations.CreateUserIdentityRequestBody](../../models/operations/createuseridentityrequestbody.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateUserIdentityResponse](../../models/operations/createuseridentityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListOrganizations

Get all organizations that the user is a member of. In each organization object, the user's roles in that organization are included in the `organizationRoles` array.

### Example Usage

```go
package main

import(
	"context"
	"os"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
    )

    res, err := s.Users.ListOrganizations(ctx, "<id>")
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
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the user.                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListUserOrganizationsResponse](../../models/operations/listuserorganizationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListMfaVerifications

Get a user's existing MFA verifications for a given user ID.

### Example Usage

```go
package main

import(
	"context"
	"os"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
    )

    res, err := s.Users.ListMfaVerifications(ctx, "<id>")
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
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the user.                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListUserMfaVerificationsResponse](../../models/operations/listusermfaverificationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateMfaVerification

Create a new MFA verification for a given user ID.

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

    res, err := s.Users.CreateMfaVerification(ctx, "<id>", operations.CreateCreateUserMfaVerificationRequestBodyCreateUserMfaVerificationRequestBody1(
        operations.CreateUserMfaVerificationRequestBody1{
            Type: "<value>",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `userID`                                                                                                           | *string*                                                                                                           | :heavy_check_mark:                                                                                                 | The unique identifier of the user.                                                                                 |
| `requestBody`                                                                                                      | [operations.CreateUserMfaVerificationRequestBody](../../models/operations/createusermfaverificationrequestbody.md) | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.CreateUserMfaVerificationResponse](../../models/operations/createusermfaverificationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteMfaVerification

Delete an MFA verification for the user with the given verification ID. The verification ID must be associated with the given user ID.

### Example Usage

```go
package main

import(
	"context"
	"os"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
    )

    res, err := s.Users.DeleteMfaVerification(ctx, "<id>", "<id>")
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
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the user.                       |
| `verificationID`                                         | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the verification.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteUserMfaVerificationResponse](../../models/operations/deleteusermfaverificationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListPersonalAccessTokens

Get all personal access tokens for the user.

### Example Usage

```go
package main

import(
	"context"
	"os"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
    )

    res, err := s.Users.ListPersonalAccessTokens(ctx, "<id>")
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
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the user.                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListUserPersonalAccessTokensResponse](../../models/operations/listuserpersonalaccesstokensresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreatePersonalAccessToken

Add a new personal access token for the user.

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

    res, err := s.Users.CreatePersonalAccessToken(ctx, "<id>", operations.CreateUserPersonalAccessTokenRequestBody{
        Name: "<value>",
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `userID`                                                                                                                   | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The unique identifier of the user.                                                                                         |
| `requestBody`                                                                                                              | [operations.CreateUserPersonalAccessTokenRequestBody](../../models/operations/createuserpersonalaccesstokenrequestbody.md) | :heavy_check_mark:                                                                                                         | N/A                                                                                                                        |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.CreateUserPersonalAccessTokenResponse](../../models/operations/createuserpersonalaccesstokenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeletePersonalAccessToken

Delete a token for the user by name.

### Example Usage

```go
package main

import(
	"context"
	"os"
	logtomanagementapi "github.com/bllli/logto-management-api"
	"log"
)

func main() {
    ctx := context.Background()

    s := logtomanagementapi.New(
        logtomanagementapi.WithSecurity(os.Getenv("LOGTOMANAGEMENTAPI_BEARER_AUTH")),
    )

    res, err := s.Users.DeletePersonalAccessToken(ctx, "<id>", "<value>")
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
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The unique identifier of the user.                       |
| `name`                                                   | *string*                                                 | :heavy_check_mark:                                       | The name of the token.                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteUserPersonalAccessTokenResponse](../../models/operations/deleteuserpersonalaccesstokenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdatePersonalAccessToken

Update a token for the user by name.

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

    res, err := s.Users.UpdatePersonalAccessToken(ctx, "<id>", "<value>", operations.UpdateUserPersonalAccessTokenRequestBody{
        Name: "<value>",
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `userID`                                                                                                                   | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The unique identifier of the user.                                                                                         |
| `name`                                                                                                                     | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | The name of the token.                                                                                                     |
| `requestBody`                                                                                                              | [operations.UpdateUserPersonalAccessTokenRequestBody](../../models/operations/updateuserpersonalaccesstokenrequestbody.md) | :heavy_check_mark:                                                                                                         | N/A                                                                                                                        |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.UpdateUserPersonalAccessTokenResponse](../../models/operations/updateuserpersonalaccesstokenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |