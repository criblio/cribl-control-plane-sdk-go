# FunctionConfSchemaLookupInField


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `EventField`                                                                            | *string*                                                                                | :heavy_check_mark:                                                                      | Field name as it appears in events                                                      |
| `LookupField`                                                                           | **string*                                                                               | :heavy_minus_sign:                                                                      | Optional: The field name as it appears in the lookup file. Defaults to event field name |