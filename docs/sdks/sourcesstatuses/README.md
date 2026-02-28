# Sources.Statuses

## Overview

### Available Operations

* [Get](#get) - Get the status of a Source
* [List](#list) - List the status of all Sources

## Get

Get the status and optional metrics for the specified Source.

### Example Usage

<!-- UsageSnippet language="go" operationID="getInputStatusById" method="get" path="/system/status/inputs/{id}" -->
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

    res, err := s.Sources.Statuses.Get(ctx, "<id>", nil, nil)
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
| `metrics`                                                                                                                    | **bool*                                                                                                                      | :heavy_minus_sign:                                                                                                           | Set to <code>true</code> to include metrics for each Source. Otherwise, <code>false</code> (default).                        |
| `type_`                                                                                                                      | **bool*                                                                                                                      | :heavy_minus_sign:                                                                                                           | Set to <code>true</code> to prefix the Source <code>id</code> with the Source type. Otherwise, <code>false</code> (default). |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.GetInputStatusByIDResponse](../../models/operations/getinputstatusbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

List status information and optional metrics for all configured Sources in the Worker Group or Edge Fleet.

### Example Usage

<!-- UsageSnippet language="go" operationID="getInputStatus" method="get" path="/system/status/inputs" -->
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

    res, err := s.Sources.Statuses.List(ctx, nil, nil)
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
| `metrics`                                                                                                                    | **bool*                                                                                                                      | :heavy_minus_sign:                                                                                                           | Set to <code>true</code> to include metrics for each Source. Otherwise, <code>false</code> (default).                        |
| `type_`                                                                                                                      | **bool*                                                                                                                      | :heavy_minus_sign:                                                                                                           | Set to <code>true</code> to prefix the Source <code>id</code> with the Source type. Otherwise, <code>false</code> (default). |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.GetInputStatusResponse](../../models/operations/getinputstatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |