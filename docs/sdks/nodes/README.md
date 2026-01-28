# Nodes

## Overview

### Available Operations

* [~~Count~~](#count) - Get a count of Worker and Edge Nodes :warning: **Deprecated**
* [~~List~~](#list) - Get detailed metadata for Worker and Edge Nodes :warning: **Deprecated**

## ~~Count~~

Get a count of all Worker and Edge Nodes. Deprecated. Use /products/{product}/summary/workers instead.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSummaryWorkers" method="get" path="/master/summary/workers" -->
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

    res, err := s.Nodes.Count(ctx, criblcontrolplanesdkgo.Pointer("<value>"))
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
| `filterExp`                                                                | **string*                                                                  | :heavy_minus_sign:                                                         | Filter expression to evaluate against Nodes for inclusion in the response. |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.GetSummaryWorkersResponse](../../models/operations/getsummaryworkersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ~~List~~

Get detailed metadata for Worker and Edge Nodes. Deprecated. Use /products/{product}/workers instead.

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="getWorkers" method="get" path="/master/workers" -->
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

    res, err := s.Nodes.List(ctx, operations.GetWorkersRequest{
        FilterExp: criblcontrolplanesdkgo.Pointer("<value>"),
        SortExp: criblcontrolplanesdkgo.Pointer("<value>"),
        Filter: criblcontrolplanesdkgo.Pointer("<value>"),
        Sort: criblcontrolplanesdkgo.Pointer("<value>"),
        Limit: criblcontrolplanesdkgo.Pointer[int64](402753),
        Offset: criblcontrolplanesdkgo.Pointer[int64](848752),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedMasterWorkerEntry != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetWorkersRequest](../../models/operations/getworkersrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.GetWorkersResponse](../../models/operations/getworkersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |