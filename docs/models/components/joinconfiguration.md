# JoinConfiguration


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Kind`                                                                   | **string*                                                                | :heavy_minus_sign:                                                       | Join kind, e.g. inner                                                    |
| `Hints`                                                                  | map[string]*string*                                                      | :heavy_minus_sign:                                                       | Hints passed to the join function                                        |
| `FieldConditions`                                                        | [][components.FieldCondition](../../models/components/fieldcondition.md) | :heavy_minus_sign:                                                       | Fields to use when joining                                               |
| `SearchJobID`                                                            | **string*                                                                | :heavy_minus_sign:                                                       | The id for this search job.                                              |
| `StageID`                                                                | **string*                                                                | :heavy_minus_sign:                                                       | The stage we are joining with.                                           |