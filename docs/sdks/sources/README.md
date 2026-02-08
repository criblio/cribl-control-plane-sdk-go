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

### Example Usage: InputCreateExamplesAppscope

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesAppscope" -->
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
### Example Usage: InputCreateExamplesAzureBlob

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesAzureBlob" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestAzureBlob(
        operations.InputAzureBlob{
            ID: "azure-blob-source",
            Type: operations.CreateInputTypeAzureBlobAzureBlob,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "azure-blob-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("<value>"),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](9845.2),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](7841.68),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](2786.75),
            ServicePeriodSecs: criblcontrolplanesdkgo.Pointer[float64](3885.37),
            SkipOnError: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](6690.08),
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](8742.16),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](5080),
            AuthType: components.AuthenticationMethodOptionsSecret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("foodstuffs consequently cannibalise like an knowingly duffel gadzooks"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesCloudflareHec" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestCloudflareHec(
        operations.InputCloudflareHec{
            ID: "cloudflare-hec-source",
            Type: operations.TypeCloudflareHecCloudflareHec,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 8088,
            AuthTokens: []operations.AuthTokenCloudflareHec{
                operations.AuthTokenCloudflareHec{
                    AuthType: operations.AuthTokenAuthenticationMethodManual.ToPointer(),
                    TokenSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("culture reflate following hateful pike whose"),
                    AllowedIndexesAtToken: []string{
                        "<value 1>",
                        "<value 2>",
                    },
                    Metadata: []components.ItemsTypeNotificationMetadata{
                        components.ItemsTypeNotificationMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](6936.75),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](226532),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](9442.97),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](5240.66),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](9114.25),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](9265.08),
            EnableHealthCheck: "<value>",
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            HecAPI: "/services/collector",
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AllowedIndexes: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](7807.01),
            AccessControlAllowOrigin: []string{
                "<value 1>",
            },
            AccessControlAllowHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            EmitTokenMetrics: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("coaxingly probate slipper smoke aw meh of"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesCollection" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestCollection(
        operations.InputCollection{
            ID: "collection-source",
            Type: operations.TypeCollectionCollection,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](5267.49),
            Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                Disabled: false,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                },
            },
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesConfluentCloud" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestConfluentCloud(
        operations.InputConfluentCloud{
            ID: "confluent-cloud-source",
            Type: operations.CreateInputTypeConfluentCloudConfluentCloud,
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
            Brokers: []string{
                "pkc-xxxxx.us-east-1.aws.confluent.cloud:9092",
            },
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
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
            Topics: []string{
                "logs",
            },
            GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
            FromBeginning: criblcontrolplanesdkgo.Pointer(true),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType{
                Disabled: false,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://ugly-availability.info"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](8658.84),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](576.51),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](7446.19),
                Auth: &components.AuthTypeKafkaSchemaRegistry{
                    Disabled: false,
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                    Disabled: criblcontrolplanesdkgo.Pointer(true),
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
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](4267.71),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](3355.4),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](7721.27),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](7462.03),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9154.04),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2946.47),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](1024.51),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](4426.91),
            Sasl: &components.AuthenticationType{
                Disabled: true,
                Username: criblcontrolplanesdkgo.Pointer("Etha51"),
                Password: criblcontrolplanesdkgo.Pointer("XukwT9HxMYXOSPg"),
                AuthType: components.AuthenticationMethodOptionsSaslSecret.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslKerberos.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(false),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://swift-laughter.com"),
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
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](449.45),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](8891),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](6176.85),
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](4138.47),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](6728.18),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](5098.28),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](1045.06),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](7315.82),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("violently throughout upwardly"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesCriblHttp" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestCriblHTTP(
        operations.InputCriblHTTP{
            ID: "cribl-http-source",
            Type: operations.CreateInputTypeCriblHTTPCriblHTTP,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10080,
            AuthTokens: []components.ItemsTypeAuthTokens{
                components.ItemsTypeAuthTokens{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("with blah omelet pfft oddly bid regarding smuggle yuck goat"),
                },
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](2294.34),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](536087),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](8312.41),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](8291.86),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](8516.65),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](4981.97),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("ha teriyaki prance"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesCriblLakeHttp" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestCriblLakeHTTP(
        operations.InputCriblLakeHTTP{
            ID: "cribl-lake-http-source",
            Type: operations.TypeCriblLakeHTTPCriblLakeHTTP,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](6791.17),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](816175),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](7079.71),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](5856.13),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](3778.34),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](4936.4),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            CriblAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            ElasticAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            SplunkHecAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            SplunkHecAcks: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthTokensExt: []operations.AuthTokensExt{
                operations.AuthTokensExt{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("neck below since blah whoa ridge"),
                    Metadata: []components.ItemsTypeNotificationMetadata{
                        components.ItemsTypeNotificationMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                    SplunkHecMetadata: &operations.SplunkHecMetadata{
                        Enabled: criblcontrolplanesdkgo.Pointer(false),
                        DefaultDataset: criblcontrolplanesdkgo.Pointer("<value>"),
                        AllowedIndexesAtToken: []string{
                            "<value 1>",
                        },
                    },
                    ElasticsearchMetadata: &operations.ElasticsearchMetadata{
                        Enabled: criblcontrolplanesdkgo.Pointer(true),
                        DefaultDataset: criblcontrolplanesdkgo.Pointer("<value>"),
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("entice substantiate bossy"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesCriblTcp" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestCriblTCP(
        operations.InputCriblTCP{
            ID: "cribl-tcp-source",
            Type: operations.CreateInputTypeCriblTCPCriblTCP,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10090,
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
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](5339.54),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](7725.89),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](7991.94),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](3827.27),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(true),
            AuthTokens: []components.ItemsTypeAuthTokens{
                components.ItemsTypeAuthTokens{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("with blah omelet pfft oddly bid regarding smuggle yuck goat"),
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("cheerfully average lovingly meh heartfelt leading scratchy make courageously"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesCrowdstrike" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestCrowdstrike(
        operations.InputCrowdstrike{
            ID: "crowdstrike-source",
            Type: operations.TypeCrowdstrikeCrowdstrike,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "crowdstrike-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](8060.83),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](1997.36),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](5226.76),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](8318.28),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](2885.87),
            SkipOnError: criblcontrolplanesdkgo.Pointer(false),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](1376.22),
            EnableSQSAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                Disabled: false,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                },
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Checkpointing: &components.CheckpointingType{
                Enabled: false,
                Retries: criblcontrolplanesdkgo.Pointer[float64](7918.57),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](9197.78),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("brr zany ha understanding upbeat consequently invite newsstand though"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TagAfterProcessing: components.TagAfterProcessingOptionsTrue.ToPointer(),
            ProcessedTagKey: criblcontrolplanesdkgo.Pointer("<value>"),
            ProcessedTagValue: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesDatadogAgent" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestDatadogAgent(
        operations.InputDatadogAgent{
            ID: "datadog-agent-source",
            Type: operations.TypeDatadogAgentDatadogAgent,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 8126,
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](8901.35),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](323919),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](9201.65),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](8623.93),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](2668.82),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5403.04),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ProxyMode: &operations.ProxyModeDatadogAgent{
                Enabled: true,
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            },
            Description: criblcontrolplanesdkgo.Pointer("meh vivacious geez however knight ah whenever bulky"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesDatagen" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestDatagen(
        operations.InputDatagen{
            ID: "datagen-source",
            Type: operations.TypeDatagenDatagen,
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
            Samples: []operations.Sample{
                operations.Sample{
                    Sample: "sample.json",
                    EventsPerSec: 10,
                },
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("practical briefly through"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesEdgePrometheus" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestEdgePrometheus(
        operations.InputEdgePrometheus{
            ID: "edge-prometheus-source",
            Type: operations.TypeEdgePrometheusEdgePrometheus,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            DimensionList: []string{
                "<value 1>",
            },
            DiscoveryType: operations.DiscoveryTypeEdgePrometheusStatic,
            Interval: 60,
            Timeout: criblcontrolplanesdkgo.Pointer[float64](1523.25),
            Persistence: &components.DiskSpoolingType{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.CompressionOptionsPersistenceNone.ToPointer(),
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthType: operations.AuthenticationMethodEdgePrometheusSecret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("forenenst submitter individual french though analogy homeschool"),
            Targets: []operations.Target{
                operations.Target{
                    Protocol: components.ProtocolOptionsTargetsItemsHTTP.ToPointer(),
                    Host: "localhost",
                    Port: criblcontrolplanesdkgo.Pointer[float64](5892.7),
                    Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                },
            },
            RecordType: components.RecordTypeOptionsAaaa.ToPointer(),
            ScrapePort: criblcontrolplanesdkgo.Pointer[float64](9784.75),
            NameList: []string{
                "<value 1>",
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
                    Values: []string{
                        "<value 1>",
                        "<value 2>",
                    },
                },
            },
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions1V2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](7529.29),
            ScrapeProtocolExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            ScrapePortExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            ScrapePathExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            PodFilter: []operations.PodFilter{
                operations.PodFilter{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("true baseboard difficult accidentally even vainly pendant appropriate supposing meh"),
                },
            },
            Username: criblcontrolplanesdkgo.Pointer("Bret.Satterfield48"),
            Password: criblcontrolplanesdkgo.Pointer("wQ8ho0Xc5FvvRdi"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesElastic" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestElastic(
        operations.InputElastic{
            ID: "elastic-source",
            Type: operations.CreateInputTypeElasticElastic,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "localhost",
            Port: 9200,
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](5262.84),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](560233),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](6514.74),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](9117.6),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](8372.68),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](8300.75),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            ElasticAPI: "/",
            AuthType: operations.AuthenticationTypeElasticAuthTokens.ToPointer(),
            APIVersion: operations.CreateInputAPIVersionCustom.ToPointer(),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ProxyMode: &operations.ProxyModeElastic{
                Enabled: false,
                AuthType: operations.ProxyModeAuthenticationMethodSecret.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("Nathanial_Vandervort"),
                Password: criblcontrolplanesdkgo.Pointer("CBiRV0jzSu7ZO46"),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                URL: criblcontrolplanesdkgo.Pointer("https://minty-diver.name"),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                RemoveHeaders: []string{
                    "<value 1>",
                },
                TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](7619.36),
            },
            Description: criblcontrolplanesdkgo.Pointer("mmm digital beneath summarise"),
            Username: criblcontrolplanesdkgo.Pointer("Melyna_Wuckert61"),
            Password: criblcontrolplanesdkgo.Pointer("kjFO782PUEcyGr_"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthTokens: []string{
                "<value 1>",
            },
            CustomAPIVersion: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesEventhub" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestEventhub(
        operations.InputEventhub{
            ID: "eventhub-source",
            Type: operations.TypeEventhubEventhub,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
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
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](4412.63),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](4590.77),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](8313.38),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](3799.93),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1213.03),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6524.55),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](5485.95),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](5062.18),
            Sasl: &components.AuthenticationType1{
                Disabled: false,
                AuthType: components.AuthenticationMethodOptionsSasl1Manual.ToPointer(),
                Password: criblcontrolplanesdkgo.Pointer("HuXGq8RA_UbNNDq"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSasl1Plain.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("Reina.Gutmann7"),
                ClientSecretAuthType: components.AuthenticationMethodOptionsSasl2Manual.ToPointer(),
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
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](7156),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](5931.63),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](1103.94),
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](2956.72),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](8644.37),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](2461.93),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](333.36),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](981.32),
            MinimizeDuplicates: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("scale mmm lecture tag overfeed out nucleotidase"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesExec" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestExec(
        operations.InputExec{
            ID: "exec-source",
            Type: operations.InputExecTypeExec,
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
            Command: "echo \"Hello World\"",
            Retries: criblcontrolplanesdkgo.Pointer[float64](8796.61),
            ScheduleType: operations.ScheduleTypeInterval.ToPointer(),
            BreakerRulesets: []string{
                "<value 1>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](3310.64),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("weary throughout rotten um even"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesFile" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestFile(
        operations.InputFile{
            ID: "file-source",
            Type: operations.InputFileTypeFile,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Mode: operations.InputFileModeManual.ToPointer(),
            Interval: criblcontrolplanesdkgo.Pointer[float64](1166.22),
            Filenames: []string{
                "<value 1>",
                "<value 2>",
            },
            FilterArchivedFiles: criblcontrolplanesdkgo.Pointer(true),
            TailOnly: criblcontrolplanesdkgo.Pointer(true),
            IdleTimeout: criblcontrolplanesdkgo.Pointer[float64](2482.99),
            MinAgeDur: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxAgeDur: criblcontrolplanesdkgo.Pointer("<value>"),
            CheckFileModTime: criblcontrolplanesdkgo.Pointer(false),
            ForceText: criblcontrolplanesdkgo.Pointer(false),
            HashLen: criblcontrolplanesdkgo.Pointer[float64](7349.43),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](1810.26),
            Description: criblcontrolplanesdkgo.Pointer("sociable dime sympathetically provided jaywalk apud covenant"),
            Path: criblcontrolplanesdkgo.Pointer("/home/user"),
            Depth: criblcontrolplanesdkgo.Pointer[float64](8602.37),
            SuppressMissingPathErrors: criblcontrolplanesdkgo.Pointer(true),
            DeleteFiles: criblcontrolplanesdkgo.Pointer(false),
            SaltHash: criblcontrolplanesdkgo.Pointer(true),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesFirehose" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestFirehose(
        operations.InputFirehose{
            ID: "firehose-source",
            Type: operations.TypeFirehoseFirehose,
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
            Host: "0.0.0.0",
            Port: 10080,
            AuthTokens: []string{
                "<value 1>",
                "<value 2>",
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](6041.36),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](312291),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](6651.29),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](9017.19),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](2313.05),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5489.18),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("substantiate mortise yowza"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesGooglePubsub" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestGooglePubsub(
        operations.InputGooglePubsub{
            ID: "google-pubsub-source",
            Type: operations.CreateInputTypeGooglePubsubGooglePubsub,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            TopicName: "my-topic",
            SubscriptionName: "my-subscription",
            MonitorSubscription: criblcontrolplanesdkgo.Pointer(false),
            CreateTopic: criblcontrolplanesdkgo.Pointer(false),
            CreateSubscription: criblcontrolplanesdkgo.Pointer(true),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            GoogleAuthMethod: components.GoogleAuthenticationMethodOptionsAuto.ToPointer(),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxBacklog: criblcontrolplanesdkgo.Pointer[float64](2713.89),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](6818.77),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](1286.98),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("dramatic ouch if issue yuck diligent emotional"),
            OrderedDelivery: criblcontrolplanesdkgo.Pointer(true),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesGrafana" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestGrafana(
        operations.CreateInputGrafanaInputGrafanaGrafana1(
            operations.InputGrafanaGrafana1{
                ID: "grafana-source",
                Type: operations.InputGrafanaType1Grafana,
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
                    MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                    CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                    MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                    Compress: components.CompressionOptionsPqGzip.ToPointer(),
                    PqControls: &components.PqTypePqControls{},
                },
                Host: "0.0.0.0",
                Port: 10080,
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
                MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](4754.71),
                MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](22276),
                EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
                CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
                ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](1083.99),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](1977.59),
                SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](657.07),
                KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](835.12),
                EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
                IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
                IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
                PrometheusAPI: "/api/prom/push",
                LokiAPI: criblcontrolplanesdkgo.Pointer("<value>"),
                PrometheusAuth: &operations.PrometheusAuth1{
                    AuthType: components.AuthenticationTypeOptionsPrometheusAuthBasic.ToPointer(),
                    Username: criblcontrolplanesdkgo.Pointer("Delbert_Yost21"),
                    Password: criblcontrolplanesdkgo.Pointer("X0GrESwkQOCR5y1"),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    LoginURL: criblcontrolplanesdkgo.Pointer("https://mysterious-ruin.org"),
                    SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
                    Secret: criblcontrolplanesdkgo.Pointer("<value>"),
                    TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
                    AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
                    TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](8859.96),
                    OauthParams: []components.ItemsTypeOauthParams{
                        components.ItemsTypeOauthParams{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                    OauthHeaders: []components.ItemsTypeOauthHeaders{
                        components.ItemsTypeOauthHeaders{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
                LokiAuth: &operations.LokiAuth1{
                    AuthType: components.AuthenticationTypeOptionsLokiAuthToken.ToPointer(),
                    Username: criblcontrolplanesdkgo.Pointer("Einar.Sawayn"),
                    Password: criblcontrolplanesdkgo.Pointer("QnAU6grcOdTFUeJ"),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    LoginURL: criblcontrolplanesdkgo.Pointer("https://velvety-loyalty.net/"),
                    SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
                    Secret: criblcontrolplanesdkgo.Pointer("<value>"),
                    TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
                    AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
                    TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](8835.66),
                    OauthParams: []components.ItemsTypeOauthParams{
                        components.ItemsTypeOauthParams{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                    OauthHeaders: []components.ItemsTypeOauthHeaders{
                        components.ItemsTypeOauthHeaders{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
                Metadata: []components.ItemsTypeNotificationMetadata{
                    components.ItemsTypeNotificationMetadata{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                Description: criblcontrolplanesdkgo.Pointer("solution satirise approximate psst jogging realistic greatly boohoo who curry"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesHttp" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestHTTP(
        operations.InputHTTP{
            ID: "http-source",
            Type: operations.TypeHTTPHTTP,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](5123.99),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](286013),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](7569.11),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](2839.46),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](5092.72),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](6429.04),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            CriblAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            ElasticAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            SplunkHecAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            SplunkHecAcks: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthTokensExt: []components.ItemsTypeAuthTokensExt{
                components.ItemsTypeAuthTokensExt{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("deeply very fiercely at total truthfully how frugal ah"),
                    Metadata: []components.ItemsTypeNotificationMetadata{
                        components.ItemsTypeNotificationMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("voluminous lest guard distant ah custom masculinize multicolored proud"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesHttpRaw" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestHTTPRaw(
        operations.InputHTTPRaw{
            ID: "http-raw-source",
            Type: operations.TypeHTTPRawHTTPRaw,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](7394.97),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](313496),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](4954.6),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](5720.36),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](222.98),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](2239.85),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](1231.87),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AllowedPaths: []string{
                "<value 1>",
            },
            AllowedMethods: []string{
                "<value 1>",
            },
            AuthTokensExt: []components.ItemsTypeAuthTokensExt{
                components.ItemsTypeAuthTokensExt{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("deeply very fiercely at total truthfully how frugal ah"),
                    Metadata: []components.ItemsTypeNotificationMetadata{
                        components.ItemsTypeNotificationMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("provider yet unsung past"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesJournalFiles" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestJournalFiles(
        operations.InputJournalFiles{
            ID: "journal-files-source",
            Type: operations.InputJournalFilesTypeJournalFiles,
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
            Path: "/var/log/journal",
            Interval: criblcontrolplanesdkgo.Pointer[float64](5000.85),
            Journals: []string{
                "system",
            },
            Rules: []operations.InputJournalFilesRule{
                operations.InputJournalFilesRule{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("instead less till ew"),
                },
            },
            CurrentBoot: criblcontrolplanesdkgo.Pointer(true),
            MaxAgeDur: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("small phew before settler fortunately meh"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesKafka" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestKafka(
        operations.InputKafka{
            ID: "kafka-source",
            Type: operations.CreateInputTypeKafkaKafka,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
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
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://ugly-availability.info"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](8658.84),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](576.51),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](7446.19),
                Auth: &components.AuthTypeKafkaSchemaRegistry{
                    Disabled: false,
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                    Disabled: criblcontrolplanesdkgo.Pointer(true),
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
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](9532.6),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](4317.54),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](8891.09),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](5534.72),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](6445.75),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](575.57),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](6409.07),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](3829.31),
            Sasl: &components.AuthenticationType{
                Disabled: true,
                Username: criblcontrolplanesdkgo.Pointer("Etha51"),
                Password: criblcontrolplanesdkgo.Pointer("XukwT9HxMYXOSPg"),
                AuthType: components.AuthenticationMethodOptionsSaslSecret.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslKerberos.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(false),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://swift-laughter.com"),
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
                Disabled: criblcontrolplanesdkgo.Pointer(true),
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
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](7980.65),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](3256.62),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](767.07),
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](1358.59),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](7136.81),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](1942.41),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](9602.7),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](9414.81),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("too rebel convection low rebuild how commonly claw of"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesKinesis" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestKinesis(
        operations.InputKinesis{
            ID: "kinesis-source",
            Type: operations.CreateInputTypeKinesisKinesis,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            StreamName: "my-stream",
            ServiceInterval: criblcontrolplanesdkgo.Pointer[float64](9559.98),
            ShardExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            ShardIteratorType: operations.ShardIteratorStartTrimHorizon.ToPointer(),
            PayloadFormat: operations.RecordDataFormatCribl.ToPointer(),
            GetRecordsLimit: criblcontrolplanesdkgo.Pointer[float64](9713.56),
            GetRecordsLimitTotal: criblcontrolplanesdkgo.Pointer[float64](4313.27),
            LoadBalancingAlgorithm: operations.ShardLoadBalancingRoundRobin.ToPointer(),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfManual.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions2V4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](1070.13),
            VerifyKPLCheckSums: criblcontrolplanesdkgo.Pointer(false),
            AvoidDuplicates: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("an eke vainly atop democratize per arrogantly gosh provision till"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesKubeEvents" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestKubeEvents(
        operations.InputKubeEvents{
            ID: "kube-events-source",
            Type: operations.TypeKubeEventsKubeEvents,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Rules: []components.ItemsTypeRules{
                components.ItemsTypeRules{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("boo a yearly why genuine extroverted merry"),
                },
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("barring sweetly whether wisely"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesKubeLogs" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestKubeLogs(
        operations.InputKubeLogs{
            ID: "kube-logs-source",
            Type: operations.TypeKubeLogsKubeLogs,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](9233.6),
            Rules: []operations.RuleKubeLogs{
                operations.RuleKubeLogs{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("during microchip wallaby gracious beside till"),
                },
            },
            Timestamps: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
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
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](9154.18),
            EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("but uselessly whereas yum vet that while cautiously"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesKubeMetrics" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestKubeMetrics(
        operations.InputKubeMetrics{
            ID: "kube-metrics-source",
            Type: operations.TypeKubeMetricsKubeMetrics,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](2715.98),
            Rules: []components.ItemsTypeRules{
                components.ItemsTypeRules{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("boo a yearly why genuine extroverted merry"),
                },
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Persistence: &operations.PersistenceKubeMetrics{
                Enable: criblcontrolplanesdkgo.Pointer(true),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.DataCompressionFormatOptionsPersistenceNone.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            Description: criblcontrolplanesdkgo.Pointer("raw dream but to"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesLoki" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestLoki(
        operations.InputLoki{
            ID: "loki-source",
            Type: operations.CreateInputTypeLokiLoki,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10080,
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](1802.52),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](767657),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](2443.33),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](6185.88),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](3691.15),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](1924.71),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            LokiAPI: "/loki/api/v1/push",
            AuthType: components.AuthenticationTypeOptionsLokiAuthToken.ToPointer(),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("which until yet"),
            Username: criblcontrolplanesdkgo.Pointer("Elvera33"),
            Password: criblcontrolplanesdkgo.Pointer("t1aJbdykpwYIGuq"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://glum-plugin.com"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](4899.83),
            OauthParams: []components.ItemsTypeOauthParams{
                components.ItemsTypeOauthParams{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            OauthHeaders: []components.ItemsTypeOauthHeaders{
                components.ItemsTypeOauthHeaders{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesMetrics" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestMetrics(
        operations.InputMetrics{
            ID: "metrics-source",
            Type: operations.TypeMetricsMetrics,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            UDPPort: criblcontrolplanesdkgo.Pointer[float64](8125),
            TCPPort: criblcontrolplanesdkgo.Pointer[float64](5345.44),
            MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](6534.85),
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
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
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](3564.68),
            Description: criblcontrolplanesdkgo.Pointer("supposing joyful investigate what"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesModelDrivenTelemetry" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestModelDrivenTelemetry(
        operations.InputModelDrivenTelemetry{
            ID: "mdt-source",
            Type: operations.TypeModelDrivenTelemetryModelDrivenTelemetry,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 57000,
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
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](7010.9),
            ShutdownTimeoutMs: criblcontrolplanesdkgo.Pointer[float64](8181.4),
            Description: criblcontrolplanesdkgo.Pointer("equally however advocate deduce thoughtfully muted categorise inasmuch suddenly"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesMsk" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestMsk(
        operations.InputMsk{
            ID: "msk-source",
            Type: operations.CreateInputTypeMskMsk,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
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
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](4833.32),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](2186.74),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](9509.09),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType{
                Disabled: false,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://ugly-availability.info"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](8658.84),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](576.51),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](7446.19),
                Auth: &components.AuthTypeKafkaSchemaRegistry{
                    Disabled: false,
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                    Disabled: criblcontrolplanesdkgo.Pointer(true),
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
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](5413.46),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](6448.74),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](8060.95),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](9433.89),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5411.14),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](4695.65),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](1757.27),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](2863.12),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto,
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsV2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](5932.72),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
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
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](7946.06),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](6583.38),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](9381.06),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](734),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](2148.16),
            Description: criblcontrolplanesdkgo.Pointer("curry softly ack reconstitute failing against yet phooey via"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesNetflow" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestNetflow(
        operations.InputNetflow{
            ID: "netflow-source",
            Type: operations.CreateInputTypeNetflowNetflow,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 2055,
            EnablePassThrough: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](3658.46),
            TemplateCacheMinutes: criblcontrolplanesdkgo.Pointer[float64](1855.48),
            V5Enabled: criblcontrolplanesdkgo.Pointer(true),
            V9Enabled: criblcontrolplanesdkgo.Pointer(false),
            IpfixEnabled: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("muscat coil greatly badly outrank represent around"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesOffice365Mgmt" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestOffice365Mgmt(
        operations.InputOffice365Mgmt{
            ID: "office365-mgmt-source",
            Type: operations.TypeOffice365MgmtOffice365Mgmt,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            PlanType: components.SubscriptionPlanOptionsEnterpriseGcc,
            TenantID: "tenant-id",
            AppID: "app-id",
            Timeout: criblcontrolplanesdkgo.Pointer[float64](3487.37),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](3940.22),
            JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](8407.86),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            PublisherIdentifier: criblcontrolplanesdkgo.Pointer("<value>"),
            ContentConfig: []operations.ContentConfigOffice365Mgmt{
                operations.ContentConfigOffice365Mgmt{
                    ContentType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Description: criblcontrolplanesdkgo.Pointer("outflank clean classic how capsize inasmuch adolescent microblog pish"),
                    Interval: criblcontrolplanesdkgo.Pointer[float64](9559.88),
                    LogLevel: components.LogLevelOptionsContentConfigItemsWarn.ToPointer(),
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                },
            },
            IngestionLag: criblcontrolplanesdkgo.Pointer[float64](7747.69),
            RetryRules: &components.RetryRulesType1{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesBackoff,
                Interval: criblcontrolplanesdkgo.Pointer[float64](9780.72),
                Limit: criblcontrolplanesdkgo.Pointer[float64](2123.94),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](107.05),
                Codes: []float64{
                    4202.47,
                    9242.01,
                    2793.86,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            AuthType: components.AuthenticationMethodOptions1Secret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("whoever whose gee blah huzzah"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesOffice365MsgTrace

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesOffice365MsgTrace" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestOffice365MsgTrace(
        operations.InputOffice365MsgTrace{
            ID: "office365-msg-trace-source",
            Type: operations.TypeOffice365MsgTraceOffice365MsgTrace,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            URL: "https://reports.office365.com/ecp/reportingwebservice/reporting.svc/MessageTrace",
            Interval: 15,
            StartDate: criblcontrolplanesdkgo.Pointer("<value>"),
            EndDate: criblcontrolplanesdkgo.Pointer("<value>"),
            Timeout: criblcontrolplanesdkgo.Pointer[float64](7116.75),
            DisableTimeFilter: criblcontrolplanesdkgo.Pointer(true),
            AuthType: operations.AuthenticationMethodOffice365MsgTraceOauth.ToPointer(),
            RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
            MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](699.11),
            LogLevel: operations.LogLevelOffice365MsgTraceWarn.ToPointer(),
            JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](7208.8),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](557.68),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            RetryRules: &components.RetryRulesType1{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesBackoff,
                Interval: criblcontrolplanesdkgo.Pointer[float64](9780.72),
                Limit: criblcontrolplanesdkgo.Pointer[float64](2123.94),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](107.05),
                Codes: []float64{
                    4202.47,
                    9242.01,
                    2793.86,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            Description: criblcontrolplanesdkgo.Pointer("deflate woefully or robust over impact"),
            Username: criblcontrolplanesdkgo.Pointer("Shaun.Bruen-Weimann"),
            Password: criblcontrolplanesdkgo.Pointer("rhOhmiS7N3YIiLF"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            ClientID: criblcontrolplanesdkgo.Pointer("<id>"),
            Resource: criblcontrolplanesdkgo.Pointer("<value>"),
            PlanType: components.SubscriptionPlanOptionsGccHigh.ToPointer(),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            CertOptions: &operations.CertOptions{
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: "<value>",
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: "<value>",
            },
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesOffice365Service" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestOffice365Service(
        operations.InputOffice365Service{
            ID: "office365-service-source",
            Type: operations.TypeOffice365ServiceOffice365Service,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            PlanType: components.SubscriptionPlanOptionsDod.ToPointer(),
            TenantID: "tenant-id",
            AppID: "app-id",
            Timeout: criblcontrolplanesdkgo.Pointer[float64](6785.41),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](9440.55),
            JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](2351.7),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ContentConfig: []operations.ContentConfigOffice365Service{
                operations.ContentConfigOffice365Service{
                    ContentType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Description: criblcontrolplanesdkgo.Pointer("heavily sharply brook misjudge legitimize"),
                    Interval: criblcontrolplanesdkgo.Pointer[float64](1247.02),
                    LogLevel: components.LogLevelOptionsContentConfigItemsWarn.ToPointer(),
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                },
            },
            RetryRules: &components.RetryRulesType1{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesBackoff,
                Interval: criblcontrolplanesdkgo.Pointer[float64](9780.72),
                Limit: criblcontrolplanesdkgo.Pointer[float64](2123.94),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](107.05),
                Codes: []float64{
                    4202.47,
                    9242.01,
                    2793.86,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            AuthType: components.AuthenticationMethodOptions1Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("above gallivant finally"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesOpenTelemetry

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesOpenTelemetry" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestOpenTelemetry(
        operations.InputOpenTelemetry{
            ID: "otel-source",
            Type: operations.CreateInputTypeOpenTelemetryOpenTelemetry,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 4317,
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](4380.05),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](446963),
            EnableProxyHeader: "<value>",
            CaptureHeaders: "<value>",
            ActivityLogSampleRate: "<value>",
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7040.74),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](3566.12),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](58.19),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            Protocol: operations.CreateInputProtocolHTTP.ToPointer(),
            ExtractSpans: criblcontrolplanesdkgo.Pointer(true),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(false),
            OtlpVersion: operations.CreateInputOTLPVersionZeroDot10Dot0.ToPointer(),
            AuthType: components.AuthenticationTypeOptionsBasic.ToPointer(),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](4758.14),
            Description: criblcontrolplanesdkgo.Pointer("furthermore while shampoo"),
            Username: criblcontrolplanesdkgo.Pointer("Demetrius_Bernhard28"),
            Password: criblcontrolplanesdkgo.Pointer("_HgrUnh66BPfzNt"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://alienated-cork.org/"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](9512.38),
            OauthParams: []components.ItemsTypeOauthParams{
                components.ItemsTypeOauthParams{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            OauthHeaders: []components.ItemsTypeOauthHeaders{
                components.ItemsTypeOauthHeaders{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ExtractLogs: criblcontrolplanesdkgo.Pointer(false),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesPrometheus" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestPrometheus(
        operations.InputPrometheus{
            ID: "prometheus-source",
            Type: operations.CreateInputTypePrometheusPrometheus,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            DimensionList: []string{
                "<value 1>",
                "<value 2>",
            },
            DiscoveryType: operations.DiscoveryTypePrometheusStatic.ToPointer(),
            Interval: 60,
            LogLevel: operations.LogLevelPrometheusInfo,
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            Timeout: criblcontrolplanesdkgo.Pointer[float64](3147.32),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](305.84),
            JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](2520.89),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthType: components.AuthenticationMethodOptionsSaslSecret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("barring anxiously break lavish whoa while really until"),
            TargetList: []string{
                "http://localhost:9090/metrics",
            },
            RecordType: components.RecordTypeOptionsAaaa.ToPointer(),
            ScrapePort: criblcontrolplanesdkgo.Pointer[float64](4700.87),
            NameList: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ScrapeProtocol: operations.MetricsProtocolHTTPS.ToPointer(),
            ScrapePath: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfManual.ToPointer(),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            UsePublicIP: criblcontrolplanesdkgo.Pointer(false),
            SearchFilter: []components.ItemsTypeSearchFilter{
                components.ItemsTypeSearchFilter{
                    Name: "<value>",
                    Values: []string{
                        "<value 1>",
                        "<value 2>",
                    },
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
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](9002.52),
            Username: criblcontrolplanesdkgo.Pointer("Luciano_Fisher99"),
            Password: criblcontrolplanesdkgo.Pointer("XWKuXRnFB74TRlD"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesPrometheusRw" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestPrometheusRw(
        operations.InputPrometheusRw{
            ID: "prometheus-rw-source",
            Type: operations.TypePrometheusRwPrometheusRw,
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
            Host: "0.0.0.0",
            Port: 10080,
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](9688.99),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](95126),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](4199.2),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](4007.06),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](2378.84),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](3340.99),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            PrometheusAPI: "/write",
            AuthType: components.AuthenticationTypeOptionsPrometheusAuthNone.ToPointer(),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("ugh frank wire circa"),
            Username: criblcontrolplanesdkgo.Pointer("Keon93"),
            Password: criblcontrolplanesdkgo.Pointer("Ph2diZRCGCwQh2B"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://cruel-bungalow.biz"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](7358.94),
            OauthParams: []components.ItemsTypeOauthParams{
                components.ItemsTypeOauthParams{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            OauthHeaders: []components.ItemsTypeOauthHeaders{
                components.ItemsTypeOauthHeaders{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesRawUdp" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestRawUDP(
        operations.InputRawUDP{
            ID: "raw-udp-source",
            Type: operations.TypeRawUDPRawUDP,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 514,
            MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2988.09),
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            SingleMsgUDPPackets: criblcontrolplanesdkgo.Pointer(true),
            IngestRawBytes: criblcontrolplanesdkgo.Pointer(true),
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](3299.36),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("following cow shrilly lavish norm yahoo torn parsnip"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesS3" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestS3(
        operations.InputS3{
            ID: "s3-source",
            Type: operations.CreateInputTypeS3S3,
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
            QueueName: "s3-notifications-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfManual.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](7117.24),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](7261.65),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](7076.16),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](3607.82),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](434.52),
            SkipOnError: criblcontrolplanesdkgo.Pointer(true),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](8815.56),
            EnableSQSAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                Disabled: false,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                },
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](7027.44),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](8114.55),
            Checkpointing: &components.CheckpointingType{
                Enabled: false,
                Retries: criblcontrolplanesdkgo.Pointer[float64](7918.57),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](4127.07),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            TagAfterProcessing: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("gah lest whoever wisely"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ProcessedTagKey: criblcontrolplanesdkgo.Pointer("<value>"),
            ProcessedTagValue: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesS3Inventory" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestS3Inventory(
        operations.InputS3Inventory{
            ID: "s3-inventory-source",
            Type: operations.TypeS3InventoryS3Inventory,
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
            QueueName: "s3-inventory-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfSecret.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](2887.68),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](437.47),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](2102.66),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](6213.22),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](6368.62),
            SkipOnError: criblcontrolplanesdkgo.Pointer(true),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](2187.22),
            EnableSQSAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                Disabled: false,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                },
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](8498.54),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](2090.07),
            Checkpointing: &components.CheckpointingType{
                Enabled: false,
                Retries: criblcontrolplanesdkgo.Pointer[float64](7918.57),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](9820.73),
            ChecksumSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxManifestSizeKB: criblcontrolplanesdkgo.Pointer[int64](713420),
            ValidateInventoryFiles: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("woot what greedily blah until content along corner"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TagAfterProcessing: components.TagAfterProcessingOptionsTrue.ToPointer(),
            ProcessedTagKey: criblcontrolplanesdkgo.Pointer("<value>"),
            ProcessedTagValue: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesSecurityLake" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestSecurityLake(
        operations.InputSecurityLake{
            ID: "security-lake-source",
            Type: operations.CreateInputTypeSecurityLakeSecurityLake,
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
            QueueName: "security-lake-queue",
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](7998.97),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](9902.64),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](6703.9),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](8995.16),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](3363.78),
            SkipOnError: criblcontrolplanesdkgo.Pointer(false),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](6949.93),
            EnableSQSAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                Disabled: false,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                },
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](1512.5),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](4251.17),
            Checkpointing: &components.CheckpointingType{
                Enabled: false,
                Retries: criblcontrolplanesdkgo.Pointer[float64](7918.57),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](5816.86),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("round since extra-large however tidy"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TagAfterProcessing: components.TagAfterProcessingOptionsTrue.ToPointer(),
            ProcessedTagKey: criblcontrolplanesdkgo.Pointer("<value>"),
            ProcessedTagValue: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesSnmp" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestSnmp(
        operations.InputSnmp{
            ID: "snmp-source",
            Type: operations.CreateInputTypeSnmpSnmp,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "192.168.1.1",
            Port: 161,
            SnmpV3Auth: &operations.SNMPv3Authentication{
                V3AuthEnabled: true,
                AllowUnmatchedTrap: criblcontrolplanesdkgo.Pointer(true),
                V3Users: []operations.V3User{
                    operations.V3User{
                        Name: "<value>",
                        AuthProtocol: components.AuthenticationProtocolOptionsV3UserSha384.ToPointer(),
                        AuthKey: criblcontrolplanesdkgo.Pointer("<value>"),
                        PrivProtocol: operations.PrivacyProtocolNone.ToPointer(),
                        PrivKey: criblcontrolplanesdkgo.Pointer("<value>"),
                    },
                },
            },
            MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](6992.98),
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](6052.66),
            VarbindsWithTypes: criblcontrolplanesdkgo.Pointer(true),
            BestEffortParsing: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("yuck enthusiastically idolized now slowly disconnection crystallize than without besides"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesSplunk" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestSplunk(
        operations.InputSplunk{
            ID: "splunk-source",
            Type: operations.CreateInputTypeSplunkSplunk,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 9997,
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
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](8429.15),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](5466.61),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](1428.04),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](478.3),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](4231.52),
            AuthTokens: []operations.AuthTokenSplunk{
                operations.AuthTokenSplunk{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("descendant absent although mortally"),
                },
            },
            MaxS2Sversion: operations.MaxS2SVersionV3.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("yahoo only formamide qua midst culminate creak lest"),
            UseFwdTimezone: criblcontrolplanesdkgo.Pointer(true),
            DropControlFields: criblcontrolplanesdkgo.Pointer(false),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(false),
            Compress: operations.CreateInputCompressionAlways.ToPointer(),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesSplunkHec" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestSplunkHec(
        operations.InputSplunkHec{
            ID: "splunk-hec-source",
            Type: operations.CreateInputTypeSplunkHecSplunkHec,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 8088,
            AuthTokens: []operations.AuthTokenSplunkHec{
                operations.AuthTokenSplunkHec{
                    AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
                    TokenSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Token: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                    Description: criblcontrolplanesdkgo.Pointer("ack swiftly unscramble parody psst during stylish gum"),
                    AllowedIndexesAtToken: []string{
                        "<value 1>",
                    },
                    Metadata: []components.ItemsTypeNotificationMetadata{
                        components.ItemsTypeNotificationMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](4327.63),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](198812),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](5619.85),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](2281.35),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](3620.99),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](7989.57),
            EnableHealthCheck: "<value>",
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            SplunkHecAPI: "/services/collector",
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AllowedIndexes: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            SplunkHecAcks: criblcontrolplanesdkgo.Pointer(false),
            BreakerRulesets: []string{
                "<value 1>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](1561.91),
            UseFwdTimezone: criblcontrolplanesdkgo.Pointer(false),
            DropControlFields: criblcontrolplanesdkgo.Pointer(true),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(true),
            AccessControlAllowOrigin: []string{
                "<value 1>",
            },
            AccessControlAllowHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            EmitTokenMetrics: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("prioritize polarisation bourgeoisie decongestant behind degenerate turret exotic murky"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesSplunkSearch" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestSplunkSearch(
        operations.InputSplunkSearch{
            ID: "splunk-search-source",
            Type: operations.TypeSplunkSearchSplunkSearch,
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
            SearchHead: "https://localhost:8089",
            Search: "index=main",
            Earliest: criblcontrolplanesdkgo.Pointer("<value>"),
            Latest: criblcontrolplanesdkgo.Pointer("<value>"),
            CronSchedule: "0 * * * *",
            Endpoint: "/services/search/jobs/export",
            OutputMode: components.OutputModeOptionsSplunkCollectorConfJSON,
            EndpointParams: []operations.EndpointParam{
                operations.EndpointParam{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            EndpointHeaders: []operations.EndpointHeader{
                operations.EndpointHeader{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            LogLevel: operations.LogLevelSplunkSearchWarn.ToPointer(),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](1700.92),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](9155.6),
            JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](6019.15),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            RetryRules: &components.RetryRulesType{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesNone,
                Interval: criblcontrolplanesdkgo.Pointer[float64](7746.55),
                Limit: criblcontrolplanesdkgo.Pointer[float64](2960.85),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](5912.74),
                Codes: []float64{
                    4045.69,
                    7782.74,
                    6687.35,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            BreakerRulesets: []string{
                "<value 1>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](9405.7),
            AuthType: operations.AuthenticationTypeSplunkSearchBasic.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("plugin roughly aboard hourly descendant pfft whoa"),
            Username: criblcontrolplanesdkgo.Pointer("Terrance_Wolff-Roberts"),
            Password: criblcontrolplanesdkgo.Pointer("DYJItRIXcRsFuTU"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://unaware-bar.biz"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](3722.95),
            OauthParams: []components.ItemsTypeOauthParams{
                components.ItemsTypeOauthParams{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            OauthHeaders: []components.ItemsTypeOauthHeaders{
                components.ItemsTypeOauthHeaders{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesSqs" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestSqs(
        operations.InputSqs{
            ID: "sqs-source",
            Type: operations.CreateInputTypeSqsSqs,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "my-queue",
            QueueType: operations.CreateInputQueueTypeStandard,
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            CreateQueue: criblcontrolplanesdkgo.Pointer(true),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions3V2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](7049.71),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](2368.62),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](2367.72),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](1405.07),
            Description: criblcontrolplanesdkgo.Pointer("instead afterwards who consequently"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](6199.83),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesSyslog" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestSyslog(
        operations.CreateInputSyslogInputSyslogSyslog1(
            operations.InputSyslogSyslog1{
                ID: "syslog-source",
                Type: operations.InputSyslogType1Syslog,
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
                Host: "0.0.0.0",
                UDPPort: 514,
                TCPPort: criblcontrolplanesdkgo.Pointer[float64](2282.1),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3011.12),
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
                MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](3194.4),
                SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](1333.98),
                SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](8139.96),
                SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](9377.11),
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
                Metadata: []components.ItemsTypeNotificationMetadata{
                    components.ItemsTypeNotificationMetadata{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](7376.65),
                EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(true),
                Description: criblcontrolplanesdkgo.Pointer("wavy however incidentally ownership valiantly climb tame"),
                EnableEnhancedProxyHeaderParsing: criblcontrolplanesdkgo.Pointer(false),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesSystemMetrics" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestSystemMetrics(
        operations.InputSystemMetrics{
            ID: "system-metrics-source",
            Type: operations.TypeSystemMetricsSystemMetrics,
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
            Interval: criblcontrolplanesdkgo.Pointer[float64](688.04),
            Host: &operations.HostSystemMetrics{
                Mode: components.ModeOptionsHostCustom.ToPointer(),
                Custom: &operations.CustomSystemMetrics{
                    System: &operations.SystemSystemMetrics{
                        Mode: operations.SystemModeSystemMetricsCustom.ToPointer(),
                        Processes: criblcontrolplanesdkgo.Pointer(true),
                    },
                    CPU: &operations.CPUSystemMetrics{
                        Mode: operations.CPUModeSystemMetricsBasic.ToPointer(),
                        PerCPU: criblcontrolplanesdkgo.Pointer(false),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Time: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Memory: &operations.MemorySystemMetrics{
                        Mode: operations.MemoryModeSystemMetricsCustom.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Network: &operations.NetworkSystemMetrics{
                        Mode: operations.NetworkModeSystemMetricsCustom.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Protocols: criblcontrolplanesdkgo.Pointer(true),
                        Devices: []string{
                            "<value 1>",
                            "<value 2>",
                            "<value 3>",
                        },
                        PerInterface: criblcontrolplanesdkgo.Pointer(true),
                    },
                    Disk: &operations.DiskSystemMetrics{
                        Mode: operations.DiskModeSystemMetricsBasic.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Inodes: criblcontrolplanesdkgo.Pointer(true),
                        Devices: []string{
                            "<value 1>",
                            "<value 2>",
                        },
                        Mountpoints: []string{
                            "<value 1>",
                        },
                        Fstypes: []string{
                            "<value 1>",
                            "<value 2>",
                            "<value 3>",
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
                        IncludeChildren: criblcontrolplanesdkgo.Pointer(false),
                    },
                },
            },
            Container: &operations.Container{
                Mode: operations.ContainerModeDisabled.ToPointer(),
                DockerSocket: []string{
                    "<value 1>",
                },
                DockerTimeout: criblcontrolplanesdkgo.Pointer[float64](9193.24),
                Filters: []operations.ContainerFilter{
                    operations.ContainerFilter{
                        Expr: "<value>",
                    },
                },
                AllContainers: criblcontrolplanesdkgo.Pointer(false),
                PerDevice: criblcontrolplanesdkgo.Pointer(false),
                Detail: criblcontrolplanesdkgo.Pointer(true),
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Persistence: &operations.PersistenceSystemMetrics{
                Enable: criblcontrolplanesdkgo.Pointer(true),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.DataCompressionFormatOptionsPersistenceNone.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            Description: criblcontrolplanesdkgo.Pointer("into hmph bah clean which status cannon lively"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesSystemState" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestSystemState(
        operations.InputSystemState{
            ID: "system-state-source",
            Type: operations.TypeSystemStateSystemState,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](4335.6),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Collectors: &operations.Collectors{
                Hostsfile: &operations.HostsFile{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Interfaces: &operations.Interfaces{
                    Enable: criblcontrolplanesdkgo.Pointer(false),
                },
                Disk: &operations.DisksAndFileSystems{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Metadata: &operations.HostInfo{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Routes: &operations.Routes{
                    Enable: criblcontrolplanesdkgo.Pointer(false),
                },
                DNS: &operations.DNS{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                User: &operations.UsersAndGroups{
                    Enable: criblcontrolplanesdkgo.Pointer(false),
                },
                Firewall: &operations.Firewall{
                    Enable: criblcontrolplanesdkgo.Pointer(false),
                },
                Services: &operations.Services{
                    Enable: criblcontrolplanesdkgo.Pointer(false),
                },
                Ports: &operations.ListeningPorts{
                    Enable: criblcontrolplanesdkgo.Pointer(false),
                },
                LoginUsers: &operations.LoggedInUsers{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
            },
            Persistence: &operations.PersistenceSystemState{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.DataCompressionFormatOptionsPersistenceNone.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            DisableNativeModule: criblcontrolplanesdkgo.Pointer(false),
            DisableNativeLastLogModule: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("plus merit the"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesTcp" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestTCP(
        operations.InputTCP{
            ID: "tcp-source",
            Type: operations.TypeTCPTCP,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10090,
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
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](2458.75),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](6864.89),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](9734.82),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](880.4),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](5128.48),
            EnableHeader: criblcontrolplanesdkgo.Pointer(false),
            Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                Disabled: false,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("ape snoopy equally common er provider concerning unimpressively pish"),
            AuthToken: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
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
### Example Usage: InputCreateExamplesTcpjson

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesTcpjson" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestTcpjson(
        operations.InputTcpjson{
            ID: "tcpjson-source",
            Type: operations.CreateInputTypeTcpjsonTcpjson,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10090,
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
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](6182.84),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](5228.06),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](1333.88),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](3762.69),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(true),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("ouch under even who dissemble gratefully slushy deceivingly"),
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
### Example Usage: InputCreateExamplesWef

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesWef" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestWef(
        operations.InputWef{
            ID: "wef-source",
            Type: operations.TypeWefWef,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 5985,
            AuthMethod: operations.AuthMethodAuthenticationMethodKerberos.ToPointer(),
            TLS: &operations.MTLSSettings{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: "<value>",
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: "<value>",
                CaPath: "<value>",
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
                OcspCheck: criblcontrolplanesdkgo.Pointer(false),
                Keytab: "<value>",
                Principal: "<value>",
                OcspCheckFailClose: criblcontrolplanesdkgo.Pointer(true),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](9267.54),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](505940),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](1430.74),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](5061.11),
            CaFingerprint: criblcontrolplanesdkgo.Pointer("<value>"),
            Keytab: criblcontrolplanesdkgo.Pointer("<value>"),
            Principal: criblcontrolplanesdkgo.Pointer("<value>"),
            AllowMachineIDMismatch: criblcontrolplanesdkgo.Pointer(true),
            Subscriptions: []operations.Subscription{
                operations.Subscription{
                    SubscriptionName: "subscription-1",
                    Version: criblcontrolplanesdkgo.Pointer("<value>"),
                    ContentFormat: operations.CreateInputFormatRenderedText,
                    HeartbeatInterval: 60,
                    BatchTimeout: 5,
                    ReadExistingEvents: criblcontrolplanesdkgo.Pointer(true),
                    SendBookmarks: criblcontrolplanesdkgo.Pointer(true),
                    Compress: criblcontrolplanesdkgo.Pointer(true),
                    Targets: []string{},
                    Locale: criblcontrolplanesdkgo.Pointer("de"),
                    QuerySelector: operations.QueryBuilderModeSimple.ToPointer(),
                    Metadata: []components.ItemsTypeNotificationMetadata{
                        components.ItemsTypeNotificationMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                    Queries: []operations.Query{
                        operations.Query{
                            Path: "/home/user/dir",
                            QueryExpression: "<value>",
                        },
                    },
                    XMLQuery: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("mockingly tenderly impressive scarification border ring"),
            LogFingerprintMismatch: criblcontrolplanesdkgo.Pointer(false),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesWinEventLogs" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestWinEventLogs(
        operations.InputWinEventLogs{
            ID: "win-event-logs-source",
            Type: operations.TypeWinEventLogsWinEventLogs,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            LogNames: []string{
                "Application",
                "System",
            },
            ReadMode: operations.ReadModeNewest.ToPointer(),
            EventFormat: operations.EventFormatXML.ToPointer(),
            DisableNativeModule: criblcontrolplanesdkgo.Pointer(true),
            Interval: criblcontrolplanesdkgo.Pointer[float64](852.74),
            BatchSize: criblcontrolplanesdkgo.Pointer[float64](5653.41),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            MaxEventBytes: criblcontrolplanesdkgo.Pointer[float64](4593.26),
            Description: criblcontrolplanesdkgo.Pointer("citizen premeditation dismal"),
            DisableJSONRendering: criblcontrolplanesdkgo.Pointer(false),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesWindowsMetrics" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestWindowsMetrics(
        operations.InputWindowsMetrics{
            ID: "windows-metrics-source",
            Type: operations.TypeWindowsMetricsWindowsMetrics,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](5250.98),
            Host: &operations.HostWindowsMetrics{
                Mode: components.ModeOptionsHostCustom.ToPointer(),
                Custom: &operations.CustomWindowsMetrics{
                    System: &operations.SystemWindowsMetrics{
                        Mode: operations.SystemModeWindowsMetricsBasic.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(true),
                    },
                    CPU: &operations.CPUWindowsMetrics{
                        Mode: operations.CPUModeWindowsMetricsCustom.ToPointer(),
                        PerCPU: criblcontrolplanesdkgo.Pointer(true),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Time: criblcontrolplanesdkgo.Pointer(true),
                    },
                    Memory: &operations.MemoryWindowsMetrics{
                        Mode: operations.MemoryModeWindowsMetricsBasic.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(true),
                    },
                    Network: &operations.NetworkWindowsMetrics{
                        Mode: operations.NetworkModeWindowsMetricsDisabled.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Protocols: criblcontrolplanesdkgo.Pointer(false),
                        Devices: []string{
                            "<value 1>",
                        },
                        PerInterface: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Disk: &operations.DiskWindowsMetrics{
                        Mode: operations.DiskModeWindowsMetricsBasic.ToPointer(),
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
                        IncludeChildren: criblcontrolplanesdkgo.Pointer(false),
                    },
                },
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Persistence: &operations.PersistenceWindowsMetrics{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.DataCompressionFormatOptionsPersistenceNone.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            DisableNativeModule: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("soon out when towards sometimes"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesWiz" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestWiz(
        operations.InputWiz{
            ID: "wiz-source",
            Type: operations.TypeWizWiz,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Endpoint: "https://api.wiz.io",
            AuthURL: "https://auth.wiz.io/oauth/token",
            AuthAudienceOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            ClientID: "client-id",
            ContentConfig: []operations.ContentConfigWiz{},
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](1355.21),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](2315.99),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](5797.18),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            RetryRules: &components.RetryRulesType{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesNone,
                Interval: criblcontrolplanesdkgo.Pointer[float64](7746.55),
                Limit: criblcontrolplanesdkgo.Pointer[float64](2960.85),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](5912.74),
                Codes: []float64{
                    4045.69,
                    7782.74,
                    6687.35,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            AuthType: components.AuthenticationMethodOptions1Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("gad wisecrack that meaningfully wherever"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesWizWebhook

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesWizWebhook" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestWizWebhook(
        operations.InputWizWebhook{
            ID: "wiz-webhook-source",
            Type: operations.TypeWizWebhookWizWebhook,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](2391.82),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](444294),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](7567.99),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](3218.14),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](2243.35),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](989.04),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](8737.45),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AllowedPaths: []string{
                "<value 1>",
            },
            AllowedMethods: []string{
                "<value 1>",
            },
            AuthTokensExt: []components.ItemsTypeAuthTokensExt{
                components.ItemsTypeAuthTokensExt{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("deeply very fiercely at total truthfully how frugal ah"),
                    Metadata: []components.ItemsTypeNotificationMetadata{
                        components.ItemsTypeNotificationMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("veg anenst underpants degenerate"),
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

<!-- UsageSnippet language="go" operationID="createInput" method="post" path="/system/inputs" example="InputCreateExamplesZscalerHec" -->
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

    res, err := s.Sources.Create(ctx, operations.CreateCreateInputRequestZscalerHec(
        operations.InputZscalerHec{
            ID: "zscaler-hec-source",
            Type: operations.TypeZscalerHecZscalerHec,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.73),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](7905.42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/opt/bin"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 8088,
            AuthTokens: []operations.AuthTokenZscalerHec{
                operations.AuthTokenZscalerHec{
                    AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
                    TokenSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Token: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                    Description: criblcontrolplanesdkgo.Pointer("about garrote complication gee woot"),
                    AllowedIndexesAtToken: []string{
                        "<value 1>",
                        "<value 2>",
                    },
                    Metadata: []components.ItemsTypeNotificationMetadata{
                        components.ItemsTypeNotificationMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](4348.05),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](172464),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](9777.07),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](3114.25),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](5229.35),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](4629.12),
            EnableHealthCheck: "<value>",
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            HecAPI: "/services/collector",
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
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
                "<value 2>",
                "<value 3>",
            },
            EmitTokenMetrics: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("self-reliant nautical fold"),
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

### Example Usage: InputCreateExamplesAppscope

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesAppscope" -->
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
### Example Usage: InputCreateExamplesAzureBlob

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesAzureBlob" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputAzureBlob(
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "azure-blob-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("<value>"),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](5294.27),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](3992.62),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](5686.02),
            ServicePeriodSecs: criblcontrolplanesdkgo.Pointer[float64](8039.05),
            SkipOnError: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](9969.24),
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](1561.4),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](1595.41),
            AuthType: components.AuthenticationMethodOptionsSecret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("understated extroverted intensely hello gentle ouch hmph rule density design"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesCloudflareHec" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputCloudflareHec(
        components.InputCloudflareHec{
            ID: criblcontrolplanesdkgo.Pointer("cloudflare-hec-source"),
            Type: components.InputCloudflareHecTypeCloudflareHec,
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
            Host: "0.0.0.0",
            Port: 8088,
            AuthTokens: []components.InputCloudflareHecAuthToken{
                components.InputCloudflareHecAuthToken{
                    AuthType: components.InputCloudflareHecAuthenticationMethodSecret.ToPointer(),
                    TokenSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("viciously left knuckle finally likewise before"),
                    AllowedIndexesAtToken: []string{
                        "<value 1>",
                        "<value 2>",
                        "<value 3>",
                    },
                    Metadata: []components.ItemsTypeNotificationMetadata{
                        components.ItemsTypeNotificationMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](3960.65),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](684894),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](5621.37),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](8713.82),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](8592.43),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](2174.26),
            EnableHealthCheck: "<value>",
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            HecAPI: "/services/collector",
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
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
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](8605.29),
            AccessControlAllowOrigin: []string{
                "<value 1>",
            },
            AccessControlAllowHeaders: []string{
                "<value 1>",
            },
            EmitTokenMetrics: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("discourse into phooey fledgling along or since gadzooks purse"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesCollection" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputCollection(
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](8452.47),
            Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
            },
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesConfluentCloud" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputConfluentCloud(
        components.InputConfluentCloud{
            ID: criblcontrolplanesdkgo.Pointer("confluent-cloud-source"),
            Type: components.InputConfluentCloudTypeConfluentCloud,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Brokers: []string{
                "pkc-xxxxx.us-east-1.aws.confluent.cloud:9092",
            },
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            Topics: []string{
                "logs",
            },
            GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
            FromBeginning: criblcontrolplanesdkgo.Pointer(true),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType{
                Disabled: true,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://concrete-petal.info"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](2814.06),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7605.14),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](8934.79),
                Auth: &components.AuthTypeKafkaSchemaRegistry{
                    Disabled: false,
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                    Disabled: criblcontrolplanesdkgo.Pointer(true),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                    CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                },
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](6822.45),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](1289.1),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](124.52),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](9242.82),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9313.77),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2563.57),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](4547.37),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](670.16),
            Sasl: &components.AuthenticationType{
                Disabled: true,
                Username: criblcontrolplanesdkgo.Pointer("Ruben.Prosacco"),
                Password: criblcontrolplanesdkgo.Pointer("9KZ67ITYDbrbAZJ"),
                AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslScramSha512.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(true),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://brilliant-sustenance.net"),
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
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](2651.67),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](8421.5),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](5595.6),
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](2888.32),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](4884.37),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](2952.38),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](9163.66),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](4977.5),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("furthermore pfft pace vague behold despite boulevard corrupt"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesCriblHttp" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputCriblHTTP(
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
            Host: "0.0.0.0",
            Port: 10080,
            AuthTokens: []components.ItemsTypeAuthTokens{
                components.ItemsTypeAuthTokens{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("fatal every rebuke ew forenenst millet"),
                },
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](539.26),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](239267),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](6330.35),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](6630.19),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](2059.35),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](9652.08),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("wearily innovation proliferate tomorrow appliance"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesCriblLakeHttp" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputCriblLakeHTTP(
        components.InputCriblLakeHTTP{
            ID: criblcontrolplanesdkgo.Pointer("cribl-lake-http-source"),
            Type: components.InputCriblLakeHTTPTypeCriblLakeHTTP,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](242.74),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](851306),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](8524.32),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](1684.81),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](7908.19),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](6036.75),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            CriblAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            ElasticAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            SplunkHecAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            SplunkHecAcks: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthTokensExt: []components.AuthTokensExt{
                components.AuthTokensExt{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("before whereas provided although bore furthermore only closely deduce tousle"),
                    Metadata: []components.ItemsTypeNotificationMetadata{
                        components.ItemsTypeNotificationMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                    SplunkHecMetadata: &components.SplunkHecMetadata{
                        Enabled: criblcontrolplanesdkgo.Pointer(true),
                        DefaultDataset: criblcontrolplanesdkgo.Pointer("<value>"),
                        AllowedIndexesAtToken: []string{
                            "<value 1>",
                            "<value 2>",
                            "<value 3>",
                        },
                    },
                    ElasticsearchMetadata: &components.ElasticsearchMetadata{
                        Enabled: criblcontrolplanesdkgo.Pointer(true),
                        DefaultDataset: criblcontrolplanesdkgo.Pointer("<value>"),
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("shrill for without lend awkwardly waft forgo monster service athletic"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesCriblTcp" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputCriblTCP(
        components.InputCriblTCP{
            ID: criblcontrolplanesdkgo.Pointer("cribl-tcp-source"),
            Type: components.InputCriblTCPTypeCriblTCP,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10090,
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
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](8039.05),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](9275.52),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](2495.06),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](858.03),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(false),
            AuthTokens: []components.ItemsTypeAuthTokens{
                components.ItemsTypeAuthTokens{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("fatal every rebuke ew forenenst millet"),
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("supposing past honesty gah made-up solace bouncy which comfortable"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesCrowdstrike" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputCrowdstrike(
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "crowdstrike-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfSecret.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](7932.13),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](999.63),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](554.2),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](3598.43),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](256.82),
            SkipOnError: criblcontrolplanesdkgo.Pointer(true),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3919.57),
            EnableSQSAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Checkpointing: &components.CheckpointingType{
                Enabled: false,
                Retries: criblcontrolplanesdkgo.Pointer[float64](3583.79),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](2925.13),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("alert bungalow with obligation supposing twine halt tomography basic rowdy"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TagAfterProcessing: components.TagAfterProcessingOptionsTrue.ToPointer(),
            ProcessedTagKey: criblcontrolplanesdkgo.Pointer("<value>"),
            ProcessedTagValue: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesDatadogAgent" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputDatadogAgent(
        components.InputDatadogAgent{
            ID: criblcontrolplanesdkgo.Pointer("datadog-agent-source"),
            Type: components.InputDatadogAgentTypeDatadogAgent,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 8126,
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](9127.42),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](469633),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](2292.96),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](5623.59),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](7916.97),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](2802.22),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ProxyMode: &components.InputDatadogAgentProxyMode{
                Enabled: false,
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            },
            Description: criblcontrolplanesdkgo.Pointer("tune-up sparkling assist impact"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesDatagen" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputDatagen(
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Samples: []components.Sample{
                components.Sample{
                    Sample: "sample.json",
                    EventsPerSec: 10,
                },
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("duh calmly meander unfortunately upright wise zowie inside inasmuch until"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesEdgePrometheus" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputEdgePrometheus(
        components.InputEdgePrometheus{
            ID: criblcontrolplanesdkgo.Pointer("edge-prometheus-source"),
            Type: components.InputEdgePrometheusTypeEdgePrometheus,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            DimensionList: []string{
                "<value 1>",
                "<value 2>",
            },
            DiscoveryType: components.InputEdgePrometheusDiscoveryTypeStatic,
            Interval: 60,
            Timeout: criblcontrolplanesdkgo.Pointer[float64](7086.07),
            Persistence: &components.DiskSpoolingType{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.CompressionOptionsPersistenceNone.ToPointer(),
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthType: components.InputEdgePrometheusAuthenticationMethodSecret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("though refine fully though because insist eek ick lest finally"),
            Targets: []components.Target{
                components.Target{
                    Protocol: components.ProtocolOptionsTargetsItemsHTTPS.ToPointer(),
                    Host: "localhost",
                    Port: criblcontrolplanesdkgo.Pointer[float64](6264.12),
                    Path: criblcontrolplanesdkgo.Pointer("/selinux"),
                },
            },
            RecordType: components.RecordTypeOptionsSrv.ToPointer(),
            ScrapePort: criblcontrolplanesdkgo.Pointer[float64](5626.37),
            NameList: []string{
                "<value 1>",
            },
            ScrapeProtocol: components.ProtocolOptionsTargetsItemsHTTP.ToPointer(),
            ScrapePath: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            UsePublicIP: criblcontrolplanesdkgo.Pointer(false),
            SearchFilter: []components.ItemsTypeSearchFilter{
                components.ItemsTypeSearchFilter{
                    Name: "<value>",
                    Values: []string{
                        "<value 1>",
                    },
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
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](8251.32),
            ScrapeProtocolExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            ScrapePortExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            ScrapePathExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            PodFilter: []components.PodFilter{
                components.PodFilter{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("flashy ick phony while who membership shell yesterday testimonial"),
                },
            },
            Username: criblcontrolplanesdkgo.Pointer("Carlos.Heaney67"),
            Password: criblcontrolplanesdkgo.Pointer("gtl6f_Qdi1zxAJl"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesElastic" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputElastic(
        components.InputElastic{
            ID: criblcontrolplanesdkgo.Pointer("elastic-source"),
            Type: components.InputElasticTypeElastic,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "localhost",
            Port: 9200,
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](7696.27),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](8041),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](5052.41),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7221.64),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](3356.99),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5056.05),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            ElasticAPI: "/",
            AuthType: components.InputElasticAuthenticationTypeCredentialsSecret.ToPointer(),
            APIVersion: components.InputElasticAPIVersionEightDot3Dot2.ToPointer(),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ProxyMode: &components.InputElasticProxyMode{
                Enabled: false,
                AuthType: components.InputElasticAuthenticationMethodNone.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("Constance.Kuhlman"),
                Password: criblcontrolplanesdkgo.Pointer("tfyqTaaDTgjB6pd"),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                URL: criblcontrolplanesdkgo.Pointer("https://fuzzy-mortise.info"),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                RemoveHeaders: []string{
                    "<value 1>",
                    "<value 2>",
                },
                TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](5673.09),
            },
            Description: criblcontrolplanesdkgo.Pointer("dependent cheerful hm plus qua developmental"),
            Username: criblcontrolplanesdkgo.Pointer("Nat23"),
            Password: criblcontrolplanesdkgo.Pointer("UtJZmJA2OMhDFcq"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthTokens: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            CustomAPIVersion: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesEventhub" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputEventhub(
        components.InputEventhub{
            ID: criblcontrolplanesdkgo.Pointer("eventhub-source"),
            Type: components.InputEventhubTypeEventhub,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
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
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](2907.13),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](2951.08),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](4807.76),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](2878.21),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5301.79),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7736.65),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](3975.06),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](3120.04),
            Sasl: &components.AuthenticationType1{
                Disabled: false,
                AuthType: components.AuthenticationMethodOptionsSasl1Manual.ToPointer(),
                Password: criblcontrolplanesdkgo.Pointer("lzctsKxgHZhl5Rc"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSasl1Oauthbearer.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("Edmond.Konopelski31"),
                ClientSecretAuthType: components.AuthenticationMethodOptionsSasl2Manual.ToPointer(),
                ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                ClientTextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEndpoint: components.MicrosoftEntraIDAuthenticationEndpointOptionsSaslHTTPSLoginMicrosoftonlineCom.ToPointer(),
                ClientID: criblcontrolplanesdkgo.Pointer("<id>"),
                TenantID: criblcontrolplanesdkgo.Pointer("<id>"),
                Scope: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            TLS: &components.TLSSettingsClientSideType{
                Disabled: true,
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            },
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](8555.39),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](1021.76),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](3760.74),
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](1722.38),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](6339.7),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](7169.23),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](8824.56),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](402.82),
            MinimizeDuplicates: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("furthermore frantically solvency likable after brr meh freely"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesExec" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputExec(
        components.InputExec{
            ID: criblcontrolplanesdkgo.Pointer("exec-source"),
            Type: components.InputExecTypeExec,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Command: "echo \"Hello World\"",
            Retries: criblcontrolplanesdkgo.Pointer[float64](9861.97),
            ScheduleType: components.ScheduleTypeCronSchedule.ToPointer(),
            BreakerRulesets: []string{
                "<value 1>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](7410.31),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("though seriously per"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesFile" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputFile(
        components.InputFile{
            ID: criblcontrolplanesdkgo.Pointer("file-source"),
            Type: components.InputFileTypeFile,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Mode: components.InputFileModeManual.ToPointer(),
            Interval: criblcontrolplanesdkgo.Pointer[float64](6083.7),
            Filenames: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            FilterArchivedFiles: criblcontrolplanesdkgo.Pointer(false),
            TailOnly: criblcontrolplanesdkgo.Pointer(true),
            IdleTimeout: criblcontrolplanesdkgo.Pointer[float64](2495.48),
            MinAgeDur: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxAgeDur: criblcontrolplanesdkgo.Pointer("<value>"),
            CheckFileModTime: criblcontrolplanesdkgo.Pointer(true),
            ForceText: criblcontrolplanesdkgo.Pointer(true),
            HashLen: criblcontrolplanesdkgo.Pointer[float64](6484.64),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](1432.48),
            Description: criblcontrolplanesdkgo.Pointer("ah afore however present hence what zowie"),
            Path: criblcontrolplanesdkgo.Pointer("/usr/bin"),
            Depth: criblcontrolplanesdkgo.Pointer[float64](17.47),
            SuppressMissingPathErrors: criblcontrolplanesdkgo.Pointer(false),
            DeleteFiles: criblcontrolplanesdkgo.Pointer(false),
            SaltHash: criblcontrolplanesdkgo.Pointer(true),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesFirehose" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputFirehose(
        components.InputFirehose{
            ID: criblcontrolplanesdkgo.Pointer("firehose-source"),
            Type: components.InputFirehoseTypeFirehose,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](7597.77),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](759138),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](6878.65),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](4516.36),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](2537.94),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](3392.09),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("sure-footed well-groomed monthly"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesGooglePubsub" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputGooglePubsub(
        components.InputGooglePubsub{
            ID: criblcontrolplanesdkgo.Pointer("google-pubsub-source"),
            Type: components.InputGooglePubsubTypeGooglePubsub,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            TopicName: "my-topic",
            SubscriptionName: "my-subscription",
            MonitorSubscription: criblcontrolplanesdkgo.Pointer(false),
            CreateTopic: criblcontrolplanesdkgo.Pointer(false),
            CreateSubscription: criblcontrolplanesdkgo.Pointer(true),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            GoogleAuthMethod: components.GoogleAuthenticationMethodOptionsManual.ToPointer(),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxBacklog: criblcontrolplanesdkgo.Pointer[float64](335.09),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](6007.89),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](33.62),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("yuck redraw psst from mmm till"),
            OrderedDelivery: criblcontrolplanesdkgo.Pointer(true),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesGrafana" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputGrafana(
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
                Host: "0.0.0.0",
                Port: 10080,
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
                MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](4763.99),
                MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](541461),
                EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
                CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
                ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](1768.11),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](5812.82),
                SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](8267.83),
                KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](4628.16),
                EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
                IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
                IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
                PrometheusAPI: "/api/prom/push",
                LokiAPI: criblcontrolplanesdkgo.Pointer("<value>"),
                PrometheusAuth: &components.PrometheusAuth1{
                    AuthType: components.AuthenticationTypeOptionsPrometheusAuthOauth.ToPointer(),
                    Username: criblcontrolplanesdkgo.Pointer("Laverne.Legros28"),
                    Password: criblcontrolplanesdkgo.Pointer("l9hg6sbGombuWWq"),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    LoginURL: criblcontrolplanesdkgo.Pointer("https://self-reliant-millet.org"),
                    SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
                    Secret: criblcontrolplanesdkgo.Pointer("<value>"),
                    TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
                    AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
                    TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](1339.18),
                    OauthParams: []components.ItemsTypeOauthParams{
                        components.ItemsTypeOauthParams{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                    OauthHeaders: []components.ItemsTypeOauthHeaders{
                        components.ItemsTypeOauthHeaders{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
                LokiAuth: &components.LokiAuth1{
                    AuthType: components.AuthenticationTypeOptionsLokiAuthBasic.ToPointer(),
                    Username: criblcontrolplanesdkgo.Pointer("Jordane.Hartmann3"),
                    Password: criblcontrolplanesdkgo.Pointer("_qO39ImWhkAe8Pw"),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    LoginURL: criblcontrolplanesdkgo.Pointer("https://realistic-league.name/"),
                    SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
                    Secret: criblcontrolplanesdkgo.Pointer("<value>"),
                    TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
                    AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
                    TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](1463.81),
                    OauthParams: []components.ItemsTypeOauthParams{
                        components.ItemsTypeOauthParams{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                    OauthHeaders: []components.ItemsTypeOauthHeaders{
                        components.ItemsTypeOauthHeaders{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
                Metadata: []components.ItemsTypeNotificationMetadata{
                    components.ItemsTypeNotificationMetadata{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                Description: criblcontrolplanesdkgo.Pointer("aha inject winged even incandescence"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesHttp" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputHTTP(
        components.InputHTTP{
            ID: criblcontrolplanesdkgo.Pointer("http-source"),
            Type: components.InputHTTPTypeHTTP,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](1811.74),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](83718),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](3908.2),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](9581.8),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](3433.67),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](1522.09),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            CriblAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            ElasticAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            SplunkHecAPI: criblcontrolplanesdkgo.Pointer("<value>"),
            SplunkHecAcks: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthTokensExt: []components.ItemsTypeAuthTokensExt{
                components.ItemsTypeAuthTokensExt{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("unless scuttle pixellate hm aw furthermore tensely pro how beloved"),
                    Metadata: []components.ItemsTypeNotificationMetadata{
                        components.ItemsTypeNotificationMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("furthermore than pomelo supposing underneath spring smoothly brightly sew"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesHttpRaw" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputHTTPRaw(
        components.InputHTTPRaw{
            ID: criblcontrolplanesdkgo.Pointer("http-raw-source"),
            Type: components.InputHTTPRawTypeHTTPRaw,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](4057.64),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](209241),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](9045.22),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](6925),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](6768.3),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](2802.36),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](9818.06),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
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
                "<value 3>",
            },
            AuthTokensExt: []components.ItemsTypeAuthTokensExt{
                components.ItemsTypeAuthTokensExt{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("unless scuttle pixellate hm aw furthermore tensely pro how beloved"),
                    Metadata: []components.ItemsTypeNotificationMetadata{
                        components.ItemsTypeNotificationMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("than helpfully adjourn shoulder gee beside preregister"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesJournalFiles" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputJournalFiles(
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Path: "/var/log/journal",
            Interval: criblcontrolplanesdkgo.Pointer[float64](9499.26),
            Journals: []string{
                "system",
            },
            Rules: []components.InputJournalFilesRule{
                components.InputJournalFilesRule{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("pocket-watch given fondly abaft fluctuate whoever punctuation simple"),
                },
            },
            CurrentBoot: criblcontrolplanesdkgo.Pointer(true),
            MaxAgeDur: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("electronics yahoo phew drat boo"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesKafka" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputKafka(
        components.InputKafka{
            ID: criblcontrolplanesdkgo.Pointer("kafka-source"),
            Type: components.InputKafkaTypeKafka,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
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
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://concrete-petal.info"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](2814.06),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7605.14),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](8934.79),
                Auth: &components.AuthTypeKafkaSchemaRegistry{
                    Disabled: false,
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                    Disabled: criblcontrolplanesdkgo.Pointer(true),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                    CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                },
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](3528.96),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](2367.78),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](3399.14),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](2734.39),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1990.67),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5356.6),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](9107.44),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](9458.51),
            Sasl: &components.AuthenticationType{
                Disabled: true,
                Username: criblcontrolplanesdkgo.Pointer("Ruben.Prosacco"),
                Password: criblcontrolplanesdkgo.Pointer("9KZ67ITYDbrbAZJ"),
                AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslScramSha512.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(true),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://brilliant-sustenance.net"),
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
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](4767.62),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](9493.08),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](8352.03),
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](8572.75),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](9008.21),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](196.45),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](4443.15),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](9682.16),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("gee anaesthetise through deployment oh where drat snoop excepting inquisitively"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesKinesis" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputKinesis(
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            StreamName: "my-stream",
            ServiceInterval: criblcontrolplanesdkgo.Pointer[float64](1279.68),
            ShardExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            ShardIteratorType: components.ShardIteratorStartTrimHorizon.ToPointer(),
            PayloadFormat: components.RecordDataFormatLine.ToPointer(),
            GetRecordsLimit: criblcontrolplanesdkgo.Pointer[float64](7102.12),
            GetRecordsLimitTotal: criblcontrolplanesdkgo.Pointer[float64](9693.93),
            LoadBalancingAlgorithm: components.ShardLoadBalancingConsistentHashing.ToPointer(),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions2V2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](913.04),
            VerifyKPLCheckSums: criblcontrolplanesdkgo.Pointer(false),
            AvoidDuplicates: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("whose husky slink where woefully shallow"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesKubeEvents" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputKubeEvents(
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Rules: []components.ItemsTypeRules{
                components.ItemsTypeRules{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("voluntarily on ack scornful narrowcast"),
                },
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("stunt entice failing owlishly scout waterspout swine upon clamp"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesKubeLogs" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputKubeLogs(
        components.InputKubeLogs{
            ID: criblcontrolplanesdkgo.Pointer("kube-logs-source"),
            Type: components.InputKubeLogsTypeKubeLogs,
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
            Interval: criblcontrolplanesdkgo.Pointer[float64](247.56),
            Rules: []components.InputKubeLogsRule{
                components.InputKubeLogsRule{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("persecute anenst although source overfeed joint"),
                },
            },
            Timestamps: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
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
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](8820.92),
            EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("affect so ah lotion since sheathe"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesKubeMetrics" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputKubeMetrics(
        components.InputKubeMetrics{
            ID: criblcontrolplanesdkgo.Pointer("kube-metrics-source"),
            Type: components.InputKubeMetricsTypeKubeMetrics,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](8394.01),
            Rules: []components.ItemsTypeRules{
                components.ItemsTypeRules{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("voluntarily on ack scornful narrowcast"),
                },
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Persistence: &components.InputKubeMetricsPersistence{
                Enable: criblcontrolplanesdkgo.Pointer(true),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            Description: criblcontrolplanesdkgo.Pointer("before beneficial versus between instead"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesLoki" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputLoki(
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10080,
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](9872.6),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](492856),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](9146.1),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](8080.01),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](4630.89),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](1005.63),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            LokiAPI: "/loki/api/v1/push",
            AuthType: components.AuthenticationTypeOptionsLokiAuthNone.ToPointer(),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("reference yuck marimba analogy excess pish defiantly investigate mill sate"),
            Username: criblcontrolplanesdkgo.Pointer("Renee_Daugherty"),
            Password: criblcontrolplanesdkgo.Pointer("NM_vAnZWTlo0qt1"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://subdued-forager.com"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](1181.96),
            OauthParams: []components.ItemsTypeOauthParams{
                components.ItemsTypeOauthParams{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            OauthHeaders: []components.ItemsTypeOauthHeaders{
                components.ItemsTypeOauthHeaders{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesMetrics" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputMetrics(
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            UDPPort: criblcontrolplanesdkgo.Pointer[float64](8125),
            TCPPort: criblcontrolplanesdkgo.Pointer[float64](3455.14),
            MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2139.64),
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
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
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](4021.68),
            Description: criblcontrolplanesdkgo.Pointer("than in unaware concerning upright fencing deserted"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesModelDrivenTelemetry" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputModelDrivenTelemetry(
        components.InputModelDrivenTelemetry{
            ID: criblcontrolplanesdkgo.Pointer("mdt-source"),
            Type: components.InputModelDrivenTelemetryTypeModelDrivenTelemetry,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 57000,
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
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](1719.64),
            ShutdownTimeoutMs: criblcontrolplanesdkgo.Pointer[float64](2136.89),
            Description: criblcontrolplanesdkgo.Pointer("pfft considering outlandish in bah yearly aw pharmacopoeia"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesMsk" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputMsk(
        components.InputMsk{
            ID: criblcontrolplanesdkgo.Pointer("msk-source"),
            Type: components.InputMskTypeMsk,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
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
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](7039.27),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](4496.31),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](9974.47),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType{
                Disabled: true,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://concrete-petal.info"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](2814.06),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7605.14),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](8934.79),
                Auth: &components.AuthTypeKafkaSchemaRegistry{
                    Disabled: false,
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                    Disabled: criblcontrolplanesdkgo.Pointer(true),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                    CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                },
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](9662.16),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](2520.43),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](2222.88),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](5072.84),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9386.56),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5006.78),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](7646.22),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](6817.9),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto,
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsV2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](8819.58),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](37.9),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](1117.15),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](7867.13),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](195.75),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](4728.15),
            Description: criblcontrolplanesdkgo.Pointer("until uh-huh grandiose notwithstanding via"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesNetflow" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputNetflow(
        components.InputNetflow{
            ID: criblcontrolplanesdkgo.Pointer("netflow-source"),
            Type: components.InputNetflowTypeNetflow,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 2055,
            EnablePassThrough: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](5371.74),
            TemplateCacheMinutes: criblcontrolplanesdkgo.Pointer[float64](9557.55),
            V5Enabled: criblcontrolplanesdkgo.Pointer(true),
            V9Enabled: criblcontrolplanesdkgo.Pointer(false),
            IpfixEnabled: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("pfft tusk until fiercely"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesOffice365Mgmt" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputOffice365Mgmt(
        components.InputOffice365Mgmt{
            ID: criblcontrolplanesdkgo.Pointer("office365-mgmt-source"),
            Type: components.InputOffice365MgmtTypeOffice365Mgmt,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            PlanType: components.SubscriptionPlanOptionsEnterpriseGcc,
            TenantID: "tenant-id",
            AppID: "app-id",
            Timeout: criblcontrolplanesdkgo.Pointer[float64](483.16),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](549.83),
            JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](5203.37),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            PublisherIdentifier: criblcontrolplanesdkgo.Pointer("<value>"),
            ContentConfig: []components.InputOffice365MgmtContentConfig{
                components.InputOffice365MgmtContentConfig{
                    ContentType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Description: criblcontrolplanesdkgo.Pointer("geez absent paintwork yippee evenly providence ha hover aw buzzing"),
                    Interval: criblcontrolplanesdkgo.Pointer[float64](3262.39),
                    LogLevel: components.LogLevelOptionsContentConfigItemsInfo.ToPointer(),
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                },
            },
            IngestionLag: criblcontrolplanesdkgo.Pointer[float64](3768.67),
            RetryRules: &components.RetryRulesType1{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesBackoff,
                Interval: criblcontrolplanesdkgo.Pointer[float64](676.59),
                Limit: criblcontrolplanesdkgo.Pointer[float64](6017.19),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](1187.9),
                Codes: []float64{
                    9919.4,
                    6531.17,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            AuthType: components.AuthenticationMethodOptions1Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("positively indeed sizzle consequently swelter midst including"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesOffice365MsgTrace

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesOffice365MsgTrace" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputOffice365MsgTrace(
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
            URL: "https://reports.office365.com/ecp/reportingwebservice/reporting.svc/MessageTrace",
            Interval: 15,
            StartDate: criblcontrolplanesdkgo.Pointer("<value>"),
            EndDate: criblcontrolplanesdkgo.Pointer("<value>"),
            Timeout: criblcontrolplanesdkgo.Pointer[float64](6172.32),
            DisableTimeFilter: criblcontrolplanesdkgo.Pointer(true),
            AuthType: components.InputOffice365MsgTraceAuthenticationMethodOauthCert.ToPointer(),
            RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
            MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](82.22),
            LogLevel: components.InputOffice365MsgTraceLogLevelError.ToPointer(),
            JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](218.99),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](652.94),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            RetryRules: &components.RetryRulesType1{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesBackoff,
                Interval: criblcontrolplanesdkgo.Pointer[float64](676.59),
                Limit: criblcontrolplanesdkgo.Pointer[float64](6017.19),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](1187.9),
                Codes: []float64{
                    9919.4,
                    6531.17,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            Description: criblcontrolplanesdkgo.Pointer("outfox absent pecan however why almighty failing"),
            Username: criblcontrolplanesdkgo.Pointer("Carole_Medhurst86"),
            Password: criblcontrolplanesdkgo.Pointer("fmYMLhi4cpq7mal"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            ClientID: criblcontrolplanesdkgo.Pointer("<id>"),
            Resource: criblcontrolplanesdkgo.Pointer("<value>"),
            PlanType: components.SubscriptionPlanOptionsEnterpriseGcc.ToPointer(),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            CertOptions: &components.CertOptions{
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: "<value>",
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: "<value>",
            },
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesOffice365Service" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputOffice365Service(
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            PlanType: components.SubscriptionPlanOptionsGcc.ToPointer(),
            TenantID: "tenant-id",
            AppID: "app-id",
            Timeout: criblcontrolplanesdkgo.Pointer[float64](6169.37),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](4019.57),
            JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](474.64),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ContentConfig: []components.InputOffice365ServiceContentConfig{
                components.InputOffice365ServiceContentConfig{
                    ContentType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Description: criblcontrolplanesdkgo.Pointer("ah whenever unlike bookcase towards gadzooks honestly"),
                    Interval: criblcontrolplanesdkgo.Pointer[float64](6065.54),
                    LogLevel: components.LogLevelOptionsContentConfigItemsWarn.ToPointer(),
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                },
            },
            RetryRules: &components.RetryRulesType1{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesBackoff,
                Interval: criblcontrolplanesdkgo.Pointer[float64](676.59),
                Limit: criblcontrolplanesdkgo.Pointer[float64](6017.19),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](1187.9),
                Codes: []float64{
                    9919.4,
                    6531.17,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            AuthType: components.AuthenticationMethodOptions1Secret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("following nice of till whose gadzooks"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesOpenTelemetry

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesOpenTelemetry" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputOpenTelemetry(
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 4317,
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](2734.59),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](231358),
            EnableProxyHeader: "<value>",
            CaptureHeaders: "<value>",
            ActivityLogSampleRate: "<value>",
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](3507.86),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](8180.94),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](3269.77),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            Protocol: components.InputOpenTelemetryProtocolHTTP.ToPointer(),
            ExtractSpans: criblcontrolplanesdkgo.Pointer(true),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(false),
            OtlpVersion: components.InputOpenTelemetryOTLPVersionOneDot3Dot1.ToPointer(),
            AuthType: components.AuthenticationTypeOptionsCredentialsSecret.ToPointer(),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](3610.61),
            Description: criblcontrolplanesdkgo.Pointer("poorly on reckon joyously"),
            Username: criblcontrolplanesdkgo.Pointer("Eleanore_Lockman27"),
            Password: criblcontrolplanesdkgo.Pointer("fbcKwHWOY7xmVhY"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://round-help.com"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](958.28),
            OauthParams: []components.ItemsTypeOauthParams{
                components.ItemsTypeOauthParams{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            OauthHeaders: []components.ItemsTypeOauthHeaders{
                components.ItemsTypeOauthHeaders{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ExtractLogs: criblcontrolplanesdkgo.Pointer(true),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesPrometheus" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputPrometheus(
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            DimensionList: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            DiscoveryType: components.InputPrometheusDiscoveryTypeStatic.ToPointer(),
            Interval: 60,
            LogLevel: components.InputPrometheusLogLevelInfo,
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            Timeout: criblcontrolplanesdkgo.Pointer[float64](6391.07),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](5417.12),
            JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](7932.03),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("so wherever after polarisation by backbone hoarse which duh inside"),
            TargetList: []string{
                "http://localhost:9090/metrics",
            },
            RecordType: components.RecordTypeOptionsAaaa.ToPointer(),
            ScrapePort: criblcontrolplanesdkgo.Pointer[float64](8978.36),
            NameList: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ScrapeProtocol: components.MetricsProtocolHTTPS.ToPointer(),
            ScrapePath: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            UsePublicIP: criblcontrolplanesdkgo.Pointer(false),
            SearchFilter: []components.ItemsTypeSearchFilter{
                components.ItemsTypeSearchFilter{
                    Name: "<value>",
                    Values: []string{
                        "<value 1>",
                    },
                },
            },
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions1V4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](6388.96),
            Username: criblcontrolplanesdkgo.Pointer("Brenden_Reynolds97"),
            Password: criblcontrolplanesdkgo.Pointer("KmtcztBL2FtzSkj"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesPrometheusRw" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputPrometheusRw(
        components.InputPrometheusRw{
            ID: criblcontrolplanesdkgo.Pointer("prometheus-rw-source"),
            Type: components.InputPrometheusRwTypePrometheusRw,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10080,
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](86.47),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](665371),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](7550.39),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](133.32),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](798.3),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](899.7),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            PrometheusAPI: "/write",
            AuthType: components.AuthenticationTypeOptionsPrometheusAuthBasic.ToPointer(),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("circa legitimize since along gah chapel when controvert overcharge"),
            Username: criblcontrolplanesdkgo.Pointer("Ernesto.Rodriguez"),
            Password: criblcontrolplanesdkgo.Pointer("332YGCpd1VZujJu"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://perky-bin.biz/"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](9331.96),
            OauthParams: []components.ItemsTypeOauthParams{
                components.ItemsTypeOauthParams{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            OauthHeaders: []components.ItemsTypeOauthHeaders{
                components.ItemsTypeOauthHeaders{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesRawUdp" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputRawUDP(
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 514,
            MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](6951.72),
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            SingleMsgUDPPackets: criblcontrolplanesdkgo.Pointer(true),
            IngestRawBytes: criblcontrolplanesdkgo.Pointer(true),
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](3490.77),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("testing king knottily magnetize um meh tepid"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesS3" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputS3(
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
            QueueName: "s3-notifications-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfManual.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](5497.69),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](1410.35),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](1210.91),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](5475.5),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](1024.79),
            SkipOnError: criblcontrolplanesdkgo.Pointer(true),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](796.75),
            EnableSQSAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](4479.73),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](9106.64),
            Checkpointing: &components.CheckpointingType{
                Enabled: false,
                Retries: criblcontrolplanesdkgo.Pointer[float64](3583.79),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](859.89),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            TagAfterProcessing: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("after airbrush oh ick safely essay"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ProcessedTagKey: criblcontrolplanesdkgo.Pointer("<value>"),
            ProcessedTagValue: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesS3Inventory" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputS3Inventory(
        components.InputS3Inventory{
            ID: criblcontrolplanesdkgo.Pointer("s3-inventory-source"),
            Type: components.InputS3InventoryTypeS3Inventory,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
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
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](8110.89),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](3914.52),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](7725.89),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](6531.06),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](2786.87),
            SkipOnError: criblcontrolplanesdkgo.Pointer(true),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](2863.98),
            EnableSQSAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](4137.96),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](1617.35),
            Checkpointing: &components.CheckpointingType{
                Enabled: false,
                Retries: criblcontrolplanesdkgo.Pointer[float64](3583.79),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](1459.26),
            ChecksumSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxManifestSizeKB: criblcontrolplanesdkgo.Pointer[int64](373833),
            ValidateInventoryFiles: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("bolster whenever phooey reckon ocelot"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TagAfterProcessing: components.TagAfterProcessingOptionsFalse.ToPointer(),
            ProcessedTagKey: criblcontrolplanesdkgo.Pointer("<value>"),
            ProcessedTagValue: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesSecurityLake" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputSecurityLake(
        components.InputSecurityLake{
            ID: criblcontrolplanesdkgo.Pointer("security-lake-source"),
            Type: components.InputSecurityLakeTypeSecurityLake,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "security-lake-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfSecret.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](8028.96),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](4156.44),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](639.7),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](2329.37),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](9281.87),
            SkipOnError: criblcontrolplanesdkgo.Pointer(true),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](312.72),
            EnableSQSAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](4503.14),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](5363.34),
            Checkpointing: &components.CheckpointingType{
                Enabled: false,
                Retries: criblcontrolplanesdkgo.Pointer[float64](3583.79),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](7033.68),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("ew apropos convection meanwhile why almost"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TagAfterProcessing: components.TagAfterProcessingOptionsTrue.ToPointer(),
            ProcessedTagKey: criblcontrolplanesdkgo.Pointer("<value>"),
            ProcessedTagValue: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesSnmp" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputSnmp(
        components.InputSnmp{
            ID: criblcontrolplanesdkgo.Pointer("snmp-source"),
            Type: components.InputSnmpTypeSnmp,
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
            Host: "192.168.1.1",
            Port: 161,
            SnmpV3Auth: &components.SNMPv3Authentication{
                V3AuthEnabled: false,
                AllowUnmatchedTrap: criblcontrolplanesdkgo.Pointer(false),
                V3Users: []components.InputSnmpV3User{
                    components.InputSnmpV3User{
                        Name: "<value>",
                        AuthProtocol: components.AuthenticationProtocolOptionsV3UserMd5.ToPointer(),
                        AuthKey: criblcontrolplanesdkgo.Pointer("<value>"),
                        PrivProtocol: components.PrivacyProtocolAes256r.ToPointer(),
                        PrivKey: criblcontrolplanesdkgo.Pointer("<value>"),
                    },
                },
            },
            MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5382.63),
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](2276.89),
            VarbindsWithTypes: criblcontrolplanesdkgo.Pointer(true),
            BestEffortParsing: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("difficult giving whose since dramatize about upon"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesSplunk" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputSplunk(
        components.InputSplunk{
            ID: criblcontrolplanesdkgo.Pointer("splunk-source"),
            Type: components.InputSplunkTypeSplunk,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 9997,
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
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](5164.85),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](2368.59),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](8933.92),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](4695.77),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](2286.63),
            AuthTokens: []components.InputSplunkAuthToken{
                components.InputSplunkAuthToken{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("catalog past whitewash seafood"),
                },
            },
            MaxS2Sversion: components.MaxS2SVersionV4.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("clinch phew indeed untrue"),
            UseFwdTimezone: criblcontrolplanesdkgo.Pointer(false),
            DropControlFields: criblcontrolplanesdkgo.Pointer(true),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(true),
            Compress: components.InputSplunkCompressionAlways.ToPointer(),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesSplunkHec" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputSplunkHec(
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
            Host: "0.0.0.0",
            Port: 8088,
            AuthTokens: []components.InputSplunkHecAuthToken{
                components.InputSplunkHecAuthToken{
                    AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
                    TokenSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Token: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                    Description: criblcontrolplanesdkgo.Pointer("wring pfft apostrophize wherever familiarize amused phooey over"),
                    AllowedIndexesAtToken: []string{
                        "<value 1>",
                        "<value 2>",
                    },
                    Metadata: []components.ItemsTypeNotificationMetadata{
                        components.ItemsTypeNotificationMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](784.34),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](120376),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](6037.04),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](2114.37),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](6749.64),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](3764.94),
            EnableHealthCheck: "<value>",
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            SplunkHecAPI: "/services/collector",
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AllowedIndexes: []string{
                "<value 1>",
                "<value 2>",
            },
            SplunkHecAcks: criblcontrolplanesdkgo.Pointer(false),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](7030.61),
            UseFwdTimezone: criblcontrolplanesdkgo.Pointer(false),
            DropControlFields: criblcontrolplanesdkgo.Pointer(true),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(false),
            AccessControlAllowOrigin: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            AccessControlAllowHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            EmitTokenMetrics: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("defenseless majestically lightly blah though unlike accentuate huzzah candid instruction"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesSplunkSearch" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputSplunkSearch(
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
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
            LogLevel: components.InputSplunkSearchLogLevelError.ToPointer(),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](355.18),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](9456.75),
            JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](2620.85),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            RetryRules: &components.RetryRulesType{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesStatic,
                Interval: criblcontrolplanesdkgo.Pointer[float64](5778.31),
                Limit: criblcontrolplanesdkgo.Pointer[float64](8782.9),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](6364.93),
                Codes: []float64{
                    2078.11,
                    6342.8,
                    4459.44,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(true),
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](7185.57),
            AuthType: components.InputSplunkSearchAuthenticationTypeNone.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("hm wicked embossing"),
            Username: criblcontrolplanesdkgo.Pointer("Trace_Schmeler1"),
            Password: criblcontrolplanesdkgo.Pointer("tQ5UHa3ZYCIdjOk"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://dull-finer.com"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](6381.01),
            OauthParams: []components.ItemsTypeOauthParams{
                components.ItemsTypeOauthParams{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            OauthHeaders: []components.ItemsTypeOauthHeaders{
                components.ItemsTypeOauthHeaders{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesSqs" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputSqs(
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "my-queue",
            QueueType: components.InputSqsQueueTypeStandard,
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            CreateQueue: criblcontrolplanesdkgo.Pointer(true),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions3V2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](8576.28),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](8042.98),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](251.38),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](4246.51),
            Description: criblcontrolplanesdkgo.Pointer("onto blah gasp circa overproduce hence when inspection"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](9993.95),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesSyslog" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputSyslog(
        components.CreateInputSyslogInputSyslogSyslog1(
            components.InputSyslogSyslog1{
                ID: criblcontrolplanesdkgo.Pointer("syslog-source"),
                Type: components.InputSyslogType1Syslog,
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
                    MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                    CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                    MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                    Compress: components.CompressionOptionsPqGzip.ToPointer(),
                    PqControls: &components.PqTypePqControls{},
                },
                Host: "0.0.0.0",
                UDPPort: 514,
                TCPPort: criblcontrolplanesdkgo.Pointer[float64](6165.25),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4800.75),
                IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
                TimestampTimezone: criblcontrolplanesdkgo.Pointer("<value>"),
                SingleMsgUDPPackets: criblcontrolplanesdkgo.Pointer(true),
                EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
                KeepFieldsList: []string{
                    "<value 1>",
                },
                OctetCounting: criblcontrolplanesdkgo.Pointer(false),
                InferFraming: criblcontrolplanesdkgo.Pointer(false),
                StrictlyInferOctetCounting: criblcontrolplanesdkgo.Pointer(false),
                AllowNonStandardAppName: criblcontrolplanesdkgo.Pointer(true),
                MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](3407.16),
                SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](3828.73),
                SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](8099.43),
                SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](2179.09),
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
                Metadata: []components.ItemsTypeNotificationMetadata{
                    components.ItemsTypeNotificationMetadata{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](748.79),
                EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(true),
                Description: criblcontrolplanesdkgo.Pointer("than apud obtrude hmph boo boo following story opera litter"),
                EnableEnhancedProxyHeaderParsing: criblcontrolplanesdkgo.Pointer(false),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesSystemMetrics" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputSystemMetrics(
        components.InputSystemMetrics{
            ID: criblcontrolplanesdkgo.Pointer("system-metrics-source"),
            Type: components.InputSystemMetricsTypeSystemMetrics,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](2717.34),
            Host: &components.InputSystemMetricsHost{
                Mode: components.ModeOptionsHostDisabled.ToPointer(),
                Custom: &components.InputSystemMetricsCustom{
                    System: &components.InputSystemMetricsSystem{
                        Mode: components.InputSystemMetricsSystemModeAll.ToPointer(),
                        Processes: criblcontrolplanesdkgo.Pointer(true),
                    },
                    CPU: &components.InputSystemMetricsCPU{
                        Mode: components.InputSystemMetricsCPUModeDisabled.ToPointer(),
                        PerCPU: criblcontrolplanesdkgo.Pointer(false),
                        Detail: criblcontrolplanesdkgo.Pointer(true),
                        Time: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Memory: &components.InputSystemMetricsMemory{
                        Mode: components.InputSystemMetricsMemoryModeBasic.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(true),
                    },
                    Network: &components.InputSystemMetricsNetwork{
                        Mode: components.InputSystemMetricsNetworkModeCustom.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Protocols: criblcontrolplanesdkgo.Pointer(false),
                        Devices: []string{
                            "<value 1>",
                            "<value 2>",
                        },
                        PerInterface: criblcontrolplanesdkgo.Pointer(true),
                    },
                    Disk: &components.InputSystemMetricsDisk{
                        Mode: components.InputSystemMetricsDiskModeBasic.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Inodes: criblcontrolplanesdkgo.Pointer(false),
                        Devices: []string{
                            "<value 1>",
                        },
                        Mountpoints: []string{
                            "<value 1>",
                        },
                        Fstypes: []string{
                            "<value 1>",
                            "<value 2>",
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
                        IncludeChildren: criblcontrolplanesdkgo.Pointer(false),
                    },
                },
            },
            Container: &components.Container{
                Mode: components.ContainerModeBasic.ToPointer(),
                DockerSocket: []string{
                    "<value 1>",
                    "<value 2>",
                },
                DockerTimeout: criblcontrolplanesdkgo.Pointer[float64](6763.69),
                Filters: []components.InputSystemMetricsFilter{
                    components.InputSystemMetricsFilter{
                        Expr: "<value>",
                    },
                },
                AllContainers: criblcontrolplanesdkgo.Pointer(true),
                PerDevice: criblcontrolplanesdkgo.Pointer(false),
                Detail: criblcontrolplanesdkgo.Pointer(false),
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Persistence: &components.InputSystemMetricsPersistence{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            Description: criblcontrolplanesdkgo.Pointer("supposing pale ack provider finding pace nippy"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesSystemState" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputSystemState(
        components.InputSystemState{
            ID: criblcontrolplanesdkgo.Pointer("system-state-source"),
            Type: components.InputSystemStateTypeSystemState,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](1268.95),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
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
                    Enable: criblcontrolplanesdkgo.Pointer(false),
                },
                DNS: &components.DNS{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                User: &components.UsersAndGroups{
                    Enable: criblcontrolplanesdkgo.Pointer(false),
                },
                Firewall: &components.Firewall{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
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
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
                Compress: components.DataCompressionFormatOptionsPersistenceNone.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            DisableNativeModule: criblcontrolplanesdkgo.Pointer(true),
            DisableNativeLastLogModule: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("um preheat unto bolster supposing besides"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesTcp" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputTCP(
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
            Host: "0.0.0.0",
            Port: 10090,
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
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](3827.91),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](5013.14),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](1943.33),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](7818.52),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](2553.44),
            EnableHeader: criblcontrolplanesdkgo.Pointer(false),
            Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("through underneath nervously rejoin collaborate mmm"),
            AuthToken: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
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
### Example Usage: InputCreateExamplesTcpjson

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesTcpjson" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputTcpjson(
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10090,
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
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](4131.12),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](8956.21),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](8941.39),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](9020.4),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(false),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("despite flashy oof gigantic ramp oh"),
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
### Example Usage: InputCreateExamplesWef

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesWef" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputWef(
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 5985,
            AuthMethod: components.InputWefAuthenticationMethodClientCert.ToPointer(),
            TLS: &components.MTLSSettings{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](7532.69),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](631259),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](9531.47),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(true),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](9065.47),
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
                    Locale: criblcontrolplanesdkgo.Pointer("hi"),
                    QuerySelector: components.QueryBuilderModeXML.ToPointer(),
                    Metadata: []components.ItemsTypeNotificationMetadata{
                        components.ItemsTypeNotificationMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                    Queries: []components.Query{
                        components.Query{
                            Path: "/opt/sbin",
                            QueryExpression: "<value>",
                        },
                    },
                    XMLQuery: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("brr dearly woot furthermore apropos save cultivated"),
            LogFingerprintMismatch: criblcontrolplanesdkgo.Pointer(false),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesWinEventLogs" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputWinEventLogs(
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
            LogNames: []string{
                "Application",
                "System",
            },
            ReadMode: components.ReadModeOldest.ToPointer(),
            EventFormat: components.EventFormatXML.ToPointer(),
            DisableNativeModule: criblcontrolplanesdkgo.Pointer(true),
            Interval: criblcontrolplanesdkgo.Pointer[float64](5532.44),
            BatchSize: criblcontrolplanesdkgo.Pointer[float64](5477.98),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            MaxEventBytes: criblcontrolplanesdkgo.Pointer[float64](7625.68),
            Description: criblcontrolplanesdkgo.Pointer("thankfully without badly eventually yet mid desk what stool jellyfish"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesWindowsMetrics" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputWindowsMetrics(
        components.InputWindowsMetrics{
            ID: criblcontrolplanesdkgo.Pointer("windows-metrics-source"),
            Type: components.InputWindowsMetricsTypeWindowsMetrics,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](359.49),
            Host: &components.InputWindowsMetricsHost{
                Mode: components.ModeOptionsHostAll.ToPointer(),
                Custom: &components.InputWindowsMetricsCustom{
                    System: &components.InputWindowsMetricsSystem{
                        Mode: components.InputWindowsMetricsSystemModeDisabled.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(true),
                    },
                    CPU: &components.InputWindowsMetricsCPU{
                        Mode: components.InputWindowsMetricsCPUModeCustom.ToPointer(),
                        PerCPU: criblcontrolplanesdkgo.Pointer(false),
                        Detail: criblcontrolplanesdkgo.Pointer(true),
                        Time: criblcontrolplanesdkgo.Pointer(true),
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
                        },
                        PerInterface: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Disk: &components.InputWindowsMetricsDisk{
                        Mode: components.InputWindowsMetricsDiskModeBasic.ToPointer(),
                        PerVolume: criblcontrolplanesdkgo.Pointer(true),
                        Detail: criblcontrolplanesdkgo.Pointer(true),
                        Volumes: []string{
                            "<value 1>",
                            "<value 2>",
                            "<value 3>",
                        },
                    },
                },
            },
            Process: &components.ProcessType{
                Sets: []components.ItemsTypeProcessSets{
                    components.ItemsTypeProcessSets{
                        Name: "<value>",
                        Filter: "<value>",
                        IncludeChildren: criblcontrolplanesdkgo.Pointer(false),
                    },
                },
            },
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
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
            DisableNativeModule: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("scout wherever criminal busily past"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesWiz" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputWiz(
        components.InputWiz{
            ID: criblcontrolplanesdkgo.Pointer("wiz-source"),
            Type: components.InputWizTypeWiz,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Endpoint: "https://api.wiz.io",
            AuthURL: "https://auth.wiz.io/oauth/token",
            AuthAudienceOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            ClientID: "client-id",
            ContentConfig: []components.InputWizContentConfig{},
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](2712.73),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](422.95),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](1889.25),
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            RetryRules: &components.RetryRulesType{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesStatic,
                Interval: criblcontrolplanesdkgo.Pointer[float64](5778.31),
                Limit: criblcontrolplanesdkgo.Pointer[float64](8782.9),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](6364.93),
                Codes: []float64{
                    2078.11,
                    6342.8,
                    4459.44,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(true),
            },
            AuthType: components.AuthenticationMethodOptions1Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("freely longingly clearly near shush busily consequently where"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: InputCreateExamplesWizWebhook

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesWizWebhook" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputWizWebhook(
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](2628.01),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](872518),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](2983.78),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](3666.47),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](3047.72),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](4969.63),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            BreakerRulesets: []string{
                "<value 1>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](4980.59),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
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
            },
            AuthTokensExt: []components.ItemsTypeAuthTokensExt{
                components.ItemsTypeAuthTokensExt{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("unless scuttle pixellate hm aw furthermore tensely pro how beloved"),
                    Metadata: []components.ItemsTypeNotificationMetadata{
                        components.ItemsTypeNotificationMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("equally tremendously lest pace starboard legal adventurously gosh because"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputCreateExamplesZscalerHec" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputZscalerHec(
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 8088,
            AuthTokens: []components.InputZscalerHecAuthToken{
                components.InputZscalerHecAuthToken{
                    AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
                    TokenSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Token: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                    Description: criblcontrolplanesdkgo.Pointer("petal pro what interesting"),
                    AllowedIndexesAtToken: []string{
                        "<value 1>",
                        "<value 2>",
                        "<value 3>",
                    },
                    Metadata: []components.ItemsTypeNotificationMetadata{
                        components.ItemsTypeNotificationMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
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
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](3918.05),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](400843),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(true),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(true),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](1080.72),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](9338.93),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](6021.36),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5654.66),
            EnableHealthCheck: "<value>",
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("<value>"),
            HecAPI: "/services/collector",
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
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
            },
            AccessControlAllowHeaders: []string{
                "<value 1>",
            },
            EmitTokenMetrics: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("supposing naturally through characterization hm blank reproach ferociously midwife selfish"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputExamplesCribl" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputCribl(
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Filter: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("now snowplow because agreement forenenst shabby treble"),
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

<!-- UsageSnippet language="go" operationID="updateInputById" method="patch" path="/system/inputs/{id}" example="InputExamplesCriblMetrics" -->
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

    res, err := s.Sources.Update(ctx, "<id>", components.CreateInputCriblmetrics(
        components.InputCriblmetrics{
            ID: criblcontrolplanesdkgo.Pointer("cribl-metrics-source"),
            Type: components.InputCriblmetricsTypeCriblmetrics,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9959.95),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](4085.76),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                Path: criblcontrolplanesdkgo.Pointer("/usr/obj"),
                Compress: components.CompressionOptionsPqGzip.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Prefix: criblcontrolplanesdkgo.Pointer("<value>"),
            FullFidelity: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeNotificationMetadata{
                components.ItemsTypeNotificationMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("waterspout lest utterly minus pomelo sandy now zowie"),
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