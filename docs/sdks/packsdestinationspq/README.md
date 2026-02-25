# Packs.Destinations.Pq

## Overview

### Available Operations

* [Clear](#clear) - Clear the persistent queue for a Destination within a Pack
* [Get](#get) - Get information about the latest job to clear the persistent queue for a Destination within a Pack

## Clear

Clear the persistent queue (PQ) for the specified Destination within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteOutputSystemPqByPackAndId" method="delete" path="/p/{pack}/system/outputs/{id}/pq" -->
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

    res, err := s.Packs.Destinations.Pq.Clear(ctx, "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |
| `id`                                                        | *string*                                                    | :heavy_check_mark:                                          | The <code>id</code> of the Destination to clear the PQ for. |
| `pack`                                                      | *string*                                                    | :heavy_check_mark:                                          | The <code>id</code> of the Pack to clear.                   |
| `opts`                                                      | [][operations.Option](../../models/operations/option.md)    | :heavy_minus_sign:                                          | The options for this request.                               |

### Response

**[*operations.DeleteOutputSystemPqByPackAndIDResponse](../../models/operations/deleteoutputsystempqbypackandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get information about the latest job to clear the persistent queue (PQ) for the specified Destination within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="getOutputSystemPqByPackAndId" method="get" path="/p/{pack}/system/outputs/{id}/pq" -->
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

    res, err := s.Packs.Destinations.Pq.Get(ctx, "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedJobInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                             | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ctx`                                                                 | [context.Context](https://pkg.go.dev/context#Context)                 | :heavy_check_mark:                                                    | The context to use for the request.                                   |
| `id`                                                                  | *string*                                                              | :heavy_check_mark:                                                    | The <code>id</code> of the Destination to get PQ job information for. |
| `pack`                                                                | *string*                                                              | :heavy_check_mark:                                                    | The <code>id</code> of the Pack to get.                               |
| `opts`                                                                | [][operations.Option](../../models/operations/option.md)              | :heavy_minus_sign:                                                    | The options for this request.                                         |

### Response

**[*operations.GetOutputSystemPqByPackAndIDResponse](../../models/operations/getoutputsystempqbypackandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |