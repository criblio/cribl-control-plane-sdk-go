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
        operations.CreateInputInputAppscope{
            ID: "appscope-source",
            Type: operations.CreateInputTypeAppscopeAppscope,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](1000),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](30),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](0),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            EnableUnixPath: criblcontrolplanesdkgo.Pointer(false),
            Filter: &operations.CreateInputFilterAppscope{
                Allow: []operations.CreateInputAllow{
                    operations.CreateInputAllow{
                        Procname: "<value>",
                        Arg: criblcontrolplanesdkgo.Pointer("<value>"),
                        Config: "<value>",
                    },
                },
                TransportURL: criblcontrolplanesdkgo.Pointer("https://drab-scrap.info/"),
            },
            Persistence: &operations.CreateInputPersistenceAppscope{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("10m"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("1GB"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("24h"),
                Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/appscope"),
            },
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("if deserted boohoo red chops excepting know stay bah"),
            Host: criblcontrolplanesdkgo.Pointer("0.0.0.0"),
            Port: criblcontrolplanesdkgo.Pointer[float64](9109),
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            UnixSocketPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/appscope.sock"),
            UnixSocketPerms: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthToken: criblcontrolplanesdkgo.Pointer(""),
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
        operations.CreateInputInputAzureBlob{
            ID: "azure-blob-source",
            Type: operations.CreateInputTypeAzureBlobAzureBlob,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "azure-blob-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("/.*/"),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](600),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](1),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](1),
            ServicePeriodSecs: criblcontrolplanesdkgo.Pointer[float64](5),
            SkipOnError: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            BreakerRulesets: []string{
                "<value 1>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](5),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](600),
            AuthType: components.AuthenticationMethodOptionsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("while brisk meanwhile kaleidoscopic ack ah above psst throughout guide"),
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
        operations.CreateInputInputCloudflareHec{
            ID: "cloudflare-hec-source",
            Type: operations.CreateInputTypeCloudflareHecCloudflareHec,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 8088,
            AuthTokens: []operations.CreateInputAuthTokenCloudflareHec{
                operations.CreateInputAuthTokenCloudflareHec{
                    AuthType: operations.CreateInputAuthTokenAuthenticationMethodSecret.ToPointer(),
                    TokenSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Token: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("aha hydrocarbon after plain"),
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
            TLS: &operations.CreateInputTLSSettingsServerSide{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: "<value>",
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
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
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
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
            Description: criblcontrolplanesdkgo.Pointer("hateful pike whose or interestingly exotic"),
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
        operations.CreateInputInputCollection{
            ID: "collection-source",
            Type: operations.CreateInputTypeCollectionCollection,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            BreakerRulesets: []string{
                "<value 1>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            Preprocess: &components.PreprocessType{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                },
            },
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
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
        operations.CreateInputInputConfluentCloud{
            ID: "confluent-cloud-source",
            Type: operations.CreateInputTypeConfluentCloudConfluentCloud,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Brokers: []string{
                "pkc-xxxxx.us-east-1.aws.confluent.cloud:9092",
            },
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
            },
            Topics: []string{
                "logs",
            },
            GroupID: criblcontrolplanesdkgo.Pointer("Cribl"),
            FromBeginning: criblcontrolplanesdkgo.Pointer(true),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType{
                Disabled: true,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("http://localhost:8081"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](1),
                Auth: &components.AuthTypeKafkaSchemaRegistry{
                    Disabled: true,
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                    CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
                },
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](30000),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](300),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](10000),
            Sasl: &components.AuthenticationType{
                Disabled: true,
                Username: criblcontrolplanesdkgo.Pointer("Monte_Thiel32"),
                Password: criblcontrolplanesdkgo.Pointer("KI_orHyojtOdRdG"),
                AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslPlain.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(false),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://dependable-pulse.net"),
                ClientID: criblcontrolplanesdkgo.Pointer("<id>"),
                OauthSecretType: criblcontrolplanesdkgo.Pointer("secret"),
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
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](3000),
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](4914.59),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](5168.01),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](1048576),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](10485760),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](0),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("comfortable whole veto certainly"),
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
        operations.CreateInputInputCriblHTTP{
            ID: "cribl-http-source",
            Type: operations.CreateInputTypeCriblHTTPCriblHTTP,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10080,
            AuthTokens: []components.ItemsTypeAuthTokens{
                components.ItemsTypeAuthTokens{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("weary likewise meh stoop upwardly amount violently throughout upwardly bathrobe"),
                },
            },
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("fledgling although substantial mmm but what immaculate hmph"),
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
        operations.CreateInputInputCriblLakeHTTP{
            ID: "cribl-lake-http-source",
            Type: operations.CreateInputTypeCriblLakeHTTPCriblLakeHTTP,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
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
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            CriblAPI: criblcontrolplanesdkgo.Pointer("/cribl"),
            ElasticAPI: criblcontrolplanesdkgo.Pointer("/elastic"),
            SplunkHecAPI: criblcontrolplanesdkgo.Pointer("/services/collector"),
            SplunkHecAcks: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthTokensExt: []operations.CreateInputAuthTokensExt{
                operations.CreateInputAuthTokensExt{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("goat thorough like"),
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                    SplunkHecMetadata: &operations.CreateInputSplunkHecMetadata{
                        Enabled: criblcontrolplanesdkgo.Pointer(true),
                        DefaultDataset: criblcontrolplanesdkgo.Pointer("<value>"),
                        AllowedIndexesAtToken: []string{
                            "<value 1>",
                            "<value 2>",
                        },
                    },
                    ElasticsearchMetadata: &operations.CreateInputElasticsearchMetadata{
                        Enabled: criblcontrolplanesdkgo.Pointer(false),
                        DefaultDataset: criblcontrolplanesdkgo.Pointer("<value>"),
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("out regarding until pull avow hunt"),
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
        operations.CreateInputInputCriblTCP{
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
            },
            Connections: []components.ItemsTypeConnectionsOptional{
                components.ItemsTypeConnectionsOptional{
                    Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                    Output: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Pq: &components.PqType{
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10090,
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](1000),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](30),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
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
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("easily bah hierarchy truthfully how brr victoriously"),
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("whenever cheerfully average lovingly meh heartfelt leading scratchy make"),
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
        operations.CreateInputInputCrowdstrike{
            ID: "crowdstrike-source",
            Type: operations.CreateInputTypeCrowdstrikeCrowdstrike,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "crowdstrike-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("/.*/"),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](1),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](21600),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](1),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](300),
            SkipOnError: criblcontrolplanesdkgo.Pointer(false),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
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
                Enabled: false,
                Retries: criblcontrolplanesdkgo.Pointer[float64](5),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](10),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("aha wherever tooth ack"),
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
        operations.CreateInputInputDatadogAgent{
            ID: "datadog-agent-source",
            Type: operations.CreateInputTypeDatadogAgentDatadogAgent,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 8126,
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ProxyMode: &operations.CreateInputProxyModeDatadogAgent{
                Enabled: false,
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            },
            Description: criblcontrolplanesdkgo.Pointer("meanwhile trusting scrutinise except settle"),
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
        operations.CreateInputInputDatagen{
            ID: "datagen-source",
            Type: operations.CreateInputTypeDatagenDatagen,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Samples: []operations.CreateInputSample{
                operations.CreateInputSample{
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
            Description: criblcontrolplanesdkgo.Pointer("quiet indeed arrogantly circumnavigate greedily wheel"),
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
        operations.CreateInputInputEdgePrometheus{
            ID: "edge-prometheus-source",
            Type: operations.CreateInputTypeEdgePrometheusEdgePrometheus,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            DimensionList: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            DiscoveryType: operations.CreateInputDiscoveryTypeEdgePrometheusStatic,
            Interval: 60,
            Timeout: criblcontrolplanesdkgo.Pointer[float64](5000),
            Persistence: &components.DiskSpoolingType{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("10m"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("1GB"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("24h"),
                Compress: components.CompressionOptionsPersistenceGzip.ToPointer(),
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthType: operations.CreateInputAuthenticationMethodEdgePrometheusManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("geez however knight ah whenever bulky throughout troubled although daintily"),
            Targets: []operations.CreateInputTarget{
                operations.CreateInputTarget{
                    Protocol: components.ProtocolOptionsTargetsItemsHTTP.ToPointer(),
                    Host: "localhost",
                    Port: criblcontrolplanesdkgo.Pointer[float64](9090),
                    Path: criblcontrolplanesdkgo.Pointer("/metrics"),
                },
            },
            RecordType: components.RecordTypeOptionsSrv.ToPointer(),
            ScrapePort: criblcontrolplanesdkgo.Pointer[float64](9090),
            NameList: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ScrapeProtocol: components.ProtocolOptionsTargetsItemsHTTP.ToPointer(),
            ScrapePath: criblcontrolplanesdkgo.Pointer("/metrics"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
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
            SignatureVersion: components.SignatureVersionOptions1V4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            ScrapeProtocolExpr: criblcontrolplanesdkgo.Pointer("metadata.annotations['prometheus.io/scheme'] || 'http'"),
            ScrapePortExpr: criblcontrolplanesdkgo.Pointer("metadata.annotations['prometheus.io/port'] || 9090"),
            ScrapePathExpr: criblcontrolplanesdkgo.Pointer("metadata.annotations['prometheus.io/path'] || '/metrics'"),
            PodFilter: []operations.CreateInputPodFilter{
                operations.CreateInputPodFilter{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("stay strictly under including corner ick hearten"),
                },
            },
            Username: criblcontrolplanesdkgo.Pointer("Hadley97"),
            Password: criblcontrolplanesdkgo.Pointer("Hkq7Phu4rXHsZ3j"),
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
        operations.CreateInputInputElastic{
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "localhost",
            Port: 9200,
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            ElasticAPI: "/",
            AuthType: operations.CreateInputAuthenticationTypeElasticNone.ToPointer(),
            APIVersion: operations.CreateInputAPIVersionEightDot3Dot2.ToPointer(),
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
            ProxyMode: &operations.CreateInputProxyModeElastic{
                Enabled: false,
                AuthType: operations.CreateInputProxyModeAuthenticationMethodNone.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("Kathryn_Senger78"),
                Password: criblcontrolplanesdkgo.Pointer("WPXuh03pB0ItRx8"),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                URL: criblcontrolplanesdkgo.Pointer("https://likely-abacus.info"),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                RemoveHeaders: []string{
                    "<value 1>",
                },
                TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](60),
                TemplateURL: criblcontrolplanesdkgo.Pointer("https://amused-glider.biz"),
            },
            Description: criblcontrolplanesdkgo.Pointer("passionate gratefully usually miserably uh-huh"),
            Username: criblcontrolplanesdkgo.Pointer("Caterina_McClure27"),
            Password: criblcontrolplanesdkgo.Pointer("5FvvRdikRVomwpo"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthTokens: []string{
                "<value 1>",
            },
            CustomAPIVersion: criblcontrolplanesdkgo.Pointer("{\n    \"name\": \"AzU84iL\",\n    \"cluster_name\": \"cribl\",\n    \"cluster_uuid\": \"Js6_Z2VKS3KbfRSxPmPbaw\",\n    \"version\": {\n        \"number\": \"8.3.2\",\n        \"build_type\": \"tar\",\n        \"build_hash\": \"bca0c8d\",\n        \"build_date\": \"2019-10-16T06:19:49.319352Z\",\n        \"build_snapshot\": false,\n        \"lucene_version\": \"9.7.2\",\n        \"minimum_wire_compatibility_version\": \"7.17.0\",\n        \"minimum_index_compatibility_version\": \"7.0.0\"\n    },\n    \"tagline\": \"You Know, for Search\"\n}"),
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
        operations.CreateInputInputEventhub{
            ID: "eventhub-source",
            Type: operations.CreateInputTypeEventhubEventhub,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Brokers: []string{
                "myeventhub.servicebus.windows.net:9093",
            },
            Topics: []string{
                "logs",
            },
            GroupID: criblcontrolplanesdkgo.Pointer("Cribl"),
            FromBeginning: criblcontrolplanesdkgo.Pointer(true),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](30000),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](300),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](10000),
            Sasl: &components.AuthenticationType1{
                Disabled: false,
                AuthType: components.AuthenticationMethodOptionsSasl1Manual.ToPointer(),
                Password: criblcontrolplanesdkgo.Pointer("rkhgMwb5YCBiRV0"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSasl1Plain.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("$ConnectionString"),
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
                Disabled: false,
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            },
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](3000),
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](7747.29),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](9425.98),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](1048576),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](10485760),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](0),
            MinimizeDuplicates: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("nerve inside gerbil orient yowza maroon ha"),
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
        operations.CreateInputInputExec{
            ID: "exec-source",
            Type: operations.CreateInputInputExecTypeExec,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Command: "echo \"Hello World\"",
            Retries: criblcontrolplanesdkgo.Pointer[float64](10),
            ScheduleType: operations.CreateInputScheduleTypeInterval.ToPointer(),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("hence after waft whoa a oof robust"),
            Interval: criblcontrolplanesdkgo.Pointer[float64](60),
            CronSchedule: criblcontrolplanesdkgo.Pointer("* * * * *"),
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
        operations.CreateInputInputFile{
            ID: "file-source",
            Type: operations.CreateInputInputFileTypeFile,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Mode: operations.CreateInputInputFileModeManual.ToPointer(),
            Interval: criblcontrolplanesdkgo.Pointer[float64](10),
            Filenames: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            FilterArchivedFiles: criblcontrolplanesdkgo.Pointer(false),
            TailOnly: criblcontrolplanesdkgo.Pointer(true),
            IdleTimeout: criblcontrolplanesdkgo.Pointer[float64](300),
            MinAgeDur: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxAgeDur: criblcontrolplanesdkgo.Pointer("<value>"),
            CheckFileModTime: criblcontrolplanesdkgo.Pointer(false),
            ForceText: criblcontrolplanesdkgo.Pointer(false),
            HashLen: criblcontrolplanesdkgo.Pointer[float64](256),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            Description: criblcontrolplanesdkgo.Pointer("oh whoa entice when phooey address"),
            Path: criblcontrolplanesdkgo.Pointer("/usr/local/src"),
            Depth: criblcontrolplanesdkgo.Pointer[float64](9351.84),
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
        operations.CreateInputInputFirehose{
            ID: "firehose-source",
            Type: operations.CreateInputTypeFirehoseFirehose,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
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
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("though instead talkative mid eek deadly these"),
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
        operations.CreateInputInputGooglePubsub{
            ID: "google-pubsub-source",
            Type: operations.CreateInputTypeGooglePubsubGooglePubsub,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
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
            MaxBacklog: criblcontrolplanesdkgo.Pointer[float64](1000),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("mortise yowza clearly er yippee taxicab never onto nor singe"),
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
        operations.CreateCreateInputInputGrafanaUnionCreateInputInputGrafanaGrafana1(
            operations.CreateInputInputGrafanaGrafana1{
                ID: "grafana-source",
                Type: operations.CreateInputInputGrafanaType1Grafana,
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
                    MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                    CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                    MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                    MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                    Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                    Compress: components.CompressionOptionsPqNone.ToPointer(),
                    PqControls: &components.PqTypePqControls{},
                },
                Host: "0.0.0.0",
                Port: 10080,
                TLS: &components.TLSSettingsServerSideType{
                    Disabled: criblcontrolplanesdkgo.Pointer(true),
                    RequestCert: criblcontrolplanesdkgo.Pointer(false),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                    CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                    CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                    PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                },
                MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
                MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
                EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
                CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
                ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
                SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
                KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
                EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
                IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
                PrometheusAPI: "/api/prom/push",
                LokiAPI: criblcontrolplanesdkgo.Pointer("/loki/api/v1/push"),
                PrometheusAuth: &operations.CreateInputPrometheusAuth1{
                    AuthType: components.AuthenticationTypeOptionsPrometheusAuthNone.ToPointer(),
                    Username: criblcontrolplanesdkgo.Pointer("Salvador1"),
                    Password: criblcontrolplanesdkgo.Pointer("ZQk_l_P10FR08Qf"),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                LokiAuth: &operations.CreateInputLokiAuth1{
                    AuthType: components.AuthenticationTypeOptionsLokiAuthNone.ToPointer(),
                    Username: criblcontrolplanesdkgo.Pointer("Philip50"),
                    Password: criblcontrolplanesdkgo.Pointer("IKe8kW3jPl5Tei7"),
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
                Description: criblcontrolplanesdkgo.Pointer("fooey after properly often charlatan"),
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
        operations.CreateInputInputHTTP{
            ID: "http-source",
            Type: operations.CreateInputTypeHTTPHTTP,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
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
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            CriblAPI: criblcontrolplanesdkgo.Pointer("/cribl"),
            ElasticAPI: criblcontrolplanesdkgo.Pointer("/elastic"),
            SplunkHecAPI: criblcontrolplanesdkgo.Pointer("/services/collector"),
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
                    Description: criblcontrolplanesdkgo.Pointer("though transplant dreary sweetly which"),
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("calmly ethyl scramble thick formamide a unfortunately"),
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
        operations.CreateInputInputHTTPRaw{
            ID: "http-raw-source",
            Type: operations.CreateInputTypeHTTPRawHTTPRaw,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
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
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
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
                    Description: criblcontrolplanesdkgo.Pointer("since reconsideration scoff"),
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("grave carnival siege uh-huh through behind excepting notwithstanding"),
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
        operations.CreateInputInputJournalFiles{
            ID: "journal-files-source",
            Type: operations.CreateInputInputJournalFilesTypeJournalFiles,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Path: "/var/log/journal",
            Interval: criblcontrolplanesdkgo.Pointer[float64](10),
            Journals: []string{
                "system",
            },
            Rules: []operations.CreateInputInputJournalFilesRule{
                operations.CreateInputInputJournalFilesRule{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("yuck drat ew our the lecture likewise"),
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
            Description: criblcontrolplanesdkgo.Pointer("oof depart loyalty reapply bog"),
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
        operations.CreateInputInputKafka{
            ID: "kafka-source",
            Type: operations.CreateInputTypeKafkaKafka,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Brokers: []string{
                "localhost:9092",
            },
            Topics: []string{
                "logs",
            },
            GroupID: criblcontrolplanesdkgo.Pointer("Cribl"),
            FromBeginning: criblcontrolplanesdkgo.Pointer(true),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType{
                Disabled: true,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("http://localhost:8081"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](1),
                Auth: &components.AuthTypeKafkaSchemaRegistry{
                    Disabled: true,
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                    CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
                },
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](30000),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](300),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](10000),
            Sasl: &components.AuthenticationType{
                Disabled: true,
                Username: criblcontrolplanesdkgo.Pointer("Monte_Thiel32"),
                Password: criblcontrolplanesdkgo.Pointer("KI_orHyojtOdRdG"),
                AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslPlain.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(false),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://dependable-pulse.net"),
                ClientID: criblcontrolplanesdkgo.Pointer("<id>"),
                OauthSecretType: criblcontrolplanesdkgo.Pointer("secret"),
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
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
            },
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](3000),
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](2313.91),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](7394.97),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](1048576),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](10485760),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](0),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("deliberately greatly hmph mom fatally"),
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
        operations.CreateInputInputKinesis{
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            StreamName: "my-stream",
            ServiceInterval: criblcontrolplanesdkgo.Pointer[float64](1),
            ShardExpr: criblcontrolplanesdkgo.Pointer("true"),
            ShardIteratorType: operations.CreateInputShardIteratorStartTrimHorizon.ToPointer(),
            PayloadFormat: operations.CreateInputRecordDataFormatCribl.ToPointer(),
            GetRecordsLimit: criblcontrolplanesdkgo.Pointer[float64](5000),
            GetRecordsLimitTotal: criblcontrolplanesdkgo.Pointer[float64](20000),
            LoadBalancingAlgorithm: operations.CreateInputShardLoadBalancingConsistentHashing.ToPointer(),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions2V4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            VerifyKPLCheckSums: criblcontrolplanesdkgo.Pointer(false),
            AvoidDuplicates: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("finally instead less till ew duh ew"),
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
        operations.CreateInputInputKubeEvents{
            ID: "kube-events-source",
            Type: operations.CreateInputTypeKubeEventsKubeEvents,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Rules: []components.ItemsTypeRules{
                components.ItemsTypeRules{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("before settler fortunately meh nice bludgeon under soybean jam-packed"),
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("stall unpleasant newsprint"),
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
        operations.CreateInputInputKubeLogs{
            ID: "kube-logs-source",
            Type: operations.CreateInputTypeKubeLogsKubeLogs,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](15),
            Rules: []operations.CreateInputRuleKubeLogs{
                operations.CreateInputRuleKubeLogs{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("over questionably yesterday"),
                },
            },
            Timestamps: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Persistence: &components.DiskSpoolingType{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("10m"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("1GB"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("24h"),
                Compress: components.CompressionOptionsPersistenceGzip.ToPointer(),
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("warp verbally through amnesty"),
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
        operations.CreateInputInputKubeMetrics{
            ID: "kube-metrics-source",
            Type: operations.CreateInputTypeKubeMetricsKubeMetrics,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](15),
            Rules: []components.ItemsTypeRules{
                components.ItemsTypeRules{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("saloon where faithfully"),
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Persistence: &operations.CreateInputPersistenceKubeMetrics{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("10m"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("1GB"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("24h"),
                Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/kube_metrics"),
            },
            Description: criblcontrolplanesdkgo.Pointer("traduce calculus where fun calculating uh-huh delightfully"),
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
        operations.CreateInputInputLoki{
            ID: "loki-source",
            Type: operations.CreateInputTypeLokiLoki,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10080,
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            LokiAPI: "/loki/api/v1/push",
            AuthType: components.AuthenticationTypeOptionsLokiAuthNone.ToPointer(),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("nudge valley instead sun er"),
            Username: criblcontrolplanesdkgo.Pointer("Werner47"),
            Password: criblcontrolplanesdkgo.Pointer("b7xIUX8OdaWkQ40"),
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
        operations.CreateInputInputMetrics{
            ID: "metrics-source",
            Type: operations.CreateInputTypeMetricsMetrics,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            UDPPort: criblcontrolplanesdkgo.Pointer[float64](8125),
            TCPPort: criblcontrolplanesdkgo.Pointer[float64](4526.8),
            MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](86.97),
            Description: criblcontrolplanesdkgo.Pointer("contrast kindly arcade total while warmly"),
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
        operations.CreateInputInputModelDrivenTelemetry{
            ID: "mdt-source",
            Type: operations.CreateInputTypeModelDrivenTelemetryModelDrivenTelemetry,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 57000,
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](1000),
            ShutdownTimeoutMs: criblcontrolplanesdkgo.Pointer[float64](5000),
            Description: criblcontrolplanesdkgo.Pointer("unnecessarily forenenst finally ick qua far or pack outlandish"),
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
        operations.CreateInputInputMsk{
            ID: "msk-source",
            Type: operations.CreateInputTypeMskMsk,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Brokers: []string{
                "b-1.example.xxxxx.c2.kafka.us-east-1.amazonaws.com:9092",
            },
            Topics: []string{
                "logs",
            },
            GroupID: criblcontrolplanesdkgo.Pointer("Cribl"),
            FromBeginning: criblcontrolplanesdkgo.Pointer(true),
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](3000),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType{
                Disabled: true,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("http://localhost:8081"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](1),
                Auth: &components.AuthTypeKafkaSchemaRegistry{
                    Disabled: true,
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                    CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
                },
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](30000),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](300),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](10000),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto,
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
            },
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](8981.95),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](7775.93),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](1048576),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](10485760),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](0),
            Description: criblcontrolplanesdkgo.Pointer("whereas gah internationalize"),
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
        operations.CreateInputInputNetflow{
            ID: "netflow-source",
            Type: operations.CreateInputTypeNetflowNetflow,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 2055,
            EnablePassThrough: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](1130.9),
            TemplateCacheMinutes: criblcontrolplanesdkgo.Pointer[float64](30),
            V5Enabled: criblcontrolplanesdkgo.Pointer(true),
            V9Enabled: criblcontrolplanesdkgo.Pointer(true),
            IpfixEnabled: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("dicker sans scarper amid which until yet yin"),
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
        operations.CreateInputInputOffice365Mgmt{
            ID: "office365-mgmt-source",
            Type: operations.CreateInputTypeOffice365MgmtOffice365Mgmt,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            PlanType: components.SubscriptionPlanOptionsEnterpriseGcc,
            TenantID: "tenant-id",
            AppID: "app-id",
            Timeout: criblcontrolplanesdkgo.Pointer[float64](300),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](30),
            JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](3),
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            PublisherIdentifier: criblcontrolplanesdkgo.Pointer("<value>"),
            ContentConfig: []operations.CreateInputContentConfigOffice365Mgmt{
                operations.CreateInputContentConfigOffice365Mgmt{
                    ContentType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Description: criblcontrolplanesdkgo.Pointer("er aha till deploy"),
                    Interval: criblcontrolplanesdkgo.Pointer[float64](3564.68),
                    LogLevel: components.LogLevelOptionsContentConfigItemsError.ToPointer(),
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                },
            },
            IngestionLag: criblcontrolplanesdkgo.Pointer[float64](0),
            RetryRules: &components.RetryRulesType1{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesBackoff,
                Interval: criblcontrolplanesdkgo.Pointer[float64](1000),
                Limit: criblcontrolplanesdkgo.Pointer[float64](5),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](2),
                Codes: []float64{
                    3152.61,
                    1476.48,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            AuthType: components.AuthenticationMethodOptions1Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("overfeed bitterly notwithstanding tidy equally however"),
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
        operations.CreateInputInputOffice365MsgTrace{
            ID: "office365-msg-trace-source",
            Type: operations.CreateInputTypeOffice365MsgTraceOffice365MsgTrace,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            URL: "https://reports.office365.com/ecp/reportingwebservice/reporting.svc/MessageTrace",
            Interval: 15,
            StartDate: criblcontrolplanesdkgo.Pointer("<value>"),
            EndDate: criblcontrolplanesdkgo.Pointer("<value>"),
            Timeout: criblcontrolplanesdkgo.Pointer[float64](300),
            DisableTimeFilter: criblcontrolplanesdkgo.Pointer(true),
            AuthType: operations.CreateInputAuthenticationMethodOffice365MsgTraceOauth.ToPointer(),
            RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
            MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1),
            LogLevel: operations.CreateInputLogLevelOffice365MsgTraceInfo.ToPointer(),
            JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](3),
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            RetryRules: &components.RetryRulesType1{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesBackoff,
                Interval: criblcontrolplanesdkgo.Pointer[float64](1000),
                Limit: criblcontrolplanesdkgo.Pointer[float64](5),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](2),
                Codes: []float64{
                    3152.61,
                    1476.48,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            Description: criblcontrolplanesdkgo.Pointer("mould schnitzel regarding yawningly"),
            Username: criblcontrolplanesdkgo.Pointer("Gennaro86"),
            Password: criblcontrolplanesdkgo.Pointer("pKemi5kWxN5TmzT"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            ClientID: criblcontrolplanesdkgo.Pointer("<id>"),
            Resource: criblcontrolplanesdkgo.Pointer("https://outlook.office365.com"),
            PlanType: components.SubscriptionPlanOptionsEnterpriseGcc.ToPointer(),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            CertOptions: &operations.CreateInputCertOptions{
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: "<value>",
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: "<value>",
            },
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://parched-iridescence.org"),
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
        operations.CreateInputInputOffice365Service{
            ID: "office365-service-source",
            Type: operations.CreateInputTypeOffice365ServiceOffice365Service,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            PlanType: components.SubscriptionPlanOptionsEnterpriseGcc.ToPointer(),
            TenantID: "tenant-id",
            AppID: "app-id",
            Timeout: criblcontrolplanesdkgo.Pointer[float64](300),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](30),
            JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](3),
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ContentConfig: []operations.CreateInputContentConfigOffice365Service{
                operations.CreateInputContentConfigOffice365Service{
                    ContentType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Description: criblcontrolplanesdkgo.Pointer("sew queasily interior seriously"),
                    Interval: criblcontrolplanesdkgo.Pointer[float64](2838.92),
                    LogLevel: components.LogLevelOptionsContentConfigItemsError.ToPointer(),
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                },
            },
            RetryRules: &components.RetryRulesType1{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesBackoff,
                Interval: criblcontrolplanesdkgo.Pointer[float64](1000),
                Limit: criblcontrolplanesdkgo.Pointer[float64](5),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](2),
                Codes: []float64{
                    3152.61,
                    1476.48,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            AuthType: components.AuthenticationMethodOptions1Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("outside bah inside"),
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
        operations.CreateInputInputOpenTelemetry{
            ID: "otel-source",
            Type: operations.CreateInputTypeOpenTelemetryOpenTelemetry,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 4317,
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: "<value>",
            CaptureHeaders: "<value>",
            ActivityLogSampleRate: "<value>",
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](15),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            Protocol: operations.CreateInputProtocolGrpc.ToPointer(),
            ExtractSpans: criblcontrolplanesdkgo.Pointer(false),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(false),
            OtlpVersion: operations.CreateInputOTLPVersionZeroDot10Dot0.ToPointer(),
            AuthType: components.AuthenticationTypeOptionsNone.ToPointer(),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](1000),
            Description: criblcontrolplanesdkgo.Pointer("indeed circa because"),
            Username: criblcontrolplanesdkgo.Pointer("Myles.Lueilwitz-Ryan"),
            Password: criblcontrolplanesdkgo.Pointer("RyYhbE0tvmVGPg6"),
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
        operations.CreateInputInputPrometheus{
            ID: "prometheus-source",
            Type: operations.CreateInputTypePrometheusPrometheus,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            DimensionList: []string{
                "<value 1>",
                "<value 2>",
            },
            DiscoveryType: operations.CreateInputDiscoveryTypePrometheusStatic.ToPointer(),
            Interval: 60,
            LogLevel: operations.CreateInputLogLevelPrometheusInfo,
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            Timeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](30),
            JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](3),
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("barring across fat who astride lowball"),
            TargetList: []string{
                "http://localhost:9090/metrics",
            },
            RecordType: components.RecordTypeOptionsSrv.ToPointer(),
            ScrapePort: criblcontrolplanesdkgo.Pointer[float64](9090),
            NameList: []string{
                "<value 1>",
            },
            ScrapeProtocol: operations.CreateInputMetricsProtocolHTTP.ToPointer(),
            ScrapePath: criblcontrolplanesdkgo.Pointer("/metrics"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            UsePublicIP: criblcontrolplanesdkgo.Pointer(true),
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
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            Username: criblcontrolplanesdkgo.Pointer("Edgardo39"),
            Password: criblcontrolplanesdkgo.Pointer("BvIzohhpUOO9Ez7"),
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
        operations.CreateInputInputPrometheusRw{
            ID: "prometheus-rw-source",
            Type: operations.CreateInputTypePrometheusRwPrometheusRw,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10080,
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            PrometheusAPI: "/write",
            AuthType: components.AuthenticationTypeOptionsPrometheusAuthNone.ToPointer(),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("openly likewise gruesome whoever whose gee"),
            Username: criblcontrolplanesdkgo.Pointer("Dean.Keebler"),
            Password: criblcontrolplanesdkgo.Pointer("0POH0dLKHeOMaH5"),
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
        operations.CreateInputInputRawUDP{
            ID: "raw-udp-source",
            Type: operations.CreateInputTypeRawUDPRawUDP,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 514,
            MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            SingleMsgUDPPackets: criblcontrolplanesdkgo.Pointer(false),
            IngestRawBytes: criblcontrolplanesdkgo.Pointer(false),
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](343.08),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("or robust over impact instead round near since joshingly except"),
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
        operations.CreateInputInputS3{
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "s3-notifications-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("/.*/"),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](1),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](600),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](1),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](300),
            SkipOnError: criblcontrolplanesdkgo.Pointer(false),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
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
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](5),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](600),
            Checkpointing: &components.CheckpointingType{
                Enabled: false,
                Retries: criblcontrolplanesdkgo.Pointer[float64](5),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](10),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            TagAfterProcessing: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("into keenly towards"),
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
        operations.CreateInputInputS3Inventory{
            ID: "s3-inventory-source",
            Type: operations.CreateInputTypeS3InventoryS3Inventory,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "s3-inventory-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("/.*/"),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](1),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](600),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](1),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](300),
            SkipOnError: criblcontrolplanesdkgo.Pointer(false),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
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
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](5),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](600),
            Checkpointing: &components.CheckpointingType{
                Enabled: false,
                Retries: criblcontrolplanesdkgo.Pointer[float64](5),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](10),
            ChecksumSuffix: criblcontrolplanesdkgo.Pointer("checksum"),
            MaxManifestSizeKB: criblcontrolplanesdkgo.Pointer[int64](4096),
            ValidateInventoryFiles: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("finally divert if lively seriously"),
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
        operations.CreateInputInputSecurityLake{
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "security-lake-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("/.*/"),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](1),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](600),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](1),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](300),
            SkipOnError: criblcontrolplanesdkgo.Pointer(false),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
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
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](5),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](600),
            Checkpointing: &components.CheckpointingType{
                Enabled: false,
                Retries: criblcontrolplanesdkgo.Pointer[float64](5),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](10),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("mmm lest pfft likewise rule what comparison amidst median gullible"),
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
        operations.CreateInputInputSnmp{
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "192.168.1.1",
            Port: 161,
            SnmpV3Auth: &operations.CreateInputSNMPv3Authentication{
                V3AuthEnabled: false,
                AllowUnmatchedTrap: criblcontrolplanesdkgo.Pointer(false),
                V3Users: []operations.CreateInputV3User{
                    operations.CreateInputV3User{
                        Name: "<value>",
                        AuthProtocol: components.AuthenticationProtocolOptionsV3UserNone.ToPointer(),
                        AuthKey: criblcontrolplanesdkgo.Pointer("<value>"),
                        PrivProtocol: operations.CreateInputPrivacyProtocolNone.ToPointer(),
                        PrivKey: criblcontrolplanesdkgo.Pointer("<value>"),
                    },
                },
            },
            MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](4917.68),
            VarbindsWithTypes: criblcontrolplanesdkgo.Pointer(false),
            BestEffortParsing: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("um pinion wash however meanwhile yesterday separately zowie"),
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
        operations.CreateInputInputSplunk{
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 9997,
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](1000),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](30),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            AuthTokens: []operations.CreateInputAuthTokenSplunk{
                operations.CreateInputAuthTokenSplunk{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("so obedience aw waft incidentally brr responsibility even furthermore kiss"),
                },
            },
            MaxS2Sversion: operations.CreateInputMaxS2SVersionV3.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("why event apropos row paralyse upbeat amidst pish joyously ack"),
            UseFwdTimezone: criblcontrolplanesdkgo.Pointer(true),
            DropControlFields: criblcontrolplanesdkgo.Pointer(true),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(false),
            Compress: operations.CreateInputCompressionDisabled.ToPointer(),
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
        operations.CreateInputInputSplunkHec{
            ID: "splunk-hec-source",
            Type: operations.CreateInputTypeSplunkHecSplunkHec,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 8088,
            AuthTokens: []operations.CreateInputAuthTokenSplunkHec{
                operations.CreateInputAuthTokenSplunkHec{
                    AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
                    TokenSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Token: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("meh earnest numeracy sardonic stack about"),
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
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: "<value>",
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            SplunkHecAPI: "/services/collector",
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AllowedIndexes: []string{
                "<value 1>",
            },
            SplunkHecAcks: criblcontrolplanesdkgo.Pointer(false),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            UseFwdTimezone: criblcontrolplanesdkgo.Pointer(true),
            DropControlFields: criblcontrolplanesdkgo.Pointer(true),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(false),
            AccessControlAllowOrigin: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            AccessControlAllowHeaders: []string{
                "<value 1>",
            },
            EmitTokenMetrics: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("executor given wheel evenly scowl arid boulevard whoa yuck"),
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
        operations.CreateInputInputSplunkSearch{
            ID: "splunk-search-source",
            Type: operations.CreateInputTypeSplunkSearchSplunkSearch,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            SearchHead: "https://localhost:8089",
            Search: "index=main",
            Earliest: criblcontrolplanesdkgo.Pointer("-16m@m"),
            Latest: criblcontrolplanesdkgo.Pointer("-1m@m"),
            CronSchedule: "0 * * * *",
            Endpoint: "/services/search/jobs/export",
            OutputMode: components.OutputModeOptionsSplunkCollectorConfJSON,
            EndpointParams: []operations.CreateInputEndpointParam{
                operations.CreateInputEndpointParam{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            EndpointHeaders: []operations.CreateInputEndpointHeader{
                operations.CreateInputEndpointHeader{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            LogLevel: operations.CreateInputLogLevelSplunkSearchError.ToPointer(),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](30),
            JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](3),
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            RetryRules: &components.RetryRulesType{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesBackoff,
                Interval: criblcontrolplanesdkgo.Pointer[float64](1000),
                Limit: criblcontrolplanesdkgo.Pointer[float64](5),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](2),
                Codes: []float64{
                    5227.09,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            AuthType: operations.CreateInputAuthenticationTypeSplunkSearchBasic.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("until content along"),
            Username: criblcontrolplanesdkgo.Pointer("Jalyn43"),
            Password: criblcontrolplanesdkgo.Pointer("2wxXTLiflvFCcIX"),
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
        operations.CreateInputInputSqs{
            ID: "sqs-source",
            Type: operations.CreateInputTypeSqsSqs,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "my-queue",
            QueueType: operations.CreateInputQueueTypeStandard,
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            CreateQueue: criblcontrolplanesdkgo.Pointer(false),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions3V4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](10),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](600),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](10),
            Description: criblcontrolplanesdkgo.Pointer("anxiously wilted gee once icebreaker unkempt"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](3),
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
        operations.CreateCreateInputInputSyslogUnionCreateInputInputSyslogSyslog1(
            operations.CreateInputInputSyslogSyslog1{
                ID: "syslog-source",
                Type: operations.CreateInputInputSyslogType1Syslog,
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
                    MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                    CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                    MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                    MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                    Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                    Compress: components.CompressionOptionsPqNone.ToPointer(),
                    PqControls: &components.PqTypePqControls{},
                },
                Host: "0.0.0.0",
                UDPPort: 514,
                TCPPort: criblcontrolplanesdkgo.Pointer[float64](2766.66),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                TimestampTimezone: criblcontrolplanesdkgo.Pointer("local"),
                SingleMsgUDPPackets: criblcontrolplanesdkgo.Pointer(false),
                EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
                KeepFieldsList: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
                OctetCounting: criblcontrolplanesdkgo.Pointer(false),
                InferFraming: criblcontrolplanesdkgo.Pointer(true),
                StrictlyInferOctetCounting: criblcontrolplanesdkgo.Pointer(true),
                AllowNonStandardAppName: criblcontrolplanesdkgo.Pointer(false),
                MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](1000),
                SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
                SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](30),
                SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](0),
                TLS: &components.TLSSettingsServerSideType{
                    Disabled: criblcontrolplanesdkgo.Pointer(true),
                    RequestCert: criblcontrolplanesdkgo.Pointer(false),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                    CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                    CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                    PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                },
                Metadata: []components.ItemsTypeMetadata{
                    components.ItemsTypeMetadata{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](2673.89),
                EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(false),
                Description: criblcontrolplanesdkgo.Pointer("enthusiastically idolized now"),
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
        operations.CreateInputInputSystemMetrics{
            ID: "system-metrics-source",
            Type: operations.CreateInputTypeSystemMetricsSystemMetrics,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](10),
            Host: &operations.CreateInputHostSystemMetrics{
                Mode: components.ModeOptionsHostBasic.ToPointer(),
                Custom: &operations.CreateInputCustomSystemMetrics{
                    System: &operations.CreateInputSystemSystemMetrics{
                        Mode: operations.CreateInputSystemModeSystemMetricsBasic.ToPointer(),
                        Processes: criblcontrolplanesdkgo.Pointer(false),
                    },
                    CPU: &operations.CreateInputCPUSystemMetrics{
                        Mode: operations.CreateInputCPUModeSystemMetricsBasic.ToPointer(),
                        PerCPU: criblcontrolplanesdkgo.Pointer(false),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Time: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Memory: &operations.CreateInputMemorySystemMetrics{
                        Mode: operations.CreateInputMemoryModeSystemMetricsBasic.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Network: &operations.CreateInputNetworkSystemMetrics{
                        Mode: operations.CreateInputNetworkModeSystemMetricsBasic.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Protocols: criblcontrolplanesdkgo.Pointer(false),
                        Devices: []string{
                            "<value 1>",
                            "<value 2>",
                            "<value 3>",
                        },
                        PerInterface: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Disk: &operations.CreateInputDiskSystemMetrics{
                        Mode: operations.CreateInputDiskModeSystemMetricsBasic.ToPointer(),
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
            Container: &operations.CreateInputContainer{
                Mode: operations.CreateInputContainerModeBasic.ToPointer(),
                DockerSocket: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
                DockerTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
                Filters: []operations.CreateInputContainerFilter{
                    operations.CreateInputContainerFilter{
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
            Persistence: &operations.CreateInputPersistenceSystemMetrics{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("10m"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("1GB"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("24h"),
                Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/system_metrics"),
            },
            Description: criblcontrolplanesdkgo.Pointer("although jealously forswear clamor"),
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
        operations.CreateInputInputSystemState{
            ID: "system-state-source",
            Type: operations.CreateInputTypeSystemStateSystemState,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](300),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Collectors: &operations.CreateInputCollectors{
                Hostsfile: &operations.CreateInputHostsFile{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Interfaces: &operations.CreateInputInterfaces{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Disk: &operations.CreateInputDisksAndFileSystems{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Metadata: &operations.CreateInputHostInfo{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Routes: &operations.CreateInputRoutes{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                DNS: &operations.CreateInputDNS{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                User: &operations.CreateInputUsersAndGroups{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Firewall: &operations.CreateInputFirewall{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Services: &operations.CreateInputServices{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Ports: &operations.CreateInputListeningPorts{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                LoginUsers: &operations.CreateInputLoggedInUsers{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
            },
            Persistence: &operations.CreateInputPersistenceSystemState{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("10m"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("1GB"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("24h"),
                Compress: components.DataCompressionFormatOptionsPersistenceNone.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/system_state"),
            },
            DisableNativeModule: criblcontrolplanesdkgo.Pointer(false),
            DisableNativeLastLogModule: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("times even descendant absent although mortally"),
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
        operations.CreateInputInputTCP{
            ID: "tcp-source",
            Type: operations.CreateInputTypeTCPTCP,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10090,
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](1000),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](30),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            EnableHeader: criblcontrolplanesdkgo.Pointer(false),
            Preprocess: &components.PreprocessType{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                    "<value 2>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("clumsy eternity than save both hover"),
            AuthToken: criblcontrolplanesdkgo.Pointer(""),
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
        operations.CreateInputInputTcpjson{
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
            },
            Connections: []components.ItemsTypeConnectionsOptional{
                components.ItemsTypeConnectionsOptional{
                    Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                    Output: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Pq: &components.PqType{
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10090,
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](1000),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](30),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(false),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("merrily scaly unless"),
            AuthToken: criblcontrolplanesdkgo.Pointer(""),
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
        operations.CreateInputInputWef{
            ID: "wef-source",
            Type: operations.CreateInputTypeWefWef,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 5985,
            AuthMethod: operations.CreateInputAuthMethodAuthenticationMethodClientCert.ToPointer(),
            TLS: &operations.CreateInputMTLSSettings{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(true),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: "<value>",
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: "<value>",
                CaPath: "<value>",
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                OcspCheck: criblcontrolplanesdkgo.Pointer(false),
                Keytab: "<value>",
                Principal: "<value>",
                OcspCheckFailClose: criblcontrolplanesdkgo.Pointer(false),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](90),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            CaFingerprint: criblcontrolplanesdkgo.Pointer("<value>"),
            Keytab: criblcontrolplanesdkgo.Pointer("<value>"),
            Principal: criblcontrolplanesdkgo.Pointer("<value>"),
            AllowMachineIDMismatch: criblcontrolplanesdkgo.Pointer(false),
            Subscriptions: []operations.CreateInputSubscription{
                operations.CreateInputSubscription{
                    SubscriptionName: "subscription-1",
                    Version: criblcontrolplanesdkgo.Pointer("<value>"),
                    ContentFormat: operations.CreateInputFormatRenderedText,
                    HeartbeatInterval: 60,
                    BatchTimeout: 5,
                    ReadExistingEvents: criblcontrolplanesdkgo.Pointer(false),
                    SendBookmarks: criblcontrolplanesdkgo.Pointer(true),
                    Compress: criblcontrolplanesdkgo.Pointer(true),
                    Targets: []string{},
                    Locale: criblcontrolplanesdkgo.Pointer("en-US"),
                    QuerySelector: operations.CreateInputQueryBuilderModeSimple.ToPointer(),
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                    Queries: []operations.CreateInputQuery{
                        operations.CreateInputQuery{
                            Path: "/usr/src",
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
            Description: criblcontrolplanesdkgo.Pointer("before whenever circa geez youthfully lest off"),
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
        operations.CreateInputInputWinEventLogs{
            ID: "win-event-logs-source",
            Type: operations.CreateInputTypeWinEventLogsWinEventLogs,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            LogNames: []string{
                "Application",
                "System",
            },
            ReadMode: operations.CreateInputReadModeNewest.ToPointer(),
            EventFormat: operations.CreateInputEventFormatJSON.ToPointer(),
            DisableNativeModule: criblcontrolplanesdkgo.Pointer(false),
            Interval: criblcontrolplanesdkgo.Pointer[float64](10),
            BatchSize: criblcontrolplanesdkgo.Pointer[float64](500),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            MaxEventBytes: criblcontrolplanesdkgo.Pointer[float64](51200),
            Description: criblcontrolplanesdkgo.Pointer("conclude sesame prioritize polarisation bourgeoisie decongestant behind"),
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
        operations.CreateInputInputWindowsMetrics{
            ID: "windows-metrics-source",
            Type: operations.CreateInputTypeWindowsMetricsWindowsMetrics,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](10),
            Host: &operations.CreateInputHostWindowsMetrics{
                Mode: components.ModeOptionsHostBasic.ToPointer(),
                Custom: &operations.CreateInputCustomWindowsMetrics{
                    System: &operations.CreateInputSystemWindowsMetrics{
                        Mode: operations.CreateInputSystemModeWindowsMetricsBasic.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                    },
                    CPU: &operations.CreateInputCPUWindowsMetrics{
                        Mode: operations.CreateInputCPUModeWindowsMetricsBasic.ToPointer(),
                        PerCPU: criblcontrolplanesdkgo.Pointer(false),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Time: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Memory: &operations.CreateInputMemoryWindowsMetrics{
                        Mode: operations.CreateInputMemoryModeWindowsMetricsBasic.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Network: &operations.CreateInputNetworkWindowsMetrics{
                        Mode: operations.CreateInputNetworkModeWindowsMetricsBasic.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Protocols: criblcontrolplanesdkgo.Pointer(false),
                        Devices: []string{
                            "<value 1>",
                            "<value 2>",
                        },
                        PerInterface: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Disk: &operations.CreateInputDiskWindowsMetrics{
                        Mode: operations.CreateInputDiskModeWindowsMetricsBasic.ToPointer(),
                        PerVolume: criblcontrolplanesdkgo.Pointer(false),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Volumes: []string{
                            "<value 1>",
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
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Persistence: &operations.CreateInputPersistenceWindowsMetrics{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("10m"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("1GB"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("24h"),
                Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/windows_metrics"),
            },
            DisableNativeModule: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("handle unsung mediocre delightfully save grown whose circa"),
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
        operations.CreateInputInputWiz{
            ID: "wiz-source",
            Type: operations.CreateInputTypeWizWiz,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Endpoint: "https://api.wiz.io",
            AuthURL: "https://auth.wiz.io/oauth/token",
            AuthAudienceOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            ClientID: "client-id",
            ContentConfig: []operations.CreateInputContentConfigWiz{},
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](300),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](3),
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            RetryRules: &components.RetryRulesType{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesBackoff,
                Interval: criblcontrolplanesdkgo.Pointer[float64](1000),
                Limit: criblcontrolplanesdkgo.Pointer[float64](5),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](2),
                Codes: []float64{
                    5227.09,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            AuthType: components.AuthenticationMethodOptions1Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("quarrelsomely furthermore deserted excitable liberalize"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateEndpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAuthURL: criblcontrolplanesdkgo.Pointer("https://misguided-pine.org"),
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
        operations.CreateInputInputWizWebhook{
            ID: "wiz-webhook-source",
            Type: operations.CreateInputTypeWizWebhookWizWebhook,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
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
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AllowedPaths: []string{
                "<value 1>",
                "<value 2>",
            },
            AllowedMethods: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            AuthTokensExt: []components.ItemsTypeAuthTokensExt{
                components.ItemsTypeAuthTokensExt{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("swath misfire object saw presume yum"),
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("above anti underneath geez wing fuel untidy verbally stir-fry lasting"),
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
        operations.CreateInputInputZscalerHec{
            ID: "zscaler-hec-source",
            Type: operations.CreateInputTypeZscalerHecZscalerHec,
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 8088,
            AuthTokens: []operations.CreateInputAuthTokenZscalerHec{
                operations.CreateInputAuthTokenZscalerHec{
                    AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
                    TokenSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Token: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("boohoo restaurant ouch braid short"),
                    AllowedIndexesAtToken: []string{
                        "<value 1>",
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
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: "<value>",
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
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
            HecAcks: criblcontrolplanesdkgo.Pointer(false),
            AccessControlAllowOrigin: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            AccessControlAllowHeaders: []string{
                "<value 1>",
            },
            EmitTokenMetrics: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("vastly against stable down pfft likely"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](1000),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](30),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](0),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            EnableUnixPath: criblcontrolplanesdkgo.Pointer(false),
            Filter: &components.InputAppscopeFilter{
                Allow: []components.Allow{
                    components.Allow{
                        Procname: "<value>",
                        Arg: criblcontrolplanesdkgo.Pointer("<value>"),
                        Config: "<value>",
                    },
                },
                TransportURL: criblcontrolplanesdkgo.Pointer("https://youthful-hammock.net/"),
            },
            Persistence: &components.InputAppscopePersistence{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("10m"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("1GB"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("24h"),
                Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/appscope"),
            },
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("beyond hidden supposing ghost fictionalize disarm geez"),
            Host: criblcontrolplanesdkgo.Pointer("0.0.0.0"),
            Port: criblcontrolplanesdkgo.Pointer[float64](9109),
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            UnixSocketPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/appscope.sock"),
            UnixSocketPerms: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthToken: criblcontrolplanesdkgo.Pointer(""),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "azure-blob-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("/.*/"),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](600),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](1),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](1),
            ServicePeriodSecs: criblcontrolplanesdkgo.Pointer[float64](5),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](5),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](600),
            AuthType: components.AuthenticationMethodOptionsManual.ToPointer(),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 8088,
            AuthTokens: []components.InputCloudflareHecAuthToken{
                components.InputCloudflareHecAuthToken{
                    AuthType: components.InputCloudflareHecAuthenticationMethodSecret.ToPointer(),
                    TokenSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Token: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("ha irresponsible patiently"),
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
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: "<value>",
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
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
            Description: criblcontrolplanesdkgo.Pointer("abaft deliberately function"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            BreakerRulesets: []string{
                "<value 1>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            Preprocess: &components.PreprocessType{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                },
            },
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Brokers: []string{
                "pkc-xxxxx.us-east-1.aws.confluent.cloud:9092",
            },
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
            Topics: []string{
                "logs",
            },
            GroupID: criblcontrolplanesdkgo.Pointer("Cribl"),
            FromBeginning: criblcontrolplanesdkgo.Pointer(true),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType{
                Disabled: true,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("http://localhost:8081"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](1),
                Auth: &components.AuthTypeKafkaSchemaRegistry{
                    Disabled: true,
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
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
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](30000),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](300),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](10000),
            Sasl: &components.AuthenticationType{
                Disabled: true,
                Username: criblcontrolplanesdkgo.Pointer("Maximus30"),
                Password: criblcontrolplanesdkgo.Pointer("6TQ9JUSPWUKvYWA"),
                AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslPlain.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(false),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://sudden-polarisation.info/"),
                ClientID: criblcontrolplanesdkgo.Pointer("<id>"),
                OauthSecretType: criblcontrolplanesdkgo.Pointer("secret"),
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
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](3000),
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](443.17),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](3647.04),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](1048576),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](10485760),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](0),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("or since gadzooks"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10080,
            AuthTokens: []components.ItemsTypeAuthTokens{
                components.ItemsTypeAuthTokens{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("but capitalise clear unfortunate ignorance gah sans despite hydrocarbon tankful"),
                },
            },
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("ride even among"),
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
            },
            Connections: []components.ItemsTypeConnectionsOptional{
                components.ItemsTypeConnectionsOptional{
                    Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                    Output: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Pq: &components.PqType{
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
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
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            CriblAPI: criblcontrolplanesdkgo.Pointer("/cribl"),
            ElasticAPI: criblcontrolplanesdkgo.Pointer("/elastic"),
            SplunkHecAPI: criblcontrolplanesdkgo.Pointer("/services/collector"),
            SplunkHecAcks: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthTokensExt: []components.AuthTokensExt{
                components.AuthTokensExt{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("extract vice instead fatal every rebuke ew forenenst millet crest"),
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
                        Enabled: criblcontrolplanesdkgo.Pointer(true),
                        DefaultDataset: criblcontrolplanesdkgo.Pointer("<value>"),
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("authentic yuck better as where splash but behind"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10090,
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](1000),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](30),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
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
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("delightfully scarily chromakey oh hm sour"),
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("refer redevelop mid reproach waterspout lest utterly"),
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "crowdstrike-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("/.*/"),
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
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](1),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](21600),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](1),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](300),
            SkipOnError: criblcontrolplanesdkgo.Pointer(false),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            EnableSQSAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            Preprocess: &components.PreprocessType{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
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
                Retries: criblcontrolplanesdkgo.Pointer[float64](5),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](10),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("corporation joyously retention only transcend save however likely shanghai"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 8126,
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ProxyMode: &components.InputDatadogAgentProxyMode{
                Enabled: false,
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            },
            Description: criblcontrolplanesdkgo.Pointer("made-up solace bouncy which comfortable gadzooks"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
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
            Description: criblcontrolplanesdkgo.Pointer("aha reflecting now"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            DimensionList: []string{
                "<value 1>",
            },
            DiscoveryType: components.InputEdgePrometheusDiscoveryTypeStatic,
            Interval: 60,
            Timeout: criblcontrolplanesdkgo.Pointer[float64](5000),
            Persistence: &components.DiskSpoolingType{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("10m"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("1GB"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("24h"),
                Compress: components.CompressionOptionsPersistenceGzip.ToPointer(),
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthType: components.InputEdgePrometheusAuthenticationMethodManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("hmph jagged minus fondly affect cinch prioritize"),
            Targets: []components.Target{
                components.Target{
                    Protocol: components.ProtocolOptionsTargetsItemsHTTP.ToPointer(),
                    Host: "localhost",
                    Port: criblcontrolplanesdkgo.Pointer[float64](9090),
                    Path: criblcontrolplanesdkgo.Pointer("/metrics"),
                },
            },
            RecordType: components.RecordTypeOptionsSrv.ToPointer(),
            ScrapePort: criblcontrolplanesdkgo.Pointer[float64](9090),
            NameList: []string{
                "<value 1>",
                "<value 2>",
            },
            ScrapeProtocol: components.ProtocolOptionsTargetsItemsHTTP.ToPointer(),
            ScrapePath: criblcontrolplanesdkgo.Pointer("/metrics"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
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
            SignatureVersion: components.SignatureVersionOptions1V4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            ScrapeProtocolExpr: criblcontrolplanesdkgo.Pointer("metadata.annotations['prometheus.io/scheme'] || 'http'"),
            ScrapePortExpr: criblcontrolplanesdkgo.Pointer("metadata.annotations['prometheus.io/port'] || 9090"),
            ScrapePathExpr: criblcontrolplanesdkgo.Pointer("metadata.annotations['prometheus.io/path'] || '/metrics'"),
            PodFilter: []components.PodFilter{
                components.PodFilter{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("severe furthermore concrete hunger muffled"),
                },
            },
            Username: criblcontrolplanesdkgo.Pointer("Orrin_Yundt20"),
            Password: criblcontrolplanesdkgo.Pointer("alXPvLywpTQbHyi"),
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "localhost",
            Port: 9200,
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            ElasticAPI: "/",
            AuthType: components.InputElasticAuthenticationTypeNone.ToPointer(),
            APIVersion: components.InputElasticAPIVersionEightDot3Dot2.ToPointer(),
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
                Enabled: false,
                AuthType: components.InputElasticAuthenticationMethodNone.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("Abigayle.Ledner38"),
                Password: criblcontrolplanesdkgo.Pointer("VDPV4GrdklqLt4A"),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                URL: criblcontrolplanesdkgo.Pointer("https://yellowish-pearl.org"),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                RemoveHeaders: []string{
                    "<value 1>",
                    "<value 2>",
                },
                TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](60),
                TemplateURL: criblcontrolplanesdkgo.Pointer("https://focused-invite.org"),
            },
            Description: criblcontrolplanesdkgo.Pointer("ouch snarling duh"),
            Username: criblcontrolplanesdkgo.Pointer("Xander_Graham"),
            Password: criblcontrolplanesdkgo.Pointer("1duIyiiwRASxBVc"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthTokens: []string{
                "<value 1>",
            },
            CustomAPIVersion: criblcontrolplanesdkgo.Pointer("{\n    \"name\": \"AzU84iL\",\n    \"cluster_name\": \"cribl\",\n    \"cluster_uuid\": \"Js6_Z2VKS3KbfRSxPmPbaw\",\n    \"version\": {\n        \"number\": \"8.3.2\",\n        \"build_type\": \"tar\",\n        \"build_hash\": \"bca0c8d\",\n        \"build_date\": \"2019-10-16T06:19:49.319352Z\",\n        \"build_snapshot\": false,\n        \"lucene_version\": \"9.7.2\",\n        \"minimum_wire_compatibility_version\": \"7.17.0\",\n        \"minimum_index_compatibility_version\": \"7.0.0\"\n    },\n    \"tagline\": \"You Know, for Search\"\n}"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Brokers: []string{
                "myeventhub.servicebus.windows.net:9093",
            },
            Topics: []string{
                "logs",
            },
            GroupID: criblcontrolplanesdkgo.Pointer("Cribl"),
            FromBeginning: criblcontrolplanesdkgo.Pointer(true),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](30000),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](300),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](10000),
            Sasl: &components.AuthenticationType1{
                Disabled: false,
                AuthType: components.AuthenticationMethodOptionsSasl1Manual.ToPointer(),
                Password: criblcontrolplanesdkgo.Pointer("xop7OD_LxhffeRA"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSasl1Plain.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("$ConnectionString"),
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
                Disabled: false,
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            },
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](3000),
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](4392.98),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](3134.99),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](1048576),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](10485760),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](0),
            MinimizeDuplicates: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("beard hoot rotten than unfortunately slushy"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Command: "echo \"Hello World\"",
            Retries: criblcontrolplanesdkgo.Pointer[float64](10),
            ScheduleType: components.ScheduleTypeInterval.ToPointer(),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("er relaunch rudely bug when"),
            Interval: criblcontrolplanesdkgo.Pointer[float64](60),
            CronSchedule: criblcontrolplanesdkgo.Pointer("* * * * *"),
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
            },
            Connections: []components.ItemsTypeConnectionsOptional{
                components.ItemsTypeConnectionsOptional{
                    Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                    Output: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Pq: &components.PqType{
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Mode: components.InputFileModeManual.ToPointer(),
            Interval: criblcontrolplanesdkgo.Pointer[float64](10),
            Filenames: []string{
                "<value 1>",
                "<value 2>",
            },
            FilterArchivedFiles: criblcontrolplanesdkgo.Pointer(false),
            TailOnly: criblcontrolplanesdkgo.Pointer(true),
            IdleTimeout: criblcontrolplanesdkgo.Pointer[float64](300),
            MinAgeDur: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxAgeDur: criblcontrolplanesdkgo.Pointer("<value>"),
            CheckFileModTime: criblcontrolplanesdkgo.Pointer(false),
            ForceText: criblcontrolplanesdkgo.Pointer(false),
            HashLen: criblcontrolplanesdkgo.Pointer[float64](256),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            Description: criblcontrolplanesdkgo.Pointer("shell yesterday testimonial afore though sedately profuse"),
            Path: criblcontrolplanesdkgo.Pointer("/Network"),
            Depth: criblcontrolplanesdkgo.Pointer[float64](4361.93),
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
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
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("fluffy now to"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
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
            MaxBacklog: criblcontrolplanesdkgo.Pointer[float64](1000),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("across splosh anticodon beautifully"),
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
                    Mode: components.ModeOptionsPqAlways.ToPointer(),
                    MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                    CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                    MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                    MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                    Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                    Compress: components.CompressionOptionsPqNone.ToPointer(),
                    PqControls: &components.PqTypePqControls{},
                },
                Host: "0.0.0.0",
                Port: 10080,
                TLS: &components.TLSSettingsServerSideType{
                    Disabled: criblcontrolplanesdkgo.Pointer(true),
                    RequestCert: criblcontrolplanesdkgo.Pointer(false),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                    CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                    CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                    PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                },
                MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
                MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
                EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
                CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
                ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
                SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
                KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
                EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
                IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
                PrometheusAPI: "/api/prom/push",
                LokiAPI: criblcontrolplanesdkgo.Pointer("/loki/api/v1/push"),
                PrometheusAuth: &components.PrometheusAuth1{
                    AuthType: components.AuthenticationTypeOptionsPrometheusAuthNone.ToPointer(),
                    Username: criblcontrolplanesdkgo.Pointer("Gage.Rippin77"),
                    Password: criblcontrolplanesdkgo.Pointer("B6pdaXMgVOmyEDo"),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                LokiAuth: &components.LokiAuth1{
                    AuthType: components.AuthenticationTypeOptionsLokiAuthNone.ToPointer(),
                    Username: criblcontrolplanesdkgo.Pointer("Zion18"),
                    Password: criblcontrolplanesdkgo.Pointer("NTgCCy2le8of9ZE"),
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
                Description: criblcontrolplanesdkgo.Pointer("ironclad meh to merrily besmirch whoa slimy"),
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
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
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            CriblAPI: criblcontrolplanesdkgo.Pointer("/cribl"),
            ElasticAPI: criblcontrolplanesdkgo.Pointer("/elastic"),
            SplunkHecAPI: criblcontrolplanesdkgo.Pointer("/services/collector"),
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
                    Description: criblcontrolplanesdkgo.Pointer("gee qua gee afterwards aboard begonia um where absent geez"),
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("via supportive so dial into valuable instructive of between"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
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
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AllowedPaths: []string{
                "<value 1>",
                "<value 2>",
            },
            AllowedMethods: []string{
                "<value 1>",
            },
            AuthTokensExt: []components.ItemsTypeAuthTokensExt{
                components.ItemsTypeAuthTokensExt{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("afore gah odd alive pant overload mob rise throughout"),
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("ultimately needily whether hope quinoa gadzooks above where hierarchy"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Path: "/var/log/journal",
            Interval: criblcontrolplanesdkgo.Pointer[float64](10),
            Journals: []string{
                "system",
            },
            Rules: []components.InputJournalFilesRule{
                components.InputJournalFilesRule{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("fedora until anguished jump famously beside consequently tender decision"),
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
            Description: criblcontrolplanesdkgo.Pointer("suffocate vamoose fortunately than kaleidoscopic"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Brokers: []string{
                "localhost:9092",
            },
            Topics: []string{
                "logs",
            },
            GroupID: criblcontrolplanesdkgo.Pointer("Cribl"),
            FromBeginning: criblcontrolplanesdkgo.Pointer(true),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType{
                Disabled: true,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("http://localhost:8081"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](1),
                Auth: &components.AuthTypeKafkaSchemaRegistry{
                    Disabled: true,
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
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
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](30000),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](300),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](10000),
            Sasl: &components.AuthenticationType{
                Disabled: true,
                Username: criblcontrolplanesdkgo.Pointer("Maximus30"),
                Password: criblcontrolplanesdkgo.Pointer("6TQ9JUSPWUKvYWA"),
                AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslPlain.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(false),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://sudden-polarisation.info/"),
                ClientID: criblcontrolplanesdkgo.Pointer("<id>"),
                OauthSecretType: criblcontrolplanesdkgo.Pointer("secret"),
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
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](3000),
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](5373.51),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](3922.91),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](1048576),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](10485760),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](0),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("to before brr including"),
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
            },
            Connections: []components.ItemsTypeConnectionsOptional{
                components.ItemsTypeConnectionsOptional{
                    Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                    Output: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Pq: &components.PqType{
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            StreamName: "my-stream",
            ServiceInterval: criblcontrolplanesdkgo.Pointer[float64](1),
            ShardExpr: criblcontrolplanesdkgo.Pointer("true"),
            ShardIteratorType: components.ShardIteratorStartTrimHorizon.ToPointer(),
            PayloadFormat: components.RecordDataFormatCribl.ToPointer(),
            GetRecordsLimit: criblcontrolplanesdkgo.Pointer[float64](5000),
            GetRecordsLimitTotal: criblcontrolplanesdkgo.Pointer[float64](20000),
            LoadBalancingAlgorithm: components.ShardLoadBalancingConsistentHashing.ToPointer(),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions2V4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            VerifyKPLCheckSums: criblcontrolplanesdkgo.Pointer(false),
            AvoidDuplicates: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("factorize than bah trained shinny regarding fooey"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Rules: []components.ItemsTypeRules{
                components.ItemsTypeRules{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("worse back second psst overspend prejudge homely gah rim"),
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("concerning know overconfidently carelessly throughout even before"),
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](15),
            Rules: []components.InputKubeLogsRule{
                components.InputKubeLogsRule{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("bravely compromise lest SUV deliberately"),
                },
            },
            Timestamps: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Persistence: &components.DiskSpoolingType{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("10m"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("1GB"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("24h"),
                Compress: components.CompressionOptionsPersistenceGzip.ToPointer(),
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("after taxicab agreeable pulverize whenever atop"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](15),
            Rules: []components.ItemsTypeRules{
                components.ItemsTypeRules{
                    Filter: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("ack fathom incidentally doing kinase"),
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
                TimeWindow: criblcontrolplanesdkgo.Pointer("10m"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("1GB"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("24h"),
                Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/kube_metrics"),
            },
            Description: criblcontrolplanesdkgo.Pointer("excepting inquisitively violin brown during frizzy aw zowie powerless"),
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10080,
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            LokiAPI: "/loki/api/v1/push",
            AuthType: components.AuthenticationTypeOptionsLokiAuthNone.ToPointer(),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("physical below elementary disclosure"),
            Username: criblcontrolplanesdkgo.Pointer("Diana.Kerluke41"),
            Password: criblcontrolplanesdkgo.Pointer("5VWIKl4RcP6DYtZ"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            UDPPort: criblcontrolplanesdkgo.Pointer[float64](8125),
            TCPPort: criblcontrolplanesdkgo.Pointer[float64](8643.34),
            MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
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
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](7392.9),
            Description: criblcontrolplanesdkgo.Pointer("reboot because clonk meh etch boo uh-huh fumigate pip coaxingly"),
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
            },
            Connections: []components.ItemsTypeConnectionsOptional{
                components.ItemsTypeConnectionsOptional{
                    Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                    Output: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Pq: &components.PqType{
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 57000,
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
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
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](1000),
            ShutdownTimeoutMs: criblcontrolplanesdkgo.Pointer[float64](5000),
            Description: criblcontrolplanesdkgo.Pointer("lotion since sheathe fooey truly"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Brokers: []string{
                "b-1.example.xxxxx.c2.kafka.us-east-1.amazonaws.com:9092",
            },
            Topics: []string{
                "logs",
            },
            GroupID: criblcontrolplanesdkgo.Pointer("Cribl"),
            FromBeginning: criblcontrolplanesdkgo.Pointer(true),
            SessionTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
            RebalanceTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            HeartbeatInterval: criblcontrolplanesdkgo.Pointer[float64](3000),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType{
                Disabled: true,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("http://localhost:8081"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](30000),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](1),
                Auth: &components.AuthTypeKafkaSchemaRegistry{
                    Disabled: true,
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
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
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](30000),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](300),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](10000),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto,
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
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
            AutoCommitInterval: criblcontrolplanesdkgo.Pointer[float64](621.07),
            AutoCommitThreshold: criblcontrolplanesdkgo.Pointer[float64](3924.6),
            MaxBytesPerPartition: criblcontrolplanesdkgo.Pointer[float64](1048576),
            MaxBytes: criblcontrolplanesdkgo.Pointer[float64](10485760),
            MaxSocketErrors: criblcontrolplanesdkgo.Pointer[float64](0),
            Description: criblcontrolplanesdkgo.Pointer("minor midst rebound left supposing ugh"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 2055,
            EnablePassThrough: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](2462.99),
            TemplateCacheMinutes: criblcontrolplanesdkgo.Pointer[float64](30),
            V5Enabled: criblcontrolplanesdkgo.Pointer(true),
            V9Enabled: criblcontrolplanesdkgo.Pointer(true),
            IpfixEnabled: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("boastfully hearten ick intend meanwhile jaywalk wallaby"),
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
            },
            Connections: []components.ItemsTypeConnectionsOptional{
                components.ItemsTypeConnectionsOptional{
                    Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                    Output: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Pq: &components.PqType{
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            PlanType: components.SubscriptionPlanOptionsEnterpriseGcc,
            TenantID: "tenant-id",
            AppID: "app-id",
            Timeout: criblcontrolplanesdkgo.Pointer[float64](300),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](30),
            JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](3),
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
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
                    Description: criblcontrolplanesdkgo.Pointer("superb consequently below"),
                    Interval: criblcontrolplanesdkgo.Pointer[float64](9014.67),
                    LogLevel: components.LogLevelOptionsContentConfigItemsInfo.ToPointer(),
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                },
            },
            IngestionLag: criblcontrolplanesdkgo.Pointer[float64](0),
            RetryRules: &components.RetryRulesType1{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesBackoff,
                Interval: criblcontrolplanesdkgo.Pointer[float64](1000),
                Limit: criblcontrolplanesdkgo.Pointer[float64](5),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](2),
                Codes: []float64{
                    2729.52,
                    6010.16,
                    5724.14,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            AuthType: components.AuthenticationMethodOptions1Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("solemnly forearm yahoo brr sweet until brown"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            URL: "https://reports.office365.com/ecp/reportingwebservice/reporting.svc/MessageTrace",
            Interval: 15,
            StartDate: criblcontrolplanesdkgo.Pointer("<value>"),
            EndDate: criblcontrolplanesdkgo.Pointer("<value>"),
            Timeout: criblcontrolplanesdkgo.Pointer[float64](300),
            DisableTimeFilter: criblcontrolplanesdkgo.Pointer(true),
            AuthType: components.InputOffice365MsgTraceAuthenticationMethodOauth.ToPointer(),
            RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
            MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1),
            LogLevel: components.InputOffice365MsgTraceLogLevelInfo.ToPointer(),
            JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](3),
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            RetryRules: &components.RetryRulesType1{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesBackoff,
                Interval: criblcontrolplanesdkgo.Pointer[float64](1000),
                Limit: criblcontrolplanesdkgo.Pointer[float64](5),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](2),
                Codes: []float64{
                    2729.52,
                    6010.16,
                    5724.14,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            Description: criblcontrolplanesdkgo.Pointer("though gadzooks pace"),
            Username: criblcontrolplanesdkgo.Pointer("Domingo.Funk20"),
            Password: criblcontrolplanesdkgo.Pointer("zEomPhRXS2K3PJN"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            ClientID: criblcontrolplanesdkgo.Pointer("<id>"),
            Resource: criblcontrolplanesdkgo.Pointer("https://outlook.office365.com"),
            PlanType: components.SubscriptionPlanOptionsEnterpriseGcc.ToPointer(),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            CertOptions: &components.CertOptions{
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: "<value>",
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: "<value>",
            },
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://valuable-stall.name/"),
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
            },
            Connections: []components.ItemsTypeConnectionsOptional{
                components.ItemsTypeConnectionsOptional{
                    Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                    Output: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Pq: &components.PqType{
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            PlanType: components.SubscriptionPlanOptionsEnterpriseGcc.ToPointer(),
            TenantID: "tenant-id",
            AppID: "app-id",
            Timeout: criblcontrolplanesdkgo.Pointer[float64](300),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](30),
            JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](3),
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ContentConfig: []components.InputOffice365ServiceContentConfig{
                components.InputOffice365ServiceContentConfig{
                    ContentType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Description: criblcontrolplanesdkgo.Pointer("actually that after hyphenation psst"),
                    Interval: criblcontrolplanesdkgo.Pointer[float64](7646.22),
                    LogLevel: components.LogLevelOptionsContentConfigItemsInfo.ToPointer(),
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                },
            },
            RetryRules: &components.RetryRulesType1{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesBackoff,
                Interval: criblcontrolplanesdkgo.Pointer[float64](1000),
                Limit: criblcontrolplanesdkgo.Pointer[float64](5),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](2),
                Codes: []float64{
                    2729.52,
                    6010.16,
                    5724.14,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            AuthType: components.AuthenticationMethodOptions1Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("or oh apropos in analogy lobotomise lender gadzooks between fully"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 4317,
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: "<value>",
            CaptureHeaders: "<value>",
            ActivityLogSampleRate: "<value>",
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](15),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            Protocol: components.InputOpenTelemetryProtocolGrpc.ToPointer(),
            ExtractSpans: criblcontrolplanesdkgo.Pointer(false),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(false),
            OtlpVersion: components.InputOpenTelemetryOTLPVersionZeroDot10Dot0.ToPointer(),
            AuthType: components.AuthenticationTypeOptionsNone.ToPointer(),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](1000),
            Description: criblcontrolplanesdkgo.Pointer("augment before serialize"),
            Username: criblcontrolplanesdkgo.Pointer("Monique.Lakin"),
            Password: criblcontrolplanesdkgo.Pointer("Ig92bVAWqo4_P8d"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
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
            Timeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](30),
            JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](3),
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("hover aw buzzing part unethically apud down scornful"),
            TargetList: []string{
                "http://localhost:9090/metrics",
            },
            RecordType: components.RecordTypeOptionsSrv.ToPointer(),
            ScrapePort: criblcontrolplanesdkgo.Pointer[float64](9090),
            NameList: []string{
                "<value 1>",
            },
            ScrapeProtocol: components.MetricsProtocolHTTP.ToPointer(),
            ScrapePath: criblcontrolplanesdkgo.Pointer("/metrics"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            UsePublicIP: criblcontrolplanesdkgo.Pointer(true),
            SearchFilter: []components.ItemsTypeSearchFilter{
                components.ItemsTypeSearchFilter{
                    Name: "<value>",
                    Values: []string{
                        "<value 1>",
                        "<value 2>",
                        "<value 3>",
                    },
                },
            },
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions1V4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            Username: criblcontrolplanesdkgo.Pointer("Martina10"),
            Password: criblcontrolplanesdkgo.Pointer("R9HovxE2l_Ve3VU"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10080,
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            PrometheusAPI: "/write",
            AuthType: components.AuthenticationTypeOptionsPrometheusAuthNone.ToPointer(),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("almost calmly against wherever gratefully versus"),
            Username: criblcontrolplanesdkgo.Pointer("Johnathon.Wolf"),
            Password: criblcontrolplanesdkgo.Pointer("xMC_bK0VYnsfmYM"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 514,
            MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            SingleMsgUDPPackets: criblcontrolplanesdkgo.Pointer(false),
            IngestRawBytes: criblcontrolplanesdkgo.Pointer(false),
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](6988.17),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("yuck trust hopelessly toward instead gadzooks oil solidly now"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "s3-notifications-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("/.*/"),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](1),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](600),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](1),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](300),
            SkipOnError: criblcontrolplanesdkgo.Pointer(false),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            EnableSQSAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            Preprocess: &components.PreprocessType{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](5),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](600),
            Checkpointing: &components.CheckpointingType{
                Enabled: false,
                Retries: criblcontrolplanesdkgo.Pointer[float64](5),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](10),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            TagAfterProcessing: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("lively oddly brood optimistically extremely soliloquy inquisitively"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "s3-inventory-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("/.*/"),
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
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](1),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](600),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](1),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](300),
            SkipOnError: criblcontrolplanesdkgo.Pointer(false),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            EnableSQSAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            Preprocess: &components.PreprocessType{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](5),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](600),
            Checkpointing: &components.CheckpointingType{
                Enabled: false,
                Retries: criblcontrolplanesdkgo.Pointer[float64](5),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](10),
            ChecksumSuffix: criblcontrolplanesdkgo.Pointer("checksum"),
            MaxManifestSizeKB: criblcontrolplanesdkgo.Pointer[int64](4096),
            ValidateInventoryFiles: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("extract until abandoned perky gladly unless yippee oh version successfully"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            QueueName: "security-lake-queue",
            FileFilter: criblcontrolplanesdkgo.Pointer("/.*/"),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](1),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](600),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](1),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](300),
            SkipOnError: criblcontrolplanesdkgo.Pointer(false),
            IncludeSqsMetadata: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            EnableSQSAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            Preprocess: &components.PreprocessType{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                },
            },
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](5),
            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](600),
            Checkpointing: &components.CheckpointingType{
                Enabled: false,
                Retries: criblcontrolplanesdkgo.Pointer[float64](5),
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](10),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("source happily how always helplessly quixotic relative what often"),
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
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
                        AuthProtocol: components.AuthenticationProtocolOptionsV3UserNone.ToPointer(),
                        AuthKey: criblcontrolplanesdkgo.Pointer("<value>"),
                        PrivProtocol: components.PrivacyProtocolNone.ToPointer(),
                        PrivKey: criblcontrolplanesdkgo.Pointer("<value>"),
                    },
                },
            },
            MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](6815.48),
            VarbindsWithTypes: criblcontrolplanesdkgo.Pointer(false),
            BestEffortParsing: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("reproach meh wallaby uncommon lotion repurpose how"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 9997,
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](1000),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](30),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](0),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            AuthTokens: []components.InputSplunkAuthToken{
                components.InputSplunkAuthToken{
                    Token: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("boo oof insignificant pfft opposite hunt"),
                },
            },
            MaxS2Sversion: components.MaxS2SVersionV3.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("until scoop fold indeed engender"),
            UseFwdTimezone: criblcontrolplanesdkgo.Pointer(true),
            DropControlFields: criblcontrolplanesdkgo.Pointer(true),
            ExtractMetrics: criblcontrolplanesdkgo.Pointer(false),
            Compress: components.InputSplunkCompressionDisabled.ToPointer(),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
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
                    Description: criblcontrolplanesdkgo.Pointer("before unless pinstripe knowledgeably courteous slide lighthearted"),
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
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: "<value>",
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
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
            SplunkHecAcks: criblcontrolplanesdkgo.Pointer(false),
            BreakerRulesets: []string{
                "<value 1>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            UseFwdTimezone: criblcontrolplanesdkgo.Pointer(true),
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
                "<value 3>",
            },
            EmitTokenMetrics: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("meh yowza manner failing"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            SearchHead: "https://localhost:8089",
            Search: "index=main",
            Earliest: criblcontrolplanesdkgo.Pointer("-16m@m"),
            Latest: criblcontrolplanesdkgo.Pointer("-1m@m"),
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
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](30),
            JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](3),
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            RetryRules: &components.RetryRulesType{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesBackoff,
                Interval: criblcontrolplanesdkgo.Pointer[float64](1000),
                Limit: criblcontrolplanesdkgo.Pointer[float64](5),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](2),
                Codes: []float64{
                    4786.8,
                    4387.89,
                    7031.34,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            BreakerRulesets: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            AuthType: components.InputSplunkSearchAuthenticationTypeBasic.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("sprinkles deflate fooey pricey below but quarrelsome"),
            Username: criblcontrolplanesdkgo.Pointer("Dulce28"),
            Password: criblcontrolplanesdkgo.Pointer("H0D8DX8APbKrKvi"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
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
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            MaxMessages: criblcontrolplanesdkgo.Pointer[float64](10),
            VisibilityTimeout: criblcontrolplanesdkgo.Pointer[float64](600),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            PollTimeout: criblcontrolplanesdkgo.Pointer[float64](10),
            Description: criblcontrolplanesdkgo.Pointer("blah yum intent er loyalty anenst gadzooks"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            NumReceivers: criblcontrolplanesdkgo.Pointer[float64](3),
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
                    MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                    CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                    MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                    MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                    Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                    Compress: components.CompressionOptionsPqNone.ToPointer(),
                    PqControls: &components.PqTypePqControls{},
                },
                Host: "0.0.0.0",
                UDPPort: 514,
                TCPPort: criblcontrolplanesdkgo.Pointer[float64](1808.19),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                TimestampTimezone: criblcontrolplanesdkgo.Pointer("local"),
                SingleMsgUDPPackets: criblcontrolplanesdkgo.Pointer(false),
                EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
                KeepFieldsList: []string{
                    "<value 1>",
                },
                OctetCounting: criblcontrolplanesdkgo.Pointer(false),
                InferFraming: criblcontrolplanesdkgo.Pointer(true),
                StrictlyInferOctetCounting: criblcontrolplanesdkgo.Pointer(true),
                AllowNonStandardAppName: criblcontrolplanesdkgo.Pointer(false),
                MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](1000),
                SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
                SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](30),
                SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](0),
                TLS: &components.TLSSettingsServerSideType{
                    Disabled: criblcontrolplanesdkgo.Pointer(true),
                    RequestCert: criblcontrolplanesdkgo.Pointer(false),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                    CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
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
                UDPSocketRxBufSize: criblcontrolplanesdkgo.Pointer[float64](1683.94),
                EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(false),
                Description: criblcontrolplanesdkgo.Pointer("omelet helpless sunbathe hmph mothball difficult"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](10),
            Host: &components.InputSystemMetricsHost{
                Mode: components.ModeOptionsHostBasic.ToPointer(),
                Custom: &components.InputSystemMetricsCustom{
                    System: &components.InputSystemMetricsSystem{
                        Mode: components.InputSystemMetricsSystemModeBasic.ToPointer(),
                        Processes: criblcontrolplanesdkgo.Pointer(false),
                    },
                    CPU: &components.InputSystemMetricsCPU{
                        Mode: components.InputSystemMetricsCPUModeBasic.ToPointer(),
                        PerCPU: criblcontrolplanesdkgo.Pointer(false),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Time: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Memory: &components.InputSystemMetricsMemory{
                        Mode: components.InputSystemMetricsMemoryModeBasic.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Network: &components.InputSystemMetricsNetwork{
                        Mode: components.InputSystemMetricsNetworkModeBasic.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Protocols: criblcontrolplanesdkgo.Pointer(false),
                        Devices: []string{
                            "<value 1>",
                            "<value 2>",
                            "<value 3>",
                        },
                        PerInterface: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Disk: &components.InputSystemMetricsDisk{
                        Mode: components.InputSystemMetricsDiskModeBasic.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Inodes: criblcontrolplanesdkgo.Pointer(false),
                        Devices: []string{
                            "<value 1>",
                            "<value 2>",
                        },
                        Mountpoints: []string{
                            "<value 1>",
                            "<value 2>",
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
            Container: &components.Container{
                Mode: components.ContainerModeBasic.ToPointer(),
                DockerSocket: []string{
                    "<value 1>",
                },
                DockerTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
                Filters: []components.InputSystemMetricsFilter{
                    components.InputSystemMetricsFilter{
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
            Persistence: &components.InputSystemMetricsPersistence{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("10m"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("1GB"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("24h"),
                Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/system_metrics"),
            },
            Description: criblcontrolplanesdkgo.Pointer("oddball boohoo perfectly libel culminate somber"),
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
            },
            Connections: []components.ItemsTypeConnectionsOptional{
                components.ItemsTypeConnectionsOptional{
                    Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                    Output: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Pq: &components.PqType{
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](300),
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
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Disk: &components.DisksAndFileSystems{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Metadata: &components.HostInfo{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Routes: &components.InputSystemStateRoutes{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                DNS: &components.DNS{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                User: &components.UsersAndGroups{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Firewall: &components.Firewall{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Services: &components.Services{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                Ports: &components.ListeningPorts{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
                LoginUsers: &components.LoggedInUsers{
                    Enable: criblcontrolplanesdkgo.Pointer(true),
                },
            },
            Persistence: &components.InputSystemStatePersistence{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("10m"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("1GB"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("24h"),
                Compress: components.DataCompressionFormatOptionsPersistenceNone.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/system_state"),
            },
            DisableNativeModule: criblcontrolplanesdkgo.Pointer(false),
            DisableNativeLastLogModule: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("catalog past whitewash seafood"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10090,
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](1000),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](30),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](0),
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
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
            EnableHeader: criblcontrolplanesdkgo.Pointer(false),
            Preprocess: &components.PreprocessType{
                Disabled: true,
                Command: criblcontrolplanesdkgo.Pointer("<value>"),
                Args: []string{
                    "<value 1>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("decongestant schnitzel supplier selfish mushy milestone heavy giant"),
            AuthToken: criblcontrolplanesdkgo.Pointer(""),
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
            },
            Connections: []components.ItemsTypeConnectionsOptional{
                components.ItemsTypeConnectionsOptional{
                    Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                    Output: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Pq: &components.PqType{
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 10090,
            TLS: &components.TLSSettingsServerSideType{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            IPWhitelistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            MaxActiveCxn: criblcontrolplanesdkgo.Pointer[float64](1000),
            SocketIdleTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketEndingMaxWait: criblcontrolplanesdkgo.Pointer[float64](30),
            SocketMaxLifespan: criblcontrolplanesdkgo.Pointer[float64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            EnableLoadBalancing: criblcontrolplanesdkgo.Pointer(false),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("abandoned sturdy phooey print charter quickly vol including towards preside"),
            AuthToken: criblcontrolplanesdkgo.Pointer(""),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 5985,
            AuthMethod: components.InputWefAuthenticationMethodClientCert.ToPointer(),
            TLS: &components.MTLSSettings{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                RequestCert: criblcontrolplanesdkgo.Pointer(true),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: "<value>",
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: "<value>",
                CaPath: "<value>",
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                OcspCheck: criblcontrolplanesdkgo.Pointer(false),
                Keytab: "<value>",
                Principal: "<value>",
                OcspCheckFailClose: criblcontrolplanesdkgo.Pointer(false),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](90),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            CaFingerprint: criblcontrolplanesdkgo.Pointer("<value>"),
            Keytab: criblcontrolplanesdkgo.Pointer("<value>"),
            Principal: criblcontrolplanesdkgo.Pointer("<value>"),
            AllowMachineIDMismatch: criblcontrolplanesdkgo.Pointer(false),
            Subscriptions: []components.Subscription{
                components.Subscription{
                    SubscriptionName: "subscription-1",
                    Version: criblcontrolplanesdkgo.Pointer("<value>"),
                    ContentFormat: components.InputWefFormatRenderedText,
                    HeartbeatInterval: 60,
                    BatchTimeout: 5,
                    ReadExistingEvents: criblcontrolplanesdkgo.Pointer(false),
                    SendBookmarks: criblcontrolplanesdkgo.Pointer(true),
                    Compress: criblcontrolplanesdkgo.Pointer(true),
                    Targets: []string{},
                    Locale: criblcontrolplanesdkgo.Pointer("en-US"),
                    QuerySelector: components.QueryBuilderModeSimple.ToPointer(),
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                    Queries: []components.Query{
                        components.Query{
                            Path: "/usr/sbin",
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
            Description: criblcontrolplanesdkgo.Pointer("eek boohoo meh heartache legend oof where reopen"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            LogNames: []string{
                "Application",
                "System",
            },
            ReadMode: components.ReadModeNewest.ToPointer(),
            EventFormat: components.EventFormatJSON.ToPointer(),
            DisableNativeModule: criblcontrolplanesdkgo.Pointer(false),
            Interval: criblcontrolplanesdkgo.Pointer[float64](10),
            BatchSize: criblcontrolplanesdkgo.Pointer[float64](500),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            MaxEventBytes: criblcontrolplanesdkgo.Pointer[float64](51200),
            Description: criblcontrolplanesdkgo.Pointer("when upward hovercraft ill trick after part"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Interval: criblcontrolplanesdkgo.Pointer[float64](10),
            Host: &components.InputWindowsMetricsHost{
                Mode: components.ModeOptionsHostBasic.ToPointer(),
                Custom: &components.InputWindowsMetricsCustom{
                    System: &components.InputWindowsMetricsSystem{
                        Mode: components.InputWindowsMetricsSystemModeBasic.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                    },
                    CPU: &components.InputWindowsMetricsCPU{
                        Mode: components.InputWindowsMetricsCPUModeBasic.ToPointer(),
                        PerCPU: criblcontrolplanesdkgo.Pointer(false),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
                        Time: criblcontrolplanesdkgo.Pointer(false),
                    },
                    Memory: &components.InputWindowsMetricsMemory{
                        Mode: components.InputWindowsMetricsMemoryModeBasic.ToPointer(),
                        Detail: criblcontrolplanesdkgo.Pointer(false),
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
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Persistence: &components.InputWindowsMetricsPersistence{
                Enable: criblcontrolplanesdkgo.Pointer(false),
                TimeWindow: criblcontrolplanesdkgo.Pointer("10m"),
                MaxDataSize: criblcontrolplanesdkgo.Pointer("1GB"),
                MaxDataTime: criblcontrolplanesdkgo.Pointer("24h"),
                Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
                DestPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/windows_metrics"),
            },
            DisableNativeModule: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("if quizzically ha zowie inferior pop now fundraising"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Endpoint: "https://api.wiz.io",
            AuthURL: "https://auth.wiz.io/oauth/token",
            AuthAudienceOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            ClientID: "client-id",
            ContentConfig: []components.InputWizContentConfig{},
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](300),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxMissedKeepAlives: criblcontrolplanesdkgo.Pointer[float64](3),
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            RetryRules: &components.RetryRulesType{
                Type: components.RetryTypeOptionsHealthCheckCollectorConfRetryRulesBackoff,
                Interval: criblcontrolplanesdkgo.Pointer[float64](1000),
                Limit: criblcontrolplanesdkgo.Pointer[float64](5),
                Multiplier: criblcontrolplanesdkgo.Pointer[float64](2),
                Codes: []float64{
                    4786.8,
                    4387.89,
                    7031.34,
                },
                EnableHeader: criblcontrolplanesdkgo.Pointer(true),
                RetryConnectTimeout: criblcontrolplanesdkgo.Pointer(false),
                RetryConnectReset: criblcontrolplanesdkgo.Pointer(false),
            },
            AuthType: components.AuthenticationMethodOptions1Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("hence when inspection um jovially comparison filter"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateEndpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAuthURL: criblcontrolplanesdkgo.Pointer("https://spiteful-ghost.name"),
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
                Mode: components.ModeOptionsPqAlways.ToPointer(),
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
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
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: criblcontrolplanesdkgo.Pointer(false),
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
            BreakerRulesets: []string{
                "<value 1>",
            },
            StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
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
                    Description: criblcontrolplanesdkgo.Pointer("mostly incidentally nearly with meanwhile"),
                    Metadata: []components.ItemsTypeMetadata{
                        components.ItemsTypeMetadata{
                            Name: "<value>",
                            Value: "<value>",
                        },
                    },
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("yowza upliftingly firm regarding"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Host: "0.0.0.0",
            Port: 8088,
            AuthTokens: []components.InputZscalerHecAuthToken{
                components.InputZscalerHecAuthToken{
                    AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
                    TokenSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Token: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("wearily after minor daintily modulo amid"),
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
                RequestCert: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CommonNameRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            MaxActiveReq: criblcontrolplanesdkgo.Pointer[float64](256),
            MaxRequestsPerSocket: criblcontrolplanesdkgo.Pointer[int64](0),
            EnableProxyHeader: criblcontrolplanesdkgo.Pointer(false),
            CaptureHeaders: criblcontrolplanesdkgo.Pointer(false),
            ActivityLogSampleRate: criblcontrolplanesdkgo.Pointer[float64](100),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            SocketTimeout: criblcontrolplanesdkgo.Pointer[float64](0),
            KeepAliveTimeout: criblcontrolplanesdkgo.Pointer[float64](5),
            EnableHealthCheck: "<value>",
            IPAllowlistRegex: criblcontrolplanesdkgo.Pointer("/.*/"),
            IPDenylistRegex: criblcontrolplanesdkgo.Pointer("/^$/"),
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
            HecAcks: criblcontrolplanesdkgo.Pointer(false),
            AccessControlAllowOrigin: []string{
                "<value 1>",
            },
            AccessControlAllowHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            EmitTokenMetrics: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("unaccountably often via till"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Filter: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("colonialism gradient acknowledge hmph yuck astride boo but boo now"),
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
                MaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1000),
                CommitFrequency: criblcontrolplanesdkgo.Pointer[float64](42),
                MaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                MaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                Path: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                Compress: components.CompressionOptionsPqNone.ToPointer(),
                PqControls: &components.PqTypePqControls{},
            },
            Prefix: criblcontrolplanesdkgo.Pointer("cribl.logstream."),
            FullFidelity: criblcontrolplanesdkgo.Pointer(true),
            Metadata: []components.ItemsTypeMetadata{
                components.ItemsTypeMetadata{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("correctly biodegrade for zowie boo irk instead mockingly considering"),
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