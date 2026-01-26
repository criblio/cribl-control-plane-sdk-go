# Sources

## Overview

Actions related to Sources

### Available Operations

* [List](#list) - List all Sources
* [Create](#create) - Create a Source
* [Get](#get) - Get a Source
* [Update](#update) - Update a Source
* [Delete](#delete) - Delete a Source

## List

Get a list of all Sources.

### Example Usage

<!-- UsageSnippet language="go" operationID="listInput" method="get" path="/system/inputs" -->
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

    res, err := s.Sources.List(ctx)
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListInputResponse](../../models/operations/listinputresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a new Source.

### Example Usage

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestAppscope(
        operations.InputAppscope{
            ID: "appscope-source",
            Type: operations.TypeAppscopeAppscope,
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
            },
            Connections: []components.ItemsTypeConnectionsOptional{
                components.ItemsTypeConnectionsOptional{
                    Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                    Output: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Pq: &components.PqType{
                Mode: components.ModeOptionsPqSmart.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](4887.41),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](2674.23),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](7415.32),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](7847.75),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](3076.3),
            EnableUnixPath: criblcontrolplanesdkgo.Pointer(false),
            Filter: &operations.FilterAppscope{
                Allow: []operations.Allow{
                    operations.Allow{
                        Procname: "<value>",
                        Arg: criblcontrolplanesdkgo.Pointer("<value>"),
                        Config: "<value>",
                    },
                },
                TransportURL: criblcontrolplanesdkgo.Pointer("https://pointed-napkin.info/"),
            },
            Persistence: &operations.PersistenceAppscope{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.DataCompressionFormatOptionsPersistenceNone.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("repeatedly urban incidentally clean up"),
            Host: criblcontrolplanesdkgo.Pointer("0.0.0.0"),
            Port: criblcontrolplanesdkgo.Pointer[float64](9109),
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
            },
            UnixSocketPath: criblcontrolplanesdkgo.Pointer("<value>"),
            UnixSocketPerms: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthToken: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.CreateInputRequest](../../models/operations/createinputrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CreateInputResponse](../../models/operations/createinputresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get the specified Source.

### Example Usage

<!-- UsageSnippet language="go" operationID="getInputById" method="get" path="/system/inputs/{id}" -->
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

    res, err := s.Sources.Get(ctx, "<id>")
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetInputByIDResponse](../../models/operations/getinputbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the specified Source.</br></br>Provide a complete representation of the Source that you want to update in the request body. This endpoint does not support partial updates. Cribl removes any omitted fields when updating the Source.</br></br>Confirm that the configuration in your request body is correct before sending the request. If the configuration is incorrect, the updated Source might not function as expected.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputAppscope(
        components.InputAppscope{
            ID: criblcontrolplanesdkgo.Pointer("appscope-source"),
            Type: components.InputAppscopeTypeAppscope,
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
            },
            Connections: []components.ItemsTypeConnectionsOptional{
                components.ItemsTypeConnectionsOptional{
                    Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                    Output: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Pq: &components.PqType{
                Mode: components.ModeOptionsPqSmart.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](3417.54),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](4799.95),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](3730.65),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](4634.53),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            BreakerRulesets: []string{
                "<value 1>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](3362.61),
            EnableUnixPath: criblcontrolplanesdkgo.Pointer(false),
            Filter: &components.InputAppscopeFilter{
                Allow: []components.Allow{
                    components.Allow{
                        Procname: "<value>",
                        Arg: criblcontrolplanesdkgo.Pointer("<value>"),
                        Config: "<value>",
                    },
                },
                TransportURL: criblcontrolplanesdkgo.Pointer("https://distorted-translation.org/"),
            },
            Persistence: &components.InputAppscopePersistence{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.DataCompressionFormatOptionsPersistenceNone.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("incidentally down versus blah"),
            Host: criblcontrolplanesdkgo.Pointer("0.0.0.0"),
            Port: criblcontrolplanesdkgo.Pointer[float64](9109),
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            UnixSocketPath: criblcontrolplanesdkgo.Pointer("<value>"),
            UnixSocketPerms: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthToken: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
| `input`                                                  | [components.Input](../../models/components/input.md)     | :heavy_check_mark:                                       | Input object                                             |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UpdateInputByIDResponse](../../models/operations/updateinputbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete the specified Source.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteInputById" method="delete" path="/system/inputs/{id}" -->
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

    res, err := s.Sources.Delete(ctx, "<id>")
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteInputByIDResponse](../../models/operations/deleteinputbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |