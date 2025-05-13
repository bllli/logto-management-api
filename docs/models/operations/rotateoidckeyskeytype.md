# RotateOidcKeysKeyType

Private keys are used to sign OIDC JWTs. Cookie keys are used to sign OIDC cookies. For clients, they do not need to know private keys to verify OIDC JWTs; they can use public keys from the JWKS endpoint instead.


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `RotateOidcKeysKeyTypePrivateKeys` | private-keys                       |
| `RotateOidcKeysKeyTypeCookieKeys`  | cookie-keys                        |