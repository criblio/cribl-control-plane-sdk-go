# Pipelines
(*Pipelines*)

## Overview

Actions related to Pipelines

### Available Operations

* [ListPipeline](#listpipeline) - Get a list of Pipeline objects
* [CreatePipeline](#createpipeline) - Create Pipeline
* [GetPipelineByID](#getpipelinebyid) - Get Pipeline by ID
* [UpdatePipelineByID](#updatepipelinebyid) - Update Pipeline
* [DeletePipelineByID](#deletepipelinebyid) - Delete Pipeline

## ListPipeline

Get a list of Pipeline objects

### Example Usage

<!-- UsageSnippet language="go" operationID="listPipeline" method="get" path="/pipelines" -->
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

    res, err := s.Pipelines.ListPipeline(ctx)
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

**[*operations.ListPipelineResponse](../../models/operations/listpipelineresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreatePipeline

Create Pipeline

### Example Usage

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" -->
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

    res, err := s.Pipelines.CreatePipeline(ctx, components.Pipeline{
        ID: "<id>",
        Conf: components.Conf{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Int64(905091),
            Description: criblcontrolplanesdkgo.String("next tightly positively"),
            Streamtags: []string{
                "<value 1>",
            },
            Functions: []components.PipelineFunctionConf{
                components.PipelineFunctionConf{
                    ID: "<id>",
                    Description: criblcontrolplanesdkgo.String("academics woot finally woot queasy bah"),
                    Disabled: criblcontrolplanesdkgo.Bool(false),
                    Final: criblcontrolplanesdkgo.Bool(false),
                    Conf: components.FunctionSpecificConfigs{},
                    GroupID: criblcontrolplanesdkgo.String("<id>"),
                },
            },
            Groups: map[string]components.PipelineGroups{
                "key": components.PipelineGroups{
                    Name: "<value>",
                    Description: criblcontrolplanesdkgo.String("yuck terribly ostrich enhance sentimental strictly whereas before reboot sleet"),
                    Disabled: criblcontrolplanesdkgo.Bool(true),
                },
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

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `request`                                                  | [components.Pipeline](../../models/components/pipeline.md) | :heavy_check_mark:                                         | The request object to use for the request.                 |
| `opts`                                                     | [][operations.Option](../../models/operations/option.md)   | :heavy_minus_sign:                                         | The options for this request.                              |

### Response

**[*operations.CreatePipelineResponse](../../models/operations/createpipelineresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetPipelineByID

Get Pipeline by ID

### Example Usage

<!-- UsageSnippet language="go" operationID="getPipelineById" method="get" path="/pipelines/{id}" -->
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

    res, err := s.Pipelines.GetPipelineByID(ctx, "<id>")
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

**[*operations.GetPipelineByIDResponse](../../models/operations/getpipelinebyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdatePipelineByID

Update Pipeline

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" -->
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

    res, err := s.Pipelines.UpdatePipelineByID(ctx, "<id>", components.Pipeline{
        ID: "<id>",
        Conf: components.Conf{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Int64(430119),
            Description: criblcontrolplanesdkgo.String("reflecting for productive extroverted instead upwardly"),
            Streamtags: []string{
                "<value 1>",
            },
            Functions: []components.PipelineFunctionConf{
                components.PipelineFunctionConf{
                    ID: "<id>",
                    Description: criblcontrolplanesdkgo.String("mozzarella boohoo possession as grok"),
                    Disabled: criblcontrolplanesdkgo.Bool(false),
                    Final: criblcontrolplanesdkgo.Bool(false),
                    Conf: components.FunctionSpecificConfigs{},
                    GroupID: criblcontrolplanesdkgo.String("<id>"),
                },
            },
            Groups: map[string]components.PipelineGroups{
                "key": components.PipelineGroups{
                    Name: "<value>",
                    Description: criblcontrolplanesdkgo.String("definitive ew but busily freely scaly indeed"),
                    Disabled: criblcontrolplanesdkgo.Bool(true),
                },
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

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `id`                                                       | *string*                                                   | :heavy_check_mark:                                         | Unique ID to PATCH                                         |
| `pipeline`                                                 | [components.Pipeline](../../models/components/pipeline.md) | :heavy_check_mark:                                         | Pipeline object to be updated                              |
| `opts`                                                     | [][operations.Option](../../models/operations/option.md)   | :heavy_minus_sign:                                         | The options for this request.                              |

### Response

**[*operations.UpdatePipelineByIDResponse](../../models/operations/updatepipelinebyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeletePipelineByID

Delete Pipeline

### Example Usage

<!-- UsageSnippet language="go" operationID="deletePipelineById" method="delete" path="/pipelines/{id}" -->
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

    res, err := s.Pipelines.DeletePipelineByID(ctx, "<id>")
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

**[*operations.DeletePipelineByIDResponse](../../models/operations/deletepipelinebyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |