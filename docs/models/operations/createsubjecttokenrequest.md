# CreateSubjectTokenRequest


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `UserID`                                                                                      | *string*                                                                                      | :heavy_check_mark:                                                                            | The ID of the user to impersonate.                                                            |
| `Context`                                                                                     | [*operations.CreateSubjectTokenContext](../../models/operations/createsubjecttokencontext.md) | :heavy_minus_sign:                                                                            | The additional context to be included in the token, this can be used in custom JWT.           |