# HealthCountType


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Green`                                                          | `*int64`                                                         | :heavy_minus_sign:                                               | Number of Worker Processes reporting a healthy (Green) status.   |
| `Red`                                                            | `*int64`                                                         | :heavy_minus_sign:                                               | Number of Worker Processes reporting a critical (Red) status.    |
| `Unknown`                                                        | `*int64`                                                         | :heavy_minus_sign:                                               | Number of Worker Processes reporting an unknown health status.   |
| `Yellow`                                                         | `*int64`                                                         | :heavy_minus_sign:                                               | Number of Worker Processes reporting a degraded (Yellow) status. |