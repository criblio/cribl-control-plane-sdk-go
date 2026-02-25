# Nodes

## Overview

### Available Operations

* [Count](#count) - Get a count of Worker or Edge Nodes
* [Get](#get) - Get detailed metadata for a Worker or Edge Node
* [List](#list) - Get detailed metadata for Worker or Edge Nodes
* [Restart](#restart) - Restart Worker or Edge Nodes

## Count

Get a count of all Worker or Edge Nodes for the specified Cribl product.

### Example Usage

<!-- UsageSnippet language="go" operationID="getProductsSummaryWorkersByProduct" method="get" path="/products/{product}/summary/workers" -->
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

    res, err := s.Nodes.Count(ctx, components.ProductsBaseEdge, criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedNumber != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `product`                                                                  | [components.ProductsBase](../../models/components/productsbase.md)         | :heavy_check_mark:                                                         | Name of the Cribl product to get the count of Worker or Edge Nodes for.    |
| `filterExp`                                                                | **string*                                                                  | :heavy_minus_sign:                                                         | Filter expression to evaluate against Nodes for inclusion in the response. |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.GetProductsSummaryWorkersByProductResponse](../../models/operations/getproductssummaryworkersbyproductresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get detailed metadata for the specified Worker or Edge Node for the specified Cribl product.

### Example Usage

<!-- UsageSnippet language="go" operationID="getProductsWorkersByProductAndId" method="get" path="/products/{product}/workers/{id}" -->
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

    res, err := s.Nodes.Get(ctx, components.ProductsBaseStream, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedMasterWorkerEntry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `product`                                                          | [components.ProductsBase](../../models/components/productsbase.md) | :heavy_check_mark:                                                 | Name of the Cribl product that contains the Node.                  |
| `id`                                                               | *string*                                                           | :heavy_check_mark:                                                 | The <code>id</code> of the Node to get the metadata for.           |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.GetProductsWorkersByProductAndIDResponse](../../models/operations/getproductsworkersbyproductandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Get detailed metadata for Worker or Edge Nodes for the specified Cribl product.

### Example Usage

<!-- UsageSnippet language="go" operationID="getProductsWorkersByProduct" method="get" path="/products/{product}/workers" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
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

    res, err := s.Nodes.List(ctx, operations.GetProductsWorkersByProductRequest{
        Product: components.ProductsBaseStream,
        FilterExp: criblcontrolplanesdkgo.Pointer("<value>"),
        SortExp: criblcontrolplanesdkgo.Pointer("<value>"),
        Filter: criblcontrolplanesdkgo.Pointer("<value>"),
        Sort: criblcontrolplanesdkgo.Pointer("<value>"),
        Limit: criblcontrolplanesdkgo.Pointer[int64](881129),
        Offset: criblcontrolplanesdkgo.Pointer[int64](990978),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedMasterWorkerEntry != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetProductsWorkersByProductRequest](../../models/operations/getproductsworkersbyproductrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.GetProductsWorkersByProductResponse](../../models/operations/getproductsworkersbyproductresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Restart

Restart all Worker or Edge Nodes for the specified Cribl product.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateProductsWorkersRestartByProduct" method="patch" path="/products/{product}/workers/restart" example="RestartWorkersExamplesRestartWorkers" -->
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

    res, err := s.Nodes.Restart(ctx, components.ProductsBaseEdge, components.RestartRequest{
        Guids: []string{
            "guid-12345678-abcd-1234-abcd-123456789abc",
            "guid-87654321-dcba-4321-dcba-cba987654321",
            "guid-11111111-2222-3333-4444-555555555555",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedRestartResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                 | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ctx`                                                                     | [context.Context](https://pkg.go.dev/context#Context)                     | :heavy_check_mark:                                                        | The context to use for the request.                                       |
| `product`                                                                 | [components.ProductsBase](../../models/components/productsbase.md)        | :heavy_check_mark:                                                        | Name of the Cribl product whose Worker or Edge Nodes you want to restart. |
| `restartRequest`                                                          | [components.RestartRequest](../../models/components/restartrequest.md)    | :heavy_check_mark:                                                        | RestartRequest object                                                     |
| `opts`                                                                    | [][operations.Option](../../models/operations/option.md)                  | :heavy_minus_sign:                                                        | The options for this request.                                             |

### Response

**[*operations.UpdateProductsWorkersRestartByProductResponse](../../models/operations/updateproductsworkersrestartbyproductresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |