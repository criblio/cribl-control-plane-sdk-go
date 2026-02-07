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

### Example Usage: InputCreateExamplesAppscope

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesAppscope" -->
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
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
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
### Example Usage: InputCreateExamplesAzureBlob

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesAzureBlob" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyAzureBlob(
        operations.CreateInputSystemByPackInputAzureBlob{
            ID: "azure-blob-source",
            Type: operations.CreateInputSystemByPackTypeAzureBlobAzureBlob,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "azure-blob-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("<value>"),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](1300.01),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](5230.25),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](9524.03),
            ServicePeriodSecs: criblcontrolplanesdkgo.Pointer[float64](3487.96),
            SkipOnError: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](2705.97),
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](473.24),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](6975.07),
            AuthType: components.AuthenticationMethodOptionsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("aw woot what how paltry fondly"),
            ConnectionString: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            StorageAccountName: criblcontrolplanesdkgo.Pointer("<value>"),
            TenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            ClientID: criblcontrolplanesdkgo.Pointer("<id>"),
            AzureCloud: criblcontrolplanesdkgo.Pointer("<value>"),
            EndpointSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            ClientTextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Certificate: &components.CertificateTypeAzureBlobAuthTypeClientCert{
                CertificateName: "<value>",
            },
            TemplateQueueName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateConnectionString: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateTenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateClientID: criblcontrolplanesdkgo.Pointer("<id>"),
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
### Example Usage: InputCreateExamplesCloudflareHec

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesCloudflareHec" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyCloudflareHec(
        operations.CreateInputSystemByPackInputCloudflareHec{
            ID: "cloudflare-hec-source",
            Type: operations.CreateInputSystemByPackTypeCloudflareHecCloudflareHec,
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 8088,
            AuthTokens: []operations.CreateInputSystemByPackAuthTokenCloudflareHec{
                operations.CreateInputSystemByPackAuthTokenCloudflareHec{
                    AuthType: operations.CreateInputSystemByPackAuthTokenAuthenticationMethodSecret.ToPointer(),
                    TokenSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Token: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("consequently meanwhile stormy after pfft willow corrupt athwart as thoughtfully"),
                    AllowedIndexesAtToken: []string{
                        "<value 1>",
                        "<value 2>",
                        "<value 3>",
                    },
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
            TLS: &operations.CreateInputSystemByPackTLSSettingsServerSide{
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](6512.35),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](679079),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](9591.92),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](6094.6),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](4715.53),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](9357.22),
            EnableHealthCheck: "<value>",
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            HecAPI: "/services/collector",
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AllowedIndexes: []string{
                "<value 1>",
                "<value 2>",
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](9219.13),
            AccessControlAllowOrigin: []string{
                "<value 1>",
                "<value 2>",
            },
            AccessControlAllowHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            EmitTokenMetrics: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("ravage perfectly perfection by"),
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
### Example Usage: InputCreateExamplesCollection

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesCollection" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyCollection(
        operations.CreateInputSystemByPackInputCollection{
            ID: "collection-source",
            Type: operations.CreateInputSystemByPackTypeCollectionCollection,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](6997.78),
            Preprocess: &components.PreprocessType{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                },
            },
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Output: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesConfluentCloud

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesConfluentCloud" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyConfluentCloud(
        operations.CreateInputSystemByPackInputConfluentCloud{
            ID: "confluent-cloud-source",
            Type: operations.CreateInputSystemByPackTypeConfluentCloudConfluentCloud,
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Brokers: []string{
                "pkc-xxxxx.us-east-1.aws.confluent.cloud:9092",
            },
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            Topics: []string{
                "logs",
            },
            GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
            FromBeginning: criblcontrolplanesdkgo.Pointer(true),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType{
                Disabled: false,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://whispered-tenant.biz"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](9267.71),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](3003.41),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](2353.64),
                Auth: &components.AuthTypeKafkaSchemaRegistry{
                    Disabled: true,
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                    CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                },
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](1224.48),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](2839.68),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](7884),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](755.04),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5844.08),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8812.41),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](8013.48),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](7182.09),
            Sasl: &components.AuthenticationType{
                Disabled: false,
                Username: criblcontrolplanesdkgo.Pointer("Cleta65"),
                Password: criblcontrolplanesdkgo.Pointer("cHVFBWuIATR5WfJ"),
                AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslPlain.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(false),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://gorgeous-marketplace.info/"),
                ClientID: criblcontrolplanesdkgo.Pointer("<id>"),
                OauthSecretType: criblcontrolplanesdkgo.Pointer("<value>"),
                ClientTextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthParams: []components.ItemsTypeSaslOauthParams{
                    components.ItemsTypeSaslOauthParams{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                SaslExtensions: []components.ItemsTypeSaslSaslExtensions{
                    components.ItemsTypeSaslSaslExtensions{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
            },
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](2225.59),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](6180.1),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](1597.69),
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](2161.53),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](1571.03),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](3391.74),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](2488.61),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](9103.56),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("dandelion loyally graceful advertisement embed"),
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
### Example Usage: InputCreateExamplesCriblHttp

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesCriblHttp" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyCriblHTTP(
        operations.CreateInputSystemByPackInputCriblHTTP{
            ID: "cribl-http-source",
            Type: operations.CreateInputSystemByPackTypeCriblHTTPCriblHTTP,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Host: "0.0.0.0",
            Port: 10080,
            AuthTokens: []components.ItemsTypeAuthTokens{
                components.ItemsTypeAuthTokens{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                    Description: criblcontrolplanesdkgo.Pointer("while onto optimistically"),
                },
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](2037.88),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](426932),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](1112.43),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](6983.36),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](8499.52),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](3693.39),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("whoever spear incidentally ceramics endow underneath"),
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
### Example Usage: InputCreateExamplesCriblLakeHttp

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesCriblLakeHttp" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyCriblLakeHTTP(
        operations.CreateInputSystemByPackInputCriblLakeHTTP{
            ID: "cribl-lake-http-source",
            Type: operations.CreateInputSystemByPackTypeCriblLakeHTTPCriblLakeHTTP,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Host: "0.0.0.0",
            Port: 10080,
            AuthTokens: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](6139.72),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](758726),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](1986.26),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](3150.95),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](8728.17),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](3056.32),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            CriblAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            ElasticAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            SplunkHecAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            SplunkHecAcks: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthTokensExt: []operations.CreateInputSystemByPackAuthTokensExt{
                operations.CreateInputSystemByPackAuthTokensExt{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("whether modulo within woot soft faithfully lounge lightly twine deeply"),
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                    SplunkHecMetadata: &operations.CreateInputSystemByPackSplunkHecMetadata{
                        Enabled: criblcontrolplanesdkgo.Pointer(true),
                        DefaultDataset: criblcontrolplanesdkgo.Pointer("<value>"),
                        AllowedIndexesAtToken: []string{
                            "<value 1>",
                            "<value 2>",
                        },
                    },
                    ElasticsearchMetadata: &operations.CreateInputSystemByPackElasticsearchMetadata{
                        Enabled: criblcontrolplanesdkgo.Pointer(false),
                        DefaultDataset: criblcontrolplanesdkgo.Pointer("<value>"),
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("competent blah athwart"),
            TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateSplunkHecAPI: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesCriblTcp

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesCriblTcp" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyCriblTCP(
        operations.CreateInputSystemByPackInputCriblTCP{
            ID: "cribl-tcp-source",
            Type: operations.CreateInputSystemByPackTypeCriblTCPCriblTCP,
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10090,
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
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](3010.37),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](7092.34),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](1608),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](284.46),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(true),
            AuthTokens: []components.ItemsTypeAuthTokens{
                components.ItemsTypeAuthTokens{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                    Description: criblcontrolplanesdkgo.Pointer("while onto optimistically"),
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("unless around bah via interestingly reckless excluding till"),
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
### Example Usage: InputCreateExamplesCrowdstrike

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesCrowdstrike" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyCrowdstrike(
        operations.CreateInputSystemByPackInputCrowdstrike{
            ID: "crowdstrike-source",
            Type: operations.CreateInputSystemByPackTypeCrowdstrikeCrowdstrike,
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "crowdstrike-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfManual.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](399.49),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](8799.77),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](8015.03),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](1114.44),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](4966.7),
            SkipOnError: criblcontrolplanesdkgo.Pointer(true),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](9794.09),
            EnableSQSAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            Preprocess: &components.PreprocessType{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Checkpointing: &components.CheckpointingType{
                Enabled: true,
                Retries: criblcontrolplanesdkgo.Pointer[float64](9654.85),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](4339.54),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("ha shout reluctantly upward meh"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TagAfterProcessing: components.TagAfterProcessingOptionsFalse.ToPointer(),
            ProcessedTagKey: criblcontrolplanesdkgo.Pointer("<value>"),
            ProcessedTagValue: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateQueueName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesDatadogAgent

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesDatadogAgent" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyDatadogAgent(
        operations.CreateInputSystemByPackInputDatadogAgent{
            ID: "datadog-agent-source",
            Type: operations.CreateInputSystemByPackTypeDatadogAgentDatadogAgent,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 8126,
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](7677.42),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](622686),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](4227.76),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](6896.39),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](8316.26),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](9658.8),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ProxyMode: &operations.CreateInputSystemByPackProxyModeDatadogAgent{
                Enabled: true,
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            },
            Description: criblcontrolplanesdkgo.Pointer("prickly across sleepily"),
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
### Example Usage: InputCreateExamplesDatagen

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesDatagen" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyDatagen(
        operations.CreateInputSystemByPackInputDatagen{
            ID: "datagen-source",
            Type: operations.CreateInputSystemByPackTypeDatagenDatagen,
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Samples: []operations.CreateInputSystemByPackSample{
                operations.CreateInputSystemByPackSample{
                    Sample: "sample.json",
                    EventsPerSec: 10,
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("scram intrigue mindless wicked some account whack digitize instead"),
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
### Example Usage: InputCreateExamplesEdgePrometheus

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesEdgePrometheus" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyEdgePrometheus(
        operations.CreateInputSystemByPackInputEdgePrometheus{
            ID: "edge-prometheus-source",
            Type: operations.CreateInputSystemByPackTypeEdgePrometheusEdgePrometheus,
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            DimensionList: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            DiscoveryType: operations.CreateInputSystemByPackDiscoveryTypeEdgePrometheusStatic,
            Interval: 60,
            Timeout: criblcontrolplanesdkgo.Pointer[float64](2723.41),
            Persistence: &components.DiskSpoolingType{
                Enable: criblcontrolplanesdkgo.Pointer(true),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.CompressionOptionsPersistenceGzip.ToPointer(),
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthType: operations.CreateInputSystemByPackAuthenticationMethodEdgePrometheusSecret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("jacket however than feather worth"),
            Targets: []operations.CreateInputSystemByPackTarget{
                operations.CreateInputSystemByPackTarget{
                    Protocol: components.ProtocolOptionsTargetsItemsHTTP.ToPointer(),
                    Host: "localhost",
                    Port: criblcontrolplanesdkgo.Pointer[float64](2109.95),
                    Path: criblcontrolplanesdkgo.Pointer("/var/yp"),
                },
            },
            RecordType: components.RecordTypeOptionsAaaa.ToPointer(),
            ScrapePort: criblcontrolplanesdkgo.Pointer[float64](2432.75),
            NameList: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ScrapeProtocol: components.ProtocolOptionsTargetsItemsHTTP.ToPointer(),
            ScrapePath: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfSecret.ToPointer(),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            UsePublicIP: criblcontrolplanesdkgo.Pointer(false),
            SearchFilter: []components.ItemsTypeSearchFilter{
                components.ItemsTypeSearchFilter{
                    Name: "<value>",
                    Values: []string{},
                },
            },
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions1V4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](4088.79),
            ScrapeProtocolExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            ScrapePortExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            ScrapePathExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            PodFilter: []operations.CreateInputSystemByPackPodFilter{
                operations.CreateInputSystemByPackPodFilter{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("depute roughly delightfully behind scratch confide"),
                },
            },
            Username: criblcontrolplanesdkgo.Pointer("Tierra.Robel61"),
            Password: criblcontrolplanesdkgo.Pointer("h4KF6fMhw8MSDfs"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
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
### Example Usage: InputCreateExamplesElastic

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesElastic" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyElastic(
        operations.CreateInputSystemByPackInputElastic{
            ID: "elastic-source",
            Type: operations.CreateInputSystemByPackTypeElasticElastic,
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
            Host: "localhost",
            Port: 9200,
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](8146.25),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](935089),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](2256.18),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](4023.32),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](6972.5),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](8055.81),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            ElasticAPI: "/",
            AuthType: operations.CreateInputSystemByPackAuthenticationTypeElasticNone.ToPointer(),
            APIVersion: operations.CreateInputSystemByPackAPIVersionEightDot3Dot2.ToPointer(),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ProxyMode: &operations.CreateInputSystemByPackProxyModeElastic{
                Enabled: false,
                AuthType: operations.CreateInputSystemByPackProxyModeAuthenticationMethodNone.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("Kendra.Boyle23"),
                Password: criblcontrolplanesdkgo.Pointer("D2e7BSQLaiQdZTl"),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                URL: criblcontrolplanesdkgo.Pointer("https://somber-lifestyle.net/"),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                RemoveHeaders: []string{
                    "<value 1>",
                },
                TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](3117.64),
                TemplateURL: criblcontrolplanesdkgo.Pointer("https://great-carnival.info"),
            },
            Description: criblcontrolplanesdkgo.Pointer("amid loudly aw warmly sunder"),
            Username: criblcontrolplanesdkgo.Pointer("Roosevelt_Harvey"),
            Password: criblcontrolplanesdkgo.Pointer("IH6wSIiQ9TSKVRp"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthTokens: []string{
                "<value 1>",
                "<value 2>",
            },
            CustomAPIVersion: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesEventhub

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesEventhub" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyEventhub(
        operations.CreateInputSystemByPackInputEventhub{
            ID: "eventhub-source",
            Type: operations.CreateInputSystemByPackTypeEventhubEventhub,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Brokers: []string{
                "myeventhub.servicebus.windows.net:9093",
            },
            Topics: []string{
                "logs",
            },
            GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
            FromBeginning: criblcontrolplanesdkgo.Pointer(false),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](3380.61),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](5386.27),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](4499.43),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](9459.16),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8116.71),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2300.62),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](2678.27),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](4698.72),
            Sasl: &components.AuthenticationType1{
                Disabled: false,
                AuthType: components.AuthenticationMethodOptionsSasl1Manual.ToPointer(),
                Password: criblcontrolplanesdkgo.Pointer("DCPZ4YAq7I_kSEB"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSasl1Plain.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("Nelda_Schimmel"),
                ClientSecretAuthType: components.AuthenticationMethodOptionsSasl2Secret.ToPointer(),
                ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                ClientTextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEndpoint: components.MicrosoftEntraIDAuthenticationEndpointOptionsSaslHTTPSLoginPartnerMicrosoftonlineCn.ToPointer(),
                ClientID: criblcontrolplanesdkgo.Pointer("<id>"),
                TenantID: criblcontrolplanesdkgo.Pointer("<id>"),
                Scope: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            TLS: &components.TLSSettingsClientSideType{
                Disabled: false,
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            },
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](2076.59),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](1206.91),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](6706.77),
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](5810.7),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](3246.62),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](188.46),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](596.83),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](1607.63),
            MinimizeDuplicates: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("where correctly yet now"),
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
### Example Usage: InputCreateExamplesExec

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesExec" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyExec(
        operations.CreateInputSystemByPackInputExec{
            ID: "exec-source",
            Type: operations.CreateInputSystemByPackInputExecTypeExec,
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
            Command: "echo \"Hello World\"",
            Retries: criblcontrolplanesdkgo.Pointer[float64](5952.7),
            ScheduleType: operations.CreateInputSystemByPackScheduleTypeCronSchedule.ToPointer(),
            BreakerRulesets: []string{
                "<value 1>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](9068.17),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("sometimes especially microchip amend abaft for shrill hydrolyze minus"),
            Interval: criblcontrolplanesdkgo.Pointer[float64](60),
            CronSchedule: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesFile

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesFile" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyFile(
        operations.CreateInputSystemByPackInputFile{
            ID: "file-source",
            Type: operations.CreateInputSystemByPackInputFileTypeFile,
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
            Mode: operations.CreateInputSystemByPackInputFileModeManual.ToPointer(),
            Interval: criblcontrolplanesdkgo.Pointer[float64](4675.92),
            Filenames: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            FilterArchivedFiles: criblcontrolplanesdkgo.Pointer(true),
            TailOnly: criblcontrolplanesdkgo.Pointer(true),
            IdleTimeout: criblcontrolplanesdkgo.Pointer[float64](1871.84),
            MinAgeDur: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxAgeDur: criblcontrolplanesdkgo.Pointer("<value>"),
            CheckFileModTime: criblcontrolplanesdkgo.Pointer(true),
            ForceText: criblcontrolplanesdkgo.Pointer(true),
            HashLen: criblcontrolplanesdkgo.Pointer[float64](7884.41),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](1656.19),
            Description: criblcontrolplanesdkgo.Pointer("gosh mysteriously inculcate guzzle"),
            Path: criblcontrolplanesdkgo.Pointer("/usr/include"),
            Depth: criblcontrolplanesdkgo.Pointer[float64](6009.53),
            SuppressMissingPathErrors: criblcontrolplanesdkgo.Pointer(true),
            DeleteFiles: criblcontrolplanesdkgo.Pointer(false),
            SaltHash: criblcontrolplanesdkgo.Pointer(true),
            IncludeUnidentifiableBinary: criblcontrolplanesdkgo.Pointer(true),
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
### Example Usage: InputCreateExamplesFirehose

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesFirehose" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyFirehose(
        operations.CreateInputSystemByPackInputFirehose{
            ID: "firehose-source",
            Type: operations.CreateInputSystemByPackTypeFirehoseFirehose,
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Host: "0.0.0.0",
            Port: 10080,
            AuthTokens: []string{
                "<value 1>",
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](3142.84),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](955945),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](4010.69),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](3023.14),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](6368.71),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](2167.16),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("solemnly beside uh-huh likewise"),
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
### Example Usage: InputCreateExamplesGooglePubsub

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesGooglePubsub" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyGooglePubsub(
        operations.CreateInputSystemByPackInputGooglePubsub{
            ID: "google-pubsub-source",
            Type: operations.CreateInputSystemByPackTypeGooglePubsubGooglePubsub,
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            TopicName: "my-topic",
            SubscriptionName: "my-subscription",
            MonitorSubscription: criblcontrolplanesdkgo.Pointer(true),
            CreateTopic: criblcontrolplanesdkgo.Pointer(false),
            CreateSubscription: criblcontrolplanesdkgo.Pointer(true),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            GoogleAuthMethod: components.GoogleAuthenticationMethodOptionsAuto.ToPointer(),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxBacklog: criblcontrolplanesdkgo.Pointer[float64](6431.72),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](2906.9),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](540.55),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("since connect from hm toward"),
            OrderedDelivery: criblcontrolplanesdkgo.Pointer(true),
            TemplateTopicName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateSubscriptionName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesGrafana

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesGrafana" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyGrafana(
        operations.CreateCreateInputSystemByPackInputGrafanaUnionCreateInputSystemByPackInputGrafanaGrafana1(
            operations.CreateInputSystemByPackInputGrafanaGrafana1{
                ID: "grafana-source",
                Type: operations.CreateInputSystemByPackInputGrafanaType1Grafana,
                Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                    Mode: components.ModeOptionsPqAlways.ToPointer(),
                    MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                    CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                    MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    Path: criblcontrolplanesdkgo.Pointer("/opt"),
                    Compress: components.CompressionOptionsPqGzip.ToPointer(),
                    PqControls: &components.PqTypePqControls{},
                },
                Host: "0.0.0.0",
                Port: 10080,
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
                MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](4438.1),
                MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](307170),
                EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
                CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
                ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](4876.48),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](426.95),
                SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](6721.71),
                KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](2204.22),
                EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
                IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
                IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
                PrometheusAPI: "/api/prom/push",
                LokiAPI: criblcontrolplanesdkgo.Pointer("<value>"),
                PrometheusAuth: &operations.CreateInputSystemByPackPrometheusAuth1{
                    AuthType: components.AuthenticationTypeOptionsPrometheusAuthTextSecret.ToPointer(),
                    Username: criblcontrolplanesdkgo.Pointer("Jaydon47"),
                    Password: criblcontrolplanesdkgo.Pointer("tGyFIJt0S4YWEdr"),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                LokiAuth: &operations.CreateInputSystemByPackLokiAuth1{
                    AuthType: components.AuthenticationTypeOptionsLokiAuthCredentialsSecret.ToPointer(),
                    Username: criblcontrolplanesdkgo.Pointer("Paolo_Bode54"),
                    Password: criblcontrolplanesdkgo.Pointer("svqKZTSbPBO1aP5"),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                Metadata: []components.ItemsTypeMetadata{
                    components.ItemsTypeMetadata{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                Description: criblcontrolplanesdkgo.Pointer("connect divert yippee upliftingly"),
                TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
            },
        ),
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInput != nil {
        // handle response
    }
}
```
### Example Usage: InputCreateExamplesHttp

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesHttp" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyHTTP(
        operations.CreateInputSystemByPackInputHTTP{
            ID: "http-source",
            Type: operations.CreateInputSystemByPackTypeHTTPHTTP,
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Host: "0.0.0.0",
            Port: 10080,
            AuthTokens: []string{
                "<value 1>",
                "<value 2>",
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](302.66),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](773604),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](3413.93),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7022.73),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](6410.62),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](3667.37),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            CriblAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            ElasticAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            SplunkHecAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            SplunkHecAcks: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthTokensExt: []components.ItemsTypeAuthTokensExt{
                components.ItemsTypeAuthTokensExt{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("beside incidentally helpfully pfft where concerned cute though"),
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("effector tinted longingly but reflate"),
            TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateSplunkHecAPI: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesHttpRaw

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesHttpRaw" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyHTTPRaw(
        operations.CreateInputSystemByPackInputHTTPRaw{
            ID: "http-raw-source",
            Type: operations.CreateInputSystemByPackTypeHTTPRawHTTPRaw,
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10080,
            AuthTokens: []string{
                "<value 1>",
                "<value 2>",
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](5088.84),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](756062),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](4982.95),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7715.91),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](6053.41),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](3626.73),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](7415.55),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AllowedPaths: []string{
                "<value 1>",
            },
            AllowedMethods: []string{
                "<value 1>",
                "<value 2>",
            },
            AuthTokensExt: []components.ItemsTypeAuthTokensExt{
                components.ItemsTypeAuthTokensExt{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("beside incidentally helpfully pfft where concerned cute though"),
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("provision fervently whoever within perfumed after eyebrow toothpick suddenly"),
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
### Example Usage: InputCreateExamplesJournalFiles

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesJournalFiles" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyJournalFiles(
        operations.CreateInputSystemByPackInputJournalFiles{
            ID: "journal-files-source",
            Type: operations.CreateInputSystemByPackInputJournalFilesTypeJournalFiles,
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Path: "/var/log/journal",
            Interval: criblcontrolplanesdkgo.Pointer[float64](7548.42),
            Journals: []string{
                "system",
            },
            Rules: []operations.CreateInputSystemByPackInputJournalFilesRule{
                operations.CreateInputSystemByPackInputJournalFilesRule{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("yippee even intellect ouch captain sans"),
                },
            },
            CurrentBoot: criblcontrolplanesdkgo.Pointer(true),
            MaxAgeDur: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("um idolized atomize"),
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
### Example Usage: InputCreateExamplesKafka

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesKafka" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyKafka(
        operations.CreateInputSystemByPackInputKafka{
            ID: "kafka-source",
            Type: operations.CreateInputSystemByPackTypeKafkaKafka,
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Brokers: []string{
                "localhost:9092",
            },
            Topics: []string{
                "logs",
            },
            GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
            FromBeginning: criblcontrolplanesdkgo.Pointer(true),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType{
                Disabled: false,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://whispered-tenant.biz"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](9267.71),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](3003.41),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](2353.64),
                Auth: &components.AuthTypeKafkaSchemaRegistry{
                    Disabled: true,
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                    CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                },
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](4621.97),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](5155.41),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](3618.58),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](982.45),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2.32),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7861.13),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](8805.41),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](6939.52),
            Sasl: &components.AuthenticationType{
                Disabled: false,
                Username: criblcontrolplanesdkgo.Pointer("Cleta65"),
                Password: criblcontrolplanesdkgo.Pointer("cHVFBWuIATR5WfJ"),
                AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslPlain.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(false),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://gorgeous-marketplace.info/"),
                ClientID: criblcontrolplanesdkgo.Pointer("<id>"),
                OauthSecretType: criblcontrolplanesdkgo.Pointer("<value>"),
                ClientTextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthParams: []components.ItemsTypeSaslOauthParams{
                    components.ItemsTypeSaslOauthParams{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                SaslExtensions: []components.ItemsTypeSaslSaslExtensions{
                    components.ItemsTypeSaslSaslExtensions{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
            },
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](9720.45),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](9699.16),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](7946.94),
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](7362.72),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](521.6),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](9012.3),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](6175.18),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](9416.76),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("entomb hm supportive mysterious because per"),
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
### Example Usage: InputCreateExamplesKinesis

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesKinesis" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyKinesis(
        operations.CreateInputSystemByPackInputKinesis{
            ID: "kinesis-source",
            Type: operations.CreateInputSystemByPackTypeKinesisKinesis,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            StreamName: "my-stream",
            ServiceInterval: criblcontrolplanesdkgo.Pointer[float64](3703.04),
            ShardExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            ShardIteratorType: operations.CreateInputSystemByPackShardIteratorStartTrimHorizon.ToPointer(),
            PayloadFormat: operations.CreateInputSystemByPackRecordDataFormatNdjson.ToPointer(),
            GetRecordsLimit: criblcontrolplanesdkgo.Pointer[float64](3184.84),
            GetRecordsLimitTotal: criblcontrolplanesdkgo.Pointer[float64](9046.79),
            LoadBalancingAlgorithm: operations.CreateInputSystemByPackShardLoadBalancingRoundRobin.ToPointer(),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions2V4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](8280.04),
            VerifyKPLCheckSums: criblcontrolplanesdkgo.Pointer(true),
            AvoidDuplicates: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("clinking tennis solemnly willfully judgementally upliftingly yuck"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateStreamName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesKubeEvents

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesKubeEvents" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyKubeEvents(
        operations.CreateInputSystemByPackInputKubeEvents{
            ID: "kube-events-source",
            Type: operations.CreateInputSystemByPackTypeKubeEventsKubeEvents,
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Rules: []components.ItemsTypeRules{
                components.ItemsTypeRules{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("until jovially tense concerning playfully inside psst"),
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("nearly married electric finished forgather"),
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
### Example Usage: InputCreateExamplesKubeLogs

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesKubeLogs" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyKubeLogs(
        operations.CreateInputSystemByPackInputKubeLogs{
            ID: "kube-logs-source",
            Type: operations.CreateInputSystemByPackTypeKubeLogsKubeLogs,
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](5360.43),
            Rules: []operations.CreateInputSystemByPackRuleKubeLogs{
                operations.CreateInputSystemByPackRuleKubeLogs{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("unimpressively per gape whoever witty"),
                },
            },
            Timestamps: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Persistence: &components.DiskSpoolingType{
                Enable: criblcontrolplanesdkgo.Pointer(true),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.CompressionOptionsPersistenceGzip.ToPointer(),
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](8648.79),
            EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("doodle through unto doubtfully"),
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
### Example Usage: InputCreateExamplesKubeMetrics

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesKubeMetrics" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyKubeMetrics(
        operations.CreateInputSystemByPackInputKubeMetrics{
            ID: "kube-metrics-source",
            Type: operations.CreateInputSystemByPackTypeKubeMetricsKubeMetrics,
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Interval: criblcontrolplanesdkgo.Pointer[float64](6774.01),
            Rules: []components.ItemsTypeRules{
                components.ItemsTypeRules{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("until jovially tense concerning playfully inside psst"),
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Persistence: &operations.CreateInputSystemByPackPersistenceKubeMetrics{
                Enable: criblcontrolplanesdkgo.Pointer(true),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            Description: criblcontrolplanesdkgo.Pointer("deck why save portly phooey"),
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
### Example Usage: InputCreateExamplesLoki

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesLoki" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyLoki(
        operations.CreateInputSystemByPackInputLoki{
            ID: "loki-source",
            Type: operations.CreateInputSystemByPackTypeLokiLoki,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10080,
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](6182.53),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](155987),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](2973.11),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](6216.26),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](1396.24),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](8913.6),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            LokiAPI: "/loki/api/v1/push",
            AuthType: components.AuthenticationTypeOptionsLokiAuthToken.ToPointer(),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("pace marketplace usefully restfully yet"),
            Username: criblcontrolplanesdkgo.Pointer("Sarah.Hartmann45"),
            Password: criblcontrolplanesdkgo.Pointer("Ks0TpSSaC_SEGSx"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesMetrics

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesMetrics" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyMetrics(
        operations.CreateInputSystemByPackInputMetrics{
            ID: "metrics-source",
            Type: operations.CreateInputSystemByPackTypeMetricsMetrics,
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
            Host: "0.0.0.0",
            UDPPort: criblcontrolplanesdkgo.Pointer[float64](8125),
            TCPPort: criblcontrolplanesdkgo.Pointer[float64](395.37),
            MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1774.9),
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
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
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](793.23),
            Description: criblcontrolplanesdkgo.Pointer("handle deceivingly as gee wholly ha square schlep mmm sans"),
            TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateUDPPort: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateTCPPort: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesModelDrivenTelemetry

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesModelDrivenTelemetry" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyModelDrivenTelemetry(
        operations.CreateInputSystemByPackInputModelDrivenTelemetry{
            ID: "mdt-source",
            Type: operations.CreateInputSystemByPackTypeModelDrivenTelemetryModelDrivenTelemetry,
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
            Host: "0.0.0.0",
            Port: 57000,
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
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](427.15),
            ShutdownTimeoutMs: criblcontrolplanesdkgo.Pointer[float64](9059.89),
            Description: criblcontrolplanesdkgo.Pointer("attribute excepting duh"),
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
### Example Usage: InputCreateExamplesMsk

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesMsk" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyMsk(
        operations.CreateInputSystemByPackInputMsk{
            ID: "msk-source",
            Type: operations.CreateInputSystemByPackTypeMskMsk,
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
            Brokers: []string{
                "b-1.example.xxxxx.c2.kafka.us-east-1.amazonaws.com:9092",
            },
            Topics: []string{
                "logs",
            },
            GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
            FromBeginning: criblcontrolplanesdkgo.Pointer(true),
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](4131.83),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](2415.81),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](5806.57),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType{
                Disabled: false,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://whispered-tenant.biz"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](9267.71),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](3003.41),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](2353.64),
                Auth: &components.AuthTypeKafkaSchemaRegistry{
                    Disabled: true,
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                    CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                },
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](8780.38),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](8587.64),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](1558.78),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](2081.32),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](3106.55),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3892.36),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](3410.14),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](6033.21),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto,
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](2987.55),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](7928.25),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](5480.06),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](4767.67),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](5592.47),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](2244.45),
            Description: criblcontrolplanesdkgo.Pointer("pluck ack which pfft afore populist"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesNetflow

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesNetflow" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyNetflow(
        operations.CreateInputSystemByPackInputNetflow{
            ID: "netflow-source",
            Type: operations.CreateInputSystemByPackTypeNetflowNetflow,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Host: "0.0.0.0",
            Port: 2055,
            EnablePassThrough: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](6270.31),
            TemplateCacheMinutes: criblcontrolplanesdkgo.Pointer[float64](6537.85),
            V5Enabled: criblcontrolplanesdkgo.Pointer(true),
            V9Enabled: criblcontrolplanesdkgo.Pointer(true),
            IpfixEnabled: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("woefully hassle fooey quixotic incidentally outgoing messy ouch circa"),
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
### Example Usage: InputCreateExamplesOffice365Mgmt

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesOffice365Mgmt" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyOffice365Mgmt(
        operations.CreateInputSystemByPackInputOffice365Mgmt{
            ID: "office365-mgmt-source",
            Type: operations.CreateInputSystemByPackTypeOffice365MgmtOffice365Mgmt,
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
            PlanType: components.SubscriptionPlanOptionsEnterpriseGcc,
            TenantID: "tenant-id",
            AppID: "app-id",
            Timeout: criblcontrolplanesdkgo.Pointer[float64](2972.66),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](2696.03),
            JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](7253.42),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            PublisherIdentifier: criblcontrolplanesdkgo.Pointer("<value>"),
            ContentConfig: []operations.CreateInputSystemByPackContentConfigOffice365Mgmt{
                operations.CreateInputSystemByPackContentConfigOffice365Mgmt{
                    ContentType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Description: criblcontrolplanesdkgo.Pointer("knuckle twist opposite"),
                    Interval: criblcontrolplanesdkgo.Pointer[float64](2830.93),
                    LogLevel: components.LogLevelOptionsContentConfigItemsError.ToPointer(),
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                },
            },
            IngestionLag: criblcontrolplanesdkgo.Pointer[float64](5013.97),
            RetryRules: &components.RetryRulesType1{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesBackoff,
                Interval: criblcontrolplanesdkgo.Pointer[float64](1360.25),
                Limit: criblcontrolplanesdkgo.Pointer[float64](2039.55),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](2148.96),
                Codes: []float64{
                    1001.17,
                    8763.82,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            AuthType: components.AuthenticationMethodOptions1Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("alert humidity ham whenever till whoever bicycle abnormally vice"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateTenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAppID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplatePublisherIdentifier: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesOffice365MsgTrace

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesOffice365MsgTrace" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyOffice365MsgTrace(
        operations.CreateInputSystemByPackInputOffice365MsgTrace{
            ID: "office365-msg-trace-source",
            Type: operations.CreateInputSystemByPackTypeOffice365MsgTraceOffice365MsgTrace,
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            URL: "https://reports.office365.com/ecp/reportingwebservice/reporting.svc/MessageTrace",
            Interval: 15,
            StartDate: criblcontrolplanesdkgo.Pointer("<value>"),
            EndDate: criblcontrolplanesdkgo.Pointer("<value>"),
            Timeout: criblcontrolplanesdkgo.Pointer[float64](267.65),
            DisableTimeFilter: criblcontrolplanesdkgo.Pointer(true),
            AuthType: operations.CreateInputSystemByPackAuthenticationMethodOffice365MsgTraceOauth.ToPointer(),
            RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
            MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](3220.01),
            LogLevel: operations.CreateInputSystemByPackLogLevelOffice365MsgTraceWarn.ToPointer(),
            JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](3233.18),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](1313.3),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            RetryRules: &components.RetryRulesType1{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesBackoff,
                Interval: criblcontrolplanesdkgo.Pointer[float64](1360.25),
                Limit: criblcontrolplanesdkgo.Pointer[float64](2039.55),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](2148.96),
                Codes: []float64{
                    1001.17,
                    8763.82,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            Description: criblcontrolplanesdkgo.Pointer("blah tremendously rural folklore responsible which"),
            Username: criblcontrolplanesdkgo.Pointer("Novella.Zieme"),
            Password: criblcontrolplanesdkgo.Pointer("EqzTKJKZW7lLPjz"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            ClientID: criblcontrolplanesdkgo.Pointer("<id>"),
            Resource: criblcontrolplanesdkgo.Pointer("<value>"),
            PlanType: components.SubscriptionPlanOptionsGccHigh.ToPointer(),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            CertOptions: &operations.CreateInputSystemByPackCertOptions{
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: "<value>",
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: "<value>",
            },
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://nocturnal-blight.info/"),
            TemplateTenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateClientID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateResource: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesOffice365Service

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesOffice365Service" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyOffice365Service(
        operations.CreateInputSystemByPackInputOffice365Service{
            ID: "office365-service-source",
            Type: operations.CreateInputSystemByPackTypeOffice365ServiceOffice365Service,
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            PlanType: components.SubscriptionPlanOptionsEnterpriseGcc.ToPointer(),
            TenantID: "tenant-id",
            AppID: "app-id",
            Timeout: criblcontrolplanesdkgo.Pointer[float64](4315.93),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](9735.61),
            JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](3387.92),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ContentConfig: []operations.CreateInputSystemByPackContentConfigOffice365Service{
                operations.CreateInputSystemByPackContentConfigOffice365Service{
                    ContentType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Description: criblcontrolplanesdkgo.Pointer("since encode minion noisily likewise log nice whoever fiercely godparent"),
                    Interval: criblcontrolplanesdkgo.Pointer[float64](2011.47),
                    LogLevel: components.LogLevelOptionsContentConfigItemsDebug.ToPointer(),
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                },
            },
            RetryRules: &components.RetryRulesType1{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesBackoff,
                Interval: criblcontrolplanesdkgo.Pointer[float64](1360.25),
                Limit: criblcontrolplanesdkgo.Pointer[float64](2039.55),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](2148.96),
                Codes: []float64{
                    1001.17,
                    8763.82,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            AuthType: components.AuthenticationMethodOptions1Secret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("unused optimistically after flame freely for save than"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateTenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAppID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesOpenTelemetry

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesOpenTelemetry" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyOpenTelemetry(
        operations.CreateInputSystemByPackInputOpenTelemetry{
            ID: "otel-source",
            Type: operations.CreateInputSystemByPackTypeOpenTelemetryOpenTelemetry,
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 4317,
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](5055.67),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](403392),
            EnableProxyHeader: "<value>",
            CaptureHeaders: "<value>",
            ActivityLogSampleRate: "<value>",
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](1257.27),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](9748.69),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](1466.36),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            Protocol: operations.CreateInputSystemByPackProtocolHTTP.ToPointer(),
            ExtractSpans: criblcontrolplanesdkgo.Pointer(true),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(false),
            OtlpVersion: operations.CreateInputSystemByPackOTLPVersionOneDot3Dot1.ToPointer(),
            AuthType: components.AuthenticationTypeOptionsCredentialsSecret.ToPointer(),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](7521.37),
            Description: criblcontrolplanesdkgo.Pointer("serene whenever successfully or gadzooks hence sticker"),
            Username: criblcontrolplanesdkgo.Pointer("Alexandria68"),
            Password: criblcontrolplanesdkgo.Pointer("dapjfc3A65SEzo6"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ExtractLogs: criblcontrolplanesdkgo.Pointer(false),
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
### Example Usage: InputCreateExamplesPrometheus

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesPrometheus" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyPrometheus(
        operations.CreateInputSystemByPackInputPrometheus{
            ID: "prometheus-source",
            Type: operations.CreateInputSystemByPackTypePrometheusPrometheus,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            DimensionList: []string{
                "<value 1>",
            },
            DiscoveryType: operations.CreateInputSystemByPackDiscoveryTypePrometheusStatic.ToPointer(),
            Interval: 60,
            LogLevel: operations.CreateInputSystemByPackLogLevelPrometheusInfo,
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            Timeout: criblcontrolplanesdkgo.Pointer[float64](818.76),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](5981.93),
            JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](2614.37),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("unto pop by owlishly bob"),
            TargetList: []string{
                "http://localhost:9090/metrics",
            },
            RecordType: components.RecordTypeOptionsAaaa.ToPointer(),
            ScrapePort: criblcontrolplanesdkgo.Pointer[float64](9590.33),
            NameList: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ScrapeProtocol: operations.CreateInputSystemByPackMetricsProtocolHTTP.ToPointer(),
            ScrapePath: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfSecret.ToPointer(),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            UsePublicIP: criblcontrolplanesdkgo.Pointer(false),
            SearchFilter: []components.ItemsTypeSearchFilter{
                components.ItemsTypeSearchFilter{
                    Name: "<value>",
                    Values: []string{},
                },
            },
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions1V2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](9641.84),
            Username: criblcontrolplanesdkgo.Pointer("Sadie66"),
            Password: criblcontrolplanesdkgo.Pointer("MIXMRs1eHGRG9mE"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateLogLevel: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
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
### Example Usage: InputCreateExamplesPrometheusRw

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesPrometheusRw" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyPrometheusRw(
        operations.CreateInputSystemByPackInputPrometheusRw{
            ID: "prometheus-rw-source",
            Type: operations.CreateInputSystemByPackTypePrometheusRwPrometheusRw,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10080,
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](3451.54),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](666276),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](345.03),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](1040.03),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](9328.18),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](4879.06),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            PrometheusAPI: "/write",
            AuthType: components.AuthenticationTypeOptionsPrometheusAuthCredentialsSecret.ToPointer(),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("etch peony deliberately veto reassemble idolized opposite vibraphone"),
            Username: criblcontrolplanesdkgo.Pointer("Rebekah89"),
            Password: criblcontrolplanesdkgo.Pointer("fucs7ltCWdWpggg"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplatePrometheusAPI: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesRawUdp

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesRawUdp" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyRawUDP(
        operations.CreateInputSystemByPackInputRawUDP{
            ID: "raw-udp-source",
            Type: operations.CreateInputSystemByPackTypeRawUDPRawUDP,
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
            Host: "0.0.0.0",
            Port: 514,
            MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1812.93),
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            SingleMsgUDPPackets: criblcontrolplanesdkgo.Pointer(false),
            IngestRawBytes: criblcontrolplanesdkgo.Pointer(false),
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](9029.64),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("huzzah clamor ack afterwards fast converse toward"),
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
### Example Usage: InputCreateExamplesS3

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesS3" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyS3(
        operations.CreateInputSystemByPackInputS3{
            ID: "s3-source",
            Type: operations.CreateInputSystemByPackTypeS3S3,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "s3-notifications-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](7305.53),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](5911.43),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](2807.36),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](8937.56),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](6211.1),
            SkipOnError: criblcontrolplanesdkgo.Pointer(false),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](6237.52),
            EnableSQSAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            Preprocess: &components.PreprocessType{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](8131.77),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](8009.46),
            Checkpointing: &components.CheckpointingType{
                Enabled: true,
                Retries: criblcontrolplanesdkgo.Pointer[float64](9654.85),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](5408.26),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            TagAfterProcessing: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("woefully lanky trench deafening zesty surprised on sunbeam vivaciously"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ProcessedTagKey: criblcontrolplanesdkgo.Pointer("<value>"),
            ProcessedTagValue: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateQueueName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesS3Inventory

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesS3Inventory" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyS3Inventory(
        operations.CreateInputSystemByPackInputS3Inventory{
            ID: "s3-inventory-source",
            Type: operations.CreateInputSystemByPackTypeS3InventoryS3Inventory,
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            QueueName: "s3-inventory-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](1222.88),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](5180.52),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](4259.65),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](110.51),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](5880.6),
            SkipOnError: criblcontrolplanesdkgo.Pointer(false),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3102.38),
            EnableSQSAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            Preprocess: &components.PreprocessType{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](7589.41),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](7741.71),
            Checkpointing: &components.CheckpointingType{
                Enabled: true,
                Retries: criblcontrolplanesdkgo.Pointer[float64](9654.85),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](2413.82),
            ChecksumSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxManifestSizeKB: criblcontrolplanesdkgo.Pointer[int64](157753),
            ValidateInventoryFiles: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("which misjudge cultivated daintily"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TagAfterProcessing: components.TagAfterProcessingOptionsTrue.ToPointer(),
            ProcessedTagKey: criblcontrolplanesdkgo.Pointer("<value>"),
            ProcessedTagValue: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateQueueName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesSecurityLake

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesSecurityLake" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodySecurityLake(
        operations.CreateInputSystemByPackInputSecurityLake{
            ID: "security-lake-source",
            Type: operations.CreateInputSystemByPackTypeSecurityLakeSecurityLake,
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
            QueueName: "security-lake-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfManual.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](3107.21),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](5612.34),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](5057.17),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](9683.62),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](61.29),
            SkipOnError: criblcontrolplanesdkgo.Pointer(false),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](979.79),
            EnableSQSAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            Preprocess: &components.PreprocessType{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](8862.21),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](822.35),
            Checkpointing: &components.CheckpointingType{
                Enabled: true,
                Retries: criblcontrolplanesdkgo.Pointer[float64](9654.85),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](8602.63),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("concerning pfft heartache knowledgeable satirize with"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TagAfterProcessing: components.TagAfterProcessingOptionsTrue.ToPointer(),
            ProcessedTagKey: criblcontrolplanesdkgo.Pointer("<value>"),
            ProcessedTagValue: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateQueueName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesSnmp

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesSnmp" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodySnmp(
        operations.CreateInputSystemByPackInputSnmp{
            ID: "snmp-source",
            Type: operations.CreateInputSystemByPackTypeSnmpSnmp,
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Host: "192.168.1.1",
            Port: 161,
            SnmpV3Auth: &operations.CreateInputSystemByPackSNMPv3Authentication{
                V3AuthEnabled: true,
                AllowUnmatchedTrap: criblcontrolplanesdkgo.Pointer(false),
                V3Users: []operations.CreateInputSystemByPackV3User{
                    operations.CreateInputSystemByPackV3User{
                        Name: "<value>",
                        AuthProtocol: components.AuthenticationProtocolOptionsV3UserMd5.ToPointer(),
                        AuthKey: criblcontrolplanesdkgo.Pointer("<value>"),
                        PrivProtocol: operations.CreateInputSystemByPackPrivacyProtocolAes256b.ToPointer(),
                        PrivKey: criblcontrolplanesdkgo.Pointer("<value>"),
                    },
                },
            },
            MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1460.46),
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](212.18),
            VarbindsWithTypes: criblcontrolplanesdkgo.Pointer(true),
            BestEffortParsing: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("concerning onset vainly obnoxiously whenever duh lazy selfishly"),
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
### Example Usage: InputCreateExamplesSplunk

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesSplunk" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodySplunk(
        operations.CreateInputSystemByPackInputSplunk{
            ID: "splunk-source",
            Type: operations.CreateInputSystemByPackTypeSplunkSplunk,
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
            Host: "0.0.0.0",
            Port: 9997,
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
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](8042.81),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](8701.69),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](4800.75),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](7418.06),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](4573.81),
            AuthTokens: []operations.CreateInputSystemByPackAuthTokenSplunk{
                operations.CreateInputSystemByPackAuthTokenSplunk{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("secret whenever indolent but yuck even aside or turbulent"),
                },
            },
            MaxS2Sversion: operations.CreateInputSystemByPackMaxS2SVersionV4.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("youthfully closely behind for"),
            UseFwdTimezone: criblcontrolplanesdkgo.Pointer(false),
            DropControlFields: criblcontrolplanesdkgo.Pointer(true),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(true),
            Compress: operations.CreateInputSystemByPackCompressionAuto.ToPointer(),
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
### Example Usage: InputCreateExamplesSplunkHec

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesSplunkHec" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodySplunkHec(
        operations.CreateInputSystemByPackInputSplunkHec{
            ID: "splunk-hec-source",
            Type: operations.CreateInputSystemByPackTypeSplunkHecSplunkHec,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Host: "0.0.0.0",
            Port: 8088,
            AuthTokens: []operations.CreateInputSystemByPackAuthTokenSplunkHec{
                operations.CreateInputSystemByPackAuthTokenSplunkHec{
                    AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
                    TokenSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Token: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                    Description: criblcontrolplanesdkgo.Pointer("oh vivacious quick than gah vibration"),
                    AllowedIndexesAtToken: []string{
                        "<value 1>",
                        "<value 2>",
                    },
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](3435.92),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](85570),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](2741.76),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](5414.67),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](369.47),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](4427.77),
            EnableHealthCheck: "<value>",
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            SplunkHecAPI: "/services/collector",
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AllowedIndexes: []string{
                "<value 1>",
                "<value 2>",
            },
            SplunkHecAcks: criblcontrolplanesdkgo.Pointer(true),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](2692.65),
            UseFwdTimezone: criblcontrolplanesdkgo.Pointer(false),
            DropControlFields: criblcontrolplanesdkgo.Pointer(false),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(false),
            AccessControlAllowOrigin: []string{
                "<value 1>",
                "<value 2>",
            },
            AccessControlAllowHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            EmitTokenMetrics: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("fondly substantial tune black-and-white aw banish and"),
            TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateSplunkHecAPI: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesSplunkSearch

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesSplunkSearch" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodySplunkSearch(
        operations.CreateInputSystemByPackInputSplunkSearch{
            ID: "splunk-search-source",
            Type: operations.CreateInputSystemByPackTypeSplunkSearchSplunkSearch,
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            SearchHead: "https://localhost:8089",
            Search: "index=main",
            Earliest: criblcontrolplanesdkgo.Pointer("<value>"),
            Latest: criblcontrolplanesdkgo.Pointer("<value>"),
            CronSchedule: "0 * * * *",
            Endpoint: "/services/search/jobs/export",
            OutputMode: components.OutputModeOptionsSplunkCollectorConfJSON,
            EndpointParams: []operations.CreateInputSystemByPackEndpointParam{
                operations.CreateInputSystemByPackEndpointParam{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            EndpointHeaders: []operations.CreateInputSystemByPackEndpointHeader{
                operations.CreateInputSystemByPackEndpointHeader{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            LogLevel: operations.CreateInputSystemByPackLogLevelSplunkSearchWarn.ToPointer(),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](9887.21),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](5030.92),
            JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](4125.49),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            RetryRules: &components.RetryRulesType{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesNone,
                Interval: criblcontrolplanesdkgo.Pointer[float64](2446.49),
                Limit: criblcontrolplanesdkgo.Pointer[float64](8984.5),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](2926.4),
                Codes: []float64{
                    9786.9,
                    1849.73,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(true),
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](2506.09),
            AuthType: operations.CreateInputSystemByPackAuthenticationTypeSplunkSearchBasic.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("plagiarise enthusiastically phew"),
            Username: criblcontrolplanesdkgo.Pointer("Darren.Murphy46"),
            Password: criblcontrolplanesdkgo.Pointer("HNMkkOA8S0ww7hv"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesSqs

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesSqs" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodySqs(
        operations.CreateInputSystemByPackInputSqs{
            ID: "sqs-source",
            Type: operations.CreateInputSystemByPackTypeSqsSqs,
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
            QueueName: "my-queue",
            QueueType: operations.CreateInputSystemByPackQueueTypeStandard,
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            CreateQueue: criblcontrolplanesdkgo.Pointer(true),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions3V2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](1332.59),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](3646.79),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](2160.76),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](988.49),
            Description: criblcontrolplanesdkgo.Pointer("which more vivaciously woot avalanche disinherit slowly"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](8199.22),
            TemplateQueueName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesSyslog

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesSyslog" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodySyslog(
        operations.CreateCreateInputSystemByPackInputSyslogUnionCreateInputSystemByPackInputSyslogSyslog1(
            operations.CreateInputSystemByPackInputSyslogSyslog1{
                ID: "syslog-source",
                Type: operations.CreateInputSystemByPackInputSyslogType1Syslog,
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
                Host: "0.0.0.0",
                UDPPort: 514,
                TCPPort: criblcontrolplanesdkgo.Pointer[float64](5767.97),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7898.91),
                IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
                TimestampTimezone: criblcontrolplanesdkgo.Pointer("<value>"),
                SingleMsgUDPPackets: criblcontrolplanesdkgo.Pointer(false),
                EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
                KeepFieldsList: []string{
                    "<value 1>",
                    "<value 2>",
                },
                OctetCounting: criblcontrolplanesdkgo.Pointer(false),
                InferFraming: criblcontrolplanesdkgo.Pointer(true),
                StrictlyInferOctetCounting: criblcontrolplanesdkgo.Pointer(false),
                AllowNonStandardAppName: criblcontrolplanesdkgo.Pointer(true),
                MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](8532.26),
                SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](5176.78),
                SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](1688.68),
                SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](7987.09),
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
                Metadata: []components.ItemsTypeMetadata{
                    components.ItemsTypeMetadata{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](40.98),
                EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(false),
                Description: criblcontrolplanesdkgo.Pointer("porter whine ethical pleasure seemingly incidentally irritably"),
                EnableEnhancedProxyHeaderParsing: criblcontrolplanesdkgo.Pointer(true),
                TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                TemplateUDPPort: criblcontrolplanesdkgo.Pointer("<value>"),
                TemplateTCPPort: criblcontrolplanesdkgo.Pointer("<value>"),
            },
        ),
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInput != nil {
        // handle response
    }
}
```
### Example Usage: InputCreateExamplesSystemMetrics

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesSystemMetrics" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodySystemMetrics(
        operations.CreateInputSystemByPackInputSystemMetrics{
            ID: "system-metrics-source",
            Type: operations.CreateInputSystemByPackTypeSystemMetricsSystemMetrics,
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](8247.04),
            Host: &operations.CreateInputSystemByPackHostSystemMetrics{
                Mode: components.ModeOptionsHostAll.ToPointer(),
                Custom: &operations.CreateInputSystemByPackCustomSystemMetrics{
                    System: &operations.CreateInputSystemByPackSystemSystemMetrics{
                        Mode: operations.CreateInputSystemByPackSystemModeSystemMetricsDisabled.ToPointer(),
                        Processes: criblcontrolplanesdkgo.Pointer(false),
                    },
                    CPU: &operations.CreateInputSystemByPackCPUSystemMetrics{
                        Mode: operations.CreateInputSystemByPackCPUModeSystemMetricsDisabled.ToPointer(),
                        PerCPU: criblcontrolplanesdkgo.Pointer(false),
                        Detail: criblcontrolplanesdkgo.Pointer(true),
                        Time: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Memory: &operations.CreateInputSystemByPackMemorySystemMetrics{
                        Mode: operations.CreateInputSystemByPackMemoryModeSystemMetricsBasic.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(true),
                    },
                    Network: &operations.CreateInputSystemByPackNetworkSystemMetrics{
                        Mode: operations.CreateInputSystemByPackNetworkModeSystemMetricsAll.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(true),
                        Protocols: criblcontrolplanesdkgo.Pointer(false),
                        Devices: []string{
                            "<value 1>",
                            "<value 2>",
                        },
                        PerInterface: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Disk: &operations.CreateInputSystemByPackDiskSystemMetrics{
                        Mode: operations.CreateInputSystemByPackDiskModeSystemMetricsAll.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Inodes: criblcontrolplanesdkgo.Pointer(false),
                        Devices: []string{
                            "<value 1>",
                            "<value 2>",
                            "<value 3>",
                        },
                        Mountpoints: []string{
                            "<value 1>",
                        },
                        Fstypes: []string{
                            "<value 1>",
                            "<value 2>",
                            "<value 3>",
                        },
                        PerDevice: criblcontrolplanesdkgo.Pointer(true),
                    },
                },
            },
            Process: &components.ProcessType{
                Sets: []components.ItemsTypeProcessSets{
                    components.ItemsTypeProcessSets{
                        Name: "<value>",
                        Filter: "<value>",
                        IncludeChildren: criblcontrolplanesdkgo.Pointer(true),
                    },
                },
            },
            Container: &operations.CreateInputSystemByPackContainer{
                Mode: operations.CreateInputSystemByPackContainerModeAll.ToPointer(),
                DockerSocket: []string{
                    "<value 1>",
                    "<value 2>",
                },
                DockerTimeout: criblcontrolplanesdkgo.Pointer[float64](3262.59),
                Filters: []operations.CreateInputSystemByPackContainerFilter{
                    operations.CreateInputSystemByPackContainerFilter{
                        Expr: "<value>",
                    },
                },
                AllContainers: criblcontrolplanesdkgo.Pointer(false),
                PerDevice: criblcontrolplanesdkgo.Pointer(false),
                Detail: criblcontrolplanesdkgo.Pointer(false),
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Persistence: &operations.CreateInputSystemByPackPersistenceSystemMetrics{
                Enable: criblcontrolplanesdkgo.Pointer(true),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            Description: criblcontrolplanesdkgo.Pointer("demobilise because lampoon equal wisely"),
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
### Example Usage: InputCreateExamplesSystemState

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesSystemState" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodySystemState(
        operations.CreateInputSystemByPackInputSystemState{
            ID: "system-state-source",
            Type: operations.CreateInputSystemByPackTypeSystemStateSystemState,
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Interval: criblcontrolplanesdkgo.Pointer[float64](8964.86),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Collectors: &operations.CreateInputSystemByPackCollectors{
                Hostsfile: &operations.CreateInputSystemByPackHostsFile{
                    Enable: criblcontrolplanesdkgo.Pointer(false),
                },
                Interfaces: &operations.CreateInputSystemByPackInterfaces{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Disk: &operations.CreateInputSystemByPackDisksAndFileSystems{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Metadata: &operations.CreateInputSystemByPackHostInfo{
                    Enable: criblcontrolplanesdkgo.Pointer(false),
                },
                Routes: &operations.CreateInputSystemByPackRoutes{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                DNS: &operations.CreateInputSystemByPackDNS{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                User: &operations.CreateInputSystemByPackUsersAndGroups{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Firewall: &operations.CreateInputSystemByPackFirewall{
                    Enable: criblcontrolplanesdkgo.Pointer(false),
                },
                Services: &operations.CreateInputSystemByPackServices{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Ports: &operations.CreateInputSystemByPackListeningPorts{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                LoginUsers: &operations.CreateInputSystemByPackLoggedInUsers{
                    Enable: criblcontrolplanesdkgo.Pointer(false),
                },
            },
            Persistence: &operations.CreateInputSystemByPackPersistenceSystemState{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            DisableNativeModule: criblcontrolplanesdkgo.Pointer(true),
            DisableNativeLastLogModule: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("longingly bah till forgo ick"),
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
### Example Usage: InputCreateExamplesTcp

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesTcp" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyTCP(
        operations.CreateInputSystemByPackInputTCP{
            ID: "tcp-source",
            Type: operations.CreateInputSystemByPackTypeTCPTCP,
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Host: "0.0.0.0",
            Port: 10090,
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
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](2500.34),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](5317.11),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](1943.71),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](6816.93),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            BreakerRulesets: []string{
                "<value 1>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](1670.1),
            EnableHeader: criblcontrolplanesdkgo.Pointer(true),
            Preprocess: &components.PreprocessType{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("given fiddle and tight without until fooey"),
            AuthToken: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
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
### Example Usage: InputCreateExamplesTcpjson

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesTcpjson" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyTcpjson(
        operations.CreateInputSystemByPackInputTcpjson{
            ID: "tcpjson-source",
            Type: operations.CreateInputSystemByPackTypeTcpjsonTcpjson,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Host: "0.0.0.0",
            Port: 10090,
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
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](6808.64),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](8728.65),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](1097.72),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](1900.76),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(true),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("rowdy triumphantly including boo toward monthly once conceal progress surprisingly"),
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
### Example Usage: InputCreateExamplesWef

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesWef" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyWef(
        operations.CreateInputSystemByPackInputWef{
            ID: "wef-source",
            Type: operations.CreateInputSystemByPackTypeWefWef,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Host: "0.0.0.0",
            Port: 5985,
            AuthMethod: operations.CreateInputSystemByPackAuthMethodAuthenticationMethodClientCert.ToPointer(),
            TLS: &operations.CreateInputSystemByPackMTLSSettings{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: "<value>",
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: "<value>",
                CaPath: "<value>",
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
                OcspCheck: criblcontrolplanesdkgo.Pointer(false),
                Keytab: "<value>",
                Principal: "<value>",
                OcspCheckFailClose: criblcontrolplanesdkgo.Pointer(false),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](9533.2),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](251300),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](6579.6),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](9544.3),
            CaFingerprint: criblcontrolplanesdkgo.Pointer("<value>"),
            Keytab: criblcontrolplanesdkgo.Pointer("<value>"),
            Principal: criblcontrolplanesdkgo.Pointer("<value>"),
            AllowMachineIDMismatch: criblcontrolplanesdkgo.Pointer(false),
            Subscriptions: []operations.CreateInputSystemByPackSubscription{
                operations.CreateInputSystemByPackSubscription{
                    SubscriptionName: "subscription-1",
                    Version: criblcontrolplanesdkgo.Pointer("<value>"),
                    ContentFormat: operations.CreateInputSystemByPackFormatRenderedText,
                    HeartbeatInterval: 60,
                    BatchTimeout: 5,
                    ReadExistingEvents: criblcontrolplanesdkgo.Pointer(true),
                    SendBookmarks: criblcontrolplanesdkgo.Pointer(false),
                    Compress: criblcontrolplanesdkgo.Pointer(true),
                    Targets: []string{},
                    Locale: criblcontrolplanesdkgo.Pointer("pt"),
                    QuerySelector: operations.CreateInputSystemByPackQueryBuilderModeSimple.ToPointer(),
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                    Queries: []operations.CreateInputSystemByPackQuery{
                        operations.CreateInputSystemByPackQuery{
                            Path: "/var",
                            QueryExpression: "<value>",
                        },
                    },
                    XMLQuery: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("jaggedly merit shell bah cruelly pale aha whose urban filthy"),
            LogFingerprintMismatch: criblcontrolplanesdkgo.Pointer(false),
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
### Example Usage: InputCreateExamplesWinEventLogs

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesWinEventLogs" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyWinEventLogs(
        operations.CreateInputSystemByPackInputWinEventLogs{
            ID: "win-event-logs-source",
            Type: operations.CreateInputSystemByPackTypeWinEventLogsWinEventLogs,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            LogNames: []string{
                "Application",
                "System",
            },
            ReadMode: operations.CreateInputSystemByPackReadModeNewest.ToPointer(),
            EventFormat: operations.CreateInputSystemByPackEventFormatJSON.ToPointer(),
            DisableNativeModule: criblcontrolplanesdkgo.Pointer(true),
            Interval: criblcontrolplanesdkgo.Pointer[float64](5728.75),
            BatchSize: criblcontrolplanesdkgo.Pointer[float64](9950.94),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            MaxEventBytes: criblcontrolplanesdkgo.Pointer[float64](2912.98),
            Description: criblcontrolplanesdkgo.Pointer("mozzarella ew pfft freezing"),
            DisableJSONRendering: criblcontrolplanesdkgo.Pointer(false),
            DisableXMLRendering: criblcontrolplanesdkgo.Pointer(false),
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
### Example Usage: InputCreateExamplesWindowsMetrics

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesWindowsMetrics" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyWindowsMetrics(
        operations.CreateInputSystemByPackInputWindowsMetrics{
            ID: "windows-metrics-source",
            Type: operations.CreateInputSystemByPackTypeWindowsMetricsWindowsMetrics,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5236.78),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](8788.99),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](4750.79),
            Host: &operations.CreateInputSystemByPackHostWindowsMetrics{
                Mode: components.ModeOptionsHostAll.ToPointer(),
                Custom: &operations.CreateInputSystemByPackCustomWindowsMetrics{
                    System: &operations.CreateInputSystemByPackSystemWindowsMetrics{
                        Mode: operations.CreateInputSystemByPackSystemModeWindowsMetricsBasic.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                    },
                    CPU: &operations.CreateInputSystemByPackCPUWindowsMetrics{
                        Mode: operations.CreateInputSystemByPackCPUModeWindowsMetricsAll.ToPointer(),
                        PerCPU: criblcontrolplanesdkgo.Pointer(true),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Time: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Memory: &operations.CreateInputSystemByPackMemoryWindowsMetrics{
                        Mode: operations.CreateInputSystemByPackMemoryModeWindowsMetricsCustom.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(true),
                    },
                    Network: &operations.CreateInputSystemByPackNetworkWindowsMetrics{
                        Mode: operations.CreateInputSystemByPackNetworkModeWindowsMetricsCustom.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(true),
                        Protocols: criblcontrolplanesdkgo.Pointer(true),
                        Devices: []string{
                            "<value 1>",
                            "<value 2>",
                            "<value 3>",
                        },
                        PerInterface: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Disk: &operations.CreateInputSystemByPackDiskWindowsMetrics{
                        Mode: operations.CreateInputSystemByPackDiskModeWindowsMetricsAll.ToPointer(),
                        PerVolume: criblcontrolplanesdkgo.Pointer(true),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Volumes: []string{
                            "<value 1>",
                            "<value 2>",
                        },
                    },
                },
            },
            Process: &components.ProcessType{
                Sets: []components.ItemsTypeProcessSets{
                    components.ItemsTypeProcessSets{
                        Name: "<value>",
                        Filter: "<value>",
                        IncludeChildren: criblcontrolplanesdkgo.Pointer(true),
                    },
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Persistence: &operations.CreateInputSystemByPackPersistenceWindowsMetrics{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            DisableNativeModule: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("settler tomorrow cappelletti unlike aw limping"),
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
### Example Usage: InputCreateExamplesWiz

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesWiz" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyWiz(
        operations.CreateInputSystemByPackInputWiz{
            ID: "wiz-source",
            Type: operations.CreateInputSystemByPackTypeWizWiz,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Endpoint: "https://api.wiz.io",
            AuthURL: "https://auth.wiz.io/oauth/token",
            AuthAudienceOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            ClientID: "client-id",
            ContentConfig: []operations.CreateInputSystemByPackContentConfigWiz{},
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](1426.41),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](3632.41),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](6059.42),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            RetryRules: &components.RetryRulesType{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesNone,
                Interval: criblcontrolplanesdkgo.Pointer[float64](2446.49),
                Limit: criblcontrolplanesdkgo.Pointer[float64](8984.5),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](2926.4),
                Codes: []float64{
                    9786.9,
                    1849.73,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(true),
            },
            AuthType: components.AuthenticationMethodOptions1Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("understanding finally furthermore than shanghai onto book scarcely incidentally apud"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateEndpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAuthURL: criblcontrolplanesdkgo.Pointer("https://tense-newsstand.net/"),
            TemplateClientID: criblcontrolplanesdkgo.Pointer("<id>"),
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
### Example Usage: InputCreateExamplesWizWebhook

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesWizWebhook" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyWizWebhook(
        operations.CreateInputSystemByPackInputWizWebhook{
            ID: "wiz-webhook-source",
            Type: operations.CreateInputSystemByPackTypeWizWebhookWizWebhook,
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Host: "0.0.0.0",
            Port: 10080,
            AuthTokens: []string{
                "<value 1>",
                "<value 2>",
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](1486.77),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](688186),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](4859.08),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](4334.25),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](7422.5),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](6074.22),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](3865.53),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AllowedPaths: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            AllowedMethods: []string{
                "<value 1>",
            },
            AuthTokensExt: []components.ItemsTypeAuthTokensExt{
                components.ItemsTypeAuthTokensExt{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("beside incidentally helpfully pfft where concerned cute though"),
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("screw male recompense deliberately even hoof lady since flawed"),
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
### Example Usage: InputCreateExamplesZscalerHec

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesZscalerHec" -->
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

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyZscalerHec(
        operations.CreateInputSystemByPackInputZscalerHec{
            ID: "zscaler-hec-source",
            Type: operations.CreateInputSystemByPackTypeZscalerHecZscalerHec,
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Host: "0.0.0.0",
            Port: 8088,
            AuthTokens: []operations.CreateInputSystemByPackAuthTokenZscalerHec{
                operations.CreateInputSystemByPackAuthTokenZscalerHec{
                    AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
                    TokenSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Token: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("slime whoa psst than yahoo"),
                    AllowedIndexesAtToken: []string{
                        "<value 1>",
                        "<value 2>",
                    },
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](1622.31),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](330485),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](4882.74),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](8838.66),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](5051.73),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](7256.83),
            EnableHealthCheck: "<value>",
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            HecAPI: "/services/collector",
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AllowedIndexes: []string{
                "<value 1>",
                "<value 2>",
            },
            HecAcks: criblcontrolplanesdkgo.Pointer(true),
            AccessControlAllowOrigin: []string{
                "<value 1>",
            },
            AccessControlAllowHeaders: []string{
                "<value 1>",
            },
            EmitTokenMetrics: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("trash underpants overplay pillory muddy highly after uh-huh"),
            TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateHecAPI: criblcontrolplanesdkgo.Pointer("<value>"),
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

### Example Usage: InputCreateExamplesAppscope

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesAppscope" -->
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
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
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
### Example Usage: InputCreateExamplesAzureBlob

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesAzureBlob" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputAzureBlob(
        components.InputAzureBlob{
            ID: criblcontrolplanesdkgo.Pointer("azure-blob-source"),
            Type: components.InputAzureBlobTypeAzureBlob,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            QueueName: "azure-blob-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("<value>"),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](3094.63),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](7570.95),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](2253.54),
            ServicePeriodSecs: criblcontrolplanesdkgo.Pointer[float64](7697.99),
            SkipOnError: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](2956.83),
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](7633.01),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](4156.74),
            AuthType: components.AuthenticationMethodOptionsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("knavishly vivaciously sympathetically"),
            ConnectionString: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            StorageAccountName: criblcontrolplanesdkgo.Pointer("<value>"),
            TenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            ClientID: criblcontrolplanesdkgo.Pointer("<id>"),
            AzureCloud: criblcontrolplanesdkgo.Pointer("<value>"),
            EndpointSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            ClientTextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Certificate: &components.CertificateTypeAzureBlobAuthTypeClientCert{
                CertificateName: "<value>",
            },
            TemplateQueueName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateConnectionString: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateTenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateClientID: criblcontrolplanesdkgo.Pointer("<id>"),
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
### Example Usage: InputCreateExamplesCloudflareHec

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesCloudflareHec" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputCloudflareHec(
        components.InputCloudflareHec{
            ID: criblcontrolplanesdkgo.Pointer("cloudflare-hec-source"),
            Type: components.InputCloudflareHecTypeCloudflareHec,
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
            Host: "0.0.0.0",
            Port: 8088,
            AuthTokens: []components.InputCloudflareHecAuthToken{
                components.InputCloudflareHecAuthToken{
                    AuthType: components.InputCloudflareHecAuthenticationMethodSecret.ToPointer(),
                    TokenSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Token: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                    Description: criblcontrolplanesdkgo.Pointer("convalesce after sympathetically"),
                    AllowedIndexesAtToken: []string{
                        "<value 1>",
                        "<value 2>",
                    },
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
            TLS: &components.TLSSettingsServerSide{
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](8504.29),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](808412),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](8898.86),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](2032.33),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](3074.69),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](8911.87),
            EnableHealthCheck: "<value>",
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            HecAPI: "/services/collector",
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AllowedIndexes: []string{
                "<value 1>",
            },
            BreakerRulesets: []string{
                "<value 1>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](7424.11),
            AccessControlAllowOrigin: []string{
                "<value 1>",
                "<value 2>",
            },
            AccessControlAllowHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            EmitTokenMetrics: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("successfully ecstatic if woefully apprehensive soulful nicely butter even deprave"),
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
### Example Usage: InputCreateExamplesCollection

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesCollection" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputCollection(
        components.InputCollection{
            ID: criblcontrolplanesdkgo.Pointer("collection-source"),
            Type: components.InputCollectionTypeCollection,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            BreakerRulesets: []string{
                "<value 1>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](2306.82),
            Preprocess: &components.PreprocessType{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
            },
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Output: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesConfluentCloud

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesConfluentCloud" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputConfluentCloud(
        components.InputConfluentCloud{
            ID: criblcontrolplanesdkgo.Pointer("confluent-cloud-source"),
            Type: components.InputConfluentCloudTypeConfluentCloud,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Brokers: []string{
                "pkc-xxxxx.us-east-1.aws.confluent.cloud:9092",
            },
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            Topics: []string{
                "logs",
            },
            GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
            FromBeginning: criblcontrolplanesdkgo.Pointer(false),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType{
                Disabled: true,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://foolish-lace.name"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](9222.03),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](2060.92),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](4747.86),
                Auth: &components.AuthTypeKafkaSchemaRegistry{
                    Disabled: true,
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                    CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                },
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](2191.22),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](9306.44),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](9430.74),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](5234.73),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5835.16),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8876.61),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](4184.48),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](2863.18),
            Sasl: &components.AuthenticationType{
                Disabled: false,
                Username: criblcontrolplanesdkgo.Pointer("Hallie.Wisozk"),
                Password: criblcontrolplanesdkgo.Pointer("oF6ULVD266SvlIK"),
                AuthType: components.AuthenticationMethodOptionsSaslSecret.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslScramSha512.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(true),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://rewarding-courtroom.net/"),
                ClientID: criblcontrolplanesdkgo.Pointer("<id>"),
                OauthSecretType: criblcontrolplanesdkgo.Pointer("<value>"),
                ClientTextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthParams: []components.ItemsTypeSaslOauthParams{
                    components.ItemsTypeSaslOauthParams{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                SaslExtensions: []components.ItemsTypeSaslSaslExtensions{
                    components.ItemsTypeSaslSaslExtensions{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
            },
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](789.56),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](6327.46),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](6024.79),
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](2783),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](9984.03),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](214.49),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](1937.09),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](7019.66),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("qua simple retention gratefully huzzah hefty crazy because"),
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
### Example Usage: InputCreateExamplesCriblHttp

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesCriblHttp" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputCriblHTTP(
        components.InputCriblHTTP{
            ID: criblcontrolplanesdkgo.Pointer("cribl-http-source"),
            Type: components.InputCriblHTTPTypeCriblHTTP,
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
            Host: "0.0.0.0",
            Port: 10080,
            AuthTokens: []components.ItemsTypeAuthTokens{
                components.ItemsTypeAuthTokens{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                    Description: criblcontrolplanesdkgo.Pointer("that convince gift ha maroon sadly above fraternise affect always"),
                },
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](94.06),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](995636),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](7620.39),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](5316.13),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](5927.22),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](984.97),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("meanwhile atomize receptor wildly whether competent intelligent"),
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
### Example Usage: InputCreateExamplesCriblLakeHttp

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesCriblLakeHttp" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputCriblLakeHTTP(
        components.InputCriblLakeHTTP{
            ID: criblcontrolplanesdkgo.Pointer("cribl-lake-http-source"),
            Type: components.InputCriblLakeHTTPTypeCriblLakeHTTP,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10080,
            AuthTokens: []string{
                "<value 1>",
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](3932.24),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](381048),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](7154.34),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](5542.72),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](5091.92),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](4354.07),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            CriblAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            ElasticAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            SplunkHecAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            SplunkHecAcks: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthTokensExt: []components.AuthTokensExt{
                components.AuthTokensExt{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("joyfully duh once among"),
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                    SplunkHecMetadata: &components.SplunkHecMetadata{
                        Enabled: criblcontrolplanesdkgo.Pointer(false),
                        DefaultDataset: criblcontrolplanesdkgo.Pointer("<value>"),
                        AllowedIndexesAtToken: []string{
                            "<value 1>",
                        },
                    },
                    ElasticsearchMetadata: &components.ElasticsearchMetadata{
                        Enabled: criblcontrolplanesdkgo.Pointer(false),
                        DefaultDataset: criblcontrolplanesdkgo.Pointer("<value>"),
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("suitcase hm um selfish"),
            TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateSplunkHecAPI: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesCriblTcp

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesCriblTcp" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputCriblTCP(
        components.InputCriblTCP{
            ID: criblcontrolplanesdkgo.Pointer("cribl-tcp-source"),
            Type: components.InputCriblTCPTypeCriblTCP,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Host: "0.0.0.0",
            Port: 10090,
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
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](3802.79),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](4455.7),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](7985.02),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](275.69),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(false),
            AuthTokens: []components.ItemsTypeAuthTokens{
                components.ItemsTypeAuthTokens{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                    Description: criblcontrolplanesdkgo.Pointer("that convince gift ha maroon sadly above fraternise affect always"),
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("consequently wise out yuck in option emboss silt"),
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
### Example Usage: InputCreateExamplesCrowdstrike

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesCrowdstrike" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputCrowdstrike(
        components.InputCrowdstrike{
            ID: criblcontrolplanesdkgo.Pointer("crowdstrike-source"),
            Type: components.InputCrowdstrikeTypeCrowdstrike,
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            QueueName: "crowdstrike-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfManual.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](1413.78),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](1324.06),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](6723.13),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](6219.39),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](8635.01),
            SkipOnError: criblcontrolplanesdkgo.Pointer(false),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](7464.28),
            EnableSQSAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            Preprocess: &components.PreprocessType{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Checkpointing: &components.CheckpointingType{
                Enabled: false,
                Retries: criblcontrolplanesdkgo.Pointer[float64](1555.53),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](1226.29),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("partridge gah and instructor plus"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TagAfterProcessing: components.TagAfterProcessingOptionsFalse.ToPointer(),
            ProcessedTagKey: criblcontrolplanesdkgo.Pointer("<value>"),
            ProcessedTagValue: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateQueueName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesDatadogAgent

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesDatadogAgent" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputDatadogAgent(
        components.InputDatadogAgent{
            ID: criblcontrolplanesdkgo.Pointer("datadog-agent-source"),
            Type: components.InputDatadogAgentTypeDatadogAgent,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 8126,
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](3701.81),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](562306),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](6427.39),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](6528.52),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](4974.13),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](3714.31),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ProxyMode: &components.InputDatadogAgentProxyMode{
                Enabled: true,
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            },
            Description: criblcontrolplanesdkgo.Pointer("gradient alongside bug fill which woot considering deadly"),
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
### Example Usage: InputCreateExamplesDatagen

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesDatagen" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputDatagen(
        components.InputDatagen{
            ID: criblcontrolplanesdkgo.Pointer("datagen-source"),
            Type: components.InputDatagenTypeDatagen,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Samples: []components.Sample{
                components.Sample{
                    Sample: "sample.json",
                    EventsPerSec: 10,
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("that apud awkwardly reclassify bid ha ew sardonic boo ameliorate"),
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
### Example Usage: InputCreateExamplesEdgePrometheus

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesEdgePrometheus" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputEdgePrometheus(
        components.InputEdgePrometheus{
            ID: criblcontrolplanesdkgo.Pointer("edge-prometheus-source"),
            Type: components.InputEdgePrometheusTypeEdgePrometheus,
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            DimensionList: []string{
                "<value 1>",
            },
            DiscoveryType: components.InputEdgePrometheusDiscoveryTypeStatic,
            Interval: 60,
            Timeout: criblcontrolplanesdkgo.Pointer[float64](256.89),
            Persistence: &components.DiskSpoolingType{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.CompressionOptionsPersistenceNone.ToPointer(),
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthType: components.InputEdgePrometheusAuthenticationMethodSecret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("those boo seafood but across around"),
            Targets: []components.Target{
                components.Target{
                    Protocol: components.ProtocolOptionsTargetsItemsHTTP.ToPointer(),
                    Host: "localhost",
                    Port: criblcontrolplanesdkgo.Pointer[float64](2116.57),
                    Path: criblcontrolplanesdkgo.Pointer("/opt/lib"),
                },
            },
            RecordType: components.RecordTypeOptionsSrv.ToPointer(),
            ScrapePort: criblcontrolplanesdkgo.Pointer[float64](8442.04),
            NameList: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ScrapeProtocol: components.ProtocolOptionsTargetsItemsHTTP.ToPointer(),
            ScrapePath: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfManual.ToPointer(),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            UsePublicIP: criblcontrolplanesdkgo.Pointer(false),
            SearchFilter: []components.ItemsTypeSearchFilter{
                components.ItemsTypeSearchFilter{
                    Name: "<value>",
                    Values: []string{},
                },
            },
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions1V4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](5608.71),
            ScrapeProtocolExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            ScrapePortExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            ScrapePathExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            PodFilter: []components.PodFilter{
                components.PodFilter{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("toaster whenever reflecting indeed rationale"),
                },
            },
            Username: criblcontrolplanesdkgo.Pointer("Karlee.Prohaska"),
            Password: criblcontrolplanesdkgo.Pointer("Lx6EeJZLVVlDkZb"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
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
### Example Usage: InputCreateExamplesElastic

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesElastic" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputElastic(
        components.InputElastic{
            ID: criblcontrolplanesdkgo.Pointer("elastic-source"),
            Type: components.InputElasticTypeElastic,
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
            Host: "localhost",
            Port: 9200,
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](2608.82),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](179171),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](7866.63),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7159.18),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](5602.34),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](3586.64),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            ElasticAPI: "/",
            AuthType: components.InputElasticAuthenticationTypeCredentialsSecret.ToPointer(),
            APIVersion: components.InputElasticAPIVersionSixDot8Dot4.ToPointer(),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ProxyMode: &components.InputElasticProxyMode{
                Enabled: true,
                AuthType: components.InputElasticAuthenticationMethodSecret.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("Courtney.Welch"),
                Password: criblcontrolplanesdkgo.Pointer("Lge4i0P6LwZldaN"),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                URL: criblcontrolplanesdkgo.Pointer("https://scientific-interior.name"),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                RemoveHeaders: []string{
                    "<value 1>",
                },
                TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](9619.23),
                TemplateURL: criblcontrolplanesdkgo.Pointer("https://exalted-turret.net"),
            },
            Description: criblcontrolplanesdkgo.Pointer("geez elderly concerning translation"),
            Username: criblcontrolplanesdkgo.Pointer("Marcus75"),
            Password: criblcontrolplanesdkgo.Pointer("Nx2JR_RGCX4Zt2A"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthTokens: []string{
                "<value 1>",
            },
            CustomAPIVersion: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesEventhub

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesEventhub" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputEventhub(
        components.InputEventhub{
            ID: criblcontrolplanesdkgo.Pointer("eventhub-source"),
            Type: components.InputEventhubTypeEventhub,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Brokers: []string{
                "myeventhub.servicebus.windows.net:9093",
            },
            Topics: []string{
                "logs",
            },
            GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
            FromBeginning: criblcontrolplanesdkgo.Pointer(true),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](1083.7),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7233.09),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](584.36),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](9713.26),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](4326.99),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8957.69),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](3151.51),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](4345.38),
            Sasl: &components.AuthenticationType1{
                Disabled: false,
                AuthType: components.AuthenticationMethodOptionsSasl1Secret.ToPointer(),
                Password: criblcontrolplanesdkgo.Pointer("5xOWKBHPhbLS62p"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSasl1Oauthbearer.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("Rosalinda89"),
                ClientSecretAuthType: components.AuthenticationMethodOptionsSasl2Certificate.ToPointer(),
                ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                ClientTextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEndpoint: components.MicrosoftEntraIDAuthenticationEndpointOptionsSaslHTTPSLoginMicrosoftonlineUs.ToPointer(),
                ClientID: criblcontrolplanesdkgo.Pointer("<id>"),
                TenantID: criblcontrolplanesdkgo.Pointer("<id>"),
                Scope: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            TLS: &components.TLSSettingsClientSideType{
                Disabled: true,
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            },
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](1137.77),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](6787.34),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](7898.58),
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](3538.75),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](8585.97),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](1564.23),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](3331.78),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](7530.35),
            MinimizeDuplicates: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("down yippee yowza meanwhile when fireplace solemnly"),
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
### Example Usage: InputCreateExamplesExec

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesExec" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputExec(
        components.InputExec{
            ID: criblcontrolplanesdkgo.Pointer("exec-source"),
            Type: components.InputExecTypeExec,
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Command: "echo \"Hello World\"",
            Retries: criblcontrolplanesdkgo.Pointer[float64](357.57),
            ScheduleType: components.ScheduleTypeCronSchedule.ToPointer(),
            BreakerRulesets: []string{
                "<value 1>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](347.07),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("rot cap eek finally"),
            Interval: criblcontrolplanesdkgo.Pointer[float64](60),
            CronSchedule: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesFile

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesFile" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputFile(
        components.InputFile{
            ID: criblcontrolplanesdkgo.Pointer("file-source"),
            Type: components.InputFileTypeFile,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Mode: components.InputFileModeManual.ToPointer(),
            Interval: criblcontrolplanesdkgo.Pointer[float64](2045.59),
            Filenames: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            FilterArchivedFiles: criblcontrolplanesdkgo.Pointer(true),
            TailOnly: criblcontrolplanesdkgo.Pointer(false),
            IdleTimeout: criblcontrolplanesdkgo.Pointer[float64](5068.88),
            MinAgeDur: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxAgeDur: criblcontrolplanesdkgo.Pointer("<value>"),
            CheckFileModTime: criblcontrolplanesdkgo.Pointer(true),
            ForceText: criblcontrolplanesdkgo.Pointer(true),
            HashLen: criblcontrolplanesdkgo.Pointer[float64](475.25),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](8347.38),
            Description: criblcontrolplanesdkgo.Pointer("to own daily redraw squiggly aside probate"),
            Path: criblcontrolplanesdkgo.Pointer("/etc/namedb"),
            Depth: criblcontrolplanesdkgo.Pointer[float64](8225.75),
            SuppressMissingPathErrors: criblcontrolplanesdkgo.Pointer(false),
            DeleteFiles: criblcontrolplanesdkgo.Pointer(false),
            SaltHash: criblcontrolplanesdkgo.Pointer(false),
            IncludeUnidentifiableBinary: criblcontrolplanesdkgo.Pointer(false),
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
### Example Usage: InputCreateExamplesFirehose

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesFirehose" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputFirehose(
        components.InputFirehose{
            ID: criblcontrolplanesdkgo.Pointer("firehose-source"),
            Type: components.InputFirehoseTypeFirehose,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10080,
            AuthTokens: []string{
                "<value 1>",
                "<value 2>",
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](4121.05),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](447263),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](8495.08),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](3063.22),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](9184.91),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](226.13),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("both merry exactly congregate"),
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
### Example Usage: InputCreateExamplesGooglePubsub

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesGooglePubsub" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputGooglePubsub(
        components.InputGooglePubsub{
            ID: criblcontrolplanesdkgo.Pointer("google-pubsub-source"),
            Type: components.InputGooglePubsubTypeGooglePubsub,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            TopicName: "my-topic",
            SubscriptionName: "my-subscription",
            MonitorSubscription: criblcontrolplanesdkgo.Pointer(false),
            CreateTopic: criblcontrolplanesdkgo.Pointer(false),
            CreateSubscription: criblcontrolplanesdkgo.Pointer(false),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            GoogleAuthMethod: components.GoogleAuthenticationMethodOptionsManual.ToPointer(),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxBacklog: criblcontrolplanesdkgo.Pointer[float64](538.05),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](42.36),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](4688.25),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("ugh mismatch soon sunny apropos hmph"),
            OrderedDelivery: criblcontrolplanesdkgo.Pointer(false),
            TemplateTopicName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateSubscriptionName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesGrafana

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesGrafana" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputGrafana(
        components.CreateInputGrafanaInputGrafanaGrafana1(
            components.InputGrafanaGrafana1{
                ID: criblcontrolplanesdkgo.Pointer("grafana-source"),
                Type: components.InputGrafanaType1Grafana,
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
                    Mode: components.ModeOptionsPqSmart.ToPointer(),
                    MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                    CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                    MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                    Compress: components.CompressionOptionsPqGzip.ToPointer(),
                    PqControls: &components.PqTypePqControls{},
                },
                Host: "0.0.0.0",
                Port: 10080,
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
                MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](7432.78),
                MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](363916),
                EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
                CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
                ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](9141.19),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](8174.29),
                SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](7036.5),
                KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](3091.06),
                EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
                IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
                IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
                PrometheusAPI: "/api/prom/push",
                LokiAPI: criblcontrolplanesdkgo.Pointer("<value>"),
                PrometheusAuth: &components.PrometheusAuth1{
                    AuthType: components.AuthenticationTypeOptionsPrometheusAuthBasic.ToPointer(),
                    Username: criblcontrolplanesdkgo.Pointer("Elody.Wintheiser"),
                    Password: criblcontrolplanesdkgo.Pointer("wz_3eL1dGvlO6eV"),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                LokiAuth: &components.LokiAuth1{
                    AuthType: components.AuthenticationTypeOptionsLokiAuthNone.ToPointer(),
                    Username: criblcontrolplanesdkgo.Pointer("Buck26"),
                    Password: criblcontrolplanesdkgo.Pointer("dlVcDmpeEG4LHgq"),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                Metadata: []components.ItemsTypeMetadata{
                    components.ItemsTypeMetadata{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                Description: criblcontrolplanesdkgo.Pointer("pilot to fatally authentic chime bah"),
                TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
            },
        ),
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInput != nil {
        // handle response
    }
}
```
### Example Usage: InputCreateExamplesHttp

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesHttp" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputHTTP(
        components.InputHTTP{
            ID: criblcontrolplanesdkgo.Pointer("http-source"),
            Type: components.InputHTTPTypeHTTP,
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
            Host: "0.0.0.0",
            Port: 10080,
            AuthTokens: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](3968.12),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](834950),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](9398.33),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](2620.6),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](689.57),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](1105.14),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            CriblAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            ElasticAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            SplunkHecAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            SplunkHecAcks: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthTokensExt: []components.ItemsTypeAuthTokensExt{
                components.ItemsTypeAuthTokensExt{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("ew phooey unless aha plus woot zowie unless"),
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("finally whole outnumber over apropos frozen"),
            TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateSplunkHecAPI: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesHttpRaw

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesHttpRaw" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputHTTPRaw(
        components.InputHTTPRaw{
            ID: criblcontrolplanesdkgo.Pointer("http-raw-source"),
            Type: components.InputHTTPRawTypeHTTPRaw,
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
                Mode: components.ModeOptionsPqSmart.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10080,
            AuthTokens: []string{
                "<value 1>",
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](2155.01),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](339648),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](4438.78),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](9645.86),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](804.45),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5141.74),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](3358.37),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AllowedPaths: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            AllowedMethods: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            AuthTokensExt: []components.ItemsTypeAuthTokensExt{
                components.ItemsTypeAuthTokensExt{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("ew phooey unless aha plus woot zowie unless"),
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("quizzically blight athwart ugh invite aha tenant multicolored if wherever"),
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
### Example Usage: InputCreateExamplesJournalFiles

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesJournalFiles" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputJournalFiles(
        components.InputJournalFiles{
            ID: criblcontrolplanesdkgo.Pointer("journal-files-source"),
            Type: components.InputJournalFilesTypeJournalFiles,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Path: "/var/log/journal",
            Interval: criblcontrolplanesdkgo.Pointer[float64](6340.49),
            Journals: []string{
                "system",
            },
            Rules: []components.InputJournalFilesRule{
                components.InputJournalFilesRule{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("mmm by yesterday"),
                },
            },
            CurrentBoot: criblcontrolplanesdkgo.Pointer(false),
            MaxAgeDur: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("likewise off atop instead never"),
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
### Example Usage: InputCreateExamplesKafka

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesKafka" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputKafka(
        components.InputKafka{
            ID: criblcontrolplanesdkgo.Pointer("kafka-source"),
            Type: components.InputKafkaTypeKafka,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Brokers: []string{
                "localhost:9092",
            },
            Topics: []string{
                "logs",
            },
            GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
            FromBeginning: criblcontrolplanesdkgo.Pointer(true),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType{
                Disabled: true,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://foolish-lace.name"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](9222.03),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](2060.92),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](4747.86),
                Auth: &components.AuthTypeKafkaSchemaRegistry{
                    Disabled: true,
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                    CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                },
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](8097.01),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](6229.98),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](6946.19),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](3258.42),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5604.23),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3668.67),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](8087.17),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](1502.06),
            Sasl: &components.AuthenticationType{
                Disabled: false,
                Username: criblcontrolplanesdkgo.Pointer("Hallie.Wisozk"),
                Password: criblcontrolplanesdkgo.Pointer("oF6ULVD266SvlIK"),
                AuthType: components.AuthenticationMethodOptionsSaslSecret.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslScramSha512.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(true),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://rewarding-courtroom.net/"),
                ClientID: criblcontrolplanesdkgo.Pointer("<id>"),
                OauthSecretType: criblcontrolplanesdkgo.Pointer("<value>"),
                ClientTextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthParams: []components.ItemsTypeSaslOauthParams{
                    components.ItemsTypeSaslOauthParams{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                SaslExtensions: []components.ItemsTypeSaslSaslExtensions{
                    components.ItemsTypeSaslSaslExtensions{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
            },
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](1415.5),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](5063.07),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](1418.21),
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](5495.4),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](3976.66),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](7857.05),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](8759.77),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](5618.92),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("tousle draw excess"),
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
### Example Usage: InputCreateExamplesKinesis

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesKinesis" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputKinesis(
        components.InputKinesis{
            ID: criblcontrolplanesdkgo.Pointer("kinesis-source"),
            Type: components.InputKinesisTypeKinesis,
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
                Mode: components.ModeOptionsPqSmart.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            StreamName: "my-stream",
            ServiceInterval: criblcontrolplanesdkgo.Pointer[float64](9387.25),
            ShardExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            ShardIteratorType: components.ShardIteratorStartLatest.ToPointer(),
            PayloadFormat: components.RecordDataFormatLine.ToPointer(),
            GetRecordsLimit: criblcontrolplanesdkgo.Pointer[float64](3821.23),
            GetRecordsLimitTotal: criblcontrolplanesdkgo.Pointer[float64](4535.77),
            LoadBalancingAlgorithm: components.ShardLoadBalancingConsistentHashing.ToPointer(),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfSecret.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions2V2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](7175.07),
            VerifyKPLCheckSums: criblcontrolplanesdkgo.Pointer(true),
            AvoidDuplicates: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("circa annually yahoo suckle the apropos divine out at too"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateStreamName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesKubeEvents

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesKubeEvents" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputKubeEvents(
        components.InputKubeEvents{
            ID: criblcontrolplanesdkgo.Pointer("kube-events-source"),
            Type: components.InputKubeEventsTypeKubeEvents,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Rules: []components.ItemsTypeRules{
                components.ItemsTypeRules{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("chow vamoose devastation"),
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("lone petticoat frightfully evil yum"),
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
### Example Usage: InputCreateExamplesKubeLogs

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesKubeLogs" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputKubeLogs(
        components.InputKubeLogs{
            ID: criblcontrolplanesdkgo.Pointer("kube-logs-source"),
            Type: components.InputKubeLogsTypeKubeLogs,
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
            Interval: criblcontrolplanesdkgo.Pointer[float64](1273.6),
            Rules: []components.InputKubeLogsRule{
                components.InputKubeLogsRule{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("suddenly foolish quiet breastplate sophisticated drat what"),
                },
            },
            Timestamps: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Persistence: &components.DiskSpoolingType{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.CompressionOptionsPersistenceNone.ToPointer(),
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](2370.87),
            EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("alb athletic bourgeoisie rapidly ew anenst license"),
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
### Example Usage: InputCreateExamplesKubeMetrics

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesKubeMetrics" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputKubeMetrics(
        components.InputKubeMetrics{
            ID: criblcontrolplanesdkgo.Pointer("kube-metrics-source"),
            Type: components.InputKubeMetricsTypeKubeMetrics,
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Interval: criblcontrolplanesdkgo.Pointer[float64](8777.96),
            Rules: []components.ItemsTypeRules{
                components.ItemsTypeRules{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("chow vamoose devastation"),
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Persistence: &components.InputKubeMetricsPersistence{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            Description: criblcontrolplanesdkgo.Pointer("duster when flood"),
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
### Example Usage: InputCreateExamplesLoki

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesLoki" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputLoki(
        components.InputLoki{
            ID: criblcontrolplanesdkgo.Pointer("loki-source"),
            Type: components.InputLokiTypeLoki,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10080,
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](2261.6),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](403885),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](391.58),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7446.1),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](7887.34),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](7085.76),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            LokiAPI: "/loki/api/v1/push",
            AuthType: components.AuthenticationTypeOptionsLokiAuthBasic.ToPointer(),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("on christen including behind annually disclosure boo whenever rubric printer"),
            Username: criblcontrolplanesdkgo.Pointer("Lou_Heathcote"),
            Password: criblcontrolplanesdkgo.Pointer("qSN2owb1cWSvjjh"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesMetrics

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesMetrics" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputMetrics(
        components.InputMetrics{
            ID: criblcontrolplanesdkgo.Pointer("metrics-source"),
            Type: components.InputMetricsTypeMetrics,
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
            Host: "0.0.0.0",
            UDPPort: criblcontrolplanesdkgo.Pointer[float64](8125),
            TCPPort: criblcontrolplanesdkgo.Pointer[float64](3449.96),
            MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](840.01),
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
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
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](6578.52),
            Description: criblcontrolplanesdkgo.Pointer("furthermore brilliant wretched geez gee"),
            TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateUDPPort: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateTCPPort: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesModelDrivenTelemetry

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesModelDrivenTelemetry" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputModelDrivenTelemetry(
        components.InputModelDrivenTelemetry{
            ID: criblcontrolplanesdkgo.Pointer("mdt-source"),
            Type: components.InputModelDrivenTelemetryTypeModelDrivenTelemetry,
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
            Host: "0.0.0.0",
            Port: 57000,
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
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](4616.35),
            ShutdownTimeoutMs: criblcontrolplanesdkgo.Pointer[float64](8224.03),
            Description: criblcontrolplanesdkgo.Pointer("boastfully past siege knavishly spellcheck scale frank"),
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
### Example Usage: InputCreateExamplesMsk

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesMsk" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputMsk(
        components.InputMsk{
            ID: criblcontrolplanesdkgo.Pointer("msk-source"),
            Type: components.InputMskTypeMsk,
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
                Mode: components.ModeOptionsPqSmart.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Brokers: []string{
                "b-1.example.xxxxx.c2.kafka.us-east-1.amazonaws.com:9092",
            },
            Topics: []string{
                "logs",
            },
            GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
            FromBeginning: criblcontrolplanesdkgo.Pointer(false),
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](9034.08),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](5113.84),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](9292.28),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType{
                Disabled: true,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://foolish-lace.name"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](9222.03),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](2060.92),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](4747.86),
                Auth: &components.AuthTypeKafkaSchemaRegistry{
                    Disabled: true,
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                    CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                },
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](5239.83),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7415.72),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](3350.38),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](1085.96),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2915.63),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5562.81),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](3485.6),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](1186.43),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto,
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3369.8),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](3782.76),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](5370.18),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](493.58),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](7523.73),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](392.88),
            Description: criblcontrolplanesdkgo.Pointer("upon deceivingly vanish profane sans furthermore um even"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesNetflow

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesNetflow" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputNetflow(
        components.InputNetflow{
            ID: criblcontrolplanesdkgo.Pointer("netflow-source"),
            Type: components.InputNetflowTypeNetflow,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 2055,
            EnablePassThrough: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](9889.4),
            TemplateCacheMinutes: criblcontrolplanesdkgo.Pointer[float64](2571.11),
            V5Enabled: criblcontrolplanesdkgo.Pointer(false),
            V9Enabled: criblcontrolplanesdkgo.Pointer(true),
            IpfixEnabled: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("fearless ha what gym when"),
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
### Example Usage: InputCreateExamplesOffice365Mgmt

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesOffice365Mgmt" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputOffice365Mgmt(
        components.InputOffice365Mgmt{
            ID: criblcontrolplanesdkgo.Pointer("office365-mgmt-source"),
            Type: components.InputOffice365MgmtTypeOffice365Mgmt,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            PlanType: components.SubscriptionPlanOptionsEnterpriseGcc,
            TenantID: "tenant-id",
            AppID: "app-id",
            Timeout: criblcontrolplanesdkgo.Pointer[float64](7322.78),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](2466.82),
            JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](3801.89),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            PublisherIdentifier: criblcontrolplanesdkgo.Pointer("<value>"),
            ContentConfig: []components.InputOffice365MgmtContentConfig{
                components.InputOffice365MgmtContentConfig{
                    ContentType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Description: criblcontrolplanesdkgo.Pointer("truthfully ragged apropos better since now whoever as oof without"),
                    Interval: criblcontrolplanesdkgo.Pointer[float64](9148.23),
                    LogLevel: components.LogLevelOptionsContentConfigItemsError.ToPointer(),
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                },
            },
            IngestionLag: criblcontrolplanesdkgo.Pointer[float64](8798.19),
            RetryRules: &components.RetryRulesType1{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesStatic,
                Interval: criblcontrolplanesdkgo.Pointer[float64](3819.89),
                Limit: criblcontrolplanesdkgo.Pointer[float64](6889.97),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](1601.39),
                Codes: []float64{
                    8528.61,
                    6036.97,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(true),
            },
            AuthType: components.AuthenticationMethodOptions1Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("aha er finally"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateTenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAppID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplatePublisherIdentifier: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesOffice365MsgTrace

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesOffice365MsgTrace" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputOffice365MsgTrace(
        components.InputOffice365MsgTrace{
            ID: criblcontrolplanesdkgo.Pointer("office365-msg-trace-source"),
            Type: components.InputOffice365MsgTraceTypeOffice365MsgTrace,
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
            URL: "https://reports.office365.com/ecp/reportingwebservice/reporting.svc/MessageTrace",
            Interval: 15,
            StartDate: criblcontrolplanesdkgo.Pointer("<value>"),
            EndDate: criblcontrolplanesdkgo.Pointer("<value>"),
            Timeout: criblcontrolplanesdkgo.Pointer[float64](2553.49),
            DisableTimeFilter: criblcontrolplanesdkgo.Pointer(false),
            AuthType: components.InputOffice365MsgTraceAuthenticationMethodOauth.ToPointer(),
            RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(false),
            MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](8433.78),
            LogLevel: components.InputOffice365MsgTraceLogLevelWarn.ToPointer(),
            JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](8444.87),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](3651.45),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            RetryRules: &components.RetryRulesType1{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesStatic,
                Interval: criblcontrolplanesdkgo.Pointer[float64](3819.89),
                Limit: criblcontrolplanesdkgo.Pointer[float64](6889.97),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](1601.39),
                Codes: []float64{
                    8528.61,
                    6036.97,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(true),
            },
            Description: criblcontrolplanesdkgo.Pointer("apropos brr deliquesce cheese ah"),
            Username: criblcontrolplanesdkgo.Pointer("Vernice_Langworth35"),
            Password: criblcontrolplanesdkgo.Pointer("_pAlCa3zWUZn969"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            ClientID: criblcontrolplanesdkgo.Pointer("<id>"),
            Resource: criblcontrolplanesdkgo.Pointer("<value>"),
            PlanType: components.SubscriptionPlanOptionsDod.ToPointer(),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            CertOptions: &components.CertOptions{
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: "<value>",
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: "<value>",
            },
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://icy-finding.name/"),
            TemplateTenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateClientID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateResource: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesOffice365Service

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesOffice365Service" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputOffice365Service(
        components.InputOffice365Service{
            ID: criblcontrolplanesdkgo.Pointer("office365-service-source"),
            Type: components.InputOffice365ServiceTypeOffice365Service,
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            PlanType: components.SubscriptionPlanOptionsGcc.ToPointer(),
            TenantID: "tenant-id",
            AppID: "app-id",
            Timeout: criblcontrolplanesdkgo.Pointer[float64](3852.82),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](5227.5),
            JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](7117.28),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ContentConfig: []components.InputOffice365ServiceContentConfig{
                components.InputOffice365ServiceContentConfig{
                    ContentType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Description: criblcontrolplanesdkgo.Pointer("alligator helpfully definitive perp scarily"),
                    Interval: criblcontrolplanesdkgo.Pointer[float64](8076.88),
                    LogLevel: components.LogLevelOptionsContentConfigItemsDebug.ToPointer(),
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                },
            },
            RetryRules: &components.RetryRulesType1{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesStatic,
                Interval: criblcontrolplanesdkgo.Pointer[float64](3819.89),
                Limit: criblcontrolplanesdkgo.Pointer[float64](6889.97),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](1601.39),
                Codes: []float64{
                    8528.61,
                    6036.97,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(true),
            },
            AuthType: components.AuthenticationMethodOptions1Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("via selfishly plus puppet contrast pfft considering"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateTenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAppID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesOpenTelemetry

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesOpenTelemetry" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputOpenTelemetry(
        components.InputOpenTelemetry{
            ID: criblcontrolplanesdkgo.Pointer("otel-source"),
            Type: components.InputOpenTelemetryTypeOpenTelemetry,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 4317,
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](3469),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](295125),
            EnableProxyHeader: "<value>",
            CaptureHeaders: "<value>",
            ActivityLogSampleRate: "<value>",
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](6472.11),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](8701.9),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](8352.78),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            Protocol: components.InputOpenTelemetryProtocolHTTP.ToPointer(),
            ExtractSpans: criblcontrolplanesdkgo.Pointer(true),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(true),
            OtlpVersion: components.InputOpenTelemetryOTLPVersionZeroDot10Dot0.ToPointer(),
            AuthType: components.AuthenticationTypeOptionsCredentialsSecret.ToPointer(),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](4972.97),
            Description: criblcontrolplanesdkgo.Pointer("lively gadzooks gestate dowse depute carelessly highly"),
            Username: criblcontrolplanesdkgo.Pointer("Jess22"),
            Password: criblcontrolplanesdkgo.Pointer("2_mROJCOhV3vY75"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ExtractLogs: criblcontrolplanesdkgo.Pointer(true),
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
### Example Usage: InputCreateExamplesPrometheus

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesPrometheus" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputPrometheus(
        components.InputPrometheus{
            ID: criblcontrolplanesdkgo.Pointer("prometheus-source"),
            Type: components.InputPrometheusTypePrometheus,
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
            DimensionList: []string{
                "<value 1>",
                "<value 2>",
            },
            DiscoveryType: components.InputPrometheusDiscoveryTypeStatic.ToPointer(),
            Interval: 60,
            LogLevel: components.InputPrometheusLogLevelInfo,
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            Timeout: criblcontrolplanesdkgo.Pointer[float64](8250.81),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](4617.44),
            JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](4814.04),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthType: components.AuthenticationMethodOptionsSaslSecret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("hm instead gosh whack undergo"),
            TargetList: []string{
                "http://localhost:9090/metrics",
            },
            RecordType: components.RecordTypeOptionsAaaa.ToPointer(),
            ScrapePort: criblcontrolplanesdkgo.Pointer[float64](8151.75),
            NameList: []string{
                "<value 1>",
            },
            ScrapeProtocol: components.MetricsProtocolHTTP.ToPointer(),
            ScrapePath: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfSecret.ToPointer(),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            UsePublicIP: criblcontrolplanesdkgo.Pointer(true),
            SearchFilter: []components.ItemsTypeSearchFilter{
                components.ItemsTypeSearchFilter{
                    Name: "<value>",
                    Values: []string{},
                },
            },
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions1V2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](5682.66),
            Username: criblcontrolplanesdkgo.Pointer("Delia_Emmerich-Hamill36"),
            Password: criblcontrolplanesdkgo.Pointer("Y7ZzsJA55UXc4BZ"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateLogLevel: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
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
### Example Usage: InputCreateExamplesPrometheusRw

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesPrometheusRw" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputPrometheusRw(
        components.InputPrometheusRw{
            ID: criblcontrolplanesdkgo.Pointer("prometheus-rw-source"),
            Type: components.InputPrometheusRwTypePrometheusRw,
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Host: "0.0.0.0",
            Port: 10080,
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](6875.17),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](794108),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](9587.33),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](1496.41),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](9308.64),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](2890.6),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            PrometheusAPI: "/write",
            AuthType: components.AuthenticationTypeOptionsPrometheusAuthTextSecret.ToPointer(),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("endow chainstay bracelet around plus"),
            Username: criblcontrolplanesdkgo.Pointer("Anastasia.Lesch69"),
            Password: criblcontrolplanesdkgo.Pointer("SRxtqO4Gdh5YJEX"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplatePrometheusAPI: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesRawUdp

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesRawUdp" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputRawUDP(
        components.InputRawUDP{
            ID: criblcontrolplanesdkgo.Pointer("raw-udp-source"),
            Type: components.InputRawUDPTypeRawUDP,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 514,
            MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1843.53),
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            SingleMsgUDPPackets: criblcontrolplanesdkgo.Pointer(false),
            IngestRawBytes: criblcontrolplanesdkgo.Pointer(true),
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](9334.16),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("healthily following redact worth aha yuck baseboard hold crazy"),
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
### Example Usage: InputCreateExamplesS3

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesS3" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputS3(
        components.InputS3{
            ID: criblcontrolplanesdkgo.Pointer("s3-source"),
            Type: components.InputS3TypeS3,
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
            QueueName: "s3-notifications-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            BreakerRulesets: []string{
                "<value 1>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](342.52),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](7313.83),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](1095.75),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](9862.68),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](8762.59),
            SkipOnError: criblcontrolplanesdkgo.Pointer(true),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](8770.51),
            EnableSQSAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            Preprocess: &components.PreprocessType{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](2205.22),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](8634.03),
            Checkpointing: &components.CheckpointingType{
                Enabled: false,
                Retries: criblcontrolplanesdkgo.Pointer[float64](1555.53),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](177.43),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            TagAfterProcessing: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("search joshingly than gullible spectacles consequently pish giggle"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ProcessedTagKey: criblcontrolplanesdkgo.Pointer("<value>"),
            ProcessedTagValue: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateQueueName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesS3Inventory

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesS3Inventory" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputS3Inventory(
        components.InputS3Inventory{
            ID: criblcontrolplanesdkgo.Pointer("s3-inventory-source"),
            Type: components.InputS3InventoryTypeS3Inventory,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "s3-inventory-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            BreakerRulesets: []string{
                "<value 1>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](5846.84),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](4965.89),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](5753.06),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](3855.04),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](226.54),
            SkipOnError: criblcontrolplanesdkgo.Pointer(false),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3140.74),
            EnableSQSAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            Preprocess: &components.PreprocessType{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](5540.32),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](4211.1),
            Checkpointing: &components.CheckpointingType{
                Enabled: false,
                Retries: criblcontrolplanesdkgo.Pointer[float64](1555.53),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](9620.31),
            ChecksumSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxManifestSizeKB: criblcontrolplanesdkgo.Pointer[int64](517204),
            ValidateInventoryFiles: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("whoa beard mythology"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TagAfterProcessing: components.TagAfterProcessingOptionsFalse.ToPointer(),
            ProcessedTagKey: criblcontrolplanesdkgo.Pointer("<value>"),
            ProcessedTagValue: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateQueueName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesSecurityLake

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesSecurityLake" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputSecurityLake(
        components.InputSecurityLake{
            ID: criblcontrolplanesdkgo.Pointer("security-lake-source"),
            Type: components.InputSecurityLakeTypeSecurityLake,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            QueueName: "security-lake-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](2414.36),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](9764.93),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](8497.5),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](9501.16),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](9459.68),
            SkipOnError: criblcontrolplanesdkgo.Pointer(false),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3820.59),
            EnableSQSAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            Preprocess: &components.PreprocessType{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](6915.79),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](118.76),
            Checkpointing: &components.CheckpointingType{
                Enabled: false,
                Retries: criblcontrolplanesdkgo.Pointer[float64](1555.53),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](428.49),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("internal likewise aha whose those upwardly gosh sharply"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TagAfterProcessing: components.TagAfterProcessingOptionsTrue.ToPointer(),
            ProcessedTagKey: criblcontrolplanesdkgo.Pointer("<value>"),
            ProcessedTagValue: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateQueueName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesSnmp

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesSnmp" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputSnmp(
        components.InputSnmp{
            ID: criblcontrolplanesdkgo.Pointer("snmp-source"),
            Type: components.InputSnmpTypeSnmp,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "192.168.1.1",
            Port: 161,
            SnmpV3Auth: &components.SNMPv3Authentication{
                V3AuthEnabled: true,
                AllowUnmatchedTrap: criblcontrolplanesdkgo.Pointer(false),
                V3Users: []components.InputSnmpV3User{
                    components.InputSnmpV3User{
                        Name: "<value>",
                        AuthProtocol: components.AuthenticationProtocolOptionsV3UserSha384.ToPointer(),
                        AuthKey: criblcontrolplanesdkgo.Pointer("<value>"),
                        PrivProtocol: components.PrivacyProtocolDes.ToPointer(),
                        PrivKey: criblcontrolplanesdkgo.Pointer("<value>"),
                    },
                },
            },
            MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](6656.71),
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](5082.57),
            VarbindsWithTypes: criblcontrolplanesdkgo.Pointer(true),
            BestEffortParsing: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("gah blue majestically hence"),
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
### Example Usage: InputCreateExamplesSplunk

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesSplunk" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputSplunk(
        components.InputSplunk{
            ID: criblcontrolplanesdkgo.Pointer("splunk-source"),
            Type: components.InputSplunkTypeSplunk,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 9997,
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
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](9578.9),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](4638.02),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](1229.66),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](4340.71),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](7556.18),
            AuthTokens: []components.InputSplunkAuthToken{
                components.InputSplunkAuthToken{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("reopen blond yawningly huzzah yuck armchair ugh"),
                },
            },
            MaxS2Sversion: components.MaxS2SVersionV3.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("than sentimental happy-go-lucky tempting"),
            UseFwdTimezone: criblcontrolplanesdkgo.Pointer(true),
            DropControlFields: criblcontrolplanesdkgo.Pointer(true),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(false),
            Compress: components.InputSplunkCompressionAlways.ToPointer(),
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
### Example Usage: InputCreateExamplesSplunkHec

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesSplunkHec" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputSplunkHec(
        components.InputSplunkHec{
            ID: criblcontrolplanesdkgo.Pointer("splunk-hec-source"),
            Type: components.InputSplunkHecTypeSplunkHec,
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
                Mode: components.ModeOptionsPqSmart.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 8088,
            AuthTokens: []components.InputSplunkHecAuthToken{
                components.InputSplunkHecAuthToken{
                    AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
                    TokenSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Token: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("better lest pfft during fervently dulcimer selfish custom whale somber"),
                    AllowedIndexesAtToken: []string{
                        "<value 1>",
                        "<value 2>",
                        "<value 3>",
                    },
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](3944.77),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](799168),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](6427.77),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](8822.66),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](1922.13),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](924.02),
            EnableHealthCheck: "<value>",
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            SplunkHecAPI: "/services/collector",
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AllowedIndexes: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            SplunkHecAcks: criblcontrolplanesdkgo.Pointer(true),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](9829.02),
            UseFwdTimezone: criblcontrolplanesdkgo.Pointer(false),
            DropControlFields: criblcontrolplanesdkgo.Pointer(true),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(false),
            AccessControlAllowOrigin: []string{
                "<value 1>",
                "<value 2>",
            },
            AccessControlAllowHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            EmitTokenMetrics: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("torn whenever who fortunately whenever"),
            TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateSplunkHecAPI: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesSplunkSearch

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesSplunkSearch" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputSplunkSearch(
        components.InputSplunkSearch{
            ID: criblcontrolplanesdkgo.Pointer("splunk-search-source"),
            Type: components.InputSplunkSearchTypeSplunkSearch,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            SearchHead: "https://localhost:8089",
            Search: "index=main",
            Earliest: criblcontrolplanesdkgo.Pointer("<value>"),
            Latest: criblcontrolplanesdkgo.Pointer("<value>"),
            CronSchedule: "0 * * * *",
            Endpoint: "/services/search/jobs/export",
            OutputMode: components.OutputModeOptionsSplunkCollectorConfJSON,
            EndpointParams: []components.EndpointParam{
                components.EndpointParam{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            EndpointHeaders: []components.EndpointHeader{
                components.EndpointHeader{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            LogLevel: components.InputSplunkSearchLogLevelWarn.ToPointer(),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](6006.12),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](1901.73),
            JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](244.27),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            RetryRules: &components.RetryRulesType{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesNone,
                Interval: criblcontrolplanesdkgo.Pointer[float64](8023.69),
                Limit: criblcontrolplanesdkgo.Pointer[float64](8538.93),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](6735.17),
                Codes: []float64{
                    3551.76,
                    3444.04,
                    2558.51,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](1276.17),
            AuthType: components.InputSplunkSearchAuthenticationTypeBasic.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("below mmm bulky gee approach phew monthly controvert volunteer if"),
            Username: criblcontrolplanesdkgo.Pointer("Milton.Gottlieb"),
            Password: criblcontrolplanesdkgo.Pointer("1Xwpask7BiTVNBz"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesSqs

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesSqs" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputSqs(
        components.InputSqs{
            ID: criblcontrolplanesdkgo.Pointer("sqs-source"),
            Type: components.InputSqsTypeSqs,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "my-queue",
            QueueType: components.InputSqsQueueTypeStandard,
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            CreateQueue: criblcontrolplanesdkgo.Pointer(false),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions3V4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](9240.08),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](8942.09),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](8052.76),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](9292.66),
            Description: criblcontrolplanesdkgo.Pointer("outlaw nervously past"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](8601.43),
            TemplateQueueName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesSyslog

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesSyslog" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputSyslog(
        components.CreateInputSyslogInputSyslogSyslog1(
            components.InputSyslogSyslog1{
                ID: criblcontrolplanesdkgo.Pointer("syslog-source"),
                Type: components.InputSyslogType1Syslog,
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
                    Mode: components.ModeOptionsPqSmart.ToPointer(),
                    MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                    CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                    MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                    Compress: components.CompressionOptionsPqGzip.ToPointer(),
                    PqControls: &components.PqTypePqControls{},
                },
                Host: "0.0.0.0",
                UDPPort: 514,
                TCPPort: criblcontrolplanesdkgo.Pointer[float64](1690.04),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](835.96),
                IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
                TimestampTimezone: criblcontrolplanesdkgo.Pointer("<value>"),
                SingleMsgUDPPackets: criblcontrolplanesdkgo.Pointer(false),
                EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
                KeepFieldsList: []string{
                    "<value 1>",
                    "<value 2>",
                },
                OctetCounting: criblcontrolplanesdkgo.Pointer(true),
                InferFraming: criblcontrolplanesdkgo.Pointer(true),
                StrictlyInferOctetCounting: criblcontrolplanesdkgo.Pointer(true),
                AllowNonStandardAppName: criblcontrolplanesdkgo.Pointer(false),
                MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](3622.72),
                SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](193.15),
                SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](3247.34),
                SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](688.7),
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
                Metadata: []components.ItemsTypeMetadata{
                    components.ItemsTypeMetadata{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](1084.02),
                EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(false),
                Description: criblcontrolplanesdkgo.Pointer("mysteriously concerning yahoo however ick kiss downchange handover out when"),
                EnableEnhancedProxyHeaderParsing: criblcontrolplanesdkgo.Pointer(true),
                TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                TemplateUDPPort: criblcontrolplanesdkgo.Pointer("<value>"),
                TemplateTCPPort: criblcontrolplanesdkgo.Pointer("<value>"),
            },
        ),
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInput != nil {
        // handle response
    }
}
```
### Example Usage: InputCreateExamplesSystemMetrics

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesSystemMetrics" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputSystemMetrics(
        components.InputSystemMetrics{
            ID: criblcontrolplanesdkgo.Pointer("system-metrics-source"),
            Type: components.InputSystemMetricsTypeSystemMetrics,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](5017.67),
            Host: &components.InputSystemMetricsHost{
                Mode: components.ModeOptionsHostCustom.ToPointer(),
                Custom: &components.InputSystemMetricsCustom{
                    System: &components.InputSystemMetricsSystem{
                        Mode: components.InputSystemMetricsSystemModeDisabled.ToPointer(),
                        Processes: criblcontrolplanesdkgo.Pointer(false),
                    },
                    CPU: &components.InputSystemMetricsCPU{
                        Mode: components.InputSystemMetricsCPUModeCustom.ToPointer(),
                        PerCPU: criblcontrolplanesdkgo.Pointer(true),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Time: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Memory: &components.InputSystemMetricsMemory{
                        Mode: components.InputSystemMetricsMemoryModeDisabled.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Network: &components.InputSystemMetricsNetwork{
                        Mode: components.InputSystemMetricsNetworkModeDisabled.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(true),
                        Protocols: criblcontrolplanesdkgo.Pointer(true),
                        Devices: []string{
                            "<value 1>",
                        },
                        PerInterface: criblcontrolplanesdkgo.Pointer(true),
                    },
                    Disk: &components.InputSystemMetricsDisk{
                        Mode: components.InputSystemMetricsDiskModeCustom.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Inodes: criblcontrolplanesdkgo.Pointer(false),
                        Devices: []string{
                            "<value 1>",
                        },
                        Mountpoints: []string{
                            "<value 1>",
                            "<value 2>",
                        },
                        Fstypes: []string{
                            "<value 1>",
                            "<value 2>",
                        },
                        PerDevice: criblcontrolplanesdkgo.Pointer(false),
                    },
                },
            },
            Process: &components.ProcessType{
                Sets: []components.ItemsTypeProcessSets{
                    components.ItemsTypeProcessSets{
                        Name: "<value>",
                        Filter: "<value>",
                        IncludeChildren: criblcontrolplanesdkgo.Pointer(true),
                    },
                },
            },
            Container: &components.Container{
                Mode: components.ContainerModeBasic.ToPointer(),
                DockerSocket: []string{
                    "<value 1>",
                    "<value 2>",
                },
                DockerTimeout: criblcontrolplanesdkgo.Pointer[float64](7531.98),
                Filters: []components.InputSystemMetricsFilter{
                    components.InputSystemMetricsFilter{
                        Expr: "<value>",
                    },
                },
                AllContainers: criblcontrolplanesdkgo.Pointer(true),
                PerDevice: criblcontrolplanesdkgo.Pointer(false),
                Detail: criblcontrolplanesdkgo.Pointer(true),
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Persistence: &components.InputSystemMetricsPersistence{
                Enable: criblcontrolplanesdkgo.Pointer(true),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.DataCompressionFormatOptionsPersistenceNone.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            Description: criblcontrolplanesdkgo.Pointer("worth alongside till during pharmacopoeia psst broadly"),
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
### Example Usage: InputCreateExamplesSystemState

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesSystemState" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputSystemState(
        components.InputSystemState{
            ID: criblcontrolplanesdkgo.Pointer("system-state-source"),
            Type: components.InputSystemStateTypeSystemState,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](3542.37),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Collectors: &components.Collectors{
                Hostsfile: &components.HostsFile{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Interfaces: &components.Interfaces{
                    Enable: criblcontrolplanesdkgo.Pointer(false),
                },
                Disk: &components.DisksAndFileSystems{
                    Enable: criblcontrolplanesdkgo.Pointer(false),
                },
                Metadata: &components.HostInfo{
                    Enable: criblcontrolplanesdkgo.Pointer(false),
                },
                Routes: &components.InputSystemStateRoutes{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                DNS: &components.DNS{
                    Enable: criblcontrolplanesdkgo.Pointer(false),
                },
                User: &components.UsersAndGroups{
                    Enable: criblcontrolplanesdkgo.Pointer(false),
                },
                Firewall: &components.Firewall{
                    Enable: criblcontrolplanesdkgo.Pointer(false),
                },
                Services: &components.Services{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Ports: &components.ListeningPorts{
                    Enable: criblcontrolplanesdkgo.Pointer(false),
                },
                LoginUsers: &components.LoggedInUsers{
                    Enable: criblcontrolplanesdkgo.Pointer(false),
                },
            },
            Persistence: &components.InputSystemStatePersistence{
                Enable: criblcontrolplanesdkgo.Pointer(true),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            DisableNativeModule: criblcontrolplanesdkgo.Pointer(true),
            DisableNativeLastLogModule: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("bustling overconfidently gleefully hmph slimy aha ultimately that as boohoo"),
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
### Example Usage: InputCreateExamplesTcp

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesTcp" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputTCP(
        components.InputTCP{
            ID: criblcontrolplanesdkgo.Pointer("tcp-source"),
            Type: components.InputTCPTypeTCP,
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
            Host: "0.0.0.0",
            Port: 10090,
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
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](1120.04),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](5926.54),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](3298.44),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](6550.73),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            BreakerRulesets: []string{
                "<value 1>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](6014.24),
            EnableHeader: criblcontrolplanesdkgo.Pointer(true),
            Preprocess: &components.PreprocessType{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("failing behind uh-huh into mentor which cop-out"),
            AuthToken: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
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
### Example Usage: InputCreateExamplesTcpjson

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesTcpjson" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputTcpjson(
        components.InputTcpjson{
            ID: criblcontrolplanesdkgo.Pointer("tcpjson-source"),
            Type: components.InputTcpjsonTypeTcpjson,
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Host: "0.0.0.0",
            Port: 10090,
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
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](8150.03),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](4509.33),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](8256.64),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](6518.95),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(true),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("up for weary although obediently inwardly"),
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
### Example Usage: InputCreateExamplesWef

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesWef" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputWef(
        components.InputWef{
            ID: criblcontrolplanesdkgo.Pointer("wef-source"),
            Type: components.InputWefTypeWef,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Host: "0.0.0.0",
            Port: 5985,
            AuthMethod: components.InputWefAuthenticationMethodKerberos.ToPointer(),
            TLS: &components.MTLSSettings{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                RequestCert: criblcontrolplanesdkgo.Pointer(true),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: "<value>",
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: "<value>",
                CaPath: "<value>",
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
                OcspCheck: criblcontrolplanesdkgo.Pointer(false),
                Keytab: "<value>",
                Principal: "<value>",
                OcspCheckFailClose: criblcontrolplanesdkgo.Pointer(true),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](7309.03),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](798308),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](6011.4),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](5503.51),
            CaFingerprint: criblcontrolplanesdkgo.Pointer("<value>"),
            Keytab: criblcontrolplanesdkgo.Pointer("<value>"),
            Principal: criblcontrolplanesdkgo.Pointer("<value>"),
            AllowMachineIDMismatch: criblcontrolplanesdkgo.Pointer(true),
            Subscriptions: []components.Subscription{
                components.Subscription{
                    SubscriptionName: "subscription-1",
                    Version: criblcontrolplanesdkgo.Pointer("<value>"),
                    ContentFormat: components.InputWefFormatRenderedText,
                    HeartbeatInterval: 60,
                    BatchTimeout: 5,
                    ReadExistingEvents: criblcontrolplanesdkgo.Pointer(true),
                    SendBookmarks: criblcontrolplanesdkgo.Pointer(false),
                    Compress: criblcontrolplanesdkgo.Pointer(false),
                    Targets: []string{},
                    Locale: criblcontrolplanesdkgo.Pointer("fi"),
                    QuerySelector: components.QueryBuilderModeSimple.ToPointer(),
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                    Queries: []components.Query{
                        components.Query{
                            Path: "/lost+found",
                            QueryExpression: "<value>",
                        },
                    },
                    XMLQuery: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("before next usually barring yuppify um pace yuck drat since"),
            LogFingerprintMismatch: criblcontrolplanesdkgo.Pointer(false),
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
### Example Usage: InputCreateExamplesWinEventLogs

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesWinEventLogs" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputWinEventLogs(
        components.InputWinEventLogs{
            ID: criblcontrolplanesdkgo.Pointer("win-event-logs-source"),
            Type: components.InputWinEventLogsTypeWinEventLogs,
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
            LogNames: []string{
                "Application",
                "System",
            },
            ReadMode: components.ReadModeOldest.ToPointer(),
            EventFormat: components.EventFormatJSON.ToPointer(),
            DisableNativeModule: criblcontrolplanesdkgo.Pointer(true),
            Interval: criblcontrolplanesdkgo.Pointer[float64](161.53),
            BatchSize: criblcontrolplanesdkgo.Pointer[float64](2676.97),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            MaxEventBytes: criblcontrolplanesdkgo.Pointer[float64](3989.15),
            Description: criblcontrolplanesdkgo.Pointer("outdo pear amidst unscramble forenenst mid analyse cow"),
            DisableJSONRendering: criblcontrolplanesdkgo.Pointer(true),
            DisableXMLRendering: criblcontrolplanesdkgo.Pointer(true),
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
### Example Usage: InputCreateExamplesWindowsMetrics

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesWindowsMetrics" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputWindowsMetrics(
        components.InputWindowsMetrics{
            ID: criblcontrolplanesdkgo.Pointer("windows-metrics-source"),
            Type: components.InputWindowsMetricsTypeWindowsMetrics,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](7895.52),
            Host: &components.InputWindowsMetricsHost{
                Mode: components.ModeOptionsHostAll.ToPointer(),
                Custom: &components.InputWindowsMetricsCustom{
                    System: &components.InputWindowsMetricsSystem{
                        Mode: components.InputWindowsMetricsSystemModeCustom.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(true),
                    },
                    CPU: &components.InputWindowsMetricsCPU{
                        Mode: components.InputWindowsMetricsCPUModeBasic.ToPointer(),
                        PerCPU: criblcontrolplanesdkgo.Pointer(true),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Time: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Memory: &components.InputWindowsMetricsMemory{
                        Mode: components.InputWindowsMetricsMemoryModeCustom.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(true),
                    },
                    Network: &components.InputWindowsMetricsNetwork{
                        Mode: components.InputWindowsMetricsNetworkModeBasic.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Protocols: criblcontrolplanesdkgo.Pointer(false),
                        Devices: []string{
                            "<value 1>",
                            "<value 2>",
                            "<value 3>",
                        },
                        PerInterface: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Disk: &components.InputWindowsMetricsDisk{
                        Mode: components.InputWindowsMetricsDiskModeDisabled.ToPointer(),
                        PerVolume: criblcontrolplanesdkgo.Pointer(false),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Volumes: []string{
                            "<value 1>",
                            "<value 2>",
                        },
                    },
                },
            },
            Process: &components.ProcessType{
                Sets: []components.ItemsTypeProcessSets{
                    components.ItemsTypeProcessSets{
                        Name: "<value>",
                        Filter: "<value>",
                        IncludeChildren: criblcontrolplanesdkgo.Pointer(true),
                    },
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Persistence: &components.InputWindowsMetricsPersistence{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            DisableNativeModule: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("shanghai bonfire cauliflower indeed"),
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
### Example Usage: InputCreateExamplesWiz

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesWiz" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputWiz(
        components.InputWiz{
            ID: criblcontrolplanesdkgo.Pointer("wiz-source"),
            Type: components.InputWizTypeWiz,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Endpoint: "https://api.wiz.io",
            AuthURL: "https://auth.wiz.io/oauth/token",
            AuthAudienceOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            ClientID: "client-id",
            ContentConfig: []components.InputWizContentConfig{},
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](968.17),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](3750.97),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](9118.86),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            RetryRules: &components.RetryRulesType{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesNone,
                Interval: criblcontrolplanesdkgo.Pointer[float64](8023.69),
                Limit: criblcontrolplanesdkgo.Pointer[float64](8538.93),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](6735.17),
                Codes: []float64{
                    3551.76,
                    3444.04,
                    2558.51,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            AuthType: components.AuthenticationMethodOptions1Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("madly whoa subtle coop or stratify"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateEndpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAuthURL: criblcontrolplanesdkgo.Pointer("https://shocked-responsibility.info"),
            TemplateClientID: criblcontrolplanesdkgo.Pointer("<id>"),
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
### Example Usage: InputCreateExamplesWizWebhook

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesWizWebhook" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputWizWebhook(
        components.InputWizWebhook{
            ID: criblcontrolplanesdkgo.Pointer("wiz-webhook-source"),
            Type: components.InputWizWebhookTypeWizWebhook,
            Disabled: criblcontrolplanesdkgo.Pointer(false),
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Host: "0.0.0.0",
            Port: 10080,
            AuthTokens: []string{
                "<value 1>",
                "<value 2>",
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](8977.16),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](153946),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](4062.3),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](4871.45),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](5909.3),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](634.62),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](3516.58),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AllowedPaths: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            AllowedMethods: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            AuthTokensExt: []components.ItemsTypeAuthTokensExt{
                components.ItemsTypeAuthTokensExt{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("ew phooey unless aha plus woot zowie unless"),
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("trusty coast furthermore finally so ape mystify impressionable"),
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
### Example Usage: InputCreateExamplesZscalerHec

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesZscalerHec" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputZscalerHec(
        components.InputZscalerHec{
            ID: criblcontrolplanesdkgo.Pointer("zscaler-hec-source"),
            Type: components.InputZscalerHecTypeZscalerHec,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 8088,
            AuthTokens: []components.InputZscalerHecAuthToken{
                components.InputZscalerHecAuthToken{
                    AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
                    TokenSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Token: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                    Description: criblcontrolplanesdkgo.Pointer("versus an thankfully and or inside playfully whether daintily"),
                    AllowedIndexesAtToken: []string{
                        "<value 1>",
                        "<value 2>",
                    },
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](77.42),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](147184),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](5157.16),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](3819.67),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](4473.94),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](665.59),
            EnableHealthCheck: "<value>",
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            HecAPI: "/services/collector",
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AllowedIndexes: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            HecAcks: criblcontrolplanesdkgo.Pointer(true),
            AccessControlAllowOrigin: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            AccessControlAllowHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            EmitTokenMetrics: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("valiantly and whisper"),
            TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateHecAPI: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputExamplesCribl

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputExamplesCribl" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputCribl(
        components.InputCribl{
            ID: criblcontrolplanesdkgo.Pointer("cribl-source"),
            Type: components.InputCriblTypeCribl,
            Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Filter: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("wherever qualified questionably fork amid"),
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
### Example Usage: InputExamplesCriblMetrics

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputExamplesCriblMetrics" -->
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

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputCriblmetrics(
        components.InputCriblmetrics{
            ID: criblcontrolplanesdkgo.Pointer("cribl-metrics-source"),
            Type: components.InputCriblmetricsTypeCriblmetrics,
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
                Mode: components.ModeOptionsPqSmart.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5318.64),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](3762.63),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/sbin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Prefix: criblcontrolplanesdkgo.Pointer("<value>"),
            FullFidelity: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("word what fall expense tromp nerve"),
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