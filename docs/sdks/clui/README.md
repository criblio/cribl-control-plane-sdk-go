# Clui
(*Clui*)

## Overview

Actions related to CLUI

### Available Operations

* [GetClui](#getclui) - Get CLUI search results

## GetClui

Get CLUI search results

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

    res, err := s.Clui.GetClui(ctx, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                 | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ctx`                                                     | [context.Context](https://pkg.go.dev/context#Context)     | :heavy_check_mark:                                        | The context to use for the request.                       |
| `query`                                                   | *string*                                                  | :heavy_check_mark:                                        | Search query                                              |
| `context_`                                                | [*operations.Context](../../models/operations/context.md) | :heavy_minus_sign:                                        | Search query context                                      |
| `opts`                                                    | [][operations.Option](../../models/operations/option.md)  | :heavy_minus_sign:                                        | The options for this request.                             |

### Response

**[*operations.GetCluiResponse](../../models/operations/getcluiresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |