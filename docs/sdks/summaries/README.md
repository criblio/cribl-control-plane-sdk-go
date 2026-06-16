# Nodes.Summaries

## Overview

### Available Operations

* [Get](#get) - Get a summary of the deployment for a specific product.

## Get

Get a summary of the deployment for the specified Cribl product (Stream or Edge).<br/><br/>The summary includes a count of Worker Groups or Edge Fleets and resources  such as Pipelines, Routes, Sources, and Destinations. For Distributed deployments,  it also includes a count and statistics for Worker or Edge Nodes.

### Example Usage

<!-- UsageSnippet language="go" operationID="getProductsSummaryByProduct" method="get" path="/products/{product}/summary" -->
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

    res, err := s.Nodes.Summaries.Get(ctx, components.ProductsBaseStream)
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedDistributedSummary != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `product`                                                          | [components.ProductsBase](../../models/components/productsbase.md) | :heavy_check_mark:                                                 | Name of the Cribl product to get the summary for.                  |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.GetProductsSummaryByProductResponse](../../models/operations/getproductssummarybyproductresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 401                | application/json   |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |