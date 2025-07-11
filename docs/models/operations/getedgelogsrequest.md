# GetEdgeLogsRequest


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Mode`                                                                  | **string*                                                               | :heavy_minus_sign:                                                      | Discovery Mode (default is "auto")                                      |
| `Allow`                                                                 | **string*                                                               | :heavy_minus_sign:                                                      | query array[string] optional List of allowed-filename wildcard patterns |
| `Path`                                                                  | **string*                                                               | :heavy_minus_sign:                                                      | Search directory for "manual" mode                                      |
| `Depth`                                                                 | **float64*                                                              | :heavy_minus_sign:                                                      | Search depth for "manual" mode                                          |