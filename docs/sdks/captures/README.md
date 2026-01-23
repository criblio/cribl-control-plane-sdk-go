# System.Captures

## Overview

### Available Operations

* [Get](#get) - Capture live incoming data

## Get

Capture live incoming data

### Example Usage

<!-- UsageSnippet language="go" operationID="createSystemCapture" method="post" path="/system/capture" -->
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

    res, err := s.System.Captures.Get(ctx, components.CaptureParams{
        Duration: 5,
        Filter: "true",
        Level: components.CaptureLevelBeforePreProcessingPipeline,
        MaxEvents: 100,
        StepDuration: criblcontrolplanesdkgo.Pointer[int64](571732),
        WorkerID: criblcontrolplanesdkgo.Pointer("<id>"),
        WorkerThreshold: criblcontrolplanesdkgo.Pointer[int64](609412),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CapturedEvent != nil {
        for res.CapturedEvent.Next() {
            event, _ := res.CapturedEvent.Value()
            log.Print(event)
            // Handle the event
	      }
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [components.CaptureParams](../../models/components/captureparams.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.CreateSystemCaptureResponse](../../models/operations/createsystemcaptureresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |