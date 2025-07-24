# CreateRoutesAppendByIDRequest


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ID`                                                                              | *string*                                                                          | :heavy_check_mark:                                                                | the route table to be appended to - currently default is the only supported value |
| `RequestBody`                                                                     | [][components.RouteConf](../../models/components/routeconf.md)                    | :heavy_check_mark:                                                                | RouteDefinitions object                                                           |