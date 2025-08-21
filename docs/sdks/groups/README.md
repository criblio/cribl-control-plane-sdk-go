# Groups
(*Groups*)

## Overview

Actions related to Groups

### Available Operations

* [List](#list) - List all Worker Groups or Edge Fleets for the specified Cribl product
* [Create](#create) - Create a Worker Group or Edge Fleet for the specified Cribl product
* [Get](#get) - Get a Worker Group or Edge Fleet
* [Update](#update) - Update a Worker Group or Edge Fleet
* [Delete](#delete) - Delete a Worker Group or Edge Fleet
* [Deploy](#deploy) - Deploy commits to a Worker Group or Edge Fleet

## List

Get a list of all Worker Groups or Edge Fleets for the specified Cribl product.

### Example Usage

<!-- UsageSnippet language="go" operationID="listConfigGroupByProduct" method="get" path="/products/{product}/groups" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Groups.List(ctx, components.ProductsCoreEdge, criblcontrolplanesdkgo.String("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                        | Type                                                                                                                                                                             | Required                                                                                                                                                                         | Description                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                            | :heavy_check_mark:                                                                                                                                                               | The context to use for the request.                                                                                                                                              |
| `product`                                                                                                                                                                        | [components.ProductsCore](../../models/components/productscore.md)                                                                                                               | :heavy_check_mark:                                                                                                                                                               | Name of the Cribl product to get the Worker Groups or Edge Fleets for.                                                                                                           |
| `fields`                                                                                                                                                                         | **string*                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                               | Comma-separated list of additional properties to include in the response. Available values are <code>git.commit</code>, <code>git.localChanges</code>, and <code>git.log</code>. |
| `opts`                                                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                                                         | :heavy_minus_sign:                                                                                                                                                               | The options for this request.                                                                                                                                                    |

### Response

**[*operations.ListConfigGroupByProductResponse](../../models/operations/listconfiggroupbyproductresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a new Worker Group or Edge Fleet for the specified Cribl product.

### Example Usage

<!-- UsageSnippet language="go" operationID="createConfigGroupByProduct" method="post" path="/products/{product}/groups" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Groups.Create(ctx, components.ProductsCoreEdge, components.ConfigGroup{
        Cloud: &components.ConfigGroupCloud{
            Provider: components.CloudProviderAws.ToPointer(),
            Region: "<value>",
        },
        ConfigVersion: "<value>",
        DeployingWorkerCount: criblcontrolplanesdkgo.Float64(393.49),
        Description: criblcontrolplanesdkgo.String("ack before fondly scent because gee without where exactly"),
        EstimatedIngestRate: criblcontrolplanesdkgo.Float64(346.37),
        Git: &components.Git{
            Commit: criblcontrolplanesdkgo.String("<value>"),
            LocalChanges: criblcontrolplanesdkgo.Float64(5255.51),
            Log: []components.Commit{
                components.Commit{
                    AuthorEmail: criblcontrolplanesdkgo.String("<value>"),
                    AuthorName: criblcontrolplanesdkgo.String("<value>"),
                    Date: "2024-06-13",
                    Hash: "<value>",
                    Message: "<value>",
                    Short: "<value>",
                },
            },
        },
        ID: "<id>",
        IncompatibleWorkerCount: criblcontrolplanesdkgo.Float64(5613.31),
        Inherits: criblcontrolplanesdkgo.String("<value>"),
        IsFleet: criblcontrolplanesdkgo.Bool(true),
        IsSearch: criblcontrolplanesdkgo.Bool(false),
        LookupDeployments: []components.ConfigGroupLookups{
            components.ConfigGroupLookups{
                Context: "<value>",
                Lookups: []components.ConfigGroupLookupsLookup{},
            },
        },
        MaxWorkerAge: criblcontrolplanesdkgo.String("<value>"),
        Name: criblcontrolplanesdkgo.String("<value>"),
        OnPrem: criblcontrolplanesdkgo.Bool(false),
        Provisioned: criblcontrolplanesdkgo.Bool(true),
        Streamtags: []string{
            "<value 1>",
            "<value 2>",
            "<value 3>",
        },
        Tags: criblcontrolplanesdkgo.String("<value>"),
        Type: components.ConfigGroupTypeLakeAccess.ToPointer(),
        UpgradeVersion: criblcontrolplanesdkgo.String("<value>"),
        WorkerCount: criblcontrolplanesdkgo.Float64(3050.1),
        WorkerRemoteAccess: criblcontrolplanesdkgo.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ctx`                                                               | [context.Context](https://pkg.go.dev/context#Context)               | :heavy_check_mark:                                                  | The context to use for the request.                                 |
| `product`                                                           | [components.ProductsCore](../../models/components/productscore.md)  | :heavy_check_mark:                                                  | Name of the Cribl product to add the Worker Group or Edge Fleet to. |
| `configGroup`                                                       | [components.ConfigGroup](../../models/components/configgroup.md)    | :heavy_check_mark:                                                  | ConfigGroup object                                                  |
| `opts`                                                              | [][operations.Option](../../models/operations/option.md)            | :heavy_minus_sign:                                                  | The options for this request.                                       |

### Response

**[*operations.CreateConfigGroupByProductResponse](../../models/operations/createconfiggroupbyproductresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get the specified Worker Group or Edge Fleet.

### Example Usage

<!-- UsageSnippet language="go" operationID="getConfigGroupByProductAndId" method="get" path="/products/{product}/groups/{id}" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Groups.Get(ctx, components.ProductsCoreEdge, "<id>", criblcontrolplanesdkgo.String("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                        | Type                                                                                                                                                                             | Required                                                                                                                                                                         | Description                                                                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                            | :heavy_check_mark:                                                                                                                                                               | The context to use for the request.                                                                                                                                              |
| `product`                                                                                                                                                                        | [components.ProductsCore](../../models/components/productscore.md)                                                                                                               | :heavy_check_mark:                                                                                                                                                               | Name of the Cribl product to get the Worker Groups or Edge Fleets for.                                                                                                           |
| `id`                                                                                                                                                                             | *string*                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                               | The <code>id</code> of the Worker Group or Edge Fleet to get.                                                                                                                    |
| `fields`                                                                                                                                                                         | **string*                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                               | Comma-separated list of additional properties to include in the response. Available values are <code>git.commit</code>, <code>git.localChanges</code>, and <code>git.log</code>. |
| `opts`                                                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                                                         | :heavy_minus_sign:                                                                                                                                                               | The options for this request.                                                                                                                                                    |

### Response

**[*operations.GetConfigGroupByProductAndIDResponse](../../models/operations/getconfiggroupbyproductandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the specified Worker Group or Edge Fleet.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateConfigGroupByProductAndId" method="patch" path="/products/{product}/groups/{id}" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Groups.Update(ctx, components.ProductsCoreStream, "<id>", components.ConfigGroup{
        Cloud: &components.ConfigGroupCloud{
            Provider: components.CloudProviderAws.ToPointer(),
            Region: "<value>",
        },
        ConfigVersion: "<value>",
        DeployingWorkerCount: criblcontrolplanesdkgo.Float64(7451.49),
        Description: criblcontrolplanesdkgo.String("verbally feminize harmful prance really"),
        EstimatedIngestRate: criblcontrolplanesdkgo.Float64(6748.35),
        Git: &components.Git{
            Commit: criblcontrolplanesdkgo.String("<value>"),
            LocalChanges: criblcontrolplanesdkgo.Float64(4475.22),
            Log: []components.Commit{
                components.Commit{
                    AuthorEmail: criblcontrolplanesdkgo.String("<value>"),
                    AuthorName: criblcontrolplanesdkgo.String("<value>"),
                    Date: "2024-01-27",
                    Hash: "<value>",
                    Message: "<value>",
                    Short: "<value>",
                },
            },
        },
        ID: "<id>",
        IncompatibleWorkerCount: criblcontrolplanesdkgo.Float64(2043.29),
        Inherits: criblcontrolplanesdkgo.String("<value>"),
        IsFleet: criblcontrolplanesdkgo.Bool(false),
        IsSearch: criblcontrolplanesdkgo.Bool(false),
        LookupDeployments: []components.ConfigGroupLookups{
            components.ConfigGroupLookups{
                Context: "<value>",
                Lookups: []components.ConfigGroupLookupsLookup{
                    components.ConfigGroupLookupsLookup{
                        DeployedVersion: criblcontrolplanesdkgo.String("<value>"),
                        File: "<value>",
                        Version: criblcontrolplanesdkgo.String("<value>"),
                    },
                },
            },
        },
        MaxWorkerAge: criblcontrolplanesdkgo.String("<value>"),
        Name: criblcontrolplanesdkgo.String("<value>"),
        OnPrem: criblcontrolplanesdkgo.Bool(false),
        Provisioned: criblcontrolplanesdkgo.Bool(true),
        Streamtags: []string{
            "<value 1>",
            "<value 2>",
        },
        Tags: criblcontrolplanesdkgo.String("<value>"),
        Type: components.ConfigGroupTypeLakeAccess.ToPointer(),
        UpgradeVersion: criblcontrolplanesdkgo.String("<value>"),
        WorkerCount: criblcontrolplanesdkgo.Float64(1557.82),
        WorkerRemoteAccess: criblcontrolplanesdkgo.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `product`                                                              | [components.ProductsCore](../../models/components/productscore.md)     | :heavy_check_mark:                                                     | Name of the Cribl product to get the Worker Groups or Edge Fleets for. |
| `id`                                                                   | *string*                                                               | :heavy_check_mark:                                                     | The <code>id</code> of the Worker Group or Edge Fleet to update.       |
| `configGroup`                                                          | [components.ConfigGroup](../../models/components/configgroup.md)       | :heavy_check_mark:                                                     | ConfigGroup object                                                     |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.UpdateConfigGroupByProductAndIDResponse](../../models/operations/updateconfiggroupbyproductandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete the specified Worker Group or Edge Fleet.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteConfigGroupByProductAndId" method="delete" path="/products/{product}/groups/{id}" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Groups.Delete(ctx, components.ProductsCoreEdge, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `product`                                                              | [components.ProductsCore](../../models/components/productscore.md)     | :heavy_check_mark:                                                     | Name of the Cribl product to get the Worker Groups or Edge Fleets for. |
| `id`                                                                   | *string*                                                               | :heavy_check_mark:                                                     | The <code>id</code> of the Worker Group or Edge Fleet to delete.       |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.DeleteConfigGroupByProductAndIDResponse](../../models/operations/deleteconfiggroupbyproductandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Deploy

Deploy commits to the specified Worker Group or Edge Fleet.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateConfigGroupDeployByProductAndId" method="patch" path="/products/{product}/groups/{id}/deploy" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Groups.Deploy(ctx, components.ProductsCoreStream, "<id>", components.DeployRequest{
        Lookups: []components.DeployRequestLookups{
            components.DeployRequestLookups{
                Context: "<value>",
                Lookups: []components.DeployRequestLookupsLookup{},
            },
        },
        Version: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                           | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `ctx`                                                                               | [context.Context](https://pkg.go.dev/context#Context)                               | :heavy_check_mark:                                                                  | The context to use for the request.                                                 |
| `product`                                                                           | [components.ProductsCore](../../models/components/productscore.md)                  | :heavy_check_mark:                                                                  | Name of the Cribl product to get the Worker Groups or Edge Fleets for.              |
| `id`                                                                                | *string*                                                                            | :heavy_check_mark:                                                                  | The <code>id</code> of the target Worker Group or Edge Fleet for commit deployment. |
| `deployRequest`                                                                     | [components.DeployRequest](../../models/components/deployrequest.md)                | :heavy_check_mark:                                                                  | DeployRequest object                                                                |
| `opts`                                                                              | [][operations.Option](../../models/operations/option.md)                            | :heavy_minus_sign:                                                                  | The options for this request.                                                       |

### Response

**[*operations.UpdateConfigGroupDeployByProductAndIDResponse](../../models/operations/updateconfiggroupdeploybyproductandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |