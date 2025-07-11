# FileSampler
(*FileSampler*)

## Overview

Actions related to FileSampler

### Available Operations

* [GetEdgeFileSample](#getedgefilesample) - Get some number of bytes from the file at the given path

## GetEdgeFileSample

Get some number of bytes from the file at the given path

### Example Usage

```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
    )

    res, err := s.FileSampler.GetEdgeFileSample(ctx, "/usr/ports", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `path`                                                                             | *string*                                                                           | :heavy_check_mark:                                                                 | The path to the file to sample                                                     |
| `bytesRequested`                                                                   | **float64*                                                                         | :heavy_minus_sign:                                                                 | The number of bytes to return;   this value could be constrained by system limits. |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetEdgeFileSampleResponse](../../models/operations/getedgefilesampleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |