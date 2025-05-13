# CreateWebAuthnAuthenticationVerificationResponseBody

WebAuthn authentication successfully initiated.


## Fields

| Field                                                                                                           | Type                                                                                                            | Required                                                                                                        | Description                                                                                                     |
| --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| `VerificationID`                                                                                                | *string*                                                                                                        | :heavy_check_mark:                                                                                              | The unique ID for the WebAuthn authentication record, required to verify the WebAuthn authentication challenge. |
| `AuthenticationOptions`                                                                                         | [operations.AuthenticationOptions](../../models/operations/authenticationoptions.md)                            | :heavy_check_mark:                                                                                              | Options for the user to authenticate with their WebAuthn credential.                                            |