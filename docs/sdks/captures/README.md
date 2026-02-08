# System.Captures

## Overview

### Available Operations

* [Create](#create) - Capture live incoming data

## Create

Initiate a live data capture from Cribl Workers.Returns a stream of captured events in NDJSON format that match the parameters specified in the request body.

### Example Usage: CaptureExamplesComplexFilter

<!-- UsageSnippet language="go" operationID="createSystemCapture" method="post" path="/system/capture" example="CaptureExamplesComplexFilter" -->
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

    res, err := s.System.Captures.Create(ctx, components.CaptureParams{
        Duration: 15,
        Filter: "__inputId.startsWith(\"http:\") && status >= 400 && status < 500",
        Level: components.CaptureLevelBeforeRoutes,
        MaxEvents: 500,
        StepDuration: criblcontrolplanesdkgo.Pointer[int64](377776),
        WorkerID: criblcontrolplanesdkgo.Pointer("<id>"),
        WorkerThreshold: criblcontrolplanesdkgo.Pointer[int64](429562),
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
### Example Usage: CaptureExamplesCompoundAndExpression

<!-- UsageSnippet language="go" operationID="createSystemCapture" method="post" path="/system/capture" example="CaptureExamplesCompoundAndExpression" -->
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

    res, err := s.System.Captures.Create(ctx, components.CaptureParams{
        Duration: 5,
        Filter: "sourcetype===\"pan:traffic\" && src_zone===\"trusted\"",
        Level: components.CaptureLevelBeforePreProcessingPipeline,
        MaxEvents: 100,
        StepDuration: criblcontrolplanesdkgo.Pointer[int64](994184),
        WorkerID: criblcontrolplanesdkgo.Pointer("<id>"),
        WorkerThreshold: criblcontrolplanesdkgo.Pointer[int64](771620),
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
### Example Usage: CaptureExamplesNestedFieldAccess

<!-- UsageSnippet language="go" operationID="createSystemCapture" method="post" path="/system/capture" example="CaptureExamplesNestedFieldAccess" -->
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

    res, err := s.System.Captures.Create(ctx, components.CaptureParams{
        Duration: 5,
        Filter: "sourcetype===\"pan:traffic\" && dest_geoip.country.iso_code === \"US\"",
        Level: components.CaptureLevelBeforePreProcessingPipeline,
        MaxEvents: 100,
        StepDuration: criblcontrolplanesdkgo.Pointer[int64](882563),
        WorkerID: criblcontrolplanesdkgo.Pointer("<id>"),
        WorkerThreshold: criblcontrolplanesdkgo.Pointer[int64](392678),
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
### Example Usage: CaptureExamplesSimpleExpression

<!-- UsageSnippet language="go" operationID="createSystemCapture" method="post" path="/system/capture" example="CaptureExamplesSimpleExpression" -->
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

    res, err := s.System.Captures.Create(ctx, components.CaptureParams{
        Duration: 5,
        Filter: "sourcetype===\"pan:traffic\"",
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