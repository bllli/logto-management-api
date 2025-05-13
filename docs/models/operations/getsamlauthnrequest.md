# GetSamlAuthnRequest


## Fields

| Field                           | Type                            | Required                        | Description                     |
| ------------------------------- | ------------------------------- | ------------------------------- | ------------------------------- |
| `ID`                            | *string*                        | :heavy_check_mark:              | The ID of the SAML application. |
| `SAMLRequest`                   | *string*                        | :heavy_check_mark:              | The SAML request message.       |
| `Signature`                     | **string*                       | :heavy_minus_sign:              | The signature of the request.   |
| `SigAlg`                        | **string*                       | :heavy_minus_sign:              | The signature algorithm.        |
| `RelayState`                    | **string*                       | :heavy_minus_sign:              | The relay state parameter.      |