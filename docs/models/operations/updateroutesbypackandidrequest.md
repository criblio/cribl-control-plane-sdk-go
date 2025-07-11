# UpdateRoutesByPackAndIDRequest


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ID`                                                             | *string*                                                         | :heavy_check_mark:                                               | Unique ID to PATCH for pack                                      |
| `Pack`                                                           | *string*                                                         | :heavy_check_mark:                                               | pack ID to PATCH                                                 |
| `Routes`                                                         | [components.RoutesInput](../../models/components/routesinput.md) | :heavy_check_mark:                                               | Routes object to be updated                                      |