# GetJobsResponseBody

a list of JobInfo objects


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Offset`                                                   | **int64*                                                   | :heavy_minus_sign:                                         | pagination offset                                          |
| `Limit`                                                    | **int64*                                                   | :heavy_minus_sign:                                         | number of items present in the items array                 |
| `Items`                                                    | [][components.JobInfo](../../models/components/jobinfo.md) | :heavy_minus_sign:                                         | the pre-limited items in the list of results               |