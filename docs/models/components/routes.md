# Routes


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `ID`                                                                          | **string*                                                                     | :heavy_minus_sign:                                                            | Routes ID                                                                     |
| `Routes`                                                                      | [][components.RoutesRoute](../../models/components/routesroute.md)            | :heavy_check_mark:                                                            | Pipeline routing rules                                                        |
| `Groups`                                                                      | map[string][components.RoutesGroups](../../models/components/routesgroups.md) | :heavy_minus_sign:                                                            | N/A                                                                           |
| `Comments`                                                                    | [][components.Comment](../../models/components/comment.md)                    | :heavy_minus_sign:                                                            | Comments                                                                      |