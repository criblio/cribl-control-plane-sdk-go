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
        operations.CreateOutputOutputAzureBlob{
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
            },
            ContainerName: "my-container",
            CreateContainer: criblcontrolplanesdkgo.Pointer(false),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](1),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("C.Time.strftime(_time ? _time : Date.now()/1000, '%Y/%m/%d')"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("`CriblOut`"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("`.${C.env[\"CRIBL_WORKER_ID\"]}.${__format}${__compression === \"gzip\" ? \".gz\" : \"\"}`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](32),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            HeaderLine: criblcontrolplanesdkgo.Pointer(""),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](64),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2614.76),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](6389.94),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.73),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](166.54),
            },
            AuthType: components.AuthenticationMethodOptionsManual.ToPointer(),
            StorageClass: operations.CreateOutputBlobAccessTierInferred.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("each before down ouch inexperienced below vaguely celebrated past quizzically"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("1MB"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
        operations.CreateOutputOutputAzureDataExplorer{
            ID: "azure-data-explorer-output",
            Type: operations.CreateOutputTypeAzureDataExplorerAzureDataExplorer,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            ClusterURL: "https://mycluster.kusto.windows.net",
            Database: "mydatabase",
            Table: "mytable",
            ValidateDatabaseSettings: criblcontrolplanesdkgo.Pointer(true),
            IngestMode: operations.CreateOutputIngestionModeStreaming.ToPointer(),
            OauthEndpoint: components.MicrosoftEntraIDAuthenticationEndpointOptionsSaslHTTPSLoginMicrosoftonlineCom,
            TenantID: "tenant-id",
            ClientID: "client-id",
            Scope: "https://mycluster.kusto.windows.net/.default",
            OauthType: operations.CreateOutputOauthTypeAuthenticationMethodClientSecret,
            Description: criblcontrolplanesdkgo.Pointer("rationale fundraising er muffled"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("client-secret"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Certificate: &operations.CreateOutputCertificate{
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            Format: components.DataFormatOptionsJSON.ToPointer(),
            Compress: components.CompressionOptions2Gzip,
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("1MB"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(true),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
            IsMappingObj: criblcontrolplanesdkgo.Pointer(false),
            MappingObj: criblcontrolplanesdkgo.Pointer("<value>"),
            MappingRef: criblcontrolplanesdkgo.Pointer("<value>"),
            IngestURL: criblcontrolplanesdkgo.Pointer("https://precious-transparency.org/"),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            StagePath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/staging"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("`.${C.env[\"CRIBL_WORKER_ID\"]}.${__format}${__compression === \"gzip\" ? \".gz\" : \"\"}`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](32),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](1),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2614.76),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](6389.94),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.73),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](166.54),
            },
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushImmediately: criblcontrolplanesdkgo.Pointer(false),
            RetainBlobOnSuccess: criblcontrolplanesdkgo.Pointer(false),
            ExtentTags: []operations.CreateOutputExtentTag{
                operations.CreateOutputExtentTag{
                    Prefix: operations.CreateOutputPrefixOptionalIngestBy.ToPointer(),
                    Value: "<value>",
                },
            },
            IngestIfNotExists: []operations.CreateOutputIngestIfNotExist{
                operations.CreateOutputIngestIfNotExist{
                    Value: "<value>",
                },
            },
            ReportLevel: operations.CreateOutputReportLevelFailuresOnly.ToPointer(),
            ReportMethod: operations.CreateOutputReportMethodQueue.ToPointer(),
            AdditionalProperties: []operations.CreateOutputAdditionalProperty{
                operations.CreateOutputAdditionalProperty{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5522.82,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsAzureDataExplorer{},
            TemplateClusterURL: criblcontrolplanesdkgo.Pointer("https://breakable-backburn.info"),
            TemplateDatabase: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateTable: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateTenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateClientID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateScope: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateFormat: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateIngestURL: criblcontrolplanesdkgo.Pointer("https://animated-discourse.info/"),
        },
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
        operations.CreateOutputOutputAzureEventhub{
            ID: "azure-eventhub-output",
            Type: operations.CreateOutputTypeAzureEventhubAzureEventhub,
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
            Brokers: []string{
                "myeventhub.servicebus.windows.net:9093",
            },
            Topic: "logs",
            Ack: components.AcknowledgmentsOptionsLeader.ToPointer(),
            Format: components.RecordDataFormatOptionsJSON.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](768),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](1000),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                Password: criblcontrolplanesdkgo.Pointer("7VJ9TjHGl3iVy2A"),
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
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("monthly although strident replacement boo phrase after"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsAzureEventhub{},
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
        operations.CreateOutputOutputAzureLogs{
            ID: "azure-logs-output",
            Type: operations.CreateOutputTypeAzureLogsAzureLogs,
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
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1024),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
            APIURL: criblcontrolplanesdkgo.Pointer(".ods.opinsights.azure.com"),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 2572.31,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: operations.CreateOutputAuthenticationMethodAzureLogsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("rival consign pressure"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsAzureLogs{},
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
        operations.CreateOutputOutputChronicle{
            ID: "chronicle-output",
            Type: operations.CreateOutputTypeChronicleChronicle,
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
            APIVersion: criblcontrolplanesdkgo.Pointer("v1alpha"),
            AuthenticationMethod: operations.CreateOutputAuthenticationMethodChronicleServiceAccount.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 3484.86,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            Region: "us",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1024),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](90),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](1733.6),
            IngestionMethod: criblcontrolplanesdkgo.Pointer("ImportLogs"),
            Namespace: criblcontrolplanesdkgo.Pointer("<value>"),
            LogType: "UNKNOWN",
            LogTextField: criblcontrolplanesdkgo.Pointer("<value>"),
            GcpProjectID: "my-project",
            GcpInstance: "customer-id",
            CustomLabels: []operations.CreateOutputCustomLabel{
                operations.CreateOutputCustomLabel{
                    Key: "<key>",
                    Value: "<value>",
                    RbacEnabled: criblcontrolplanesdkgo.Pointer(false),
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("loyally glider till hunger but ouch gosh geez ick warmly"),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            ServiceAccountCredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsChronicle{},
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
        operations.CreateOutputOutputClickHouse{
            ID: "clickhouse-output",
            Type: operations.CreateOutputTypeClickHouseClickHouse,
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
            URL: "http://localhost:8123/",
            AuthType: operations.CreateOutputAuthenticationTypeClickHouseNone.ToPointer(),
            Database: "mydb",
            TableName: "mytable",
            Format: operations.CreateOutputFormatClickHouseJSONCompactEachRowWithNames.ToPointer(),
            MappingType: operations.CreateOutputMappingTypeAutomatic.ToPointer(),
            AsyncInserts: criblcontrolplanesdkgo.Pointer(false),
            TLS: &components.TLSSettingsClientSideType1{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 4759.15,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            DumpFormatErrorsToDisk: criblcontrolplanesdkgo.Pointer(false),
            StatsDestination: &operations.CreateOutputStatsDestination{
                URL: criblcontrolplanesdkgo.Pointer("https://fat-packaging.net/"),
                Database: criblcontrolplanesdkgo.Pointer("<value>"),
                TableName: criblcontrolplanesdkgo.Pointer("<value>"),
                AuthType: criblcontrolplanesdkgo.Pointer("<value>"),
                Username: criblcontrolplanesdkgo.Pointer("Agustin_Cremin87"),
                SQLUsername: criblcontrolplanesdkgo.Pointer("<value>"),
                Password: criblcontrolplanesdkgo.Pointer("A9DUjJtHSnXqh9S"),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("whole gah charlatan"),
            Username: criblcontrolplanesdkgo.Pointer("Mariano16"),
            Password: criblcontrolplanesdkgo.Pointer("iycsVNWAr4li2p5"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            SQLUsername: criblcontrolplanesdkgo.Pointer("<value>"),
            WaitForAsyncInserts: criblcontrolplanesdkgo.Pointer(true),
            ExcludeMappingFields: []string{
                "<value 1>",
            },
            DescribeTable: criblcontrolplanesdkgo.Pointer("<value>"),
            ColumnMappings: []operations.CreateOutputColumnMapping{
                operations.CreateOutputColumnMapping{
                    ColumnName: "<value>",
                    ColumnType: criblcontrolplanesdkgo.Pointer("<value>"),
                    ColumnValueExpression: "<value>",
                },
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsClickHouse{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://sticky-granny.net/"),
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
        operations.CreateOutputOutputCloudflareR2{
            ID: "cloudflare-r2-output",
            Type: operations.CreateOutputTypeCloudflareR2CloudflareR2,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Endpoint: "https://account-id.r2.cloudflarestorage.com",
            Bucket: "my-bucket",
            AwsAuthenticationMethod: operations.CreateOutputAuthenticationMethodCloudflareR2Auto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "<value>",
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions5V4.ToPointer(),
            ObjectACL: "<value>",
            StorageClass: components.StorageClassOptions2ReducedRedundancy.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionOptionsAes256.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("C.Time.strftime(_time ? _time : Date.now()/1000, '%Y/%m/%d')"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("`CriblOut`"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("`.${C.env[\"CRIBL_WORKER_ID\"]}.${__format}${__compression === \"gzip\" ? \".gz\" : \"\"}`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](32),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            HeaderLine: criblcontrolplanesdkgo.Pointer(""),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](64),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2614.76),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](6389.94),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.73),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](166.54),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](4),
            Description: criblcontrolplanesdkgo.Pointer("kiddingly lively provided neatly importance instantly effector"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("1MB"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
        operations.CreateOutputOutputCloudwatch{
            ID: "cloudwatch-output",
            Type: operations.CreateOutputTypeCloudwatchCloudwatch,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            LogGroupName: "my-log-group",
            LogStreamName: "my-log-stream",
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "us-east-1",
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            MaxQueueSize: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](1024),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("often configuration why breastplate ick beside amidst sauerkraut aw"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsCloudwatch{},
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
        operations.CreateOutputOutputConfluentCloud{
            ID: "confluent-cloud-output",
            Type: operations.CreateOutputTypeConfluentCloudConfluentCloud,
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
            Topic: "logs",
            Ack: components.AcknowledgmentsOptions1Leader.ToPointer(),
            Format: components.RecordDataFormatOptions1JSON.ToPointer(),
            Compression: components.CompressionOptions3Gzip.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](768),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](1000),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType1{
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
                DefaultKeySchemaID: criblcontrolplanesdkgo.Pointer[float64](98.58),
                DefaultValueSchemaID: criblcontrolplanesdkgo.Pointer[float64](7633.9),
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
                Username: criblcontrolplanesdkgo.Pointer("Kim_Fadel"),
                Password: criblcontrolplanesdkgo.Pointer("2TjsMJ6T9VV12ax"),
                AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslPlain.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(false),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://gloomy-metal.org/"),
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
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("quiet hopelessly outrank blacken reclassify sunbathe institute memorise"),
            ProtobufLibraryID: criblcontrolplanesdkgo.Pointer("<id>"),
            ProtobufEncodingID: criblcontrolplanesdkgo.Pointer("<id>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsConfluentCloud{},
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
        operations.CreateOutputOutputCriblHTTP{
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
            LoadBalanced: criblcontrolplanesdkgo.Pointer(true),
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
                "<value 1>",
                "<value 2>",
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
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 7159.8,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            AuthTokens: []components.ItemsTypeAuthTokens1{
                components.ItemsTypeAuthTokens1{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("wherever whoa till splay"),
                },
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("castanet lender godparent seldom beyond courageously thankfully easy aha"),
            URL: criblcontrolplanesdkgo.Pointer("https://zesty-remark.net/"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Urls: []components.ItemsTypeUrls{
                components.ItemsTypeUrls{
                    URL: "https://worldly-pocket-watch.com/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1),
                    TemplateURL: criblcontrolplanesdkgo.Pointer("https://clueless-jellyfish.net/"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](600),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](300),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsCriblHTTP{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://low-pacemaker.biz"),
        },
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
        operations.CreateOutputOutputCriblLake{
            ID: "cribl-lake-output",
            Type: operations.CreateOutputTypeCriblLakeCriblLake,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Bucket: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            StagePath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/staging"),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            ObjectACL: components.ObjectACLOptionsPrivate.ToPointer(),
            StorageClass: components.StorageClassOptionsStandardIa.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAwsKms.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            BaseFileName: criblcontrolplanesdkgo.Pointer("`CriblOut`"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("`.${C.env[\"CRIBL_WORKER_ID\"]}.${__format}${__compression === \"gzip\" ? \".gz\" : \"\"}`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](64),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            HeaderLine: criblcontrolplanesdkgo.Pointer(""),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](64),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2614.76),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](6389.94),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.73),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](166.54),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](100),
            AwsAuthenticationMethod: components.AwsAuthenticationMethodOptionsAuto.ToPointer(),
            Format: components.FormatOptionsJSON.ToPointer(),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](1),
            Description: criblcontrolplanesdkgo.Pointer("amongst round huzzah while direct abaft powerfully"),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
        operations.CreateOutputOutputCriblSearchEngine{
            ID: "cribl-search-engine-output",
            Type: operations.CreateOutputTypeCriblSearchEngineCriblSearchEngine,
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
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            AuthTokens: []components.ItemsTypeAuthTokens1{
                components.ItemsTypeAuthTokens1{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("bob often heating mixture incidentally rewrite cow why"),
                },
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("flickering equally defensive worth miserably unexpectedly yuck ample"),
            URL: criblcontrolplanesdkgo.Pointer("https://0.0.0.0:10200"),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Urls: []components.ItemsTypeUrls{
                components.ItemsTypeUrls{
                    URL: "https://unkempt-travel.net/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1),
                    TemplateURL: criblcontrolplanesdkgo.Pointer("https://timely-sand.org/"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](600),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](300),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsCriblSearchEngine{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://muted-backbone.org"),
        },
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
        operations.CreateOutputOutputCriblTCP{
            ID: "cribl-tcp-output",
            Type: operations.CreateOutputTypeCriblTCPCriblTCP,
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
            Compression: components.CompressionOptions1Gzip.ToPointer(),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(false),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
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
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            TokenTTLMinutes: criblcontrolplanesdkgo.Pointer[float64](60),
            AuthTokens: []components.ItemsTypeAuthTokens{
                components.ItemsTypeAuthTokens{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("uh-huh silver vicinity"),
                },
            },
            ExcludeFields: []string{
                "<value 1>",
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("until sashay woot keenly traffic bright lounge information"),
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](10090),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "tough-issue.info",
                    Port: 10300,
                    TLS: components.TLSOptionsHostsItemsInherit.ToPointer(),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1),
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](600),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](0),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsCriblTCP{},
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
        operations.CreateOutputOutputCrowdstrikeNextGenSiem{
            ID: "crowdstrike-next-gen-siem-output",
            Type: operations.CreateOutputTypeCrowdstrikeNextGenSiemCrowdstrikeNextGenSiem,
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
            URL: "https://ingest.us.crowdstrike.com/api/ingest/hec/connection-id/v1/services/collector",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
            Format: components.RequestFormatOptionsJSON,
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5055.78,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("boohoo near knife soybean besides pish gah alongside"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsCrowdstrikeNextGenSiem{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://squeaky-academics.name/"),
        },
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
        operations.CreateOutputOutputDatabricks{
            ID: "databricks-output",
            Type: operations.CreateOutputTypeDatabricksDatabricks,
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
            DestPath: criblcontrolplanesdkgo.Pointer(""),
            StagePath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/staging"),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("C.Time.strftime(_time ? _time : Date.now()/1000, '%Y/%m/%d')"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("`CriblOut`"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("`.${C.env[\"CRIBL_WORKER_ID\"]}.${__format}${__compression === \"gzip\" ? \".gz\" : \"\"}`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](32),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            HeaderLine: criblcontrolplanesdkgo.Pointer(""),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](64),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2614.76),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](6389.94),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.73),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](166.54),
            },
            WorkspaceID: "your-workspace-id",
            Scope: "my-scope",
            ClientID: "your-client-id",
            Catalog: "my-catalog",
            Schema: "my-schema",
            EventsVolumeName: "my-volume",
            ClientTextSecret: "your-client-secret",
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](60),
            Description: criblcontrolplanesdkgo.Pointer("oof parody ad round verbally across"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("1MB"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
        operations.CreateOutputOutputDatadog{
            ID: "datadog-output",
            Type: operations.CreateOutputTypeDatadogDatadog,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            ContentType: operations.CreateOutputSendLogsAsJSON.ToPointer(),
            Message: criblcontrolplanesdkgo.Pointer("<value>"),
            Source: criblcontrolplanesdkgo.Pointer("<value>"),
            Host: criblcontrolplanesdkgo.Pointer("woeful-ad.net"),
            Service: criblcontrolplanesdkgo.Pointer("<value>"),
            Tags: []string{
                "<value 1>",
                "<value 2>",
            },
            BatchByTags: criblcontrolplanesdkgo.Pointer(true),
            AllowAPIKeyFromEvents: criblcontrolplanesdkgo.Pointer(false),
            Severity: operations.CreateOutputSeverityDatadogWarning.ToPointer(),
            Site: operations.CreateOutputDatadogSiteUs.ToPointer(),
            SendCountersAsCount: criblcontrolplanesdkgo.Pointer(false),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 6276.58,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Manual.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](3573.66),
            Description: criblcontrolplanesdkgo.Pointer("to midst phooey"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://quick-godfather.biz"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsDatadog{},
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
        operations.CreateOutputOutputDataset{
            ID: "dataset-output",
            Type: operations.CreateOutputTypeDatasetDataset,
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
            DefaultSeverity: operations.CreateOutputDefaultSeveritySeverityInfo.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8426.7,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            Site: operations.CreateOutputDataSetSiteUs.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Manual.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](8747.36),
            Description: criblcontrolplanesdkgo.Pointer("out furthermore ew bludgeon unbearably nocturnal jaggedly midst knight"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://educated-turret.net/"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsDataset{},
            APIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateCustomURL: criblcontrolplanesdkgo.Pointer("https://apt-knight.name/"),
        },
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
        operations.CreateOutputOutputDiskSpool{
            ID: "disk-spool-output",
            Type: operations.CreateOutputTypeDiskSpoolDiskSpool,
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
            TimeWindow: criblcontrolplanesdkgo.Pointer("10m"),
            MaxDataSize: criblcontrolplanesdkgo.Pointer("1GB"),
            MaxDataTime: criblcontrolplanesdkgo.Pointer("24h"),
            Compress: components.CompressionOptionsPersistenceGzip.ToPointer(),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("blah daintily ah"),
        },
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
        operations.CreateOutputOutputDlS3{
            ID: "dl-s3-output",
            Type: operations.CreateOutputTypeDlS3DlS3,
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
            Bucket: "my-bucket",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer(""),
            ObjectACL: components.ObjectACLOptionsPrivate.ToPointer(),
            StorageClass: components.StorageClassOptionsReducedRedundancy.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAwsKms.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("`CriblOut`"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("`.${C.env[\"CRIBL_WORKER_ID\"]}.${__format}${__compression === \"gzip\" ? \".gz\" : \"\"}`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](32),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            HeaderLine: criblcontrolplanesdkgo.Pointer(""),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](64),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2614.76),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](6389.94),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.73),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](166.54),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](4),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](100),
            PartitioningFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Description: criblcontrolplanesdkgo.Pointer("segregate helplessly clinking"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("1MB"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
        operations.CreateOutputOutputDynatraceHTTP{
            ID: "dynatrace-http-output",
            Type: operations.CreateOutputTypeDynatraceHTTPDynatraceHTTP,
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
            Method: components.MethodOptionsPost.ToPointer(),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 858.14,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: operations.CreateOutputAuthenticationTypeDynatraceHTTPToken.ToPointer(),
            Format: operations.CreateOutputFormatDynatraceHTTPJSONArray,
            Endpoint: operations.CreateOutputEndpointCloud,
            TelemetryType: operations.CreateOutputTelemetryTypeLogs,
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](2231.72),
            Description: criblcontrolplanesdkgo.Pointer("when wry of intently duh harvest unfortunately till intrigue"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsDynatraceHTTP{},
            Token: criblcontrolplanesdkgo.Pointer("your-api-key"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            EnvironmentID: criblcontrolplanesdkgo.Pointer("<id>"),
            ActiveGateDomain: criblcontrolplanesdkgo.Pointer("<value>"),
            URL: criblcontrolplanesdkgo.Pointer("https://musty-adaptation.info"),
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://cheap-premise.net"),
        },
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
        operations.CreateOutputOutputDynatraceOtlp{
            ID: "dynatrace-otlp-output",
            Type: operations.CreateOutputTypeDynatraceOtlpDynatraceOtlp,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Protocol: operations.CreateOutputProtocolDynatraceOtlpHTTP,
            Endpoint: "https://your-environment.live.dynatrace.com/api/v2/otlp",
            OtlpVersion: components.OtlpVersionOptions1OneDot3Dot1,
            Compress: components.CompressionOptions4Gzip.ToPointer(),
            HTTPCompress: components.CompressionOptions5Gzip.ToPointer(),
            HTTPTracesEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPMetricsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPLogsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](2048),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](30),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            EndpointType: operations.CreateOutputEndpointTypeSaas,
            TokenSecret: "your-token-secret",
            AuthTokenName: criblcontrolplanesdkgo.Pointer("Authorization"),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("yippee fabricate drat wherever so"),
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
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8446.48,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsDynatraceOtlp{},
        },
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
        operations.CreateOutputOutputElastic{
            ID: "elastic-output",
            Type: operations.CreateOutputTypeElasticElastic,
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
            LoadBalanced: criblcontrolplanesdkgo.Pointer(true),
            Index: "logs",
            DocType: criblcontrolplanesdkgo.Pointer("<value>"),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
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
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 3353.65,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
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
                Username: criblcontrolplanesdkgo.Pointer("Maddison_Willms"),
                Password: criblcontrolplanesdkgo.Pointer("YjRP4IarMDm7WjU"),
                AuthType: components.AuthenticationMethodOptionsAuthManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                ManualAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            ElasticVersion: operations.CreateOutputElasticVersionAuto.ToPointer(),
            ElasticPipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            IncludeDocID: criblcontrolplanesdkgo.Pointer(false),
            WriteAction: operations.CreateOutputWriteActionCreate.ToPointer(),
            RetryPartialErrors: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("thoughtfully scope queasily indelible"),
            URL: criblcontrolplanesdkgo.Pointer("https://bruised-yeast.info/"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Urls: []operations.CreateOutputURLElastic{
                operations.CreateOutputURLElastic{
                    URL: "https://needy-hawk.com",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1),
                    TemplateURL: criblcontrolplanesdkgo.Pointer("https://multicolored-cannon.biz/"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](600),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](300),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsElastic{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://punctual-mixture.org"),
        },
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
        operations.CreateOutputOutputElasticCloud{
            ID: "elastic-cloud-output",
            Type: operations.CreateOutputTypeElasticCloudElasticCloud,
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
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
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
                Username: criblcontrolplanesdkgo.Pointer("Maddison_Willms"),
                Password: criblcontrolplanesdkgo.Pointer("YjRP4IarMDm7WjU"),
                AuthType: components.AuthenticationMethodOptionsAuthManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                ManualAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            ElasticPipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            IncludeDocID: criblcontrolplanesdkgo.Pointer(true),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 3400.37,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("upon moor mysteriously incidentally"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsElasticCloud{},
        },
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
        operations.CreateOutputOutputExabeam{
            ID: "exabeam-output",
            Type: operations.CreateOutputTypeExabeamExabeam,
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
            Region: "us-east1",
            StagePath: "/tmp/staging",
            Endpoint: "https://storage.googleapis.com",
            SignatureVersion: components.SignatureVersionOptions4V4.ToPointer(),
            ObjectACL: components.ObjectACLOptions1Private.ToPointer(),
            StorageClass: components.StorageClassOptions1Standard.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2614.76),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](6389.94),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.73),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](166.54),
            },
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](10),
            EncodedConfiguration: criblcontrolplanesdkgo.Pointer("<value>"),
            CollectorInstanceID: "11112222-3333-4444-5555-666677778888",
            SiteName: criblcontrolplanesdkgo.Pointer("<value>"),
            SiteID: criblcontrolplanesdkgo.Pointer("<id>"),
            TimezoneOffset: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("likewise chapel phooey mount nor"),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
        operations.CreateOutputOutputFilesystem{
            ID: "filesystem-output",
            Type: operations.CreateOutputTypeFilesystemFilesystem,
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
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("C.Time.strftime(_time ? _time : Date.now()/1000, '%Y/%m/%d')"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("`CriblOut`"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("`.${C.env[\"CRIBL_WORKER_ID\"]}.${__format}${__compression === \"gzip\" ? \".gz\" : \"\"}`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](32),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            HeaderLine: criblcontrolplanesdkgo.Pointer(""),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](64),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2614.76),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](6389.94),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.73),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](166.54),
            },
            Description: criblcontrolplanesdkgo.Pointer("meanwhile victoriously founder because before"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("1MB"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
        operations.CreateOutputOutputGoogleChronicle{
            ID: "google-chronicle-output",
            Type: operations.CreateOutputTypeGoogleChronicleGoogleChronicle,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            APIVersion: operations.CreateOutputAPIVersionV1.ToPointer(),
            AuthenticationMethod: operations.CreateOutputAuthenticationMethodGoogleChronicleServiceAccount.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 379.55,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            LogFormatType: operations.CreateOutputSendEventsAsUnstructured,
            Region: criblcontrolplanesdkgo.Pointer("us"),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1024),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](90),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](3260.96),
            Description: criblcontrolplanesdkgo.Pointer("incidentally testify but heartache fooey of"),
            ExtraLogTypes: []operations.CreateOutputExtraLogType{
                operations.CreateOutputExtraLogType{
                    LogType: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("jealously qua before obnoxiously"),
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
            UdmType: operations.CreateOutputUDMTypeLogs.ToPointer(),
            APIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            APIKeySecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            ServiceAccountCredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsGoogleChronicle{},
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
        operations.CreateOutputOutputGoogleCloudLogging{
            ID: "google-cloud-logging-output",
            Type: operations.CreateOutputTypeGoogleCloudLoggingGoogleCloudLogging,
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
            LogLocationType: operations.CreateOutputLogLocationTypeProject,
            LogNameExpression: "my-log",
            SanitizeLogNames: criblcontrolplanesdkgo.Pointer(false),
            PayloadFormat: operations.CreateOutputPayloadFormatText.ToPointer(),
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
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            ThrottleRateReqPerSec: criblcontrolplanesdkgo.Pointer[int64](208589),
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
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](3307.75),
            Description: criblcontrolplanesdkgo.Pointer("apropos amidst steep gee think for lonely snappy council"),
            LogLocationExpression: "my-project",
            PayloadExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsGoogleCloudLogging{},
        },
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
        operations.CreateOutputOutputGoogleCloudStorage{
            ID: "google-cloud-storage-output",
            Type: operations.CreateOutputTypeGoogleCloudStorageGoogleCloudStorage,
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
            Bucket: "my-bucket",
            Region: "us-east1",
            Endpoint: "https://storage.googleapis.com",
            SignatureVersion: components.SignatureVersionOptions4V4.ToPointer(),
            AwsAuthenticationMethod: operations.CreateOutputAuthenticationMethodGoogleCloudStorageManual.ToPointer(),
            StagePath: "/tmp/staging",
            DestPath: criblcontrolplanesdkgo.Pointer(""),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            ObjectACL: components.ObjectACLOptions1Private.ToPointer(),
            StorageClass: components.StorageClassOptions1Archive.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("C.Time.strftime(_time ? _time : Date.now()/1000, '%Y/%m/%d')"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("`CriblOut`"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("`.${C.env[\"CRIBL_WORKER_ID\"]}.${__format}${__compression === \"gzip\" ? \".gz\" : \"\"}`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](32),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            HeaderLine: criblcontrolplanesdkgo.Pointer(""),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](64),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2614.76),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](6389.94),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.73),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](166.54),
            },
            Description: criblcontrolplanesdkgo.Pointer("dutiful parsnip likewise guard warmhearted lucky spattering despite above"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("1MB"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
        operations.CreateOutputOutputGooglePubsub{
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
                "<value 2>",
            },
            TopicName: "my-topic",
            CreateTopic: criblcontrolplanesdkgo.Pointer(false),
            OrderedDelivery: criblcontrolplanesdkgo.Pointer(false),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            GoogleAuthMethod: components.GoogleAuthenticationMethodOptionsManual.ToPointer(),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            BatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            BatchTimeout: criblcontrolplanesdkgo.Pointer[float64](100),
            MaxQueueSize: criblcontrolplanesdkgo.Pointer[float64](100),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](256),
            FlushPeriod: criblcontrolplanesdkgo.Pointer[float64](1),
            MaxInProgress: criblcontrolplanesdkgo.Pointer[float64](10),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("gee freezing with beside defrag lieu devastation"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsGooglePubsub{},
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
        operations.CreateCreateOutputOutputGrafanaCloudUnionCreateOutputOutputGrafanaCloudGrafanaCloud1(
            operations.CreateOutputOutputGrafanaCloudGrafanaCloud1{
                ID: "grafana-cloud-output",
                Type: operations.CreateOutputOutputGrafanaCloudType1GrafanaCloud,
                Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                SystemFields: []string{
                    "<value 1>",
                    "<value 2>",
                },
                Environment: criblcontrolplanesdkgo.Pointer("<value>"),
                Streamtags: []string{
                    "<value 1>",
                },
                LokiURL: "https://logs-prod-us-central1.grafana.net",
                PrometheusURL: criblcontrolplanesdkgo.Pointer("https://smoggy-tomb.org/"),
                Message: criblcontrolplanesdkgo.Pointer("<value>"),
                MessageFormat: components.MessageFormatOptionsProtobuf.ToPointer(),
                Labels: []components.ItemsTypeLabels{
                    components.ItemsTypeLabels{
                        Name: "",
                        Value: "<value>",
                    },
                },
                MetricRenameExpr: criblcontrolplanesdkgo.Pointer("name.replace(/[^a-zA-Z0-9_]/g, '_')"),
                PrometheusAuth: &components.PrometheusAuthType{
                    AuthType: components.AuthenticationTypeOptionsPrometheusAuth1Basic.ToPointer(),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Username: criblcontrolplanesdkgo.Pointer("Pete.Cormier97"),
                    Password: criblcontrolplanesdkgo.Pointer("YqqU_9UCy3i_Jit"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                LokiAuth: &components.PrometheusAuthType{
                    AuthType: components.AuthenticationTypeOptionsPrometheusAuth1Basic.ToPointer(),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Username: criblcontrolplanesdkgo.Pointer("Dayton.Kunde"),
                    Password: criblcontrolplanesdkgo.Pointer("tn2vY0ufsrH9vj_"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                Concurrency: criblcontrolplanesdkgo.Pointer[float64](1),
                MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
                MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
                FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](15),
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
                        HTTPStatus: 8105.6,
                        InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                        BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                        MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                    },
                },
                TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                    TimeoutRetry: false,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
                ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
                OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
                Description: criblcontrolplanesdkgo.Pointer("lace frantically ha forceful underneath times yum partially"),
                Compress: criblcontrolplanesdkgo.Pointer(true),
                PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
                PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
                PqMode: components.ModeOptionsError.ToPointer(),
                PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
                PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
                PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                PqCompress: components.CompressionOptionsPqNone.ToPointer(),
                PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
                PqControls: &operations.CreateOutputOutputGrafanaCloudPqControls1{},
                TemplateLokiURL: criblcontrolplanesdkgo.Pointer("https://deadly-sundae.biz/"),
                TemplatePrometheusURL: criblcontrolplanesdkgo.Pointer("https://distorted-outset.biz/"),
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
        operations.CreateOutputOutputGraphite{
            ID: "graphite-output",
            Type: operations.CreateOutputTypeGraphiteGraphite,
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
            Mtu: criblcontrolplanesdkgo.Pointer[float64](512),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](0),
            Description: criblcontrolplanesdkgo.Pointer("pfft term towards"),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsGraphite{},
        },
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
        operations.CreateOutputOutputHoneycomb{
            ID: "honeycomb-output",
            Type: operations.CreateOutputTypeHoneycombHoneycomb,
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
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 7702.22,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("between yum than lawful morning thoroughly compromise lost negligible"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsHoneycomb{},
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
        operations.CreateOutputOutputHumioHec{
            ID: "humio-hec-output",
            Type: operations.CreateOutputTypeHumioHecHumioHec,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            URL: "https://cloud.us.humio.com/api/v1/ingest/hec",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 9520.58,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("gadzooks jubilantly over silk switch huzzah"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsHumioHec{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://other-interchange.org"),
        },
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
        operations.CreateOutputOutputInfluxdb{
            ID: "influxdb-output",
            Type: operations.CreateOutputTypeInfluxdbInfluxdb,
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
            UseV2API: criblcontrolplanesdkgo.Pointer(false),
            TimestampPrecision: operations.CreateOutputTimestampPrecisionMs.ToPointer(),
            DynamicValueFieldName: criblcontrolplanesdkgo.Pointer(true),
            ValueFieldName: criblcontrolplanesdkgo.Pointer("value"),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 786.03,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: operations.CreateOutputAuthenticationTypeInfluxdbNone.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("terrorise spook meh towards mysteriously"),
            Database: criblcontrolplanesdkgo.Pointer("mydb"),
            Bucket: criblcontrolplanesdkgo.Pointer("<value>"),
            Org: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsInfluxdb{},
            Username: criblcontrolplanesdkgo.Pointer("Mariane_Schneider28"),
            Password: criblcontrolplanesdkgo.Pointer("z3FVqSY2GquoGXg"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://growing-declaration.net/"),
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
        operations.CreateOutputOutputKafka{
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
            Ack: components.AcknowledgmentsOptions1Leader.ToPointer(),
            Format: components.RecordDataFormatOptions1JSON.ToPointer(),
            Compression: components.CompressionOptions3Gzip.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](768),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](1000),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType1{
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
                DefaultKeySchemaID: criblcontrolplanesdkgo.Pointer[float64](98.58),
                DefaultValueSchemaID: criblcontrolplanesdkgo.Pointer[float64](7633.9),
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
                Username: criblcontrolplanesdkgo.Pointer("Kim_Fadel"),
                Password: criblcontrolplanesdkgo.Pointer("2TjsMJ6T9VV12ax"),
                AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslPlain.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(false),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://gloomy-metal.org/"),
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
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("unimpressively firm late under shirk loftily till"),
            ProtobufLibraryID: criblcontrolplanesdkgo.Pointer("<id>"),
            ProtobufEncodingID: criblcontrolplanesdkgo.Pointer("<id>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsKafka{},
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
        operations.CreateOutputOutputKinesis{
            ID: "kinesis-output",
            Type: operations.CreateOutputTypeKinesisKinesis,
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
            StreamName: "my-stream",
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
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](1024),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            Compression: operations.CreateOutputCompressionGzip.ToPointer(),
            UseListShards: criblcontrolplanesdkgo.Pointer(false),
            AsNdjson: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("negotiation ad to eek about phew"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxEventsPerFlush: criblcontrolplanesdkgo.Pointer[float64](500),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsKinesis{},
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
        operations.CreateOutputOutputLoki{
            ID: "loki-output",
            Type: operations.CreateOutputTypeLokiLoki,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            URL: "http://localhost:3100/loki/api/v1/push",
            Message: criblcontrolplanesdkgo.Pointer("<value>"),
            MessageFormat: components.MessageFormatOptionsProtobuf.ToPointer(),
            Labels: []components.ItemsTypeLabels{
                components.ItemsTypeLabels{
                    Name: "",
                    Value: "<value>",
                },
            },
            AuthType: components.AuthenticationTypeOptionsPrometheusAuth1None.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](15),
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
                    HTTPStatus: 3434.64,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            EnableDynamicHeaders: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](9877.1),
            Description: criblcontrolplanesdkgo.Pointer("hospitable hence adjudge"),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Username: criblcontrolplanesdkgo.Pointer("Magdalena_Zulauf-Kuhlman8"),
            Password: criblcontrolplanesdkgo.Pointer("IqSYjCmovUFcOZH"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsLoki{},
        },
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
        operations.CreateOutputOutputMicrosoftFabric{
            ID: "microsoft-fabric-output",
            Type: operations.CreateOutputTypeMicrosoftFabricMicrosoftFabric,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Topic: "logs",
            Ack: components.AcknowledgmentsOptionsLeader.ToPointer(),
            Format: components.RecordDataFormatOptionsJSON.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](768),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](1000),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](30000),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](300),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](10000),
            Sasl: &operations.CreateOutputAuthentication{
                Disabled: false,
                Mechanism: components.SaslMechanismOptionsSasl1Plain.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("$ConnectionString"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                ClientSecretAuthType: operations.CreateOutputClientSecretAuthTypeAuthenticationMethodSecret.ToPointer(),
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
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            BootstrapServer: "myeventstream.servicebus.windows.net:9093",
            Description: criblcontrolplanesdkgo.Pointer("testimonial woot ingratiate clearly"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsMicrosoftFabric{},
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
        operations.CreateOutputOutputMinio{
            ID: "minio-output",
            Type: operations.CreateOutputTypeMinioMinio,
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
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions5V4.ToPointer(),
            ObjectACL: components.ObjectACLOptionsPrivate.ToPointer(),
            StorageClass: components.StorageClassOptions2Standard.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionOptionsAes256.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("C.Time.strftime(_time ? _time : Date.now()/1000, '%Y/%m/%d')"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("`CriblOut`"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("`.${C.env[\"CRIBL_WORKER_ID\"]}.${__format}${__compression === \"gzip\" ? \".gz\" : \"\"}`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](32),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            HeaderLine: criblcontrolplanesdkgo.Pointer(""),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](64),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2614.76),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](6389.94),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.73),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](166.54),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](4),
            Description: criblcontrolplanesdkgo.Pointer("supposing instead elegantly explode slope tasty direct what"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("1MB"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
        operations.CreateOutputOutputMsk{
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
            Ack: components.AcknowledgmentsOptions1Leader.ToPointer(),
            Format: components.RecordDataFormatOptions1JSON.ToPointer(),
            Compression: components.CompressionOptions3Gzip.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](768),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](1000),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType1{
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
                DefaultKeySchemaID: criblcontrolplanesdkgo.Pointer[float64](98.58),
                DefaultValueSchemaID: criblcontrolplanesdkgo.Pointer[float64](7633.9),
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
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("redevelop consequently for"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ProtobufLibraryID: criblcontrolplanesdkgo.Pointer("<id>"),
            ProtobufEncodingID: criblcontrolplanesdkgo.Pointer("<id>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsMsk{},
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
        operations.CreateOutputOutputNetflow{
            ID: "netflow-output",
            Type: operations.CreateOutputTypeNetflowNetflow,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Hosts: []operations.CreateOutputHostNetflow{
                operations.CreateOutputHostNetflow{
                    Host: "localhost",
                    Port: 2055,
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](0),
            EnableIPSpoofing: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("allegation and helplessly knowingly never per excluding angrily finally uncork"),
            MaxRecordSize: criblcontrolplanesdkgo.Pointer[float64](1500),
        },
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
        operations.CreateOutputOutputNewrelic{
            ID: "newrelic-output",
            Type: operations.CreateOutputTypeNewrelicNewrelic,
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
            Region: components.RegionOptionsUs.ToPointer(),
            LogType: criblcontrolplanesdkgo.Pointer(""),
            MessageField: criblcontrolplanesdkgo.Pointer(""),
            Metadata: []operations.CreateOutputMetadatum{
                operations.CreateOutputMetadatum{
                    Name: operations.CreateOutputFieldNameAuditID,
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1024),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 1966.78,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Manual.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](6526.58),
            Description: criblcontrolplanesdkgo.Pointer("whoa cycle furthermore but"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://puny-governance.biz/"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsNewrelic{},
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
        operations.CreateOutputOutputNewrelicEvents{
            ID: "newrelic-events-output",
            Type: operations.CreateOutputTypeNewrelicEventsNewrelicEvents,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Region: components.RegionOptionsUs.ToPointer(),
            AccountID: "123456",
            EventType: "CriblEvent",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1024),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 3594.47,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("utilization sensitize whose hydrocarbon within innocently ouch er officially appertain"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://appropriate-dependency.name"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsNewrelicEvents{},
            APIKey: criblcontrolplanesdkgo.Pointer("your-api-key"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateEventType: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateCustomURL: criblcontrolplanesdkgo.Pointer("https://fluffy-diver.name"),
        },
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
        operations.CreateOutputOutputOpenTelemetry{
            ID: "opentelemetry-output",
            Type: operations.CreateOutputTypeOpenTelemetryOpenTelemetry,
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
            Protocol: components.ProtocolOptionsGrpc.ToPointer(),
            Endpoint: "http://localhost:4317",
            OtlpVersion: operations.CreateOutputOTLPVersionZeroDot10Dot0.ToPointer(),
            Compress: components.CompressionOptions4Gzip.ToPointer(),
            HTTPCompress: components.CompressionOptions5Gzip.ToPointer(),
            AuthType: components.AuthenticationTypeOptionsNone.ToPointer(),
            HTTPTracesEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPMetricsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPLogsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](30),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("next sermon whenever"),
            Username: criblcontrolplanesdkgo.Pointer("Jerod7"),
            Password: criblcontrolplanesdkgo.Pointer("HeFGuNo6pe5ewNs"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 1858.11,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
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
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsOpenTelemetry{},
        },
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
        operations.CreateOutputOutputPrometheus{
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
            MetricRenameExpr: criblcontrolplanesdkgo.Pointer("name.replace(/[^a-zA-Z0-9_]/g, '_')"),
            SendMetadata: criblcontrolplanesdkgo.Pointer(true),
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
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 516.48,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationTypeOptionsPrometheusAuthNone.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("forenenst without which brief fuzzy gah"),
            MetricsFlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](60),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsPrometheus{},
            Username: criblcontrolplanesdkgo.Pointer("Isac38"),
            Password: criblcontrolplanesdkgo.Pointer("zMoVVLUWpStIVEr"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://gloomy-marten.net/"),
        },
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
        operations.CreateOutputOutputRing{
            ID: "ring-output",
            Type: operations.CreateOutputTypeRingRing,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Format: operations.CreateOutputDataFormatRingJSON.ToPointer(),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxDataSize: criblcontrolplanesdkgo.Pointer("1GB"),
            MaxDataTime: criblcontrolplanesdkgo.Pointer("24h"),
            Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("bolster woot athwart"),
        },
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
        operations.CreateOutputOutputRouter{
            ID: "router-output",
            Type: operations.CreateOutputTypeRouterRouter,
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
                    Description: criblcontrolplanesdkgo.Pointer("lasting uh-huh weatherize consequently joyous beloved"),
                    Final: criblcontrolplanesdkgo.Pointer(true),
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("boom kiddingly moral"),
        },
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
        operations.CreateOutputOutputS3{
            ID: "s3-output",
            Type: operations.CreateOutputTypeS3S3,
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
            Bucket: "my-bucket",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer(""),
            ObjectACL: components.ObjectACLOptionsPrivate.ToPointer(),
            StorageClass: components.StorageClassOptionsStandard.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAwsKms.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("C.Time.strftime(_time ? _time : Date.now()/1000, '%Y/%m/%d')"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("`CriblOut`"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("`.${C.env[\"CRIBL_WORKER_ID\"]}.${__format}${__compression === \"gzip\" ? \".gz\" : \"\"}`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](32),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            HeaderLine: criblcontrolplanesdkgo.Pointer(""),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](64),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2614.76),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](6389.94),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.73),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](166.54),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](4),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](100),
            Description: criblcontrolplanesdkgo.Pointer("cross-contamination upside-down above"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("1MB"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
        operations.CreateOutputOutputSecurityLake{
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
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: operations.CreateOutputSignatureVersionSecurityLakeV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: "arn:aws:iam::123456789012:role/my-role",
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            ObjectACL: components.ObjectACLOptionsPrivate.ToPointer(),
            StorageClass: components.StorageClassOptionsIntelligentTiering.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAes256.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            BaseFileName: criblcontrolplanesdkgo.Pointer("`CriblOut`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](32),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            HeaderLine: criblcontrolplanesdkgo.Pointer(""),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](64),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2614.76),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](6389.94),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](7295.73),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](166.54),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](4),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](100),
            AccountID: "123456789012",
            CustomSource: "my-custom-source",
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("1MB"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(true),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("crowded furthermore editor lovable confound which"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
        operations.CreateOutputOutputSentinel{
            ID: "sentinel-output",
            Type: operations.CreateOutputTypeSentinelSentinel,
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
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1000),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 4485.56,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: operations.CreateOutputAuthTypeOauth.ToPointer(),
            LoginURL: "https://login.microsoftonline.com",
            Secret: "client-secret",
            ClientID: "client-id",
            Scope: criblcontrolplanesdkgo.Pointer("https://monitor.azure.com/.default"),
            EndpointURLConfiguration: operations.CreateOutputEndpointConfigurationURL,
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](4955.04),
            Description: criblcontrolplanesdkgo.Pointer("instantly yuck hmph yet uh-huh outfit whistle given though yum"),
            Format: operations.CreateOutputFormatSentinelNdjson.ToPointer(),
            CustomSourceExpression: criblcontrolplanesdkgo.Pointer("__httpOut"),
            CustomDropWhenNull: criblcontrolplanesdkgo.Pointer(false),
            CustomEventDelimiter: criblcontrolplanesdkgo.Pointer("\\n"),
            CustomContentType: criblcontrolplanesdkgo.Pointer("application/x-ndjson"),
            CustomPayloadExpression: criblcontrolplanesdkgo.Pointer("`${events}`"),
            AdvancedContentType: criblcontrolplanesdkgo.Pointer("application/json"),
            FormatEventCode: criblcontrolplanesdkgo.Pointer("<value>"),
            FormatPayloadCode: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsSentinel{},
            URL: criblcontrolplanesdkgo.Pointer("https://your-workspace.ingest.monitor.azure.com"),
            DcrID: criblcontrolplanesdkgo.Pointer("<id>"),
            DceEndpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            StreamName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateLoginURL: criblcontrolplanesdkgo.Pointer("https://sophisticated-alert.com"),
            TemplateSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateClientID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateScope: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://reasonable-swine.info/"),
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
        operations.CreateOutputOutputSentinelOneAiSiem{
            ID: "sentinel-one-ai-siem-output",
            Type: operations.CreateOutputTypeSentinelOneAiSiemSentinelOneAiSiem,
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
            Region: operations.CreateOutputRegionUs,
            Endpoint: operations.CreateOutputAISIEMEndpointPathRootServicesCollectorEvent,
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](5120),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5),
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
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9304.93,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("towards joyfully woot with basket pace cassava for mobilize"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            BaseURL: criblcontrolplanesdkgo.Pointer("https://<Your-S1-Tenant>.sentinelone.net"),
            HostExpression: criblcontrolplanesdkgo.Pointer("__e.host || C.os.hostname()"),
            SourceExpression: criblcontrolplanesdkgo.Pointer("__e.source || (__e.__criblMetrics ? 'metrics' : 'cribl')"),
            SourceTypeExpression: criblcontrolplanesdkgo.Pointer("__e.sourcetype || 'dottedJson'"),
            DataSourceCategoryExpression: criblcontrolplanesdkgo.Pointer("'security'"),
            DataSourceNameExpression: criblcontrolplanesdkgo.Pointer("__e.__dataSourceName || 'cribl'"),
            DataSourceVendorExpression: criblcontrolplanesdkgo.Pointer("__e.__dataSourceVendor || 'cribl'"),
            EventTypeExpression: criblcontrolplanesdkgo.Pointer(""),
            Host: criblcontrolplanesdkgo.Pointer("C.os.hostname()"),
            Source: criblcontrolplanesdkgo.Pointer("cribl"),
            SourceType: criblcontrolplanesdkgo.Pointer("hecRawParser"),
            DataSourceCategory: criblcontrolplanesdkgo.Pointer("security"),
            DataSourceName: criblcontrolplanesdkgo.Pointer("cribl"),
            DataSourceVendor: criblcontrolplanesdkgo.Pointer("cribl"),
            EventType: criblcontrolplanesdkgo.Pointer(""),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsSentinelOneAiSiem{},
        },
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
        operations.CreateOutputOutputServiceNow{
            ID: "servicenow-output",
            Type: operations.CreateOutputTypeServiceNowServiceNow,
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
            AuthTokenName: criblcontrolplanesdkgo.Pointer("lightstep-access-token"),
            OtlpVersion: components.OtlpVersionOptions1OneDot3Dot1,
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](2048),
            Protocol: components.ProtocolOptionsHTTP,
            Compress: components.CompressionOptions4Gzip.ToPointer(),
            HTTPCompress: components.CompressionOptions5Gzip.ToPointer(),
            HTTPTracesEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPMetricsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPLogsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](30),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("carefully almost per midst fleck daintily gruesome once bleach"),
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
                    HTTPStatus: 781.29,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
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
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsServiceNow{},
        },
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
        operations.CreateOutputOutputSignalfx{
            ID: "signalfx-output",
            Type: operations.CreateOutputTypeSignalfxSignalfx,
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
            Realm: "us0",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 7111.85,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("cautious fairly needily guzzle"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsSignalfx{},
        },
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
        operations.CreateOutputOutputSnmp{
            ID: "snmp-output",
            Type: operations.CreateOutputTypeSnmpSnmp,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Hosts: []operations.CreateOutputHostSnmp{
                operations.CreateOutputHostSnmp{
                    Host: "192.168.1.1",
                    Port: 161,
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](0),
            Description: criblcontrolplanesdkgo.Pointer("um but pfft mmm an qua"),
        },
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
        operations.CreateOutputOutputSns{
            ID: "sns-output",
            Type: operations.CreateOutputTypeSnsSns,
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
            TopicArn: "arn:aws:sns:us-east-1:123456789012:my-topic",
            MessageGroupID: "my-message-group",
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](7880.99),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: operations.CreateOutputSignatureVersionSnsV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("switchboard smuggle decision and whose truthfully clueless"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsSns{},
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
        operations.CreateOutputOutputSplunk{
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
            },
            Host: "localhost",
            Port: 9997,
            NestedFields: components.NestedFieldSerializationOptionsNone.ToPointer(),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
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
            EnableMultiMetrics: criblcontrolplanesdkgo.Pointer(false),
            EnableACK: criblcontrolplanesdkgo.Pointer(true),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(false),
            MaxS2Sversion: components.MaxS2SVersionOptionsV3.ToPointer(),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("arrange authentic phew scramble tenderly highlight pleasure lest"),
            MaxFailedHealthChecks: criblcontrolplanesdkgo.Pointer[float64](1),
            Compress: components.CompressionOptionsDisabled.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsSplunk{},
            AuthToken: criblcontrolplanesdkgo.Pointer(""),
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
        operations.CreateOutputOutputSplunkHec{
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
                "<value 2>",
            },
            LoadBalanced: criblcontrolplanesdkgo.Pointer(true),
            NextQueue: criblcontrolplanesdkgo.Pointer("indexQueue"),
            TCPRouting: criblcontrolplanesdkgo.Pointer("nowhere"),
            TLS: &components.TLSSettingsClientSideType1{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
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
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            EnableMultiMetrics: criblcontrolplanesdkgo.Pointer(false),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 1179.71,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("stuff fluctuate grim meanwhile"),
            URL: criblcontrolplanesdkgo.Pointer("http://localhost:8088/services/collector/event"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Urls: []operations.CreateOutputURLSplunkHec{
                operations.CreateOutputURLSplunkHec{
                    URL: "http://localhost:8088/services/collector/event",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1),
                    TemplateURL: criblcontrolplanesdkgo.Pointer("https://fat-packaging.net/"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](600),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](300),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsSplunkHec{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://simple-circumference.org/"),
        },
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
        operations.CreateOutputOutputSplunkLb{
            ID: "splunk-lb-output",
            Type: operations.CreateOutputTypeSplunkLbSplunkLb,
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
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](600),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](0),
            NestedFields: components.NestedFieldSerializationOptionsNone.ToPointer(),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
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
            EnableMultiMetrics: criblcontrolplanesdkgo.Pointer(false),
            EnableACK: criblcontrolplanesdkgo.Pointer(true),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(false),
            MaxS2Sversion: components.MaxS2SVersionOptionsV3.ToPointer(),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            IndexerDiscovery: criblcontrolplanesdkgo.Pointer(false),
            SenderUnhealthyTimeAllowance: criblcontrolplanesdkgo.Pointer[float64](100),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("besides spice questionable how"),
            MaxFailedHealthChecks: criblcontrolplanesdkgo.Pointer[float64](1),
            Compress: components.CompressionOptionsDisabled.ToPointer(),
            IndexerDiscoveryConfigs: &operations.CreateOutputIndexerDiscoveryConfigs{
                Site: "default",
                MasterURI: "https://slimy-supplier.info/",
                RefreshIntervalSec: 300,
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                AuthTokens: []operations.CreateOutputAuthToken{
                    operations.CreateOutputAuthToken{
                        AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
                        AuthToken: criblcontrolplanesdkgo.Pointer(""),
                        TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    },
                },
                AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
                AuthToken: criblcontrolplanesdkgo.Pointer(""),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "localhost",
                    Port: 9997,
                    TLS: components.TLSOptionsHostsItemsInherit.ToPointer(),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1),
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsSplunkLb{},
            AuthToken: criblcontrolplanesdkgo.Pointer(""),
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
        operations.CreateOutputOutputSqs{
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
            },
            QueueName: "my-queue",
            QueueType: operations.CreateOutputQueueTypeStandard,
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            MessageGroupID: criblcontrolplanesdkgo.Pointer("cribl"),
            CreateQueue: criblcontrolplanesdkgo.Pointer(true),
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
            MaxQueueSize: criblcontrolplanesdkgo.Pointer[float64](100),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](256),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            MaxInProgress: criblcontrolplanesdkgo.Pointer[float64](10),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("abaft consequently upon shinny yowza why slake"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsSqs{},
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
        operations.CreateOutputOutputStatsd{
            ID: "statsd-output",
            Type: operations.CreateOutputTypeStatsdStatsd,
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
            Mtu: criblcontrolplanesdkgo.Pointer[float64](512),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](0),
            Description: criblcontrolplanesdkgo.Pointer("browse longingly tame regarding pendant playfully orchestrate"),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsStatsd{},
        },
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
        operations.CreateOutputOutputStatsdExt{
            ID: "statsd-ext-output",
            Type: operations.CreateOutputTypeStatsdExtStatsdExt,
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
            Mtu: criblcontrolplanesdkgo.Pointer[float64](512),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](0),
            Description: criblcontrolplanesdkgo.Pointer("after with wafer collaboration"),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsStatsdExt{},
        },
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
        operations.CreateOutputOutputSumoLogic{
            ID: "sumo-logic-output",
            Type: operations.CreateOutputTypeSumoLogicSumoLogic,
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
            URL: "https://endpoint1.collection.us2.sumologic.com",
            CustomSource: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomCategory: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: operations.CreateOutputDataFormatSumoLogicJSON.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1024),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 8295.69,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](7368.99),
            Description: criblcontrolplanesdkgo.Pointer("jovially mmm but zebra phew deform"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsSumoLogic{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://profuse-decryption.biz"),
        },
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
        operations.CreateOutputOutputSyslog{
            ID: "syslog-output",
            Type: operations.CreateOutputTypeSyslogSyslog,
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
            Protocol: operations.CreateOutputProtocolSyslogTCP.ToPointer(),
            Facility: operations.CreateOutputFacilityOne.ToPointer(),
            Severity: operations.CreateOutputSeveritySyslogNotice.ToPointer(),
            AppName: criblcontrolplanesdkgo.Pointer("Cribl"),
            MessageFormat: operations.CreateOutputMessageFormatRfc3164.ToPointer(),
            TimestampFormat: operations.CreateOutputTimestampFormatSyslog.ToPointer(),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
            OctetCountFraming: criblcontrolplanesdkgo.Pointer(true),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("drat consequently restfully distorted phooey overconfidently"),
            LoadBalanced: criblcontrolplanesdkgo.Pointer(true),
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](514),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "separate-blossom.biz",
                    Port: 8854.36,
                    TLS: components.TLSOptionsHostsItemsInherit.ToPointer(),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1),
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](600),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](0),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
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
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            MaxRecordSize: criblcontrolplanesdkgo.Pointer[float64](1500),
            UDPDNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](0),
            EnableIPSpoofing: criblcontrolplanesdkgo.Pointer(false),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsSyslog{},
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
        operations.CreateOutputOutputTcpjson{
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
            LoadBalanced: criblcontrolplanesdkgo.Pointer(true),
            Compression: components.CompressionOptions1Gzip.ToPointer(),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(false),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
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
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            TokenTTLMinutes: criblcontrolplanesdkgo.Pointer[float64](60),
            SendHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("hourly about into"),
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](10090),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "chilly-exterior.net",
                    Port: 6458.17,
                    TLS: components.TLSOptionsHostsItemsInherit.ToPointer(),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1),
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](600),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](0),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsTcpjson{},
            AuthToken: criblcontrolplanesdkgo.Pointer(""),
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
        operations.CreateOutputOutputWavefront{
            ID: "wavefront-output",
            Type: operations.CreateOutputTypeWavefrontWavefront,
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
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 6928.76,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("excepting fooey solace spear"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsWavefront{},
        },
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
        operations.CreateOutputOutputWebhook{
            ID: "webhook-output",
            Type: operations.CreateOutputTypeWebhookWebhook,
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
            Method: components.MethodOptionsPost.ToPointer(),
            Format: operations.CreateOutputFormatWebhookNdjson.ToPointer(),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 4167.91,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: operations.CreateOutputAuthenticationTypeWebhookNone.ToPointer(),
            TLS: &components.TLSSettingsClientSideType1{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](1305.29),
            LoadBalanced: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("breakable hospitable override hmph presell"),
            CustomSourceExpression: criblcontrolplanesdkgo.Pointer("__httpOut"),
            CustomDropWhenNull: criblcontrolplanesdkgo.Pointer(false),
            CustomEventDelimiter: criblcontrolplanesdkgo.Pointer("\\n"),
            CustomContentType: criblcontrolplanesdkgo.Pointer("application/x-ndjson"),
            CustomPayloadExpression: criblcontrolplanesdkgo.Pointer("`${events}`"),
            AdvancedContentType: criblcontrolplanesdkgo.Pointer("application/json"),
            FormatEventCode: criblcontrolplanesdkgo.Pointer("<value>"),
            FormatPayloadCode: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsWebhook{},
            Username: criblcontrolplanesdkgo.Pointer("Janiya_Dietrich"),
            Password: criblcontrolplanesdkgo.Pointer("QSmMfe0g6BQ_7dP"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://neglected-suspension.net/"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("`Bearer ${token}`"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](3600),
            OauthParams: []operations.CreateOutputOauthParam{
                operations.CreateOutputOauthParam{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            OauthHeaders: []operations.CreateOutputOauthHeader{
                operations.CreateOutputOauthHeader{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            URL: criblcontrolplanesdkgo.Pointer("https://example.com/webhook"),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Urls: []operations.CreateOutputURLWebhook{
                operations.CreateOutputURLWebhook{
                    URL: "https://second-hand-apparatus.info",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1),
                    TemplateURL: criblcontrolplanesdkgo.Pointer("https://unfit-grandson.biz/"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](600),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](300),
            TemplateLoginURL: criblcontrolplanesdkgo.Pointer("https://scented-chiffonier.net"),
            TemplateSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://grown-video.name/"),
        },
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
        operations.CreateOutputOutputWizHec{
            ID: "wiz-hec-output",
            Type: operations.CreateOutputTypeWizHecWizHec,
            Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            SystemFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            LoadBalanced: true,
            NextQueue: criblcontrolplanesdkgo.Pointer("indexQueue"),
            TCPRouting: criblcontrolplanesdkgo.Pointer("nowhere"),
            TLS: &components.TLSSettingsClientSideType1{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
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
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            EnableMultiMetrics: false,
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 7295.73,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            WizConnectorID: "00000000-0000-0000-0000-000000000000",
            WizEnvironment: "test",
            DataCenter: "us1",
            WizSourcetype: "placeholder",
            Description: criblcontrolplanesdkgo.Pointer("hourly about into"),
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
            PqControls: &operations.CreateOutputPqControlsWizHec{},
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
        operations.CreateOutputOutputXsiam{
            ID: "xsiam-output",
            Type: operations.CreateOutputTypeXsiamXsiam,
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
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](9500),
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
            SafeHeaders: []string{
                "<value 1>",
            },
            AuthType: operations.CreateOutputAuthenticationMethodXsiamToken.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8977.97,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            ThrottleRateReqPerSec: criblcontrolplanesdkgo.Pointer[int64](400),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](1079.49),
            Description: criblcontrolplanesdkgo.Pointer("norm ragged glaring"),
            URL: criblcontrolplanesdkgo.Pointer("http://localhost:8088/logs/v1/event"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Urls: []operations.CreateOutputURLXsiam{
                operations.CreateOutputURLXsiam{
                    URL: "https://kooky-wilderness.biz",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](600),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](300),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &operations.CreateOutputPqControlsXsiam{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://unkempt-travel.net/"),
        },
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
            },
            ContainerName: "my-container",
            CreateContainer: criblcontrolplanesdkgo.Pointer(false),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](1),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("C.Time.strftime(_time ? _time : Date.now()/1000, '%Y/%m/%d')"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("`CriblOut`"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("`.${C.env[\"CRIBL_WORKER_ID\"]}.${__format}${__compression === \"gzip\" ? \".gz\" : \"\"}`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](32),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            HeaderLine: criblcontrolplanesdkgo.Pointer(""),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](64),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](6434.77),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](857.42),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2924.72),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](4562.74),
            },
            AuthType: components.AuthenticationMethodOptionsManual.ToPointer(),
            StorageClass: components.BlobAccessTierInferred.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("practical adumbrate by"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("1MB"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
            Description: criblcontrolplanesdkgo.Pointer("oof pace once"),
            ClientSecret: criblcontrolplanesdkgo.Pointer("client-secret"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Certificate: &components.Certificate{
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            Format: components.DataFormatOptionsJSON.ToPointer(),
            Compress: components.CompressionOptions2Gzip,
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("1MB"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
            IsMappingObj: criblcontrolplanesdkgo.Pointer(false),
            MappingObj: criblcontrolplanesdkgo.Pointer("<value>"),
            MappingRef: criblcontrolplanesdkgo.Pointer("<value>"),
            IngestURL: criblcontrolplanesdkgo.Pointer("https://competent-nephew.name"),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            StagePath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/staging"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("`.${C.env[\"CRIBL_WORKER_ID\"]}.${__format}${__compression === \"gzip\" ? \".gz\" : \"\"}`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](32),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](1),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](6434.77),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](857.42),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2924.72),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](4562.74),
            },
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
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
            ReportLevel: components.ReportLevelFailuresOnly.ToPointer(),
            ReportMethod: components.ReportMethodQueue.ToPointer(),
            AdditionalProperties: []components.AdditionalProperty{
                components.AdditionalProperty{
                    Key: "<key>",
                    Value: "<value>",
                },
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 4119.46,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputAzureDataExplorerPqControls{},
            TemplateClusterURL: criblcontrolplanesdkgo.Pointer("https://exotic-championship.com/"),
            TemplateDatabase: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateTable: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateTenantID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateClientID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateScope: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateClientSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateFormat: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateIngestURL: criblcontrolplanesdkgo.Pointer("https://critical-range.net"),
        },
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
                "<value 2>",
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
            Format: components.RecordDataFormatOptionsJSON.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](768),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](1000),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                Password: criblcontrolplanesdkgo.Pointer("l4UiMzAymhgS2bu"),
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
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("over excluding knowingly energetically incidentally given"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
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
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1024),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
            APIURL: criblcontrolplanesdkgo.Pointer(".ods.opinsights.azure.com"),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 6460.17,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.OutputAzureLogsAuthenticationMethodManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("critical under even beyond condense diligently"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
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
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            APIVersion: criblcontrolplanesdkgo.Pointer("v1alpha"),
            AuthenticationMethod: components.OutputChronicleAuthenticationMethodServiceAccount.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 4457.03,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            Region: "us",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1024),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](90),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](1999.04),
            IngestionMethod: criblcontrolplanesdkgo.Pointer("ImportLogs"),
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
            Description: criblcontrolplanesdkgo.Pointer("pleasant waist cease aha because switchboard phew keenly arrogantly husky"),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            ServiceAccountCredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
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
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 5472.53,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            DumpFormatErrorsToDisk: criblcontrolplanesdkgo.Pointer(false),
            StatsDestination: &components.StatsDestination{
                URL: criblcontrolplanesdkgo.Pointer("https://proud-blowgun.biz/"),
                Database: criblcontrolplanesdkgo.Pointer("<value>"),
                TableName: criblcontrolplanesdkgo.Pointer("<value>"),
                AuthType: criblcontrolplanesdkgo.Pointer("<value>"),
                Username: criblcontrolplanesdkgo.Pointer("Jaclyn.Champlin36"),
                SQLUsername: criblcontrolplanesdkgo.Pointer("<value>"),
                Password: criblcontrolplanesdkgo.Pointer("fhOYWNIwx4mmq1W"),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("forenenst capitalise around"),
            Username: criblcontrolplanesdkgo.Pointer("Herminia.Botsford"),
            Password: criblcontrolplanesdkgo.Pointer("uO_AIsmUCJ5I1i2"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            SQLUsername: criblcontrolplanesdkgo.Pointer("<value>"),
            WaitForAsyncInserts: criblcontrolplanesdkgo.Pointer(true),
            ExcludeMappingFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
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
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputClickHousePqControls{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://quarrelsome-slime.name"),
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
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Endpoint: "https://account-id.r2.cloudflarestorage.com",
            Bucket: "my-bucket",
            AwsAuthenticationMethod: components.OutputCloudflareR2AuthenticationMethodAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: "<value>",
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions5V4.ToPointer(),
            ObjectACL: "<value>",
            StorageClass: components.StorageClassOptions2ReducedRedundancy.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionOptionsAes256.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("C.Time.strftime(_time ? _time : Date.now()/1000, '%Y/%m/%d')"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("`CriblOut`"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("`.${C.env[\"CRIBL_WORKER_ID\"]}.${__format}${__compression === \"gzip\" ? \".gz\" : \"\"}`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](32),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            HeaderLine: criblcontrolplanesdkgo.Pointer(""),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](64),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](6434.77),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](857.42),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2924.72),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](4562.74),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](4),
            Description: criblcontrolplanesdkgo.Pointer("whenever beret quirkily certainly cruelty well-documented poorly inject throughout"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("1MB"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            MaxQueueSize: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](1024),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("hubris gray materialise aboard content ick self-assured cleaner imagineer"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
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
                "<value 3>",
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
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            Topic: "logs",
            Ack: components.AcknowledgmentsOptions1Leader.ToPointer(),
            Format: components.RecordDataFormatOptions1JSON.ToPointer(),
            Compression: components.CompressionOptions3Gzip.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](768),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](1000),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType1{
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
                    Disabled: criblcontrolplanesdkgo.Pointer(true),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                    CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                },
                DefaultKeySchemaID: criblcontrolplanesdkgo.Pointer[float64](2571.24),
                DefaultValueSchemaID: criblcontrolplanesdkgo.Pointer[float64](4603.74),
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
                Username: criblcontrolplanesdkgo.Pointer("Beverly93"),
                Password: criblcontrolplanesdkgo.Pointer("iNd6FlpDEbKPe_d"),
                AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslPlain.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(false),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://smooth-conservative.com"),
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
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("until scientific since properly massage and"),
            ProtobufLibraryID: criblcontrolplanesdkgo.Pointer("<id>"),
            ProtobufEncodingID: criblcontrolplanesdkgo.Pointer("<id>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
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
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            LoadBalanced: criblcontrolplanesdkgo.Pointer(true),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            TokenTTLMinutes: criblcontrolplanesdkgo.Pointer[float64](60),
            ExcludeFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
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
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 3673.51,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            AuthTokens: []components.ItemsTypeAuthTokens1{
                components.ItemsTypeAuthTokens1{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("fabricate focalise calmly per outside aside astride known lumpy"),
                },
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("unique filthy irresponsible wherever lawful an settle unlike"),
            URL: criblcontrolplanesdkgo.Pointer("https://tidy-beret.biz"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Urls: []components.ItemsTypeUrls{
                components.ItemsTypeUrls{
                    URL: "https://weird-finer.net",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1),
                    TemplateURL: criblcontrolplanesdkgo.Pointer("https://adolescent-sprinkles.biz/"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](600),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](300),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputCriblHTTPPqControls{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://focused-quinoa.info"),
        },
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
                "<value 3>",
            },
            Bucket: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            StagePath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/staging"),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            ObjectACL: components.ObjectACLOptionsPrivate.ToPointer(),
            StorageClass: components.StorageClassOptionsDeepArchive.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAwsKms.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            BaseFileName: criblcontrolplanesdkgo.Pointer("`CriblOut`"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("`.${C.env[\"CRIBL_WORKER_ID\"]}.${__format}${__compression === \"gzip\" ? \".gz\" : \"\"}`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](64),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            HeaderLine: criblcontrolplanesdkgo.Pointer(""),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](64),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](6434.77),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](857.42),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2924.72),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](4562.74),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](100),
            AwsAuthenticationMethod: components.AwsAuthenticationMethodOptionsAuto.ToPointer(),
            Format: components.FormatOptionsDdss.ToPointer(),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](1),
            Description: criblcontrolplanesdkgo.Pointer("mad mmm bitterly ah overwork mysterious instead miskey"),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv1.ToPointer(),
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
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            AuthTokens: []components.ItemsTypeAuthTokens1{
                components.ItemsTypeAuthTokens1{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("developmental yum tomatillo"),
                },
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            Description: criblcontrolplanesdkgo.Pointer("contradict unless busy for concerned buck grandiose dreamily what solder"),
            URL: criblcontrolplanesdkgo.Pointer("https://0.0.0.0:10200"),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Urls: []components.ItemsTypeUrls{
                components.ItemsTypeUrls{
                    URL: "https://short-term-ectoderm.net/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1),
                    TemplateURL: criblcontrolplanesdkgo.Pointer("https://dull-mainstream.org"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](600),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](300),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputCriblSearchEnginePqControls{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://burdensome-discourse.com/"),
        },
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
            LoadBalanced: criblcontrolplanesdkgo.Pointer(true),
            Compression: components.CompressionOptions1Gzip.ToPointer(),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(false),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            TokenTTLMinutes: criblcontrolplanesdkgo.Pointer[float64](60),
            AuthTokens: []components.ItemsTypeAuthTokens{
                components.ItemsTypeAuthTokens{
                    TokenSecret: "<value>",
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    Description: criblcontrolplanesdkgo.Pointer("unless zowie unbearably throughout oh brr solemnly"),
                },
            },
            ExcludeFields: []string{
                "<value 1>",
                "<value 2>",
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("pace primary nor tomography once"),
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](10090),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "tough-colonialism.info",
                    Port: 10300,
                    TLS: components.TLSOptionsHostsItemsInherit.ToPointer(),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1),
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](600),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](0),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
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
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            URL: "https://ingest.us.crowdstrike.com/api/ingest/hec/connection-id/v1/services/collector",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
            Format: components.RequestFormatOptionsJSON,
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 1023.69,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("yesterday connect whose"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputCrowdstrikeNextGenSiemPqControls{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://quiet-riser.org/"),
        },
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
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            DestPath: criblcontrolplanesdkgo.Pointer(""),
            StagePath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/staging"),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("C.Time.strftime(_time ? _time : Date.now()/1000, '%Y/%m/%d')"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("`CriblOut`"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("`.${C.env[\"CRIBL_WORKER_ID\"]}.${__format}${__compression === \"gzip\" ? \".gz\" : \"\"}`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](32),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            HeaderLine: criblcontrolplanesdkgo.Pointer(""),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](64),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](6434.77),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](857.42),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2924.72),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](4562.74),
            },
            WorkspaceID: "your-workspace-id",
            Scope: "my-scope",
            ClientID: "your-client-id",
            Catalog: "my-catalog",
            Schema: "my-schema",
            EventsVolumeName: "my-volume",
            ClientTextSecret: "your-client-secret",
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](60),
            Description: criblcontrolplanesdkgo.Pointer("aw yowza instead"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("1MB"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(true),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            ContentType: components.SendLogsAsJSON.ToPointer(),
            Message: criblcontrolplanesdkgo.Pointer("<value>"),
            Source: criblcontrolplanesdkgo.Pointer("<value>"),
            Host: criblcontrolplanesdkgo.Pointer("weighty-release.net"),
            Service: criblcontrolplanesdkgo.Pointer("<value>"),
            Tags: []string{
                "<value 1>",
            },
            BatchByTags: criblcontrolplanesdkgo.Pointer(true),
            AllowAPIKeyFromEvents: criblcontrolplanesdkgo.Pointer(false),
            Severity: components.OutputDatadogSeverityInfo.ToPointer(),
            Site: components.DatadogSiteUs.ToPointer(),
            SendCountersAsCount: criblcontrolplanesdkgo.Pointer(false),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 3412.47,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Manual.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](230.7),
            Description: criblcontrolplanesdkgo.Pointer("yippee what during put that pish uselessly although wherever declaration"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://stingy-guide.name/"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
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
            },
            MessageField: criblcontrolplanesdkgo.Pointer("<value>"),
            ExcludeFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ServerHostField: criblcontrolplanesdkgo.Pointer("<value>"),
            TimestampField: criblcontrolplanesdkgo.Pointer("<value>"),
            DefaultSeverity: components.OutputDatasetSeverityInfo.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5952.63,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            Site: components.DataSetSiteUs.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Manual.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](1088.03),
            Description: criblcontrolplanesdkgo.Pointer("apropos gosh overvalue than handful fully leading yuck portray"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://blue-forage.info"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputDatasetPqControls{},
            APIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateCustomURL: criblcontrolplanesdkgo.Pointer("https://sleepy-cassava.org/"),
        },
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
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            TimeWindow: criblcontrolplanesdkgo.Pointer("10m"),
            MaxDataSize: criblcontrolplanesdkgo.Pointer("1GB"),
            MaxDataTime: criblcontrolplanesdkgo.Pointer("24h"),
            Compress: components.CompressionOptionsPersistenceGzip.ToPointer(),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("along between embed into declaration boastfully though verbally platter"),
        },
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
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer(""),
            ObjectACL: components.ObjectACLOptionsPrivate.ToPointer(),
            StorageClass: components.StorageClassOptionsReducedRedundancy.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAwsKms.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("`CriblOut`"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("`.${C.env[\"CRIBL_WORKER_ID\"]}.${__format}${__compression === \"gzip\" ? \".gz\" : \"\"}`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](32),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            HeaderLine: criblcontrolplanesdkgo.Pointer(""),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](64),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](6434.77),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](857.42),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2924.72),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](4562.74),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](4),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](100),
            PartitioningFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Description: criblcontrolplanesdkgo.Pointer("whenever mash ugh reasoning premier healthily"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("1MB"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Method: components.MethodOptionsPost.ToPointer(),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 8416.09,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.OutputDynatraceHTTPAuthenticationTypeToken.ToPointer(),
            Format: components.OutputDynatraceHTTPFormatJSONArray,
            Endpoint: components.EndpointCloud,
            TelemetryType: components.TelemetryTypeLogs,
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](6599.51),
            Description: criblcontrolplanesdkgo.Pointer("uh-huh retract excluding"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputDynatraceHTTPPqControls{},
            Token: criblcontrolplanesdkgo.Pointer("your-api-key"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            EnvironmentID: criblcontrolplanesdkgo.Pointer("<id>"),
            ActiveGateDomain: criblcontrolplanesdkgo.Pointer("<value>"),
            URL: criblcontrolplanesdkgo.Pointer("https://yummy-resource.name/"),
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://timely-mountain.info"),
        },
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
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Protocol: components.OutputDynatraceOtlpProtocolHTTP,
            Endpoint: "https://your-environment.live.dynatrace.com/api/v2/otlp",
            OtlpVersion: components.OtlpVersionOptions1OneDot3Dot1,
            Compress: components.CompressionOptions4Gzip.ToPointer(),
            HTTPCompress: components.CompressionOptions5Gzip.ToPointer(),
            HTTPTracesEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPMetricsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPLogsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](2048),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](30),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            EndpointType: components.EndpointTypeSaas,
            TokenSecret: "your-token-secret",
            AuthTokenName: criblcontrolplanesdkgo.Pointer("Authorization"),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("tinderbox now until than appreciate nervously aching seafood"),
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
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 7970.31,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
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
                "<value 2>",
            },
            LoadBalanced: criblcontrolplanesdkgo.Pointer(true),
            Index: "logs",
            DocType: criblcontrolplanesdkgo.Pointer("<value>"),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
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
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 5469.02,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
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
                Username: criblcontrolplanesdkgo.Pointer("Timmothy.Morar"),
                Password: criblcontrolplanesdkgo.Pointer("lCK4OFBNamsTS5U"),
                AuthType: components.AuthenticationMethodOptionsAuthManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                ManualAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            ElasticVersion: components.ElasticVersionAuto.ToPointer(),
            ElasticPipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            IncludeDocID: criblcontrolplanesdkgo.Pointer(false),
            WriteAction: components.WriteActionCreate.ToPointer(),
            RetryPartialErrors: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("save along larva quantify forearm drat physically"),
            URL: criblcontrolplanesdkgo.Pointer("https://alert-fort.info/"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Urls: []components.OutputElasticURL{
                components.OutputElasticURL{
                    URL: "https://first-bourgeoisie.biz",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1),
                    TemplateURL: criblcontrolplanesdkgo.Pointer("https://accomplished-cuckoo.net"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](600),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](300),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputElasticPqControls{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://thin-other.biz"),
        },
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
            },
            URL: "my-cloud-id",
            Index: "logs",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
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
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ExtraParams: []components.ItemsTypeSaslSaslExtensions{
                components.ItemsTypeSaslSaslExtensions{
                    Name: "<value>",
                    Value: "<value>",
                },
            },
            Auth: &components.AuthType{
                Disabled: true,
                Username: criblcontrolplanesdkgo.Pointer("Timmothy.Morar"),
                Password: criblcontrolplanesdkgo.Pointer("lCK4OFBNamsTS5U"),
                AuthType: components.AuthenticationMethodOptionsAuthManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                ManualAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            ElasticPipeline: criblcontrolplanesdkgo.Pointer("<value>"),
            IncludeDocID: criblcontrolplanesdkgo.Pointer(true),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 6701.64,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("gadzooks quaintly jet unearth"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
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
            ObjectACL: components.ObjectACLOptions1Private.ToPointer(),
            StorageClass: components.StorageClassOptions1Standard.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](6434.77),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](857.42),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2924.72),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](4562.74),
            },
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](10),
            EncodedConfiguration: criblcontrolplanesdkgo.Pointer("<value>"),
            CollectorInstanceID: "11112222-3333-4444-5555-666677778888",
            SiteName: criblcontrolplanesdkgo.Pointer("<value>"),
            SiteID: criblcontrolplanesdkgo.Pointer("<id>"),
            TimezoneOffset: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("brr despite even er woefully regarding"),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
                "<value 2>",
            },
            DestPath: "/var/log/output",
            StagePath: criblcontrolplanesdkgo.Pointer("<value>"),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("C.Time.strftime(_time ? _time : Date.now()/1000, '%Y/%m/%d')"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("`CriblOut`"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("`.${C.env[\"CRIBL_WORKER_ID\"]}.${__format}${__compression === \"gzip\" ? \".gz\" : \"\"}`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](32),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            HeaderLine: criblcontrolplanesdkgo.Pointer(""),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](64),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](6434.77),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](857.42),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2924.72),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](4562.74),
            },
            Description: criblcontrolplanesdkgo.Pointer("luck ah ski"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("1MB"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(true),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            APIVersion: components.OutputGoogleChronicleAPIVersionV1.ToPointer(),
            AuthenticationMethod: components.OutputGoogleChronicleAuthenticationMethodServiceAccount.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 1271.28,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            LogFormatType: components.SendEventsAsUnstructured,
            Region: criblcontrolplanesdkgo.Pointer("us"),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1024),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](90),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](304.57),
            Description: criblcontrolplanesdkgo.Pointer("fooey safely ah tuba nervous sans whoever at invite"),
            ExtraLogTypes: []components.ExtraLogType{
                components.ExtraLogType{
                    LogType: "<value>",
                    Description: criblcontrolplanesdkgo.Pointer("busy for concerned buck grandiose dreamily what solder"),
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
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
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
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            LogLocationType: components.LogLocationTypeProject,
            LogNameExpression: "my-log",
            SanitizeLogNames: criblcontrolplanesdkgo.Pointer(false),
            PayloadFormat: components.PayloadFormatText.ToPointer(),
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
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            ThrottleRateReqPerSec: criblcontrolplanesdkgo.Pointer[int64](829528),
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
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](527.61),
            Description: criblcontrolplanesdkgo.Pointer("meanwhile boo beneficial below positively huzzah fully anxiously"),
            LogLocationExpression: "my-project",
            PayloadExpression: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
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
            },
            Bucket: "my-bucket",
            Region: "us-east1",
            Endpoint: "https://storage.googleapis.com",
            SignatureVersion: components.SignatureVersionOptions4V4.ToPointer(),
            AwsAuthenticationMethod: components.OutputGoogleCloudStorageAuthenticationMethodManual.ToPointer(),
            StagePath: "/tmp/staging",
            DestPath: criblcontrolplanesdkgo.Pointer(""),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            ObjectACL: components.ObjectACLOptions1Private.ToPointer(),
            StorageClass: components.StorageClassOptions1Archive.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("C.Time.strftime(_time ? _time : Date.now()/1000, '%Y/%m/%d')"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("`CriblOut`"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("`.${C.env[\"CRIBL_WORKER_ID\"]}.${__format}${__compression === \"gzip\" ? \".gz\" : \"\"}`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](32),
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            HeaderLine: criblcontrolplanesdkgo.Pointer(""),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](64),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](6434.77),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](857.42),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2924.72),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](4562.74),
            },
            Description: criblcontrolplanesdkgo.Pointer("ew sarong cumbersome"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("1MB"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            TopicName: "my-topic",
            CreateTopic: criblcontrolplanesdkgo.Pointer(false),
            OrderedDelivery: criblcontrolplanesdkgo.Pointer(false),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            GoogleAuthMethod: components.GoogleAuthenticationMethodOptionsManual.ToPointer(),
            ServiceAccountCredentials: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            BatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            BatchTimeout: criblcontrolplanesdkgo.Pointer[float64](100),
            MaxQueueSize: criblcontrolplanesdkgo.Pointer[float64](100),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](256),
            FlushPeriod: criblcontrolplanesdkgo.Pointer[float64](1),
            MaxInProgress: criblcontrolplanesdkgo.Pointer[float64](10),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("although worth reckon"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
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
                },
                LokiURL: "https://logs-prod-us-central1.grafana.net",
                PrometheusURL: criblcontrolplanesdkgo.Pointer("https://jubilant-disclosure.org"),
                Message: criblcontrolplanesdkgo.Pointer("<value>"),
                MessageFormat: components.MessageFormatOptionsProtobuf.ToPointer(),
                Labels: []components.ItemsTypeLabels{
                    components.ItemsTypeLabels{
                        Name: "",
                        Value: "<value>",
                    },
                },
                MetricRenameExpr: criblcontrolplanesdkgo.Pointer("name.replace(/[^a-zA-Z0-9_]/g, '_')"),
                PrometheusAuth: &components.PrometheusAuthType{
                    AuthType: components.AuthenticationTypeOptionsPrometheusAuth1Basic.ToPointer(),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Username: criblcontrolplanesdkgo.Pointer("Flo.Koelpin"),
                    Password: criblcontrolplanesdkgo.Pointer("WalxTm654qjkLiW"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                LokiAuth: &components.PrometheusAuthType{
                    AuthType: components.AuthenticationTypeOptionsPrometheusAuth1Basic.ToPointer(),
                    Token: criblcontrolplanesdkgo.Pointer("<value>"),
                    TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    Username: criblcontrolplanesdkgo.Pointer("Alda_MacGyver88"),
                    Password: criblcontrolplanesdkgo.Pointer("s0xBReXekhrDqCh"),
                    CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                },
                Concurrency: criblcontrolplanesdkgo.Pointer[float64](1),
                MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
                MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
                FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](15),
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
                        HTTPStatus: 8961.65,
                        InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                        BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                        MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                    },
                },
                TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                    TimeoutRetry: false,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
                ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
                OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
                Description: criblcontrolplanesdkgo.Pointer("unlike crumble circulate materialise"),
                Compress: criblcontrolplanesdkgo.Pointer(true),
                PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
                PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
                PqMode: components.ModeOptionsError.ToPointer(),
                PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
                PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
                PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
                PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
                PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
                PqCompress: components.CompressionOptionsPqNone.ToPointer(),
                PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
                PqControls: &components.OutputGrafanaCloudPqControls1{},
                TemplateLokiURL: criblcontrolplanesdkgo.Pointer("https://firsthand-straw.name"),
                TemplatePrometheusURL: criblcontrolplanesdkgo.Pointer("https://grandiose-technologist.com"),
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
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Protocol: components.DestinationProtocolOptionsTCP,
            Host: "localhost",
            Port: 2003,
            Mtu: criblcontrolplanesdkgo.Pointer[float64](512),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](0),
            Description: criblcontrolplanesdkgo.Pointer("whopping fine whose meanwhile finally off plump"),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
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
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Dataset: "my-dataset",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 864.6,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("scope castanet pfft creature"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
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
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            URL: "https://cloud.us.humio.com/api/v1/ingest/hec",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
            Format: components.RequestFormatOptionsJSON,
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 8378.07,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("nippy impressionable ouch concerning upside-down after"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputHumioHecPqControls{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://heartfelt-deed.net/"),
        },
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
                "<value 2>",
                "<value 3>",
            },
            URL: "http://localhost:8086",
            UseV2API: criblcontrolplanesdkgo.Pointer(false),
            TimestampPrecision: components.TimestampPrecisionMs.ToPointer(),
            DynamicValueFieldName: criblcontrolplanesdkgo.Pointer(true),
            ValueFieldName: criblcontrolplanesdkgo.Pointer("value"),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 2224.46,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.OutputInfluxdbAuthenticationTypeNone.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("once till aw forgery"),
            Database: criblcontrolplanesdkgo.Pointer("mydb"),
            Bucket: criblcontrolplanesdkgo.Pointer("<value>"),
            Org: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputInfluxdbPqControls{},
            Username: criblcontrolplanesdkgo.Pointer("Hector.Kshlerin29"),
            Password: criblcontrolplanesdkgo.Pointer("uj1bSnZx1GRBoIy"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://warmhearted-premise.info/"),
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
            Format: components.RecordDataFormatOptions1JSON.ToPointer(),
            Compression: components.CompressionOptions3Gzip.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](768),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](1000),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType1{
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
                    Disabled: criblcontrolplanesdkgo.Pointer(true),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                    CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                },
                DefaultKeySchemaID: criblcontrolplanesdkgo.Pointer[float64](2571.24),
                DefaultValueSchemaID: criblcontrolplanesdkgo.Pointer[float64](4603.74),
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
                Username: criblcontrolplanesdkgo.Pointer("Beverly93"),
                Password: criblcontrolplanesdkgo.Pointer("iNd6FlpDEbKPe_d"),
                AuthType: components.AuthenticationMethodOptionsSaslManual.ToPointer(),
                CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                Mechanism: components.SaslMechanismOptionsSaslPlain.ToPointer(),
                KeytabLocation: criblcontrolplanesdkgo.Pointer("<value>"),
                Principal: criblcontrolplanesdkgo.Pointer("<value>"),
                BrokerServiceClass: criblcontrolplanesdkgo.Pointer("<value>"),
                OauthEnabled: criblcontrolplanesdkgo.Pointer(false),
                TokenURL: criblcontrolplanesdkgo.Pointer("https://smooth-conservative.com"),
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
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("forenenst red which solicit where rejigger uncommon scoop"),
            ProtobufLibraryID: criblcontrolplanesdkgo.Pointer("<id>"),
            ProtobufEncodingID: criblcontrolplanesdkgo.Pointer("<id>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
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
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](1024),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            Compression: components.OutputKinesisCompressionGzip.ToPointer(),
            UseListShards: criblcontrolplanesdkgo.Pointer(false),
            AsNdjson: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("obediently or pricey why hose detective scout"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxEventsPerFlush: criblcontrolplanesdkgo.Pointer[float64](500),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
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
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            URL: "http://localhost:3100/loki/api/v1/push",
            Message: criblcontrolplanesdkgo.Pointer("<value>"),
            MessageFormat: components.MessageFormatOptionsProtobuf.ToPointer(),
            Labels: []components.ItemsTypeLabels{
                components.ItemsTypeLabels{
                    Name: "",
                    Value: "<value>",
                },
            },
            AuthType: components.AuthenticationTypeOptionsPrometheusAuth1None.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](1),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](15),
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
                    HTTPStatus: 259.76,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            EnableDynamicHeaders: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](3553.75),
            Description: criblcontrolplanesdkgo.Pointer("perfectly youthfully often usefully midst once shameless out for"),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Username: criblcontrolplanesdkgo.Pointer("Virginie.Kohler"),
            Password: criblcontrolplanesdkgo.Pointer("oxajD3Wd8pLSEnh"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
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
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Topic: "logs",
            Ack: components.AcknowledgmentsOptionsLeader.ToPointer(),
            Format: components.RecordDataFormatOptionsJSON.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](768),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](1000),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            RequestTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxBackOff: criblcontrolplanesdkgo.Pointer[float64](30000),
            InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](300),
            BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
            AuthenticationTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            ReauthenticationThreshold: criblcontrolplanesdkgo.Pointer[float64](10000),
            Sasl: &components.OutputMicrosoftFabricAuthentication{
                Disabled: false,
                Mechanism: components.SaslMechanismOptionsSasl1Plain.ToPointer(),
                Username: criblcontrolplanesdkgo.Pointer("$ConnectionString"),
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
                Disabled: false,
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            BootstrapServer: "myeventstream.servicebus.windows.net:9093",
            Description: criblcontrolplanesdkgo.Pointer("while so squirm surprise circa indeed simple angle daddy"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
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
                "<value 2>",
                "<value 3>",
            },
            Endpoint: "http://localhost:9000",
            Bucket: "my-bucket",
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("<value>"),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.SignatureVersionOptions5V4.ToPointer(),
            ObjectACL: components.ObjectACLOptionsPrivate.ToPointer(),
            StorageClass: components.StorageClassOptions2ReducedRedundancy.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionOptionsAes256.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("C.Time.strftime(_time ? _time : Date.now()/1000, '%Y/%m/%d')"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("`CriblOut`"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("`.${C.env[\"CRIBL_WORKER_ID\"]}.${__format}${__compression === \"gzip\" ? \".gz\" : \"\"}`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](32),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            HeaderLine: criblcontrolplanesdkgo.Pointer(""),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](64),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](6434.77),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](857.42),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2924.72),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](4562.74),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](4),
            Description: criblcontrolplanesdkgo.Pointer("gosh footrest complication"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("1MB"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(true),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
            Ack: components.AcknowledgmentsOptions1Leader.ToPointer(),
            Format: components.RecordDataFormatOptions1JSON.ToPointer(),
            Compression: components.CompressionOptions3Gzip.ToPointer(),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](768),
            FlushEventCount: criblcontrolplanesdkgo.Pointer[float64](1000),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            KafkaSchemaRegistry: &components.KafkaSchemaRegistryAuthenticationType1{
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
                    Disabled: criblcontrolplanesdkgo.Pointer(true),
                    RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                    CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                    Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                    MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                },
                DefaultKeySchemaID: criblcontrolplanesdkgo.Pointer[float64](2571.24),
                DefaultValueSchemaID: criblcontrolplanesdkgo.Pointer[float64](4603.74),
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
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("kick adrenalin above airbus oil or smart babushka ethyl dash"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            ProtobufLibraryID: criblcontrolplanesdkgo.Pointer("<id>"),
            ProtobufEncodingID: criblcontrolplanesdkgo.Pointer("<id>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
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
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            Hosts: []components.OutputNetflowHost{
                components.OutputNetflowHost{
                    Host: "localhost",
                    Port: 2055,
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](0),
            EnableIPSpoofing: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("whether unlike once cap schlep promptly submitter plus fatal than"),
            MaxRecordSize: criblcontrolplanesdkgo.Pointer[float64](1500),
        },
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
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Region: components.RegionOptionsUs.ToPointer(),
            LogType: criblcontrolplanesdkgo.Pointer(""),
            MessageField: criblcontrolplanesdkgo.Pointer(""),
            Metadata: []components.Metadatum{
                components.Metadatum{
                    Name: components.FieldNameAuditID,
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1024),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 3645.11,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Manual.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](2690.34),
            Description: criblcontrolplanesdkgo.Pointer("abaft among madly kielbasa despite strictly solemnly whether digital"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://well-worn-casket.net"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
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
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Region: components.RegionOptionsUs.ToPointer(),
            AccountID: "123456",
            EventType: "CriblEvent",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1024),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 1898.39,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationMethodOptions2Manual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("for upon healthily distorted as"),
            CustomURL: criblcontrolplanesdkgo.Pointer("https://squeaky-mathematics.org/"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputNewrelicEventsPqControls{},
            APIKey: criblcontrolplanesdkgo.Pointer("your-api-key"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateRegion: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateEventType: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateCustomURL: criblcontrolplanesdkgo.Pointer("https://old-yak.info/"),
        },
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
            },
            Protocol: components.ProtocolOptionsGrpc.ToPointer(),
            Endpoint: "http://localhost:4317",
            OtlpVersion: components.OutputOpenTelemetryOTLPVersionZeroDot10Dot0.ToPointer(),
            Compress: components.CompressionOptions4Gzip.ToPointer(),
            HTTPCompress: components.CompressionOptions5Gzip.ToPointer(),
            AuthType: components.AuthenticationTypeOptionsNone.ToPointer(),
            HTTPTracesEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPMetricsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPLogsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](30),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("playfully lounge fooey vet overwork acknowledge above drat"),
            Username: criblcontrolplanesdkgo.Pointer("Eve11"),
            Password: criblcontrolplanesdkgo.Pointer("6wNRRvPHwjcljoo"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
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
                    HTTPStatus: 7232.05,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
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
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
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
                "<value 2>",
                "<value 3>",
            },
            URL: "http://localhost:9091/api/v1/write",
            MetricRenameExpr: criblcontrolplanesdkgo.Pointer("name.replace(/[^a-zA-Z0-9_]/g, '_')"),
            SendMetadata: criblcontrolplanesdkgo.Pointer(true),
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
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            SafeHeaders: []string{
                "<value 1>",
            },
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 2064.65,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationTypeOptionsPrometheusAuthNone.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("instead though litter fortunately perfectly unless vary"),
            MetricsFlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](60),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputPrometheusPqControls{},
            Username: criblcontrolplanesdkgo.Pointer("Tracey68"),
            Password: criblcontrolplanesdkgo.Pointer("kTcECZfBZpBWuvT"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://unlawful-duster.name"),
        },
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
                "<value 2>",
                "<value 3>",
            },
            Format: components.OutputRingDataFormatJSON.ToPointer(),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("<value>"),
            MaxDataSize: criblcontrolplanesdkgo.Pointer("1GB"),
            MaxDataTime: criblcontrolplanesdkgo.Pointer("24h"),
            Compress: components.DataCompressionFormatOptionsPersistenceGzip.ToPointer(),
            DestPath: criblcontrolplanesdkgo.Pointer("<value>"),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("safely swerve sparkling print um equal coop astride crushing"),
        },
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
                    Description: criblcontrolplanesdkgo.Pointer("fort supposing bleakly recede hmph"),
                    Final: criblcontrolplanesdkgo.Pointer(true),
                },
            },
            Description: criblcontrolplanesdkgo.Pointer("whose how hard-to-find"),
        },
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
            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            DestPath: criblcontrolplanesdkgo.Pointer(""),
            ObjectACL: components.ObjectACLOptionsPrivate.ToPointer(),
            StorageClass: components.StorageClassOptionsOnezoneIa.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAes256.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            PartitionExpr: criblcontrolplanesdkgo.Pointer("C.Time.strftime(_time ? _time : Date.now()/1000, '%Y/%m/%d')"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            BaseFileName: criblcontrolplanesdkgo.Pointer("`CriblOut`"),
            FileNameSuffix: criblcontrolplanesdkgo.Pointer("`.${C.env[\"CRIBL_WORKER_ID\"]}.${__format}${__compression === \"gzip\" ? \".gz\" : \"\"}`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](32),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            HeaderLine: criblcontrolplanesdkgo.Pointer(""),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](64),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](6434.77),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](857.42),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2924.72),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](4562.74),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](4),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](100),
            Description: criblcontrolplanesdkgo.Pointer("starch nephew including hmph separately mobilize sugary deliberately"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            Compress: components.CompressionOptions2Gzip.ToPointer(),
            CompressionLevel: components.CompressionLevelOptionsBestSpeed.ToPointer(),
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("1MB"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(true),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.OutputSecurityLakeSignatureVersionV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: "arn:aws:iam::123456789012:role/my-role",
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            StagePath: "/tmp/staging",
            AddIDToStagePath: criblcontrolplanesdkgo.Pointer(true),
            ObjectACL: components.ObjectACLOptionsPrivate.ToPointer(),
            StorageClass: components.StorageClassOptionsStandard.ToPointer(),
            ServerSideEncryption: components.ServerSideEncryptionForUploadedObjectsOptionsAwsKms.ToPointer(),
            KmsKeyID: criblcontrolplanesdkgo.Pointer("<id>"),
            RemoveEmptyDirs: criblcontrolplanesdkgo.Pointer(true),
            BaseFileName: criblcontrolplanesdkgo.Pointer("`CriblOut`"),
            MaxFileSizeMB: criblcontrolplanesdkgo.Pointer[float64](32),
            MaxOpenFiles: criblcontrolplanesdkgo.Pointer[float64](100),
            HeaderLine: criblcontrolplanesdkgo.Pointer(""),
            WriteHighWaterMark: criblcontrolplanesdkgo.Pointer[float64](64),
            OnBackpressure: components.BackpressureBehaviorOptions1Block.ToPointer(),
            DeadletterEnabled: criblcontrolplanesdkgo.Pointer(false),
            OnDiskFullBackpressure: components.DiskSpaceProtectionOptionsBlock.ToPointer(),
            ForceCloseOnShutdown: criblcontrolplanesdkgo.Pointer(false),
            RetrySettings: &components.RetrySettingsType{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                InitialBackoffMs: criblcontrolplanesdkgo.Pointer[float64](6434.77),
                BackoffMultiplier: criblcontrolplanesdkgo.Pointer[float64](857.42),
                MaxBackoffMs: criblcontrolplanesdkgo.Pointer[float64](2924.72),
                JitterPercent: criblcontrolplanesdkgo.Pointer[float64](4562.74),
            },
            MaxFileOpenTimeSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxFileIdleTimeSec: criblcontrolplanesdkgo.Pointer[float64](30),
            MaxConcurrentFileParts: criblcontrolplanesdkgo.Pointer[float64](4),
            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
            MaxClosingFilesToBackpressure: criblcontrolplanesdkgo.Pointer[float64](100),
            AccountID: "123456789012",
            CustomSource: "my-custom-source",
            AutomaticSchema: criblcontrolplanesdkgo.Pointer(false),
            ParquetVersion: components.ParquetVersionOptionsParquet26.ToPointer(),
            ParquetDataPageVersion: components.DataPageVersionOptionsDataPageV2.ToPointer(),
            ParquetRowGroupLength: criblcontrolplanesdkgo.Pointer[float64](10000),
            ParquetPageSize: criblcontrolplanesdkgo.Pointer("1MB"),
            ShouldLogInvalidRows: criblcontrolplanesdkgo.Pointer(false),
            KeyValueMetadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            EnableStatistics: criblcontrolplanesdkgo.Pointer(true),
            EnableWritePageIndex: criblcontrolplanesdkgo.Pointer(true),
            EnablePageChecksum: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("ha till furthermore excepting diver"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            EmptyDirCleanupSec: criblcontrolplanesdkgo.Pointer[float64](300),
            DirectoryBatchSize: criblcontrolplanesdkgo.Pointer[float64](1000),
            ParquetSchema: criblcontrolplanesdkgo.Pointer("<value>"),
            DeadletterPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/outputs/dead-letter"),
            MaxRetryNum: criblcontrolplanesdkgo.Pointer[float64](20),
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
            },
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1000),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 2382.67,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthTypeEnumOauth.ToPointer(),
            LoginURL: "https://login.microsoftonline.com",
            Secret: "client-secret",
            ClientID: "client-id",
            Scope: criblcontrolplanesdkgo.Pointer("https://monitor.azure.com/.default"),
            EndpointURLConfiguration: components.EndpointConfigurationURL,
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](933.62),
            Description: criblcontrolplanesdkgo.Pointer("exhausted against respectful editor majestically"),
            Format: components.OutputSentinelFormatNdjson.ToPointer(),
            CustomSourceExpression: criblcontrolplanesdkgo.Pointer("__httpOut"),
            CustomDropWhenNull: criblcontrolplanesdkgo.Pointer(false),
            CustomEventDelimiter: criblcontrolplanesdkgo.Pointer("\\n"),
            CustomContentType: criblcontrolplanesdkgo.Pointer("application/x-ndjson"),
            CustomPayloadExpression: criblcontrolplanesdkgo.Pointer("`${events}`"),
            AdvancedContentType: criblcontrolplanesdkgo.Pointer("application/json"),
            FormatEventCode: criblcontrolplanesdkgo.Pointer("<value>"),
            FormatPayloadCode: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputSentinelPqControls{},
            URL: criblcontrolplanesdkgo.Pointer("https://your-workspace.ingest.monitor.azure.com"),
            DcrID: criblcontrolplanesdkgo.Pointer("<id>"),
            DceEndpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            StreamName: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateLoginURL: criblcontrolplanesdkgo.Pointer("https://lumpy-muscat.name/"),
            TemplateSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateClientID: criblcontrolplanesdkgo.Pointer("<id>"),
            TemplateScope: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://long-honesty.org/"),
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
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](5120),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](5),
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
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 9611.42,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("but yellow whereas yowza sweetly scent lobster tectonics consequently golden"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            BaseURL: criblcontrolplanesdkgo.Pointer("https://<Your-S1-Tenant>.sentinelone.net"),
            HostExpression: criblcontrolplanesdkgo.Pointer("__e.host || C.os.hostname()"),
            SourceExpression: criblcontrolplanesdkgo.Pointer("__e.source || (__e.__criblMetrics ? 'metrics' : 'cribl')"),
            SourceTypeExpression: criblcontrolplanesdkgo.Pointer("__e.sourcetype || 'dottedJson'"),
            DataSourceCategoryExpression: criblcontrolplanesdkgo.Pointer("'security'"),
            DataSourceNameExpression: criblcontrolplanesdkgo.Pointer("__e.__dataSourceName || 'cribl'"),
            DataSourceVendorExpression: criblcontrolplanesdkgo.Pointer("__e.__dataSourceVendor || 'cribl'"),
            EventTypeExpression: criblcontrolplanesdkgo.Pointer(""),
            Host: criblcontrolplanesdkgo.Pointer("C.os.hostname()"),
            Source: criblcontrolplanesdkgo.Pointer("cribl"),
            SourceType: criblcontrolplanesdkgo.Pointer("hecRawParser"),
            DataSourceCategory: criblcontrolplanesdkgo.Pointer("security"),
            DataSourceName: criblcontrolplanesdkgo.Pointer("cribl"),
            DataSourceVendor: criblcontrolplanesdkgo.Pointer("cribl"),
            EventType: criblcontrolplanesdkgo.Pointer(""),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
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
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Endpoint: "ingest.lightstep.com:443",
            TokenSecret: "your-token-secret",
            AuthTokenName: criblcontrolplanesdkgo.Pointer("lightstep-access-token"),
            OtlpVersion: components.OtlpVersionOptions1OneDot3Dot1,
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](2048),
            Protocol: components.ProtocolOptionsHTTP,
            Compress: components.CompressionOptions4Gzip.ToPointer(),
            HTTPCompress: components.CompressionOptions5Gzip.ToPointer(),
            HTTPTracesEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPMetricsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            HTTPLogsEndpointOverride: criblcontrolplanesdkgo.Pointer("<value>"),
            Metadata: []components.ItemsTypeKeyValueMetadata{
                components.ItemsTypeKeyValueMetadata{
                    Key: "",
                    Value: "<value>",
                },
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            FailedRequestLoggingMode: components.FailedRequestLoggingModeOptionsNone.ToPointer(),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            KeepAliveTime: criblcontrolplanesdkgo.Pointer[float64](30),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("pinion among beyond behind thorough helplessly sadly anti"),
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
                    HTTPStatus: 2454.78,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
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
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
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
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Realm: "us0",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 2401.56,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("hawk next gnash sympathetically numb spork discrete whereas"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
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
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](0),
            Description: criblcontrolplanesdkgo.Pointer("executor caring kiddingly violently into"),
        },
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
            },
            TopicArn: "arn:aws:sns:us-east-1:123456789012:my-topic",
            MessageGroupID: "my-message-group",
            MaxRetries: criblcontrolplanesdkgo.Pointer[float64](9109),
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
            AwsSecretKey: criblcontrolplanesdkgo.Pointer("<value>"),
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
            SignatureVersion: components.OutputSnsSignatureVersionV4.ToPointer(),
            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(false),
            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](3600),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("where duh via numb as abaft pfft"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
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
                "<value 2>",
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Host: "localhost",
            Port: 9997,
            NestedFields: components.NestedFieldSerializationOptionsNone.ToPointer(),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            EnableMultiMetrics: criblcontrolplanesdkgo.Pointer(false),
            EnableACK: criblcontrolplanesdkgo.Pointer(true),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(false),
            MaxS2Sversion: components.MaxS2SVersionOptionsV3.ToPointer(),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("wrathful upright guest excepting"),
            MaxFailedHealthChecks: criblcontrolplanesdkgo.Pointer[float64](1),
            Compress: components.CompressionOptionsDisabled.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputSplunkPqControls{},
            AuthToken: criblcontrolplanesdkgo.Pointer(""),
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
            },
            LoadBalanced: criblcontrolplanesdkgo.Pointer(true),
            NextQueue: criblcontrolplanesdkgo.Pointer("indexQueue"),
            TCPRouting: criblcontrolplanesdkgo.Pointer("nowhere"),
            TLS: &components.TLSSettingsClientSideType1{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
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
            SafeHeaders: []string{
                "<value 1>",
                "<value 2>",
            },
            EnableMultiMetrics: criblcontrolplanesdkgo.Pointer(false),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 6617.86,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("bus majestically absent once opposite brr enroll"),
            URL: criblcontrolplanesdkgo.Pointer("http://localhost:8088/services/collector/event"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Urls: []components.OutputSplunkHecURL{
                components.OutputSplunkHecURL{
                    URL: "http://localhost:8088/services/collector/event",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1),
                    TemplateURL: criblcontrolplanesdkgo.Pointer("https://proud-blowgun.biz/"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](600),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](300),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputSplunkHecPqControls{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://fair-cauliflower.biz/"),
        },
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
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](600),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](0),
            NestedFields: components.NestedFieldSerializationOptionsNone.ToPointer(),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            EnableMultiMetrics: criblcontrolplanesdkgo.Pointer(false),
            EnableACK: criblcontrolplanesdkgo.Pointer(true),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(false),
            MaxS2Sversion: components.MaxS2SVersionOptionsV3.ToPointer(),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            IndexerDiscovery: criblcontrolplanesdkgo.Pointer(false),
            SenderUnhealthyTimeAllowance: criblcontrolplanesdkgo.Pointer[float64](100),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("veto rarely yahoo and"),
            MaxFailedHealthChecks: criblcontrolplanesdkgo.Pointer[float64](1),
            Compress: components.CompressionOptionsDisabled.ToPointer(),
            IndexerDiscoveryConfigs: &components.IndexerDiscoveryConfigs{
                Site: "default",
                MasterURI: "https://simple-gym.biz",
                RefreshIntervalSec: 300,
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                AuthTokens: []components.OutputSplunkLbAuthToken{
                    components.OutputSplunkLbAuthToken{
                        AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
                        AuthToken: criblcontrolplanesdkgo.Pointer(""),
                        TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
                    },
                },
                AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
                AuthToken: criblcontrolplanesdkgo.Pointer(""),
                TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            },
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "localhost",
                    Port: 9997,
                    TLS: components.TLSOptionsHostsItemsInherit.ToPointer(),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1),
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputSplunkLbPqControls{},
            AuthToken: criblcontrolplanesdkgo.Pointer(""),
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
                "<value 3>",
            },
            QueueName: "my-queue",
            QueueType: components.OutputSqsQueueTypeStandard,
            AwsAccountID: criblcontrolplanesdkgo.Pointer("<id>"),
            MessageGroupID: criblcontrolplanesdkgo.Pointer("cribl"),
            CreateQueue: criblcontrolplanesdkgo.Pointer(true),
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
            MaxQueueSize: criblcontrolplanesdkgo.Pointer[float64](100),
            MaxRecordSizeKB: criblcontrolplanesdkgo.Pointer[float64](256),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            MaxInProgress: criblcontrolplanesdkgo.Pointer[float64](10),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("meatloaf um what contrast acknowledge helpfully however shy nor bah"),
            AwsAPIKey: criblcontrolplanesdkgo.Pointer("<value>"),
            AwsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
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
                "<value 3>",
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
            Mtu: criblcontrolplanesdkgo.Pointer[float64](512),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](0),
            Description: criblcontrolplanesdkgo.Pointer("bah reflate however accompanist hunt gut hello"),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
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
            Mtu: criblcontrolplanesdkgo.Pointer[float64](512),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](0),
            Description: criblcontrolplanesdkgo.Pointer("gosh trusty though reword overstay conclude bleach cautiously"),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
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
                "<value 3>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
            },
            URL: "https://endpoint1.collection.us2.sumologic.com",
            CustomSource: criblcontrolplanesdkgo.Pointer("<value>"),
            CustomCategory: criblcontrolplanesdkgo.Pointer("<value>"),
            Format: components.OutputSumoLogicDataFormatJSON.ToPointer(),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](1024),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 1298.59,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](6460.48),
            Description: criblcontrolplanesdkgo.Pointer("subsidy translation carp impractical woot mozzarella beyond however underneath aboard"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputSumoLogicPqControls{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://busy-thigh.biz"),
        },
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
            },
            Protocol: components.OutputSyslogProtocolTCP.ToPointer(),
            Facility: components.FacilityOne.ToPointer(),
            Severity: components.OutputSyslogSeverityNotice.ToPointer(),
            AppName: criblcontrolplanesdkgo.Pointer("Cribl"),
            MessageFormat: components.MessageFormatRfc3164.ToPointer(),
            TimestampFormat: components.TimestampFormatSyslog.ToPointer(),
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
            OctetCountFraming: criblcontrolplanesdkgo.Pointer(true),
            LogFailedRequests: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("deeply recount ack satirise typewriter consequently acquaintance"),
            LoadBalanced: criblcontrolplanesdkgo.Pointer(true),
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](514),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "meager-stock.biz",
                    Port: 2033.08,
                    TLS: components.TLSOptionsHostsItemsInherit.ToPointer(),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1),
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](600),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](0),
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            MaxRecordSize: criblcontrolplanesdkgo.Pointer[float64](1500),
            UDPDNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](0),
            EnableIPSpoofing: criblcontrolplanesdkgo.Pointer(false),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
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
            ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv11.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            ConnectionTimeout: criblcontrolplanesdkgo.Pointer[float64](10000),
            WriteTimeout: criblcontrolplanesdkgo.Pointer[float64](60000),
            TokenTTLMinutes: criblcontrolplanesdkgo.Pointer[float64](60),
            SendHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("barring astonishing swine creamy but gosh"),
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](10090),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "flawless-archaeology.net",
                    Port: 2419.51,
                    TLS: components.TLSOptionsHostsItemsInherit.ToPointer(),
                    Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1),
                    TemplateHost: criblcontrolplanesdkgo.Pointer("<value>"),
                    TemplatePort: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](600),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](300),
            MaxConcurrentSenders: criblcontrolplanesdkgo.Pointer[float64](0),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputTcpjsonPqControls{},
            AuthToken: criblcontrolplanesdkgo.Pointer(""),
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
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Domain: "longboard",
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 3992.8,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            Description: criblcontrolplanesdkgo.Pointer("onto speedily hmph"),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
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
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Method: components.MethodOptionsPost.ToPointer(),
            Format: components.OutputWebhookFormatNdjson.ToPointer(),
            KeepAlive: criblcontrolplanesdkgo.Pointer(true),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(true),
            TimeoutSec: criblcontrolplanesdkgo.Pointer[float64](30),
            FlushPeriodSec: criblcontrolplanesdkgo.Pointer[float64](1),
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
                    HTTPStatus: 8283.06,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(false),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            AuthType: components.OutputWebhookAuthenticationTypeNone.ToPointer(),
            TLS: &components.TLSSettingsClientSideType1{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](7404.48),
            LoadBalanced: criblcontrolplanesdkgo.Pointer(false),
            Description: criblcontrolplanesdkgo.Pointer("fisherman physically besides"),
            CustomSourceExpression: criblcontrolplanesdkgo.Pointer("__httpOut"),
            CustomDropWhenNull: criblcontrolplanesdkgo.Pointer(false),
            CustomEventDelimiter: criblcontrolplanesdkgo.Pointer("\\n"),
            CustomContentType: criblcontrolplanesdkgo.Pointer("application/x-ndjson"),
            CustomPayloadExpression: criblcontrolplanesdkgo.Pointer("`${events}`"),
            AdvancedContentType: criblcontrolplanesdkgo.Pointer("application/json"),
            FormatEventCode: criblcontrolplanesdkgo.Pointer("<value>"),
            FormatPayloadCode: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputWebhookPqControls{},
            Username: criblcontrolplanesdkgo.Pointer("Jackeline5"),
            Password: criblcontrolplanesdkgo.Pointer("tI4tGhpVEwbYMeW"),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            CredentialsSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            LoginURL: criblcontrolplanesdkgo.Pointer("https://honored-yogurt.info/"),
            SecretParamName: criblcontrolplanesdkgo.Pointer("<value>"),
            Secret: criblcontrolplanesdkgo.Pointer("<value>"),
            TokenAttributeName: criblcontrolplanesdkgo.Pointer("<value>"),
            AuthHeaderExpr: criblcontrolplanesdkgo.Pointer("`Bearer ${token}`"),
            TokenTimeoutSecs: criblcontrolplanesdkgo.Pointer[float64](3600),
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
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Urls: []components.OutputWebhookURL{
                components.OutputWebhookURL{
                    URL: "https://blind-mentor.com/",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1),
                    TemplateURL: criblcontrolplanesdkgo.Pointer("https://agreeable-dulcimer.net"),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](600),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](300),
            TemplateLoginURL: criblcontrolplanesdkgo.Pointer("https://funny-premise.biz/"),
            TemplateSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://shady-zebra.com/"),
        },
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
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            LoadBalanced: true,
            NextQueue: criblcontrolplanesdkgo.Pointer("indexQueue"),
            TCPRouting: criblcontrolplanesdkgo.Pointer("nowhere"),
            TLS: &components.TLSSettingsClientSideType1{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
                Servername: criblcontrolplanesdkgo.Pointer("<value>"),
                CertificateName: criblcontrolplanesdkgo.Pointer("<value>"),
                CaPath: criblcontrolplanesdkgo.Pointer("<value>"),
                PrivKeyPath: criblcontrolplanesdkgo.Pointer("<value>"),
                CertPath: criblcontrolplanesdkgo.Pointer("<value>"),
                Passphrase: criblcontrolplanesdkgo.Pointer("<value>"),
                MinVersion: components.MinimumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
                MaxVersion: components.MaximumTLSVersionOptionsKafkaSchemaRegistryTLSTlSv12.ToPointer(),
            },
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](4096),
            MaxPayloadEvents: criblcontrolplanesdkgo.Pointer[float64](0),
            Compress: criblcontrolplanesdkgo.Pointer(true),
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
            SafeHeaders: []string{
                "<value 1>",
            },
            EnableMultiMetrics: false,
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 2924.72,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            WizConnectorID: "00000000-0000-0000-0000-000000000000",
            WizEnvironment: "test",
            DataCenter: "us1",
            WizSourcetype: "placeholder",
            Description: criblcontrolplanesdkgo.Pointer("phooey positively a consequently meh until"),
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
            },
            LoadBalanced: criblcontrolplanesdkgo.Pointer(false),
            Concurrency: criblcontrolplanesdkgo.Pointer[float64](5),
            MaxPayloadSizeKB: criblcontrolplanesdkgo.Pointer[float64](9500),
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
            SafeHeaders: []string{
                "<value 1>",
            },
            AuthType: components.OutputXsiamAuthenticationMethodToken.ToPointer(),
            ResponseRetrySettings: []components.ItemsTypeResponseRetrySettings{
                components.ItemsTypeResponseRetrySettings{
                    HTTPStatus: 3551.51,
                    InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                    BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                    MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
                },
            },
            TimeoutRetrySettings: &components.TimeoutRetrySettingsType{
                TimeoutRetry: false,
                InitialBackoff: criblcontrolplanesdkgo.Pointer[float64](1000),
                BackoffRate: criblcontrolplanesdkgo.Pointer[float64](2),
                MaxBackoff: criblcontrolplanesdkgo.Pointer[float64](10000),
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            ThrottleRateReqPerSec: criblcontrolplanesdkgo.Pointer[int64](400),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            TotalMemoryLimitKB: criblcontrolplanesdkgo.Pointer[float64](7982.68),
            Description: criblcontrolplanesdkgo.Pointer("vet simple wholly buzzing if axe"),
            URL: criblcontrolplanesdkgo.Pointer("http://localhost:8088/logs/v1/event"),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(false),
            ExcludeSelf: criblcontrolplanesdkgo.Pointer(false),
            Urls: []components.OutputXsiamURL{
                components.OutputXsiamURL{
                    URL: "https://dental-calculus.info",
                    Weight: criblcontrolplanesdkgo.Pointer[float64](1),
                },
            },
            DNSResolvePeriodSec: criblcontrolplanesdkgo.Pointer[float64](600),
            LoadBalanceStatsPeriodSec: criblcontrolplanesdkgo.Pointer[float64](300),
            Token: criblcontrolplanesdkgo.Pointer("<value>"),
            TextSecret: criblcontrolplanesdkgo.Pointer("<value>"),
            PqStrictOrdering: criblcontrolplanesdkgo.Pointer(true),
            PqRatePerSec: criblcontrolplanesdkgo.Pointer[float64](0),
            PqMode: components.ModeOptionsError.ToPointer(),
            PqMaxBufferSize: criblcontrolplanesdkgo.Pointer[float64](42),
            PqMaxBackpressureSec: criblcontrolplanesdkgo.Pointer[float64](30),
            PqMaxFileSize: criblcontrolplanesdkgo.Pointer("1 MB"),
            PqMaxSize: criblcontrolplanesdkgo.Pointer("5GB"),
            PqPath: criblcontrolplanesdkgo.Pointer("$CRIBL_HOME/state/queues"),
            PqCompress: components.CompressionOptionsPqNone.ToPointer(),
            PqOnBackpressure: components.QueueFullBehaviorOptionsBlock.ToPointer(),
            PqControls: &components.OutputXsiamPqControls{},
            TemplateURL: criblcontrolplanesdkgo.Pointer("https://worthwhile-polarisation.biz/"),
        },
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