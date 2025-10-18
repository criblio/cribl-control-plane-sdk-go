# Results
(*Search.Jobs.Results*)

## Overview

### Available Operations

* [Get](#get) - List search results, when lower/upper bound is provided, offset is relative to the time range.
* [Poll](#poll) - List search results

## Get

List search results, when lower/upper bound is provided, offset is relative to the time range.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSearchJobsResultsById" method="get" path="/search/jobs/{id}/results" -->
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

    res, err := s.Search.Jobs.Results.Get(ctx, operations.GetSearchJobsResultsByIDRequest{
        ID: "<id>",
        Limit: criblcontrolplanesdkgo.Pointer[float64](9933.5),
        Offset: criblcontrolplanesdkgo.Pointer[float64](9757.07),
        LowerBound: criblcontrolplanesdkgo.Pointer[float64](6377.21),
        UpperBound: criblcontrolplanesdkgo.Pointer[float64](1238.19),
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetSearchJobsResultsByIDRequest](../../models/operations/getsearchjobsresultsbyidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetSearchJobsResultsByIDResponse](../../models/operations/getsearchjobsresultsbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Poll

List search results

### Example Usage

<!-- UsageSnippet language="go" operationID="getSearchJobsResultsPollById" method="get" path="/search/jobs/{id}/results-poll" -->
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

    res, err := s.Search.Jobs.Results.Poll(ctx, operations.GetSearchJobsResultsPollByIDRequest{
        ID: "<id>",
        Limit: criblcontrolplanesdkgo.Pointer[float64](7966.81),
        Offset: criblcontrolplanesdkgo.Pointer[float64](8591.71),
        LowerBound: criblcontrolplanesdkgo.Pointer[float64](9350.22),
        UpperBound: criblcontrolplanesdkgo.Pointer[float64](9654.15),
        LastJobStatus: criblcontrolplanesdkgo.Pointer("<value>"),
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetSearchJobsResultsPollByIDRequest](../../models/operations/getsearchjobsresultspollbyidrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.GetSearchJobsResultsPollByIDResponse](../../models/operations/getsearchjobsresultspollbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |