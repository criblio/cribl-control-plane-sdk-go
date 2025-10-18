# Queries
(*Search.Queries*)

## Overview

### Available Operations

* [Create](#create) - Runs the query and returns the results

## Create

Runs the query and returns the results

### Example Usage

<!-- UsageSnippet language="go" operationID="getSearchQuery" method="get" path="/search/query" -->
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

    res, err := s.Search.Queries.Create(ctx, operations.GetSearchQueryRequest{
        QueryID: criblcontrolplanesdkgo.Pointer("<id>"),
        JobID: criblcontrolplanesdkgo.Pointer("<id>"),
        Query: criblcontrolplanesdkgo.Pointer("<value>"),
        Earliest: criblcontrolplanesdkgo.Pointer[float64](8608.12),
        Latest: criblcontrolplanesdkgo.Pointer[float64](3938.83),
        SampleRate: criblcontrolplanesdkgo.Pointer[float64](2315.98),
        Force: criblcontrolplanesdkgo.Pointer(true),
        Offset: criblcontrolplanesdkgo.Pointer[float64](6530.23),
        Limit: criblcontrolplanesdkgo.Pointer[float64](49.47),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchJobResults != nil {
        for res.SearchJobResults.Next() {
            event, _ := res.SearchJobResults.Value()
            log.Print(event)
            // Handle the event
	      }
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetSearchQueryRequest](../../models/operations/getsearchqueryrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetSearchQueryResponse](../../models/operations/getsearchqueryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |