# Groups

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

    res, err := s.Groups.List(ctx, components.ProductsCoreEdge, criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedConfigGroup != nil {
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

    res, err := s.Groups.Create(ctx, components.ProductsCoreEdge, components.GroupCreateRequest{
        Cloud: &components.ConfigGroupCloud{
            Provider: components.CloudProviderAws.ToPointer(),
            Region: "us-west-2",
        },
        DeployingWorkerCount: criblcontrolplanesdkgo.Pointer[float64](5631.58),
        Description: criblcontrolplanesdkgo.Pointer("ack resort boohoo"),
        EstimatedIngestRate: components.EstimatedIngestRateOptionsConfigGroupRate24MbPerSec.ToPointer(),
        Git: &components.GitTypeConfigGroup{
            Commit: criblcontrolplanesdkgo.Pointer("<value>"),
            LocalChanges: criblcontrolplanesdkgo.Pointer[float64](2413.01),
            Log: []components.Commit{
                components.Commit{
                    AuthorEmail: criblcontrolplanesdkgo.Pointer("<value>"),
                    AuthorName: criblcontrolplanesdkgo.Pointer("<value>"),
                    Date: "2024-04-03",
                    Hash: "<value>",
                    Message: "<value>",
                    Short: "<value>",
                },
            },
        },
        ID: "goatCloudIanWg",
        IncompatibleWorkerCount: criblcontrolplanesdkgo.Pointer[float64](7174.43),
        Inherits: criblcontrolplanesdkgo.Pointer("<value>"),
        IsFleet: criblcontrolplanesdkgo.Pointer(false),
        IsSearch: criblcontrolplanesdkgo.Pointer(false),
        LookupDeployments: []components.ConfigGroupLookups{
            components.ConfigGroupLookups{
                Context: "<value>",
                Lookups: []components.ConfigGroupLookupsLookup{
                    components.ConfigGroupLookupsLookup{
                        DeployedVersion: criblcontrolplanesdkgo.Pointer("<value>"),
                        File: "<value>",
                        Version: criblcontrolplanesdkgo.Pointer("<value>"),
                    },
                },
            },
        },
        MaxWorkerAge: criblcontrolplanesdkgo.Pointer("<value>"),
        Name: criblcontrolplanesdkgo.Pointer("goatCloudIanWg"),
        OnPrem: criblcontrolplanesdkgo.Pointer(false),
        Provisioned: criblcontrolplanesdkgo.Pointer(true),
        SourceGroupID: criblcontrolplanesdkgo.Pointer("<id>"),
        Streamtags: []string{
            "<value 1>",
            "<value 2>",
            "<value 3>",
        },
        Tags: criblcontrolplanesdkgo.Pointer("<value>"),
        Type: components.TypeOptionsConfigGroupLakeAccess.ToPointer(),
        UpgradeVersion: criblcontrolplanesdkgo.Pointer("<value>"),
        WorkerCount: criblcontrolplanesdkgo.Pointer[float64](4980.41),
        WorkerRemoteAccess: criblcontrolplanesdkgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedConfigGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `product`                                                                      | [components.ProductsCore](../../models/components/productscore.md)             | :heavy_check_mark:                                                             | Name of the Cribl product to add the Worker Group or Edge Fleet to.            |
| `groupCreateRequest`                                                           | [components.GroupCreateRequest](../../models/components/groupcreaterequest.md) | :heavy_check_mark:                                                             | GroupCreateRequest object                                                      |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

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

    res, err := s.Groups.Get(ctx, components.ProductsCoreEdge, "<id>", criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedConfigGroup != nil {
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

    res, err := s.Groups.Update(ctx, components.ProductsCoreStream, "<id>", components.ConfigGroup{
        Cloud: &components.ConfigGroupCloud{
            Provider: components.CloudProviderAws.ToPointer(),
            Region: "us-west-2",
        },
        ConfigVersion: criblcontrolplanesdkgo.Pointer("<value>"),
        DeployingWorkerCount: criblcontrolplanesdkgo.Pointer[float64](7786.61),
        Description: criblcontrolplanesdkgo.Pointer("Scaled Worker Group with estimated ingest rate of 4096 (48 MB/s, 21 Worker Processes) for increased capacity"),
        EstimatedIngestRate: components.EstimatedIngestRateOptionsConfigGroupRate48MbPerSec.ToPointer(),
        Git: &components.GitTypeConfigGroup{
            Commit: criblcontrolplanesdkgo.Pointer("<value>"),
            LocalChanges: criblcontrolplanesdkgo.Pointer[float64](776.15),
            Log: []components.Commit{
                components.Commit{
                    AuthorEmail: criblcontrolplanesdkgo.Pointer("<value>"),
                    AuthorName: criblcontrolplanesdkgo.Pointer("<value>"),
                    Date: "2024-09-29",
                    Hash: "<value>",
                    Message: "<value>",
                    Short: "<value>",
                },
            },
        },
        ID: "goatCloudIanWg",
        IncompatibleWorkerCount: criblcontrolplanesdkgo.Pointer[float64](2874.65),
        Inherits: criblcontrolplanesdkgo.Pointer("<value>"),
        IsFleet: criblcontrolplanesdkgo.Pointer(false),
        IsSearch: criblcontrolplanesdkgo.Pointer(false),
        LookupDeployments: []components.ConfigGroupLookups{
            components.ConfigGroupLookups{
                Context: "<value>",
                Lookups: []components.ConfigGroupLookupsLookup{},
            },
        },
        MaxWorkerAge: criblcontrolplanesdkgo.Pointer("<value>"),
        Name: criblcontrolplanesdkgo.Pointer("goatCloudIanWg"),
        OnPrem: criblcontrolplanesdkgo.Pointer(false),
        Provisioned: criblcontrolplanesdkgo.Pointer(true),
        Streamtags: []string{
            "<value 1>",
            "<value 2>",
            "<value 3>",
        },
        Tags: criblcontrolplanesdkgo.Pointer("<value>"),
        Type: components.TypeOptionsConfigGroupLakeAccess.ToPointer(),
        UpgradeVersion: criblcontrolplanesdkgo.Pointer("<value>"),
        WorkerCount: criblcontrolplanesdkgo.Pointer[float64](835.08),
        WorkerRemoteAccess: criblcontrolplanesdkgo.Pointer(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedConfigGroup != nil {
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

    res, err := s.Groups.Delete(ctx, components.ProductsCoreEdge, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedConfigGroup != nil {
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
    if res.CountedConfigGroup != nil {
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