# UpdateSystemProjectsPipelinesByGroupIDAndProjectIDAndPipelineIDRequest


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `GroupID`                                                  | *string*                                                   | :heavy_check_mark:                                         | Group Id                                                   |
| `ProjectID`                                                | *string*                                                   | :heavy_check_mark:                                         | Project Id                                                 |
| `PipelineID`                                               | *string*                                                   | :heavy_check_mark:                                         | Pipeline iD                                                |
| `Pipeline`                                                 | [components.Pipeline](../../models/components/pipeline.md) | :heavy_check_mark:                                         | Pipeline object to be updated in specified Project         |