# Packs.Sources

## Overview

### Available Operations

* [List](#list) - List all Sources within a Pack
* [Create](#create) - Create a Source within a Pack
* [Get](#get) - Get a Source within a Pack
* [Update](#update) - Update a Source within a Pack
* [Delete](#delete) - Delete a Source within a Pack

## List

Get a list of all Sources within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="getInputSystemByPack" method="get" path="/p/{pack}/system/inputs" -->
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

    res, err := s.Packs.Sources.List(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInput != nil {
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

**[*operations.GetInputSystemByPackResponse](../../models/operations/getinputsystembypackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a new Source within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyAppscope(
        operations.CreateInputSystemByPackInputAppscope{
            ID: "appscope-source",
            Type: operations.CreateInputSystemByPackTypeAppscopeAppscope,
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Connections: []components.ItemsTypeConnectionsOptional{
                components.ItemsTypeConnectionsOptional{
                    Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                    Output: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Pq: &components.PqType{
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](4067.15),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](1373.9),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](5473.3),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](7426.73),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](4309.15),
            EnableUnixPath: criblcontrolplanesdkgo.Pointer(true),
            Filter: &operations.CreateInputSystemByPackFilterAppscope{
                Allow: []operations.CreateInputSystemByPackAllow{
                    operations.CreateInputSystemByPackAllow{
                        Procname: "<value>",
                        Arg: criblcontrolplanesdkgo.Pointer("<value>"),
                        Config: "<value>",
                    },
                },
                TransportURL: criblcontrolplanesdkgo.Pointer("https://negative-asset.info/"),
            },
            Persistence: &operations.CreateInputSystemByPackPersistenceAppscope{
                Enable: criblcontrolplanesdkgo.Pointer(true),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.DataCompressionFormatOptionsPersistenceNone.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("char antagonize yuck"),
            Host: criblcontrolplanesdkgo.Pointer("0.0.0.0"),
            Port: criblcontrolplanesdkgo.Pointer[float64](9109),
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                RequestCert: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            UnixSocketPath: criblcontrolplanesdkgo.Pointer("<value>"),
            UnixSocketPerms: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthToken: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `pack`                                                                                                         | *string*                                                                                                       | :heavy_check_mark:                                                                                             | The <code>id</code> of the Pack to create.                                                                     |
| `requestBody`                                                                                                  | [operations.CreateInputSystemByPackRequestBody](../../models/operations/createinputsystembypackrequestbody.md) | :heavy_check_mark:                                                                                             | Input object                                                                                                   |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.CreateInputSystemByPackResponse](../../models/operations/createinputsystembypackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get the specified Source within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="getInputSystemByPackAndId" method="get" path="/p/{pack}/system/inputs/{id}" -->
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

    res, err := s.Packs.Sources.Get(ctx, "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Source to get.                |
| `pack`                                                   | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Pack to get.                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetInputSystemByPackAndIDResponse](../../models/operations/getinputsystembypackandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the specified Source.</br></br>Provide a complete representation of the Source that you want to update in the request body. This endpoint does not support partial updates. Cribl removes any omitted fields when updating the Source.</br></br>Confirm that the configuration in your request body is correct before sending the request. If the configuration is incorrect, the updated Source might not function as expected within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputAppscope(
        components.InputAppscope{
            ID: criblcontrolplanesdkgo.Pointer("appscope-source"),
            Type: components.InputAppscopeTypeAppscope,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Connections: []components.ItemsTypeConnectionsOptional{
                components.ItemsTypeConnectionsOptional{
                    Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                    Output: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Pq: &components.PqType{
                Mode: components.ModeOptionsPqSmart.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](2177.87),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](1131.76),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](2476.4),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](8980.2),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            BreakerRulesets: []string{
                "<value 1>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](5201.14),
            EnableUnixPath: criblcontrolplanesdkgo.Pointer(false),
            Filter: &components.InputAppscopeFilter{
                Allow: []components.Allow{
                    components.Allow{
                        Procname: "<value>",
                        Arg: criblcontrolplanesdkgo.Pointer("<value>"),
                        Config: "<value>",
                    },
                },
                TransportURL: criblcontrolplanesdkgo.Pointer("https://stable-transparency.org/"),
            },
            Persistence: &components.InputAppscopePersistence{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.DataCompressionFormatOptionsPersistenceNone.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("graffiti yawningly absent when times sonata or uselessly woeful amidst"),
            Host: criblcontrolplanesdkgo.Pointer("0.0.0.0"),
            Port: criblcontrolplanesdkgo.Pointer[float64](9109),
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            UnixSocketPath: criblcontrolplanesdkgo.Pointer("<value>"),
            UnixSocketPerms: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthToken: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Source to update.             |
| `pack`                                                   | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Pack to update.               |
| `input`                                                  | [components.Input](../../models/components/input.md)     | :heavy_check_mark:                                       | Input object                                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UpdateInputSystemByPackAndIDResponse](../../models/operations/updateinputsystembypackandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete the specified Source within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteInputSystemByPackAndId" method="delete" path="/p/{pack}/system/inputs/{id}" -->
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

    res, err := s.Packs.Sources.Delete(ctx, "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Source to delete.             |
| `pack`                                                   | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Pack to delete.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteInputSystemByPackAndIDResponse](../../models/operations/deleteinputsystembypackandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |