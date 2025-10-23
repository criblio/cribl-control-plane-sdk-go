# GetHealthInfoResponse


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `HTTPMeta`                                                          | [components.HTTPMetadata](../../models/components/httpmetadata.md)  | :heavy_check_mark:                                                  | N/A                                                                 |
| `HealthStatus`                                                      | [*components.HealthStatus](../../models/components/healthstatus.md) | :heavy_minus_sign:                                                  | Healthy status                                                      |
| `Error`                                                             | [*components.Error](../../models/components/error.md)               | :heavy_minus_sign:                                                  | Unexpected error                                                    |