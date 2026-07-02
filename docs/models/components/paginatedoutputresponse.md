# PaginatedOutputResponse


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Items`                                                                  | [][components.OutputResponse](../../models/components/outputresponse.md) | :heavy_check_mark:                                                       | the pre-limited items in the list of results                             |
| `Count`                                                                  | `int64`                                                                  | :heavy_check_mark:                                                       | number of items present in the items array                               |
| `Offset`                                                                 | `int64`                                                                  | :heavy_check_mark:                                                       | pagination offset                                                        |
| `Limit`                                                                  | `int64`                                                                  | :heavy_check_mark:                                                       | pagination limit                                                         |
| `TotalCount`                                                             | `*int64`                                                                 | :heavy_minus_sign:                                                       | total number of items available (present when limit is set)              |