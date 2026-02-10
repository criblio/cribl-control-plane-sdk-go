# Packs.Sources.Statuses

## Overview

### Available Operations

* [Get](#get) - Get the status of a Source within a Pack
* [List](#list) - List the status of all Sources within a Pack

## Get

Get the status and optional metrics for the specified Source within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="getInputStatusSystemInputsByPackAndId" method="get" path="/p/{pack}/system/status/inputs/{id}" -->
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

    res, err := s.Packs.Sources.Statuses.Get(ctx, "<id>", "<value>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `id`                                                                                                                         | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The <code>id</code> of the Source to get the status for.                                                                     |
| `pack`                                                                                                                       | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The <code>id</code> of the Pack to get.                                                                                      |
| `metrics`                                                                                                                    | **bool*                                                                                                                      | :heavy_minus_sign:                                                                                                           | Set to true <code>true</code> to include metrics for each Source. Otherwise, <code>false</code> (default).                   |
| `type_`                                                                                                                      | **bool*                                                                                                                      | :heavy_minus_sign:                                                                                                           | Set to <code>true</code> to prefix the Source <code>id</code> with the Source type. Otherwise, <code>false</code> (default). |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.GetInputStatusSystemInputsByPackAndIDResponse](../../models/operations/getinputstatussysteminputsbypackandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Get a list of status information and optional metrics for all configured Sources in the Worker Group or Edge Fleet within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="getInputStatusSystemInputsByPack" method="get" path="/p/{pack}/system/status/inputs" -->
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

    res, err := s.Packs.Sources.Statuses.List(ctx, "<value>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `pack`                                                                                                                       | *string*                                                                                                                     | :heavy_check_mark:                                                                                                           | The <code>id</code> of the Pack to list.                                                                                     |
| `metrics`                                                                                                                    | **bool*                                                                                                                      | :heavy_minus_sign:                                                                                                           | Set to true <code>true</code> to include metrics for each Source. Otherwise, <code>false</code> (default).                   |
| `type_`                                                                                                                      | **bool*                                                                                                                      | :heavy_minus_sign:                                                                                                           | Set to <code>true</code> to prefix the Source <code>id</code> with the Source type. Otherwise, <code>false</code> (default). |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.GetInputStatusSystemInputsByPackResponse](../../models/operations/getinputstatussysteminputsbypackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |