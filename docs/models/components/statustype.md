# StatusType

Runtime status: health, metrics, and optional persistent-queue info. Fields may be absent when data is unavailable.


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `Error`                                                                     | [*components.StatusError](../../models/components/statuserror.md)           | :heavy_minus_sign:                                                          | N/A                                                                         |
| `Health`                                                                    | [*components.HealthStringType](../../models/components/healthstringtype.md) | :heavy_minus_sign:                                                          | N/A                                                                         |
| `Metrics`                                                                   | map[string]`any`                                                            | :heavy_minus_sign:                                                          | Metrics data for the Source or Destination.                                 |
| `Pq`                                                                        | [*components.WorkerPQStatus](../../models/components/workerpqstatus.md)     | :heavy_minus_sign:                                                          | N/A                                                                         |
| `Timestamp`                                                                 | `*float64`                                                                  | :heavy_minus_sign:                                                          | Timestamp (in Unix time) when the status was last updated.                  |
| `UseStatusFromLB`                                                           | `*bool`                                                                     | :heavy_minus_sign:                                                          | Set to prefer status from the LB process, not from the worker process.      |