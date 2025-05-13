# GetSamlApplicationCallbackRequest


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `ID`                                           | *string*                                       | :heavy_check_mark:                             | The unique identifier of the saml application. |
| `Code`                                         | **string*                                      | :heavy_minus_sign:                             | The authorization code from OIDC callback.     |
| `State`                                        | **string*                                      | :heavy_minus_sign:                             | The state parameter from OIDC callback.        |
| `RedirectURI`                                  | **string*                                      | :heavy_minus_sign:                             | The redirect URI for the callback.             |
| `Error`                                        | **string*                                      | :heavy_minus_sign:                             | N/A                                            |
| `ErrorDescription`                             | **string*                                      | :heavy_minus_sign:                             | N/A                                            |