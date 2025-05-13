# ListResourceScopesRequest


## Fields

| Field                                  | Type                                   | Required                               | Description                            |
| -------------------------------------- | -------------------------------------- | -------------------------------------- | -------------------------------------- |
| `ResourceID`                           | *string*                               | :heavy_check_mark:                     | The unique identifier of the resource. |
| `Page`                                 | **int64*                               | :heavy_minus_sign:                     | Page number (starts from 1).           |
| `PageSize`                             | **int64*                               | :heavy_minus_sign:                     | Entries per page.                      |
| `SearchParams`                         | map[string]*string*                    | :heavy_minus_sign:                     | Search query parameters.               |