# PipelineFunctionUnrollConf


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `SrcExpr`                                                                             | *string*                                                                              | :heavy_check_mark:                                                                    | Field in which to find/calculate the array to unroll. Example: _raw, _raw.split(/\n/) |
| `DstField`                                                                            | *string*                                                                              | :heavy_check_mark:                                                                    | Field in destination event in which to place the unrolled value                       |