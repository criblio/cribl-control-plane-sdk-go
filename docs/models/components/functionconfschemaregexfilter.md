# FunctionConfSchemaRegexFilter


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Regex`                                                        | **string*                                                      | :heavy_minus_sign:                                             | Regex to test against                                          |
| `RegexList`                                                    | [][components.RegexList](../../models/components/regexlist.md) | :heavy_minus_sign:                                             | N/A                                                            |
| `Field`                                                        | **string*                                                      | :heavy_minus_sign:                                             | Name of the field to apply the regex on (defaults to _raw)     |