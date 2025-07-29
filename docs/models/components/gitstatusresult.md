# GitStatusResult


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Ahead`                                                    | *float64*                                                  | :heavy_check_mark:                                         | N/A                                                        |
| `Behind`                                                   | *float64*                                                  | :heavy_check_mark:                                         | N/A                                                        |
| `Conflicted`                                               | []*string*                                                 | :heavy_check_mark:                                         | N/A                                                        |
| `Created`                                                  | []*string*                                                 | :heavy_check_mark:                                         | N/A                                                        |
| `Current`                                                  | *string*                                                   | :heavy_check_mark:                                         | N/A                                                        |
| `Deleted`                                                  | []*string*                                                 | :heavy_check_mark:                                         | N/A                                                        |
| `Files`                                                    | [][components.File](../../models/components/file.md)       | :heavy_check_mark:                                         | N/A                                                        |
| `Modified`                                                 | []*string*                                                 | :heavy_check_mark:                                         | N/A                                                        |
| `NotAdded`                                                 | []*string*                                                 | :heavy_check_mark:                                         | N/A                                                        |
| `Renamed`                                                  | [][components.Renamed](../../models/components/renamed.md) | :heavy_check_mark:                                         | N/A                                                        |
| `Staged`                                                   | []*string*                                                 | :heavy_check_mark:                                         | N/A                                                        |
| `Tracking`                                                 | *string*                                                   | :heavy_check_mark:                                         | N/A                                                        |