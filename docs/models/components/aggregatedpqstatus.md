# AggregatedPQStatus


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `Error`                                                                     | [*components.StatusError](../../models/components/statuserror.md)           | :heavy_minus_sign:                                                          | N/A                                                                         |
| `Health`                                                                    | [components.HealthStringType](../../models/components/healthstringtype.md)  | :heavy_check_mark:                                                          | N/A                                                                         |
| `HealthCounts`                                                              | [components.HealthCountType](../../models/components/healthcounttype.md)    | :heavy_check_mark:                                                          | N/A                                                                         |
| `Timestamp`                                                                 | `float64`                                                                   | :heavy_check_mark:                                                          | Timestamp (in Unix time) when the persistent queue status was last updated. |