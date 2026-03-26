# JobStatus

Status of a job, including its current state and failure reason.


## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `Reason`                                                                                 | map[string]`any`                                                                         | :heavy_minus_sign:                                                                       | Reason the job entered its current <code>state</code>, typically populated upon failure. |
| `State`                                                                                  | [components.State](../../models/components/state.md)                                     | :heavy_check_mark:                                                                       | State of the Job                                                                         |