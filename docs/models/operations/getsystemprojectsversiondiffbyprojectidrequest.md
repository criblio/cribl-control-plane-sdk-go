# GetSystemProjectsVersionDiffByProjectIDRequest


## Fields

| Field                           | Type                            | Required                        | Description                     |
| ------------------------------- | ------------------------------- | ------------------------------- | ------------------------------- |
| `ProjectID`                     | *string*                        | :heavy_check_mark:              | Project Id                      |
| `Commit`                        | **string*                       | :heavy_minus_sign:              | Commit hash (default is HEAD)   |
| `Filename`                      | **string*                       | :heavy_minus_sign:              | Filename                        |
| `DiffLineLimit`                 | **float64*                      | :heavy_minus_sign:              | Limit maximum lines in the diff |