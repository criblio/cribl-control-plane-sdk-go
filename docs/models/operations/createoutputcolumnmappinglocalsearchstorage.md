# CreateOutputColumnMappingLocalSearchStorage


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ColumnName`                                                         | *string*                                                             | :heavy_check_mark:                                                   | Name of the column that will store field value                       |
| `ColumnType`                                                         | **string*                                                            | :heavy_minus_sign:                                                   | Type of the column in the database                                   |
| `ColumnValueExpression`                                              | *string*                                                             | :heavy_check_mark:                                                   | JavaScript expression to compute value to be inserted into the table |