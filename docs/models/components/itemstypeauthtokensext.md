# ItemsTypeAuthTokensExt


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Token`                                                                        | *string*                                                                       | :heavy_check_mark:                                                             | Shared secret to be provided by any client (Authorization: <token>)            |
| `Description`                                                                  | **string*                                                                      | :heavy_minus_sign:                                                             | N/A                                                                            |
| `Metadata`                                                                     | [][components.ItemsTypeMetadata](../../models/components/itemstypemetadata.md) | :heavy_minus_sign:                                                             | Fields to add to events referencing this token                                 |