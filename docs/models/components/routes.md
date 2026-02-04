# Routes


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Comments`                                                                    | [][components.RouteComment](../../models/components/routecomment.md)          | :heavy_minus_sign:                                                            | Comments                                                                      |
| `Groups`                                                                      | map[string][components.RoutesGroups](../../models/components/routesgroups.md) | :heavy_minus_sign:                                                            | Map of route groups                                                           |
| `ID`                                                                          | *string*                                                                      | :heavy_check_mark:                                                            | Routes ID                                                                     |
| `Routes`                                                                      | [][components.RouteConf](../../models/components/routeconf.md)                | :heavy_check_mark:                                                            | Pipeline routing rules                                                        |