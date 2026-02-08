# Routes

## Overview

Actions related to Routes

### Available Operations

* [List](#list) - List all Routes
* [Get](#get) - Get a Routing table
* [Update](#update) - Update a Route
* [Append](#append) - Add a Route to the end of the Routing table

## List

Get a list of all Routes.

### Example Usage

<!-- UsageSnippet language="go" operationID="listRoutes" method="get" path="/routes" -->
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

**[*operations.ListRoutesResponse](../../models/operations/listroutesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

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
| `id`                                                                                          | *string*                                                                                      | :heavy_check_mark:                                                                            | The <code>id</code> of the Routing table to get. The supported value is <code>default</code>. |
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

    res, err := s.Routes.Update(ctx, "<id>", components.Routes{
        ID: criblcontrolplanesdkgo.Pointer("default"),
        Routes: []components.RoutesRoute{
            components.RoutesRoute{
                ID: criblcontrolplanesdkgo.Pointer("default"),
                Name: "my-route",
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                Filter: criblcontrolplanesdkgo.Pointer("source == \"access.log\""),
                Pipeline: "main",
                EnableOutputExpression: criblcontrolplanesdkgo.Pointer(false),
                Output: "<value>",
                OutputExpression: "<value>",
                Description: criblcontrolplanesdkgo.Pointer("Route access logs to main pipeline"),
                Final: criblcontrolplanesdkgo.Pointer(true),
            },
        },
        Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{
            "key": components.AdditionalPropertiesTypePipelineConfGroups{
                Name: "<value>",
                Description: criblcontrolplanesdkgo.Pointer("drat yet spectacles ha"),
                Disabled: criblcontrolplanesdkgo.Pointer(false),
            },
        },
        Comments: []components.Comment{
            components.Comment{
                Comment: criblcontrolplanesdkgo.Pointer("The Football Is Good For Training And Recreational Purposes"),
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

    res, err := s.Routes.Update(ctx, "<id>", components.Routes{
        ID: criblcontrolplanesdkgo.Pointer("default"),
        Routes: []components.RoutesRoute{
            components.RoutesRoute{
                ID: criblcontrolplanesdkgo.Pointer("route-speedtest"),
                Name: "speedtest",
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                Filter: criblcontrolplanesdkgo.Pointer("source == \"speedtest.log\""),
                Pipeline: "main",
                EnableOutputExpression: criblcontrolplanesdkgo.Pointer(true),
                Output: "default",
                OutputExpression: "<value>",
                Description: criblcontrolplanesdkgo.Pointer("Route speedtest logs"),
                Final: criblcontrolplanesdkgo.Pointer(false),
            },
            components.RoutesRoute{
                ID: criblcontrolplanesdkgo.Pointer("route-mtr"),
                Name: "mtr",
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                Filter: criblcontrolplanesdkgo.Pointer("source == \"mtr.log\""),
                Pipeline: "passthru",
                EnableOutputExpression: criblcontrolplanesdkgo.Pointer(true),
                Output: "default",
                OutputExpression: "<value>",
                Description: criblcontrolplanesdkgo.Pointer("Route mtr logs"),
                Final: criblcontrolplanesdkgo.Pointer(false),
            },
            components.RoutesRoute{
                ID: criblcontrolplanesdkgo.Pointer("route-statsd"),
                Name: "statsd",
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                Filter: criblcontrolplanesdkgo.Pointer("source == \"statsd.log\""),
                Pipeline: "prometheus_metrics",
                EnableOutputExpression: criblcontrolplanesdkgo.Pointer(false),
                Output: "devnull",
                OutputExpression: "<value>",
                Description: criblcontrolplanesdkgo.Pointer("Route statsd metrics"),
                Final: criblcontrolplanesdkgo.Pointer(false),
            },
            components.RoutesRoute{
                ID: criblcontrolplanesdkgo.Pointer("route-default"),
                Name: "default",
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                Filter: criblcontrolplanesdkgo.Pointer("true"),
                Pipeline: "main",
                EnableOutputExpression: criblcontrolplanesdkgo.Pointer(false),
                Output: "default",
                OutputExpression: "<value>",
                Description: criblcontrolplanesdkgo.Pointer("Catch-all Route for all other events"),
                Final: criblcontrolplanesdkgo.Pointer(true),
            },
        },
        Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{
            "key": components.AdditionalPropertiesTypePipelineConfGroups{
                Name: "<value>",
                Description: criblcontrolplanesdkgo.Pointer("drat yet spectacles ha"),
                Disabled: criblcontrolplanesdkgo.Pointer(false),
            },
        },
        Comments: []components.Comment{
            components.Comment{
                Comment: criblcontrolplanesdkgo.Pointer("The Football Is Good For Training And Recreational Purposes"),
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

    res, err := s.Routes.Update(ctx, "<id>", components.Routes{
        ID: criblcontrolplanesdkgo.Pointer("default"),
        Routes: []components.RoutesRoute{
            components.RoutesRoute{
                ID: criblcontrolplanesdkgo.Pointer("route-dynamic"),
                Name: "dynamic-output",
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                Filter: criblcontrolplanesdkgo.Pointer("source == \"dynamic.log\""),
                Pipeline: "main",
                EnableOutputExpression: criblcontrolplanesdkgo.Pointer(true),
                Output: "<value>",
                OutputExpression: "`myDest_${C.logStreamEnv}`",
                Description: criblcontrolplanesdkgo.Pointer("Route with dynamic Destination based on environment"),
                Final: criblcontrolplanesdkgo.Pointer(true),
            },
        },
        Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{
            "key": components.AdditionalPropertiesTypePipelineConfGroups{
                Name: "<value>",
                Description: criblcontrolplanesdkgo.Pointer("drat yet spectacles ha"),
                Disabled: criblcontrolplanesdkgo.Pointer(false),
            },
        },
        Comments: []components.Comment{
            components.Comment{
                Comment: criblcontrolplanesdkgo.Pointer("The Football Is Good For Training And Recreational Purposes"),
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
| `routes`                                                                                                                 | [components.Routes](../../models/components/routes.md)                                                                   | :heavy_check_mark:                                                                                                       | Routes object                                                                                                            |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.UpdateRoutesByIDResponse](../../models/operations/updateroutesbyidresponse.md), error**

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

    res, err := s.Routes.Append(ctx, "<id>", []components.RouteConf{
        components.RouteConf{
            Clones: []map[string]string{
                map[string]string{
                    "key": "<value>",
                    "key1": "<value>",
                },
            },
            Context: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("Route audit logs"),
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            EnableOutputExpression: criblcontrolplanesdkgo.Pointer(false),
            Filter: criblcontrolplanesdkgo.Pointer("source == \"audit.log\""),
            Final: false,
            GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
            ID: "route-audit",
            Name: "audit",
            Output: criblcontrolplanesdkgo.Pointer("default"),
            OutputExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            Pipeline: "main",
        },
        components.RouteConf{
            Clones: []map[string]string{
                map[string]string{

                },
                map[string]string{

                },
                map[string]string{

                },
            },
            Context: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("Route security logs"),
            Disabled: criblcontrolplanesdkgo.Pointer(true),
            EnableOutputExpression: criblcontrolplanesdkgo.Pointer(false),
            Filter: criblcontrolplanesdkgo.Pointer("source == \"security.log\""),
            Final: false,
            GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
            ID: "route-security",
            Name: "security",
            Output: criblcontrolplanesdkgo.Pointer("devnull"),
            OutputExpression: criblcontrolplanesdkgo.Pointer("<value>"),
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

    res, err := s.Routes.Append(ctx, "<id>", []components.RouteConf{
        components.RouteConf{
            Clones: []map[string]string{
                map[string]string{
                    "key": "<value>",
                    "key1": "<value>",
                },
            },
            Context: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("Route with dynamic Destination based on environment"),
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            EnableOutputExpression: criblcontrolplanesdkgo.Pointer(true),
            Filter: criblcontrolplanesdkgo.Pointer("source == \"dynamic.log\""),
            Final: true,
            GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
            ID: "route-dynamic-append",
            Name: "dynamic-append",
            Output: criblcontrolplanesdkgo.Pointer("<value>"),
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

    res, err := s.Routes.Append(ctx, "<id>", []components.RouteConf{
        components.RouteConf{
            Clones: []map[string]string{
                map[string]string{
                    "key": "<value>",
                },
                map[string]string{

                },
            },
            Context: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("Route new logs to main pipeline"),
            Disabled: criblcontrolplanesdkgo.Pointer(true),
            EnableOutputExpression: criblcontrolplanesdkgo.Pointer(true),
            Filter: criblcontrolplanesdkgo.Pointer("source == \"new.log\""),
            Final: true,
            GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
            ID: "route-new",
            Name: "new-route",
            Output: criblcontrolplanesdkgo.Pointer("<value>"),
            OutputExpression: criblcontrolplanesdkgo.Pointer("<value>"),
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
| `requestBody`                                                                                              | [][components.RouteConf](../../models/components/routeconf.md)                                             | :heavy_check_mark:                                                                                         | RouteDefinitions object                                                                                    |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateRoutesAppendByIDResponse](../../models/operations/createroutesappendbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |