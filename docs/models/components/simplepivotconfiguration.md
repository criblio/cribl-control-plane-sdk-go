# SimplePivotConfiguration


## Fields

| Field                                         | Type                                          | Required                                      | Description                                   |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| `LabelField`                                  | *string*                                      | :heavy_check_mark:                            | Fields to be used for the left-most column.   |
| `DataFields`                                  | []*string*                                    | :heavy_check_mark:                            | Fields with the cell values (i.e. aggregates) |
| `QualifierFields`                             | []*string*                                    | :heavy_check_mark:                            | Fields to qualify or group data fields        |