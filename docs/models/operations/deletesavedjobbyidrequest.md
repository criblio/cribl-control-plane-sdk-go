# DeleteSavedJobByIDRequest


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ID`                                                                           | *string*                                                                       | :heavy_check_mark:                                                             | The <code>id</code> of the Collector to delete.                                |
| `CriblPack`                                                                    | **string*                                                                      | :heavy_minus_sign:                                                             | The <code>id</code> of the Pack that includes the Collector to delete.         |
| `GroupID`                                                                      | **string*                                                                      | :heavy_minus_sign:                                                             | The <code>id</code> of the Worker Group that includes the Collector to delete. |