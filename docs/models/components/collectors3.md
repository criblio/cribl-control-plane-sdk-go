# CollectorS3

S3 collector configuration


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Type`                                                                   | [components.CollectorS3Type](../../models/components/collectors3type.md) | :heavy_check_mark:                                                       | Collector type                                                           |
| `Conf`                                                                   | [components.S3CollectorConf](../../models/components/s3collectorconf.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `Destructive`                                                            | **bool*                                                                  | :heavy_minus_sign:                                                       | Delete any files collected (where applicable)                            |
| `Encoding`                                                               | **string*                                                                | :heavy_minus_sign:                                                       | Character encoding to use when parsing ingested data.                    |