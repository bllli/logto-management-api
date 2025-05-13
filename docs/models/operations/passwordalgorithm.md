# PasswordAlgorithm

The hash algorithm used for the password. It should be one of the supported algorithms: argon2, md5, sha1, sha256. Should the encryption algorithm differ from argon2, it will automatically be upgraded to argon2 upon the user's next sign-in.


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `PasswordAlgorithmArgon2i`  | Argon2i                     |
| `PasswordAlgorithmArgon2id` | Argon2id                    |
| `PasswordAlgorithmArgon2d`  | Argon2d                     |
| `PasswordAlgorithmSha1`     | SHA1                        |
| `PasswordAlgorithmSha256`   | SHA256                      |
| `PasswordAlgorithmMd5`      | MD5                         |
| `PasswordAlgorithmBcrypt`   | Bcrypt                      |
| `PasswordAlgorithmLegacy`   | Legacy                      |