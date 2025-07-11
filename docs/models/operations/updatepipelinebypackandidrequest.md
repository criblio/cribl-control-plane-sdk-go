# UpdatePipelineByPackAndIDRequest


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ID`                                                       | *string*                                                   | :heavy_check_mark:                                         | Unique ID to PATCH for pack                                |
| `Pack`                                                     | *string*                                                   | :heavy_check_mark:                                         | pack ID to PATCH                                           |
| `Pipeline`                                                 | [components.Pipeline](../../models/components/pipeline.md) | :heavy_check_mark:                                         | Pipeline object to be updated                              |