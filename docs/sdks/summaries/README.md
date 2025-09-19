# Summaries
(*Nodes.Summaries*)

## Overview

### Available Operations

* [Get](#get) - Get a summary of the Distributed deployment

## Get

Get a summary of the Distributed deployment. The response includes counts of Worker Groups, Edge Fleets, Pipelines, Routes, Sources, Destinations, and Worker and Edge Nodes, as well as statistics for the Worker and Edge Nodes.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSummary" method="get" path="/master/summary" -->
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
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Nodes.Summaries.Get(ctx, components.WorkerTypesWorker.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `mode`                                                                                                                           | [*components.WorkerTypes](../../models/components/workertypes.md)                                                                | :heavy_minus_sign:                                                                                                               | Filter for limiting the response by Cribl product: Cribl Stream (<code>worker</code>) or Cribl Edge (<code>managed-edge</code>). |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.GetSummaryResponse](../../models/operations/getsummaryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |