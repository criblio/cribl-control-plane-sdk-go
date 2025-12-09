# Pipelines

## Overview

Actions related to Pipelines

### Available Operations

* [List](#list) - List all Pipelines
* [Create](#create) - Create a Pipeline
* [Get](#get) - Get a Pipeline
* [Update](#update) - Update a Pipeline
* [Delete](#delete) - Delete a Pipeline

## List

Get a list of all Pipelines.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPipeline" method="get" path="/pipelines" -->
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

    res, err := s.Pipelines.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
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

## Create

Create a new Pipeline.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" -->
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

    res, err := s.Pipelines.Create(ctx, components.Pipeline{
        ID: "<id>",
        Conf: components.Conf{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](905091),
            Description: criblcontrolplanesdkgo.Pointer("next tightly positively"),
            Streamtags: []string{
                "<value 1>",
            },
            Functions: []components.PipelineFunctionConf{
                components.PipelineFunctionConf{
                    ID: "<id>",
                    Description: criblcontrolplanesdkgo.Pointer("academics woot finally woot queasy bah"),
                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                    Final: criblcontrolplanesdkgo.Pointer(false),
                    Conf: components.FunctionSpecificConfigs{},
                    GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                },
            },
            Groups: map[string]components.PipelineGroups{
                "key": components.PipelineGroups{
                    Name: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("yuck terribly ostrich enhance sentimental strictly whereas before reboot sleet"),
                    Disabled: criblcontrolplanesdkgo.Pointer(true),
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
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

## Get

Get the specified Pipeline.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPipelineById" method="get" path="/pipelines/{id}" -->
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

    res, err := s.Pipelines.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Pipeline to get.              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetPipelineByIDResponse](../../models/operations/getpipelinebyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the specified Pipeline.</br></br>Provide a complete representation of the Pipeline that you want to update in the request body. This endpoint does not support partial updates. Cribl removes any omitted fields when updating the Pipeline.</br></br>Confirm that the configuration in your request body is correct before sending the request. If the configuration is incorrect, the updated Pipeline might not function as expected.

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" -->
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

    res, err := s.Pipelines.Update(ctx, "<id>", components.Pipeline{
        ID: "<id>",
        Conf: components.Conf{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](430119),
            Description: criblcontrolplanesdkgo.Pointer("reflecting for productive extroverted instead upwardly"),
            Streamtags: []string{
                "<value 1>",
            },
            Functions: []components.PipelineFunctionConf{
                components.PipelineFunctionConf{
                    ID: "<id>",
                    Description: criblcontrolplanesdkgo.Pointer("mozzarella boohoo possession as grok"),
                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                    Final: criblcontrolplanesdkgo.Pointer(false),
                    Conf: components.FunctionSpecificConfigs{},
                    GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                },
            },
            Groups: map[string]components.PipelineGroups{
                "key": components.PipelineGroups{
                    Name: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("definitive ew but busily freely scaly indeed"),
                    Disabled: criblcontrolplanesdkgo.Pointer(true),
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `id`                                                       | *string*                                                   | :heavy_check_mark:                                         | The <code>id</code> of the Pipeline to update.             |
| `pipeline`                                                 | [components.Pipeline](../../models/components/pipeline.md) | :heavy_check_mark:                                         | Pipeline object                                            |
| `opts`                                                     | [][operations.Option](../../models/operations/option.md)   | :heavy_minus_sign:                                         | The options for this request.                              |

### Response

**[*operations.UpdatePipelineByIDResponse](../../models/operations/updatepipelinebyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete the specified Pipeline.

### Example Usage

<!-- UsageSnippet language="go" operationID="deletePipelineById" method="delete" path="/pipelines/{id}" -->
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

    res, err := s.Pipelines.Delete(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Pipeline to delete.           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeletePipelineByIDResponse](../../models/operations/deletepipelinebyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |