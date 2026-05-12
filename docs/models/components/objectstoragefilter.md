# ObjectStorageFilter


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `DataPathFormat`                                                                    | [*components.PathFilterDataFormat](../../models/components/pathfilterdataformat.md) | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `DataTypeID`                                                                        | `string`                                                                            | :heavy_check_mark:                                                                  | Datatype identifier that maps filtered objects to a data type definition.           |
| `Filter`                                                                            | `string`                                                                            | :heavy_check_mark:                                                                  | Glob pattern for selecting files within the storage path.                           |