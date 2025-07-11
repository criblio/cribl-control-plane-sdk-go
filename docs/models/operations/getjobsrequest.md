# GetJobsRequest


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `Offset`                                                  | **int64*                                                  | :heavy_minus_sign:                                        | Pagination offset                                         |
| `Limit`                                                   | **int64*                                                  | :heavy_minus_sign:                                        | Maximum number of items to return                         |
| `RunType`                                                 | [*operations.RunType](../../models/operations/runtype.md) | :heavy_minus_sign:                                        | Filter by job run type                                    |
| `State`                                                   | **string*                                                 | :heavy_minus_sign:                                        | Filter by current job state, e.g. "running"               |
| `ID`                                                      | **string*                                                 | :heavy_minus_sign:                                        | Filter by job id, e.g. "id=1608713335.3&id=1608713326.1"  |
| `CollectorID`                                             | **string*                                                 | :heavy_minus_sign:                                        | Filter by collector id, e.g. "collectorId=Prometheus-in"  |
| `GroupID`                                                 | **string*                                                 | :heavy_minus_sign:                                        | Filter by worker group id, e.g. "defaultHybrid"           |