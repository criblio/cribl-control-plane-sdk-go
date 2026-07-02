# NotificationConfigForSMTPTarget1

Simple Mail Transfer Protocol (SMTP) configuration for the Notification target.


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `Subject`                                                                 | `*string`                                                                 | :heavy_minus_sign:                                                        | Email subject                                                             |
| `Body`                                                                    | `*string`                                                                 | :heavy_minus_sign:                                                        | Email body                                                                |
| `EmailRecipient`                                                          | [*components.EmailRecipient1](../../models/components/emailrecipient1.md) | :heavy_minus_sign:                                                        | Email recipient settings for the Notification target.                     |