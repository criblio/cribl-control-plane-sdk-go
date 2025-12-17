# Condition


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Key`                                                      | *string*                                                   | :heavy_check_mark:                                         | Event field name to match against                          |
| `Operator`                                                 | [components.Operator](../../models/components/operator.md) | :heavy_check_mark:                                         | Comparison operator                                        |
| `Value`                                                    | [components.Value](../../models/components/value.md)       | :heavy_check_mark:                                         | Value to compare against (string, number, boolean)         |