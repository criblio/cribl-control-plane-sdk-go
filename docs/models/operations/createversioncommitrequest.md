# CreateVersionCommitRequest


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `GroupID`                                                                         | **string*                                                                         | :heavy_minus_sign:                                                                | The <code>id</code> of the Worker Group or Edge Fleet to create a new commit for. |
| `GitCommitParams`                                                                 | [components.GitCommitParams](../../models/components/gitcommitparams.md)          | :heavy_check_mark:                                                                | GitCommitParams object                                                            |