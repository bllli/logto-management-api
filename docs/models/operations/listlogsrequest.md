# ListLogsRequest


## Fields

| Field                          | Type                           | Required                       | Description                    |
| ------------------------------ | ------------------------------ | ------------------------------ | ------------------------------ |
| `UserID`                       | **string*                      | :heavy_minus_sign:             | Filter logs by user ID.        |
| `ApplicationID`                | **string*                      | :heavy_minus_sign:             | Filter logs by application ID. |
| `LogKey`                       | **string*                      | :heavy_minus_sign:             | Filter logs by log key.        |
| `Page`                         | **int64*                       | :heavy_minus_sign:             | Page number (starts from 1).   |
| `PageSize`                     | **int64*                       | :heavy_minus_sign:             | Entries per page.              |