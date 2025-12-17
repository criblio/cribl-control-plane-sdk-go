# FunctionConfSchemaPivot


## Fields

| Field                                         | Type                                          | Required                                      | Description                                   |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| `LabelField`                                  | **string*                                     | :heavy_minus_sign:                            | Fields to be used for the left-most column.   |
| `DataFields`                                  | []*string*                                    | :heavy_minus_sign:                            | Fields with the cell values (i.e. aggregates) |
| `QualifierFields`                             | []*string*                                    | :heavy_minus_sign:                            | Fields to qualify or group data fields        |