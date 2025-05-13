# AssignOrganizationRolesToUsersRequestBody


## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `UserIds`                                                                                   | []*string*                                                                                  | :heavy_check_mark:                                                                          | An array of user IDs to assign roles.                                                       |
| `OrganizationRoleIds`                                                                       | []*string*                                                                                  | :heavy_check_mark:                                                                          | An array of organization role IDs to assign. User existed roles assignment will be ignored. |