# GitFilesResponse


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `CommitMessage`                                            | `*string`                                                  | :heavy_minus_sign:                                         | Commit message of the specified commit.                    |
| `Count`                                                    | `int64`                                                    | :heavy_check_mark:                                         | Number of files returned.                                  |
| `Items`                                                    | [][components.GitFile](../../models/components/gitfile.md) | :heavy_check_mark:                                         | Array of files that changed since the specified commit.    |