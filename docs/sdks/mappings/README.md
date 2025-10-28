# Mappings
(*Groups.Mappings*)

## Overview

### Available Operations

* [Activate](#activate) - Set a Mapping Ruleset as the active configuration for the specified Cribl product

## Activate

Set a specific Mapping Ruleset as the currently active configuration for the specified Cribl product.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAdminProductsMappingsActivateByProduct" method="post" path="/admin/products/{product}/mappings/activate" -->
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

    res, err := s.Groups.Mappings.Activate(ctx, components.ProductsCoreStream, components.RulesetID{
        ID: "default",
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

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `product`                                                          | [components.ProductsCore](../../models/components/productscore.md) | :heavy_check_mark:                                                 | Name of the Cribl product to activate the Mapping Ruleset for      |
| `rulesetID`                                                        | [components.RulesetID](../../models/components/rulesetid.md)       | :heavy_check_mark:                                                 | RulesetId object                                                   |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.CreateAdminProductsMappingsActivateByProductResponse](../../models/operations/createadminproductsmappingsactivatebyproductresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |