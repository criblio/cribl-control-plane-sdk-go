# Packs.Sources.HecTokens

## Overview

### Available Operations

* [Create](#create) - Add an HEC token and optional metadata to a Splunk HEC Source within a Pack
* [Update](#update) - Update metadata for an HEC token for a Splunk HEC Source within a Pack

## Create

Add an HEC token and optional metadata to the specified Splunk HEC Source within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="createInputSystemHecTokenByPackAndId" method="post" path="/p/{pack}/system/inputs/{id}/hectoken" -->
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

    res, err := s.Packs.Sources.HecTokens.Create(ctx, "<id>", "<value>", components.AddHecTokenRequest{
        AllowedIndexesAtToken: []string{
            "<value 1>",
        },
        Description: criblcontrolplanesdkgo.Pointer("SUV velvety without"),
        Enabled: criblcontrolplanesdkgo.Pointer(true),
        Metadata: []components.EventBreakerRuleFields{
            components.EventBreakerRuleFields{
                Name: "fieldX",
                Value: "valueX",
            },
        },
        Token: "12345678901",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputSplunkHec != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `id`                                                                           | *string*                                                                       | :heavy_check_mark:                                                             | The <code>id</code> of the Splunk HEC Source.                                  |
| `pack`                                                                         | *string*                                                                       | :heavy_check_mark:                                                             | The <code>id</code> of the Pack to create.                                     |
| `addHecTokenRequest`                                                           | [components.AddHecTokenRequest](../../models/components/addhectokenrequest.md) | :heavy_check_mark:                                                             | AddHecTokenRequest object                                                      |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CreateInputSystemHecTokenByPackAndIDResponse](../../models/operations/createinputsystemhectokenbypackandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the metadata for the specified HEC token for the specified Splunk HEC Source within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateInputSystemHecTokenByPackAndIdAndToken" method="patch" path="/p/{pack}/system/inputs/{id}/hectoken/{token}" -->
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

    res, err := s.Packs.Sources.HecTokens.Update(ctx, "<id>", "<value>", "<value>", components.UpdateHecTokenRequest{
        AllowedIndexesAtToken: []string{
            "<value 1>",
            "<value 2>",
            "<value 3>",
        },
        Description: criblcontrolplanesdkgo.Pointer("straw critical famously dream after delightfully"),
        Enabled: criblcontrolplanesdkgo.Pointer(true),
        Metadata: []components.EventBreakerRuleFields{
            components.EventBreakerRuleFields{
                Name: "fieldX",
                Value: "valueX",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputSplunkHec != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `id`                                                                                 | *string*                                                                             | :heavy_check_mark:                                                                   | The <code>id</code> of the Splunk HEC Source.                                        |
| `token`                                                                              | *string*                                                                             | :heavy_check_mark:                                                                   | The <code>id</code> of the HEC token to update.                                      |
| `pack`                                                                               | *string*                                                                             | :heavy_check_mark:                                                                   | The <code>id</code> of the Pack to update.                                           |
| `updateHecTokenRequest`                                                              | [components.UpdateHecTokenRequest](../../models/components/updatehectokenrequest.md) | :heavy_check_mark:                                                                   | UpdateHecTokenRequest object                                                         |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdateInputSystemHecTokenByPackAndIDAndTokenResponse](../../models/operations/updateinputsystemhectokenbypackandidandtokenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |