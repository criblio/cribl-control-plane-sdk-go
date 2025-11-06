# HecTokens
(*Sources.HecTokens*)

## Overview

### Available Operations

* [Create](#create) - Add an HEC token and optional metadata to a Splunk HEC Source
* [Update](#update) - Update metadata for an HEC token for a Splunk HEC Source

## Create

Add an HEC token and optional metadata to the specified Splunk HEC Source.

### Example Usage

<!-- UsageSnippet language="go" operationID="createInputHecTokenById" method="post" path="/system/inputs/{id}/hectoken" -->
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

    res, err := s.Sources.HecTokens.Create(ctx, "<id>", components.AddHecTokenRequest{
        Description: criblcontrolplanesdkgo.Pointer("bah ick stingy"),
        Enabled: criblcontrolplanesdkgo.Pointer(false),
        Metadata: []components.AddHecTokenRequestMetadatum{
            components.AddHecTokenRequestMetadatum{
                Name: "<value>",
                Value: "<value>",
            },
        },
        Token: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedListInputSplunkHec != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `id`                                                                           | *string*                                                                       | :heavy_check_mark:                                                             | The <code>id</code> of the Splunk HEC Source.                                  |
| `addHecTokenRequest`                                                           | [components.AddHecTokenRequest](../../models/components/addhectokenrequest.md) | :heavy_check_mark:                                                             | AddHecTokenRequest object                                                      |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CreateInputHecTokenByIDResponse](../../models/operations/createinputhectokenbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the metadata for the specified HEC token for the specified Splunk HEC Source.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateInputHecTokenByIdAndToken" method="patch" path="/system/inputs/{id}/hectoken/{token}" -->
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

    res, err := s.Sources.HecTokens.Update(ctx, "<id>", "<value>", components.UpdateHecTokenRequest{
        Description: criblcontrolplanesdkgo.Pointer("by bleakly fortunately phew barring"),
        Enabled: criblcontrolplanesdkgo.Pointer(false),
        Metadata: []components.UpdateHecTokenRequestMetadatum{
            components.UpdateHecTokenRequestMetadatum{
                Name: "<value>",
                Value: "<value>",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedListInputSplunkHec != nil {
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
| `updateHecTokenRequest`                                                              | [components.UpdateHecTokenRequest](../../models/components/updatehectokenrequest.md) | :heavy_check_mark:                                                                   | UpdateHecTokenRequest object                                                         |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.UpdateInputHecTokenByIDAndTokenResponse](../../models/operations/updateinputhectokenbyidandtokenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |