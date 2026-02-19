# PipelineFunctionPackConf


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `UnpackedFields`                                       | []*string*                                             | :heavy_check_mark:                                     | List of fields to keep, everything else will be packed |
| `Target`                                               | **string*                                              | :heavy_minus_sign:                                     | Name of the (packed) target field                      |