# UpdateSavedJobByIDRequest


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ID`                                                                               | *string*                                                                           | :heavy_check_mark:                                                                 | The <code>id</code> of the Collector to update.                                    |
| `CriblPack`                                                                        | **string*                                                                          | :heavy_minus_sign:                                                                 | The <code>id</code> of the Pack that includes the Collector to update.             |
| `SavedJobCreateUpdate`                                                             | [components.SavedJobCreateUpdate](../../models/components/savedjobcreateupdate.md) | :heavy_check_mark:                                                                 | SavedJobCreateUpdate object                                                        |