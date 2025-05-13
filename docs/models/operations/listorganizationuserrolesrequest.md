# ListOrganizationUserRolesRequest


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `ID`                                       | *string*                                   | :heavy_check_mark:                         | The unique identifier of the organization. |
| `UserID`                                   | *string*                                   | :heavy_check_mark:                         | The unique identifier of the user.         |
| `Page`                                     | **int64*                                   | :heavy_minus_sign:                         | Page number (starts from 1).               |
| `PageSize`                                 | **int64*                                   | :heavy_minus_sign:                         | Entries per page.                          |