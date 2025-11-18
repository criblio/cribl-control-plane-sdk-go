# Destinations
(*Destinations*)

## Overview

Actions related to Destinations

### Available Operations

* [List](#list) - List all Destinations
* [Create](#create) - Create a Destination
* [Get](#get) - Get a Destination
* [Update](#update) - Update a Destination
* [Delete](#delete) - Delete a Destination

## List

Get a list of all Destinations.

### Example Usage

<!-- UsageSnippet language="go" operationID="listOutput" method="get" path="/system/outputs" -->
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

    res, err := s.Destinations.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedOutput != nil {
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

## Create

Create a new Destination.

### Example Usage

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" -->
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

    res, err := s.Destinations.Create(ctx, components.CreateOutputElasticCloud(
        components.OutputElasticCloud{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Type: components.TypeElasticCloudElasticCloud,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            URL: "https://probable-rationale.com/",
            Index: "<value>",
            ExtraHTTPHeaders: []components.ExtraHTTPHeaderElasticCloud{
                components.ExtraHTTPHeaderElasticCloud{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            ExtraParams: []components.ExtraParamElasticCloud{
                components.ExtraParamElasticCloud{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Auth: &components.AuthElasticCloud{
                Username: criblcontrolplanesdkgo.Pointer("Imogene1"),
                Password: criblcontrolplanesdkgo.Pointer("tA9DUjJtHSnXqh9"),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                ManualAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            ElasticPipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            ResponseRetrySettings: []components.ResponseRetrySettingElasticCloud{
                components.ResponseRetrySettingElasticCloud{
                    HTTPStatus: 7295.73,
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsElasticCloud{},
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("hourly about into"),
            PqControls: &components.PqControlsElasticCloud{},
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedOutput != nil {
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

## Get

Get the specified Destination.

### Example Usage

<!-- UsageSnippet language="go" operationID="getOutputById" method="get" path="/system/outputs/{id}" -->
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

    res, err := s.Destinations.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Destination to get.           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetOutputByIDResponse](../../models/operations/getoutputbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the specified Destination.</br></br>Provide a complete representation of the Destination that you want to update in the request body. This endpoint does not support partial updates. Cribl removes any omitted fields when updating the Destination.</br></br>Confirm that the configuration in your request body is correct before sending the request. If the configuration is incorrect, the updated Destination might not function as expected.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" -->
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

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputSignalfx(
        components.OutputSignalfx{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Type: components.TypeSignalfxSignalfx,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            ExtraHTTPHeaders: []components.ExtraHTTPHeaderSignalfx{
                components.ExtraHTTPHeaderSignalfx{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            SafeHeaders: []string{
                "<value 1>",
            },
            ResponseRetrySettings: []components.ResponseRetrySettingSignalfx{
                components.ResponseRetrySettingSignalfx{
                    HTTPStatus: 2924.72,
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsSignalfx{},
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("phooey positively a consequently meh until"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqControls: &components.PqControlsSignalfx{},
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Destination to update.        |
| `output`                                                 | [components.Output](../../models/components/output.md)   | :heavy_check_mark:                                       | Output object                                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UpdateOutputByIDResponse](../../models/operations/updateoutputbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete the specified Destination.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteOutputById" method="delete" path="/system/outputs/{id}" -->
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

    res, err := s.Destinations.Delete(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Destination to delete.        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteOutputByIDResponse](../../models/operations/deleteoutputbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |