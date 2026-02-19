# CollectorCriblLake

CriblLake collector configuration


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `Type`                                                                                 | [components.CollectorCriblLakeType](../../models/components/collectorcribllaketype.md) | :heavy_check_mark:                                                                     | Collector type                                                                         |
| `Conf`                                                                                 | [components.CriblLakeCollectorConf](../../models/components/cribllakecollectorconf.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `Destructive`                                                                          | **bool*                                                                                | :heavy_minus_sign:                                                                     | Delete any files collected (where applicable)                                          |
| `Encoding`                                                                             | **string*                                                                              | :heavy_minus_sign:                                                                     | Character encoding to use when parsing ingested data.                                  |