# GetSignInExpSentinelPolicy

Custom sentinel policy settings. Use this field to customize the user lockout policy. The default value is 100 failed attempts within one hour. The user will be locked out for 60 minutes after exceeding the limit.


## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `MaxAttempts`      | **float64*         | :heavy_minus_sign: | N/A                |
| `LockoutDuration`  | **float64*         | :heavy_minus_sign: | N/A                |