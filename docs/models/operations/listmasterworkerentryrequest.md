# ListMasterWorkerEntryRequest


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `FilterExp`                                                          | **string*                                                            | :heavy_minus_sign:                                                   | Filter expression evaluated against nodes                            |
| `Sort`                                                               | **string*                                                            | :heavy_minus_sign:                                                   | Sorting object (JSON stringified) expression evaluated against nodes |
| `SortExp`                                                            | **string*                                                            | :heavy_minus_sign:                                                   | Sorting expression evaluated against nodes                           |
| `Limit`                                                              | **int64*                                                             | :heavy_minus_sign:                                                   | Maximum number of nodes to return                                    |
| `Offset`                                                             | **int64*                                                             | :heavy_minus_sign:                                                   | Pagination offset                                                    |
| `Filter`                                                             | **string*                                                            | :heavy_minus_sign:                                                   | Filter object (JSON stringified) to select nodes                     |