# ShutdownTypeSystemSettingsConf

Graceful shutdown configuration.


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `DrainTimeout`                                                                                | `int64`                                                                                       | :heavy_check_mark:                                                                            | Maximum time in milliseconds to wait for in-flight events to drain before forcing a shutdown. |