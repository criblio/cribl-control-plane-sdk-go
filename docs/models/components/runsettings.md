# RunSettings


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `MaxTaskReschedule`                                       | **float64*                                                | :heavy_minus_sign:                                        | Maximum number of times a task can be rescheduled.        |
| `Now`                                                     | [*time.Time](https://pkg.go.dev/time#Time)                | :heavy_minus_sign:                                        | N/A                                                       |
| `RescheduleDroppedTasks`                                  | **bool*                                                   | :heavy_minus_sign:                                        | Reschedule tasks that failed with non-fatal errors.       |
| `TaskHeartbeatPeriod`                                     | **float64*                                                | :heavy_minus_sign:                                        | N/A                                                       |
| `Type`                                                    | [*components.RunType](../../models/components/runtype.md) | :heavy_minus_sign:                                        | N/A                                                       |
| `AdditionalProperties`                                    | map[string]*any*                                          | :heavy_minus_sign:                                        | N/A                                                       |