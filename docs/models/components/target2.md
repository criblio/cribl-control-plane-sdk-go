# Target2


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Protocol`                                                    | [*components.Protocol2](../../models/components/protocol2.md) | :heavy_minus_sign:                                            | Protocol to use when collecting metrics                       |
| `Host`                                                        | *string*                                                      | :heavy_check_mark:                                            | Name of host from which to pull metrics.                      |
| `Port`                                                        | **float64*                                                    | :heavy_minus_sign:                                            | The port number in the metrics URL for discovered targets.    |
| `Path`                                                        | **string*                                                     | :heavy_minus_sign:                                            | Path to use when collecting metrics from discovered targets   |