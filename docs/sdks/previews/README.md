# System.Previews

## Overview

### Available Operations

* [Create](#create) - Send sample events through a Pipeline and review results

## Create

Send sample events through the specified Pipeline.The response contains an array of objects, each showing the transformations,additions, and removals for an event after it passes through the pipeline.Useful for testing Pipeline logic with sample data before deploying changes.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPreview" method="post" path="/preview" -->
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

    res, err := s.System.Previews.Create(ctx, components.PreviewDataParams{
        Mode: components.PreviewDataParamsModePipe,
        PipelineID: "<id>",
        SampleID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PreviewResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.PreviewDataParams](../../models/components/previewdataparams.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.CreatePreviewResponse](../../models/operations/createpreviewresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |