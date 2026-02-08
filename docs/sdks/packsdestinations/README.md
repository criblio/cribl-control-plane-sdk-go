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

### Example Usage: OutputCreateExamplesAzureBlob

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesAzureBlob" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyAzureBlob(
        operations.CreateOutputSystemByPackOutputAzureBlob{
            ID: "azure-blob-output",
            Type: operations.CreateOutputSystemByPackTypeAzureBlobAzureBlob,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ContainerName: "my-container",
            CreateContainer: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](1789.95),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](5235.4),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](373.78),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](2011.77),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](9670.5),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](3614.52),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](27.11),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4258.54),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1146.47),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1403.61),
            },
            AuthType: components.AuthenticationMethodOptionsClientCert.ToPointer(),
            StorageClass: operations.CreateOutputSystemByPackBlobAccessTierHot.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("concentration writhing defiantly shrill unfortunately recklessly when"),
            Compress: components.CompressionOptions2None.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestCompression.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV1.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](1367.63),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(true),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(false),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](9688.12),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](5926.86),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](7443.01),
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
            TemplateContainerName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateFormat: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateConnectionString: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateTenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateClientID: criblcontrolplanesdkgo.Pointer("<id>"),
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
### Example Usage: OutputCreateExamplesAzureDataExplorer

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesAzureDataExplorer" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyAzureDataExplorer(
        operations.CreateOutputSystemByPackOutputAzureDataExplorer{
            ID: "azure-data-explorer-output",
            Type: operations.CreateOutputSystemByPackTypeAzureDataExplorerAzureDataExplorer,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            ClusterURL: "https://mycluster.kusto.windows.net",
            Database: "mydatabase",
            Table: "mytable",
            ValidateDatabaseSettings: criblcontrolplanesdkgo.Pointer(true),
            IngestMode: operations.CreateOutputSystemByPackIngestionModeStreaming.ToPointer(),
            OauthEndpoint: components.MicrosoftEntraIDAuthenticationEndpointOptionsSaslHTTPSLoginMicrosoftonlineCom,
            TenantID: "tenant-id",
            ClientID: "client-id",
            Scope: "https://mycluster.kusto.windows.net/.default",
            OauthType: operations.CreateOutputSystemByPackOauthTypeAuthenticationMethodClientSecret,
            Description: criblcontrolplanesdkgo.Pointer("allegation pleasure orient anti rout forecast absent hence unto"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("client-secret"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Certificate: &operations.CreateOutputSystemByPackCertificate{
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            Format: components.DataFormatOptionsJSON.ToPointer(),
            Compress: components.CompressionOptions2Gzip,
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet10.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](8673.65),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(true),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](2046.39),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](6195.79),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](49.27),
            IsMappingObj: criblcontrolplanesdkgo.Pointer(false),
            MappingObj: criblcontrolplanesdkgo.Pointer("<value>"),
            MappingRef: criblcontrolplanesdkgo.Pointer("<value>"),
            IngestURL: criblcontrolplanesdkgo.Pointer("https://talkative-edge.biz/"),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            StagePath: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](3693.13),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](369.97),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](9166.41),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](6949.39),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](1551.41),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](27.11),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4258.54),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1146.47),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1403.61),
            },
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](7481.82),
            FlushImmediately: criblcontrolplanesdkgo.Pointer(false),
            RetainBlobOnSuccess: criblcontrolplanesdkgo.Pointer(false),
            ExtentTags: []operations.CreateOutputSystemByPackExtentTag{
                operations.CreateOutputSystemByPackExtentTag{
                    Prefix: operations.CreateOutputSystemByPackPrefixOptionalIngestBy.ToPointer(),
                    Value: "<value>",
                },
            },
            IngestIfNotExists: []operations.CreateOutputSystemByPackIngestIfNotExist{
                operations.CreateOutputSystemByPackIngestIfNotExist{
                    Value: "<value>",
                },
            },
            ReportLevel: operations.CreateOutputSystemByPackReportLevelFailuresAndSuccesses.ToPointer(),
            ReportMethod: operations.CreateOutputSystemByPackReportMethodQueue.ToPointer(),
            AdditionalProperties: []operations.CreateOutputSystemByPackAdditionalProperty{
                operations.CreateOutputSystemByPackAdditionalProperty{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](6050.98),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](9743.06),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](4864.39),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](9238.64),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8217.74),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2192.23),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](342.7),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsAzureDataExplorer{},
            TemplateClusterURL: criblcontrolplanesdkgo.Pointer("https://slow-airline.name"),
            TemplateDatabase: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateTable: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateTenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateClientID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateScope: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateFormat: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateIngestURL: criblcontrolplanesdkgo.Pointer("https://imaginative-contractor.com/"),
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
### Example Usage: OutputCreateExamplesAzureEventhub

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesAzureEventhub" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyAzureEventhub(
        operations.CreateOutputSystemByPackOutputAzureEventhub{
            ID: "azure-eventhub-output",
            Type: operations.CreateOutputSystemByPackTypeAzureEventhubAzureEventhub,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Brokers: []string{
                "myeventhub.servicebus.windows.net:9093",
            },
            Topic: "logs",
            Ack: components.AcknowledgmentsOptionsLeader.ToPointer(),
            Format: components.RecordDataFormatOptionsRaw.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](8675.97),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](8556.87),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8039.59),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](8255),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](2433.39),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](5930.87),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](2311.35),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2640.47),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6022.68),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](8625.98),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](9613.74),
            Sasl: &components.AuthenticationType1{
                Disabled: true,
                AuthType: components.AuthenticationMethodOptionsSasl1Manual.ToPointer(),
                Password: criblcontrolplanesdkgo.Pointer("AuRj3AcLPI0x9Ju"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSasl1Plain.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("Annetta33"),
                ClientSecretAuthType: components.AuthenticationMethodOptionsSasl2Certificate.ToPointer(),
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
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("academics where queasily woot nervously weird oh after pecan"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6690.26),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](802.2),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8801.27),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsAzureEventhub{},
            TemplateTopic: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesAzureLogs

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesAzureLogs" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyAzureLogs(
        operations.CreateOutputSystemByPackOutputAzureLogs{
            ID: "azure-logs-output",
            Type: operations.CreateOutputSystemByPackTypeAzureLogsAzureLogs,
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
            LogType: "Cribl",
            ResourceID: criblcontrolplanesdkgo.Pointer("<id>"),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](9102.79),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4801.03),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](8307.36),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](1571.25),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5174.17),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            APIURL: criblcontrolplanesdkgo.Pointer("https://ample-pliers.net"),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            AuthType: operations.CreateOutputSystemByPackAuthenticationMethodAzureLogsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("eek testing boo that"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8407),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2910.6),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](4740.31),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsAzureLogs{},
            WorkspaceID: criblcontrolplanesdkgo.Pointer("workspace-id"),
            WorkspaceKey: criblcontrolplanesdkgo.Pointer("workspace-key"),
            KeypairSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateWorkspaceID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateWorkspaceKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesChronicle

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesChronicle" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyChronicle(
        operations.CreateOutputSystemByPackOutputChronicle{
            ID: "chronicle-output",
            Type: operations.CreateOutputSystemByPackTypeChronicleChronicle,
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
            APIVersion: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthenticationMethod: operations.CreateOutputSystemByPackAuthenticationMethodChronicleServiceAccountSecret.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            Region: "us",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](6953.73),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](2955.01),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](5676.03),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](4660),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7930.45),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](2639.28),
            IngestionMethod: criblcontrolplanesdkgo.Pointer("<value>"),
            Namespace: criblcontrolplanesdkgo.Pointer("<value>"),
            LogType: "UNKNOWN",
            LogTextField: criblcontrolplanesdkgo.Pointer("<value>"),
            GcpProjectID: "my-project",
            GcpInstance: "customer-id",
            CustomLabels: []operations.CreateOutputSystemByPackCustomLabel{
                operations.CreateOutputSystemByPackCustomLabel{
                    Key: "<key>",
                    Value: "<value>",
                    RbacEnabled: criblcontrolplanesdkgo.Pointer(true),
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("testimonial conceal ah whose wetly except resource"),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            ServiceAccountCredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1950.82),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7546.11),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1188.64),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsChronicle{},
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesClickHouse

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesClickHouse" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyClickHouse(
        operations.CreateOutputSystemByPackOutputClickHouse{
            ID: "clickhouse-output",
            Type: operations.CreateOutputSystemByPackTypeClickHouseClickHouse,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            URL: "http://localhost:8123/",
            AuthType: operations.CreateOutputSystemByPackAuthenticationTypeClickHouseSslUserCertificate.ToPointer(),
            Database: "mydb",
            TableName: "mytable",
            Format: operations.CreateOutputSystemByPackFormatClickHouseJSONCompactEachRowWithNames.ToPointer(),
            MappingType: operations.CreateOutputSystemByPackMappingTypeCustom.ToPointer(),
            AsyncInserts: criblcontrolplanesdkgo.Pointer(true),
            TLS: &components.TLSSettingsClientSideType1{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](3199.11),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](392.9),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](5643.64),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](294.98),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7689.72),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            DumpFormatErrorsToDisk: criblcontrolplanesdkgo.Pointer(true),
            StatsDestination: &operations.CreateOutputSystemByPackStatsDestination{
                URL: criblcontrolplanesdkgo.Pointer("https://supportive-veto.com/"),
                Database: criblcontrolplanesdkgo.Pointer("<value>"),
                TableName: criblcontrolplanesdkgo.Pointer("<value>"),
                AuthType: criblcontrolplanesdkgo.Pointer("<value>"),
                Username: criblcontrolplanesdkgo.Pointer("Earline4"),
                SQLUsername: criblcontrolplanesdkgo.Pointer("<value>"),
                Password: criblcontrolplanesdkgo.Pointer("fn_ubMAFMtOwBpx"),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("instead afterwards chunter exalted oh as"),
            Username: criblcontrolplanesdkgo.Pointer("Bethel26"),
            Password: criblcontrolplanesdkgo.Pointer("dN33b8ZkjXUCnHP"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            SQLUsername: criblcontrolplanesdkgo.Pointer("<value>"),
            WaitForAsyncInserts: criblcontrolplanesdkgo.Pointer(true),
            ExcludeMappingFields: []string{
                "<value 1>",
            },
            DescribeTable: criblcontrolplanesdkgo.Pointer("<value>"),
            ColumnMappings: []operations.CreateOutputSystemByPackColumnMapping{
                operations.CreateOutputSystemByPackColumnMapping{
                    ColumnName: "<value>",
                    ColumnType: criblcontrolplanesdkgo.Pointer("<value>"),
                    ColumnValueExpression: "<value>",
                },
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1442.22),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](6319.08),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6685.88),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsClickHouse{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://early-finer.info/"),
            TemplateDatabase: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateTableName: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesCloudflareR2

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesCloudflareR2" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyCloudflareR2(
        operations.CreateOutputSystemByPackOutputCloudflareR2{
            ID: "cloudflare-r2-output",
            Type: operations.CreateOutputSystemByPackTypeCloudflareR2CloudflareR2,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Endpoint: "https://account-id.r2.cloudflarestorage.com",
            Bucket: "my-bucket",
            AwsAuthenticationMethod: operations.CreateOutputSystemByPackAuthenticationMethodCloudflareR2Auto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "<value>",
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions5V4.ToPointer(),
            ObjectACL: "<value>",
            StorageClass: components.StorageClassOptions2Standard.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionOptionsAes256.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsParquet.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](8892.26),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](525.29),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](7938.95),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(true),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](27.11),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4258.54),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1146.47),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1403.61),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](1135.53),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](4195.99),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](3062.39),
            Description: criblcontrolplanesdkgo.Pointer("opera catalyze absentmindedly shudder lyre nor familiar what notwithstanding considering"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2None.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestCompression.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV1.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](3900.48),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(false),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](2838.89),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](8968.38),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](5274.03),
            TemplateBucket: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateFormat: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesCloudwatch

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesCloudwatch" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyCloudwatch(
        operations.CreateOutputSystemByPackOutputCloudwatch{
            ID: "cloudwatch-output",
            Type: operations.CreateOutputSystemByPackTypeCloudwatchCloudwatch,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            LogGroupName: "my-log-group",
            LogStreamName: "my-log-stream",
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](7012.18),
            MaxQueueSize: criblcontrolplanesdkgo.Pointer[float64](3270.64),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](3913.32),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](4120.36),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("qua ugly hypothesise direct er vivaciously"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](4860.51),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7884.96),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1997.19),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsCloudwatch{},
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
    if res.CountedOutput != nil {
        // handle response
    }
}
```
### Example Usage: OutputCreateExamplesConfluentCloud

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesConfluentCloud" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyConfluentCloud(
        operations.CreateOutputSystemByPackOutputConfluentCloud{
            ID: "confluent-cloud-output",
            Type: operations.CreateOutputSystemByPackTypeConfluentCloudConfluentCloud,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
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
            Topic: "logs",
            Ack: components.AcknowledgmentsOptions1Leader.ToPointer(),
            Format: components.RecordDataFormatOptions1Raw.ToPointer(),
            Compression: components.CompressionOptions3Zstd.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](6366.62),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](5374.13),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5869.76),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType1{
                Disabled: false,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://orderly-knitting.info/"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](65.04),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7620.77),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](1627.05),
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
                DefaultKeySchemaID: criblcontrolplanesdkgo.Pointer[float64](3110.3),
                DefaultValueSchemaID: criblcontrolplanesdkgo.Pointer[float64](4216.12),
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](5394.03),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](2190.11),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](1441.82),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](8482.93),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1824.83),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7880.22),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](5479.88),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](9839.91),
            Sasl: &components.AuthenticationType{
                Disabled: true,
                Username: criblcontrolplanesdkgo.Pointer("Jeromy34"),
                Password: criblcontrolplanesdkgo.Pointer("mkbiANI3XTLISi6"),
                AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslScramSha512.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(true),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://definitive-kielbasa.com"),
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
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("possible configuration headline sparse outside forgo"),
            ProtobufLibraryID: criblcontrolplanesdkgo.Pointer("<id>"),
            ProtobufEncodingID: criblcontrolplanesdkgo.Pointer("<id>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1939.48),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3342.59),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6520.74),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsConfluentCloud{},
            TemplateTopic: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesCriblHttp

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesCriblHttp" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyCriblHTTP(
        operations.CreateOutputSystemByPackOutputCriblHTTP{
            ID: "cribl-http-output",
            Type: operations.CreateOutputSystemByPackTypeCriblHTTPCriblHTTP,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            LoadBalanced: criblcontrolplanesdkgo.Pointer(false),
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
            TokenTTLMinutes: criblcontrolplanesdkgo.Pointer[float64](1134.2),
            ExcludeFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Compression: components.CompressionOptions1Gzip.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1653.74),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](9494.28),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](8070.97),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](6216.85),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5963.98),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            AuthTokens: []components.ItemsTypeAuthTokens1{
                components.ItemsTypeAuthTokens1{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                    Description: criblcontrolplanesdkgo.Pointer("as patiently incidentally favorite failing festival blah upwardly modulo embed"),
                },
            },
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("around status liquid"),
            URL: criblcontrolplanesdkgo.Pointer("https://circular-disconnection.com"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(true),
            Urls: []components.ItemsTypeUrls{
                components.ItemsTypeUrls{
                    URL: "https://serene-challenge.name/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](50.28),
                    TemplateURL: criblcontrolplanesdkgo.Pointer("https://hefty-hawk.name"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](5021.98),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](4608.29),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](145.09),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7044.3),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](4123.65),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsCriblHTTP{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://awesome-availability.info"),
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
### Example Usage: OutputCreateExamplesCriblLake

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesCriblLake" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyCriblLake(
        operations.CreateOutputSystemByPackOutputCriblLake{
            ID: "cribl-lake-output",
            Type: operations.CreateOutputSystemByPackTypeCriblLakeCriblLake,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Bucket: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](5759.58),
            StagePath: criblcontrolplanesdkgo.Pointer("<value>"),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            ObjectACL: components.ObjectACLOptionsBucketOwnerRead.ToPointer(),
            StorageClass: components.StorageClassOptionsDeepArchive.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAes256.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](3018.8),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](3645.7),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](8651.05),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(true),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](27.11),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4258.54),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1146.47),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1403.61),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](9215.57),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](8392.15),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](5089.57),
            AwsAuthenticationMethod: components.AwsAuthenticationMethodOptionsAutoRPC.ToPointer(),
            Format: components.FormatOptionsJSON.ToPointer(),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](9776.75),
            Description: criblcontrolplanesdkgo.Pointer("immediately sham given quirkily since lean"),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](6637.62),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](7553.43),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](3491.79),
            TemplateBucket: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateDestPath: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesCriblSearchEngine

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesCriblSearchEngine" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyCriblSearchEngine(
        operations.CreateOutputSystemByPackOutputCriblSearchEngine{
            ID: "cribl-search-engine-output",
            Type: operations.CreateOutputSystemByPackTypeCriblSearchEngineCriblSearchEngine,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "cribl_pipe",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{},
            LoadBalanced: criblcontrolplanesdkgo.Pointer(false),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            TokenTTLMinutes: criblcontrolplanesdkgo.Pointer[float64](60),
            ExcludeFields: []string{
                "__kube_*",
                "__metadata",
                "__winEvent",
            },
            Compression: components.CompressionOptions1Gzip.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            SafeHeaders: []string{},
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 401,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](20000),
                },
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 403,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](20000),
                },
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 408,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](250),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 429,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 500,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](250),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 502,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](250),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 503,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](250),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 504,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](250),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 509,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](250),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5510.34),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6720.13),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](214.03),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            AuthTokens: []components.ItemsTypeAuthTokens1{
                components.ItemsTypeAuthTokens1{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                    Description: criblcontrolplanesdkgo.Pointer("as patiently incidentally favorite failing festival blah upwardly modulo embed"),
                },
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("lest lest oof"),
            URL: criblcontrolplanesdkgo.Pointer("https://0.0.0.0:10200"),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(true),
            Urls: []components.ItemsTypeUrls{
                components.ItemsTypeUrls{
                    URL: "https://serene-challenge.name/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](50.28),
                    TemplateURL: criblcontrolplanesdkgo.Pointer("https://hefty-hawk.name"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](9867.26),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](562.01),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](918.86),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1938.18),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](7222.85),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsCriblSearchEngine{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://bossy-vestment.name/"),
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
### Example Usage: OutputCreateExamplesCriblTcp

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesCriblTcp" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyCriblTCP(
        operations.CreateOutputSystemByPackOutputCriblTCP{
            ID: "cribl-tcp-output",
            Type: operations.CreateOutputSystemByPackTypeCriblTCPCriblTCP,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
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
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](9223.46),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](3115.03),
            TokenTTLMinutes: criblcontrolplanesdkgo.Pointer[float64](5115.16),
            AuthTokens: []components.ItemsTypeAuthTokens{
                components.ItemsTypeAuthTokens{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                    Description: criblcontrolplanesdkgo.Pointer("tuba while yuck ah boastfully"),
                },
            },
            ExcludeFields: []string{
                "<value 1>",
            },
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("ashamed voluminous mmm"),
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](10090),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
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
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](3954.84),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1051.16),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](7291.71),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](569.47),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](6128.38),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5664.51),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsCriblTCP{},
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
### Example Usage: OutputCreateExamplesCrowdstrikeNextGenSiem

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesCrowdstrikeNextGenSiem" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyCrowdstrikeNextGenSiem(
        operations.CreateOutputSystemByPackOutputCrowdstrikeNextGenSiem{
            ID: "crowdstrike-next-gen-siem-output",
            Type: operations.CreateOutputSystemByPackTypeCrowdstrikeNextGenSiemCrowdstrikeNextGenSiem,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            URL: "https://ingest.us.crowdstrike.com/api/ingest/hec/connection-id/v1/services/collector",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](6085.06),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](7765.33),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](596.15),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](2515.16),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1760.23),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            Format: components.RequestFormatOptionsJSON,
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("apropos instead masquerade gray numeric polarisation"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8044.56),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2332.3),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](2926.21),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsCrowdstrikeNextGenSiem{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://impressive-vista.com"),
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
### Example Usage: OutputCreateExamplesDatabricks

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesDatabricks" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyDatabricks(
        operations.CreateOutputSystemByPackOutputDatabricks{
            ID: "databricks-output",
            Type: operations.CreateOutputSystemByPackTypeDatabricksDatabricks,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            StagePath: criblcontrolplanesdkgo.Pointer("<value>"),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsRaw.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](2478.02),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](7056.55),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](7694.51),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](6708.36),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](3215.06),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](27.11),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4258.54),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1146.47),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1403.61),
            },
            WorkspaceID: "your-workspace-id",
            Scope: "my-scope",
            ClientID: "your-client-id",
            Catalog: "my-catalog",
            Schema: "my-schema",
            EventsVolumeName: "my-volume",
            ClientTextSecret: "your-client-secret",
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](7590.35),
            Description: criblcontrolplanesdkgo.Pointer("where pish what astride underpants upright veg solemnly"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](673.74),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(true),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(false),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(true),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](7208.74),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1011.97),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](3748.53),
            TemplateFormat: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesDatadog

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesDatadog" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyDatadog(
        operations.CreateOutputSystemByPackOutputDatadog{
            ID: "datadog-output",
            Type: operations.CreateOutputSystemByPackTypeDatadogDatadog,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ContentType: operations.CreateOutputSystemByPackSendLogsAsJSON.ToPointer(),
            Message: criblcontrolplanesdkgo.Pointer("<value>"),
            Source: criblcontrolplanesdkgo.Pointer("<value>"),
            Host: criblcontrolplanesdkgo.Pointer("primary-casement.name"),
            Service: criblcontrolplanesdkgo.Pointer("<value>"),
            Tags: []string{
                "<value 1>",
                "<value 2>",
            },
            BatchByTags: criblcontrolplanesdkgo.Pointer(false),
            AllowAPIKeyFromEvents: criblcontrolplanesdkgo.Pointer(false),
            Severity: operations.CreateOutputSystemByPackSeverityDatadogError.ToPointer(),
            Site: operations.CreateOutputSystemByPackDatadogSiteCustom.ToPointer(),
            SendCountersAsCount: criblcontrolplanesdkgo.Pointer(false),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](4996.19),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](783.8),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](3498.58),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](6306.6),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3548.4),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Secret.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](2050.93),
            Description: criblcontrolplanesdkgo.Pointer("well-worn sternly silk"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://tall-mouser.net"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8870.68),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3588.47),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](2596.94),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsDatadog{},
            APIKey: criblcontrolplanesdkgo.Pointer("your-api-key"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesDataset

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesDataset" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyDataset(
        operations.CreateOutputSystemByPackOutputDataset{
            ID: "dataset-output",
            Type: operations.CreateOutputSystemByPackTypeDatasetDataset,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            MessageField: criblcontrolplanesdkgo.Pointer("<value>"),
            ExcludeFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ServerHostField: criblcontrolplanesdkgo.Pointer("<value>"),
            TimestampField: criblcontrolplanesdkgo.Pointer("<value>"),
            DefaultSeverity: operations.CreateOutputSystemByPackDefaultSeveritySeverityFinest.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            Site: operations.CreateOutputSystemByPackDataSetSiteEu.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](975.15),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](7090.23),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](6642.59),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](6897.1),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7457.7),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Secret.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](7282.67),
            Description: criblcontrolplanesdkgo.Pointer("of pfft gladly"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://harmful-intellect.info/"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](3554),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](6448.6),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](3676.42),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsDataset{},
            APIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateCustomURL: criblcontrolplanesdkgo.Pointer("https://acceptable-compromise.biz/"),
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
### Example Usage: OutputCreateExamplesDiskSpool

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesDiskSpool" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyDiskSpool(
        operations.CreateOutputSystemByPackOutputDiskSpool{
            ID: "disk-spool-output",
            Type: operations.CreateOutputSystemByPackTypeDiskSpoolDiskSpool,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptionsPersistenceNone.ToPointer(),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("sorrowful excepting tattered"),
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
### Example Usage: OutputCreateExamplesDlS3

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesDlS3" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyDlS3(
        operations.CreateOutputSystemByPackOutputDlS3{
            ID: "dl-s3-output",
            Type: operations.CreateOutputSystemByPackTypeDlS3DlS3,
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
            Bucket: "my-bucket",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfManual.ToPointer(),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](2701.53),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            ObjectACL: components.ObjectACLOptionsAuthenticatedRead.ToPointer(),
            StorageClass: components.StorageClassOptionsDeepArchive.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAwsKms.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](2614.56),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](8785.34),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](5785.31),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](27.11),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4258.54),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1146.47),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1403.61),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](8096.65),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](3785.78),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](4826.39),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(false),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](913.63),
            PartitioningFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Description: criblcontrolplanesdkgo.Pointer("beside furthermore merry rightfully boldly gadzooks lumpy inborn"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsNormal.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet10.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV1.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](6089.28),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(true),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](6278.92),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](6525.85),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](9619.17),
            TemplateBucket: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateFormat: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesDynatraceHttp

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesDynatraceHttp" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyDynatraceHTTP(
        operations.CreateOutputSystemByPackOutputDynatraceHTTP{
            ID: "dynatrace-http-output",
            Type: operations.CreateOutputSystemByPackTypeDynatraceHTTPDynatraceHTTP,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Method: components.MethodOptionsPut.ToPointer(),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](630.09),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](6057.03),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](5392.66),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](6660.13),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8150.62),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            AuthType: operations.CreateOutputSystemByPackAuthenticationTypeDynatraceHTTPToken.ToPointer(),
            Format: operations.CreateOutputSystemByPackFormatDynatraceHTTPJSONArray,
            Endpoint: operations.CreateOutputSystemByPackEndpointCloud,
            TelemetryType: operations.CreateOutputSystemByPackTelemetryTypeLogs,
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](6491.32),
            Description: criblcontrolplanesdkgo.Pointer("aw beyond than so"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2384.36),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](6495.59),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8006.32),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsDynatraceHTTP{},
            Token: criblcontrolplanesdkgo.Pointer("your-api-key"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            EnvironmentID: criblcontrolplanesdkgo.Pointer("<id>"),
            ActiveGateDomain: criblcontrolplanesdkgo.Pointer("<value>"),
            URL: criblcontrolplanesdkgo.Pointer("https://frail-doorpost.biz/"),
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://altruistic-inspection.com"),
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
### Example Usage: OutputCreateExamplesDynatraceOtlp

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesDynatraceOtlp" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyDynatraceOtlp(
        operations.CreateOutputSystemByPackOutputDynatraceOtlp{
            ID: "dynatrace-otlp-output",
            Type: operations.CreateOutputSystemByPackTypeDynatraceOtlpDynatraceOtlp,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Protocol: operations.CreateOutputSystemByPackProtocolDynatraceOtlpHTTP,
            Endpoint: "https://your-environment.live.dynatrace.com/api/v2/otlp",
            OtlpVersion: components.OtlpVersionOptions1OneDot3Dot1,
            Compress: components.CompressionOptions4None.ToPointer(),
            HTTPCompress: components.CompressionOptions5None.ToPointer(),
            HTTPTracesEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPMetricsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPLogsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5411.14),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](678.56),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](3256.42),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](2512.28),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](8923.37),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](612.86),
            KeepAlive: criblcontrolplanesdkgo.Pointer(false),
            EndpointType: operations.CreateOutputSystemByPackEndpointTypeSaas,
            TokenSecret: "your-token-secret",
            AuthTokenName: criblcontrolplanesdkgo.Pointer("<value>"),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("petticoat toward ack weight breakable until"),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            SafeHeaders: []string{
                "<value 1>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8121.78),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5538.56),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6321.23),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsDynatraceOtlp{},
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
### Example Usage: OutputCreateExamplesElastic

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesElastic" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyElastic(
        operations.CreateOutputSystemByPackOutputElastic{
            ID: "elastic-output",
            Type: operations.CreateOutputSystemByPackTypeElasticElastic,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            LoadBalanced: criblcontrolplanesdkgo.Pointer(true),
            Index: "logs",
            DocType: criblcontrolplanesdkgo.Pointer("<value>"),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1502.29),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](3134.83),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](137.38),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](4867.48),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5125.77),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            ExtraParams: []components.ItemsTypeSaslSaslExtensions{
                components.ItemsTypeSaslSaslExtensions{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Auth: &components.AuthType{
                Disabled: true,
                Username: criblcontrolplanesdkgo.Pointer("Lou_Pfeffer24"),
                Password: criblcontrolplanesdkgo.Pointer("tuLWEvVJ7xoDNhs"),
                AuthType: components.AuthenticationMethodOptionsAuthManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                ManualAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            ElasticVersion: operations.CreateOutputSystemByPackElasticVersionSix.ToPointer(),
            ElasticPipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            IncludeDocID: criblcontrolplanesdkgo.Pointer(true),
            WriteAction: operations.CreateOutputSystemByPackWriteActionIndex.ToPointer(),
            RetryPartialErrors: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("hm irritably traditionalism amazing"),
            URL: criblcontrolplanesdkgo.Pointer("https://intelligent-management.org"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(true),
            Urls: []operations.CreateOutputSystemByPackURLElastic{
                operations.CreateOutputSystemByPackURLElastic{
                    URL: "https://legal-traditionalism.biz/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](4899.28),
                    TemplateURL: criblcontrolplanesdkgo.Pointer("https://hard-to-find-kinase.info"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](3663.2),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6869.14),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9533.05),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5259.8),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](7333.84),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsElastic{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://orange-legging.info"),
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
### Example Usage: OutputCreateExamplesElasticCloud

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesElasticCloud" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyElasticCloud(
        operations.CreateOutputSystemByPackOutputElasticCloud{
            ID: "elastic-cloud-output",
            Type: operations.CreateOutputSystemByPackTypeElasticCloudElasticCloud,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            URL: "my-cloud-id",
            Index: "logs",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1121.92),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](9781.82),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](791.89),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](8807.42),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](326.05),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            ExtraParams: []components.ItemsTypeSaslSaslExtensions{
                components.ItemsTypeSaslSaslExtensions{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Auth: &components.AuthType{
                Disabled: true,
                Username: criblcontrolplanesdkgo.Pointer("Lou_Pfeffer24"),
                Password: criblcontrolplanesdkgo.Pointer("tuLWEvVJ7xoDNhs"),
                AuthType: components.AuthenticationMethodOptionsAuthManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                ManualAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            ElasticPipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            IncludeDocID: criblcontrolplanesdkgo.Pointer(false),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("sans rationalize object sniveling towards persecute jubilantly kindly"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8354.98),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7384.07),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](3294.09),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsElasticCloud{},
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
### Example Usage: OutputCreateExamplesExabeam

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesExabeam" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyExabeam(
        operations.CreateOutputSystemByPackOutputExabeam{
            ID: "exabeam-output",
            Type: operations.CreateOutputSystemByPackTypeExabeamExabeam,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Bucket: "my-bucket",
            Region: "us-east1",
            StagePath: "/tmp/staging",
            Endpoint: "https://storage.googleapis.com",
            SignatureVersion: components.SignatureVersionOptions4V2.ToPointer(),
            ObjectACL: components.ObjectACLOptions1ProjectPrivate.ToPointer(),
            StorageClass: components.StorageClassOptions1Archive.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](9154.74),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](2771.71),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](7467.87),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](27.11),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4258.54),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1146.47),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1403.61),
            },
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](5791.63),
            EncodedConfiguration: criblcontrolplanesdkgo.Pointer("<value>"),
            CollectorInstanceID: "11112222-3333-4444-5555-666677778888",
            SiteName: criblcontrolplanesdkgo.Pointer("<value>"),
            SiteID: criblcontrolplanesdkgo.Pointer("<id>"),
            TimezoneOffset: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("likewise home ick"),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](1922.62),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](8157.19),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](5052.55),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesFilesystem

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesFilesystem" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyFilesystem(
        operations.CreateOutputSystemByPackOutputFilesystem{
            ID: "filesystem-output",
            Type: operations.CreateOutputSystemByPackTypeFilesystemFilesystem,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            DestPath: "/var/log/output",
            StagePath: criblcontrolplanesdkgo.Pointer("<value>"),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsRaw.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](2839.52),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](9635.12),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](3867.38),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](7228.29),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](3951.73),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(true),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](27.11),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4258.54),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1146.47),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1403.61),
            },
            Description: criblcontrolplanesdkgo.Pointer("huddle solemnly divert information"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet10.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV1.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](9998.38),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(false),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(true),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](1168.77),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](7818.69),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](9073.02),
            TemplateFormat: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesGoogleChronicle

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesGoogleChronicle" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyGoogleChronicle(
        operations.CreateOutputSystemByPackOutputGoogleChronicle{
            ID: "google-chronicle-output",
            Type: operations.CreateOutputSystemByPackTypeGoogleChronicleGoogleChronicle,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            APIVersion: operations.CreateOutputSystemByPackAPIVersionV2.ToPointer(),
            AuthenticationMethod: operations.CreateOutputSystemByPackAuthenticationMethodGoogleChronicleManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            LogFormatType: operations.CreateOutputSystemByPackSendEventsAsUnstructured,
            Region: criblcontrolplanesdkgo.Pointer("us"),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](6836.19),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4752.71),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](3157.86),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](4825.55),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3633.7),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](8409.11),
            Description: criblcontrolplanesdkgo.Pointer("juggernaut instead sympathetically"),
            ExtraLogTypes: []operations.CreateOutputSystemByPackExtraLogType{
                operations.CreateOutputSystemByPackExtraLogType{
                    LogType: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("surface respectful when popularize"),
                },
            },
            LogType: criblcontrolplanesdkgo.Pointer("<value>"),
            LogTextField: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomerID: criblcontrolplanesdkgo.Pointer("customer-id"),
            Namespace: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomLabels: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            UdmType: operations.CreateOutputSystemByPackUDMTypeEntities.ToPointer(),
            APIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            APIKeySecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            ServiceAccountCredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](3416.4),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](635.3),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8140.42),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsGoogleChronicle{},
            TemplateAPIVersion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateCustomerID: criblcontrolplanesdkgo.Pointer("<id>"),
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
### Example Usage: OutputCreateExamplesGoogleCloudLogging

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesGoogleCloudLogging" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyGoogleCloudLogging(
        operations.CreateOutputSystemByPackOutputGoogleCloudLogging{
            ID: "google-cloud-logging-output",
            Type: operations.CreateOutputSystemByPackTypeGoogleCloudLoggingGoogleCloudLogging,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            LogLocationType: operations.CreateOutputSystemByPackLogLocationTypeProject,
            LogNameExpression: "my-log",
            SanitizeLogNames: criblcontrolplanesdkgo.Pointer(true),
            PayloadFormat: operations.CreateOutputSystemByPackPayloadFormatJSON.ToPointer(),
            LogLabels: []components.ItemsTypeLogLabels{
                components.ItemsTypeLogLabels{
                    Label: "<value>",
                    ValueExpression: "<value>",
                },
            },
            ResourceTypeExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            ResourceTypeLabels: []components.ItemsTypeLogLabels{
                components.ItemsTypeLogLabels{
                    Label: "<value>",
                    ValueExpression: "<value>",
                },
            },
            SeverityExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            InsertIDExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            GoogleAuthMethod: components.GoogleAuthenticationMethodOptionsAuto.ToPointer(),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1744.96),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](3426.33),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5934.82),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](2939.76),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](7127.91),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](1175.45),
            ThrottleRateReqPerSec: criblcontrolplanesdkgo.Pointer[int64](625818),
            RequestMethodExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            RequestURLExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            RequestSizeExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            StatusExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            ResponseSizeExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            UserAgentExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            RemoteIPExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            ServerIPExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            RefererExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            LatencyExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            CacheLookupExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            CacheHitExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            CacheValidatedExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            CacheFillBytesExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            ProtocolExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            IDExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            ProducerExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            FirstExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            LastExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            FileExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            LineExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            FunctionExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            UIDExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            IndexExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            TotalSplitsExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            TraceExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            SpanIDExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            TraceSampledExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](2579.72),
            Description: criblcontrolplanesdkgo.Pointer("ack instead at whether aw frizzy inasmuch husk motionless"),
            LogLocationExpression: "my-project",
            PayloadExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9650.84),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3600.26),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8654.63),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsGoogleCloudLogging{},
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
### Example Usage: OutputCreateExamplesGoogleCloudStorage

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesGoogleCloudStorage" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyGoogleCloudStorage(
        operations.CreateOutputSystemByPackOutputGoogleCloudStorage{
            ID: "google-cloud-storage-output",
            Type: operations.CreateOutputSystemByPackTypeGoogleCloudStorageGoogleCloudStorage,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Bucket: "my-bucket",
            Region: "us-east1",
            Endpoint: "https://storage.googleapis.com",
            SignatureVersion: components.SignatureVersionOptions4V4.ToPointer(),
            AwsAuthenticationMethod: operations.CreateOutputSystemByPackAuthenticationMethodGoogleCloudStorageManual.ToPointer(),
            StagePath: "/tmp/staging",
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            ObjectACL: components.ObjectACLOptions1BucketOwnerRead.ToPointer(),
            StorageClass: components.StorageClassOptions1Nearline.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsParquet.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](1542.83),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](83.66),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](2030.04),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](4148.44),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](5005.38),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](27.11),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4258.54),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1146.47),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1403.61),
            },
            Description: criblcontrolplanesdkgo.Pointer("chase untrue maintainer searchingly athwart bootleg"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet10.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](5409.47),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(false),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(true),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](5064.36),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](7802.24),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](8400.2),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateBucket: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateFormat: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesGooglePubsub

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesGooglePubsub" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyGooglePubsub(
        operations.CreateOutputSystemByPackOutputGooglePubsub{
            ID: "google-pubsub-output",
            Type: operations.CreateOutputSystemByPackTypeGooglePubsubGooglePubsub,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            TopicName: "my-topic",
            CreateTopic: criblcontrolplanesdkgo.Pointer(true),
            OrderedDelivery: criblcontrolplanesdkgo.Pointer(true),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            GoogleAuthMethod: components.GoogleAuthenticationMethodOptionsAuto.ToPointer(),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            BatchSize: criblcontrolplanesdkgo.Pointer[float64](5376.75),
            BatchTimeout: criblcontrolplanesdkgo.Pointer[float64](8928.52),
            MaxQueueSize: criblcontrolplanesdkgo.Pointer[float64](2707.47),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](1466.77),
            FlushPeriod: criblcontrolplanesdkgo.Pointer[float64](9515.96),
            MaxInProgress: criblcontrolplanesdkgo.Pointer[float64](3285.97),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("scarcely bathhouse notwithstanding unless gadzooks across wound in-joke infinite"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9743.6),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4533.61),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](230.29),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsGooglePubsub{},
            TemplateTopicName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesGrafanaCloud

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesGrafanaCloud" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyGrafanaCloud(
        operations.CreateCreateOutputSystemByPackOutputGrafanaCloudUnionCreateOutputSystemByPackOutputGrafanaCloudGrafanaCloud1(
            operations.CreateOutputSystemByPackOutputGrafanaCloudGrafanaCloud1{
                ID: "grafana-cloud-output",
                Type: operations.CreateOutputSystemByPackOutputGrafanaCloudType1GrafanaCloud,
                Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                SystemFields: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
                Environment: criblcontrolplanesdkgo.Pointer("<value>"),
                Streamtags: []string{
                    "<value 1>",
                    "<value 2>",
                },
                LokiURL: "https://logs-prod-us-central1.grafana.net",
                PrometheusURL: criblcontrolplanesdkgo.Pointer("https://steep-amendment.name"),
                Message: criblcontrolplanesdkgo.Pointer("<value>"),
                MessageFormat: components.MessageFormatOptionsJSON.ToPointer(),
                Labels: []components.ItemsTypeLabels{
                    components.ItemsTypeLabels{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                MetricRenameExpr: criblcontrolplanesdkgo.Pointer("<value>"),
                PrometheusAuth: &components.PrometheusAuthType{
                    AuthType: components.AuthenticationTypeOptionsPrometheusAuth1CredentialsSecret.ToPointer(),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Username: criblcontrolplanesdkgo.Pointer("Jamey_Pagac"),
                    Password: criblcontrolplanesdkgo.Pointer("FunP8vUHMJ7SCJw"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                LokiAuth: &components.PrometheusAuthType{
                    AuthType: components.AuthenticationTypeOptionsPrometheusAuth1CredentialsSecret.ToPointer(),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Username: criblcontrolplanesdkgo.Pointer("Ford_Funk55"),
                    Password: criblcontrolplanesdkgo.Pointer("Qsq6QXVCZ8Kn7zO"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                Concurrency: criblcontrolplanesdkgo.Pointer[float64](957.76),
                MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](7013.67),
                MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](1661.58),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](6284.86),
                FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8783.91),
                ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                    components.ItemsTypeExtraHTTPHeaders{
                        Name: criblcontrolplanesdkgo.Pointer("<value>"),
                        Value: "<value>",
                    },
                },
                UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
                FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
                SafeHeaders: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
                ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                    components.ItemsTypeResponseRetrySettings{
                        HTTPStatus: 8177.04,
                        InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                        BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                        MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                    },
                },
                TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                    TimeoutRetry: true,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
                },
                ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
                OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
                Description: criblcontrolplanesdkgo.Pointer("glisten but apropos label solder meanwhile"),
                Compress: criblcontrolplanesdkgo.Pointer(true),
                PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
                PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2106.95),
                PqMode: components.ModeOptionsError.ToPointer(),
                PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3413.11),
                PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](7713.97),
                PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
                PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
                PqControls: &operations.CreateOutputSystemByPackOutputGrafanaCloudPqControls1{},
                TemplateLokiURL: criblcontrolplanesdkgo.Pointer("https://excellent-grass.name"),
                TemplatePrometheusURL: criblcontrolplanesdkgo.Pointer("https://nocturnal-programme.com"),
            },
        ),
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedOutput != nil {
        // handle response
    }
}
```
### Example Usage: OutputCreateExamplesGraphite

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesGraphite" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyGraphite(
        operations.CreateOutputSystemByPackOutputGraphite{
            ID: "graphite-output",
            Type: operations.CreateOutputSystemByPackTypeGraphiteGraphite,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Protocol: components.DestinationProtocolOptionsTCP,
            Host: "localhost",
            Port: 2003,
            Mtu: criblcontrolplanesdkgo.Pointer[float64](7866.83),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1426.35),
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](8899.18),
            Description: criblcontrolplanesdkgo.Pointer("incidentally inwardly boohoo absent license fibre yet aha"),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](6999.86),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](6874.71),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](468.36),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8559.24),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](238.67),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsGraphite{},
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
### Example Usage: OutputCreateExamplesHoneycomb

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesHoneycomb" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyHoneycomb(
        operations.CreateOutputSystemByPackOutputHoneycomb{
            ID: "honeycomb-output",
            Type: operations.CreateOutputSystemByPackTypeHoneycombHoneycomb,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Dataset: "my-dataset",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](3521.46),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](2922.05),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](160.1),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](4128.27),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1304.16),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Secret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("clonk consequently newsletter that excepting although unlearn"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](7952.29),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2340.96),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8250.64),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsHoneycomb{},
            Team: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesHumioHec

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesHumioHec" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyHumioHec(
        operations.CreateOutputSystemByPackOutputHumioHec{
            ID: "humio-hec-output",
            Type: operations.CreateOutputSystemByPackTypeHumioHecHumioHec,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            URL: "https://cloud.us.humio.com/api/v1/ingest/hec",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](4843.25),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](5786.95),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](1997.07),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](7504.67),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](9338.27),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Format: components.RequestFormatOptionsJSON,
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("gum inasmuch promptly legislature deer caring near operating"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8938.16),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1435),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5502.93),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsHumioHec{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://worst-forager.biz/"),
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
### Example Usage: OutputCreateExamplesInfluxdb

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesInfluxdb" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyInfluxdb(
        operations.CreateOutputSystemByPackOutputInfluxdb{
            ID: "influxdb-output",
            Type: operations.CreateOutputSystemByPackTypeInfluxdbInfluxdb,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            URL: "http://localhost:8086",
            UseV2API: criblcontrolplanesdkgo.Pointer(true),
            TimestampPrecision: operations.CreateOutputSystemByPackTimestampPrecisionS.ToPointer(),
            DynamicValueFieldName: criblcontrolplanesdkgo.Pointer(false),
            ValueFieldName: criblcontrolplanesdkgo.Pointer("<value>"),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](4865.38),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](7713.92),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](1654.04),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](9548.5),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](9122.06),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: operations.CreateOutputSystemByPackAuthenticationTypeInfluxdbToken.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("even after reporter"),
            Database: criblcontrolplanesdkgo.Pointer("mydb"),
            Bucket: criblcontrolplanesdkgo.Pointer("<value>"),
            Org: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](644.64),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4345.91),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8044.48),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsInfluxdb{},
            Username: criblcontrolplanesdkgo.Pointer("Graham_Cummerata"),
            Password: criblcontrolplanesdkgo.Pointer("QoazYQzuIXrYLgy"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://astonishing-disconnection.name"),
            TemplateDatabase: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateBucket: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesKafka

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesKafka" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyKafka(
        operations.CreateOutputSystemByPackOutputKafka{
            ID: "kafka-output",
            Type: operations.CreateOutputSystemByPackTypeKafkaKafka,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Brokers: []string{
                "localhost:9092",
            },
            Topic: "logs",
            Ack: components.AcknowledgmentsOptions1Leader.ToPointer(),
            Format: components.RecordDataFormatOptions1Raw.ToPointer(),
            Compression: components.CompressionOptions3Zstd.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](7314.96),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](2019.51),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3377.23),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType1{
                Disabled: false,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://orderly-knitting.info/"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](65.04),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7620.77),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](1627.05),
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
                DefaultKeySchemaID: criblcontrolplanesdkgo.Pointer[float64](3110.3),
                DefaultValueSchemaID: criblcontrolplanesdkgo.Pointer[float64](4216.12),
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](7181.04),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](3724.78),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](4755.27),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](7144.48),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](6810.09),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](4556.71),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](6671.44),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](5906.1),
            Sasl: &components.AuthenticationType{
                Disabled: true,
                Username: criblcontrolplanesdkgo.Pointer("Jeromy34"),
                Password: criblcontrolplanesdkgo.Pointer("mkbiANI3XTLISi6"),
                AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslScramSha512.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(true),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://definitive-kielbasa.com"),
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
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("than never yum shoddy finger"),
            ProtobufLibraryID: criblcontrolplanesdkgo.Pointer("<id>"),
            ProtobufEncodingID: criblcontrolplanesdkgo.Pointer("<id>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5957.13),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3159),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1778.58),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsKafka{},
            TemplateTopic: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesKinesis

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesKinesis" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyKinesis(
        operations.CreateOutputSystemByPackOutputKinesis{
            ID: "kinesis-output",
            Type: operations.CreateOutputSystemByPackTypeKinesisKinesis,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            StreamName: "my-stream",
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfSecret.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions2V2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](2206.96),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](578.15),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](1353.1),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7828.01),
            Compression: operations.CreateOutputSystemByPackCompressionGzip.ToPointer(),
            UseListShards: criblcontrolplanesdkgo.Pointer(true),
            AsNdjson: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("never quizzically ecliptic"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxEventsPerFlush: criblcontrolplanesdkgo.Pointer[float64](7848.06),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6095.42),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8419.65),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](4700.16),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsKinesis{},
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
    if res.CountedOutput != nil {
        // handle response
    }
}
```
### Example Usage: OutputCreateExamplesLoki

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesLoki" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyLoki(
        operations.CreateOutputSystemByPackOutputLoki{
            ID: "loki-output",
            Type: operations.CreateOutputSystemByPackTypeLokiLoki,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            URL: "http://localhost:3100/loki/api/v1/push",
            Message: criblcontrolplanesdkgo.Pointer("<value>"),
            MessageFormat: components.MessageFormatOptionsProtobuf.ToPointer(),
            Labels: []components.ItemsTypeLabels{
                components.ItemsTypeLabels{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthType: components.AuthenticationTypeOptionsPrometheusAuth1Token.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](200.1),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1452.29),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](8898.77),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](917.42),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](4481.48),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            EnableDynamicHeaders: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](5475.99),
            Description: criblcontrolplanesdkgo.Pointer("pfft chapel swat save phew aside dress"),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Username: criblcontrolplanesdkgo.Pointer("Aliza.Schiller7"),
            Password: criblcontrolplanesdkgo.Pointer("llUzXbfL955jZKg"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1961.13),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](726.31),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6710.87),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsLoki{},
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
### Example Usage: OutputCreateExamplesMicrosoftFabric

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesMicrosoftFabric" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyMicrosoftFabric(
        operations.CreateOutputSystemByPackOutputMicrosoftFabric{
            ID: "microsoft-fabric-output",
            Type: operations.CreateOutputSystemByPackTypeMicrosoftFabricMicrosoftFabric,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Topic: "logs",
            Ack: components.AcknowledgmentsOptionsAll.ToPointer(),
            Format: components.RecordDataFormatOptionsRaw.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](4124.7),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](5573.22),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1504.21),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](4318.81),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](2670.54),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](7669.88),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](3448.63),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2514.2),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](9653.92),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](5832.31),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](8601.04),
            Sasl: &operations.CreateOutputSystemByPackAuthentication{
                Disabled: true,
                Mechanism: components.SaslMechanismOptionsSasl1Oauthbearer.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("Arthur_Pfeffer"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                ClientSecretAuthType: operations.CreateOutputSystemByPackClientSecretAuthTypeAuthenticationMethodSecret.ToPointer(),
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
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            BootstrapServer: "myeventstream.servicebus.windows.net:9093",
            Description: criblcontrolplanesdkgo.Pointer("violent incomparable yahoo pendant diligently square innocently"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](154.96),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9316.56),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](7819.76),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsMicrosoftFabric{},
            TemplateTopic: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateBootstrapServer: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesMinio

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesMinio" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyMinio(
        operations.CreateOutputSystemByPackOutputMinio{
            ID: "minio-output",
            Type: operations.CreateOutputSystemByPackTypeMinioMinio,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Endpoint: "http://localhost:9000",
            Bucket: "my-bucket",
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfSecret.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions5V2.ToPointer(),
            ObjectACL: components.ObjectACLOptionsPublicRead.ToPointer(),
            StorageClass: components.StorageClassOptions2ReducedRedundancy.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionOptionsAes256.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(false),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsParquet.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](2618.15),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](6371.52),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](193.27),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](27.11),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4258.54),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1146.47),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1403.61),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](2803.98),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](3019.56),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](519.06),
            Description: criblcontrolplanesdkgo.Pointer("flat phew abaft junior hunt eyebrow disposer robust"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](587.46),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(true),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(true),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](5832.98),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](6279.33),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](1198.46),
            TemplateBucket: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateFormat: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesMsk

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesMsk" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyMsk(
        operations.CreateOutputSystemByPackOutputMsk{
            ID: "msk-output",
            Type: operations.CreateOutputSystemByPackTypeMskMsk,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Brokers: []string{
                "b-1.example.xxxxx.c2.kafka.us-east-1.amazonaws.com:9092",
            },
            Topic: "logs",
            Ack: components.AcknowledgmentsOptions1None.ToPointer(),
            Format: components.RecordDataFormatOptions1JSON.ToPointer(),
            Compression: components.CompressionOptions3Lz4.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](6028.8),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](5526.35),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3412.48),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType1{
                Disabled: false,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://orderly-knitting.info/"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](65.04),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7620.77),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](1627.05),
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
                DefaultKeySchemaID: criblcontrolplanesdkgo.Pointer[float64](3110.3),
                DefaultValueSchemaID: criblcontrolplanesdkgo.Pointer[float64](4216.12),
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](593.31),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](175.51),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](4501.98),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](9120.57),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1617.95),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](651.48),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](1474.54),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](3357.1),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto,
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsV2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](699.67),
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
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("slake about intend innocently resource"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ProtobufLibraryID: criblcontrolplanesdkgo.Pointer("<id>"),
            ProtobufEncodingID: criblcontrolplanesdkgo.Pointer("<id>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](4002.99),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9919.57),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5946.41),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsMsk{},
            TemplateTopic: criblcontrolplanesdkgo.Pointer("<value>"),
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
    if res.CountedOutput != nil {
        // handle response
    }
}
```
### Example Usage: OutputCreateExamplesNetflow

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesNetflow" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyNetflow(
        operations.CreateOutputSystemByPackOutputNetflow{
            ID: "netflow-output",
            Type: operations.CreateOutputSystemByPackTypeNetflowNetflow,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Hosts: []operations.CreateOutputSystemByPackHostNetflow{
                operations.CreateOutputSystemByPackHostNetflow{
                    Host: "localhost",
                    Port: 2055,
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](5309.41),
            EnableIPSpoofing: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("making sneak store lyre rough like"),
            MaxRecordSize: criblcontrolplanesdkgo.Pointer[float64](9396.95),
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
### Example Usage: OutputCreateExamplesNewrelic

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesNewrelic" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyNewrelic(
        operations.CreateOutputSystemByPackOutputNewrelic{
            ID: "newrelic-output",
            Type: operations.CreateOutputSystemByPackTypeNewrelicNewrelic,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Region: components.RegionOptionsCustom.ToPointer(),
            LogType: criblcontrolplanesdkgo.Pointer("<value>"),
            MessageField: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []operations.CreateOutputSystemByPackMetadatum{
                operations.CreateOutputSystemByPackMetadatum{
                    Name: operations.CreateOutputSystemByPackFieldNameHostname,
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](4757.96),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](6250.48),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](4654.17),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](3745.63),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3355.24),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Manual.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](4813.25),
            Description: criblcontrolplanesdkgo.Pointer("notwithstanding by interchange reservation that downright so cruelly"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://fluffy-bump.biz/"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6694.11),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2562.6),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1412.75),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsNewrelic{},
            APIKey: criblcontrolplanesdkgo.Pointer("your-api-key"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateLogType: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateMessageField: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesNewrelicEvents

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesNewrelicEvents" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyNewrelicEvents(
        operations.CreateOutputSystemByPackOutputNewrelicEvents{
            ID: "newrelic-events-output",
            Type: operations.CreateOutputSystemByPackTypeNewrelicEventsNewrelicEvents,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Region: components.RegionOptionsUs.ToPointer(),
            AccountID: "123456",
            EventType: "CriblEvent",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](83.93),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4622.8),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](6007.1),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](9181.26),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8180.04),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Secret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("yowza ick claw fussy singe swing dirty"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://concrete-possession.org"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](4362.97),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](955.1),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](2757.81),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsNewrelicEvents{},
            APIKey: criblcontrolplanesdkgo.Pointer("your-api-key"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateEventType: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateCustomURL: criblcontrolplanesdkgo.Pointer("https://leading-palate.name"),
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
### Example Usage: OutputCreateExamplesOpenTelemetry

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesOpenTelemetry" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyOpenTelemetry(
        operations.CreateOutputSystemByPackOutputOpenTelemetry{
            ID: "opentelemetry-output",
            Type: operations.CreateOutputSystemByPackTypeOpenTelemetryOpenTelemetry,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Protocol: components.ProtocolOptionsGrpc.ToPointer(),
            Endpoint: "http://localhost:4317",
            OtlpVersion: operations.CreateOutputSystemByPackOTLPVersionOneDot3Dot1.ToPointer(),
            Compress: components.CompressionOptions4Deflate.ToPointer(),
            HTTPCompress: components.CompressionOptions5None.ToPointer(),
            AuthType: components.AuthenticationTypeOptionsTextSecret.ToPointer(),
            HTTPTracesEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPMetricsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPLogsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5277.69),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](8479.36),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](1500.56),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7118.59),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](236.26),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](1152.01),
            KeepAlive: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("lightly merit consequently impossible"),
            Username: criblcontrolplanesdkgo.Pointer("Annamae_Ondricka"),
            Password: criblcontrolplanesdkgo.Pointer("F_ZED65h2bmasxD"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            TLS: &components.TLSSettingsClientSideType2{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1156.32),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5935.54),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5842.73),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsOpenTelemetry{},
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
### Example Usage: OutputCreateExamplesPrometheus

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesPrometheus" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyPrometheus(
        operations.CreateOutputSystemByPackOutputPrometheus{
            ID: "prometheus-output",
            Type: operations.CreateOutputSystemByPackTypePrometheusPrometheus,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            URL: "http://localhost:9091/api/v1/write",
            MetricRenameExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            SendMetadata: criblcontrolplanesdkgo.Pointer(false),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](8539.58),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](5457.42),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](6165.24),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](131.09),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7914.63),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            AuthType: components.AuthenticationTypeOptionsPrometheusAuthNone.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("official free brr pish wherever drat awkwardly yet"),
            MetricsFlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1546.38),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](3024.12),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4777.5),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1256.83),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsPrometheus{},
            Username: criblcontrolplanesdkgo.Pointer("Dina_Dooley"),
            Password: criblcontrolplanesdkgo.Pointer("jWIvX2o_9GpWhf7"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://athletic-zen.com"),
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
### Example Usage: OutputCreateExamplesRing

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesRing" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyRing(
        operations.CreateOutputSystemByPackOutputRing{
            ID: "ring-output",
            Type: operations.CreateOutputSystemByPackTypeRingRing,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Format: operations.CreateOutputSystemByPackDataFormatRingJSON.ToPointer(),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("interesting well-documented casket pfft cassava joyously"),
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
### Example Usage: OutputCreateExamplesRouter

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesRouter" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyRouter(
        operations.CreateOutputSystemByPackOutputRouter{
            ID: "router-output",
            Type: operations.CreateOutputSystemByPackTypeRouterRouter,
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
            Rules: []operations.CreateOutputSystemByPackRule{
                operations.CreateOutputSystemByPackRule{
                    Filter: "true",
                    Output: "my-output",
                    Description: criblcontrolplanesdkgo.Pointer("testify morning secret and rigidly diligently emulsify"),
                    Final: criblcontrolplanesdkgo.Pointer(false),
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("platter yum across"),
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
### Example Usage: OutputCreateExamplesS3

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesS3" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyS3(
        operations.CreateOutputSystemByPackOutputS3{
            ID: "s3-output",
            Type: operations.CreateOutputSystemByPackTypeS3S3,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Bucket: "my-bucket",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](1137.22),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            ObjectACL: components.ObjectACLOptionsAwsExecRead.ToPointer(),
            StorageClass: components.StorageClassOptionsIntelligentTiering.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAwsKms.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsRaw.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](4861.56),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](979.15),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](5009.3),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(true),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](27.11),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4258.54),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1146.47),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1403.61),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](7299.34),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](7678.19),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](4461.04),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](5681.16),
            Description: criblcontrolplanesdkgo.Pointer("acquire unless phooey furthermore blah pish of limply via even"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2None.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV1.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](5602.83),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(true),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(false),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(true),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](4474.24),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](8131.76),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](4149.04),
            TemplateBucket: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateFormat: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesSecurityLake

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesSecurityLake" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodySecurityLake(
        operations.CreateOutputSystemByPackOutputSecurityLake{
            ID: "security-lake-output",
            Type: operations.CreateOutputSystemByPackTypeSecurityLakeSecurityLake,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Bucket: "my-bucket",
            Region: "us-east-1",
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfSecret.ToPointer(),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: operations.CreateOutputSystemByPackSignatureVersionSecurityLakeV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: "arn:aws:iam::123456789012:role/my-role",
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](7869.73),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            ObjectACL: components.ObjectACLOptionsAuthenticatedRead.ToPointer(),
            StorageClass: components.StorageClassOptionsStandardIa.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAes256.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](4079.53),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](2520.16),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](7042.03),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(true),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](27.11),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4258.54),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1146.47),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1403.61),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](161.23),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](9675.63),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](1500.74),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](2660.43),
            AccountID: "123456789012",
            CustomSource: "my-custom-source",
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetVersion: components.ParquetVersionOptionsParquet24.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV1.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](608.83),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(false),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("unlike brisk license"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](6857.05),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1431.06),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](8351.21),
            TemplateBucket: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesSentinel

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesSentinel" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodySentinel(
        operations.CreateOutputSystemByPackOutputSentinel{
            ID: "sentinel-output",
            Type: operations.CreateOutputSystemByPackTypeSentinelSentinel,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            KeepAlive: criblcontrolplanesdkgo.Pointer(false),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1449.12),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4077.75),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](4029.17),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](2469.04),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8583.42),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: operations.CreateOutputSystemByPackAuthTypeOauth.ToPointer(),
            LoginURL: "https://login.microsoftonline.com",
            Secret: "client-secret",
            ClientID: "client-id",
            Scope: criblcontrolplanesdkgo.Pointer("<value>"),
            EndpointURLConfiguration: operations.CreateOutputSystemByPackEndpointConfigurationURL,
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](9753.03),
            Description: criblcontrolplanesdkgo.Pointer("holster oh ha next forecast sheepishly"),
            Format: operations.CreateOutputSystemByPackFormatSentinelCustom.ToPointer(),
            CustomSourceExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomDropWhenNull: criblcontrolplanesdkgo.Pointer(false),
            CustomEventDelimiter: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomContentType: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomPayloadExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            AdvancedContentType: criblcontrolplanesdkgo.Pointer("<value>"),
            FormatEventCode: criblcontrolplanesdkgo.Pointer("<value>"),
            FormatPayloadCode: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6973.94),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9191.66),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6570.19),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsSentinel{},
            URL: criblcontrolplanesdkgo.Pointer("https://your-workspace.ingest.monitor.azure.com"),
            DcrID: criblcontrolplanesdkgo.Pointer("<id>"),
            DceEndpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            StreamName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateLoginURL: criblcontrolplanesdkgo.Pointer("https://grave-amendment.org"),
            TemplateSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateClientID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateScope: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://talkative-legislature.info/"),
            TemplateDcrID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateDceEndpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateStreamName: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesSentinelOneAiSiem

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesSentinelOneAiSiem" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodySentinelOneAiSiem(
        operations.CreateOutputSystemByPackOutputSentinelOneAiSiem{
            ID: "sentinel-one-ai-siem-output",
            Type: operations.CreateOutputSystemByPackTypeSentinelOneAiSiemSentinelOneAiSiem,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Region: operations.CreateOutputSystemByPackRegionUs,
            Endpoint: operations.CreateOutputSystemByPackAISIEMEndpointPathRootServicesCollectorEvent,
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](246.07),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](8345.77),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](6926.81),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](8657.49),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](9009.46),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("decent gosh hotfoot focalise jet"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            BaseURL: criblcontrolplanesdkgo.Pointer("https://golden-guidance.biz/"),
            HostExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            SourceExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            SourceTypeExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceCategoryExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceNameExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceVendorExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            EventTypeExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            Host: criblcontrolplanesdkgo.Pointer("humble-scrap.net"),
            Source: criblcontrolplanesdkgo.Pointer("<value>"),
            SourceType: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceCategory: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceName: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceVendor: criblcontrolplanesdkgo.Pointer("<value>"),
            EventType: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](4392.85),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2026.56),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6084.94),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsSentinelOneAiSiem{},
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
### Example Usage: OutputCreateExamplesServiceNow

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesServiceNow" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyServiceNow(
        operations.CreateOutputSystemByPackOutputServiceNow{
            ID: "servicenow-output",
            Type: operations.CreateOutputSystemByPackTypeServiceNowServiceNow,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Endpoint: "ingest.lightstep.com:443",
            TokenSecret: "your-token-secret",
            AuthTokenName: criblcontrolplanesdkgo.Pointer("<value>"),
            OtlpVersion: components.OtlpVersionOptions1OneDot3Dot1,
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](503.77),
            Protocol: components.ProtocolOptionsHTTP,
            Compress: components.CompressionOptions4Gzip.ToPointer(),
            HTTPCompress: components.CompressionOptions5None.ToPointer(),
            HTTPTracesEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPMetricsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPLogsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](2154.24),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](4368.66),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1760.74),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](4443.02),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](5437.31),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("woefully ick marathon mostly oof amid if inwardly commemorate"),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            TLS: &components.TLSSettingsClientSideType2{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9783.22),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3299.29),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8193.24),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsServiceNow{},
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
### Example Usage: OutputCreateExamplesSignalfx

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesSignalfx" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodySignalfx(
        operations.CreateOutputSystemByPackOutputSignalfx{
            ID: "signalfx-output",
            Type: operations.CreateOutputSystemByPackTypeSignalfxSignalfx,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Realm: "us0",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](6645.7),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](9487.59),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](7248.16),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](9800.19),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3414.66),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("inwardly milestone reopen meanwhile leading radiant hole mature drive cheerfully"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8643.23),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1709.64),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5809.61),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsSignalfx{},
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
### Example Usage: OutputCreateExamplesSnmp

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesSnmp" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodySnmp(
        operations.CreateOutputSystemByPackOutputSnmp{
            ID: "snmp-output",
            Type: operations.CreateOutputSystemByPackTypeSnmpSnmp,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Hosts: []operations.CreateOutputSystemByPackHostSnmp{
                operations.CreateOutputSystemByPackHostSnmp{
                    Host: "192.168.1.1",
                    Port: 161,
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](4878.98),
            Description: criblcontrolplanesdkgo.Pointer("exactly representation officially yuck easily absent jacket"),
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
### Example Usage: OutputCreateExamplesSns

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesSns" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodySns(
        operations.CreateOutputSystemByPackOutputSns{
            ID: "sns-output",
            Type: operations.CreateOutputSystemByPackTypeSnsSns,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            TopicArn: "arn:aws:sns:us-east-1:123456789012:my-topic",
            MessageGroupID: "my-message-group",
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](9220.51),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfManual.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: operations.CreateOutputSystemByPackSignatureVersionSnsV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](6435.8),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("supposing oh hastily give fortunately hammock"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2030.75),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4614.55),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](874.72),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsSns{},
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
    if res.CountedOutput != nil {
        // handle response
    }
}
```
### Example Usage: OutputCreateExamplesSplunk

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesSplunk" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodySplunk(
        operations.CreateOutputSystemByPackOutputSplunk{
            ID: "splunk-output",
            Type: operations.CreateOutputSystemByPackTypeSplunkSplunk,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Host: "localhost",
            Port: 9997,
            NestedFields: components.NestedFieldSerializationOptionsJSON.ToPointer(),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](5454.45),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](2267.18),
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
            EnableMultiMetrics: criblcontrolplanesdkgo.Pointer(true),
            EnableACK: criblcontrolplanesdkgo.Pointer(true),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(false),
            MaxS2Sversion: components.MaxS2SVersionOptionsV3.ToPointer(),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("certification yum likely throughout seriously stupendous jellyfish certification lazily better"),
            MaxFailedHealthChecks: criblcontrolplanesdkgo.Pointer[float64](8593.9),
            Compress: components.CompressionOptionsAuto.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5137.13),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3578.55),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](626.61),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsSplunk{},
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
### Example Usage: OutputCreateExamplesSplunkHec

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesSplunkHec" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodySplunkHec(
        operations.CreateOutputSystemByPackOutputSplunkHec{
            ID: "splunk-hec-output",
            Type: operations.CreateOutputSystemByPackTypeSplunkHecSplunkHec,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            LoadBalanced: criblcontrolplanesdkgo.Pointer(true),
            NextQueue: criblcontrolplanesdkgo.Pointer("<value>"),
            TCPRouting: criblcontrolplanesdkgo.Pointer("<value>"),
            TLS: &components.TLSSettingsClientSideType1{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](3689.15),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](6879.41),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](4653.04),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](8311.1),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6170.31),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            EnableMultiMetrics: criblcontrolplanesdkgo.Pointer(true),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("querulous painfully what beside while unlucky gee overstay"),
            URL: criblcontrolplanesdkgo.Pointer("https://orderly-cruelty.net/"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(true),
            Urls: []operations.CreateOutputSystemByPackURLSplunkHec{
                operations.CreateOutputSystemByPackURLSplunkHec{
                    URL: "https://illustrious-permafrost.net",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](6169.06),
                    TemplateURL: criblcontrolplanesdkgo.Pointer("https://careless-casement.com"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](7364.12),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](4627.82),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1872.69),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9087.95),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5363.05),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsSplunkHec{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://breakable-championship.net"),
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
### Example Usage: OutputCreateExamplesSplunkLb

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesSplunkLb" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodySplunkLb(
        operations.CreateOutputSystemByPackOutputSplunkLb{
            ID: "splunk-lb-output",
            Type: operations.CreateOutputSystemByPackTypeSplunkLbSplunkLb,
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
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](7056.71),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7257.28),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](8308.87),
            NestedFields: components.NestedFieldSerializationOptionsJSON.ToPointer(),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](4504.9),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](3216.79),
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
            EnableMultiMetrics: criblcontrolplanesdkgo.Pointer(false),
            EnableACK: criblcontrolplanesdkgo.Pointer(false),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(true),
            MaxS2Sversion: components.MaxS2SVersionOptionsV3.ToPointer(),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            IndexerDiscovery: criblcontrolplanesdkgo.Pointer(false),
            SenderUnhealthyTimeAllowance: criblcontrolplanesdkgo.Pointer[float64](9874.35),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("impractical powerfully pulverize round zowie"),
            MaxFailedHealthChecks: criblcontrolplanesdkgo.Pointer[float64](7299.56),
            Compress: components.CompressionOptionsDisabled.ToPointer(),
            IndexerDiscoveryConfigs: &operations.CreateOutputSystemByPackIndexerDiscoveryConfigs{
                Site: "<value>",
                MasterURI: "https://mild-bonfire.biz",
                RefreshIntervalSec: 481.47,
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                AuthTokens: []operations.CreateOutputSystemByPackAuthToken{
                    operations.CreateOutputSystemByPackAuthToken{
                        AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
                        AuthToken: criblcontrolplanesdkgo.Pointer("<value>"),
                        TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    },
                },
                AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
                AuthToken: criblcontrolplanesdkgo.Pointer("<value>"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(true),
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "localhost",
                    Port: 9997,
                    TLS: components.TLSOptionsHostsItemsInherit.ToPointer(),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    Weight: criblcontrolplanesdkgo.Pointer[float64](7243.49),
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](3109.58),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3899.85),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](4967.6),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsSplunkLb{},
            AuthToken: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesSqs

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesSqs" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodySqs(
        operations.CreateOutputSystemByPackOutputSqs{
            ID: "sqs-output",
            Type: operations.CreateOutputSystemByPackTypeSqsSqs,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            QueueName: "my-queue",
            QueueType: operations.CreateOutputSystemByPackQueueTypeStandard,
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            MessageGroupID: criblcontrolplanesdkgo.Pointer("<id>"),
            CreateQueue: criblcontrolplanesdkgo.Pointer(false),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfManual.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions3V2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](7240.49),
            MaxQueueSize: criblcontrolplanesdkgo.Pointer[float64](7569.17),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](7532.91),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](9657.12),
            MaxInProgress: criblcontrolplanesdkgo.Pointer[float64](6097.73),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("ecstatic opposite or around phew ouch whether whose consequently versus"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](41.51),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4240.87),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](9550.59),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsSqs{},
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
    if res.CountedOutput != nil {
        // handle response
    }
}
```
### Example Usage: OutputCreateExamplesStatsd

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesStatsd" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyStatsd(
        operations.CreateOutputSystemByPackOutputStatsd{
            ID: "statsd-output",
            Type: operations.CreateOutputSystemByPackTypeStatsdStatsd,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Protocol: components.DestinationProtocolOptionsUDP,
            Host: "localhost",
            Port: 8125,
            Mtu: criblcontrolplanesdkgo.Pointer[float64](8881.31),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1139.86),
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](2858.69),
            Description: criblcontrolplanesdkgo.Pointer("scuttle upright well-off weary zowie procurement after debit"),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](2682.71),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](1267.8),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2479.41),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7900.85),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](2414),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsStatsd{},
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
### Example Usage: OutputCreateExamplesStatsdExt

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesStatsdExt" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyStatsdExt(
        operations.CreateOutputSystemByPackOutputStatsdExt{
            ID: "statsd-ext-output",
            Type: operations.CreateOutputSystemByPackTypeStatsdExtStatsdExt,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Protocol: components.DestinationProtocolOptionsUDP,
            Host: "localhost",
            Port: 8125,
            Mtu: criblcontrolplanesdkgo.Pointer[float64](1819.2),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](4621.8),
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](5593.14),
            Description: criblcontrolplanesdkgo.Pointer("consequently playfully hoick realistic"),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](2291.64),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](1484.48),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6506.07),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5507.17),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](2725.91),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsStatsdExt{},
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
### Example Usage: OutputCreateExamplesSumoLogic

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesSumoLogic" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodySumoLogic(
        operations.CreateOutputSystemByPackOutputSumoLogic{
            ID: "sumo-logic-output",
            Type: operations.CreateOutputSystemByPackTypeSumoLogicSumoLogic,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            URL: "https://endpoint1.collection.us2.sumologic.com",
            CustomSource: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomCategory: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: operations.CreateOutputSystemByPackDataFormatSumoLogicJSON.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1424.31),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](922.32),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](5993.15),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](1717.56),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](2003.39),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](5961.13),
            Description: criblcontrolplanesdkgo.Pointer("down worriedly at around easily or inasmuch outbid"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](3481.25),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2319.77),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](4362.24),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsSumoLogic{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://understated-sediment.info"),
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
### Example Usage: OutputCreateExamplesSyslog

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesSyslog" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodySyslog(
        operations.CreateOutputSystemByPackOutputSyslog{
            ID: "syslog-output",
            Type: operations.CreateOutputSystemByPackTypeSyslogSyslog,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Protocol: operations.CreateOutputSystemByPackProtocolSyslogTCP.ToPointer(),
            Facility: operations.CreateOutputSystemByPackFacilitySix.ToPointer(),
            Severity: operations.CreateOutputSystemByPackSeveritySyslogDebug.ToPointer(),
            AppName: criblcontrolplanesdkgo.Pointer("<value>"),
            MessageFormat: operations.CreateOutputSystemByPackMessageFormatRfc5424.ToPointer(),
            TimestampFormat: operations.CreateOutputSystemByPackTimestampFormatSyslog.ToPointer(),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            OctetCountFraming: criblcontrolplanesdkgo.Pointer(true),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("fairly pleasure jiggle"),
            LoadBalanced: criblcontrolplanesdkgo.Pointer(true),
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](514),
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
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](6656.86),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](4768.36),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](748.42),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](9832.65),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](1279.92),
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
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            MaxRecordSize: criblcontrolplanesdkgo.Pointer[float64](6152.25),
            UDPDNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](3762.93),
            EnableIPSpoofing: criblcontrolplanesdkgo.Pointer(true),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5395.19),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7289.32),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](4600.99),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsSyslog{},
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
### Example Usage: OutputCreateExamplesTcpjson

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesTcpjson" -->
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
### Example Usage: OutputCreateExamplesWavefront

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesWavefront" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyWavefront(
        operations.CreateOutputSystemByPackOutputWavefront{
            ID: "wavefront-output",
            Type: operations.CreateOutputSystemByPackTypeWavefrontWavefront,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Domain: "longboard",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](8786.17),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](2965.79),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](3712.73),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](3372.05),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](2182.98),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("through yieldingly inquisitively pfft readies whenever minus louse alliance"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1288.03),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2799.01),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](2440.55),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsWavefront{},
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
### Example Usage: OutputCreateExamplesWebhook

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesWebhook" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyWebhook(
        operations.CreateOutputSystemByPackOutputWebhook{
            ID: "webhook-output",
            Type: operations.CreateOutputSystemByPackTypeWebhookWebhook,
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
            Method: components.MethodOptionsPut.ToPointer(),
            Format: operations.CreateOutputSystemByPackFormatWebhookAdvanced.ToPointer(),
            KeepAlive: criblcontrolplanesdkgo.Pointer(false),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](9520.49),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](5120.47),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](1110.11),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](9465.68),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1261.02),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            AuthType: operations.CreateOutputSystemByPackAuthenticationTypeWebhookTextSecret.ToPointer(),
            TLS: &components.TLSSettingsClientSideType1{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](8639.61),
            LoadBalanced: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("gosh lobotomise weakly contravene print"),
            CustomSourceExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomDropWhenNull: criblcontrolplanesdkgo.Pointer(false),
            CustomEventDelimiter: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomContentType: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomPayloadExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            AdvancedContentType: criblcontrolplanesdkgo.Pointer("<value>"),
            FormatEventCode: criblcontrolplanesdkgo.Pointer("<value>"),
            FormatPayloadCode: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6552.81),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2560.36),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](4650.34),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsWebhook{},
            Username: criblcontrolplanesdkgo.Pointer("Icie69"),
            Password: criblcontrolplanesdkgo.Pointer("9AbOW9tqERCfhMp"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://unaware-fowl.info"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](3789.51),
            OauthParams: []operations.CreateOutputSystemByPackOauthParam{
                operations.CreateOutputSystemByPackOauthParam{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            OauthHeaders: []operations.CreateOutputSystemByPackOauthHeader{
                operations.CreateOutputSystemByPackOauthHeader{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            URL: criblcontrolplanesdkgo.Pointer("https://example.com/webhook"),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Urls: []operations.CreateOutputSystemByPackURLWebhook{
                operations.CreateOutputSystemByPackURLWebhook{
                    URL: "https://gloomy-publication.com/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1833.76),
                    TemplateURL: criblcontrolplanesdkgo.Pointer("https://courageous-saloon.net/"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](2546.17),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](2294.07),
            TemplateLoginURL: criblcontrolplanesdkgo.Pointer("https://fake-louse.biz/"),
            TemplateSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://submissive-heartache.name/"),
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
### Example Usage: OutputCreateExamplesWizHec

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesWizHec" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyWizHec(
        operations.CreateOutputSystemByPackOutputWizHec{
            ID: "wiz-hec-output",
            Type: operations.CreateOutputSystemByPackTypeWizHecWizHec,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            LoadBalanced: "<value>",
            NextQueue: criblcontrolplanesdkgo.Pointer("<value>"),
            TCPRouting: criblcontrolplanesdkgo.Pointer("<value>"),
            TLS: &components.TLSSettingsClientSideType1{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](4222.91),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](2163.92),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](5534.34),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](1901),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5774.63),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            EnableMultiMetrics: "<value>",
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            WizConnectorID: "00000000-0000-0000-0000-000000000000",
            WizEnvironment: "test",
            DataCenter: "us1",
            WizSourcetype: "placeholder",
            Description: criblcontrolplanesdkgo.Pointer("excluding hydrolyze eyeglasses ha kookily zowie home blah during battle"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8687.97),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5698.85),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](465.8),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsWizHec{},
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateWizEnvironment: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateDataCenter: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateWizSourcetype: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesXsiam

<!-- UsageSnippet language="go" operationID="createOutputSystemByPack" method="post" path="/p/{pack}/system/outputs" example="OutputCreateExamplesXsiam" -->
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

    res, err := s.Packs.Destinations.Create(ctx, "<value>", operations.CreateCreateOutputSystemByPackRequestBodyXsiam(
        operations.CreateOutputSystemByPackOutputXsiam{
            ID: "xsiam-output",
            Type: operations.CreateOutputSystemByPackTypeXsiamXsiam,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            LoadBalanced: criblcontrolplanesdkgo.Pointer(false),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5763.61),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](6925.51),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](4318.09),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](9888.04),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6723.71),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            AuthType: operations.CreateOutputSystemByPackAuthenticationMethodXsiamSecret.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8177.04,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](7866.98),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6178.16),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3348.05),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: true,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9325.54),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3205.55),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](1352.21),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            ThrottleRateReqPerSec: criblcontrolplanesdkgo.Pointer[int64](188000),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](664.63),
            Description: criblcontrolplanesdkgo.Pointer("soliloquy abnegate numb"),
            URL: criblcontrolplanesdkgo.Pointer("https://fatal-typewriter.net"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(true),
            Urls: []operations.CreateOutputSystemByPackURLXsiam{
                operations.CreateOutputSystemByPackURLXsiam{
                    URL: "https://sour-willow.net",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](2247.73),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](4453.39),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](2927.5),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5713.34),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7532.3),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6857.66),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputSystemByPackPqControlsXsiam{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://acceptable-napkin.biz/"),
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

### Example Usage: OutputCreateExamplesAzureBlob

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesAzureBlob" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputAzureBlob(
        components.OutputAzureBlob{
            ID: criblcontrolplanesdkgo.Pointer("azure-blob-output"),
            Type: components.OutputAzureBlobTypeAzureBlob,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            ContainerName: "my-container",
            CreateContainer: criblcontrolplanesdkgo.Pointer(false),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](9626.81),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsParquet.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](1190.21),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](5799.8),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](2643.71),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](4914.27),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](460.88),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](9741.71),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7684.82),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1991.93),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](5686.77),
            },
            AuthType: components.AuthenticationMethodOptionsClientCert.ToPointer(),
            StorageClass: components.BlobAccessTierHot.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("outside duh publicize neatly unto nun"),
            Compress: components.CompressionOptions2None.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](8090.72),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(true),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(true),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](3115.77),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](5180.41),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](7908.54),
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
            TemplateContainerName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateFormat: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateConnectionString: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateTenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateClientID: criblcontrolplanesdkgo.Pointer("<id>"),
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
### Example Usage: OutputCreateExamplesAzureDataExplorer

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesAzureDataExplorer" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputAzureDataExplorer(
        components.OutputAzureDataExplorer{
            ID: criblcontrolplanesdkgo.Pointer("azure-data-explorer-output"),
            Type: components.OutputAzureDataExplorerTypeAzureDataExplorer,
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
            ClusterURL: "https://mycluster.kusto.windows.net",
            Database: "mydatabase",
            Table: "mytable",
            ValidateDatabaseSettings: criblcontrolplanesdkgo.Pointer(false),
            IngestMode: components.IngestionModeStreaming.ToPointer(),
            OauthEndpoint: components.MicrosoftEntraIDAuthenticationEndpointOptionsSaslHTTPSLoginMicrosoftonlineCom,
            TenantID: "tenant-id",
            ClientID: "client-id",
            Scope: "https://mycluster.kusto.windows.net/.default",
            OauthType: components.OutputAzureDataExplorerAuthenticationMethodClientSecret,
            Description: criblcontrolplanesdkgo.Pointer("judgementally gee now phooey"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("client-secret"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Certificate: &components.Certificate{
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            Format: components.DataFormatOptionsJSON.ToPointer(),
            Compress: components.CompressionOptions2Gzip,
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV1.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](4890.71),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(false),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](7222.75),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](6416.36),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](2049.99),
            IsMappingObj: criblcontrolplanesdkgo.Pointer(false),
            MappingObj: criblcontrolplanesdkgo.Pointer("<value>"),
            MappingRef: criblcontrolplanesdkgo.Pointer("<value>"),
            IngestURL: criblcontrolplanesdkgo.Pointer("https://scornful-premier.info"),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            StagePath: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](4304.04),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](5478.64),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](5753.43),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](4384.79),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](1628.67),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](9741.71),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7684.82),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1991.93),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](5686.77),
            },
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](5395.45),
            FlushImmediately: criblcontrolplanesdkgo.Pointer(true),
            RetainBlobOnSuccess: criblcontrolplanesdkgo.Pointer(false),
            ExtentTags: []components.ExtentTag{
                components.ExtentTag{
                    Prefix: components.PrefixOptionalIngestBy.ToPointer(),
                    Value: "<value>",
                },
            },
            IngestIfNotExists: []components.IngestIfNotExist{
                components.IngestIfNotExist{
                    Value: "<value>",
                },
            },
            ReportLevel: components.ReportLevelFailuresOnly.ToPointer(),
            ReportMethod: components.ReportMethodTable.ToPointer(),
            AdditionalProperties: []components.AdditionalProperty{
                components.AdditionalProperty{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](8457.32),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](6553.47),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](1458.55),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6284.6),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5229.32),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3774.4),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](9855.68),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputAzureDataExplorerPqControls{},
            TemplateClusterURL: criblcontrolplanesdkgo.Pointer("https://yellowish-trick.net/"),
            TemplateDatabase: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateTable: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateTenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateClientID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateScope: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateFormat: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateIngestURL: criblcontrolplanesdkgo.Pointer("https://blushing-colonialism.name"),
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
### Example Usage: OutputCreateExamplesAzureEventhub

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesAzureEventhub" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputAzureEventhub(
        components.OutputAzureEventhub{
            ID: criblcontrolplanesdkgo.Pointer("azure-eventhub-output"),
            Type: components.OutputAzureEventhubTypeAzureEventhub,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Brokers: []string{
                "myeventhub.servicebus.windows.net:9093",
            },
            Topic: "logs",
            Ack: components.AcknowledgmentsOptionsNone.ToPointer(),
            Format: components.RecordDataFormatOptionsRaw.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](7301.52),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](5790.81),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6221.86),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](6833.13),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](2611.24),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](7039.15),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](4527.16),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](4140.17),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5925.09),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](3354.64),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](8673.22),
            Sasl: &components.AuthenticationType1{
                Disabled: false,
                AuthType: components.AuthenticationMethodOptionsSasl1Secret.ToPointer(),
                Password: criblcontrolplanesdkgo.Pointer("7UquuWKnI4T8yEp"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSasl1Oauthbearer.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("Otha.Wolf36"),
                ClientSecretAuthType: components.AuthenticationMethodOptionsSasl2Certificate.ToPointer(),
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
                Disabled: true,
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("bind unethically behind velocity"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9002.1),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4029.74),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5403.8),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputAzureEventhubPqControls{},
            TemplateTopic: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesAzureLogs

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesAzureLogs" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputAzureLogs(
        components.OutputAzureLogs{
            ID: criblcontrolplanesdkgo.Pointer("azure-logs-output"),
            Type: components.OutputAzureLogsTypeAzureLogs,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            LogType: "Cribl",
            ResourceID: criblcontrolplanesdkgo.Pointer("<id>"),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](7445.48),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](9222.37),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](9758.33),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](9069.11),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8626.46),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            APIURL: criblcontrolplanesdkgo.Pointer("https://delicious-assist.org/"),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            AuthType: components.OutputAzureLogsAuthenticationMethodManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("while sleet yahoo subsidy drat ripe moralise meh clonk"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9055.24),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3094.09),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8231.22),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputAzureLogsPqControls{},
            WorkspaceID: criblcontrolplanesdkgo.Pointer("workspace-id"),
            WorkspaceKey: criblcontrolplanesdkgo.Pointer("workspace-key"),
            KeypairSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateWorkspaceID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateWorkspaceKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesChronicle

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesChronicle" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputChronicle(
        components.OutputChronicle{
            ID: criblcontrolplanesdkgo.Pointer("chronicle-output"),
            Type: components.OutputChronicleTypeChronicle,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            APIVersion: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthenticationMethod: components.OutputChronicleAuthenticationMethodServiceAccountSecret.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            Region: "us",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1591.96),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](360.82),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](669.48),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](917.86),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3344.53),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](5872.26),
            IngestionMethod: criblcontrolplanesdkgo.Pointer("<value>"),
            Namespace: criblcontrolplanesdkgo.Pointer("<value>"),
            LogType: "UNKNOWN",
            LogTextField: criblcontrolplanesdkgo.Pointer("<value>"),
            GcpProjectID: "my-project",
            GcpInstance: "customer-id",
            CustomLabels: []components.CustomLabel{
                components.CustomLabel{
                    Key: "<key>",
                    Value: "<value>",
                    RbacEnabled: criblcontrolplanesdkgo.Pointer(true),
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("quit regulate deliberately if far oh a"),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            ServiceAccountCredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6992.05),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9241.54),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](7581.91),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputChroniclePqControls{},
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesClickHouse

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesClickHouse" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputClickHouse(
        components.OutputClickHouse{
            ID: criblcontrolplanesdkgo.Pointer("clickhouse-output"),
            Type: components.OutputClickHouseTypeClickHouse,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            URL: "http://localhost:8123/",
            AuthType: components.OutputClickHouseAuthenticationTypeNone.ToPointer(),
            Database: "mydb",
            TableName: "mytable",
            Format: components.OutputClickHouseFormatJSONCompactEachRowWithNames.ToPointer(),
            MappingType: components.MappingTypeAutomatic.ToPointer(),
            AsyncInserts: criblcontrolplanesdkgo.Pointer(false),
            TLS: &components.TLSSettingsClientSideType1{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](6496.06),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](6754.34),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](2144.97),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](8229.2),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5763.32),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            DumpFormatErrorsToDisk: criblcontrolplanesdkgo.Pointer(true),
            StatsDestination: &components.StatsDestination{
                URL: criblcontrolplanesdkgo.Pointer("https://apprehensive-exasperation.name"),
                Database: criblcontrolplanesdkgo.Pointer("<value>"),
                TableName: criblcontrolplanesdkgo.Pointer("<value>"),
                AuthType: criblcontrolplanesdkgo.Pointer("<value>"),
                Username: criblcontrolplanesdkgo.Pointer("Jett.Waelchi4"),
                SQLUsername: criblcontrolplanesdkgo.Pointer("<value>"),
                Password: criblcontrolplanesdkgo.Pointer("w2NGyozKB5spoFh"),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("yet impish keenly insecure instructor dead perfectly"),
            Username: criblcontrolplanesdkgo.Pointer("Lulu.Spinka46"),
            Password: criblcontrolplanesdkgo.Pointer("NZUA9GkXcdy50id"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            SQLUsername: criblcontrolplanesdkgo.Pointer("<value>"),
            WaitForAsyncInserts: criblcontrolplanesdkgo.Pointer(true),
            ExcludeMappingFields: []string{
                "<value 1>",
            },
            DescribeTable: criblcontrolplanesdkgo.Pointer("<value>"),
            ColumnMappings: []components.ColumnMapping{
                components.ColumnMapping{
                    ColumnName: "<value>",
                    ColumnType: criblcontrolplanesdkgo.Pointer("<value>"),
                    ColumnValueExpression: "<value>",
                },
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](4194.7),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3660.82),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6559.7),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputClickHousePqControls{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://standard-pine.info/"),
            TemplateDatabase: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateTableName: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesCloudflareR2

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesCloudflareR2" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputCloudflareR2(
        components.OutputCloudflareR2{
            ID: criblcontrolplanesdkgo.Pointer("cloudflare-r2-output"),
            Type: components.OutputCloudflareR2TypeCloudflareR2,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Endpoint: "https://account-id.r2.cloudflarestorage.com",
            Bucket: "my-bucket",
            AwsAuthenticationMethod: components.OutputCloudflareR2AuthenticationMethodAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "<value>",
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions5V4.ToPointer(),
            ObjectACL: "<value>",
            StorageClass: components.StorageClassOptions2Standard.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionOptionsAes256.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsParquet.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](1251.74),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](2969.43),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](8826.77),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](9741.71),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7684.82),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1991.93),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](5686.77),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](8446.52),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](2426.05),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](5933.43),
            Description: criblcontrolplanesdkgo.Pointer("amid loosely linear"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV1.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](6645.34),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(false),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](3135.25),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1154.61),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](185.03),
            TemplateBucket: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateFormat: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesCloudwatch

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesCloudwatch" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputCloudwatch(
        components.OutputCloudwatch{
            ID: criblcontrolplanesdkgo.Pointer("cloudwatch-output"),
            Type: components.OutputCloudwatchTypeCloudwatch,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            LogGroupName: "my-log-group",
            LogStreamName: "my-log-stream",
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfSecret.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](4520.57),
            MaxQueueSize: criblcontrolplanesdkgo.Pointer[float64](4309.43),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](9350.28),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6984.69),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("birdcage fooey sternly whether majestically excluding"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5109.17),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1085.97),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](9827),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputCloudwatchPqControls{},
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
    if res.CountedOutput != nil {
        // handle response
    }
}
```
### Example Usage: OutputCreateExamplesConfluentCloud

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesConfluentCloud" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputConfluentCloud(
        components.OutputConfluentCloud{
            ID: criblcontrolplanesdkgo.Pointer("confluent-cloud-output"),
            Type: components.OutputConfluentCloudTypeConfluentCloud,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Brokers: []string{
                "pkc-xxxxx.us-east-1.aws.confluent.cloud:9092",
            },
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
            },
            Topic: "logs",
            Ack: components.AcknowledgmentsOptions1All.ToPointer(),
            Format: components.RecordDataFormatOptions1JSON.ToPointer(),
            Compression: components.CompressionOptions3Snappy.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](5407.63),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](715.33),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5546.59),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType1{
                Disabled: true,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://lighthearted-dredger.org"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](7210.32),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7627.12),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](7494.42),
                Auth: &components.AuthTypeKafkaSchemaRegistry{
                    Disabled: false,
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                    Disabled: criblcontrolplanesdkgo.Pointer(true),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                    CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
                },
                DefaultKeySchemaID: criblcontrolplanesdkgo.Pointer[float64](5091.61),
                DefaultValueSchemaID: criblcontrolplanesdkgo.Pointer[float64](2259.59),
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](6448.17),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7558.24),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](3448.64),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](2938.63),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5846.5),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](4656.1),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](9684.87),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](4341.07),
            Sasl: &components.AuthenticationType{
                Disabled: false,
                Username: criblcontrolplanesdkgo.Pointer("Vaughn_Botsford"),
                Password: criblcontrolplanesdkgo.Pointer("J3to4z1Pba3wD7B"),
                AuthType: components.AuthenticationMethodOptionsSaslSecret.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslScramSha256.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(false),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://runny-saw.name"),
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
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("astride impressionable trial for stake godfather gosh aside kindly supposing"),
            ProtobufLibraryID: criblcontrolplanesdkgo.Pointer("<id>"),
            ProtobufEncodingID: criblcontrolplanesdkgo.Pointer("<id>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6978.04),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7228.5),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](3343.79),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputConfluentCloudPqControls{},
            TemplateTopic: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesCriblHttp

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesCriblHttp" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputCriblHTTP(
        components.OutputCriblHTTP{
            ID: criblcontrolplanesdkgo.Pointer("cribl-http-output"),
            Type: components.OutputCriblHTTPTypeCriblHTTP,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            LoadBalanced: criblcontrolplanesdkgo.Pointer(true),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
            },
            TokenTTLMinutes: criblcontrolplanesdkgo.Pointer[float64](3197.79),
            ExcludeFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Compression: components.CompressionOptions1Gzip.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](3056.71),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](3747.12),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](7787.01),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](3082.83),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8834.82),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            AuthTokens: []components.ItemsTypeAuthTokens1{
                components.ItemsTypeAuthTokens1{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                    Description: criblcontrolplanesdkgo.Pointer("tenement until where pfft anti closely quarterly"),
                },
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("modulo firsthand appropriate impressionable unnaturally label"),
            URL: criblcontrolplanesdkgo.Pointer("https://shrill-cake.net/"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Urls: []components.ItemsTypeUrls{
                components.ItemsTypeUrls{
                    URL: "https://unkempt-charlatan.net/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](7027.74),
                    TemplateURL: criblcontrolplanesdkgo.Pointer("https://married-chiffonier.net"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](2705.77),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](2439.7),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9095.75),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4144.65),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](543.85),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputCriblHTTPPqControls{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://silky-councilman.com/"),
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
### Example Usage: OutputCreateExamplesCriblLake

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesCriblLake" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputCriblLake(
        components.OutputCriblLake{
            ID: criblcontrolplanesdkgo.Pointer("cribl-lake-output"),
            Type: components.OutputCriblLakeTypeCriblLake,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Bucket: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](8301.38),
            StagePath: criblcontrolplanesdkgo.Pointer("<value>"),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            ObjectACL: components.ObjectACLOptionsAuthenticatedRead.ToPointer(),
            StorageClass: components.StorageClassOptionsDeepArchive.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAwsKms.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](6872.45),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](922.06),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](2405.54),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](9741.71),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7684.82),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1991.93),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](5686.77),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](8321.52),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](3354.55),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(false),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](8402.1),
            AwsAuthenticationMethod: components.AwsAuthenticationMethodOptionsAuto.ToPointer(),
            Format: components.FormatOptionsDdss.ToPointer(),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](5969.66),
            Description: criblcontrolplanesdkgo.Pointer("excluding now mallard reorganisation grok busily"),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](1546.23),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](9248.69),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](3518.28),
            TemplateBucket: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateDestPath: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesCriblSearchEngine

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesCriblSearchEngine" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputCriblSearchEngine(
        components.OutputCriblSearchEngine{
            ID: criblcontrolplanesdkgo.Pointer("cribl-search-engine-output"),
            Type: components.OutputCriblSearchEngineTypeCriblSearchEngine,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "cribl_pipe",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{},
            LoadBalanced: criblcontrolplanesdkgo.Pointer(false),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            TokenTTLMinutes: criblcontrolplanesdkgo.Pointer[float64](60),
            ExcludeFields: []string{
                "__kube_*",
                "__metadata",
                "__winEvent",
            },
            Compression: components.CompressionOptions1Gzip.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            SafeHeaders: []string{},
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 401,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](20000),
                },
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 403,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](20000),
                },
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 408,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](250),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 429,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 500,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](250),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 502,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](250),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 503,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](250),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 504,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](250),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 509,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](250),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9897.86),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5067.51),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](90.87),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            AuthTokens: []components.ItemsTypeAuthTokens1{
                components.ItemsTypeAuthTokens1{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                    Description: criblcontrolplanesdkgo.Pointer("tenement until where pfft anti closely quarterly"),
                },
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("boohoo beneath scamper"),
            URL: criblcontrolplanesdkgo.Pointer("https://0.0.0.0:10200"),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Urls: []components.ItemsTypeUrls{
                components.ItemsTypeUrls{
                    URL: "https://unkempt-charlatan.net/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](7027.74),
                    TemplateURL: criblcontrolplanesdkgo.Pointer("https://married-chiffonier.net"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](494.5),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1392.24),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5384.36),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9654.22),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](38.09),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputCriblSearchEnginePqControls{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://puzzled-cope.org/"),
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
### Example Usage: OutputCreateExamplesCriblTcp

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesCriblTcp" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputCriblTCP(
        components.OutputCriblTCP{
            ID: criblcontrolplanesdkgo.Pointer("cribl-tcp-output"),
            Type: components.OutputCriblTCPTypeCriblTCP,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            LoadBalanced: criblcontrolplanesdkgo.Pointer(true),
            Compression: components.CompressionOptions1None.ToPointer(),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(true),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](4499.75),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](9083.6),
            TokenTTLMinutes: criblcontrolplanesdkgo.Pointer[float64](3194.42),
            AuthTokens: []components.ItemsTypeAuthTokens{
                components.ItemsTypeAuthTokens{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("by hence why step-mother bleakly thankfully whose prohibition whisper"),
                },
            },
            ExcludeFields: []string{
                "<value 1>",
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("into enormously mmm uh-huh rudely mispronounce pigpen ew huzzah"),
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](10090),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "webbed-import.com",
                    Port: 5669.2,
                    TLS: components.TLSOptionsHostsItemsOff.ToPointer(),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    Weight: criblcontrolplanesdkgo.Pointer[float64](8579.92),
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](4679.8),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](4291.78),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](1845.71),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5642.93),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](50.68),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6764.87),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputCriblTCPPqControls{},
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
### Example Usage: OutputCreateExamplesCrowdstrikeNextGenSiem

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesCrowdstrikeNextGenSiem" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputCrowdstrikeNextGenSiem(
        components.OutputCrowdstrikeNextGenSiem{
            ID: criblcontrolplanesdkgo.Pointer("crowdstrike-next-gen-siem-output"),
            Type: components.OutputCrowdstrikeNextGenSiemTypeCrowdstrikeNextGenSiem,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            URL: "https://ingest.us.crowdstrike.com/api/ingest/hec/connection-id/v1/services/collector",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1332.31),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](8421.11),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](4372.03),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](9653.19),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6146.45),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Format: components.RequestFormatOptionsJSON,
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("huzzah eek enfold"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9665.01),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8335.23),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](4317.26),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputCrowdstrikeNextGenSiemPqControls{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://well-worn-finer.net/"),
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
### Example Usage: OutputCreateExamplesDatabricks

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesDatabricks" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputDatabricks(
        components.OutputDatabricks{
            ID: criblcontrolplanesdkgo.Pointer("databricks-output"),
            Type: components.OutputDatabricksTypeDatabricks,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            StagePath: criblcontrolplanesdkgo.Pointer("<value>"),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsRaw.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](970.54),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](9610.63),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](7268.2),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](7514.57),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](9064.92),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](9741.71),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7684.82),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1991.93),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](5686.77),
            },
            WorkspaceID: "your-workspace-id",
            Scope: "my-scope",
            ClientID: "your-client-id",
            Catalog: "my-catalog",
            Schema: "my-schema",
            EventsVolumeName: "my-volume",
            ClientTextSecret: "your-client-secret",
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](371.07),
            Description: criblcontrolplanesdkgo.Pointer("hubris eulogise react finer clamour pigpen fervently"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet24.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](8628.52),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(true),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](9170.98),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](4426.15),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](8094.33),
            TemplateFormat: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesDatadog

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesDatadog" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputDatadog(
        components.OutputDatadog{
            ID: criblcontrolplanesdkgo.Pointer("datadog-output"),
            Type: components.OutputDatadogTypeDatadog,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            ContentType: components.SendLogsAsJSON.ToPointer(),
            Message: criblcontrolplanesdkgo.Pointer("<value>"),
            Source: criblcontrolplanesdkgo.Pointer("<value>"),
            Host: criblcontrolplanesdkgo.Pointer("animated-tentacle.biz"),
            Service: criblcontrolplanesdkgo.Pointer("<value>"),
            Tags: []string{
                "<value 1>",
                "<value 2>",
            },
            BatchByTags: criblcontrolplanesdkgo.Pointer(false),
            AllowAPIKeyFromEvents: criblcontrolplanesdkgo.Pointer(true),
            Severity: components.OutputDatadogSeverityEmergency.ToPointer(),
            Site: components.DatadogSiteUs5.ToPointer(),
            SendCountersAsCount: criblcontrolplanesdkgo.Pointer(true),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](7305.28),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](8510.77),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](6981.39),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](9822.43),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](9155.3),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Manual.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](8670.83),
            Description: criblcontrolplanesdkgo.Pointer("subtract duh general almost"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://glaring-membership.name/"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](511.66),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2626.58),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8406.46),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputDatadogPqControls{},
            APIKey: criblcontrolplanesdkgo.Pointer("your-api-key"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesDataset

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesDataset" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputDataset(
        components.OutputDataset{
            ID: criblcontrolplanesdkgo.Pointer("dataset-output"),
            Type: components.OutputDatasetTypeDataset,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            MessageField: criblcontrolplanesdkgo.Pointer("<value>"),
            ExcludeFields: []string{
                "<value 1>",
            },
            ServerHostField: criblcontrolplanesdkgo.Pointer("<value>"),
            TimestampField: criblcontrolplanesdkgo.Pointer("<value>"),
            DefaultSeverity: components.OutputDatasetSeverityFinest.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            Site: components.DataSetSiteUs.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](9193.77),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](9808.87),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](253.36),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](2723.58),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3591.85),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Manual.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](7404.58),
            Description: criblcontrolplanesdkgo.Pointer("unlike by belabor waist barracks superficial vol"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://wavy-mainstream.com"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6272.47),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8207.74),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](9156.37),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputDatasetPqControls{},
            APIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateCustomURL: criblcontrolplanesdkgo.Pointer("https://short-validity.org/"),
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
### Example Usage: OutputCreateExamplesDiskSpool

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesDiskSpool" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputDiskSpool(
        components.OutputDiskSpool{
            ID: criblcontrolplanesdkgo.Pointer("disk-spool-output"),
            Type: components.OutputDiskSpoolTypeDiskSpool,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptionsPersistenceNone.ToPointer(),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("sunny what likewise vivaciously rarely"),
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
### Example Usage: OutputCreateExamplesDlS3

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesDlS3" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputDlS3(
        components.OutputDlS3{
            ID: criblcontrolplanesdkgo.Pointer("dl-s3-output"),
            Type: components.OutputDlS3TypeDlS3,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Bucket: "my-bucket",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](6863.47),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            ObjectACL: components.ObjectACLOptionsPrivate.ToPointer(),
            StorageClass: components.StorageClassOptionsOnezoneIa.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAes256.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            Format: components.DataFormatOptionsRaw.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](9099.5),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](3454.84),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](8359.92),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(true),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](9741.71),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7684.82),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1991.93),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](5686.77),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](1612.93),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](1991.26),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](8183.89),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](9401.43),
            PartitioningFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Description: criblcontrolplanesdkgo.Pointer("before across seemingly courageously yuck providence underneath"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsNormal.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet10.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV1.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](6793.06),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(true),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(false),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](5188.13),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](8658.32),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](2633.72),
            TemplateBucket: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateFormat: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesDynatraceHttp

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesDynatraceHttp" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputDynatraceHTTP(
        components.OutputDynatraceHTTP{
            ID: criblcontrolplanesdkgo.Pointer("dynatrace-http-output"),
            Type: components.OutputDynatraceHTTPTypeDynatraceHTTP,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Method: components.MethodOptionsPut.ToPointer(),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](3097.84),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](6714.87),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](1141.45),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](7758.15),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6285.52),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            AuthType: components.OutputDynatraceHTTPAuthenticationTypeToken.ToPointer(),
            Format: components.OutputDynatraceHTTPFormatJSONArray,
            Endpoint: components.EndpointCloud,
            TelemetryType: components.TelemetryTypeLogs,
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](3794.77),
            Description: criblcontrolplanesdkgo.Pointer("round affectionate likely whether tooth"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1551.62),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](6826.29),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5720.73),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputDynatraceHTTPPqControls{},
            Token: criblcontrolplanesdkgo.Pointer("your-api-key"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            EnvironmentID: criblcontrolplanesdkgo.Pointer("<id>"),
            ActiveGateDomain: criblcontrolplanesdkgo.Pointer("<value>"),
            URL: criblcontrolplanesdkgo.Pointer("https://gullible-attraction.info/"),
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://well-lit-costume.net"),
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
### Example Usage: OutputCreateExamplesDynatraceOtlp

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesDynatraceOtlp" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputDynatraceOtlp(
        components.OutputDynatraceOtlp{
            ID: criblcontrolplanesdkgo.Pointer("dynatrace-otlp-output"),
            Type: components.OutputDynatraceOtlpTypeDynatraceOtlp,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Protocol: components.OutputDynatraceOtlpProtocolHTTP,
            Endpoint: "https://your-environment.live.dynatrace.com/api/v2/otlp",
            OtlpVersion: components.OtlpVersionOptions1OneDot3Dot1,
            Compress: components.CompressionOptions4None.ToPointer(),
            HTTPCompress: components.CompressionOptions5None.ToPointer(),
            HTTPTracesEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPMetricsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPLogsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](210.47),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](9413.88),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](4820.34),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](86.91),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](5704.41),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](5789.58),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            EndpointType: components.EndpointTypeSaas,
            TokenSecret: "your-token-secret",
            AuthTokenName: criblcontrolplanesdkgo.Pointer("<value>"),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("quick-witted while phew fooey helplessly pfft"),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](242.91),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4385.39),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5538.13),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputDynatraceOtlpPqControls{},
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
### Example Usage: OutputCreateExamplesElastic

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesElastic" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputElastic(
        components.OutputElastic{
            ID: criblcontrolplanesdkgo.Pointer("elastic-output"),
            Type: components.OutputElasticTypeElastic,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            LoadBalanced: criblcontrolplanesdkgo.Pointer(true),
            Index: "logs",
            DocType: criblcontrolplanesdkgo.Pointer("<value>"),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](733.26),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](5746.62),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](9728.73),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](9678.92),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3467.47),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            ExtraParams: []components.ItemsTypeSaslSaslExtensions{
                components.ItemsTypeSaslSaslExtensions{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Auth: &components.AuthType{
                Disabled: false,
                Username: criblcontrolplanesdkgo.Pointer("Odie.Skiles"),
                Password: criblcontrolplanesdkgo.Pointer("RyWsR_yFaKceblE"),
                AuthType: components.AuthenticationMethodOptionsAuthTextSecret.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                ManualAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            ElasticVersion: components.ElasticVersionSeven.ToPointer(),
            ElasticPipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            IncludeDocID: criblcontrolplanesdkgo.Pointer(true),
            WriteAction: components.WriteActionIndex.ToPointer(),
            RetryPartialErrors: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("whereas oof modulo impossible mmm till"),
            URL: criblcontrolplanesdkgo.Pointer("https://cuddly-printer.info"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Urls: []components.OutputElasticURL{
                components.OutputElasticURL{
                    URL: "https://easy-dandelion.info/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](9314.75),
                    TemplateURL: criblcontrolplanesdkgo.Pointer("https://proud-footrest.info"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](5398),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](14.64),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](3077.68),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3308.34),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1449.39),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputElasticPqControls{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://whirlwind-tennis.org"),
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
### Example Usage: OutputCreateExamplesElasticCloud

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesElasticCloud" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputElasticCloud(
        components.OutputElasticCloud{
            ID: criblcontrolplanesdkgo.Pointer("elastic-cloud-output"),
            Type: components.OutputElasticCloudTypeElasticCloud,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            URL: "my-cloud-id",
            Index: "logs",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](4319.98),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](9163.36),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](2697.78),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](2028.06),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8449.21),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            ExtraParams: []components.ItemsTypeSaslSaslExtensions{
                components.ItemsTypeSaslSaslExtensions{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Auth: &components.AuthType{
                Disabled: false,
                Username: criblcontrolplanesdkgo.Pointer("Odie.Skiles"),
                Password: criblcontrolplanesdkgo.Pointer("RyWsR_yFaKceblE"),
                AuthType: components.AuthenticationMethodOptionsAuthTextSecret.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                ManualAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            ElasticPipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            IncludeDocID: criblcontrolplanesdkgo.Pointer(false),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("or well-to-do diligently joyfully negotiation greatly bewail gadzooks considering"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9392.65),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](556.12),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](2451.19),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputElasticCloudPqControls{},
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
### Example Usage: OutputCreateExamplesExabeam

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesExabeam" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputExabeam(
        components.OutputExabeam{
            ID: criblcontrolplanesdkgo.Pointer("exabeam-output"),
            Type: components.OutputExabeamTypeExabeam,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Bucket: "my-bucket",
            Region: "us-east1",
            StagePath: "/tmp/staging",
            Endpoint: "https://storage.googleapis.com",
            SignatureVersion: components.SignatureVersionOptions4V2.ToPointer(),
            ObjectACL: components.ObjectACLOptions1BucketOwnerFullControl.ToPointer(),
            StorageClass: components.StorageClassOptions1Nearline.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](8587.19),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](5462.55),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](2343.76),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](9741.71),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7684.82),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1991.93),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](5686.77),
            },
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](1337.73),
            EncodedConfiguration: criblcontrolplanesdkgo.Pointer("<value>"),
            CollectorInstanceID: "11112222-3333-4444-5555-666677778888",
            SiteName: criblcontrolplanesdkgo.Pointer("<value>"),
            SiteID: criblcontrolplanesdkgo.Pointer("<id>"),
            TimezoneOffset: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("seldom gadzooks snarling forecast"),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](9206.29),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](3791.32),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](1606.45),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesFilesystem

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesFilesystem" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputFilesystem(
        components.OutputFilesystem{
            ID: criblcontrolplanesdkgo.Pointer("filesystem-output"),
            Type: components.OutputFilesystemTypeFilesystem,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            DestPath: "/var/log/output",
            StagePath: criblcontrolplanesdkgo.Pointer("<value>"),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](9141.81),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](6269.67),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](4900.17),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](3178.13),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](5888.68),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(true),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](9741.71),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7684.82),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1991.93),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](5686.77),
            },
            Description: criblcontrolplanesdkgo.Pointer("dicker unto scrap young exotic unless drat joyfully"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsNormal.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet24.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV1.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](3704.88),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(false),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](5667.47),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](9357.47),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](6117.35),
            TemplateFormat: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesGoogleChronicle

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesGoogleChronicle" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputGoogleChronicle(
        components.OutputGoogleChronicle{
            ID: criblcontrolplanesdkgo.Pointer("google-chronicle-output"),
            Type: components.OutputGoogleChronicleTypeGoogleChronicle,
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
            APIVersion: components.OutputGoogleChronicleAPIVersionV1.ToPointer(),
            AuthenticationMethod: components.OutputGoogleChronicleAuthenticationMethodManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            LogFormatType: components.SendEventsAsUnstructured,
            Region: criblcontrolplanesdkgo.Pointer("us"),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](8784.21),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](8397.4),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](6095.89),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](6957.39),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8927.22),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](6657.33),
            Description: criblcontrolplanesdkgo.Pointer("despite ouch yum duh"),
            ExtraLogTypes: []components.ExtraLogType{
                components.ExtraLogType{
                    LogType: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("incidentally amid if vicinity ribbon including whether toward"),
                },
            },
            LogType: criblcontrolplanesdkgo.Pointer("<value>"),
            LogTextField: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomerID: criblcontrolplanesdkgo.Pointer("customer-id"),
            Namespace: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomLabels: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            UdmType: components.UDMTypeLogs.ToPointer(),
            APIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            APIKeySecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            ServiceAccountCredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1368.71),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1168.56),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5538.69),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputGoogleChroniclePqControls{},
            TemplateAPIVersion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateCustomerID: criblcontrolplanesdkgo.Pointer("<id>"),
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
### Example Usage: OutputCreateExamplesGoogleCloudLogging

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesGoogleCloudLogging" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputGoogleCloudLogging(
        components.OutputGoogleCloudLogging{
            ID: criblcontrolplanesdkgo.Pointer("google-cloud-logging-output"),
            Type: components.OutputGoogleCloudLoggingTypeGoogleCloudLogging,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            LogLocationType: components.LogLocationTypeProject,
            LogNameExpression: "my-log",
            SanitizeLogNames: criblcontrolplanesdkgo.Pointer(true),
            PayloadFormat: components.PayloadFormatJSON.ToPointer(),
            LogLabels: []components.ItemsTypeLogLabels{
                components.ItemsTypeLogLabels{
                    Label: "<value>",
                    ValueExpression: "<value>",
                },
            },
            ResourceTypeExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            ResourceTypeLabels: []components.ItemsTypeLogLabels{
                components.ItemsTypeLogLabels{
                    Label: "<value>",
                    ValueExpression: "<value>",
                },
            },
            SeverityExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            InsertIDExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            GoogleAuthMethod: components.GoogleAuthenticationMethodOptionsSecret.ToPointer(),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](5491.4),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](8655.58),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](4379.89),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](6997),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](9702.69),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](1224.89),
            ThrottleRateReqPerSec: criblcontrolplanesdkgo.Pointer[int64](306923),
            RequestMethodExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            RequestURLExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            RequestSizeExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            StatusExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            ResponseSizeExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            UserAgentExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            RemoteIPExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            ServerIPExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            RefererExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            LatencyExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            CacheLookupExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            CacheHitExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            CacheValidatedExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            CacheFillBytesExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            ProtocolExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            IDExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            ProducerExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            FirstExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            LastExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            FileExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            LineExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            FunctionExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            UIDExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            IndexExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            TotalSplitsExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            TraceExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            SpanIDExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            TraceSampledExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](7152.99),
            Description: criblcontrolplanesdkgo.Pointer("correctly whirlwind consequently that um nucleotidase nor round independence"),
            LogLocationExpression: "my-project",
            PayloadExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2278.08),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1187.19),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1895.52),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputGoogleCloudLoggingPqControls{},
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
### Example Usage: OutputCreateExamplesGoogleCloudStorage

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesGoogleCloudStorage" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputGoogleCloudStorage(
        components.OutputGoogleCloudStorage{
            ID: criblcontrolplanesdkgo.Pointer("google-cloud-storage-output"),
            Type: components.OutputGoogleCloudStorageTypeGoogleCloudStorage,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Bucket: "my-bucket",
            Region: "us-east1",
            Endpoint: "https://storage.googleapis.com",
            SignatureVersion: components.SignatureVersionOptions4V4.ToPointer(),
            AwsAuthenticationMethod: components.OutputGoogleCloudStorageAuthenticationMethodAuto.ToPointer(),
            StagePath: "/tmp/staging",
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            ObjectACL: components.ObjectACLOptions1AuthenticatedRead.ToPointer(),
            StorageClass: components.StorageClassOptions1Coldline.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsParquet.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](3344.08),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](3310.81),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](1124.18),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](5414.81),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](216.54),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(true),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](9741.71),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7684.82),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1991.93),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](5686.77),
            },
            Description: criblcontrolplanesdkgo.Pointer("slow actually readily"),
            Compress: components.CompressionOptions2None.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestCompression.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet10.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV1.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](4167.75),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(true),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](9443.56),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](318.2),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](8876.52),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateBucket: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateFormat: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesGooglePubsub

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesGooglePubsub" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputGooglePubsub(
        components.OutputGooglePubsub{
            ID: criblcontrolplanesdkgo.Pointer("google-pubsub-output"),
            Type: components.OutputGooglePubsubTypeGooglePubsub,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            TopicName: "my-topic",
            CreateTopic: criblcontrolplanesdkgo.Pointer(true),
            OrderedDelivery: criblcontrolplanesdkgo.Pointer(false),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            GoogleAuthMethod: components.GoogleAuthenticationMethodOptionsSecret.ToPointer(),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            BatchSize: criblcontrolplanesdkgo.Pointer[float64](9482.92),
            BatchTimeout: criblcontrolplanesdkgo.Pointer[float64](7263.98),
            MaxQueueSize: criblcontrolplanesdkgo.Pointer[float64](5971.61),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](3431.33),
            FlushPeriod: criblcontrolplanesdkgo.Pointer[float64](288.76),
            MaxInProgress: criblcontrolplanesdkgo.Pointer[float64](1946.73),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("maroon deduct truly"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](7475.46),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8048.15),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6977.08),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputGooglePubsubPqControls{},
            TemplateTopicName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesGrafanaCloud

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesGrafanaCloud" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputGrafanaCloud(
        components.CreateOutputGrafanaCloudOutputGrafanaCloudGrafanaCloud1(
            components.OutputGrafanaCloudGrafanaCloud1{
                ID: criblcontrolplanesdkgo.Pointer("grafana-cloud-output"),
                Type: components.OutputGrafanaCloudType1GrafanaCloud,
                Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                SystemFields: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
                Environment: criblcontrolplanesdkgo.Pointer("<value>"),
                Streamtags: []string{
                    "<value 1>",
                },
                LokiURL: "https://logs-prod-us-central1.grafana.net",
                PrometheusURL: criblcontrolplanesdkgo.Pointer("https://difficult-mentor.info"),
                Message: criblcontrolplanesdkgo.Pointer("<value>"),
                MessageFormat: components.MessageFormatOptionsProtobuf.ToPointer(),
                Labels: []components.ItemsTypeLabels{
                    components.ItemsTypeLabels{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                MetricRenameExpr: criblcontrolplanesdkgo.Pointer("<value>"),
                PrometheusAuth: &components.PrometheusAuthType{
                    AuthType: components.AuthenticationTypeOptionsPrometheusAuth1CredentialsSecret.ToPointer(),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Username: criblcontrolplanesdkgo.Pointer("Jedidiah.Hammes"),
                    Password: criblcontrolplanesdkgo.Pointer("vldG5B7AGD7DT_h"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                LokiAuth: &components.PrometheusAuthType{
                    AuthType: components.AuthenticationTypeOptionsPrometheusAuth1Basic.ToPointer(),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Username: criblcontrolplanesdkgo.Pointer("Berry.Lowe"),
                    Password: criblcontrolplanesdkgo.Pointer("pUBpjkjynnpVdwt"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                Concurrency: criblcontrolplanesdkgo.Pointer[float64](2881.29),
                MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4737.16),
                MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](5944.74),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](1009.95),
                FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](9956.31),
                ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                    components.ItemsTypeExtraHTTPHeaders{
                        Name: criblcontrolplanesdkgo.Pointer("<value>"),
                        Value: "<value>",
                    },
                },
                UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
                FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
                SafeHeaders: []string{
                    "<value 1>",
                    "<value 2>",
                },
                ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                    components.ItemsTypeResponseRetrySettings{
                        HTTPStatus: 9145.72,
                        InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                        BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                        MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                    },
                },
                TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                    TimeoutRetry: false,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
                },
                ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
                OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
                Description: criblcontrolplanesdkgo.Pointer("stir-fry inquisitively developing coaxingly icy gee"),
                Compress: criblcontrolplanesdkgo.Pointer(false),
                PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
                PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6236.79),
                PqMode: components.ModeOptionsBackpressure.ToPointer(),
                PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2055.87),
                PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5302.97),
                PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PqCompress: components.CompressionOptionsPqNone.ToPointer(),
                PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
                PqControls: &components.OutputGrafanaCloudPqControls1{},
                TemplateLokiURL: criblcontrolplanesdkgo.Pointer("https://wasteful-summary.org"),
                TemplatePrometheusURL: criblcontrolplanesdkgo.Pointer("https://zesty-hyphenation.net"),
            },
        ),
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedOutput != nil {
        // handle response
    }
}
```
### Example Usage: OutputCreateExamplesGraphite

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesGraphite" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputGraphite(
        components.OutputGraphite{
            ID: criblcontrolplanesdkgo.Pointer("graphite-output"),
            Type: components.OutputGraphiteTypeGraphite,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Protocol: components.DestinationProtocolOptionsTCP,
            Host: "localhost",
            Port: 2003,
            Mtu: criblcontrolplanesdkgo.Pointer[float64](1108.13),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3328.06),
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](2083.11),
            Description: criblcontrolplanesdkgo.Pointer("affectionate about clear save dwell"),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](2176.12),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](1709.73),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9742.96),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](6208.93),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](3617.53),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputGraphitePqControls{},
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
### Example Usage: OutputCreateExamplesHoneycomb

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesHoneycomb" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputHoneycomb(
        components.OutputHoneycomb{
            ID: criblcontrolplanesdkgo.Pointer("honeycomb-output"),
            Type: components.OutputHoneycombTypeHoneycomb,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Dataset: "my-dataset",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1911.92),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](6791.12),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](2263.49),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](4717.56),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6241.91),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("singing pacemaker abaft finger help meh um and"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5731.2),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4952.91),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](2623.72),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputHoneycombPqControls{},
            Team: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesHumioHec

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesHumioHec" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputHumioHec(
        components.OutputHumioHec{
            ID: criblcontrolplanesdkgo.Pointer("humio-hec-output"),
            Type: components.OutputHumioHecTypeHumioHec,
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
            URL: "https://cloud.us.humio.com/api/v1/ingest/hec",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](8748.35),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](5912.04),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](1669.79),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](3955.5),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8784.08),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            Format: components.RequestFormatOptionsJSON,
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("on humble that provided"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1560.43),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7920.24),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](77.53),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputHumioHecPqControls{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://free-in-joke.net"),
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
### Example Usage: OutputCreateExamplesInfluxdb

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesInfluxdb" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputInfluxdb(
        components.OutputInfluxdb{
            ID: criblcontrolplanesdkgo.Pointer("influxdb-output"),
            Type: components.OutputInfluxdbTypeInfluxdb,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            URL: "http://localhost:8086",
            UseV2API: criblcontrolplanesdkgo.Pointer(true),
            TimestampPrecision: components.TimestampPrecisionM.ToPointer(),
            DynamicValueFieldName: criblcontrolplanesdkgo.Pointer(false),
            ValueFieldName: criblcontrolplanesdkgo.Pointer("<value>"),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](3829.03),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](6069.71),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](5528.1),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](33.06),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](9775.28),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.OutputInfluxdbAuthenticationTypeTextSecret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("unsteady famously however"),
            Database: criblcontrolplanesdkgo.Pointer("mydb"),
            Bucket: criblcontrolplanesdkgo.Pointer("<value>"),
            Org: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6508.18),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5310.59),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](9488.72),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputInfluxdbPqControls{},
            Username: criblcontrolplanesdkgo.Pointer("Jackson_Hettinger40"),
            Password: criblcontrolplanesdkgo.Pointer("cNBgRT46DROEKDU"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://utter-nectarine.com/"),
            TemplateDatabase: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateBucket: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesKafka

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesKafka" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputKafka(
        components.OutputKafka{
            ID: criblcontrolplanesdkgo.Pointer("kafka-output"),
            Type: components.OutputKafkaTypeKafka,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Brokers: []string{
                "localhost:9092",
            },
            Topic: "logs",
            Ack: components.AcknowledgmentsOptions1None.ToPointer(),
            Format: components.RecordDataFormatOptions1Raw.ToPointer(),
            Compression: components.CompressionOptions3Lz4.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](2385.96),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](4905.03),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7310),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType1{
                Disabled: true,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://lighthearted-dredger.org"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](7210.32),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7627.12),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](7494.42),
                Auth: &components.AuthTypeKafkaSchemaRegistry{
                    Disabled: false,
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                    Disabled: criblcontrolplanesdkgo.Pointer(true),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                    CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
                },
                DefaultKeySchemaID: criblcontrolplanesdkgo.Pointer[float64](5091.61),
                DefaultValueSchemaID: criblcontrolplanesdkgo.Pointer[float64](2259.59),
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](6414.38),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](6220.17),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](9336.12),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](1613.32),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](6542.35),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](4920.44),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](2234.31),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](1737.55),
            Sasl: &components.AuthenticationType{
                Disabled: false,
                Username: criblcontrolplanesdkgo.Pointer("Vaughn_Botsford"),
                Password: criblcontrolplanesdkgo.Pointer("J3to4z1Pba3wD7B"),
                AuthType: components.AuthenticationMethodOptionsSaslSecret.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslScramSha256.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(false),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://runny-saw.name"),
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
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("meanwhile aching in ecliptic than through"),
            ProtobufLibraryID: criblcontrolplanesdkgo.Pointer("<id>"),
            ProtobufEncodingID: criblcontrolplanesdkgo.Pointer("<id>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](558.33),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5057.63),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5220.97),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputKafkaPqControls{},
            TemplateTopic: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesKinesis

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesKinesis" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputKinesis(
        components.OutputKinesis{
            ID: criblcontrolplanesdkgo.Pointer("kinesis-output"),
            Type: components.OutputKinesisTypeKinesis,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            StreamName: "my-stream",
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfManual.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions2V2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](1103.42),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](3976.04),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](2294.33),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](209.99),
            Compression: components.OutputKinesisCompressionNone.ToPointer(),
            UseListShards: criblcontrolplanesdkgo.Pointer(true),
            AsNdjson: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("oddly through gadzooks whoa"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxEventsPerFlush: criblcontrolplanesdkgo.Pointer[float64](4963.5),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5757.02),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5915.4),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](9169.27),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputKinesisPqControls{},
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
    if res.CountedOutput != nil {
        // handle response
    }
}
```
### Example Usage: OutputCreateExamplesLoki

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesLoki" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputLoki(
        components.OutputLoki{
            ID: criblcontrolplanesdkgo.Pointer("loki-output"),
            Type: components.OutputLokiTypeLoki,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            URL: "http://localhost:3100/loki/api/v1/push",
            Message: criblcontrolplanesdkgo.Pointer("<value>"),
            MessageFormat: components.MessageFormatOptionsJSON.ToPointer(),
            Labels: []components.ItemsTypeLabels{
                components.ItemsTypeLabels{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthType: components.AuthenticationTypeOptionsPrometheusAuth1Token.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5209.75),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4899.83),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](866.12),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](661.18),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8113.72),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            EnableDynamicHeaders: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](5730.87),
            Description: criblcontrolplanesdkgo.Pointer("why wherever that anxiously"),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Username: criblcontrolplanesdkgo.Pointer("Jeremy12"),
            Password: criblcontrolplanesdkgo.Pointer("Kaun7gP2PP7nR8C"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](7524.6),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7287.04),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5827.98),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputLokiPqControls{},
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
### Example Usage: OutputCreateExamplesMicrosoftFabric

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesMicrosoftFabric" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputMicrosoftFabric(
        components.OutputMicrosoftFabric{
            ID: criblcontrolplanesdkgo.Pointer("microsoft-fabric-output"),
            Type: components.OutputMicrosoftFabricTypeMicrosoftFabric,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Topic: "logs",
            Ack: components.AcknowledgmentsOptionsNone.ToPointer(),
            Format: components.RecordDataFormatOptionsRaw.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](5491.84),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](7850.09),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](9993.97),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](6233.74),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](5606.2),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](4955.03),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](9077.83),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9941.69),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](6244.55),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](4271.75),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](8338.17),
            Sasl: &components.OutputMicrosoftFabricAuthentication{
                Disabled: false,
                Mechanism: components.SaslMechanismOptionsSasl1Oauthbearer.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("Abner_Fritsch79"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                ClientSecretAuthType: components.OutputMicrosoftFabricAuthenticationMethodSecret.ToPointer(),
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
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            BootstrapServer: "myeventstream.servicebus.windows.net:9093",
            Description: criblcontrolplanesdkgo.Pointer("mutate surprisingly worriedly after fatally bleach worldly"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6423.23),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4887.41),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](3522.04),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputMicrosoftFabricPqControls{},
            TemplateTopic: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateBootstrapServer: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesMinio

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesMinio" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputMinio(
        components.OutputMinio{
            ID: criblcontrolplanesdkgo.Pointer("minio-output"),
            Type: components.OutputMinioTypeMinio,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Endpoint: "http://localhost:9000",
            Bucket: "my-bucket",
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfSecret.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions5V2.ToPointer(),
            ObjectACL: components.ObjectACLOptionsPublicRead.ToPointer(),
            StorageClass: components.StorageClassOptions2Standard.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionOptionsAes256.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(false),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsRaw.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](6437.35),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](3981.69),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](8983.21),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](9741.71),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7684.82),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1991.93),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](5686.77),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](8039.1),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](2262.73),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](457.87),
            Description: criblcontrolplanesdkgo.Pointer("slipper meanwhile excepting giving dispense provided prime extroverted outside lest"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2None.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet24.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](1393.11),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(true),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(false),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(false),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](6647.23),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1311.23),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](2082.71),
            TemplateBucket: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateFormat: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesMsk

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesMsk" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputMsk(
        components.OutputMsk{
            ID: criblcontrolplanesdkgo.Pointer("msk-output"),
            Type: components.OutputMskTypeMsk,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Brokers: []string{
                "b-1.example.xxxxx.c2.kafka.us-east-1.amazonaws.com:9092",
            },
            Topic: "logs",
            Ack: components.AcknowledgmentsOptions1Leader.ToPointer(),
            Format: components.RecordDataFormatOptions1JSON.ToPointer(),
            Compression: components.CompressionOptions3None.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](1115.17),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](3467.83),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](2384.7),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType1{
                Disabled: true,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://lighthearted-dredger.org"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](7210.32),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7627.12),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](7494.42),
                Auth: &components.AuthTypeKafkaSchemaRegistry{
                    Disabled: false,
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                    Disabled: criblcontrolplanesdkgo.Pointer(true),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                    CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
                },
                DefaultKeySchemaID: criblcontrolplanesdkgo.Pointer[float64](5091.61),
                DefaultValueSchemaID: criblcontrolplanesdkgo.Pointer[float64](2259.59),
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](5007.18),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](581.3),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](636.26),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](6902.71),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](6572.27),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](4364.27),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](4969.27),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](5111.56),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto,
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3344.39),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("till motivate near lady shadowbox pale aside naturally until"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ProtobufLibraryID: criblcontrolplanesdkgo.Pointer("<id>"),
            ProtobufEncodingID: criblcontrolplanesdkgo.Pointer("<id>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](7271.21),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1225.14),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6320.84),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputMskPqControls{},
            TemplateTopic: criblcontrolplanesdkgo.Pointer("<value>"),
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
    if res.CountedOutput != nil {
        // handle response
    }
}
```
### Example Usage: OutputCreateExamplesNetflow

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesNetflow" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputNetflow(
        components.OutputNetflow{
            ID: criblcontrolplanesdkgo.Pointer("netflow-output"),
            Type: components.OutputNetflowTypeNetflow,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Hosts: []components.OutputNetflowHost{
                components.OutputNetflowHost{
                    Host: "localhost",
                    Port: 2055,
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](9038.19),
            EnableIPSpoofing: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("stealthily lest along"),
            MaxRecordSize: criblcontrolplanesdkgo.Pointer[float64](7756.42),
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
### Example Usage: OutputCreateExamplesNewrelic

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesNewrelic" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputNewrelic(
        components.OutputNewrelic{
            ID: criblcontrolplanesdkgo.Pointer("newrelic-output"),
            Type: components.OutputNewrelicTypeNewrelic,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Region: components.RegionOptionsEu.ToPointer(),
            LogType: criblcontrolplanesdkgo.Pointer("<value>"),
            MessageField: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.Metadatum{
                components.Metadatum{
                    Name: components.FieldNameAuditID,
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](9558.92),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](5115.13),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](6174.66),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](4711.57),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](46.1),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Manual.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](2240.62),
            Description: criblcontrolplanesdkgo.Pointer("tectonics than briefly pish wildly"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://quiet-wombat.info"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](602.38),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7480.49),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1506.62),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputNewrelicPqControls{},
            APIKey: criblcontrolplanesdkgo.Pointer("your-api-key"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateLogType: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateMessageField: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesNewrelicEvents

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesNewrelicEvents" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputNewrelicEvents(
        components.OutputNewrelicEvents{
            ID: criblcontrolplanesdkgo.Pointer("newrelic-events-output"),
            Type: components.OutputNewrelicEventsTypeNewrelicEvents,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Region: components.RegionOptionsCustom.ToPointer(),
            AccountID: "123456",
            EventType: "CriblEvent",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1490.13),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](23.35),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](2495.92),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](5294.51),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](4825.42),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("internalise greedy meanwhile wetly"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://vague-hyphenation.net/"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](7505.35),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7681.23),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](3570.46),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputNewrelicEventsPqControls{},
            APIKey: criblcontrolplanesdkgo.Pointer("your-api-key"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateEventType: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateCustomURL: criblcontrolplanesdkgo.Pointer("https://twin-ferret.com"),
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
### Example Usage: OutputCreateExamplesOpenTelemetry

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesOpenTelemetry" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputOpenTelemetry(
        components.OutputOpenTelemetry{
            ID: criblcontrolplanesdkgo.Pointer("opentelemetry-output"),
            Type: components.OutputOpenTelemetryTypeOpenTelemetry,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Protocol: components.ProtocolOptionsHTTP.ToPointer(),
            Endpoint: "http://localhost:4317",
            OtlpVersion: components.OutputOpenTelemetryOTLPVersionZeroDot10Dot0.ToPointer(),
            Compress: components.CompressionOptions4Gzip.ToPointer(),
            HTTPCompress: components.CompressionOptions5None.ToPointer(),
            AuthType: components.AuthenticationTypeOptionsNone.ToPointer(),
            HTTPTracesEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPMetricsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPLogsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](7438.69),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](5961.44),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](8866.44),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3115.59),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](1681.74),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](783.35),
            KeepAlive: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("wherever after so grimy"),
            Username: criblcontrolplanesdkgo.Pointer("Newell_Block"),
            Password: criblcontrolplanesdkgo.Pointer("Pog3DCwYAVHYvXx"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            TLS: &components.TLSSettingsClientSideType2{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1273.92),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3586.04),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](3775.16),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputOpenTelemetryPqControls{},
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
### Example Usage: OutputCreateExamplesPrometheus

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesPrometheus" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputPrometheus(
        components.OutputPrometheus{
            ID: criblcontrolplanesdkgo.Pointer("prometheus-output"),
            Type: components.OutputPrometheusTypePrometheus,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            URL: "http://localhost:9091/api/v1/write",
            MetricRenameExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            SendMetadata: criblcontrolplanesdkgo.Pointer(false),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](3188.74),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](2790.51),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](542.04),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](3180.33),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7593.86),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationTypeOptionsPrometheusAuthTextSecret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("unless sleepily instantly dwell per pitiful under broadly axe reporter"),
            MetricsFlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8599.01),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5568.14),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8622.36),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](9787),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputPrometheusPqControls{},
            Username: criblcontrolplanesdkgo.Pointer("Alison31"),
            Password: criblcontrolplanesdkgo.Pointer("zoUKDif9prN2TD0"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://sociable-case.name/"),
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
### Example Usage: OutputCreateExamplesRing

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesRing" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputRing(
        components.OutputRing{
            ID: criblcontrolplanesdkgo.Pointer("ring-output"),
            Type: components.OutputRingTypeRing,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Format: components.OutputRingDataFormatRaw.ToPointer(),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("vice bind for bah whenever below following"),
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
### Example Usage: OutputCreateExamplesRouter

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesRouter" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputRouter(
        components.OutputRouter{
            ID: criblcontrolplanesdkgo.Pointer("router-output"),
            Type: components.OutputRouterTypeRouter,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Rules: []components.OutputRouterRule{
                components.OutputRouterRule{
                    Filter: "true",
                    Output: "my-output",
                    Description: criblcontrolplanesdkgo.Pointer("yearly hot accelerator darling energetically sedately without detective"),
                    Final: criblcontrolplanesdkgo.Pointer(true),
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("dental oh how"),
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
### Example Usage: OutputCreateExamplesS3

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesS3" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputS3(
        components.OutputS3{
            ID: criblcontrolplanesdkgo.Pointer("s3-output"),
            Type: components.OutputS3TypeS3,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Bucket: "my-bucket",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](4718.83),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            ObjectACL: components.ObjectACLOptionsPublicReadWrite.ToPointer(),
            StorageClass: components.StorageClassOptionsStandard.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAwsKms.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](6192.09),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](3926.44),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](8748.54),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(true),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](9741.71),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7684.82),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1991.93),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](5686.77),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](4324.66),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](2081.88),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](9218.14),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(false),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](6709.29),
            Description: criblcontrolplanesdkgo.Pointer("furthermore iterate within yuck inconsequential personalise puritan oh synthesise"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestCompression.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](6277.7),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(true),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](1725.86),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](4697.22),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](5211.8),
            TemplateBucket: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateFormat: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesSecurityLake

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesSecurityLake" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputSecurityLake(
        components.OutputSecurityLake{
            ID: criblcontrolplanesdkgo.Pointer("security-lake-output"),
            Type: components.OutputSecurityLakeTypeSecurityLake,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Bucket: "my-bucket",
            Region: "us-east-1",
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfManual.ToPointer(),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.OutputSecurityLakeSignatureVersionV2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: "arn:aws:iam::123456789012:role/my-role",
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](1265.03),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            ObjectACL: components.ObjectACLOptionsAwsExecRead.ToPointer(),
            StorageClass: components.StorageClassOptionsIntelligentTiering.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAwsKms.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](1025.66),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](8406.64),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](6587.5),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](9741.71),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7684.82),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](1991.93),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](5686.77),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](3870.7),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](8758.85),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](2545.81),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](5493.06),
            AccountID: "123456789012",
            CustomSource: "my-custom-source",
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetVersion: components.ParquetVersionOptionsParquet24.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV1.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](4946.58),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(true),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(false),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("nor astride behind drat drum proliferate whether"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](5174.85),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](6602.4),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](641.9),
            TemplateBucket: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateAwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesSentinel

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesSentinel" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputSentinel(
        components.OutputSentinel{
            ID: criblcontrolplanesdkgo.Pointer("sentinel-output"),
            Type: components.OutputSentinelTypeSentinel,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            KeepAlive: criblcontrolplanesdkgo.Pointer(false),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](7452.54),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](259.52),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](1734.02),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](8639.21),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](9332.14),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            AuthType: components.AuthTypeEnumOauth.ToPointer(),
            LoginURL: "https://login.microsoftonline.com",
            Secret: "client-secret",
            ClientID: "client-id",
            Scope: criblcontrolplanesdkgo.Pointer("<value>"),
            EndpointURLConfiguration: components.EndpointConfigurationURL,
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](4882.5),
            Description: criblcontrolplanesdkgo.Pointer("best psst modulo pivot paralyse exotic prime outside hungrily"),
            Format: components.OutputSentinelFormatCustom.ToPointer(),
            CustomSourceExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomDropWhenNull: criblcontrolplanesdkgo.Pointer(true),
            CustomEventDelimiter: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomContentType: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomPayloadExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            AdvancedContentType: criblcontrolplanesdkgo.Pointer("<value>"),
            FormatEventCode: criblcontrolplanesdkgo.Pointer("<value>"),
            FormatPayloadCode: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9126.79),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](136.07),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](7727.78),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputSentinelPqControls{},
            URL: criblcontrolplanesdkgo.Pointer("https://your-workspace.ingest.monitor.azure.com"),
            DcrID: criblcontrolplanesdkgo.Pointer("<id>"),
            DceEndpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            StreamName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateLoginURL: criblcontrolplanesdkgo.Pointer("https://ill-lifestyle.name"),
            TemplateSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateClientID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateScope: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://irresponsible-suspension.net"),
            TemplateDcrID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateDceEndpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateStreamName: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesSentinelOneAiSiem

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesSentinelOneAiSiem" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputSentinelOneAiSiem(
        components.OutputSentinelOneAiSiem{
            ID: criblcontrolplanesdkgo.Pointer("sentinel-one-ai-siem-output"),
            Type: components.OutputSentinelOneAiSiemTypeSentinelOneAiSiem,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Region: components.RegionUs,
            Endpoint: components.AISIEMEndpointPathRootServicesCollectorEvent,
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](9061.93),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](8105.66),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](5244.2),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](629.95),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](891.65),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("apropos lace ugh illustrious lovingly attraction skeleton whenever wetly"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            BaseURL: criblcontrolplanesdkgo.Pointer("https://tidy-lox.net"),
            HostExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            SourceExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            SourceTypeExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceCategoryExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceNameExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceVendorExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            EventTypeExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            Host: criblcontrolplanesdkgo.Pointer("empty-asset.biz"),
            Source: criblcontrolplanesdkgo.Pointer("<value>"),
            SourceType: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceCategory: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceName: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceVendor: criblcontrolplanesdkgo.Pointer("<value>"),
            EventType: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9088.42),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](6644.04),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](7125.85),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputSentinelOneAiSiemPqControls{},
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
### Example Usage: OutputCreateExamplesServiceNow

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesServiceNow" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputServiceNow(
        components.OutputServiceNow{
            ID: criblcontrolplanesdkgo.Pointer("servicenow-output"),
            Type: components.OutputServiceNowTypeServiceNow,
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
            Endpoint: "ingest.lightstep.com:443",
            TokenSecret: "your-token-secret",
            AuthTokenName: criblcontrolplanesdkgo.Pointer("<value>"),
            OtlpVersion: components.OtlpVersionOptions1OneDot3Dot1,
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](9648.28),
            Protocol: components.ProtocolOptionsHTTP,
            Compress: components.CompressionOptions4Gzip.ToPointer(),
            HTTPCompress: components.CompressionOptions5None.ToPointer(),
            HTTPTracesEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPMetricsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPLogsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](4338.53),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](2341.74),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6471.56),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](1985.64),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](3282.21),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("towards broadcast um er however wasabi worriedly afore"),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            TLS: &components.TLSSettingsClientSideType2{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1041.55),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2544.94),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](4506.9),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputServiceNowPqControls{},
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
### Example Usage: OutputCreateExamplesSignalfx

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesSignalfx" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputSignalfx(
        components.OutputSignalfx{
            ID: criblcontrolplanesdkgo.Pointer("signalfx-output"),
            Type: components.OutputSignalfxTypeSignalfx,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Realm: "us0",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](8034.73),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](9984.9),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](3261.65),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](1420.03),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7567.96),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("knavishly whale an impressionable in scholarship penalise below opposite along"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](7221.57),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5552.89),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](3344.59),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputSignalfxPqControls{},
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
### Example Usage: OutputCreateExamplesSnmp

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesSnmp" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputSnmp(
        components.OutputSnmp{
            ID: criblcontrolplanesdkgo.Pointer("snmp-output"),
            Type: components.OutputSnmpTypeSnmp,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Hosts: []components.OutputSnmpHost{
                components.OutputSnmpHost{
                    Host: "192.168.1.1",
                    Port: 161,
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](9651.38),
            Description: criblcontrolplanesdkgo.Pointer("zowie but once outdo runny meh whereas"),
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
### Example Usage: OutputCreateExamplesSns

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesSns" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputSns(
        components.OutputSns{
            ID: criblcontrolplanesdkgo.Pointer("sns-output"),
            Type: components.OutputSnsTypeSns,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            TopicArn: "arn:aws:sns:us-east-1:123456789012:my-topic",
            MessageGroupID: "my-message-group",
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](2785.69),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfSecret.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.OutputSnsSignatureVersionV2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](8775.69),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("bewail clumsy bank delightfully doodle splash collectivization gosh observe"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5474.65),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](6128.21),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](4668.8),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputSnsPqControls{},
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
    if res.CountedOutput != nil {
        // handle response
    }
}
```
### Example Usage: OutputCreateExamplesSplunk

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesSplunk" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputSplunk(
        components.OutputSplunk{
            ID: criblcontrolplanesdkgo.Pointer("splunk-output"),
            Type: components.OutputSplunkTypeSplunk,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Host: "localhost",
            Port: 9997,
            NestedFields: components.NestedFieldSerializationOptionsJSON.ToPointer(),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](8728.82),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](9415.24),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
            },
            EnableMultiMetrics: criblcontrolplanesdkgo.Pointer(true),
            EnableACK: criblcontrolplanesdkgo.Pointer(true),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(false),
            MaxS2Sversion: components.MaxS2SVersionOptionsV4.ToPointer(),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("woot solemnly bah plus supposing reassuringly summary perky but"),
            MaxFailedHealthChecks: criblcontrolplanesdkgo.Pointer[float64](7847.74),
            Compress: components.CompressionOptionsDisabled.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2414.19),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9268.97),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](7468.74),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputSplunkPqControls{},
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
### Example Usage: OutputCreateExamplesSplunkHec

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesSplunkHec" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputSplunkHec(
        components.OutputSplunkHec{
            ID: criblcontrolplanesdkgo.Pointer("splunk-hec-output"),
            Type: components.OutputSplunkHecTypeSplunkHec,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            LoadBalanced: criblcontrolplanesdkgo.Pointer(false),
            NextQueue: criblcontrolplanesdkgo.Pointer("<value>"),
            TCPRouting: criblcontrolplanesdkgo.Pointer("<value>"),
            TLS: &components.TLSSettingsClientSideType1{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](2075.02),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](9369.83),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](351.37),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](2414.52),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3436.47),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            EnableMultiMetrics: criblcontrolplanesdkgo.Pointer(true),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("oof often truly difficult pish yuck"),
            URL: criblcontrolplanesdkgo.Pointer("https://forceful-compromise.net"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Urls: []components.OutputSplunkHecURL{
                components.OutputSplunkHecURL{
                    URL: "https://puny-case.net/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](5487.59),
                    TemplateURL: criblcontrolplanesdkgo.Pointer("https://stable-humor.net"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](5280.98),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](4197.66),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9986.63),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3014.85),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](9782.24),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputSplunkHecPqControls{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://sudden-tributary.net"),
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
### Example Usage: OutputCreateExamplesSplunkLb

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesSplunkLb" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputSplunkLb(
        components.OutputSplunkLb{
            ID: criblcontrolplanesdkgo.Pointer("splunk-lb-output"),
            Type: components.OutputSplunkLbTypeSplunkLb,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](5646.62),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](9217.44),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](8658.99),
            NestedFields: components.NestedFieldSerializationOptionsJSON.ToPointer(),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](2955.33),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](8574.03),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
            },
            EnableMultiMetrics: criblcontrolplanesdkgo.Pointer(false),
            EnableACK: criblcontrolplanesdkgo.Pointer(true),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(true),
            MaxS2Sversion: components.MaxS2SVersionOptionsV4.ToPointer(),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            IndexerDiscovery: criblcontrolplanesdkgo.Pointer(true),
            SenderUnhealthyTimeAllowance: criblcontrolplanesdkgo.Pointer[float64](7463.13),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("gadzooks founder in"),
            MaxFailedHealthChecks: criblcontrolplanesdkgo.Pointer[float64](5884.8),
            Compress: components.CompressionOptionsDisabled.ToPointer(),
            IndexerDiscoveryConfigs: &components.IndexerDiscoveryConfigs{
                Site: "<value>",
                MasterURI: "https://intrepid-warming.biz/",
                RefreshIntervalSec: 1878.89,
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                AuthTokens: []components.OutputSplunkLbAuthToken{
                    components.OutputSplunkLbAuthToken{
                        AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
                        AuthToken: criblcontrolplanesdkgo.Pointer("<value>"),
                        TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    },
                },
                AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
                AuthToken: criblcontrolplanesdkgo.Pointer("<value>"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(true),
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "localhost",
                    Port: 9997,
                    TLS: components.TLSOptionsHostsItemsOff.ToPointer(),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1733.48),
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](827.91),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2331.28),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](4424.21),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputSplunkLbPqControls{},
            AuthToken: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesSqs

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesSqs" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputSqs(
        components.OutputSqs{
            ID: criblcontrolplanesdkgo.Pointer("sqs-output"),
            Type: components.OutputSqsTypeSqs,
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
            QueueName: "my-queue",
            QueueType: components.OutputSqsQueueTypeStandard,
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            MessageGroupID: criblcontrolplanesdkgo.Pointer("<id>"),
            CreateQueue: criblcontrolplanesdkgo.Pointer(false),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfSecret.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions3V2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](8027.27),
            MaxQueueSize: criblcontrolplanesdkgo.Pointer[float64](1545.05),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](2828.49),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3670.2),
            MaxInProgress: criblcontrolplanesdkgo.Pointer[float64](2921.36),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("fooey swordfish worst connect what hm paltry tributary"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6796.85),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4862.56),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](4675.07),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputSqsPqControls{},
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
    if res.CountedOutput != nil {
        // handle response
    }
}
```
### Example Usage: OutputCreateExamplesStatsd

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesStatsd" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputStatsd(
        components.OutputStatsd{
            ID: criblcontrolplanesdkgo.Pointer("statsd-output"),
            Type: components.OutputStatsdTypeStatsd,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Protocol: components.DestinationProtocolOptionsUDP,
            Host: "localhost",
            Port: 8125,
            Mtu: criblcontrolplanesdkgo.Pointer[float64](393.48),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5.11),
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](2355.45),
            Description: criblcontrolplanesdkgo.Pointer("great familiar make"),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](1162.08),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](6512.05),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8267.14),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](6245.22),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6436.69),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputStatsdPqControls{},
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
### Example Usage: OutputCreateExamplesStatsdExt

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesStatsdExt" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputStatsdExt(
        components.OutputStatsdExt{
            ID: criblcontrolplanesdkgo.Pointer("statsd-ext-output"),
            Type: components.OutputStatsdExtTypeStatsdExt,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Protocol: components.DestinationProtocolOptionsUDP,
            Host: "localhost",
            Port: 8125,
            Mtu: criblcontrolplanesdkgo.Pointer[float64](9805.29),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8070.55),
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](6532.22),
            Description: criblcontrolplanesdkgo.Pointer("phew meanwhile gee emulsify ew and next mechanically whereas"),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](2274.33),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](1209.76),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9641.16),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9213.48),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5873.87),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputStatsdExtPqControls{},
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
### Example Usage: OutputCreateExamplesSumoLogic

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesSumoLogic" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputSumoLogic(
        components.OutputSumoLogic{
            ID: criblcontrolplanesdkgo.Pointer("sumo-logic-output"),
            Type: components.OutputSumoLogicTypeSumoLogic,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            URL: "https://endpoint1.collection.us2.sumologic.com",
            CustomSource: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomCategory: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.OutputSumoLogicDataFormatRaw.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](8191.14),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](2659.7),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](7967.49),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](812.41),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6680.93),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](1901.24),
            Description: criblcontrolplanesdkgo.Pointer("heighten blushing finally bah"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](3450.55),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](835.24),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](491.18),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputSumoLogicPqControls{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://bowed-provider.com"),
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
### Example Usage: OutputCreateExamplesSyslog

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesSyslog" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputSyslog(
        components.OutputSyslog{
            ID: criblcontrolplanesdkgo.Pointer("syslog-output"),
            Type: components.OutputSyslogTypeSyslog,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Protocol: components.OutputSyslogProtocolUDP.ToPointer(),
            Facility: components.FacilityTen.ToPointer(),
            Severity: components.OutputSyslogSeverityInfo.ToPointer(),
            AppName: criblcontrolplanesdkgo.Pointer("<value>"),
            MessageFormat: components.MessageFormatRfc5424.ToPointer(),
            TimestampFormat: components.TimestampFormatIso8601.ToPointer(),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            OctetCountFraming: criblcontrolplanesdkgo.Pointer(true),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("shudder mockingly dispense gloomy anenst gosh until whenever"),
            LoadBalanced: criblcontrolplanesdkgo.Pointer(false),
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](514),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "webbed-import.com",
                    Port: 5669.2,
                    TLS: components.TLSOptionsHostsItemsOff.ToPointer(),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    Weight: criblcontrolplanesdkgo.Pointer[float64](8579.92),
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](9446.86),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](417.62),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](9989.57),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](6911.57),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](2379.31),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            MaxRecordSize: criblcontrolplanesdkgo.Pointer[float64](6755.68),
            UDPDNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](4579.44),
            EnableIPSpoofing: criblcontrolplanesdkgo.Pointer(true),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9178.98),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2070.88),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6498.55),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputSyslogPqControls{},
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
### Example Usage: OutputCreateExamplesTcpjson

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesTcpjson" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputTcpjson(
        components.OutputTcpjson{
            ID: criblcontrolplanesdkgo.Pointer("tcpjson-output"),
            Type: components.OutputTcpjsonTypeTcpjson,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            LoadBalanced: criblcontrolplanesdkgo.Pointer(true),
            Compression: components.CompressionOptions1Gzip.ToPointer(),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(true),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](8843.62),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](3301.22),
            TokenTTLMinutes: criblcontrolplanesdkgo.Pointer[float64](1770.13),
            SendHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("aside unabashedly while sometimes descent slump worriedly weary revoke"),
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](10090),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "webbed-import.com",
                    Port: 5669.2,
                    TLS: components.TLSOptionsHostsItemsOff.ToPointer(),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    Weight: criblcontrolplanesdkgo.Pointer[float64](8579.92),
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](1892.95),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](2219.18),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](5741.69),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2434.51),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2098.8),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8241.17),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputTcpjsonPqControls{},
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
### Example Usage: OutputCreateExamplesWavefront

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesWavefront" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputWavefront(
        components.OutputWavefront{
            ID: criblcontrolplanesdkgo.Pointer("wavefront-output"),
            Type: components.OutputWavefrontTypeWavefront,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
            Domain: "longboard",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](6780.94),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](8710.76),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](4977.62),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](1014.39),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1863.25),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("minority fork diligently consequently though absent whose ignorance dutiful"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](439.55),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](984.45),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6714.61),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputWavefrontPqControls{},
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
### Example Usage: OutputCreateExamplesWebhook

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesWebhook" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputWebhook(
        components.OutputWebhook{
            ID: criblcontrolplanesdkgo.Pointer("webhook-output"),
            Type: components.OutputWebhookTypeWebhook,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Method: components.MethodOptionsPut.ToPointer(),
            Format: components.OutputWebhookFormatAdvanced.ToPointer(),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](8755.55),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](6177.86),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](3404.23),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](3455.4),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5135.61),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            AuthType: components.OutputWebhookAuthenticationTypeCredentialsSecret.ToPointer(),
            TLS: &components.TLSSettingsClientSideType1{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](7737.05),
            LoadBalanced: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("notarize terribly bashfully aw woot"),
            CustomSourceExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomDropWhenNull: criblcontrolplanesdkgo.Pointer(true),
            CustomEventDelimiter: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomContentType: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomPayloadExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            AdvancedContentType: criblcontrolplanesdkgo.Pointer("<value>"),
            FormatEventCode: criblcontrolplanesdkgo.Pointer("<value>"),
            FormatPayloadCode: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](3983.53),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9465.91),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](7319.75),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputWebhookPqControls{},
            Username: criblcontrolplanesdkgo.Pointer("Helene_Brown"),
            Password: criblcontrolplanesdkgo.Pointer("1IQ4If1hcFrV5EE"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://hateful-hutch.name"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](6873.12),
            OauthParams: []components.OauthParam{
                components.OauthParam{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            OauthHeaders: []components.OauthHeader{
                components.OauthHeader{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            URL: criblcontrolplanesdkgo.Pointer("https://example.com/webhook"),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(true),
            Urls: []components.OutputWebhookURL{
                components.OutputWebhookURL{
                    URL: "https://weekly-supplier.org",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](7951.07),
                    TemplateURL: criblcontrolplanesdkgo.Pointer("https://rectangular-accountability.name"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](1730.18),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8338.46),
            TemplateLoginURL: criblcontrolplanesdkgo.Pointer("https://carefree-cork.net"),
            TemplateSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://excitable-bourgeoisie.biz/"),
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
### Example Usage: OutputCreateExamplesWizHec

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesWizHec" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputWizHec(
        components.OutputWizHec{
            ID: criblcontrolplanesdkgo.Pointer("wiz-hec-output"),
            Type: components.OutputWizHecTypeWizHec,
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
            LoadBalanced: "<value>",
            NextQueue: criblcontrolplanesdkgo.Pointer("<value>"),
            TCPRouting: criblcontrolplanesdkgo.Pointer("<value>"),
            TLS: &components.TLSSettingsClientSideType1{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](653.76),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](8807.31),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](2723.31),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](1822.3),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3768.44),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            EnableMultiMetrics: "<value>",
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            WizConnectorID: "00000000-0000-0000-0000-000000000000",
            WizEnvironment: "test",
            DataCenter: "us1",
            WizSourcetype: "placeholder",
            Description: criblcontrolplanesdkgo.Pointer("midst unexpectedly boiling uh-huh reflecting veg pfft deserted excepting"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](7926.84),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7504.19),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](4345),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputWizHecPqControls{},
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateWizEnvironment: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateDataCenter: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateWizSourcetype: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesXsiam

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputCreateExamplesXsiam" -->
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

    res, err := s.Packs.Destinations.Update(ctx, "<id>", "<value>", components.CreateOutputXsiam(
        components.OutputXsiam{
            ID: criblcontrolplanesdkgo.Pointer("xsiam-output"),
            Type: components.OutputXsiamTypeXsiam,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            LoadBalanced: criblcontrolplanesdkgo.Pointer(false),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](8207.39),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1553.6),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](5682.57),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](9456.19),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1832.2),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            AuthType: components.OutputXsiamAuthenticationMethodSecret.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9145.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8049.35),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7506.18),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](8392.23),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](9705.48),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8513.57),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](3688.43),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            ThrottleRateReqPerSec: criblcontrolplanesdkgo.Pointer[int64](668286),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](9706.26),
            Description: criblcontrolplanesdkgo.Pointer("qua cornet outside oil candid seemingly"),
            URL: criblcontrolplanesdkgo.Pointer("https://inexperienced-pigsty.biz/"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Urls: []components.OutputXsiamURL{
                components.OutputXsiamURL{
                    URL: "https://delirious-tackle.info/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1722.99),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](6218.76),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8569.53),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2675.25),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8561.04),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](325.7),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputXsiamPqControls{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://distinct-management.biz/"),
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
### Example Usage: OutputExamplesDefault

<!-- UsageSnippet language="go" operationID="updateOutputSystemByPackAndId" method="patch" path="/p/{pack}/system/outputs/{id}" example="OutputExamplesDefault" -->
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