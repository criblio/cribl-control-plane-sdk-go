# Diag
(*Diag*)

## Overview

Actions related to diagnostics

### Available Operations

* [GetHealthInfo](#gethealthinfo) - Provides health info for REST server

## GetHealthInfo

Provides health info for REST server

### Example Usage

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

    res, err := s.Diag.GetHealthInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.HealthStatus != nil {
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

**[*operations.GetHealthInfoResponse](../../models/operations/gethealthinforesponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| apierrors.HealthStatusError | 420                         | application/json            |
| apierrors.APIError          | 4XX, 5XX                    | \*/\*                       |