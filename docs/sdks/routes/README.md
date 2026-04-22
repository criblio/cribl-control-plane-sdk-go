# Routes

## Overview

Actions related to Routes

### Available Operations

* [Get](#get) - Get a Routing table
* [Update](#update) - Update a Route
* [List](#list) - List all Routes
* [Append](#append) - Add a Route to the end of the Routing table

## Get

Get the specified Routing table.

### Example Usage

<!-- UsageSnippet language="go" operationID="getRoutesById" method="get" path="/routes/{id}" -->
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
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Routes.Get(ctx, "<id>")
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
| `id`                                                                                          | `string`                                                                                      | :heavy_check_mark:                                                                            | The <code>id</code> of the Routing table to get. The supported value is <code>default</code>. |
| `opts`                                                                                        | [][operations.Option](../../models/operations/option.md)                                      | :heavy_minus_sign:                                                                            | The options for this request.                                                                 |

### Response

**[*operations.GetRoutesByIDResponse](../../models/operations/getroutesbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update a Route in the specified Routing table.</br></br>Provide a complete representation of the Routing table, including the Route that you want to update, in the request body. This endpoint does not support partial updates. Cribl removes any omitted Routes and fields when updating.</br></br>Confirm that the configuration in your request body is correct before sending the request. If the configuration is incorrect, the Routing table might not function as expected.

### Example Usage: RoutesUpdateExamplesBasicRoute

<!-- UsageSnippet language="go" operationID="updateRoutesById" method="patch" path="/routes/{id}" example="RoutesUpdateExamplesBasicRoute" -->
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
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Routes.Update(ctx, "<id>", components.RoutesInput{
        ID: "default",
        Routes: []components.RouteConfInput{
            components.RouteConfInput{
                Description: criblcontrolplanesdkgo.Pointer("Route access logs to main Pipeline"),
                Filter: criblcontrolplanesdkgo.Pointer("source == \"access.log\""),
                Name: "my-route",
                Pipeline: "main",
                Final: criblcontrolplanesdkgo.Pointer(true),
                ID: criblcontrolplanesdkgo.Pointer("default"),
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

<!-- UsageSnippet language="go" operationID="updateRoutesById" method="patch" path="/routes/{id}" example="RoutesUpdateExamplesMultipleRoutes" -->
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
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Routes.Update(ctx, "<id>", components.RoutesInput{
        ID: "default",
        Routes: []components.RouteConfInput{
            components.RouteConfInput{
                Description: criblcontrolplanesdkgo.Pointer("Route speedtest logs"),
                Filter: criblcontrolplanesdkgo.Pointer("source == \"speedtest.log\""),
                Name: "speedtest",
                Output: criblcontrolplanesdkgo.Pointer("default"),
                Pipeline: "main",
                Final: criblcontrolplanesdkgo.Pointer(false),
                ID: criblcontrolplanesdkgo.Pointer("route-speedtest"),
            },
            components.RouteConfInput{
                Description: criblcontrolplanesdkgo.Pointer("Route mtr logs"),
                Filter: criblcontrolplanesdkgo.Pointer("source == \"mtr.log\""),
                Name: "mtr",
                Output: criblcontrolplanesdkgo.Pointer("default"),
                Pipeline: "passthru",
                Final: criblcontrolplanesdkgo.Pointer(false),
                ID: criblcontrolplanesdkgo.Pointer("route-mtr"),
            },
            components.RouteConfInput{
                Description: criblcontrolplanesdkgo.Pointer("Route statsd metrics"),
                Filter: criblcontrolplanesdkgo.Pointer("source == \"statsd.log\""),
                Name: "statsd",
                Output: criblcontrolplanesdkgo.Pointer("devnull"),
                Pipeline: "prometheus_metrics",
                Final: criblcontrolplanesdkgo.Pointer(false),
                ID: criblcontrolplanesdkgo.Pointer("route-statsd"),
            },
            components.RouteConfInput{
                Description: criblcontrolplanesdkgo.Pointer("Catch-all Route for all other events"),
                Filter: criblcontrolplanesdkgo.Pointer("true"),
                Name: "default",
                Output: criblcontrolplanesdkgo.Pointer("default"),
                Pipeline: "main",
                Final: criblcontrolplanesdkgo.Pointer(true),
                ID: criblcontrolplanesdkgo.Pointer("route-default"),
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
### Example Usage: RoutesUpdateExamplesRouteWithDefaults

<!-- UsageSnippet language="go" operationID="updateRoutesById" method="patch" path="/routes/{id}" example="RoutesUpdateExamplesRouteWithDefaults" -->
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
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Routes.Update(ctx, "<id>", components.RoutesInput{
        ID: "default",
        Routes: []components.RouteConfInput{
            components.RouteConfInput{
                Description: criblcontrolplanesdkgo.Pointer("Route access logs to main Pipeline"),
                Filter: criblcontrolplanesdkgo.Pointer("source == \"access.log\""),
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
### Example Usage: RoutesUpdateExamplesRouteWithOutputExpression

<!-- UsageSnippet language="go" operationID="updateRoutesById" method="patch" path="/routes/{id}" example="RoutesUpdateExamplesRouteWithOutputExpression" -->
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
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Routes.Update(ctx, "<id>", components.RoutesInput{
        ID: "default",
        Routes: []components.RouteConfInput{
            components.RouteConfInput{
                Description: criblcontrolplanesdkgo.Pointer("Route with dynamic Destination based on environment"),
                EnableOutputExpression: criblcontrolplanesdkgo.Pointer(true),
                Filter: criblcontrolplanesdkgo.Pointer("source == \"dynamic.log\""),
                Name: "dynamic-output",
                OutputExpression: criblcontrolplanesdkgo.Pointer("`myDest_${C.logStreamEnv}`"),
                Pipeline: "main",
                Final: criblcontrolplanesdkgo.Pointer(true),
                ID: criblcontrolplanesdkgo.Pointer("route-dynamic"),
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
| `id`                                                                                                                     | `string`                                                                                                                 | :heavy_check_mark:                                                                                                       | The <code>id</code> of the Routing table that contains the Route to update. The supported value is <code>default</code>. |
| `routesInput`                                                                                                            | [components.RoutesInput](../../models/components/routesinput.md)                                                         | :heavy_check_mark:                                                                                                       | RoutesInput object                                                                                                       |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.UpdateRoutesByIDResponse](../../models/operations/updateroutesbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Get a list of all Routes.

### Example Usage

<!-- UsageSnippet language="go" operationID="getRoutes" method="get" path="/routes" -->
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
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Routes.List(ctx)
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetRoutesResponse](../../models/operations/getroutesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Append

Add a Route to the end of the specified Routing table.

### Example Usage: RoutesAppendExamplesMultipleRoutes

<!-- UsageSnippet language="go" operationID="createRoutesAppendById" method="post" path="/routes/{id}/append" example="RoutesAppendExamplesMultipleRoutes" -->
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
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Routes.Append(ctx, "<id>", []components.RouteConfInput{
        components.RouteConfInput{
            Description: criblcontrolplanesdkgo.Pointer("Route audit logs"),
            Filter: criblcontrolplanesdkgo.Pointer("source == \"audit.log\""),
            Name: "audit",
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Pipeline: "main",
            Final: criblcontrolplanesdkgo.Pointer(false),
            ID: criblcontrolplanesdkgo.Pointer("route-audit"),
        },
        components.RouteConfInput{
            Description: criblcontrolplanesdkgo.Pointer("Route security logs"),
            Filter: criblcontrolplanesdkgo.Pointer("source == \"security.log\""),
            Name: "security",
            Output: criblcontrolplanesdkgo.Pointer("devnull"),
            Pipeline: "passthru",
            Final: criblcontrolplanesdkgo.Pointer(false),
            ID: criblcontrolplanesdkgo.Pointer("route-security"),
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
### Example Usage: RoutesAppendExamplesRouteWithDefaults

<!-- UsageSnippet language="go" operationID="createRoutesAppendById" method="post" path="/routes/{id}/append" example="RoutesAppendExamplesRouteWithDefaults" -->
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
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Routes.Append(ctx, "<id>", []components.RouteConfInput{
        components.RouteConfInput{
            Description: criblcontrolplanesdkgo.Pointer("Route with server-generated id and default final value"),
            Filter: criblcontrolplanesdkgo.Pointer("source == \"new.log\""),
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
### Example Usage: RoutesAppendExamplesRouteWithOutputExpression

<!-- UsageSnippet language="go" operationID="createRoutesAppendById" method="post" path="/routes/{id}/append" example="RoutesAppendExamplesRouteWithOutputExpression" -->
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
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Routes.Append(ctx, "<id>", []components.RouteConfInput{
        components.RouteConfInput{
            Description: criblcontrolplanesdkgo.Pointer("Route with dynamic Destination based on environment"),
            EnableOutputExpression: criblcontrolplanesdkgo.Pointer(true),
            Filter: criblcontrolplanesdkgo.Pointer("source == \"dynamic.log\""),
            Name: "dynamic-append",
            OutputExpression: criblcontrolplanesdkgo.Pointer("`myDest_${C.logStreamEnv}`"),
            Pipeline: "main",
            Final: criblcontrolplanesdkgo.Pointer(true),
            ID: criblcontrolplanesdkgo.Pointer("route-dynamic-append"),
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

<!-- UsageSnippet language="go" operationID="createRoutesAppendById" method="post" path="/routes/{id}/append" example="RoutesAppendExamplesSingleRoute" -->
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
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Routes.Append(ctx, "<id>", []components.RouteConfInput{
        components.RouteConfInput{
            Description: criblcontrolplanesdkgo.Pointer("Route new logs to main pipeline"),
            Filter: criblcontrolplanesdkgo.Pointer("source == \"new.log\""),
            Name: "new-route",
            Pipeline: "main",
            Final: criblcontrolplanesdkgo.Pointer(true),
            ID: criblcontrolplanesdkgo.Pointer("route-new"),
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
| `id`                                                                                                       | `string`                                                                                                   | :heavy_check_mark:                                                                                         | The <code>id</code> of the Routing table to add the Route to. The supported value is <code>default</code>. |
| `requestBody`                                                                                              | [][components.RouteConfInput](../../models/components/routeconfinput.md)                                   | :heavy_check_mark:                                                                                         | RouteDefinitions object                                                                                    |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateRoutesAppendByIDResponse](../../models/operations/createroutesappendbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |