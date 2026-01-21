# CollectorAzureBlob

AzureBlob collector configuration


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `Type`                                                                                 | [components.CollectorAzureBlobType](../../models/components/collectorazureblobtype.md) | :heavy_check_mark:                                                                     | Collector type                                                                         |
| `Conf`                                                                                 | [components.AzureBlobCollectorConf](../../models/components/azureblobcollectorconf.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `Destructive`                                                                          | **bool*                                                                                | :heavy_minus_sign:                                                                     | Delete any files collected (where applicable)                                          |
| `Encoding`                                                                             | **string*                                                                              | :heavy_minus_sign:                                                                     | Character encoding to use when parsing ingested data.                                  |