# InputDatadogAgentSamplingRule


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Service`                                                        | `string`                                                         | :heavy_check_mark:                                               | Datadog service name                                             |
| `Environment`                                                    | `string`                                                         | :heavy_check_mark:                                               | Datadog environment name (example: prod, staging)                |
| `Rate`                                                           | `float64`                                                        | :heavy_check_mark:                                               | Sampling rate for this service/environment combination (0.0–1.0) |