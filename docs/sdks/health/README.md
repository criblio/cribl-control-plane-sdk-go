# Health

## Overview

Actions related to REST server health

### Available Operations

* [Get](#get) - Retrieve health status of the server

## Get

Get the current health status of the server.

### Example Usage

<!-- UsageSnippet language="go" operationID="getHealth" method="get" path="/health" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
    )

    res, err := s.Health.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthServerStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetHealthResponse](../../models/operations/gethealthresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.HealthServerStatusError | 420                               | application/json                  |
| apierrors.Error                   | 500                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |