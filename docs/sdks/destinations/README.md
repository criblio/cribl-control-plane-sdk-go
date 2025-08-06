# Destinations
(*Destinations*)

## Overview

### Available Operations

* [ListDestination](#listdestination) - Get a list of Destination objects
* [CreateDestination](#createdestination) - Create Destination
* [GetDestinationByID](#getdestinationbyid) - Get Destination by ID
* [UpdateDestinationByID](#updatedestinationbyid) - Update Destination
* [DeleteDestinationByID](#deletedestinationbyid) - Delete Destination
* [DeleteDestinationPqByID](#deletedestinationpqbyid) - Clears destination persistent queue
* [GetDestinationPqByID](#getdestinationpqbyid) - Retrieves status of latest clear PQ job for a destination
* [GetDestinationSamplesByID](#getdestinationsamplesbyid) - Retrieve samples data for the specified destination. Used to get sample data for the test action.
* [CreateDestinationTestByID](#createdestinationtestbyid) - Send sample data to a destination to validate configuration or test connectivity

## ListDestination

Get a list of Destination objects

### Example Usage

<!-- UsageSnippet language="go" operationID="listOutput" method="get" path="/system/outputs" -->
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.ListDestination(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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

**[*operations.ListOutputResponse](../../models/operations/listoutputresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateDestination

Create Destination

### Example Usage

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" -->
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.CreateDestination(ctx, components.CreateOutputOutputElasticCloud(
        components.OutputElasticCloud{
            ID: criblcontrolplanesdkgo.String("<id>"),
            Type: components.OutputElasticCloudTypeElasticCloud.ToPointer(),
            Pipeline: criblcontrolplanesdkgo.String("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.String("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            URL: "https://probable-rationale.com/",
            Index: "<value>",
            ExtraHTTPHeaders: []components.OutputElasticCloudExtraHTTPHeader{
                components.OutputElasticCloudExtraHTTPHeader{
                    Name: criblcontrolplanesdkgo.String("<value>"),
                    Value: "<value>",
                },
            },
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            ExtraParams: []components.OutputElasticCloudExtraParam{
                components.OutputElasticCloudExtraParam{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Auth: &components.OutputElasticCloudAuth{},
            ElasticPipeline: criblcontrolplanesdkgo.String("<value>"),
            ResponseRetrySettings: []components.OutputElasticCloudResponseRetrySetting{
                components.OutputElasticCloudResponseRetrySetting{
                    HTTPStatus: 7295.73,
                },
            },
            TimeoutRetrySettings: &components.OutputElasticCloudTimeoutRetrySettings{},
            Description: criblcontrolplanesdkgo.String("hourly about into"),
            PqControls: &components.OutputElasticCloudPqControls{},
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [components.Output](../../models/components/output.md)   | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CreateOutputResponse](../../models/operations/createoutputresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetDestinationByID

Get Destination by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="getOutputById" method="get" path="/system/outputs/{id}" -->
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.GetDestinationByID(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Unique ID to GET                                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetOutputByIDResponse](../../models/operations/getoutputbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateDestinationByID

Update Destination

### Example Usage

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" -->
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.UpdateDestinationByID(ctx, "<id>", components.CreateOutputOutputSignalfx(
        components.OutputSignalfx{
            ID: criblcontrolplanesdkgo.String("<id>"),
            Type: components.OutputSignalfxTypeSignalfx,
            Pipeline: criblcontrolplanesdkgo.String("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.String("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            ExtraHTTPHeaders: []components.OutputSignalfxExtraHTTPHeader{
                components.OutputSignalfxExtraHTTPHeader{
                    Name: criblcontrolplanesdkgo.String("<value>"),
                    Value: "<value>",
                },
            },
            SafeHeaders: []string{
                "<value 1>",
            },
            ResponseRetrySettings: []components.OutputSignalfxResponseRetrySetting{
                components.OutputSignalfxResponseRetrySetting{
                    HTTPStatus: 2924.72,
                },
            },
            TimeoutRetrySettings: &components.OutputSignalfxTimeoutRetrySettings{},
            Description: criblcontrolplanesdkgo.String("phooey positively a consequently meh until"),
            Token: criblcontrolplanesdkgo.String("<value>"),
            TextSecret: criblcontrolplanesdkgo.String("<value>"),
            PqControls: &components.OutputSignalfxPqControls{},
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Unique ID to PATCH                                       |
| `output`                                                 | [components.Output](../../models/components/output.md)   | :heavy_check_mark:                                       | Output object                                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UpdateOutputByIDResponse](../../models/operations/updateoutputbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteDestinationByID

Delete Destination

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteOutputById" method="delete" path="/system/outputs/{id}" -->
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.DeleteDestinationByID(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Unique ID to DELETE                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteOutputByIDResponse](../../models/operations/deleteoutputbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteDestinationPqByID

Clears destination persistent queue

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteOutputPqById" method="delete" path="/system/outputs/{id}/pq" -->
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.DeleteDestinationPqByID(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Destination Id                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteOutputPqByIDResponse](../../models/operations/deleteoutputpqbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetDestinationPqByID

Retrieves status of latest clear PQ job for a destination

### Example Usage

<!-- UsageSnippet language="go" operationID="getOutputPqById" method="get" path="/system/outputs/{id}/pq" -->
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.GetDestinationPqByID(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Destination Id                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetOutputPqByIDResponse](../../models/operations/getoutputpqbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetDestinationSamplesByID

Retrieve samples data for the specified destination. Used to get sample data for the test action.

### Example Usage

<!-- UsageSnippet language="go" operationID="getOutputSamplesById" method="get" path="/system/outputs/{id}/samples" -->
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.GetDestinationSamplesByID(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Destination Id                                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetOutputSamplesByIDResponse](../../models/operations/getoutputsamplesbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateDestinationTestByID

Send sample data to a destination to validate configuration or test connectivity

### Example Usage

<!-- UsageSnippet language="go" operationID="createOutputTestById" method="post" path="/system/outputs/{id}/test" -->
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.CreateDestinationTestByID(ctx, "<id>", components.OutputTestRequest{
        Events: []components.CriblEvent{
            components.CriblEvent{
                Raw: "<value>",
            },
        },
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
| `id`                                                                         | *string*                                                                     | :heavy_check_mark:                                                           | Destination Id                                                               |
| `outputTestRequest`                                                          | [components.OutputTestRequest](../../models/components/outputtestrequest.md) | :heavy_check_mark:                                                           | OutputTestRequest object                                                     |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.CreateOutputTestByIDResponse](../../models/operations/createoutputtestbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |