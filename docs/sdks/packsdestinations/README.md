# Packs.Destinations

## Overview

### Available Operations

* [List](#list) - List all Destinations within a Pack
* [Create](#create) - Create a Destination within a Pack
* [Get](#get) - Get a Destination within a Pack
* [Update](#update) - Update a Destination within a Pack
* [Delete](#delete) - Delete a Destination within a Pack

## List

Get a list of all Destinations within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="getOutputSystemByPack" method="get" path="/p/{pack}/system/outputs" -->
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

    res, err := s.Packs.Destinations.List(ctx, "<value>")
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
| `pack`                                                   | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Pack to list.                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetOutputSystemByPackResponse](../../models/operations/getoutputsystembypackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a new Destination within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyTcpjson(
        operations.CreateOutputSystemByPackOutputTcpjson{
            ID: "tcpjson-output",
            Type: operations.CreateOutputSystemByPackTypeTcpjsonTcpjson,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            LoadBalanced: criblcontrolplanesdkgo.Pointer(true),
            Compression: components.CompressionOptions1Gzip.ToPointer(),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(false),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](4098.5),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](3807.97),
            TokenTTLMinutes: criblcontrolplanesdkgo.Pointer[float64](9215.23),
            SendHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("unaccountably hard-to-find same kindly deprave the moisten task"),
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](10090),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(true),
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "adolescent-cemetery.org",
                    Port: 5699.66,
                    TLS: components.TLSOptionsHostsItemsOff.ToPointer(),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    Weight: criblcontrolplanesdkgo.Pointer[float64](7378.74),
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](1392.89),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](9185.86),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](9360.71),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1119.07),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4924.47),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](7420.74),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsTcpjson{},
            AuthToken: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `pack`                                                                                                           | *string*                                                                                                         | :heavy_check_mark:                                                                                               | The <code>id</code> of the Pack to create.                                                                       |
| `requestBody`                                                                                                    | [operations.CreateOutputSystemByPackRequestBody](../../models/operations/createoutputsystembypackrequestbody.md) | :heavy_check_mark:                                                                                               | Output object                                                                                                    |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.CreateOutputSystemByPackResponse](../../models/operations/createoutputsystembypackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get the specified Destination within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="getOutputSystemByPackAndId" method="get" path="/p/{pack}/system/outputs/{id}" -->
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

    res, err := s.Packs.Destinations.Get(ctx, "<id>", "<value>")
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
| `pack`                                                   | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Pack to get.                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetOutputSystemByPackAndIDResponse](../../models/operations/getoutputsystembypackandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the specified Destination.</br></br>Provide a complete representation of the Destination that you want to update in the request body. This endpoint does not support partial updates. Cribl removes any omitted fields when updating the Destination.</br></br>Confirm that the configuration in your request body is correct before sending the request. If the configuration is incorrect, the updated Destination might not function as expected within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputDefault(
        components.OutputDefault{
            ID: criblcontrolplanesdkgo.Pointer("default-output"),
            Type: components.OutputDefaultTypeDefault,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            DefaultID: criblcontrolplanesdkgo.Pointer("my-default-output"),
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
| `pack`                                                   | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Pack to update.               |
| `output`                                                 | [components.Output](../../models/components/output.md)   | :heavy_check_mark:                                       | Output object                                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UpdateOutputSystemByPackAndIDResponse](../../models/operations/updateoutputsystembypackandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete the specified Destination within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteOutputSystemByPackAndId" method="delete" path="/p/{pack}/system/outputs/{id}" -->
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

    res, err := s.Packs.Destinations.Delete(ctx, "<id>", "<value>")
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
| `pack`                                                   | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Pack to delete.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteOutputSystemByPackAndIDResponse](../../models/operations/deleteoutputsystembypackandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |