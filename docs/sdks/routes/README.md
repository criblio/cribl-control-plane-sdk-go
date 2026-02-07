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
        Comments: []components.RouteComment{
            components.RouteComment{
                Comment: "New range of formal shirts are designed keeping you in mind. With fits and styling that will make you stand apart",
                GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                ID: "<id>",
                Index: 3844.57,
            },
        },
        Groups: map[string]components.RoutesGroups{
            "key": components.RoutesGroups{
                Description: criblcontrolplanesdkgo.Pointer("ugh eyeliner authorized even burgeon chime expansion boldly midst and"),
                Index: 8899.36,
                Name: "<value>",
            },
        },
        ID: "default",
        Routes: []components.RouteConf{
            components.RouteConf{
                Clones: []map[string]string{
                    map[string]string{
                        "key": "<value>",
                    },
                },
                Context: criblcontrolplanesdkgo.Pointer("<value>"),
                Description: criblcontrolplanesdkgo.Pointer("Route access logs to main pipeline"),
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                EnableOutputExpression: criblcontrolplanesdkgo.Pointer(false),
                Filter: criblcontrolplanesdkgo.Pointer("source == \"access.log\""),
                Final: true,
                GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                ID: "default",
                Name: "my-route",
                Output: criblcontrolplanesdkgo.Pointer("<value>"),
                OutputExpression: criblcontrolplanesdkgo.Pointer("<value>"),
                Pipeline: "main",
                TargetContext: components.TargetContextGroup.ToPointer(),
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
        Comments: []components.RouteComment{
            components.RouteComment{
                Comment: "New range of formal shirts are designed keeping you in mind. With fits and styling that will make you stand apart",
                GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                ID: "<id>",
                Index: 1800.08,
            },
        },
        Groups: map[string]components.RoutesGroups{
            "key": components.RoutesGroups{
                Description: criblcontrolplanesdkgo.Pointer("ugh eyeliner authorized even burgeon chime expansion boldly midst and"),
                Index: 557.98,
                Name: "<value>",
            },
        },
        ID: "default",
        Routes: []components.RouteConf{
            components.RouteConf{
                Clones: []map[string]string{
                    map[string]string{
                        "key": "<value>",
                    },
                    map[string]string{
                        "key": "<value>",
                        "key1": "<value>",
                        "key2": "<value>",
                    },
                },
                Context: criblcontrolplanesdkgo.Pointer("<value>"),
                Description: criblcontrolplanesdkgo.Pointer("Route speedtest logs"),
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                EnableOutputExpression: criblcontrolplanesdkgo.Pointer(false),
                Filter: criblcontrolplanesdkgo.Pointer("source == \"speedtest.log\""),
                Final: false,
                GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                ID: "route-speedtest",
                Name: "speedtest",
                Output: criblcontrolplanesdkgo.Pointer("default"),
                OutputExpression: criblcontrolplanesdkgo.Pointer("<value>"),
                Pipeline: "main",
                TargetContext: components.TargetContextPack.ToPointer(),
            },
            components.RouteConf{
                Clones: []map[string]string{
                    map[string]string{
                        "key": "<value>",
                        "key1": "<value>",
                        "key2": "<value>",
                    },
                    map[string]string{
                        "key": "<value>",
                    },
                    map[string]string{
                        "key": "<value>",
                        "key1": "<value>",
                        "key2": "<value>",
                    },
                },
                Context: criblcontrolplanesdkgo.Pointer("<value>"),
                Description: criblcontrolplanesdkgo.Pointer("Route mtr logs"),
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                EnableOutputExpression: criblcontrolplanesdkgo.Pointer(false),
                Filter: criblcontrolplanesdkgo.Pointer("source == \"mtr.log\""),
                Final: false,
                GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                ID: "route-mtr",
                Name: "mtr",
                Output: criblcontrolplanesdkgo.Pointer("default"),
                OutputExpression: criblcontrolplanesdkgo.Pointer("<value>"),
                Pipeline: "passthru",
                TargetContext: components.TargetContextGroup.ToPointer(),
            },
            components.RouteConf{
                Clones: []map[string]string{
                    map[string]string{
                        "key": "<value>",
                        "key1": "<value>",
                        "key2": "<value>",
                    },
                    map[string]string{
                        "key": "<value>",
                        "key1": "<value>",
                    },
                    map[string]string{

                    },
                },
                Context: criblcontrolplanesdkgo.Pointer("<value>"),
                Description: criblcontrolplanesdkgo.Pointer("Route statsd metrics"),
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                EnableOutputExpression: criblcontrolplanesdkgo.Pointer(false),
                Filter: criblcontrolplanesdkgo.Pointer("source == \"statsd.log\""),
                Final: false,
                GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                ID: "route-statsd",
                Name: "statsd",
                Output: criblcontrolplanesdkgo.Pointer("devnull"),
                OutputExpression: criblcontrolplanesdkgo.Pointer("<value>"),
                Pipeline: "prometheus_metrics",
                TargetContext: components.TargetContextGroup.ToPointer(),
            },
            components.RouteConf{
                Clones: []map[string]string{
                    map[string]string{

                    },
                    map[string]string{
                        "key": "<value>",
                        "key1": "<value>",
                    },
                    map[string]string{
                        "key": "<value>",
                        "key1": "<value>",
                    },
                },
                Context: criblcontrolplanesdkgo.Pointer("<value>"),
                Description: criblcontrolplanesdkgo.Pointer("Catch-all Route for all other events"),
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                EnableOutputExpression: criblcontrolplanesdkgo.Pointer(false),
                Filter: criblcontrolplanesdkgo.Pointer("true"),
                Final: true,
                GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                ID: "route-default",
                Name: "default",
                Output: criblcontrolplanesdkgo.Pointer("default"),
                OutputExpression: criblcontrolplanesdkgo.Pointer("<value>"),
                Pipeline: "main",
                TargetContext: components.TargetContextPack.ToPointer(),
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
        Comments: []components.RouteComment{
            components.RouteComment{
                Comment: "New range of formal shirts are designed keeping you in mind. With fits and styling that will make you stand apart",
                GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                ID: "<id>",
                Index: 2035.39,
            },
        },
        Groups: map[string]components.RoutesGroups{
            "key": components.RoutesGroups{
                Description: criblcontrolplanesdkgo.Pointer("ugh eyeliner authorized even burgeon chime expansion boldly midst and"),
                Index: 3962.06,
                Name: "<value>",
            },
        },
        ID: "default",
        Routes: []components.RouteConf{
            components.RouteConf{
                Clones: []map[string]string{
                    map[string]string{

                    },
                    map[string]string{
                        "key": "<value>",
                        "key1": "<value>",
                        "key2": "<value>",
                    },
                    map[string]string{

                    },
                },
                Context: criblcontrolplanesdkgo.Pointer("<value>"),
                Description: criblcontrolplanesdkgo.Pointer("Route with dynamic Destination based on environment"),
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                EnableOutputExpression: criblcontrolplanesdkgo.Pointer(true),
                Filter: criblcontrolplanesdkgo.Pointer("source == \"dynamic.log\""),
                Final: true,
                GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                ID: "route-dynamic",
                Name: "dynamic-output",
                Output: criblcontrolplanesdkgo.Pointer("<value>"),
                OutputExpression: criblcontrolplanesdkgo.Pointer("`myDest_${C.logStreamEnv}`"),
                Pipeline: "main",
                TargetContext: components.TargetContextPack.ToPointer(),
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
            TargetContext: components.TargetContextGroup.ToPointer(),
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
            TargetContext: components.TargetContextGroup.ToPointer(),
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
            TargetContext: components.TargetContextGroup.ToPointer(),
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
            TargetContext: components.TargetContextGroup.ToPointer(),
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