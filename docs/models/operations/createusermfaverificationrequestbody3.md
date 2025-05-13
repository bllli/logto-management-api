# CreateUserMfaVerificationRequestBody3


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `Type`                                                                                | *string*                                                                              | :heavy_check_mark:                                                                    | The type of MFA verification to create.                                               |
| `Secret`                                                                              | **string*                                                                             | :heavy_minus_sign:                                                                    | The secret for the MFA verification, if not provided, a new secret will be generated. |