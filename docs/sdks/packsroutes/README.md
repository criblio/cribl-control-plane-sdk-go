# Packs.Routes

## Overview

### Available Operations

* [Get](#get) - Get a Routing table within a Pack
* [Update](#update) - Update a Route within a Pack
* [List](#list) - List all Routes within a Pack
* [Append](#append) - Add a Route to the end of the Routing table within a Pack

## Get

Get the specified Routing table within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="getRoutesByPackAndId" method="get" path="/p/{pack}/routes/{id}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Routes.Get(ctx, "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedRoutes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `ctx`                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                         | :heavy_check_mark:                                                                            | The context to use for the request.                                                           |
| `id`                                                                                          | *string*                                                                                      | :heavy_check_mark:                                                                            | The <code>id</code> of the Routing table to get. The supported value is <code>default</code>. |
| `pack`                                                                                        | *string*                                                                                      | :heavy_check_mark:                                                                            | The <code>id</code> of the Pack to get.                                                       |
| `opts`                                                                                        | [][operations.Option](../../models/operations/option.md)                                      | :heavy_minus_sign:                                                                            | The options for this request.                                                                 |

### Response

**[*operations.GetRoutesByPackAndIDResponse](../../models/operations/getroutesbypackandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the specified Route within the specified Pack.</br></br>Provide a complete representation of the Routing table, including the Route that you want to update, in the request body. This endpoint does not support partial updates. Cribl removes any omitted Routes and fields when updating.</br></br>Confirm that the configuration in your request body is correct before sending the request. If the configuration is incorrect, the Routing table might not function as expected.

### Example Usage: RoutesUpdateExamplesBasicRoute

<!-- UsageSnippet language="go" operationID="updateRoutesByPackAndId" method="patch" path="/p/{pack}/routes/{id}" example="RoutesUpdateExamplesBasicRoute" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Routes.Update(ctx, "<id>", "<value>", components.Routes{
        ID: "default",
        Routes: []components.RouteConf{
            components.RouteConf{
                Description: criblcontrolplanesdkgo.Pointer("Route access logs to main pipeline"),
                Filter: criblcontrolplanesdkgo.Pointer("source == \"access.log\""),
                Final: true,
                ID: "default",
                Name: "my-route",
                Pipeline: "main",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedRoutes != nil {
        // handle response
    }
}
```
### Example Usage: RoutesUpdateExamplesMultipleRoutes

<!-- UsageSnippet language="go" operationID="updateRoutesByPackAndId" method="patch" path="/p/{pack}/routes/{id}" example="RoutesUpdateExamplesMultipleRoutes" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Routes.Update(ctx, "<id>", "<value>", components.Routes{
        ID: "default",
        Routes: []components.RouteConf{
            components.RouteConf{
                Description: criblcontrolplanesdkgo.Pointer("Route speedtest logs"),
                Filter: criblcontrolplanesdkgo.Pointer("source == \"speedtest.log\""),
                Final: false,
                ID: "route-speedtest",
                Name: "speedtest",
                Output: criblcontrolplanesdkgo.Pointer("default"),
                Pipeline: "main",
            },
            components.RouteConf{
                Description: criblcontrolplanesdkgo.Pointer("Route mtr logs"),
                Filter: criblcontrolplanesdkgo.Pointer("source == \"mtr.log\""),
                Final: false,
                ID: "route-mtr",
                Name: "mtr",
                Output: criblcontrolplanesdkgo.Pointer("default"),
                Pipeline: "passthru",
            },
            components.RouteConf{
                Description: criblcontrolplanesdkgo.Pointer("Route statsd metrics"),
                Filter: criblcontrolplanesdkgo.Pointer("source == \"statsd.log\""),
                Final: false,
                ID: "route-statsd",
                Name: "statsd",
                Output: criblcontrolplanesdkgo.Pointer("devnull"),
                Pipeline: "prometheus_metrics",
            },
            components.RouteConf{
                Description: criblcontrolplanesdkgo.Pointer("Catch-all Route for all other events"),
                Filter: criblcontrolplanesdkgo.Pointer("true"),
                Final: true,
                ID: "route-default",
                Name: "default",
                Output: criblcontrolplanesdkgo.Pointer("default"),
                Pipeline: "main",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedRoutes != nil {
        // handle response
    }
}
```
### Example Usage: RoutesUpdateExamplesRouteWithOutputExpression

<!-- UsageSnippet language="go" operationID="updateRoutesByPackAndId" method="patch" path="/p/{pack}/routes/{id}" example="RoutesUpdateExamplesRouteWithOutputExpression" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Routes.Update(ctx, "<id>", "<value>", components.Routes{
        ID: "default",
        Routes: []components.RouteConf{
            components.RouteConf{
                Description: criblcontrolplanesdkgo.Pointer("Route with dynamic Destination based on environment"),
                EnableOutputExpression: criblcontrolplanesdkgo.Pointer(true),
                Filter: criblcontrolplanesdkgo.Pointer("source == \"dynamic.log\""),
                Final: true,
                ID: "route-dynamic",
                Name: "dynamic-output",
                OutputExpression: criblcontrolplanesdkgo.Pointer("`myDest_${C.logStreamEnv}`"),
                Pipeline: "main",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedRoutes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `id`                                                                                                                     | *string*                                                                                                                 | :heavy_check_mark:                                                                                                       | The <code>id</code> of the Routing table that contains the Route to update. The supported value is <code>default</code>. |
| `pack`                                                                                                                   | *string*                                                                                                                 | :heavy_check_mark:                                                                                                       | The <code>id</code> of the Pack to update.                                                                               |
| `routes`                                                                                                                 | [components.Routes](../../models/components/routes.md)                                                                   | :heavy_check_mark:                                                                                                       | Routes object                                                                                                            |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.UpdateRoutesByPackAndIDResponse](../../models/operations/updateroutesbypackandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Get a list of all Routes within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="getRoutesByPack" method="get" path="/p/{pack}/routes" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Routes.List(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedRoutes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `pack`                                                   | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Pack to list.                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetRoutesByPackResponse](../../models/operations/getroutesbypackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Append

Add a Route to the end of the specified Routing table within the specified Pack.

### Example Usage: RoutesAppendExamplesMultipleRoutes

<!-- UsageSnippet language="go" operationID="createRoutesAppendByPackAndId" method="post" path="/p/{pack}/routes/{id}/append" example="RoutesAppendExamplesMultipleRoutes" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Routes.Append(ctx, "<id>", "<value>", []components.RouteConf{
        components.RouteConf{
            Description: criblcontrolplanesdkgo.Pointer("Route audit logs"),
            Filter: criblcontrolplanesdkgo.Pointer("source == \"audit.log\""),
            Final: false,
            ID: "route-audit",
            Name: "audit",
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Pipeline: "main",
        },
        components.RouteConf{
            Description: criblcontrolplanesdkgo.Pointer("Route security logs"),
            Filter: criblcontrolplanesdkgo.Pointer("source == \"security.log\""),
            Final: false,
            ID: "route-security",
            Name: "security",
            Output: criblcontrolplanesdkgo.Pointer("devnull"),
            Pipeline: "passthru",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedRoutes != nil {
        // handle response
    }
}
```
### Example Usage: RoutesAppendExamplesRouteWithOutputExpression

<!-- UsageSnippet language="go" operationID="createRoutesAppendByPackAndId" method="post" path="/p/{pack}/routes/{id}/append" example="RoutesAppendExamplesRouteWithOutputExpression" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Routes.Append(ctx, "<id>", "<value>", []components.RouteConf{
        components.RouteConf{
            Description: criblcontrolplanesdkgo.Pointer("Route with dynamic Destination based on environment"),
            EnableOutputExpression: criblcontrolplanesdkgo.Pointer(true),
            Filter: criblcontrolplanesdkgo.Pointer("source == \"dynamic.log\""),
            Final: true,
            ID: "route-dynamic-append",
            Name: "dynamic-append",
            OutputExpression: criblcontrolplanesdkgo.Pointer("`myDest_${C.logStreamEnv}`"),
            Pipeline: "main",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedRoutes != nil {
        // handle response
    }
}
```
### Example Usage: RoutesAppendExamplesSingleRoute

<!-- UsageSnippet language="go" operationID="createRoutesAppendByPackAndId" method="post" path="/p/{pack}/routes/{id}/append" example="RoutesAppendExamplesSingleRoute" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Routes.Append(ctx, "<id>", "<value>", []components.RouteConf{
        components.RouteConf{
            Description: criblcontrolplanesdkgo.Pointer("Route new logs to main pipeline"),
            Filter: criblcontrolplanesdkgo.Pointer("source == \"new.log\""),
            Final: true,
            ID: "route-new",
            Name: "new-route",
            Pipeline: "main",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedRoutes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `id`                                                                                                       | *string*                                                                                                   | :heavy_check_mark:                                                                                         | The <code>id</code> of the Routing table to add the Route to. The supported value is <code>default</code>. |
| `pack`                                                                                                     | *string*                                                                                                   | :heavy_check_mark:                                                                                         | The <code>id</code> of the Pack to append.                                                                 |
| `requestBody`                                                                                              | [][components.RouteConf](../../models/components/routeconf.md)                                             | :heavy_check_mark:                                                                                         | RouteDefinitions object                                                                                    |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateRoutesAppendByPackAndIDResponse](../../models/operations/createroutesappendbypackandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |