# HBLeaderInfo

Connection parameters for the Leader Node, as reported in a Worker heartbeat.


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `Host`                                                          | `string`                                                        | :heavy_check_mark:                                              | Leader hostname or IP address.                                  |
| `Port`                                                          | `int64`                                                         | :heavy_check_mark:                                              | Leader TCP port.                                                |
| `Servername`                                                    | `*string`                                                       | :heavy_minus_sign:                                              | TLS server name (SNI) for the Leader connection.                |
| `TLS`                                                           | `*bool`                                                         | :heavy_minus_sign:                                              | If <code>true</code>, TLS is enabled for the Leader connection. |