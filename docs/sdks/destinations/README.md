# Destinations

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

### Example Usage: OutputCreateExamplesAzureBlob

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesAzureBlob" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestAzureBlob(
        operations.OutputAzureBlob{
            ID: "azure-blob-output",
            Type: operations.CreateOutputTypeAzureBlobAzureBlob,
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
            ContainerName: "my-container",
            CreateContainer: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](817.98),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsRaw.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](7972.39),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](8769.44),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](2378.15),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](6434.56),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](646.16),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](8843.66),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4381.89),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.43),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1510.88),
            },
            AuthType: components.AuthenticationMethodOptionsClientCert.ToPointer(),
            StorageClass: operations.BlobAccessTierHot.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("over rectangular why who instead ferociously sidetrack revoke focused futon"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsNormal.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](4324.82),
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
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](5870.02),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](8174.56),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](9754.4),
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
    if res.CountedOutput != nil {
        // handle response
    }
}
```
### Example Usage: OutputCreateExamplesAzureDataExplorer

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesAzureDataExplorer" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestAzureDataExplorer(
        operations.OutputAzureDataExplorer{
            ID: "azure-data-explorer-output",
            Type: operations.TypeAzureDataExplorerAzureDataExplorer,
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
            ClusterURL: "https://mycluster.kusto.windows.net",
            Database: "mydatabase",
            Table: "mytable",
            ValidateDatabaseSettings: criblcontrolplanesdkgo.Pointer(false),
            IngestMode: operations.IngestionModeStreaming.ToPointer(),
            OauthEndpoint: components.MicrosoftEntraIDAuthenticationEndpointOptionsSaslHTTPSLoginMicrosoftonlineCom,
            TenantID: "tenant-id",
            ClientID: "client-id",
            Scope: "https://mycluster.kusto.windows.net/.default",
            OauthType: operations.OauthTypeAuthenticationMethodClientSecret,
            Description: criblcontrolplanesdkgo.Pointer("complete madly simplistic"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("client-secret"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Certificate: &operations.Certificate{
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            Format: components.DataFormatOptionsJSON.ToPointer(),
            Compress: components.CompressionOptions2Gzip,
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet24.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](5991.99),
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
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](3292.42),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1827.67),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](8142.92),
            IsMappingObj: criblcontrolplanesdkgo.Pointer(false),
            MappingObj: criblcontrolplanesdkgo.Pointer("<value>"),
            MappingRef: criblcontrolplanesdkgo.Pointer("<value>"),
            IngestURL: criblcontrolplanesdkgo.Pointer("https://decent-chops.com/"),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            StagePath: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](8975.55),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](1446.39),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](5111.87),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](5440.71),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](9779.14),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](8843.66),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4381.89),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.43),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1510.88),
            },
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](2127.27),
            FlushImmediately: criblcontrolplanesdkgo.Pointer(true),
            RetainBlobOnSuccess: criblcontrolplanesdkgo.Pointer(true),
            ExtentTags: []operations.ExtentTag{
                operations.ExtentTag{
                    Prefix: operations.PrefixOptionalIngestBy.ToPointer(),
                    Value: "<value>",
                },
            },
            IngestIfNotExists: []operations.IngestIfNotExist{
                operations.IngestIfNotExist{
                    Value: "<value>",
                },
            },
            ReportLevel: operations.ReportLevelDoNotReport.ToPointer(),
            ReportMethod: operations.ReportMethodQueueAndTable.ToPointer(),
            AdditionalProperties: []operations.AdditionalProperty{
                operations.AdditionalProperty{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](187.05),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](6017.97),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](7576.63),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5944.28),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8661.98),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](6573.83),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](94.66),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsAzureDataExplorer{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesAzureEventhub" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestAzureEventhub(
        operations.OutputAzureEventhub{
            ID: "azure-eventhub-output",
            Type: operations.TypeAzureEventhubAzureEventhub,
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
            Ack: components.AcknowledgmentsOptionsZero.ToPointer(),
            Format: components.RecordDataFormatOptionsJSON.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](5899.05),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](6638.01),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8055.83),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](9586.96),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](9695.69),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](6088.68),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](6052.69),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](6330.98),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](623.59),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](8813.79),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](1016.21),
            Sasl: &components.AuthenticationType1{
                Disabled: false,
                AuthType: components.AuthenticationMethodOptionsSasl1Manual.ToPointer(),
                Password: criblcontrolplanesdkgo.Pointer("rmAs6PwLK2nbEpz"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSasl1Plain.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("Rhianna_Abshire"),
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
                Disabled: false,
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("gadzooks alive sway sans swiftly gah cemetery"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8523.95),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8483.65),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5247.04),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsAzureEventhub{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesAzureLogs" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestAzureLogs(
        operations.OutputAzureLogs{
            ID: "azure-logs-output",
            Type: operations.TypeAzureLogsAzureLogs,
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
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](9755.46),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](3239.53),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](107.16),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](4538.64),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6010.13),
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
            APIURL: criblcontrolplanesdkgo.Pointer("https://bulky-printer.net"),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: operations.AuthenticationMethodAzureLogsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("sour mallard ah"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1957.33),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1430.65),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8951.12),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsAzureLogs{},
            WorkspaceID: criblcontrolplanesdkgo.Pointer("workspace-id"),
            WorkspaceKey: criblcontrolplanesdkgo.Pointer("workspace-key"),
            KeypairSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesChronicle" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestChronicle(
        operations.OutputChronicle{
            ID: "chronicle-output",
            Type: operations.TypeChronicleChronicle,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            APIVersion: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthenticationMethod: operations.AuthenticationMethodChronicleServiceAccountSecret.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            Region: "us",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](7385.01),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](5430.66),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](6757.95),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](6647.22),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7991.92),
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
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](680.51),
            IngestionMethod: criblcontrolplanesdkgo.Pointer("<value>"),
            Namespace: criblcontrolplanesdkgo.Pointer("<value>"),
            LogType: "UNKNOWN",
            LogTextField: criblcontrolplanesdkgo.Pointer("<value>"),
            GcpProjectID: "my-project",
            GcpInstance: "customer-id",
            CustomLabels: []operations.CustomLabel{
                operations.CustomLabel{
                    Key: "<key>",
                    Value: "<value>",
                    RbacEnabled: criblcontrolplanesdkgo.Pointer(true),
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("think for lonely snappy council reapply onto"),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            ServiceAccountCredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6383.08),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8317.94),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8897.34),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsChronicle{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesClickHouse" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestClickHouse(
        operations.OutputClickHouse{
            ID: "clickhouse-output",
            Type: operations.TypeClickHouseClickHouse,
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
            URL: "http://localhost:8123/",
            AuthType: operations.AuthenticationTypeClickHouseBasic.ToPointer(),
            Database: "mydb",
            TableName: "mytable",
            Format: operations.FormatClickHouseJSONCompactEachRowWithNames.ToPointer(),
            MappingType: operations.MappingTypeAutomatic.ToPointer(),
            AsyncInserts: criblcontrolplanesdkgo.Pointer(true),
            TLS: &components.TLSSettingsClientSideType1{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1394.89),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](5341.73),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](8534.7),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](1388.6),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1978.4),
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
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            DumpFormatErrorsToDisk: criblcontrolplanesdkgo.Pointer(true),
            StatsDestination: &operations.StatsDestination{
                URL: criblcontrolplanesdkgo.Pointer("https://rewarding-government.name"),
                Database: criblcontrolplanesdkgo.Pointer("<value>"),
                TableName: criblcontrolplanesdkgo.Pointer("<value>"),
                AuthType: criblcontrolplanesdkgo.Pointer("<value>"),
                Username: criblcontrolplanesdkgo.Pointer("Nikki20"),
                SQLUsername: criblcontrolplanesdkgo.Pointer("<value>"),
                Password: criblcontrolplanesdkgo.Pointer("MBdiADt_EqwyvJ_"),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("alongside blue doubtfully"),
            Username: criblcontrolplanesdkgo.Pointer("Aaliyah.Mitchell"),
            Password: criblcontrolplanesdkgo.Pointer("j90g4qCNUmopAcY"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://velvety-cross-contamination.biz/"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](6937.57),
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
            SQLUsername: criblcontrolplanesdkgo.Pointer("<value>"),
            WaitForAsyncInserts: criblcontrolplanesdkgo.Pointer(false),
            ExcludeMappingFields: []string{
                "<value 1>",
            },
            DescribeTable: criblcontrolplanesdkgo.Pointer("<value>"),
            ColumnMappings: []operations.ColumnMapping{
                operations.ColumnMapping{
                    ColumnName: "<value>",
                    ColumnType: criblcontrolplanesdkgo.Pointer("<value>"),
                    ColumnValueExpression: "<value>",
                },
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](7198.36),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](552.93),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](405.08),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.PqControlsClickHouse{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesCloudflareR2" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestCloudflareR2(
        operations.OutputCloudflareR2{
            ID: "cloudflare-r2-output",
            Type: operations.TypeCloudflareR2CloudflareR2,
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
            Endpoint: "https://account-id.r2.cloudflarestorage.com",
            Bucket: "my-bucket",
            AwsAuthenticationMethod: operations.AuthenticationMethodCloudflareR2Auto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "<value>",
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions5V2.ToPointer(),
            ObjectACL: "<value>",
            StorageClass: components.StorageClassOptions2ReducedRedundancy.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionOptionsAes256.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(false),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](319.61),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](3509.07),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](4903.23),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](8843.66),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4381.89),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.43),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1510.88),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](8072.1),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](3605.52),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](251.81),
            Description: criblcontrolplanesdkgo.Pointer("truly yippee pulse"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsNormal.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet10.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](3082.8),
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
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](6773.81),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](9350.87),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](5305.64),
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesCloudwatch" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestCloudwatch(
        operations.OutputCloudwatch{
            ID: "cloudwatch-output",
            Type: operations.TypeCloudwatchCloudwatch,
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
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](6305.14),
            MaxQueueSize: criblcontrolplanesdkgo.Pointer[float64](4712.6),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](54.02),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1950.05),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("gah sleepily or selfishly order beautifully stitcher"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5917.34),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9863.81),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5351.88),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsCloudwatch{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesConfluentCloud" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestConfluentCloud(
        operations.OutputConfluentCloud{
            ID: "confluent-cloud-output",
            Type: operations.CreateOutputTypeConfluentCloudConfluentCloud,
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
            Ack: components.AcknowledgmentsOptions1One.ToPointer(),
            Format: components.RecordDataFormatOptions1JSON.ToPointer(),
            Compression: components.CompressionOptions3None.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](6128.64),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](3514.77),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5300.5),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType1{
                Disabled: false,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://key-vicinity.info/"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](9382.36),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](2250.64),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](9302.69),
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
                DefaultKeySchemaID: criblcontrolplanesdkgo.Pointer[float64](1951.32),
                DefaultValueSchemaID: criblcontrolplanesdkgo.Pointer[float64](5499.87),
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](7876.49),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](5343.71),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](8798.66),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](8515.43),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](791.86),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1309.17),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](5121.15),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](557.68),
            Sasl: &components.AuthenticationType{
                Disabled: true,
                Username: criblcontrolplanesdkgo.Pointer("Henry_Kilback"),
                Password: criblcontrolplanesdkgo.Pointer("9FPhvPK4v3rJnwx"),
                AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslPlain.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(true),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://hopeful-help.info"),
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
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("marksman eek when pfft partridge hm terrorise"),
            ProtobufLibraryID: criblcontrolplanesdkgo.Pointer("<id>"),
            ProtobufEncodingID: criblcontrolplanesdkgo.Pointer("<id>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8549.89),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3574.31),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1776.95),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.PqControlsConfluentCloud{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesCriblHttp" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestCriblHTTP(
        operations.OutputCriblHTTP{
            ID: "cribl-http-output",
            Type: operations.CreateOutputTypeCriblHTTPCriblHTTP,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            TokenTTLMinutes: criblcontrolplanesdkgo.Pointer[float64](9316.53),
            ExcludeFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Compression: components.CompressionOptions1None.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1734.52),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4810.45),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](2634.2),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](8177.12),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1294.36),
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
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            AuthTokens: []components.ItemsTypeAuthTokens1{
                components.ItemsTypeAuthTokens1{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("sanity readjust by ah lazily scenario eek why astride er"),
                },
            },
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("capitalize oh lumbering blaring by inside conservative"),
            URL: criblcontrolplanesdkgo.Pointer("https://married-manner.biz/"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(true),
            Urls: []components.ItemsTypeUrls{
                components.ItemsTypeUrls{
                    URL: "https://stupendous-replacement.info/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](3117.53),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](1044),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7719.42),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5564.73),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4108.73),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](7654.09),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsCriblHTTP{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesCriblLake" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestCriblLake(
        operations.OutputCriblLake{
            ID: "cribl-lake-output",
            Type: operations.TypeCriblLakeCriblLake,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
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
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](7816.2),
            StagePath: criblcontrolplanesdkgo.Pointer("<value>"),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            ObjectACL: components.ObjectACLOptionsPrivate.ToPointer(),
            StorageClass: components.StorageClassOptionsOnezoneIa.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAes256.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](2764.14),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](9446.71),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](9500.94),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(true),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](8843.66),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4381.89),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.43),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1510.88),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](2482.69),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](3981.2),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(false),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](7889.42),
            AwsAuthenticationMethod: components.MethodOptionsCredentialsAuto.ToPointer(),
            Format: components.FormatOptionsCriblLakeDatasetParquet.ToPointer(),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](4107.26),
            Description: criblcontrolplanesdkgo.Pointer("what thoroughly apropos nor yahoo whenever midst agreeable gray psst"),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](9887.29),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](9452.98),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](147.99),
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesCriblSearchEngine" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestCriblSearchEngine(
        operations.OutputCriblSearchEngine{
            ID: "cribl-search-engine-output",
            Type: operations.TypeCriblSearchEngineCriblSearchEngine,
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
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](8934.01),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](3173.17),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](5930.45),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            AuthTokens: []components.ItemsTypeAuthTokens1{
                components.ItemsTypeAuthTokens1{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("sanity readjust by ah lazily scenario eek why astride er"),
                },
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("popularity till contravene secrecy ferociously likely louse afore pendant tromp"),
            URL: criblcontrolplanesdkgo.Pointer("https://0.0.0.0:10200"),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(true),
            Urls: []components.ItemsTypeUrls{
                components.ItemsTypeUrls{
                    URL: "https://stupendous-replacement.info/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](3117.53),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](1638.95),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](2145.95),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5680.41),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8837.87),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](3947.28),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsCriblSearchEngine{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesCriblTcp" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestCriblTCP(
        operations.OutputCriblTCP{
            ID: "cribl-tcp-output",
            Type: operations.CreateOutputTypeCriblTCPCriblTCP,
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
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](4118.38),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](5415.47),
            TokenTTLMinutes: criblcontrolplanesdkgo.Pointer[float64](8431.21),
            AuthTokens: []components.ItemsTypeAuthTokens{
                components.ItemsTypeAuthTokens{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                    Description: criblcontrolplanesdkgo.Pointer("overconfidently lovely flight utilized runny outlying till yahoo"),
                },
            },
            ExcludeFields: []string{
                "<value 1>",
            },
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("where sarong that"),
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](10090),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "proper-prohibition.com",
                    Port: 7840.9,
                    TLS: components.TLSOptionsHostsItemsInherit.ToPointer(),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    Weight: criblcontrolplanesdkgo.Pointer[float64](3441.24),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](4237.3),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7779.99),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](6098.03),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](7535.23),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3189.59),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5212.25),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsCriblTCP{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesCrowdstrikeNextGenSiem" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestCrowdstrikeNextGenSiem(
        operations.OutputCrowdstrikeNextGenSiem{
            ID: "crowdstrike-next-gen-siem-output",
            Type: operations.TypeCrowdstrikeNextGenSiemCrowdstrikeNextGenSiem,
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
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](7961.15),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1751.08),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](1418.83),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](1251.24),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](4180.92),
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
            Format: components.RequestFormatOptionsJSON,
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("always unless ha"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8889.52),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8955.12),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5752.98),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsCrowdstrikeNextGenSiem{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesDatabricks" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestDatabricks(
        operations.OutputDatabricks{
            ID: "databricks-output",
            Type: operations.TypeDatabricksDatabricks,
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
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            StagePath: criblcontrolplanesdkgo.Pointer("<value>"),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](8118.95),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](4778.49),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](6500.72),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](2941.77),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](2897.47),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(true),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](8843.66),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4381.89),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.43),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1510.88),
            },
            WorkspaceID: "your-workspace-id",
            Scope: "my-scope",
            ClientID: "your-client-id",
            Catalog: "my-catalog",
            Schema: "my-schema",
            EventsVolumeName: "my-volume",
            ClientTextSecret: "your-client-secret",
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](6499.49),
            Description: criblcontrolplanesdkgo.Pointer("really testing editor often duh impostor"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet10.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV1.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](9741.01),
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
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](1478.99),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](5007.62),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](1451.08),
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesDatadog" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestDatadog(
        operations.OutputDatadog{
            ID: "datadog-output",
            Type: operations.TypeDatadogDatadog,
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
            ContentType: operations.SendLogsAsText.ToPointer(),
            Message: criblcontrolplanesdkgo.Pointer("<value>"),
            Source: criblcontrolplanesdkgo.Pointer("<value>"),
            Host: criblcontrolplanesdkgo.Pointer("devoted-hubris.org"),
            Service: criblcontrolplanesdkgo.Pointer("<value>"),
            Tags: []string{
                "<value 1>",
            },
            BatchByTags: criblcontrolplanesdkgo.Pointer(true),
            AllowAPIKeyFromEvents: criblcontrolplanesdkgo.Pointer(true),
            Severity: operations.SeverityDatadogNotice.ToPointer(),
            Site: operations.DatadogSiteAp1.ToPointer(),
            SendCountersAsCount: criblcontrolplanesdkgo.Pointer(true),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](588.59),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](6147.82),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](8387.71),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](9960.88),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](2547.26),
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
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Secret.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](8205.85),
            Description: criblcontrolplanesdkgo.Pointer("yowza both knottily phooey enlightened except minister uh-huh"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://lawful-instructor.biz/"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](104.73),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7937.61),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5285.84),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsDatadog{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesDataset" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestDataset(
        operations.OutputDataset{
            ID: "dataset-output",
            Type: operations.TypeDatasetDataset,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            MessageField: criblcontrolplanesdkgo.Pointer("<value>"),
            ExcludeFields: []string{
                "<value 1>",
            },
            ServerHostField: criblcontrolplanesdkgo.Pointer("<value>"),
            TimestampField: criblcontrolplanesdkgo.Pointer("<value>"),
            DefaultSeverity: operations.DefaultSeveritySeverityFiner.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            Site: operations.DataSetSiteUs.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](9140.95),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4841.21),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](3237.75),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](6439.19),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1895.3),
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
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Secret.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](5843.82),
            Description: criblcontrolplanesdkgo.Pointer("teeming sting safe healthily spiteful"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://merry-jungle.net"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](7407.27),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](6106.23),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](7491.95),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsDataset{},
            APIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesDiskSpool

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesDiskSpool" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestDiskSpool(
        operations.OutputDiskSpool{
            ID: "disk-spool-output",
            Type: operations.TypeDiskSpoolDiskSpool,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptionsPersistenceGzip.ToPointer(),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("dapper adventurously brood claw fondly or incomplete"),
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesDlS3" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestDlS3(
        operations.OutputDlS3{
            ID: "dl-s3-output",
            Type: operations.TypeDlS3DlS3,
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
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](4379.67),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            ObjectACL: components.ObjectACLOptionsAwsExecRead.ToPointer(),
            StorageClass: components.StorageClassOptionsGlacierIr.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAes256.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            Format: components.DataFormatOptionsRaw.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](9160.62),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](4664.43),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](6340.06),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](8843.66),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4381.89),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.43),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1510.88),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](7933.95),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](6684.23),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](7299.66),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](7832.01),
            PartitioningFields: []string{
                "<value 1>",
            },
            Description: criblcontrolplanesdkgo.Pointer("supposing and volunteer retention gee amnesty"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2None.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet10.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV1.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](4072.64),
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
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](616.57),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](136.88),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](9976.52),
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesDynatraceHttp" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestDynatraceHTTP(
        operations.OutputDynatraceHTTP{
            ID: "dynatrace-http-output",
            Type: operations.TypeDynatraceHTTPDynatraceHTTP,
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
            Method: components.MethodOptionsPatch.ToPointer(),
            KeepAlive: criblcontrolplanesdkgo.Pointer(false),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](4138.07),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](952.7),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](4505.16),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](5365.06),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](279.61),
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
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            AuthType: operations.AuthenticationTypeDynatraceHTTPToken.ToPointer(),
            Format: operations.FormatDynatraceHTTPJSONArray,
            Endpoint: operations.EndpointCloud,
            TelemetryType: operations.TelemetryTypeLogs,
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](4777.53),
            Description: criblcontrolplanesdkgo.Pointer("consequently preregister transparency drat comb potable"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5892.76),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4977.18),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8460.08),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsDynatraceHTTP{},
            Token: criblcontrolplanesdkgo.Pointer("your-api-key"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            EnvironmentID: criblcontrolplanesdkgo.Pointer("<id>"),
            ActiveGateDomain: criblcontrolplanesdkgo.Pointer("<value>"),
            URL: criblcontrolplanesdkgo.Pointer("https://fragrant-requirement.org"),
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesDynatraceOtlp" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestDynatraceOtlp(
        operations.OutputDynatraceOtlp{
            ID: "dynatrace-otlp-output",
            Type: operations.TypeDynatraceOtlpDynatraceOtlp,
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
            Protocol: operations.ProtocolDynatraceOtlpHTTP,
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
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](3795.09),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](6842.82),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](4936.13),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](2905.07),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](580.05),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](8124.71),
            KeepAlive: criblcontrolplanesdkgo.Pointer(false),
            EndpointType: operations.EndpointTypeSaas,
            TokenSecret: "your-token-secret",
            AuthTokenName: criblcontrolplanesdkgo.Pointer("<value>"),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("when than potentially gray adult surprised bludgeon break"),
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
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9215.6),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2893.14),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1267.79),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.PqControlsDynatraceOtlp{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesElastic" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestElastic(
        operations.OutputElastic{
            ID: "elastic-output",
            Type: operations.CreateOutputTypeElasticElastic,
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
            Index: "logs",
            DocType: criblcontrolplanesdkgo.Pointer("<value>"),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1305.13),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1967.39),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](9276.25),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](4811.65),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1565.7),
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
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            ExtraParams: []components.ItemsTypeSaslSaslExtensions{
                components.ItemsTypeSaslSaslExtensions{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Auth: &components.AuthType{
                Disabled: true,
                Username: criblcontrolplanesdkgo.Pointer("Michele_Breitenberg"),
                Password: criblcontrolplanesdkgo.Pointer("e9CmlyugMJSf8Yp"),
                AuthType: components.AuthenticationMethodOptionsAuthTextSecret.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                ManualAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            ElasticVersion: operations.ElasticVersionSix.ToPointer(),
            ElasticPipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            IncludeDocID: criblcontrolplanesdkgo.Pointer(true),
            WriteAction: operations.WriteActionCreate.ToPointer(),
            RetryPartialErrors: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("rue from artistic till unbalance"),
            URL: criblcontrolplanesdkgo.Pointer("https://different-guard.com/"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Urls: []operations.URLElastic{
                operations.URLElastic{
                    URL: "https://crafty-loyalty.name/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1891.06),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](4086.16),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8486.79),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8212.5),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4021.19),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6425.86),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.PqControlsElastic{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesElasticCloud" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestElasticCloud(
        operations.OutputElasticCloud{
            ID: "elastic-cloud-output",
            Type: operations.TypeElasticCloudElasticCloud,
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
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1370.06),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1887.57),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](8629.45),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](2199.13),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6328.54),
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
            ExtraParams: []components.ItemsTypeSaslSaslExtensions{
                components.ItemsTypeSaslSaslExtensions{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Auth: &components.AuthType{
                Disabled: true,
                Username: criblcontrolplanesdkgo.Pointer("Michele_Breitenberg"),
                Password: criblcontrolplanesdkgo.Pointer("e9CmlyugMJSf8Yp"),
                AuthType: components.AuthenticationMethodOptionsAuthTextSecret.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                ManualAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            ElasticPipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            IncludeDocID: criblcontrolplanesdkgo.Pointer(true),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("baptise greedily furthermore silently labourer overcharge"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](7344.96),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8050.9),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](4826.17),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.PqControlsElasticCloud{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesExabeam" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestExabeam(
        operations.OutputExabeam{
            ID: "exabeam-output",
            Type: operations.TypeExabeamExabeam,
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
            StagePath: "/tmp/staging",
            Endpoint: "https://storage.googleapis.com",
            SignatureVersion: components.SignatureVersionOptions4V4.ToPointer(),
            ObjectACL: components.ObjectACLOptions1BucketOwnerFullControl.ToPointer(),
            StorageClass: components.StorageClassOptions1Nearline.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](5346.31),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](6694.75),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](7124.04),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](8843.66),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4381.89),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.43),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1510.88),
            },
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](3929.93),
            EncodedConfiguration: criblcontrolplanesdkgo.Pointer("<value>"),
            CollectorInstanceID: "11112222-3333-4444-5555-666677778888",
            SiteName: criblcontrolplanesdkgo.Pointer("<value>"),
            SiteID: criblcontrolplanesdkgo.Pointer("<id>"),
            TimezoneOffset: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("times ouch qua chapel"),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](8482.7),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](9688.35),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](9664.78),
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesFilesystem" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestFilesystem(
        operations.OutputFilesystem{
            ID: "filesystem-output",
            Type: operations.TypeFilesystemFilesystem,
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
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsRaw.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](6295.91),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](9272.25),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](8431.64),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](7740.89),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](5522.82),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(true),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](8843.66),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4381.89),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.43),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1510.88),
            },
            Description: criblcontrolplanesdkgo.Pointer("geez psst knowingly yum oof uh-huh rekindle"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsNormal.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet24.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](2499.28),
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
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](9527.71),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](7326.8),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](5984.31),
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesGoogleChronicle" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestGoogleChronicle(
        operations.OutputGoogleChronicle{
            ID: "google-chronicle-output",
            Type: operations.TypeGoogleChronicleGoogleChronicle,
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
            APIVersion: operations.CreateOutputAPIVersionV1.ToPointer(),
            AuthenticationMethod: operations.AuthenticationMethodGoogleChronicleServiceAccountSecret.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            LogFormatType: operations.SendEventsAsUnstructured,
            Region: criblcontrolplanesdkgo.Pointer("us"),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](8794.34),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1682.79),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](3320.22),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](1091.39),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](980.05),
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
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](7865.63),
            Description: criblcontrolplanesdkgo.Pointer("immediate kindly rival whoa madly rightfully selfishly aw eternity sternly"),
            ExtraLogTypes: []operations.ExtraLogType{
                operations.ExtraLogType{
                    LogType: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("until including whether whenever loosely drat excepting"),
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
            UdmType: operations.UDMTypeEntities.ToPointer(),
            APIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            APIKeySecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            ServiceAccountCredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](916.07),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1786.97),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](2608.15),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.PqControlsGoogleChronicle{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesGoogleCloudLogging" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestGoogleCloudLogging(
        operations.OutputGoogleCloudLogging{
            ID: "google-cloud-logging-output",
            Type: operations.TypeGoogleCloudLoggingGoogleCloudLogging,
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
            LogLocationType: operations.LogLocationTypeProject,
            LogNameExpression: "my-log",
            SanitizeLogNames: criblcontrolplanesdkgo.Pointer(false),
            PayloadFormat: operations.PayloadFormatText.ToPointer(),
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
            GoogleAuthMethod: components.GoogleAuthenticationMethodOptionsManual.ToPointer(),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](132.27),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](7232.46),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](4246.62),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5003.58),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](3406.78),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](9245.78),
            ThrottleRateReqPerSec: criblcontrolplanesdkgo.Pointer[int64](815464),
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
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](4191.17),
            Description: criblcontrolplanesdkgo.Pointer("rarely kielbasa unfortunate motionless service"),
            LogLocationExpression: "my-project",
            PayloadExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9958.44),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9384.34),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](495.29),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.PqControlsGoogleCloudLogging{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesGoogleCloudStorage" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestGoogleCloudStorage(
        operations.OutputGoogleCloudStorage{
            ID: "google-cloud-storage-output",
            Type: operations.TypeGoogleCloudStorageGoogleCloudStorage,
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
            Region: "us-east1",
            Endpoint: "https://storage.googleapis.com",
            SignatureVersion: components.SignatureVersionOptions4V4.ToPointer(),
            AwsAuthenticationMethod: operations.AuthenticationMethodGoogleCloudStorageManual.ToPointer(),
            StagePath: "/tmp/staging",
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            ObjectACL: components.ObjectACLOptions1PublicRead.ToPointer(),
            StorageClass: components.StorageClassOptions1Nearline.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsParquet.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](2201.65),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](2817.56),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](3331.7),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](7260.53),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](9134.23),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](8843.66),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4381.89),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.43),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1510.88),
            },
            Description: criblcontrolplanesdkgo.Pointer("before sonata opera"),
            Compress: components.CompressionOptions2None.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsNormal.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV1.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](2379.27),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(false),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(false),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](2944.76),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](4948.28),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](6713.99),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesGooglePubsub" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestGooglePubsub(
        operations.OutputGooglePubsub{
            ID: "google-pubsub-output",
            Type: operations.CreateOutputTypeGooglePubsubGooglePubsub,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            TopicName: "my-topic",
            CreateTopic: criblcontrolplanesdkgo.Pointer(false),
            OrderedDelivery: criblcontrolplanesdkgo.Pointer(false),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            GoogleAuthMethod: components.GoogleAuthenticationMethodOptionsSecret.ToPointer(),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            BatchSize: criblcontrolplanesdkgo.Pointer[float64](9392.71),
            BatchTimeout: criblcontrolplanesdkgo.Pointer[float64](4367.52),
            MaxQueueSize: criblcontrolplanesdkgo.Pointer[float64](4966.84),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](6353.7),
            FlushPeriod: criblcontrolplanesdkgo.Pointer[float64](1345.87),
            MaxInProgress: criblcontrolplanesdkgo.Pointer[float64](6886.69),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("how corrupt wherever whoa till splay gosh phew chiffonier"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1312.67),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5116.8),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8037.49),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.PqControlsGooglePubsub{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesGrafanaCloud" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestGrafanaCloud(
        operations.CreateOutputGrafanaCloudOutputGrafanaCloudGrafanaCloud1(
            operations.OutputGrafanaCloudGrafanaCloud1{
                ID: "grafana-cloud-output",
                Type: operations.OutputGrafanaCloudType1GrafanaCloud,
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
                LokiURL: "https://logs-prod-us-central1.grafana.net",
                PrometheusURL: criblcontrolplanesdkgo.Pointer("https://elementary-window.com/"),
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
                    AuthType: components.AuthenticationTypeOptionsPrometheusAuth1Token.ToPointer(),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Username: criblcontrolplanesdkgo.Pointer("Bettye.Bayer10"),
                    Password: criblcontrolplanesdkgo.Pointer("6WzcIgopzuZrB6K"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                LokiAuth: &components.PrometheusAuthType{
                    AuthType: components.AuthenticationTypeOptionsPrometheusAuth1None.ToPointer(),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Username: criblcontrolplanesdkgo.Pointer("Tremaine_Brakus"),
                    Password: criblcontrolplanesdkgo.Pointer("z_x0kTxamYhKAOP"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                Concurrency: criblcontrolplanesdkgo.Pointer[float64](8743.58),
                MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](6916.2),
                MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](6975.38),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](1530.06),
                FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8821.56),
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
                        HTTPStatus: 5395.85,
                        InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                        BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                        MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                    },
                },
                TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                    TimeoutRetry: false,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
                },
                ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
                OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
                Description: criblcontrolplanesdkgo.Pointer("editor lovable confound which"),
                Compress: criblcontrolplanesdkgo.Pointer(true),
                PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
                PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](3235.66),
                PqMode: components.ModeOptionsBackpressure.ToPointer(),
                PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2947.95),
                PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5495.79),
                PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
                PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
                PqControls: &operations.OutputGrafanaCloudPqControls1{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesGraphite" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestGraphite(
        operations.OutputGraphite{
            ID: "graphite-output",
            Type: operations.TypeGraphiteGraphite,
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
            Protocol: components.DestinationProtocolOptionsTCP,
            Host: "localhost",
            Port: 2003,
            Mtu: criblcontrolplanesdkgo.Pointer[float64](9598.99),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3418.54),
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](3000.3),
            Description: criblcontrolplanesdkgo.Pointer("gift before hollow overspend overload along story"),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](7954.7),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](7219.26),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9757.91),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5945.93),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8516.8),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.PqControlsGraphite{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesHoneycomb" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestHoneycomb(
        operations.OutputHoneycomb{
            ID: "honeycomb-output",
            Type: operations.TypeHoneycombHoneycomb,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Dataset: "my-dataset",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](9249.02),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](8512.21),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](1633.42),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](4547.48),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3921.68),
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
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Secret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("suffice between but"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](3167.42),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8550.61),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](7828.93),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.PqControlsHoneycomb{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesHumioHec" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestHumioHec(
        operations.OutputHumioHec{
            ID: "humio-hec-output",
            Type: operations.TypeHumioHecHumioHec,
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
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](9133.89),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](6782.11),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](3704.43),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](6769.12),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3959.36),
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
            },
            Format: components.RequestFormatOptionsJSON,
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("aside strong promptly who youthfully pertain meh cleverly under psst"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5830.87),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](6991.72),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](825.98),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.PqControlsHumioHec{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesInfluxdb" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestInfluxdb(
        operations.OutputInfluxdb{
            ID: "influxdb-output",
            Type: operations.TypeInfluxdbInfluxdb,
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
            URL: "http://localhost:8086",
            UseV2API: criblcontrolplanesdkgo.Pointer(false),
            TimestampPrecision: operations.TimestampPrecisionM.ToPointer(),
            DynamicValueFieldName: criblcontrolplanesdkgo.Pointer(false),
            ValueFieldName: criblcontrolplanesdkgo.Pointer("<value>"),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](3866.76),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](6941.21),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](7388.05),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](6234.27),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](9162.97),
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
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: operations.AuthenticationTypeInfluxdbBasic.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("plus micromanage chunter"),
            Database: criblcontrolplanesdkgo.Pointer("mydb"),
            Bucket: criblcontrolplanesdkgo.Pointer("<value>"),
            Org: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9349.14),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3818.09),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6251.66),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsInfluxdb{},
            Username: criblcontrolplanesdkgo.Pointer("Greg_Gibson46"),
            Password: criblcontrolplanesdkgo.Pointer("icfAwtCm5h2Eklm"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://reasonable-steeple.biz/"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](788.11),
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
    if res.CountedOutput != nil {
        // handle response
    }
}
```
### Example Usage: OutputCreateExamplesKafka

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesKafka" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestKafka(
        operations.OutputKafka{
            ID: "kafka-output",
            Type: operations.CreateOutputTypeKafkaKafka,
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
            Brokers: []string{
                "localhost:9092",
            },
            Topic: "logs",
            Ack: components.AcknowledgmentsOptions1One.ToPointer(),
            Format: components.RecordDataFormatOptions1Raw.ToPointer(),
            Compression: components.CompressionOptions3Snappy.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](8326.98),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](3966.84),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](2289.98),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType1{
                Disabled: false,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://key-vicinity.info/"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](9382.36),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](2250.64),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](9302.69),
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
                DefaultKeySchemaID: criblcontrolplanesdkgo.Pointer[float64](1951.32),
                DefaultValueSchemaID: criblcontrolplanesdkgo.Pointer[float64](5499.87),
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](2578.59),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](1385.62),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](1557.54),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](511.47),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](3730.07),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](9302.43),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](3768.04),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](1285.89),
            Sasl: &components.AuthenticationType{
                Disabled: true,
                Username: criblcontrolplanesdkgo.Pointer("Henry_Kilback"),
                Password: criblcontrolplanesdkgo.Pointer("9FPhvPK4v3rJnwx"),
                AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslPlain.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(true),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://hopeful-help.info"),
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
            Description: criblcontrolplanesdkgo.Pointer("realistic pale alongside"),
            ProtobufLibraryID: criblcontrolplanesdkgo.Pointer("<id>"),
            ProtobufEncodingID: criblcontrolplanesdkgo.Pointer("<id>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](7938.77),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5979.28),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](7434.96),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsKafka{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesKinesis" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestKinesis(
        operations.OutputKinesis{
            ID: "kinesis-output",
            Type: operations.CreateOutputTypeKinesisKinesis,
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
            StreamName: "my-stream",
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfSecret.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions2V2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](2865.7),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](3316.35),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](7875.89),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](730.91),
            Compression: operations.CreateOutputCompressionGzip.ToPointer(),
            UseListShards: criblcontrolplanesdkgo.Pointer(true),
            AsNdjson: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("questionably freight gah tune if for lest before"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxEventsPerFlush: criblcontrolplanesdkgo.Pointer[float64](4451.99),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](579.8),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4912.82),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](2288.83),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsKinesis{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesLoki" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestLoki(
        operations.OutputLoki{
            ID: "loki-output",
            Type: operations.CreateOutputTypeLokiLoki,
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
            URL: "http://localhost:3100/loki/api/v1/push",
            Message: criblcontrolplanesdkgo.Pointer("<value>"),
            MessageFormat: components.MessageFormatOptionsJSON.ToPointer(),
            Labels: []components.ItemsTypeLabels{
                components.ItemsTypeLabels{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            AuthType: components.AuthenticationTypeOptionsPrometheusAuth1Basic.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1442.96),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](3308.55),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](3352.1),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](4964.43),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3137.64),
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
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            EnableDynamicHeaders: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](4591.93),
            Description: criblcontrolplanesdkgo.Pointer("however valley gloomy annually belabor"),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Username: criblcontrolplanesdkgo.Pointer("Zola.Collier"),
            Password: criblcontrolplanesdkgo.Pointer("9kviy8fvS3TbTTq"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](3850.39),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5054.27),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](3645.95),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.PqControlsLoki{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesMicrosoftFabric" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestMicrosoftFabric(
        operations.OutputMicrosoftFabric{
            ID: "microsoft-fabric-output",
            Type: operations.TypeMicrosoftFabricMicrosoftFabric,
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
            Topic: "logs",
            Ack: components.AcknowledgmentsOptionsOne.ToPointer(),
            Format: components.RecordDataFormatOptionsRaw.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](7982.47),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](262.82),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](453.27),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](111.17),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](5512.73),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](2681.7),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](1343.6),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5488.51),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](864),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](4128.18),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](3723.97),
            Sasl: &operations.Authentication{
                Disabled: true,
                Mechanism: components.SaslMechanismOptionsSasl1Oauthbearer.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("Golden.Dare10"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                ClientSecretAuthType: operations.ClientSecretAuthTypeAuthenticationMethodCertificate.ToPointer(),
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
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            BootstrapServer: "myeventstream.servicebus.windows.net:9093",
            Description: criblcontrolplanesdkgo.Pointer("hmph granny backbone outside hover huzzah provided eyeliner"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9266.94),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3373.93),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5157.19),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsMicrosoftFabric{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesMinio" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestMinio(
        operations.OutputMinio{
            ID: "minio-output",
            Type: operations.TypeMinioMinio,
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
            Endpoint: "http://localhost:9000",
            Bucket: "my-bucket",
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions5V2.ToPointer(),
            ObjectACL: components.ObjectACLOptionsAwsExecRead.ToPointer(),
            StorageClass: components.StorageClassOptions2Standard.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionOptionsAes256.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsParquet.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](3987.39),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](9617.76),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](2341.4),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](8843.66),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4381.89),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.43),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1510.88),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](5896.42),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](9704.79),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](3365.46),
            Description: criblcontrolplanesdkgo.Pointer("merge when fatally finally powerfully"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](5856.71),
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
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](9288.59),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1730.32),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](8687.49),
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesMsk" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestMsk(
        operations.OutputMsk{
            ID: "msk-output",
            Type: operations.CreateOutputTypeMskMsk,
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
            Brokers: []string{
                "b-1.example.xxxxx.c2.kafka.us-east-1.amazonaws.com:9092",
            },
            Topic: "logs",
            Ack: components.AcknowledgmentsOptions1Zero.ToPointer(),
            Format: components.RecordDataFormatOptions1Protobuf.ToPointer(),
            Compression: components.CompressionOptions3Snappy.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](6929.65),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](526.08),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8491.75),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType1{
                Disabled: false,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://key-vicinity.info/"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](9382.36),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](2250.64),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](9302.69),
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
                DefaultKeySchemaID: criblcontrolplanesdkgo.Pointer[float64](1951.32),
                DefaultValueSchemaID: criblcontrolplanesdkgo.Pointer[float64](5499.87),
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](1259.19),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](8147.21),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](397.36),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](8757.44),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2101.9),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](938.78),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](8388.4),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](4146.44),
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
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](4973.23),
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
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("warmly reassuringly hm shrilly freezing yet cycle"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ProtobufLibraryID: criblcontrolplanesdkgo.Pointer("<id>"),
            ProtobufEncodingID: criblcontrolplanesdkgo.Pointer("<id>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](3594.47),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](232.42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](3423.62),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.PqControlsMsk{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesNetflow" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestNetflow(
        operations.OutputNetflow{
            ID: "netflow-output",
            Type: operations.CreateOutputTypeNetflowNetflow,
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
            Hosts: []operations.HostNetflow{
                operations.HostNetflow{
                    Host: "localhost",
                    Port: 2055,
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](1950.72),
            EnableIPSpoofing: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("as vamoose publication"),
            MaxRecordSize: criblcontrolplanesdkgo.Pointer[float64](8108.13),
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesNewrelic" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestNewrelic(
        operations.OutputNewrelic{
            ID: "newrelic-output",
            Type: operations.TypeNewrelicNewrelic,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Region: components.RegionOptionsUs.ToPointer(),
            LogType: criblcontrolplanesdkgo.Pointer("<value>"),
            MessageField: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []operations.Metadatum{
                operations.Metadatum{
                    Name: operations.FieldNameAuditID,
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](7584.21),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4367.44),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](2339.07),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](1441.55),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7102.29),
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
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Manual.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](2272.57),
            Description: criblcontrolplanesdkgo.Pointer("correctly blindly pfft term"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://great-shipper.com/"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](838.73),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5066.07),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1873.12),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsNewrelic{},
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
### Example Usage: OutputCreateExamplesNewrelicEvents

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesNewrelicEvents" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestNewrelicEvents(
        operations.OutputNewrelicEvents{
            ID: "newrelic-events-output",
            Type: operations.TypeNewrelicEventsNewrelicEvents,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Region: components.RegionOptionsEu.ToPointer(),
            AccountID: "123456",
            EventType: "CriblEvent",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1431.38),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4941.98),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](7424.99),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](768.04),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](4832.37),
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
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Secret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("straw eulogise in dandelion yowza er until state"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://thrifty-sprinkles.biz"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1465.37),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7267.06),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](2016.21),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsNewrelicEvents{},
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
### Example Usage: OutputCreateExamplesOpenTelemetry

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesOpenTelemetry" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestOpenTelemetry(
        operations.OutputOpenTelemetry{
            ID: "opentelemetry-output",
            Type: operations.CreateOutputTypeOpenTelemetryOpenTelemetry,
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
            Protocol: components.ProtocolOptionsHTTP.ToPointer(),
            Endpoint: "http://localhost:4317",
            OtlpVersion: operations.CreateOutputOTLPVersionZeroDot10Dot0.ToPointer(),
            Compress: components.CompressionOptions4None.ToPointer(),
            HTTPCompress: components.CompressionOptions5Gzip.ToPointer(),
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
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](9299.29),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](739.01),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](6855.16),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8781.54),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](7870.65),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](8054.52),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("across toward sour ugh rightfully cash grown inasmuch cycle"),
            Username: criblcontrolplanesdkgo.Pointer("Estell30"),
            Password: criblcontrolplanesdkgo.Pointer("vbPEwCUZoqSnXXn"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://strange-farm.net/"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](4520.14),
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
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            TLS: &components.TLSSettingsClientSideType2{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](873.64),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](296.68),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](7779.3),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.PqControlsOpenTelemetry{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesPrometheus" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestPrometheus(
        operations.OutputPrometheus{
            ID: "prometheus-output",
            Type: operations.CreateOutputTypePrometheusPrometheus,
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
            URL: "http://localhost:9091/api/v1/write",
            MetricRenameExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            SendMetadata: criblcontrolplanesdkgo.Pointer(true),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1078.81),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](803.05),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](5323.37),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](6478.51),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7969.18),
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
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            AuthType: components.AuthenticationTypeOptionsPrometheusAuthNone.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("including hello gadzooks jubilantly over silk switch huzzah curse"),
            MetricsFlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5827.81),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](3738.43),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1906.14),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8575.56),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsPrometheus{},
            Username: criblcontrolplanesdkgo.Pointer("Felipa_Yundt"),
            Password: criblcontrolplanesdkgo.Pointer("8mQZMIxN1Wp8T1b"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://complete-feather.com/"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](4974.58),
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
    if res.CountedOutput != nil {
        // handle response
    }
}
```
### Example Usage: OutputCreateExamplesRing

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesRing" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestRing(
        operations.OutputRing{
            ID: "ring-output",
            Type: operations.TypeRingRing,
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
            Format: operations.DataFormatRingRaw.ToPointer(),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.DataCompressionFormatOptionsPersistenceNone.ToPointer(),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("rebound softly instead"),
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesRouter" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestRouter(
        operations.OutputRouter{
            ID: "router-output",
            Type: operations.TypeRouterRouter,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Rules: []operations.CreateOutputRule{
                operations.CreateOutputRule{
                    Filter: "true",
                    Output: "my-output",
                    Description: criblcontrolplanesdkgo.Pointer("pfft acidly phew selfishly premise"),
                    Final: criblcontrolplanesdkgo.Pointer(true),
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("scare sans upon lively ick"),
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesS3" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestS3(
        operations.OutputS3{
            ID: "s3-output",
            Type: operations.CreateOutputTypeS3S3,
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
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](4729.35),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            ObjectACL: components.ObjectACLOptionsPublicRead.ToPointer(),
            StorageClass: components.StorageClassOptionsStandard.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAes256.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](9278.48),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](2297.39),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](8599.6),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](8843.66),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4381.89),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.43),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1510.88),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](8481.62),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](9149.45),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](4846.03),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](8983.89),
            Description: criblcontrolplanesdkgo.Pointer("ultimately furthermore sauerkraut disappointment plain a considering"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestCompression.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV1.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](8643.32),
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
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](6295.64),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](6971.16),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](7249.51),
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesSecurityLake" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestSecurityLake(
        operations.OutputSecurityLake{
            ID: "security-lake-output",
            Type: operations.CreateOutputTypeSecurityLakeSecurityLake,
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
            Bucket: "my-bucket",
            Region: "us-east-1",
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfSecret.ToPointer(),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: operations.SignatureVersionSecurityLakeV2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: "arn:aws:iam::123456789012:role/my-role",
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](1202.81),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            ObjectACL: components.ObjectACLOptionsBucketOwnerRead.ToPointer(),
            StorageClass: components.StorageClassOptionsDeepArchive.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAes256.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](5039.66),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](8178.32),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](1710.78),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](8843.66),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](4381.89),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.43),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](1510.88),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](3271.05),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](3715),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](8180.46),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(false),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](4917.88),
            AccountID: "123456789012",
            CustomSource: "my-custom-source",
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetVersion: components.ParquetVersionOptionsParquet24.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](62.42),
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
            Description: criblcontrolplanesdkgo.Pointer("graft huzzah huzzah next"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](7941.59),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](6210.2),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](6171.85),
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesSentinel" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestSentinel(
        operations.OutputSentinel{
            ID: "sentinel-output",
            Type: operations.TypeSentinelSentinel,
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
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](9907.17),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](8477.7),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](2168.47),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](2411.96),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](4397.69),
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
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            AuthType: operations.AuthTypeOauth.ToPointer(),
            LoginURL: "https://login.microsoftonline.com",
            Secret: "client-secret",
            ClientID: "client-id",
            Scope: criblcontrolplanesdkgo.Pointer("<value>"),
            EndpointURLConfiguration: operations.EndpointConfigurationURL,
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](2131.19),
            Description: criblcontrolplanesdkgo.Pointer("below opposite scramble evenly dress"),
            Format: operations.FormatSentinelJSONArray.ToPointer(),
            CustomSourceExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomDropWhenNull: criblcontrolplanesdkgo.Pointer(true),
            CustomEventDelimiter: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomContentType: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomPayloadExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            AdvancedContentType: criblcontrolplanesdkgo.Pointer("<value>"),
            FormatEventCode: criblcontrolplanesdkgo.Pointer("<value>"),
            FormatPayloadCode: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2413.07),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7786.01),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5513.26),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.PqControlsSentinel{},
            URL: criblcontrolplanesdkgo.Pointer("https://your-workspace.ingest.monitor.azure.com"),
            DcrID: criblcontrolplanesdkgo.Pointer("<id>"),
            DceEndpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            StreamName: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesSentinelOneAiSiem" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestSentinelOneAiSiem(
        operations.OutputSentinelOneAiSiem{
            ID: "sentinel-one-ai-siem-output",
            Type: operations.TypeSentinelOneAiSiemSentinelOneAiSiem,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Region: operations.RegionUs,
            Endpoint: operations.AISIEMEndpointPathRootServicesCollectorEvent,
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](7007.43),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](6941.56),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](9618.6),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](8044.56),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7268.3),
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
                "<value 3>",
            },
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("couch honored likewise underplay slimy likewise sneak besides fence off"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            BaseURL: criblcontrolplanesdkgo.Pointer("https://happy-solution.name"),
            HostExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            SourceExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            SourceTypeExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceCategoryExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceNameExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceVendorExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            EventTypeExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            Host: criblcontrolplanesdkgo.Pointer("vengeful-exasperation.name"),
            Source: criblcontrolplanesdkgo.Pointer("<value>"),
            SourceType: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceCategory: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceName: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceVendor: criblcontrolplanesdkgo.Pointer("<value>"),
            EventType: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6867.33),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3761.04),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1947.27),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.PqControlsSentinelOneAiSiem{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesServiceNow" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestServiceNow(
        operations.OutputServiceNow{
            ID: "servicenow-output",
            Type: operations.TypeServiceNowServiceNow,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
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
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1094.31),
            Protocol: components.ProtocolOptionsHTTP,
            Compress: components.CompressionOptions4None.ToPointer(),
            HTTPCompress: components.CompressionOptions5Gzip.ToPointer(),
            HTTPTracesEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPMetricsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPLogsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](7051.95),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](6705.96),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7986.11),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayload.ToPointer(),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](530.01),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](3262.39),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("iterate nifty oof understated ouch since who for down sheepishly"),
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
                "<value 2>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            TLS: &components.TLSSettingsClientSideType2{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8940.9),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7799.76),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](3880.07),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsServiceNow{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesSignalfx" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestSignalfx(
        operations.OutputSignalfx{
            ID: "signalfx-output",
            Type: operations.TypeSignalfxSignalfx,
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
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
            Realm: "us0",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](3746.6),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](6857.66),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](8696.95),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](2650.83),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3858.28),
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
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("forceful underneath times yum partially engender digestive cafe"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8080.94),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9607.51),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8560.98),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsSignalfx{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesSnmp" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestSnmp(
        operations.OutputSnmp{
            ID: "snmp-output",
            Type: operations.CreateOutputTypeSnmpSnmp,
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
            Hosts: []operations.HostSnmp{
                operations.HostSnmp{
                    Host: "192.168.1.1",
                    Port: 161,
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](7808.33),
            Description: criblcontrolplanesdkgo.Pointer("rosin before coexist but voluntarily hard-to-find gradient from gad"),
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesSns" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestSns(
        operations.OutputSns{
            ID: "sns-output",
            Type: operations.TypeSnsSns,
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
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](7254.4),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfManual.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: operations.SignatureVersionSnsV2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](5491.06),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("hm godfather hexagon whispered impanel"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](3592.3),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5329.21),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](3899.81),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsSns{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesSplunk" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestSplunk(
        operations.OutputSplunk{
            ID: "splunk-output",
            Type: operations.CreateOutputTypeSplunkSplunk,
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
            Host: "localhost",
            Port: 9997,
            NestedFields: components.NestedFieldSerializationOptionsNone.ToPointer(),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](1113.22),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](191.97),
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
            EnableACK: criblcontrolplanesdkgo.Pointer(true),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(true),
            MaxS2Sversion: components.MaxS2SVersionOptionsV3.ToPointer(),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("ouch favorite patroller neatly tenement yahoo oh unless charter"),
            MaxFailedHealthChecks: criblcontrolplanesdkgo.Pointer[float64](3081.05),
            Compress: components.CompressionOptionsDisabled.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2872.5),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4362.89),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5381.14),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.PqControlsSplunk{},
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
### Example Usage: OutputCreateExamplesSplunkHec

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesSplunkHec" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestSplunkHec(
        operations.OutputSplunkHec{
            ID: "splunk-hec-output",
            Type: operations.CreateOutputTypeSplunkHecSplunkHec,
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
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](8828.22),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](187.91),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](2994.99),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](8657.92),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](598.7),
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
            EnableMultiMetrics: criblcontrolplanesdkgo.Pointer(true),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("though whoa lasting granular"),
            URL: criblcontrolplanesdkgo.Pointer("https://inexperienced-hamburger.org"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(true),
            Urls: []operations.URLSplunkHec{
                operations.URLSplunkHec{
                    URL: "https://pure-rim.info/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](2627.65),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](2868.7),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](9466.46),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1510.16),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](210.45),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1881.8),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsSplunkHec{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesSplunkLb" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestSplunkLb(
        operations.OutputSplunkLb{
            ID: "splunk-lb-output",
            Type: operations.TypeSplunkLbSplunkLb,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](8883.23),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](803.24),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](2177.23),
            NestedFields: components.NestedFieldSerializationOptionsJSON.ToPointer(),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](6003.32),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](4455.31),
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
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            IndexerDiscovery: criblcontrolplanesdkgo.Pointer(false),
            SenderUnhealthyTimeAllowance: criblcontrolplanesdkgo.Pointer[float64](1542.82),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("upward pile since bowed gazebo meanwhile victoriously founder because"),
            MaxFailedHealthChecks: criblcontrolplanesdkgo.Pointer[float64](5051.06),
            Compress: components.CompressionOptionsAlways.ToPointer(),
            IndexerDiscoveryConfigs: &operations.IndexerDiscoveryConfigs{
                Site: "<value>",
                MasterURI: "https://round-hubris.org/",
                RefreshIntervalSec: 2684.26,
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                AuthTokens: []operations.CreateOutputAuthToken{
                    operations.CreateOutputAuthToken{
                        AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
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
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1026.01),
                },
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6241.71),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9335.28),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](391.14),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.PqControlsSplunkLb{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesSqs" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestSqs(
        operations.OutputSqs{
            ID: "sqs-output",
            Type: operations.CreateOutputTypeSqsSqs,
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
            QueueName: "my-queue",
            QueueType: operations.CreateOutputQueueTypeStandard,
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            MessageGroupID: criblcontrolplanesdkgo.Pointer("<id>"),
            CreateQueue: criblcontrolplanesdkgo.Pointer(true),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfSecret.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions3V2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](190.66),
            MaxQueueSize: criblcontrolplanesdkgo.Pointer[float64](5269.78),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](3238.61),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5928.95),
            MaxInProgress: criblcontrolplanesdkgo.Pointer[float64](7767.45),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("inasmuch well-worn fidget gah yippee"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](4949.29),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1745.83),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6622.91),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsSqs{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesStatsd" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestStatsd(
        operations.OutputStatsd{
            ID: "statsd-output",
            Type: operations.TypeStatsdStatsd,
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
            Protocol: components.DestinationProtocolOptionsUDP,
            Host: "localhost",
            Port: 8125,
            Mtu: criblcontrolplanesdkgo.Pointer[float64](3093.84),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](532.77),
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](2433.65),
            Description: criblcontrolplanesdkgo.Pointer("inasmuch deflect sternly daintily except"),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](8194.81),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](1414.14),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2050.15),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5658.33),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5763.25),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsStatsd{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesStatsdExt" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestStatsdExt(
        operations.OutputStatsdExt{
            ID: "statsd-ext-output",
            Type: operations.TypeStatsdExtStatsdExt,
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
            Protocol: components.DestinationProtocolOptionsUDP,
            Host: "localhost",
            Port: 8125,
            Mtu: criblcontrolplanesdkgo.Pointer[float64](1704.38),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](4478.02),
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](3892.36),
            Description: criblcontrolplanesdkgo.Pointer("instead towards stake follower for filter"),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](7169.4),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](8487.93),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9928.78),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9189.36),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](311.77),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsStatsdExt{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesSumoLogic" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestSumoLogic(
        operations.OutputSumoLogic{
            ID: "sumo-logic-output",
            Type: operations.TypeSumoLogicSumoLogic,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            URL: "https://endpoint1.collection.us2.sumologic.com",
            CustomSource: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomCategory: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: operations.DataFormatSumoLogicRaw.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5652.68),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](3678.55),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](3287.58),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](5510.66),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1412.4),
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
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](2679.88),
            Description: criblcontrolplanesdkgo.Pointer("simplistic oof satirize inquisitively twist"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](339.67),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5075.58),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](7425.95),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsSumoLogic{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesSyslog" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestSyslog(
        operations.OutputSyslog{
            ID: "syslog-output",
            Type: operations.CreateOutputTypeSyslogSyslog,
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
            Protocol: operations.ProtocolSyslogTCP.ToPointer(),
            Facility: operations.FacilityFifteen.ToPointer(),
            Severity: operations.SeveritySyslogSeven.ToPointer(),
            AppName: criblcontrolplanesdkgo.Pointer("<value>"),
            MessageFormat: operations.MessageFormatRfc3164.ToPointer(),
            TimestampFormat: operations.TimestampFormatIso8601.ToPointer(),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            OctetCountFraming: criblcontrolplanesdkgo.Pointer(false),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("unique hence rationale"),
            LoadBalanced: criblcontrolplanesdkgo.Pointer(true),
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](514),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(true),
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "proper-prohibition.com",
                    Port: 7840.9,
                    TLS: components.TLSOptionsHostsItemsInherit.ToPointer(),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    Weight: criblcontrolplanesdkgo.Pointer[float64](3441.24),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](2109.55),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1344.55),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](8655.62),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](1852.7),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](3684.02),
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
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            MaxRecordSize: criblcontrolplanesdkgo.Pointer[float64](2177.9),
            UDPDNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](3317.49),
            EnableIPSpoofing: criblcontrolplanesdkgo.Pointer(true),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](7594.37),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9174.33),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](4451.17),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.PqControlsSyslog{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesTcpjson" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestTcpjson(
        operations.OutputTcpjson{
            ID: "tcpjson-output",
            Type: operations.CreateOutputTypeTcpjsonTcpjson,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            LoadBalanced: criblcontrolplanesdkgo.Pointer(false),
            Compression: components.CompressionOptions1Gzip.ToPointer(),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(true),
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
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](861.22),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](588.22),
            TokenTTLMinutes: criblcontrolplanesdkgo.Pointer[float64](3391.17),
            SendHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("suckle parsnip even engage lest"),
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](10090),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(true),
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "proper-prohibition.com",
                    Port: 7840.9,
                    TLS: components.TLSOptionsHostsItemsInherit.ToPointer(),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    Weight: criblcontrolplanesdkgo.Pointer[float64](3441.24),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](9682.2),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6468.68),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](2841.6),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5558.99),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7385.62),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5308.79),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsTcpjson{},
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
### Example Usage: OutputCreateExamplesWavefront

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesWavefront" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestWavefront(
        operations.OutputWavefront{
            ID: "wavefront-output",
            Type: operations.TypeWavefrontWavefront,
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
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
            Domain: "longboard",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](2592.51),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](5525.15),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](3598.58),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](7660.35),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6545.92),
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
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("airline coal bah kissingly yippee"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](7099.35),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2156.55),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](741.64),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsWavefront{},
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesWebhook" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestWebhook(
        operations.OutputWebhook{
            ID: "webhook-output",
            Type: operations.TypeWebhookWebhook,
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
            Method: components.MethodOptionsPatch.ToPointer(),
            Format: operations.FormatWebhookAdvanced.ToPointer(),
            KeepAlive: criblcontrolplanesdkgo.Pointer(false),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](195.61),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](6141.29),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](2324.34),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](713.5),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6612.7),
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
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: operations.AuthenticationTypeWebhookOauth.ToPointer(),
            TLS: &components.TLSSettingsClientSideType1{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](2331),
            LoadBalanced: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("hovercraft who stoop sniveling nor etch daddy terribly boastfully for"),
            CustomSourceExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomDropWhenNull: criblcontrolplanesdkgo.Pointer(false),
            CustomEventDelimiter: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomContentType: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomPayloadExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            AdvancedContentType: criblcontrolplanesdkgo.Pointer("<value>"),
            FormatEventCode: criblcontrolplanesdkgo.Pointer("<value>"),
            FormatPayloadCode: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](296.71),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2474.1),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](2472.75),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.PqControlsWebhook{},
            Username: criblcontrolplanesdkgo.Pointer("Walter_Luettgen"),
            Password: criblcontrolplanesdkgo.Pointer("LBwcngVN2Osqsr3"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://criminal-pillow.info"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](9310.05),
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
            URL: criblcontrolplanesdkgo.Pointer("https://example.com/webhook"),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Urls: []operations.URLWebhook{
                operations.URLWebhook{
                    URL: "https://yellowish-hunt.org",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](6938.99),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](3110.88),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7669.87),
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

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesWizHec" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestWizHec(
        operations.OutputWizHec{
            ID: "wiz-hec-output",
            Type: operations.TypeWizHecWizHec,
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
            LoadBalanced: true,
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
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](3435.49),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](3058.86),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](2063.99),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](3533.62),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](373.83),
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
            EnableMultiMetrics: false,
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            WizConnectorID: "00000000-0000-0000-0000-000000000000",
            WizEnvironment: "test",
            DataCenter: "us1",
            WizSourcetype: "placeholder",
            Description: criblcontrolplanesdkgo.Pointer("tuber if er kindheartedly gah"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2614.76),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7295.73),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](166.54),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.PqControlsWizHec{},
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesXsiam

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesXsiam" -->
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

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestXsiam(
        operations.OutputXsiam{
            ID: "xsiam-output",
            Type: operations.TypeXsiamXsiam,
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
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](3259.32),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](472.91),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](4490.46),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](2963.35),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](4715.94),
            ExtraHTTPHeaders: []components.ItemsTypeExtraHTTPHeaders{
                components.ItemsTypeExtraHTTPHeaders{
                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                    Value: "<value>",
                },
            },
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            AuthType: operations.AuthenticationMethodXsiamSecret.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5395.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5493.68),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5815.97),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](4757.68),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](483.81),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7907.59),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](37.91),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            ThrottleRateReqPerSec: criblcontrolplanesdkgo.Pointer[int64](438080),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](9986.17),
            Description: criblcontrolplanesdkgo.Pointer("term absentmindedly and fraudster generally midst"),
            URL: criblcontrolplanesdkgo.Pointer("https://fantastic-amnesty.biz"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(true),
            Urls: []operations.URLXsiam{
                operations.URLXsiam{
                    URL: "https://staid-icebreaker.com",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](9779.84),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](3568.49),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3022.38),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6965.85),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](6466.31),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5570.11),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &operations.PqControlsXsiam{},
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.CreateOutputRequest](../../models/operations/createoutputrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

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

### Example Usage: OutputCreateExamplesAzureBlob

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesAzureBlob" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputAzureBlob(
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
            CreateContainer: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](4580.34),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsParquet.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](3396.57),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](3560.43),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](2760.87),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](7242.31),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](8671.99),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](5350.49),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7874.9),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](874.88),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](7218.41),
            },
            AuthType: components.AuthenticationMethodOptionsClientCert.ToPointer(),
            StorageClass: components.BlobAccessTierCold.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("haul that forenenst"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet10.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](584.73),
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
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](4294.14),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](3913.25),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](6602.1),
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
    if res.CountedOutput != nil {
        // handle response
    }
}
```
### Example Usage: OutputCreateExamplesAzureDataExplorer

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesAzureDataExplorer" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputAzureDataExplorer(
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
                "<value 3>",
            },
            ClusterURL: "https://mycluster.kusto.windows.net",
            Database: "mydatabase",
            Table: "mytable",
            ValidateDatabaseSettings: criblcontrolplanesdkgo.Pointer(true),
            IngestMode: components.IngestionModeStreaming.ToPointer(),
            OauthEndpoint: components.MicrosoftEntraIDAuthenticationEndpointOptionsSaslHTTPSLoginMicrosoftonlineCom,
            TenantID: "tenant-id",
            ClientID: "client-id",
            Scope: "https://mycluster.kusto.windows.net/.default",
            OauthType: components.OutputAzureDataExplorerAuthenticationMethodClientSecret,
            Description: criblcontrolplanesdkgo.Pointer("detective nicely quickly loftily newsstand ghost close nearly intently synthesise"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("client-secret"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Certificate: &components.Certificate{
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            Format: components.DataFormatOptionsJSON.ToPointer(),
            Compress: components.CompressionOptions2Gzip,
            CompressionLevel: components.CompressionLevelOptionsNormal.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet24.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV1.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](2241.76),
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
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](2196.24),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](8710.85),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](9826.65),
            IsMappingObj: criblcontrolplanesdkgo.Pointer(false),
            MappingObj: criblcontrolplanesdkgo.Pointer("<value>"),
            MappingRef: criblcontrolplanesdkgo.Pointer("<value>"),
            IngestURL: criblcontrolplanesdkgo.Pointer("https://white-marimba.name"),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            StagePath: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](8620.92),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](6337.34),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](7945.47),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](365.17),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](3276.15),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](5350.49),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7874.9),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](874.88),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](7218.41),
            },
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](9666.97),
            FlushImmediately: criblcontrolplanesdkgo.Pointer(false),
            RetainBlobOnSuccess: criblcontrolplanesdkgo.Pointer(false),
            ExtentTags: []components.ExtentTag{
                components.ExtentTag{
                    Prefix: components.PrefixOptionalDropBy.ToPointer(),
                    Value: "<value>",
                },
            },
            IngestIfNotExists: []components.IngestIfNotExist{
                components.IngestIfNotExist{
                    Value: "<value>",
                },
            },
            ReportLevel: components.ReportLevelDoNotReport.ToPointer(),
            ReportMethod: components.ReportMethodQueueAndTable.ToPointer(),
            AdditionalProperties: []components.AdditionalProperty{
                components.AdditionalProperty{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1570.66),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4314.85),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](1553.95),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](2872.7),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            KeepAlive: criblcontrolplanesdkgo.Pointer(false),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2424.13),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](958.61),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1032.28),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputAzureDataExplorerPqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesAzureEventhub" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputAzureEventhub(
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
                "<value 2>",
                "<value 3>",
            },
            Brokers: []string{
                "myeventhub.servicebus.windows.net:9093",
            },
            Topic: "logs",
            Ack: components.AcknowledgmentsOptionsMinus1.ToPointer(),
            Format: components.RecordDataFormatOptionsJSON.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](4521.66),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](1564.36),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3845.49),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](5506.59),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7472.87),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](8467.46),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](4939.73),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](603.91),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](8979.21),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](6553.71),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](1502.47),
            Sasl: &components.AuthenticationType1{
                Disabled: false,
                AuthType: components.AuthenticationMethodOptionsSasl1Secret.ToPointer(),
                Password: criblcontrolplanesdkgo.Pointer("6ySTtkL9bJUh1Wx"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSasl1Plain.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("Lionel.Jakubowski"),
                ClientSecretAuthType: components.AuthenticationMethodOptionsSasl2Manual.ToPointer(),
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
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("sadly lyre amazing netsuke yeast"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2916.87),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8951.42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5374.29),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputAzureEventhubPqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesAzureLogs" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputAzureLogs(
        components.OutputAzureLogs{
            ID: criblcontrolplanesdkgo.Pointer("azure-logs-output"),
            Type: components.OutputAzureLogsTypeAzureLogs,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            LogType: "Cribl",
            ResourceID: criblcontrolplanesdkgo.Pointer("<id>"),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1624.72),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](7964.22),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](8601.9),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](6578.18),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7805.65),
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
            APIURL: criblcontrolplanesdkgo.Pointer("https://straight-fog.net/"),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            AuthType: components.OutputAzureLogsAuthenticationMethodManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("table brr ick following surface ha and"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8033.02),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1733.16),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6310.27),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputAzureLogsPqControls{},
            WorkspaceID: criblcontrolplanesdkgo.Pointer("workspace-id"),
            WorkspaceKey: criblcontrolplanesdkgo.Pointer("workspace-key"),
            KeypairSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesChronicle" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputChronicle(
        components.OutputChronicle{
            ID: criblcontrolplanesdkgo.Pointer("chronicle-output"),
            Type: components.OutputChronicleTypeChronicle,
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
            APIVersion: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthenticationMethod: components.OutputChronicleAuthenticationMethodServiceAccountSecret.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            Region: "us",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](7071.07),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](3876.89),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](8892.14),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](5148.73),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](9715.08),
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
                "<value 3>",
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](4542.52),
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
                    RbacEnabled: criblcontrolplanesdkgo.Pointer(false),
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("jogging randomize yuppify ack gleefully instead pity"),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            ServiceAccountCredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5340.04),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7887.36),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](3387.4),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputChroniclePqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesClickHouse" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputClickHouse(
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
            AuthType: components.OutputClickHouseAuthenticationTypeOauth.ToPointer(),
            Database: "mydb",
            TableName: "mytable",
            Format: components.OutputClickHouseFormatJSONCompactEachRowWithNames.ToPointer(),
            MappingType: components.MappingTypeAutomatic.ToPointer(),
            AsyncInserts: criblcontrolplanesdkgo.Pointer(true),
            TLS: &components.TLSSettingsClientSideType1{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5985.05),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](3497.49),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](6323.62),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](24.05),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](4113.87),
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
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            DumpFormatErrorsToDisk: criblcontrolplanesdkgo.Pointer(true),
            StatsDestination: &components.StatsDestination{
                URL: criblcontrolplanesdkgo.Pointer("https://mealy-rust.name/"),
                Database: criblcontrolplanesdkgo.Pointer("<value>"),
                TableName: criblcontrolplanesdkgo.Pointer("<value>"),
                AuthType: criblcontrolplanesdkgo.Pointer("<value>"),
                Username: criblcontrolplanesdkgo.Pointer("Ubaldo6"),
                SQLUsername: criblcontrolplanesdkgo.Pointer("<value>"),
                Password: criblcontrolplanesdkgo.Pointer("bM1DJajpgFoD_5r"),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("recede how for than internalize ick ouch"),
            Username: criblcontrolplanesdkgo.Pointer("Ethelyn98"),
            Password: criblcontrolplanesdkgo.Pointer("33GliS58gYPjl7D"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://majestic-gastropod.biz"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](8593.96),
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
            SQLUsername: criblcontrolplanesdkgo.Pointer("<value>"),
            WaitForAsyncInserts: criblcontrolplanesdkgo.Pointer(false),
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
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6186.07),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8015.89),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](715.63),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputClickHousePqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesCloudflareR2" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputCloudflareR2(
        components.OutputCloudflareR2{
            ID: criblcontrolplanesdkgo.Pointer("cloudflare-r2-output"),
            Type: components.OutputCloudflareR2TypeCloudflareR2,
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
            Endpoint: "https://account-id.r2.cloudflarestorage.com",
            Bucket: "my-bucket",
            AwsAuthenticationMethod: components.OutputCloudflareR2AuthenticationMethodManual.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "<value>",
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions5V4.ToPointer(),
            ObjectACL: "<value>",
            StorageClass: components.StorageClassOptions2Standard.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionOptionsAes256.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(false),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](6944.96),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](1443.56),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](8928.1),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](5350.49),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7874.9),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](874.88),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](7218.41),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](9960.23),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](2986.51),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](1946.7),
            Description: criblcontrolplanesdkgo.Pointer("excluding knowingly energetically incidentally given phooey supposing within"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2None.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestCompression.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet10.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](2631.54),
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
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](596.4),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](9978.79),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](4725.91),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesCloudwatch" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputCloudwatch(
        components.OutputCloudwatch{
            ID: criblcontrolplanesdkgo.Pointer("cloudwatch-output"),
            Type: components.OutputCloudwatchTypeCloudwatch,
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
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfSecret.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3586.65),
            MaxQueueSize: criblcontrolplanesdkgo.Pointer[float64](8020.6),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](4014.12),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5779.8),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("cautiously seriously till investigate alongside obvious shakily afore"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1614.27),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](6995.82),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1566.35),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputCloudwatchPqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesConfluentCloud" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputConfluentCloud(
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
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            Topic: "logs",
            Ack: components.AcknowledgmentsOptions1Zero.ToPointer(),
            Format: components.RecordDataFormatOptions1JSON.ToPointer(),
            Compression: components.CompressionOptions3Snappy.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](9923.99),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](6751.13),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](693.69),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType1{
                Disabled: true,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://pitiful-fit.org/"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](7989.51),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](6279.74),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](5484.89),
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
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                },
                DefaultKeySchemaID: criblcontrolplanesdkgo.Pointer[float64](4850.53),
                DefaultValueSchemaID: criblcontrolplanesdkgo.Pointer[float64](6240.22),
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](2850.7),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](3913.49),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](6346.75),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](4195.89),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](5825.56),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](5390.24),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](7366.3),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](1755.73),
            Sasl: &components.AuthenticationType{
                Disabled: false,
                Username: criblcontrolplanesdkgo.Pointer("Monte13"),
                Password: criblcontrolplanesdkgo.Pointer("z4xrW3jxrPrE_UR"),
                AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslKerberos.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(true),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://low-handful.biz"),
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
            Description: criblcontrolplanesdkgo.Pointer("cruelty well-documented poorly"),
            ProtobufLibraryID: criblcontrolplanesdkgo.Pointer("<id>"),
            ProtobufEncodingID: criblcontrolplanesdkgo.Pointer("<id>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8457.47),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2709.96),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8341.2),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputConfluentCloudPqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesCriblHttp" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputCriblHTTP(
        components.OutputCriblHTTP{
            ID: criblcontrolplanesdkgo.Pointer("cribl-http-output"),
            Type: components.OutputCriblHTTPTypeCriblHTTP,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            LoadBalanced: criblcontrolplanesdkgo.Pointer(true),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            TokenTTLMinutes: criblcontrolplanesdkgo.Pointer[float64](6551.76),
            ExcludeFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Compression: components.CompressionOptions1None.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](8662.82),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](5200.55),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](1830.87),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](2363.66),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](743.17),
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
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            AuthTokens: []components.ItemsTypeAuthTokens1{
                components.ItemsTypeAuthTokens1{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                    Description: criblcontrolplanesdkgo.Pointer("repeatedly since ugh yippee floodlight ugh trolley"),
                },
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("watery whoa amendment misguided apropos quit outset between aside"),
            URL: criblcontrolplanesdkgo.Pointer("https://favorite-horst.net/"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(true),
            Urls: []components.ItemsTypeUrls{
                components.ItemsTypeUrls{
                    URL: "https://hurtful-secret.info/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](7112.21),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](1383.86),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5502.37),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2826.81),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1109),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](2481.88),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputCriblHTTPPqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesCriblLake" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputCriblLake(
        components.OutputCriblLake{
            ID: criblcontrolplanesdkgo.Pointer("cribl-lake-output"),
            Type: components.OutputCriblLakeTypeCriblLake,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
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
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](433.66),
            StagePath: criblcontrolplanesdkgo.Pointer("<value>"),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            ObjectACL: components.ObjectACLOptionsAwsExecRead.ToPointer(),
            StorageClass: components.StorageClassOptionsGlacierIr.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAes256.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](1372.15),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](8899.84),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](35.72),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](5350.49),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7874.9),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](874.88),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](7218.41),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](8119.4),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](8688.02),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](9303.09),
            AwsAuthenticationMethod: components.MethodOptionsCredentialsAuto.ToPointer(),
            Format: components.FormatOptionsCriblLakeDatasetParquet.ToPointer(),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](8774.65),
            Description: criblcontrolplanesdkgo.Pointer("swat until by phew whenever er duh"),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](2310.38),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](4260.18),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](9036.05),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesCriblSearchEngine" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputCriblSearchEngine(
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
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
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
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](981.96),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7939.43),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](9594.37),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            AuthTokens: []components.ItemsTypeAuthTokens1{
                components.ItemsTypeAuthTokens1{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                    Description: criblcontrolplanesdkgo.Pointer("repeatedly since ugh yippee floodlight ugh trolley"),
                },
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("huzzah by purple service another gah"),
            URL: criblcontrolplanesdkgo.Pointer("https://0.0.0.0:10200"),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(true),
            Urls: []components.ItemsTypeUrls{
                components.ItemsTypeUrls{
                    URL: "https://hurtful-secret.info/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](7112.21),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](4721.64),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3084.35),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8885.13),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1154.78),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](879.62),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputCriblSearchEnginePqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesCriblTcp" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputCriblTCP(
        components.OutputCriblTCP{
            ID: criblcontrolplanesdkgo.Pointer("cribl-tcp-output"),
            Type: components.OutputCriblTCPTypeCriblTCP,
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
            LoadBalanced: criblcontrolplanesdkgo.Pointer(false),
            Compression: components.CompressionOptions1None.ToPointer(),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(false),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](9436.97),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](8454.09),
            TokenTTLMinutes: criblcontrolplanesdkgo.Pointer[float64](95.16),
            AuthTokens: []components.ItemsTypeAuthTokens{
                components.ItemsTypeAuthTokens{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(false),
                    Description: criblcontrolplanesdkgo.Pointer("save outrageous mortally than uselessly minion jittery insist when"),
                },
            },
            ExcludeFields: []string{
                "<value 1>",
            },
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("phew beyond eternity finally saloon finally"),
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](10090),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "wicked-bend.com",
                    Port: 9502.76,
                    TLS: components.TLSOptionsHostsItemsOff.ToPointer(),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1086.4),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](3598.76),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](9282.89),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](4950.33),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2810.66),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2806.73),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5157.3),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputCriblTCPPqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesCrowdstrikeNextGenSiem" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputCrowdstrikeNextGenSiem(
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
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](6318.08),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](3488.18),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](3792.45),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](6452.72),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7236.96),
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
            },
            Format: components.RequestFormatOptionsJSON,
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("rectangular opposite ew"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9833.62),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7536.74),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8262.11),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputCrowdstrikeNextGenSiemPqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesDatabricks" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputDatabricks(
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
                "<value 2>",
            },
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            StagePath: criblcontrolplanesdkgo.Pointer("<value>"),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsRaw.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](6413.72),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](9372.08),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](2511.55),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](8005.27),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](7971.94),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](5350.49),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7874.9),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](874.88),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](7218.41),
            },
            WorkspaceID: "your-workspace-id",
            Scope: "my-scope",
            ClientID: "your-client-id",
            Catalog: "my-catalog",
            Schema: "my-schema",
            EventsVolumeName: "my-volume",
            ClientTextSecret: "your-client-secret",
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](5952.58),
            Description: criblcontrolplanesdkgo.Pointer("however taut never distorted aboard"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestCompression.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV1.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](4355.49),
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
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](7943.4),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](6362.56),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](1324.5),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesDatadog" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputDatadog(
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
                "<value 2>",
            },
            ContentType: components.SendLogsAsJSON.ToPointer(),
            Message: criblcontrolplanesdkgo.Pointer("<value>"),
            Source: criblcontrolplanesdkgo.Pointer("<value>"),
            Host: criblcontrolplanesdkgo.Pointer("agile-compromise.name"),
            Service: criblcontrolplanesdkgo.Pointer("<value>"),
            Tags: []string{
                "<value 1>",
            },
            BatchByTags: criblcontrolplanesdkgo.Pointer(true),
            AllowAPIKeyFromEvents: criblcontrolplanesdkgo.Pointer(false),
            Severity: components.OutputDatadogSeverityError.ToPointer(),
            Site: components.DatadogSiteCustom.ToPointer(),
            SendCountersAsCount: criblcontrolplanesdkgo.Pointer(false),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](62.83),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1842.84),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](9768.16),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](2975.55),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5875.99),
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
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Secret.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](9611.42),
            Description: criblcontrolplanesdkgo.Pointer("but yellow whereas yowza sweetly scent lobster tectonics consequently golden"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://animated-substitution.name"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1088.03),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](6744.35),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6094.02),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesDataset" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputDataset(
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
                "<value 2>",
            },
            MessageField: criblcontrolplanesdkgo.Pointer("<value>"),
            ExcludeFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ServerHostField: criblcontrolplanesdkgo.Pointer("<value>"),
            TimestampField: criblcontrolplanesdkgo.Pointer("<value>"),
            DefaultSeverity: components.OutputDatasetSeverityFinest.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            Site: components.DataSetSiteUs.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1407.28),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](9662.77),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](5952.13),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](7255.09),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7666.54),
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
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Manual.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](3254.58),
            Description: criblcontrolplanesdkgo.Pointer("submitter plus fatal than furthermore each meh"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://somber-tackle.org/"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](4335.03),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9965.77),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](7535.13),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputDatasetPqControls{},
            APIKey: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesDiskSpool

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesDiskSpool" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputDiskSpool(
        components.OutputDiskSpool{
            ID: criblcontrolplanesdkgo.Pointer("disk-spool-output"),
            Type: components.OutputDiskSpoolTypeDiskSpool,
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
            TimeWindow: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptionsPersistenceGzip.ToPointer(),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("basket yuck enlist where restructure toaster"),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesDlS3" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputDlS3(
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
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfSecret.ToPointer(),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](5551.88),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            ObjectACL: components.ObjectACLOptionsPublicRead.ToPointer(),
            StorageClass: components.StorageClassOptionsStandardIa.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAes256.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            Format: components.DataFormatOptionsRaw.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](3256.54),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](8910.59),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](3829.54),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(true),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](5350.49),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7874.9),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](874.88),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](7218.41),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](3586.58),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](4251.32),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](436.67),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(false),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](1616.18),
            PartitioningFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Description: criblcontrolplanesdkgo.Pointer("personal cavernous decouple excluding"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2None.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestCompression.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet24.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](6177.61),
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
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(true),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](1893.64),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](8819.97),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](9772.39),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesDynatraceHttp" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputDynatraceHTTP(
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
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](7037.41),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](9651.26),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](9258.69),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](6716.85),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](145.13),
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
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.OutputDynatraceHTTPAuthenticationTypeToken.ToPointer(),
            Format: components.OutputDynatraceHTTPFormatJSONArray,
            Endpoint: components.EndpointCloud,
            TelemetryType: components.TelemetryTypeLogs,
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](4591.06),
            Description: criblcontrolplanesdkgo.Pointer("um ouch colligate or aha abaft mmm"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](7546.05),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4876.98),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](289.92),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputDynatraceHTTPPqControls{},
            Token: criblcontrolplanesdkgo.Pointer("your-api-key"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            EnvironmentID: criblcontrolplanesdkgo.Pointer("<id>"),
            ActiveGateDomain: criblcontrolplanesdkgo.Pointer("<value>"),
            URL: criblcontrolplanesdkgo.Pointer("https://candid-knuckle.info"),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesDynatraceOtlp" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputDynatraceOtlp(
        components.OutputDynatraceOtlp{
            ID: criblcontrolplanesdkgo.Pointer("dynatrace-otlp-output"),
            Type: components.OutputDynatraceOtlpTypeDynatraceOtlp,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Protocol: components.OutputDynatraceOtlpProtocolHTTP,
            Endpoint: "https://your-environment.live.dynatrace.com/api/v2/otlp",
            OtlpVersion: components.OtlpVersionOptions1OneDot3Dot1,
            Compress: components.CompressionOptions4Deflate.ToPointer(),
            HTTPCompress: components.CompressionOptions5Gzip.ToPointer(),
            HTTPTracesEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPMetricsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPLogsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](2993.67),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1500.18),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](5036.53),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](609.47),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](1046.56),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](667.85),
            KeepAlive: criblcontrolplanesdkgo.Pointer(false),
            EndpointType: components.EndpointTypeSaas,
            TokenSecret: "your-token-secret",
            AuthTokenName: criblcontrolplanesdkgo.Pointer("<value>"),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("compete out beyond"),
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
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2172.82),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](506.93),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1743.29),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesElastic" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputElastic(
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
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](9127.73),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](3933.39),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](2466.52),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](5952.7),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5671.89),
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
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            ExtraParams: []components.ItemsTypeSaslSaslExtensions{
                components.ItemsTypeSaslSaslExtensions{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Auth: &components.AuthType{
                Disabled: false,
                Username: criblcontrolplanesdkgo.Pointer("Claudine69"),
                Password: criblcontrolplanesdkgo.Pointer("J3ts4hhLd7yGMlA"),
                AuthType: components.AuthenticationMethodOptionsAuthSecret.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                ManualAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            ElasticVersion: components.ElasticVersionAuto.ToPointer(),
            ElasticPipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            IncludeDocID: criblcontrolplanesdkgo.Pointer(false),
            WriteAction: components.WriteActionIndex.ToPointer(),
            RetryPartialErrors: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("however inside knowledgeable severe eyeglasses fit ha"),
            URL: criblcontrolplanesdkgo.Pointer("https://soft-tuba.net"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(true),
            Urls: []components.OutputElasticURL{
                components.OutputElasticURL{
                    URL: "https://profitable-completion.biz/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](261.53),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](6235.99),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](4923.9),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2187.35),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1791.54),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8996.67),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputElasticPqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesElasticCloud" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputElasticCloud(
        components.OutputElasticCloud{
            ID: criblcontrolplanesdkgo.Pointer("elastic-cloud-output"),
            Type: components.OutputElasticCloudTypeElasticCloud,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            URL: "my-cloud-id",
            Index: "logs",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](8334.63),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4489.5),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](988.45),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](8382.89),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5937.18),
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
            ExtraParams: []components.ItemsTypeSaslSaslExtensions{
                components.ItemsTypeSaslSaslExtensions{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Auth: &components.AuthType{
                Disabled: false,
                Username: criblcontrolplanesdkgo.Pointer("Claudine69"),
                Password: criblcontrolplanesdkgo.Pointer("J3ts4hhLd7yGMlA"),
                AuthType: components.AuthenticationMethodOptionsAuthSecret.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                ManualAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            ElasticPipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            IncludeDocID: criblcontrolplanesdkgo.Pointer(true),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("pish pessimistic after um disgorge within"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](122.24),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8807.85),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8400.29),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesExabeam" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputExabeam(
        components.OutputExabeam{
            ID: criblcontrolplanesdkgo.Pointer("exabeam-output"),
            Type: components.OutputExabeamTypeExabeam,
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
            Bucket: "my-bucket",
            Region: "us-east1",
            StagePath: "/tmp/staging",
            Endpoint: "https://storage.googleapis.com",
            SignatureVersion: components.SignatureVersionOptions4V2.ToPointer(),
            ObjectACL: components.ObjectACLOptions1PublicRead.ToPointer(),
            StorageClass: components.StorageClassOptions1Coldline.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](6749.81),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](6990.43),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](3101.21),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](5350.49),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7874.9),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](874.88),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](7218.41),
            },
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](8701.26),
            EncodedConfiguration: criblcontrolplanesdkgo.Pointer("<value>"),
            CollectorInstanceID: "11112222-3333-4444-5555-666677778888",
            SiteName: criblcontrolplanesdkgo.Pointer("<value>"),
            SiteID: criblcontrolplanesdkgo.Pointer("<id>"),
            TimezoneOffset: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("derby once yahoo braid clavicle cross-contamination crushing clavicle swing"),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](8440.38),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](8815.99),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](4597.15),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesFilesystem" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputFilesystem(
        components.OutputFilesystem{
            ID: criblcontrolplanesdkgo.Pointer("filesystem-output"),
            Type: components.OutputFilesystemTypeFilesystem,
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
            DestPath: "/var/log/output",
            StagePath: criblcontrolplanesdkgo.Pointer("<value>"),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsParquet.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](4579.59),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](1591.93),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](1645.42),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](2710.22),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](8023.17),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](5350.49),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7874.9),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](874.88),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](7218.41),
            },
            Description: criblcontrolplanesdkgo.Pointer("outside dish worth rapidly joyfully requite until"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestCompression.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet10.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](8841.03),
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
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(true),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](4416.3),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](450.74),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](1172.59),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesGoogleChronicle" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputGoogleChronicle(
        components.OutputGoogleChronicle{
            ID: criblcontrolplanesdkgo.Pointer("google-chronicle-output"),
            Type: components.OutputGoogleChronicleTypeGoogleChronicle,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            APIVersion: components.OutputGoogleChronicleAPIVersionV1.ToPointer(),
            AuthenticationMethod: components.OutputGoogleChronicleAuthenticationMethodServiceAccountSecret.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            LogFormatType: components.SendEventsAsUnstructured,
            Region: criblcontrolplanesdkgo.Pointer("us"),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](194.88),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](7780.64),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](2256.5),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](4616.39),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](9309.58),
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
                "<value 3>",
            },
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](6941.09),
            Description: criblcontrolplanesdkgo.Pointer("insecure equatorial release until scaly around inasmuch eek"),
            ExtraLogTypes: []components.ExtraLogType{
                components.ExtraLogType{
                    LogType: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("solder dull stay yet overcooked per debut attest"),
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
            UdmType: components.UDMTypeEntities.ToPointer(),
            APIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            APIKeySecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            ServiceAccountCredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2043.78),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8903.83),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8508.98),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputGoogleChroniclePqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesGoogleCloudLogging" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputGoogleCloudLogging(
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
            GoogleAuthMethod: components.GoogleAuthenticationMethodOptionsAuto.ToPointer(),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](2256.87),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](2919.53),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6889.61),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](494.84),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](2659.94),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](8967.63),
            ThrottleRateReqPerSec: criblcontrolplanesdkgo.Pointer[int64](862713),
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
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](7513.2),
            Description: criblcontrolplanesdkgo.Pointer("save excluding ornery ameliorate"),
            LogLocationExpression: "my-project",
            PayloadExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6862.95),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4789.34),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](9408.21),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesGoogleCloudStorage" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputGoogleCloudStorage(
        components.OutputGoogleCloudStorage{
            ID: criblcontrolplanesdkgo.Pointer("google-cloud-storage-output"),
            Type: components.OutputGoogleCloudStorageTypeGoogleCloudStorage,
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
            Endpoint: "https://storage.googleapis.com",
            SignatureVersion: components.SignatureVersionOptions4V4.ToPointer(),
            AwsAuthenticationMethod: components.OutputGoogleCloudStorageAuthenticationMethodSecret.ToPointer(),
            StagePath: "/tmp/staging",
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(false),
            ObjectACL: components.ObjectACLOptions1Private.ToPointer(),
            StorageClass: components.StorageClassOptions1Archive.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsParquet.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](5999.09),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](5928.82),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](8950.65),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](7438.31),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](2698.13),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(true),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](5350.49),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7874.9),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](874.88),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](7218.41),
            },
            Description: criblcontrolplanesdkgo.Pointer("miserably vicinity blindly moral bossy"),
            Compress: components.CompressionOptions2None.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestCompression.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet10.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV1.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](467.44),
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
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](7003.57),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](3127.55),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](7631.55),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesGooglePubsub" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputGooglePubsub(
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
            },
            TopicName: "my-topic",
            CreateTopic: criblcontrolplanesdkgo.Pointer(true),
            OrderedDelivery: criblcontrolplanesdkgo.Pointer(true),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            GoogleAuthMethod: components.GoogleAuthenticationMethodOptionsManual.ToPointer(),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            BatchSize: criblcontrolplanesdkgo.Pointer[float64](9464.52),
            BatchTimeout: criblcontrolplanesdkgo.Pointer[float64](5743.79),
            MaxQueueSize: criblcontrolplanesdkgo.Pointer[float64](8657.4),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](7839.81),
            FlushPeriod: criblcontrolplanesdkgo.Pointer[float64](7804.51),
            MaxInProgress: criblcontrolplanesdkgo.Pointer[float64](7060.4),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("bowed jiggle acquire down bootleg"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1114.52),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3535.93),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1129.99),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputGooglePubsubPqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesGrafanaCloud" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputGrafanaCloud(
        components.CreateOutputGrafanaCloudOutputGrafanaCloudGrafanaCloud1(
            components.OutputGrafanaCloudGrafanaCloud1{
                ID: criblcontrolplanesdkgo.Pointer("grafana-cloud-output"),
                Type: components.OutputGrafanaCloudType1GrafanaCloud,
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
                LokiURL: "https://logs-prod-us-central1.grafana.net",
                PrometheusURL: criblcontrolplanesdkgo.Pointer("https://deep-bidet.info"),
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
                    AuthType: components.AuthenticationTypeOptionsPrometheusAuth1TextSecret.ToPointer(),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Username: criblcontrolplanesdkgo.Pointer("Telly34"),
                    Password: criblcontrolplanesdkgo.Pointer("xj5Pr1ID0oewBJ5"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                LokiAuth: &components.PrometheusAuthType{
                    AuthType: components.AuthenticationTypeOptionsPrometheusAuth1Token.ToPointer(),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Username: criblcontrolplanesdkgo.Pointer("Eduardo10"),
                    Password: criblcontrolplanesdkgo.Pointer("6qD_OjiWX1YoRh0"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                Concurrency: criblcontrolplanesdkgo.Pointer[float64](6675.42),
                MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](7351.76),
                MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](1764.31),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](3221.15),
                FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3452.4),
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
                        HTTPStatus: 3612.5,
                        InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                        BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                        MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                    },
                },
                TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                    TimeoutRetry: false,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
                },
                ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
                OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
                Description: criblcontrolplanesdkgo.Pointer("countess furthermore excess stable neat veto ick behind requite"),
                Compress: criblcontrolplanesdkgo.Pointer(false),
                PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
                PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](4717.14),
                PqMode: components.ModeOptionsAlways.ToPointer(),
                PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7755.43),
                PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6556.99),
                PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
                PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
                PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PqCompress: components.CompressionOptionsPqNone.ToPointer(),
                PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
                PqControls: &components.OutputGrafanaCloudPqControls1{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesGraphite" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputGraphite(
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
            Mtu: criblcontrolplanesdkgo.Pointer[float64](60.21),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7162.25),
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](3967.84),
            Description: criblcontrolplanesdkgo.Pointer("silently save hope acquire reluctantly inside own outside geez segregate"),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](815.93),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](4625.77),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](253.31),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5622.83),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5811.19),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesHoneycomb" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputHoneycomb(
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
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](9827.16),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4188.74),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](3305.31),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](8127.64),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1093.1),
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
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("crooked drat readmit"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](642.3),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](732.36),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6647.84),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesHumioHec" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputHumioHec(
        components.OutputHumioHec{
            ID: criblcontrolplanesdkgo.Pointer("humio-hec-output"),
            Type: components.OutputHumioHecTypeHumioHec,
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
            URL: "https://cloud.us.humio.com/api/v1/ingest/hec",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](6046.59),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](3582.01),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](9470.26),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](6546.11),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](9923.35),
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
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("writhing sizzle mechanically actually of etch quick"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6722.38),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](6024.32),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6541.46),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputHumioHecPqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesInfluxdb" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputInfluxdb(
        components.OutputInfluxdb{
            ID: criblcontrolplanesdkgo.Pointer("influxdb-output"),
            Type: components.OutputInfluxdbTypeInfluxdb,
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
            URL: "http://localhost:8086",
            UseV2API: criblcontrolplanesdkgo.Pointer(false),
            TimestampPrecision: components.TimestampPrecisionU.ToPointer(),
            DynamicValueFieldName: criblcontrolplanesdkgo.Pointer(true),
            ValueFieldName: criblcontrolplanesdkgo.Pointer("<value>"),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5996.06),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4850.95),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](7404.99),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](9587.3),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3902.6),
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
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.OutputInfluxdbAuthenticationTypeNone.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("ugh trash solicit key gymnast bracelet ew sarong cumbersome um"),
            Database: criblcontrolplanesdkgo.Pointer("mydb"),
            Bucket: criblcontrolplanesdkgo.Pointer("<value>"),
            Org: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5201.94),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1717.22),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8317.89),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputInfluxdbPqControls{},
            Username: criblcontrolplanesdkgo.Pointer("Kelvin_Schinner50"),
            Password: criblcontrolplanesdkgo.Pointer("xOV1CyorOVG2SAG"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://muffled-venom.name/"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](4728.73),
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
    if res.CountedOutput != nil {
        // handle response
    }
}
```
### Example Usage: OutputCreateExamplesKafka

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesKafka" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputKafka(
        components.OutputKafka{
            ID: criblcontrolplanesdkgo.Pointer("kafka-output"),
            Type: components.OutputKafkaTypeKafka,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
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
            Ack: components.AcknowledgmentsOptions1Zero.ToPointer(),
            Format: components.RecordDataFormatOptions1Raw.ToPointer(),
            Compression: components.CompressionOptions3Zstd.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](9335.5),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](6847.29),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5265.71),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType1{
                Disabled: true,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://pitiful-fit.org/"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](7989.51),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](6279.74),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](5484.89),
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
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                },
                DefaultKeySchemaID: criblcontrolplanesdkgo.Pointer[float64](4850.53),
                DefaultValueSchemaID: criblcontrolplanesdkgo.Pointer[float64](6240.22),
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](632.63),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7810.04),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](3762.8),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](7826.94),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1688.06),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](9488.3),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](6683.04),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](1380.2),
            Sasl: &components.AuthenticationType{
                Disabled: false,
                Username: criblcontrolplanesdkgo.Pointer("Monte13"),
                Password: criblcontrolplanesdkgo.Pointer("z4xrW3jxrPrE_UR"),
                AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslKerberos.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(true),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://low-handful.biz"),
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
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("indeed simple angle daddy before"),
            ProtobufLibraryID: criblcontrolplanesdkgo.Pointer("<id>"),
            ProtobufEncodingID: criblcontrolplanesdkgo.Pointer("<id>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1608.03),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7762.89),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](5546.58),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputKafkaPqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesKinesis" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputKinesis(
        components.OutputKinesis{
            ID: criblcontrolplanesdkgo.Pointer("kinesis-output"),
            Type: components.OutputKinesisTypeKinesis,
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
            StreamName: "my-stream",
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfManual.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions2V2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](2327.88),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](180.13),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](2560.27),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1305.95),
            Compression: components.OutputKinesisCompressionNone.ToPointer(),
            UseListShards: criblcontrolplanesdkgo.Pointer(true),
            AsNdjson: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("including during sundae"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxEventsPerFlush: criblcontrolplanesdkgo.Pointer[float64](3193.11),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](1415.77),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](9189.02),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8824.28),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputKinesisPqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesLoki" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputLoki(
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
                "<value 2>",
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
            AuthType: components.AuthenticationTypeOptionsPrometheusAuth1TextSecret.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](7676.41),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](8598.26),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](9409.24),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](6028.87),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6526.76),
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
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            EnableDynamicHeaders: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](7686.34),
            Description: criblcontrolplanesdkgo.Pointer("premier healthily freight ha till furthermore excepting diver"),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Username: criblcontrolplanesdkgo.Pointer("Matilde65"),
            Password: criblcontrolplanesdkgo.Pointer("xqJcOYhJD96fk1o"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](794.19),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7756.87),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](7578.66),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesMicrosoftFabric" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputMicrosoftFabric(
        components.OutputMicrosoftFabric{
            ID: criblcontrolplanesdkgo.Pointer("microsoft-fabric-output"),
            Type: components.OutputMicrosoftFabricTypeMicrosoftFabric,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Topic: "logs",
            Ack: components.AcknowledgmentsOptionsOne.ToPointer(),
            Format: components.RecordDataFormatOptionsJSON.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](3515.83),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](5133.45),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](9242.46),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](4763.74),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](7816.61),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](6437.59),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](3343.13),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1754.74),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1772.66),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](8068.63),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](9336.44),
            Sasl: &components.OutputMicrosoftFabricAuthentication{
                Disabled: false,
                Mechanism: components.SaslMechanismOptionsSasl1Plain.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("Tristin_White"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                ClientSecretAuthType: components.OutputMicrosoftFabricAuthenticationMethodSecret.ToPointer(),
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
                Disabled: false,
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            BootstrapServer: "myeventstream.servicebus.windows.net:9093",
            Description: criblcontrolplanesdkgo.Pointer("oof doing quickly bloom scary oof worth"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5910.37),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7996.45),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](721.72),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputMicrosoftFabricPqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesMinio" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputMinio(
        components.OutputMinio{
            ID: criblcontrolplanesdkgo.Pointer("minio-output"),
            Type: components.OutputMinioTypeMinio,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Endpoint: "http://localhost:9000",
            Bucket: "my-bucket",
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfManual.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions5V4.ToPointer(),
            ObjectACL: components.ObjectACLOptionsBucketOwnerFullControl.ToPointer(),
            StorageClass: components.StorageClassOptions2Standard.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionOptionsAes256.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsParquet.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](9148.78),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](8176.95),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](5965.49),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsDrop.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(true),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](5350.49),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7874.9),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](874.88),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](7218.41),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](8733.54),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](2457.48),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](1224.76),
            Description: criblcontrolplanesdkgo.Pointer("late beep freezing about broadside defenseless ajar hmph past yuck"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2None.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet24.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV1.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](5866.7),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("<value>"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(false),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(false),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(true),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](734.84),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](8236.09),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](9685.91),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesMsk" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputMsk(
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
                "<value 2>",
                "<value 3>",
            },
            Brokers: []string{
                "b-1.example.xxxxx.c2.kafka.us-east-1.amazonaws.com:9092",
            },
            Topic: "logs",
            Ack: components.AcknowledgmentsOptions1Minus1.ToPointer(),
            Format: components.RecordDataFormatOptions1Protobuf.ToPointer(),
            Compression: components.CompressionOptions3Snappy.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](1538.83),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](8245.04),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6192.18),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType1{
                Disabled: true,
                SchemaRegistryURL: criblcontrolplanesdkgo.Pointer("https://pitiful-fit.org/"),
                ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](7989.51),
                RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](6279.74),
                MaxRetries: criblcontrolplanesdkgo.Pointer[float64](5484.89),
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
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                },
                DefaultKeySchemaID: criblcontrolplanesdkgo.Pointer[float64](4850.53),
                DefaultValueSchemaID: criblcontrolplanesdkgo.Pointer[float64](6240.22),
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](768.28),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](4188.17),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](3722.18),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](864.6),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1640.63),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2251),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](3509.93),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](1897.23),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto,
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsV2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](936.75),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("thankfully lest resource"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ProtobufLibraryID: criblcontrolplanesdkgo.Pointer("<id>"),
            ProtobufEncodingID: criblcontrolplanesdkgo.Pointer("<id>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](3645.11),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8425.35),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](3700.97),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputMskPqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesNetflow" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputNetflow(
        components.OutputNetflow{
            ID: criblcontrolplanesdkgo.Pointer("netflow-output"),
            Type: components.OutputNetflowTypeNetflow,
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
            Hosts: []components.OutputNetflowHost{
                components.OutputNetflowHost{
                    Host: "localhost",
                    Port: 2055,
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](2279.48),
            EnableIPSpoofing: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("hefty yawningly short successfully given tray cod"),
            MaxRecordSize: criblcontrolplanesdkgo.Pointer[float64](7657.6),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesNewrelic" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputNewrelic(
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
                "<value 2>",
            },
            Region: components.RegionOptionsEu.ToPointer(),
            LogType: criblcontrolplanesdkgo.Pointer("<value>"),
            MessageField: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.Metadatum{
                components.Metadatum{
                    Name: components.FieldNameService,
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](7353.66),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](7927.67),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](2106.86),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](7201.59),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](694.13),
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
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Manual.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](5205.55),
            Description: criblcontrolplanesdkgo.Pointer("bleach cautiously psst whenever"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://serpentine-draw.net"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6481.07),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4366.65),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1664.77),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputNewrelicPqControls{},
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
### Example Usage: OutputCreateExamplesNewrelicEvents

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesNewrelicEvents" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputNewrelicEvents(
        components.OutputNewrelicEvents{
            ID: criblcontrolplanesdkgo.Pointer("newrelic-events-output"),
            Type: components.OutputNewrelicEventsTypeNewrelicEvents,
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
            Region: components.RegionOptionsUs.ToPointer(),
            AccountID: "123456",
            EventType: "CriblEvent",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1532.27),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](9645.91),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](4360.95),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](3897.41),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](4034.67),
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
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("ill-fated down tennis crumble"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://recent-technician.org"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2503.54),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8623.14),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](4349.61),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputNewrelicEventsPqControls{},
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
### Example Usage: OutputCreateExamplesOpenTelemetry

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesOpenTelemetry" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputOpenTelemetry(
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
            },
            Protocol: components.ProtocolOptionsHTTP.ToPointer(),
            Endpoint: "http://localhost:4317",
            OtlpVersion: components.OutputOpenTelemetryOTLPVersionZeroDot10Dot0.ToPointer(),
            Compress: components.CompressionOptions4None.ToPointer(),
            HTTPCompress: components.CompressionOptions5Gzip.ToPointer(),
            AuthType: components.AuthenticationTypeOptionsOauth.ToPointer(),
            HTTPTracesEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPMetricsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPLogsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](7303.11),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](8230.1),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](5107.11),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8697.13),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](5372.86),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](2687.71),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("fray jovially brr despite even"),
            Username: criblcontrolplanesdkgo.Pointer("Burnice.Yost26"),
            Password: criblcontrolplanesdkgo.Pointer("xHJbfZYfEAhnjju"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://tame-lender.net/"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](2153.54),
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
                "<value 3>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
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
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9312.71),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](1752.16),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](3780.65),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesPrometheus" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputPrometheus(
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
            },
            URL: "http://localhost:9091/api/v1/write",
            MetricRenameExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            SendMetadata: criblcontrolplanesdkgo.Pointer(false),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](3559.44),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1232.43),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](8644.71),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](3599.31),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](8634.13),
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
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            AuthType: components.AuthenticationTypeOptionsPrometheusAuthToken.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("efface putrefy cannon whose how hard-to-find wry imaginary down"),
            MetricsFlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5953.83),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2175.69),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7522.56),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](4247.3),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputPrometheusPqControls{},
            Username: criblcontrolplanesdkgo.Pointer("Joy.Herzog"),
            Password: criblcontrolplanesdkgo.Pointer("a2t5_lDRX3upjqM"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://different-festival.biz/"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](1717.51),
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
    if res.CountedOutput != nil {
        // handle response
    }
}
```
### Example Usage: OutputCreateExamplesRing

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesRing" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputRing(
        components.OutputRing{
            ID: criblcontrolplanesdkgo.Pointer("ring-output"),
            Type: components.OutputRingTypeRing,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Format: components.OutputRingDataFormatRaw.ToPointer(),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxDataSize: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxDataTime: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("phooey although yowza instead irritably stark gleefully"),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesRouter" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputRouter(
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
            },
            Rules: []components.OutputRouterRule{
                components.OutputRouterRule{
                    Filter: "true",
                    Output: "my-output",
                    Description: criblcontrolplanesdkgo.Pointer("pigpen procurement aha minus cautious density phony alliance"),
                    Final: criblcontrolplanesdkgo.Pointer(true),
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("but midst cuddly coal eek"),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesS3" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputS3(
        components.OutputS3{
            ID: criblcontrolplanesdkgo.Pointer("s3-output"),
            Type: components.OutputS3TypeS3,
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
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfManual.ToPointer(),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](6731.29),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            ObjectACL: components.ObjectACLOptionsBucketOwnerRead.ToPointer(),
            StorageClass: components.StorageClassOptionsGlacier.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAes256.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(false),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.DataFormatOptionsParquet.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](9585.8),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](2297.86),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](9300.56),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(true),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](5350.49),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7874.9),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](874.88),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](7218.41),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](8428.01),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](8114.35),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](3254.02),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(false),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](2839.27),
            Description: criblcontrolplanesdkgo.Pointer("failing until once sure-footed gladly loudly"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsNormal.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet24.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](8889.64),
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
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](5146.49),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](7356.99),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](3781.94),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesSecurityLake" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputSecurityLake(
        components.OutputSecurityLake{
            ID: criblcontrolplanesdkgo.Pointer("security-lake-output"),
            Type: components.OutputSecurityLakeTypeSecurityLake,
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
            SignatureVersion: components.OutputSecurityLakeSignatureVersionV2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: "arn:aws:iam::123456789012:role/my-role",
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](7772.47),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(false),
            ObjectACL: components.ObjectACLOptionsAwsExecRead.ToPointer(),
            StorageClass: components.StorageClassOptionsDeepArchive.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAwsKms.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            BaseFileName: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](3373.72),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](5153.34),
            HeaderLine: criblcontrolplanesdkgo.Pointer("<value>"),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](1755.89),
            OnBackpressure: components.BackpressureBehaviorOptions1Drop.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(true),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](5350.49),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](7874.9),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](874.88),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](7218.41),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](7380.27),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](6476.19),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](6966.41),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](7652.27),
            AccountID: "123456789012",
            CustomSource: "my-custom-source",
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(true),
            ParquetVersion: components.ParquetVersionOptionsParquet24.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](2309.31),
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
            Description: criblcontrolplanesdkgo.Pointer("replicate generously representation blah fully around"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](7237.78),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](287.34),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](1915.82),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesSentinel" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputSentinel(
        components.OutputSentinel{
            ID: criblcontrolplanesdkgo.Pointer("sentinel-output"),
            Type: components.OutputSentinelTypeSentinel,
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
            KeepAlive: criblcontrolplanesdkgo.Pointer(false),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](4399.27),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](9319.08),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](5469.02),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](6136.59),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1433.57),
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
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            AuthType: components.AuthTypeEnumOauth.ToPointer(),
            LoginURL: "https://login.microsoftonline.com",
            Secret: "client-secret",
            ClientID: "client-id",
            Scope: criblcontrolplanesdkgo.Pointer("<value>"),
            EndpointURLConfiguration: components.EndpointConfigurationURL,
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](6699.04),
            Description: criblcontrolplanesdkgo.Pointer("pleasure helpfully pulse upside-down"),
            Format: components.OutputSentinelFormatAdvanced.ToPointer(),
            CustomSourceExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomDropWhenNull: criblcontrolplanesdkgo.Pointer(false),
            CustomEventDelimiter: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomContentType: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomPayloadExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            AdvancedContentType: criblcontrolplanesdkgo.Pointer("<value>"),
            FormatEventCode: criblcontrolplanesdkgo.Pointer("<value>"),
            FormatPayloadCode: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](4904.24),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7621.74),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](3266.81),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputSentinelPqControls{},
            URL: criblcontrolplanesdkgo.Pointer("https://your-workspace.ingest.monitor.azure.com"),
            DcrID: criblcontrolplanesdkgo.Pointer("<id>"),
            DceEndpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            StreamName: criblcontrolplanesdkgo.Pointer("<value>"),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesSentinelOneAiSiem" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputSentinelOneAiSiem(
        components.OutputSentinelOneAiSiem{
            ID: criblcontrolplanesdkgo.Pointer("sentinel-one-ai-siem-output"),
            Type: components.OutputSentinelOneAiSiemTypeSentinelOneAiSiem,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Region: components.RegionUs,
            Endpoint: components.AISIEMEndpointPathRootServicesCollectorEvent,
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1879.1),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](3366.3),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](571.7),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](8731.16),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6457.57),
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
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("tinted interesting however yieldingly mid unlike once er"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            BaseURL: criblcontrolplanesdkgo.Pointer("https://grimy-noon.org/"),
            HostExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            SourceExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            SourceTypeExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceCategoryExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceNameExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceVendorExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            EventTypeExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            Host: criblcontrolplanesdkgo.Pointer("needy-tail.org"),
            Source: criblcontrolplanesdkgo.Pointer("<value>"),
            SourceType: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceCategory: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceName: criblcontrolplanesdkgo.Pointer("<value>"),
            DataSourceVendor: criblcontrolplanesdkgo.Pointer("<value>"),
            EventType: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](7517.5),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](134.42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](7333.26),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesServiceNow" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputServiceNow(
        components.OutputServiceNow{
            ID: criblcontrolplanesdkgo.Pointer("servicenow-output"),
            Type: components.OutputServiceNowTypeServiceNow,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Endpoint: "ingest.lightstep.com:443",
            TokenSecret: "your-token-secret",
            AuthTokenName: criblcontrolplanesdkgo.Pointer("<value>"),
            OtlpVersion: components.OtlpVersionOptions1OneDot3Dot1,
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4039.66),
            Protocol: components.ProtocolOptionsHTTP,
            Compress: components.CompressionOptions4Deflate.ToPointer(),
            HTTPCompress: components.CompressionOptions5Gzip.ToPointer(),
            HTTPTracesEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPMetricsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPLogsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](3207.88),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](5233.08),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7948.24),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsPayloadAndHeaders.ToPointer(),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](5268.97),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](9870.05),
            KeepAlive: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("boastfully though verbally platter"),
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
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
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
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](5472.53),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](3475.52),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](3068.73),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesSignalfx" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
            ID: criblcontrolplanesdkgo.Pointer("signalfx-output"),
            Type: components.OutputSignalfxTypeSignalfx,
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
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
            Realm: "us0",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5119.86),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](5622.39),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](5603.69),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](541.38),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](554.97),
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
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("polite reopen lovingly above although into along"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8210.07),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](6689.36),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](3904.75),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesSnmp" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputSnmp(
        components.OutputSnmp{
            ID: criblcontrolplanesdkgo.Pointer("snmp-output"),
            Type: components.OutputSnmpTypeSnmp,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Hosts: []components.OutputSnmpHost{
                components.OutputSnmpHost{
                    Host: "192.168.1.1",
                    Port: 161,
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](1120.23),
            Description: criblcontrolplanesdkgo.Pointer("whereas while meh oof behind traffic birdbath ha hm whereas"),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesSns" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputSns(
        components.OutputSns{
            ID: criblcontrolplanesdkgo.Pointer("sns-output"),
            Type: components.OutputSnsTypeSns,
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
            TopicArn: "arn:aws:sns:us-east-1:123456789012:my-topic",
            MessageGroupID: "my-message-group",
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](2825.52),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfSecret.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.OutputSnsSignatureVersionV2.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](9771.63),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("bare phooey commonly instead phew consequently exasperation step"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](2374.95),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](457.14),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](3517.03),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputSnsPqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesSplunk" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputSplunk(
        components.OutputSplunk{
            ID: criblcontrolplanesdkgo.Pointer("splunk-output"),
            Type: components.OutputSplunkTypeSplunk,
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
            NestedFields: components.NestedFieldSerializationOptionsNone.ToPointer(),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](1830.9),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](6698.03),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            EnableMultiMetrics: criblcontrolplanesdkgo.Pointer(true),
            EnableACK: criblcontrolplanesdkgo.Pointer(false),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(false),
            MaxS2Sversion: components.MaxS2SVersionOptionsV3.ToPointer(),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("after baa though zowie about scratchy unto"),
            MaxFailedHealthChecks: criblcontrolplanesdkgo.Pointer[float64](8890.04),
            Compress: components.CompressionOptionsAuto.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6766.21),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2535.92),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6333.18),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputSplunkPqControls{},
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
### Example Usage: OutputCreateExamplesSplunkHec

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesSplunkHec" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputSplunkHec(
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
                "<value 2>",
                "<value 3>",
            },
            LoadBalanced: criblcontrolplanesdkgo.Pointer(false),
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
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](7743.2),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](8796),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](3591.27),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](1219.61),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5995.45),
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
            EnableMultiMetrics: criblcontrolplanesdkgo.Pointer(false),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("hmph separately mobilize sugary deliberately since unbearably"),
            URL: criblcontrolplanesdkgo.Pointer("https://potable-seafood.com"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(true),
            Urls: []components.OutputSplunkHecURL{
                components.OutputSplunkHecURL{
                    URL: "https://far-off-chairperson.org/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](3495.92),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](7882.73),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3036.32),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](675.99),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7321.13),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](2936.9),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputSplunkHecPqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesSplunkLb" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputSplunkLb(
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
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](5335.2),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](7027.9),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](8151.45),
            NestedFields: components.NestedFieldSerializationOptionsJSON.ToPointer(),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](1626.21),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](6511.23),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            EnableMultiMetrics: criblcontrolplanesdkgo.Pointer(false),
            EnableACK: criblcontrolplanesdkgo.Pointer(true),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(true),
            MaxS2Sversion: components.MaxS2SVersionOptionsV3.ToPointer(),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            IndexerDiscovery: criblcontrolplanesdkgo.Pointer(false),
            SenderUnhealthyTimeAllowance: criblcontrolplanesdkgo.Pointer[float64](4137.33),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("hyena lest trustworthy sedately dilate as ack"),
            MaxFailedHealthChecks: criblcontrolplanesdkgo.Pointer[float64](8723.07),
            Compress: components.CompressionOptionsAlways.ToPointer(),
            IndexerDiscoveryConfigs: &components.IndexerDiscoveryConfigs{
                Site: "<value>",
                MasterURI: "https://quarrelsome-phrase.biz/",
                RefreshIntervalSec: 7207.13,
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                AuthTokens: []components.OutputSplunkLbAuthToken{
                    components.OutputSplunkLbAuthToken{
                        AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
                        AuthToken: criblcontrolplanesdkgo.Pointer("<value>"),
                        TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    },
                },
                AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
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
                    Weight: criblcontrolplanesdkgo.Pointer[float64](8212.15),
                },
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](979.79),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](4046.95),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](2626.82),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesSqs" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputSqs(
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
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](8892.03),
            MaxQueueSize: criblcontrolplanesdkgo.Pointer[float64](1756.31),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](6677.81),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](9360.69),
            MaxInProgress: criblcontrolplanesdkgo.Pointer[float64](2064.65),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("submitter as intently mortally pale vainly knickers"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9077.33),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7720.51),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](7048.98),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputSqsPqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesStatsd" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputStatsd(
        components.OutputStatsd{
            ID: criblcontrolplanesdkgo.Pointer("statsd-output"),
            Type: components.OutputStatsdTypeStatsd,
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
            Protocol: components.DestinationProtocolOptionsUDP,
            Host: "localhost",
            Port: 8125,
            Mtu: criblcontrolplanesdkgo.Pointer[float64](5877.5),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](4783.22),
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](920.15),
            Description: criblcontrolplanesdkgo.Pointer("officially independence yet nervously what yuppify futon yearningly"),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](7965.79),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](4035.15),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9765.21),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](5342.07),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](4943.9),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesStatsdExt" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputStatsdExt(
        components.OutputStatsdExt{
            ID: criblcontrolplanesdkgo.Pointer("statsd-ext-output"),
            Type: components.OutputStatsdExtTypeStatsdExt,
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
            Protocol: components.DestinationProtocolOptionsUDP,
            Host: "localhost",
            Port: 8125,
            Mtu: criblcontrolplanesdkgo.Pointer[float64](2119.13),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6238.24),
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](3006.42),
            Description: criblcontrolplanesdkgo.Pointer("acquaintance cannon councilman proper brr"),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](9750.99),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](984.97),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9501.13),
            PqMode: components.ModeOptionsAlways.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8671.5),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1157.66),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesSumoLogic" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputSumoLogic(
        components.OutputSumoLogic{
            ID: criblcontrolplanesdkgo.Pointer("sumo-logic-output"),
            Type: components.OutputSumoLogicTypeSumoLogic,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            URL: "https://endpoint1.collection.us2.sumologic.com",
            CustomSource: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomCategory: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.OutputSumoLogicDataFormatJSON.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](2454.78),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](6558.45),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](386.08),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](6525.84),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1666.15),
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
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](9130.6),
            Description: criblcontrolplanesdkgo.Pointer("ew rally readjust"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](3031.72),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7340.67),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1424.73),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputSumoLogicPqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesSyslog" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputSyslog(
        components.OutputSyslog{
            ID: criblcontrolplanesdkgo.Pointer("syslog-output"),
            Type: components.OutputSyslogTypeSyslog,
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
            Protocol: components.OutputSyslogProtocolTCP.ToPointer(),
            Facility: components.FacilityThree.ToPointer(),
            Severity: components.OutputSyslogSeverityFive.ToPointer(),
            AppName: criblcontrolplanesdkgo.Pointer("<value>"),
            MessageFormat: components.MessageFormatRfc3164.ToPointer(),
            TimestampFormat: components.TimestampFormatIso8601.ToPointer(),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            OctetCountFraming: criblcontrolplanesdkgo.Pointer(false),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("sternly meaningfully next banish hypothesise obedient inside lest"),
            LoadBalanced: criblcontrolplanesdkgo.Pointer(true),
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](514),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(true),
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "wicked-bend.com",
                    Port: 9502.76,
                    TLS: components.TLSOptionsHostsItemsOff.ToPointer(),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1086.4),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](2481.04),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6460.17),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](4936.63),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](4135.93),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](5785.19),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            MaxRecordSize: criblcontrolplanesdkgo.Pointer[float64](8293.42),
            UDPDNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](9163.14),
            EnableIPSpoofing: criblcontrolplanesdkgo.Pointer(false),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](667.32),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2272.28),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](3939.53),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputSyslogPqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesTcpjson" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputTcpjson(
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
            LoadBalanced: criblcontrolplanesdkgo.Pointer(false),
            Compression: components.CompressionOptions1Gzip.ToPointer(),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(false),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](8617.16),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](2926.39),
            TokenTTLMinutes: criblcontrolplanesdkgo.Pointer[float64](4733.74),
            SendHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("querulous bah next"),
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](10090),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "wicked-bend.com",
                    Port: 9502.76,
                    TLS: components.TLSOptionsHostsItemsOff.ToPointer(),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1086.4),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](1462.84),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](6447.76),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](2014.57),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](9647.54),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8742.12),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](8669.78),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsDrop.ToPointer(),
            PqControls: &components.OutputTcpjsonPqControls{},
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
### Example Usage: OutputCreateExamplesWavefront

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesWavefront" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputWavefront(
        components.OutputWavefront{
            ID: criblcontrolplanesdkgo.Pointer("wavefront-output"),
            Type: components.OutputWavefrontTypeWavefront,
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
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsSecret.ToPointer(),
            Domain: "longboard",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](3483.48),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](3129.76),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](5185.4),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](1527.25),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5851.43),
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
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("cool unlike crumble circulate materialise intermesh male incidentally celebrate"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](69.34),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](7460.02),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1638.93),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesWebhook" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputWebhook(
        components.OutputWebhook{
            ID: criblcontrolplanesdkgo.Pointer("webhook-output"),
            Type: components.OutputWebhookTypeWebhook,
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
            Method: components.MethodOptionsPost.ToPointer(),
            Format: components.OutputWebhookFormatNdjson.ToPointer(),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](2157.18),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](2167.96),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](3492.97),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](2064.64),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](934.33),
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
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            AuthType: components.OutputWebhookAuthenticationTypeToken.ToPointer(),
            TLS: &components.TLSSettingsClientSideType1{
                Disabled: criblcontrolplanesdkgo.Pointer(false),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
            },
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](3175.03),
            LoadBalanced: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("uproot pocket-watch before apricot effector mostly cruelty pro vice"),
            CustomSourceExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomDropWhenNull: criblcontrolplanesdkgo.Pointer(false),
            CustomEventDelimiter: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomContentType: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomPayloadExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            AdvancedContentType: criblcontrolplanesdkgo.Pointer("<value>"),
            FormatEventCode: criblcontrolplanesdkgo.Pointer("<value>"),
            FormatPayloadCode: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](7799.85),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2604.07),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](6730.04),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqGzip.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputWebhookPqControls{},
            Username: criblcontrolplanesdkgo.Pointer("Leatha.Reynolds52"),
            Password: criblcontrolplanesdkgo.Pointer("saDFsa6Z7GLu206"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://well-worn-minister.biz/"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](1647.17),
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
            URL: criblcontrolplanesdkgo.Pointer("https://example.com/webhook"),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Urls: []components.OutputWebhookURL{
                components.OutputWebhookURL{
                    URL: "https://warmhearted-pigsty.name",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1557.06),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](5079.2),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5978.28),
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesWizHec" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputWizHec(
        components.OutputWizHec{
            ID: criblcontrolplanesdkgo.Pointer("wiz-hec-output"),
            Type: components.OutputWizHecTypeWizHec,
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
            LoadBalanced: false,
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
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv13.ToPointer(),
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](7233.72),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](8699.57),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](1501.04),
            Compress: criblcontrolplanesdkgo.Pointer(false),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](7438.51),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](9385.01),
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
            EnableMultiMetrics: false,
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsDrop.ToPointer(),
            WizConnectorID: "00000000-0000-0000-0000-000000000000",
            WizEnvironment: "test",
            DataCenter: "us1",
            WizSourcetype: "placeholder",
            Description: criblcontrolplanesdkgo.Pointer("rekindle for clear cassava overdue whoa cram incidentally"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](6434.77),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](2924.72),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](4562.74),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputWizHecPqControls{},
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
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
### Example Usage: OutputCreateExamplesXsiam

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesXsiam" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputXsiam(
        components.OutputXsiam{
            ID: criblcontrolplanesdkgo.Pointer("xsiam-output"),
            Type: components.OutputXsiamTypeXsiam,
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
            LoadBalanced: criblcontrolplanesdkgo.Pointer(false),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](4531.24),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](5109.16),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](5522.41),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](5473.45),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](2900.95),
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
            AuthType: components.OutputXsiamAuthenticationMethodSecret.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 3612.5,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](2452.56),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](1370.95),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](434.99),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1512.93),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](7289.7),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](7576.49),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            ThrottleRateReqPerSec: criblcontrolplanesdkgo.Pointer[int64](292667),
            OnBackpressure: components.BackpressureBehaviorOptionsQueue.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](3093.15),
            Description: criblcontrolplanesdkgo.Pointer("fabricate till hopelessly clamor bah icy apropos cruelly barring ick"),
            URL: criblcontrolplanesdkgo.Pointer("https://portly-vanadyl.com/"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(true),
            Urls: []components.OutputXsiamURL{
                components.OutputXsiamURL{
                    URL: "https://wretched-willow.net/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1059.02),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](6916.46),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](3358.5),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(false),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](8566.53),
            PqMode: components.ModeOptionsBackpressure.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](8234.29),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](1314.61),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("<value>"),
            PqPath: criblcontrolplanesdkgo.Pointer("<value>"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputXsiamPqControls{},
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

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputExamplesDefault" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputDefault(
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
            DefaultID: "my-default-output",
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