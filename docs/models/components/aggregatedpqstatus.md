# AggregatedPQStatus


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `Error`                                                                     | [*components.StatusError](../../models/components/statuserror.md)           | :heavy_minus_sign:                                                          | N/A                                                                         |
| `Health`                                                                    | [components.Health](../../models/components/health.md)                      | :heavy_check_mark:                                                          | Health status of the persistent queue.                                      |
| `HealthCounts`                                                              | [components.HealthCountType](../../models/components/healthcounttype.md)    | :heavy_check_mark:                                                          | N/A                                                                         |
| `Timestamp`                                                                 | `int64`                                                                     | :heavy_check_mark:                                                          | Timestamp (in Unix time) when the persistent queue status was last updated. |