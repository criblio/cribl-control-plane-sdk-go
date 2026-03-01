# Packs.Destinations.Statuses

## Overview

### Available Operations

* [Get](#get) - Get the status of a Destination within a Pack
* [List](#list) - List the status of all Destinations within a Pack

## Get

Get the status and optional metrics for the specified Destination within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="getOutputStatusSystemOutputsByPackAndId" method="get" path="/p/{pack}/system/status/outputs/{id}" -->
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

    res, err := s.Packs.Destinations.Statuses.Get(ctx, "<id>", "<value>", nil, nil)
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
| `id`                                                                                                                                   | *string*                                                                                                                               | :heavy_check_mark:                                                                                                                     | The <code>id</code> of the Destination to get the status for.                                                                          |
| `pack`                                                                                                                                 | *string*                                                                                                                               | :heavy_check_mark:                                                                                                                     | The <code>id</code> of the Pack to get.                                                                                                |
| `metrics`                                                                                                                              | **bool*                                                                                                                                | :heavy_minus_sign:                                                                                                                     | Set to <code>true</code> to include metrics for each Destination. Otherwise, <code>false</code> (default).                             |
| `type_`                                                                                                                                | **bool*                                                                                                                                | :heavy_minus_sign:                                                                                                                     | Set to <code>true</code> to prefix the Destination <code>id</code> with the Destination type. Otherwise, <code>false</code> (default). |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.GetOutputStatusSystemOutputsByPackAndIDResponse](../../models/operations/getoutputstatussystemoutputsbypackandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

List status information and optional health metrics for all configured Destinations in the Worker Group or Edge Fleet within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="getOutputStatusSystemOutputsByPack" method="get" path="/p/{pack}/system/status/outputs" -->
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

    res, err := s.Packs.Destinations.Statuses.List(ctx, "<value>", nil, nil)
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
| `pack`                                                                                                                                 | *string*                                                                                                                               | :heavy_check_mark:                                                                                                                     | The <code>id</code> of the Pack to list.                                                                                               |
| `metrics`                                                                                                                              | **bool*                                                                                                                                | :heavy_minus_sign:                                                                                                                     | Set to <code>true</code> to include metrics for each Destination. Otherwise, <code>false</code> (default).                             |
| `type_`                                                                                                                                | **bool*                                                                                                                                | :heavy_minus_sign:                                                                                                                     | Set to <code>true</code> to prefix the Destination <code>id</code> with the Destination type. Otherwise, <code>false</code> (default). |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.GetOutputStatusSystemOutputsByPackResponse](../../models/operations/getoutputstatussystemoutputsbypackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |