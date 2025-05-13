# CreateNewPasswordIdentityVerificationIdentifier

The unique user identifier.  <br/> Currently, only `username` is accepted. For `email` or `phone` registration, a `CodeVerification` record must be created and used to verify the user's email or phone number identifier.


## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Type`             | *string*           | :heavy_check_mark: | N/A                |
| `Value`            | *string*           | :heavy_check_mark: | N/A                |