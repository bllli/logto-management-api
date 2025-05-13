# Template


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `LanguageTag`                                                                    | *string*                                                                         | :heavy_check_mark:                                                               | The language tag of the email template, e.g., `en` or `fr`.                      |
| `TemplateType`                                                                   | [operations.TemplateTypeRequest](../../models/operations/templatetyperequest.md) | :heavy_check_mark:                                                               | The type of the email template, e.g. `SignIn` or `ForgotPassword`                |
| `Details`                                                                        | [operations.TemplateDetails](../../models/operations/templatedetails.md)         | :heavy_check_mark:                                                               | The details of the email template.                                               |