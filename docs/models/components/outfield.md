# OutField


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `LookupField`                                                        | *string*                                                             | :heavy_check_mark:                                                   | The field name as it appears in the lookup file                      |
| `EventField`                                                         | **string*                                                            | :heavy_minus_sign:                                                   | Optional: Field name to add to event. Defaults to lookup field name. |
| `DefaultValue`                                                       | **string*                                                            | :heavy_minus_sign:                                                   | Optional: Value to assign if lookup entry is not found               |