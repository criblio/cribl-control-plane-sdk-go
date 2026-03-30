# HealthServerStatusError

Health status of the Leader or Worker Node.


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `Role`                                                                                     | [*components.Role](../../models/components/role.md)                                        | :heavy_minus_sign:                                                                         | Leader Node role: <code>primary</code> or <code>standby</code>.                            |
| `StartTime`                                                                                | `int64`                                                                                    | :heavy_check_mark:                                                                         | Timestamp (in Unix time) when the Cribl process started.                                   |
| `Status`                                                                                   | [components.HealthServerStatusStatus](../../models/components/healthserverstatusstatus.md) | :heavy_check_mark:                                                                         | Health state: <code>healthy</code>, <code>standby</code>, or <code>shutting down</code>.   |
| `HTTPMeta`                                                                                 | [components.HTTPMetadata](../../models/components/httpmetadata.md)                         | :heavy_check_mark:                                                                         | N/A                                                                                        |