# UpdateGlobalVariableLibVarsByPackAndIDRequest


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ID`                                                         | *string*                                                     | :heavy_check_mark:                                           | Unique ID to PATCH for pack                                  |
| `Pack`                                                       | *string*                                                     | :heavy_check_mark:                                           | pack ID to PATCH                                             |
| `GlobalVar`                                                  | [components.GlobalVar](../../models/components/globalvar.md) | :heavy_check_mark:                                           | Global Variable object to be updated                         |