# Destinations.Statuses

## Overview

### Available Operations

* [List](#list) - List the status of all Destinations
* [Get](#get) - Get the status of a Destination

## List

List status information and optional metrics for all configured Destinations in the Worker Group or Edge Fleet.

### Example Usage: OutputStatusResponseExamplesGreenDestination

<!-- UsageSnippet language="go" operationID="getOutputStatus" method="get" path="/system/status/outputs" example="OutputStatusResponseExamplesGreenDestination" -->
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

    res, err := s.Destinations.Statuses.List(ctx, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedOutputStatus != nil {
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
### Example Usage: OutputStatusResponseExamplesYellowDestination

<!-- UsageSnippet language="go" operationID="getOutputStatus" method="get" path="/system/status/outputs" example="OutputStatusResponseExamplesYellowDestination" -->
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

    res, err := s.Destinations.Statuses.List(ctx, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedOutputStatus != nil {
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

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `metrics`                                                                                                                                          | `*bool`                                                                                                                                            | :heavy_minus_sign:                                                                                                                                 | Set to <code>true</code> to include metrics for each Destination. Otherwise, <code>false</code> (default).                                         |
| `type_`                                                                                                                                            | `*bool`                                                                                                                                            | :heavy_minus_sign:                                                                                                                                 | Set to <code>true</code> to prefix the Destination <code>id</code> with the Destination type. Otherwise, <code>false</code> (default).             |
| `offset`                                                                                                                                           | `*int64`                                                                                                                                           | :heavy_minus_sign:                                                                                                                                 | Starting point from which to retrieve results for this request. Use with <code>limit</code> to paginate the response into manageable batches.      |
| `limit`                                                                                                                                            | `*int64`                                                                                                                                           | :heavy_minus_sign:                                                                                                                                 | Maximum number of items to return in the response for this request. Use with <code>offset</code> to paginate the response into manageable batches. |
| `opts`                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.GetOutputStatusResponse](../../models/operations/getoutputstatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 401                | application/json   |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get the status and optional metrics for the specified Destination.

### Example Usage: OutputStatusResponseExamplesGreenDestination

<!-- UsageSnippet language="go" operationID="getOutputStatusById" method="get" path="/system/status/outputs/{id}" example="OutputStatusResponseExamplesGreenDestination" -->
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

    res, err := s.Destinations.Statuses.Get(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedOutputStatus != nil {
        // handle response
    }
}
```
### Example Usage: OutputStatusResponseExamplesYellowDestination

<!-- UsageSnippet language="go" operationID="getOutputStatusById" method="get" path="/system/status/outputs/{id}" example="OutputStatusResponseExamplesYellowDestination" -->
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

    res, err := s.Destinations.Statuses.Get(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedOutputStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `id`                                                                                                                                   | `string`                                                                                                                               | :heavy_check_mark:                                                                                                                     | The <code>id</code> of the Destination to get the status for.                                                                          |
| `metrics`                                                                                                                              | `*bool`                                                                                                                                | :heavy_minus_sign:                                                                                                                     | Set to <code>true</code> to include metrics for each Destination. Otherwise, <code>false</code> (default).                             |
| `type_`                                                                                                                                | `*bool`                                                                                                                                | :heavy_minus_sign:                                                                                                                     | Set to <code>true</code> to prefix the Destination <code>id</code> with the Destination type. Otherwise, <code>false</code> (default). |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.GetOutputStatusByIDResponse](../../models/operations/getoutputstatusbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 401                | application/json   |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |