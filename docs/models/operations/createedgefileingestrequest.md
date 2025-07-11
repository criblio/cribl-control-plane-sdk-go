# CreateEdgeFileIngestRequest


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `FilePath`                                                      | **string*                                                       | :heavy_minus_sign:                                              | Absolute path to file to ingest.                                |
| `PipelineID`                                                    | **string*                                                       | :heavy_minus_sign:                                              | Id of the pipeline to use.                                      |
| `OutputID`                                                      | **string*                                                       | :heavy_minus_sign:                                              | Destination to send events to.                                  |
| `PreProcessingPipelineID`                                       | **string*                                                       | :heavy_minus_sign:                                              | Id to the pre-processing pipeline to use for routes.            |
| `SendToRoutes`                                                  | **string*                                                       | :heavy_minus_sign:                                              | boolean condition required on whether to send events to routes. |
| `BreakerRuleSet`                                                | **string*                                                       | :heavy_minus_sign:                                              | Breaker rules to use on the file.                               |