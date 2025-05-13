# GetOidcKeysResponse


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `HTTPMeta`                                                                                 | [components.HTTPMetadata](../../models/components/httpmetadata.md)                         | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `ResponseBodies`                                                                           | [][operations.GetOidcKeysResponseBody](../../models/operations/getoidckeysresponsebody.md) | :heavy_minus_sign:                                                                         | An array of OIDC signing keys for the given key type.                                      |