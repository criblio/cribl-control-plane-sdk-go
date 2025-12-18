# FunctionConfSchemaGeoipAdditionalField


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ExtraInField`                                               | *string*                                                     | :heavy_check_mark:                                           | Field name in which to find an IP to look up. Can be nested. |
| `ExtraOutField`                                              | *string*                                                     | :heavy_check_mark:                                           | Field name in which to store the GeoIP lookup results        |