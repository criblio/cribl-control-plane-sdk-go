# Packs.Destinations.Samples

## Overview

### Available Operations

* [Get](#get) - Get sample event data for a Destination within a Pack
* [Create](#create) - Send sample event data to a Destination within a Pack

## Get

Get sample event data for the specified Destination to validate the configuration or test connectivity within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="getOutputSystemSamplesByPackAndId" method="get" path="/p/{pack}/system/outputs/{id}/samples" -->
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

    res, err := s.Packs.Destinations.Samples.Get(ctx, "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedOutputSamplesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `id`                                                                 | *string*                                                             | :heavy_check_mark:                                                   | The <code>id</code> of the Destination to get sample event data for. |
| `pack`                                                               | *string*                                                             | :heavy_check_mark:                                                   | The <code>id</code> of the Pack to get.                              |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.GetOutputSystemSamplesByPackAndIDResponse](../../models/operations/getoutputsystemsamplesbypackandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Send sample event data to the specified Destination to validate the configuration or test connectivity within the specified Pack.

### Example Usage: OutputTestExamplesMultipleEvents

<!-- UsageSnippet language="go" operationID="createOutputSystemTestByPackAndId" method="post" path="/p/{pack}/system/outputs/{id}/test" example="OutputTestExamplesMultipleEvents" -->
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

    res, err := s.Packs.Destinations.Samples.Create(ctx, "<id>", "<value>", components.OutputTestRequest{
        Events: []map[string]any{
            map[string]any{
                "_raw": "Test event 1",
                "source": "test",
                "sourcetype": "manual",
            },
            map[string]any{
                "_raw": "Test event 2",
                "source": "test",
                "sourcetype": "manual",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedOutputTestResponse != nil {
        // handle response
    }
}
```
### Example Usage: OutputTestExamplesSingleEvent

<!-- UsageSnippet language="go" operationID="createOutputSystemTestByPackAndId" method="post" path="/p/{pack}/system/outputs/{id}/test" example="OutputTestExamplesSingleEvent" -->
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

    res, err := s.Packs.Destinations.Samples.Create(ctx, "<id>", "<value>", components.OutputTestRequest{
        Events: []map[string]any{
            map[string]any{
                "_raw": "This is a test event",
                "source": "test",
                "sourcetype": "manual",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedOutputTestResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `id`                                                                         | *string*                                                                     | :heavy_check_mark:                                                           | The <code>id</code> of the Destination to send sample event data to.         |
| `pack`                                                                       | *string*                                                                     | :heavy_check_mark:                                                           | The <code>id</code> of the Pack to create.                                   |
| `outputTestRequest`                                                          | [components.OutputTestRequest](../../models/components/outputtestrequest.md) | :heavy_check_mark:                                                           | OutputTestRequest object                                                     |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.CreateOutputSystemTestByPackAndIDResponse](../../models/operations/createoutputsystemtestbypackandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |