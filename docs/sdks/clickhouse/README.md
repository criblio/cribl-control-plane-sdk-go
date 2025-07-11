# ClickHouse
(*ClickHouse*)

## Overview

Actions related to ClickHouse

### Available Operations

* [CreateOutputClickHouseDescribeTable](#createoutputclickhousedescribetable) - Retrieve the description of the configured ClickHouse table

## CreateOutputClickHouseDescribeTable

Retrieve the description of the configured ClickHouse table

### Example Usage

```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
    )

    res, err := s.ClickHouse.CreateOutputClickHouseDescribeTable(ctx, components.CHOutConfig{
        AsyncInserts: true,
        Database: "<value>",
        FlushPeriodSec: 6320.13,
        Format: components.FormatEnumJSONEachRow,
        LoadBalanced: true,
        MappingType: components.MappingTypeCustom,
        TableName: "<value>",
        URL: "https://blue-fundraising.name",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [components.CHOutConfig](../../models/components/choutconfig.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.CreateOutputClickHouseDescribeTableResponse](../../models/operations/createoutputclickhousedescribetableresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |