# ListApplicationsRequest


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `Types`                                               | [*operations.Types](../../models/operations/types.md) | :heavy_minus_sign:                                    | An array of application types to filter applications. |
| `ExcludeRoleID`                                       | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   |
| `ExcludeOrganizationID`                               | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   |
| `IsThirdParty`                                        | **string*                                             | :heavy_minus_sign:                                    | N/A                                                   |
| `Page`                                                | **int64*                                              | :heavy_minus_sign:                                    | Page number (starts from 1).                          |
| `PageSize`                                            | **int64*                                              | :heavy_minus_sign:                                    | Entries per page.                                     |
| `SearchParams`                                        | map[string]*string*                                   | :heavy_minus_sign:                                    | Search query parameters.                              |