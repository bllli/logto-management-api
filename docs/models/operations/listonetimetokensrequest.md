# ListOneTimeTokensRequest


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `Email`                                                                     | **string*                                                                   | :heavy_minus_sign:                                                          | Filter one-time tokens by email address.                                    |
| `Status`                                                                    | [*operations.QueryParamStatus](../../models/operations/queryparamstatus.md) | :heavy_minus_sign:                                                          | Filter one-time tokens by status.                                           |
| `Page`                                                                      | **int64*                                                                    | :heavy_minus_sign:                                                          | Page number (starts from 1).                                                |
| `PageSize`                                                                  | **int64*                                                                    | :heavy_minus_sign:                                                          | Entries per page.                                                           |