# Preview
(*Preview*)

## Overview

Actions related to data preview

### Available Operations

* [CreateSystemProjectsSubscriptionsCaptureByGroupIDAndSubscriptionID](#createsystemprojectssubscriptionscapturebygroupidandsubscriptionid) - Capture live incoming data from a particular Project and Subscription at the Subscription
* [CreateSystemProjectsCaptureByGroupIDAndProjectID](#createsystemprojectscapturebygroupidandprojectid) - Capture live incoming data from a particular Project and Subscription at the Destination
* [CreateSystemProjectsPreviewByGroupIDAndProjectID](#createsystemprojectspreviewbygroupidandprojectid) - Sends sample events through a Pipeline  for specified Project and returns the results
* [CreateSystemCapture](#createsystemcapture) - Capture live incoming data
* [CreateSystemProjectsCaptureByProjectID](#createsystemprojectscapturebyprojectid) - Capture live incoming data from a particular project and subscription at the destination
* [CreateSystemProjectsSubscriptionsCaptureByProjectIDAndSubscriptionID](#createsystemprojectssubscriptionscapturebyprojectidandsubscriptionid) - Capture live incoming data from a particular project and subscription at the subscription
* [CreatePreview](#createpreview) - Sends sample events through a pipeline and returns the results

## CreateSystemProjectsSubscriptionsCaptureByGroupIDAndSubscriptionID

Capture live incoming data from a particular Project and Subscription at the Subscription

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

    res, err := s.Preview.CreateSystemProjectsSubscriptionsCaptureByGroupIDAndSubscriptionID(ctx, "<id>", "<id>", "<id>", components.CaptureParams{
        Duration: 9977.94,
        Filter: "<value>",
        Level: components.CaptureLevelTwo,
        MaxEvents: 5232.64,
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

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `groupID`                                                            | *string*                                                             | :heavy_check_mark:                                                   | Group Id                                                             |
| `projectID`                                                          | *string*                                                             | :heavy_check_mark:                                                   | Project Id                                                           |
| `subscriptionID`                                                     | *string*                                                             | :heavy_check_mark:                                                   | Subscription Id                                                      |
| `captureParams`                                                      | [components.CaptureParams](../../models/components/captureparams.md) | :heavy_check_mark:                                                   | CaptureParams object                                                 |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.CreateSystemProjectsSubscriptionsCaptureByGroupIDAndSubscriptionIDResponse](../../models/operations/createsystemprojectssubscriptionscapturebygroupidandsubscriptionidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateSystemProjectsCaptureByGroupIDAndProjectID

Capture live incoming data from a particular Project and Subscription at the Destination

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

    res, err := s.Preview.CreateSystemProjectsCaptureByGroupIDAndProjectID(ctx, "<id>", "<id>", components.CaptureParams{
        Duration: 7020.88,
        Filter: "<value>",
        Level: components.CaptureLevelOne,
        MaxEvents: 7103.16,
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

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `groupID`                                                            | *string*                                                             | :heavy_check_mark:                                                   | Group Id                                                             |
| `projectID`                                                          | *string*                                                             | :heavy_check_mark:                                                   | Project Id                                                           |
| `captureParams`                                                      | [components.CaptureParams](../../models/components/captureparams.md) | :heavy_check_mark:                                                   | CaptureParams object                                                 |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.CreateSystemProjectsCaptureByGroupIDAndProjectIDResponse](../../models/operations/createsystemprojectscapturebygroupidandprojectidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateSystemProjectsPreviewByGroupIDAndProjectID

Sends sample events through a Pipeline  for specified Project and returns the results

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

    res, err := s.Preview.CreateSystemProjectsPreviewByGroupIDAndProjectID(ctx, "<id>", "<id>", components.PreviewDataParams{
        Mode: components.PreviewDataParamsModeRoute,
        PipelineID: "<id>",
        SampleID: "<id>",
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `groupID`                                                                    | *string*                                                                     | :heavy_check_mark:                                                           | Group Id                                                                     |
| `projectID`                                                                  | *string*                                                                     | :heavy_check_mark:                                                           | Project Id                                                                   |
| `previewDataParams`                                                          | [components.PreviewDataParams](../../models/components/previewdataparams.md) | :heavy_check_mark:                                                           | PreviewDataParams object                                                     |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.CreateSystemProjectsPreviewByGroupIDAndProjectIDResponse](../../models/operations/createsystemprojectspreviewbygroupidandprojectidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateSystemCapture

Capture live incoming data

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

    res, err := s.Preview.CreateSystemCapture(ctx, components.CaptureParams{
        Duration: 5717.32,
        Filter: "<value>",
        Level: components.CaptureLevelTwo,
        MaxEvents: 9941.84,
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

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [components.CaptureParams](../../models/components/captureparams.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.CreateSystemCaptureResponse](../../models/operations/createsystemcaptureresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateSystemProjectsCaptureByProjectID

Capture live incoming data from a particular project and subscription at the destination

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

    res, err := s.Preview.CreateSystemProjectsCaptureByProjectID(ctx, "<id>", components.CaptureParams{
        Duration: 4855.02,
        Filter: "<value>",
        Level: components.CaptureLevelThree,
        MaxEvents: 5475.12,
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

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `projectID`                                                          | *string*                                                             | :heavy_check_mark:                                                   | Project ID                                                           |
| `captureParams`                                                      | [components.CaptureParams](../../models/components/captureparams.md) | :heavy_check_mark:                                                   | CaptureParams object                                                 |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.CreateSystemProjectsCaptureByProjectIDResponse](../../models/operations/createsystemprojectscapturebyprojectidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateSystemProjectsSubscriptionsCaptureByProjectIDAndSubscriptionID

Capture live incoming data from a particular project and subscription at the subscription

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

    res, err := s.Preview.CreateSystemProjectsSubscriptionsCaptureByProjectIDAndSubscriptionID(ctx, "<id>", "<id>", components.CaptureParams{
        Duration: 9984.14,
        Filter: "<value>",
        Level: components.CaptureLevelThree,
        MaxEvents: 502.1,
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

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `projectID`                                                          | *string*                                                             | :heavy_check_mark:                                                   | Project Id                                                           |
| `subscriptionID`                                                     | *string*                                                             | :heavy_check_mark:                                                   | SubscriptionId Id                                                    |
| `captureParams`                                                      | [components.CaptureParams](../../models/components/captureparams.md) | :heavy_check_mark:                                                   | CaptureParams object                                                 |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.CreateSystemProjectsSubscriptionsCaptureByProjectIDAndSubscriptionIDResponse](../../models/operations/createsystemprojectssubscriptionscapturebyprojectidandsubscriptionidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreatePreview

Sends sample events through a pipeline and returns the results

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

    res, err := s.Preview.CreatePreview(ctx, components.PreviewDataParams{
        Mode: components.PreviewDataParamsModePipe,
        PipelineID: "<id>",
        SampleID: "<id>",
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