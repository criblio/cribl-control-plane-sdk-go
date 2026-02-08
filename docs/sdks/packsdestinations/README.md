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
            ContainerName: "my-container",
            StagePath: "/tmp/staging",
        },
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
            ClusterURL: "https://mycluster.kusto.windows.net",
            Database: "mydatabase",
            Table: "mytable",
            IngestMode: operations.CreateOutputSystemByPackIngestionModeStreaming.ToPointer(),
            OauthEndpoint: components.MicrosoftEntraIDAuthenticationEndpointOptionsSaslHTTPSLoginMicrosoftonlineCom,
            TenantID: "tenant-id",
            ClientID: "client-id",
            Scope: "https://mycluster.kusto.windows.net/.default",
            OauthType: operations.CreateOutputSystemByPackOauthTypeAuthenticationMethodClientSecret,
            ClientSecret: criblcontrolplanesdkgo.Pointer("client-secret"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            Compress: components.CompressionOptions2Gzip,
        },
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
            Brokers: []string{
                "myeventhub.servicebus.windows.net:9093",
            },
            Topic: "logs",
        },
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
            LogType: "Cribl",
            AuthType: operations.CreateOutputSystemByPackAuthenticationMethodAzureLogsManual.ToPointer(),
            WorkspaceID: criblcontrolplanesdkgo.Pointer("workspace-id"),
            WorkspaceKey: criblcontrolplanesdkgo.Pointer("workspace-key"),
        },
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
            Region: "us",
            LogType: "UNKNOWN",
            GcpProjectID: "my-project",
            GcpInstance: "customer-id",
        },
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
            URL: "http://localhost:8123/",
            Database: "mydb",
            TableName: "mytable",
        },
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
            Endpoint: "https://account-id.r2.cloudflarestorage.com",
            Bucket: "my-bucket",
            StagePath: "/tmp/staging",
        },
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
            LogGroupName: "my-log-group",
            LogStreamName: "my-log-stream",
            Region: "us-east-1",
        },
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
            Brokers: []string{
                "pkc-xxxxx.us-east-1.aws.confluent.cloud:9092",
            },
            Topic: "logs",
        },
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
        },
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
        },
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
            SystemFields: []string{
                "cribl_pipe",
            },
            Streamtags: []string{},
            LoadBalanced: criblcontrolplanesdkgo.Pointer(false),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
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
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            URL: criblcontrolplanesdkgo.Pointer("https://0.0.0.0:10200"),
        },
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
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](10090),
        },
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
            URL: "https://ingest.us.crowdstrike.com/api/ingest/hec/connection-id/v1/services/collector",
            Format: components.RequestFormatOptionsJSON,
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
        },
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
            WorkspaceID: "your-workspace-id",
            Scope: "my-scope",
            ClientID: "your-client-id",
            Catalog: "my-catalog",
            Schema: "my-schema",
            EventsVolumeName: "my-volume",
            ClientTextSecret: "your-client-secret",
        },
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
            APIKey: criblcontrolplanesdkgo.Pointer("your-api-key"),
        },
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
        },
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
        },
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
            Bucket: "my-bucket",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            StagePath: "/tmp/staging",
        },
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
            AuthType: operations.CreateOutputSystemByPackAuthenticationTypeDynatraceHTTPToken.ToPointer(),
            Format: operations.CreateOutputSystemByPackFormatDynatraceHTTPJSONArray,
            Endpoint: operations.CreateOutputSystemByPackEndpointCloud,
            TelemetryType: operations.CreateOutputSystemByPackTelemetryTypeLogs,
            Token: criblcontrolplanesdkgo.Pointer("your-api-key"),
        },
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
            Protocol: operations.CreateOutputSystemByPackProtocolDynatraceOtlpHTTP,
            Endpoint: "https://your-environment.live.dynatrace.com/api/v2/otlp",
            OtlpVersion: components.OtlpVersionOptions1OneDot3Dot1,
            EndpointType: operations.CreateOutputSystemByPackEndpointTypeSaas,
            TokenSecret: "your-token-secret",
        },
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
            Index: "logs",
        },
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
            URL: "my-cloud-id",
            Index: "logs",
        },
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
            Bucket: "my-bucket",
            Region: "us-east1",
            StagePath: "/tmp/staging",
            Endpoint: "https://storage.googleapis.com",
            CollectorInstanceID: "11112222-3333-4444-5555-666677778888",
        },
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
            DestPath: "/var/log/output",
        },
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
            LogFormatType: operations.CreateOutputSystemByPackSendEventsAsUnstructured,
            Region: criblcontrolplanesdkgo.Pointer("us"),
            CustomerID: criblcontrolplanesdkgo.Pointer("customer-id"),
        },
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
            LogLocationType: operations.CreateOutputSystemByPackLogLocationTypeProject,
            LogNameExpression: "my-log",
            LogLocationExpression: "my-project",
        },
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
            Bucket: "my-bucket",
            Region: "us-east1",
            Endpoint: "https://storage.googleapis.com",
            StagePath: "/tmp/staging",
        },
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
            TopicName: "my-topic",
        },
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
                LokiURL: "https://logs-prod-us-central1.grafana.net",
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
            Protocol: components.DestinationProtocolOptionsTCP,
            Host: "localhost",
            Port: 2003,
        },
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
            Dataset: "my-dataset",
        },
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
            URL: "https://cloud.us.humio.com/api/v1/ingest/hec",
            Format: components.RequestFormatOptionsJSON,
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
        },
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
            URL: "http://localhost:8086",
            Database: criblcontrolplanesdkgo.Pointer("mydb"),
        },
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
            Brokers: []string{
                "localhost:9092",
            },
            Topic: "logs",
        },
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
            StreamName: "my-stream",
            Region: "us-east-1",
        },
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
            URL: "http://localhost:3100/loki/api/v1/push",
        },
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
            Topic: "logs",
            BootstrapServer: "myeventstream.servicebus.windows.net:9093",
        },
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
            Endpoint: "http://localhost:9000",
            Bucket: "my-bucket",
            StagePath: "/tmp/staging",
        },
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
            Brokers: []string{
                "b-1.example.xxxxx.c2.kafka.us-east-1.amazonaws.com:9092",
            },
            Topic: "logs",
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto,
            Region: "us-east-1",
        },
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
            Hosts: []operations.CreateOutputSystemByPackHostNetflow{
                operations.CreateOutputSystemByPackHostNetflow{
                    Host: "localhost",
                    Port: 2055,
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
            APIKey: criblcontrolplanesdkgo.Pointer("your-api-key"),
        },
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
            AccountID: "123456",
            EventType: "CriblEvent",
            APIKey: criblcontrolplanesdkgo.Pointer("your-api-key"),
        },
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
            Endpoint: "http://localhost:4317",
        },
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
            URL: "http://localhost:9091/api/v1/write",
        },
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
        },
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
            Rules: []operations.CreateOutputSystemByPackRule{
                operations.CreateOutputSystemByPackRule{
                    Filter: "true",
                    Output: "my-output",
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
            Bucket: "my-bucket",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            StagePath: "/tmp/staging",
        },
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
            Bucket: "my-bucket",
            Region: "us-east-1",
            AssumeRoleArn: "arn:aws:iam::123456789012:role/my-role",
            StagePath: "/tmp/staging",
            AccountID: "123456789012",
            CustomSource: "my-custom-source",
        },
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
            LoginURL: "https://login.microsoftonline.com",
            Secret: "client-secret",
            ClientID: "client-id",
            EndpointURLConfiguration: operations.CreateOutputSystemByPackEndpointConfigurationURL,
            URL: criblcontrolplanesdkgo.Pointer("https://your-workspace.ingest.monitor.azure.com"),
        },
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
            Region: operations.CreateOutputSystemByPackRegionUs,
            Endpoint: operations.CreateOutputSystemByPackAISIEMEndpointPathRootServicesCollectorEvent,
        },
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
            Endpoint: "ingest.lightstep.com:443",
            TokenSecret: "your-token-secret",
            OtlpVersion: components.OtlpVersionOptions1OneDot3Dot1,
            Protocol: components.ProtocolOptionsHTTP,
        },
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
            Realm: "us0",
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
        },
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
            Hosts: []operations.CreateOutputSystemByPackHostSnmp{
                operations.CreateOutputSystemByPackHostSnmp{
                    Host: "192.168.1.1",
                    Port: 161,
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
            TopicArn: "arn:aws:sns:us-east-1:123456789012:my-topic",
            MessageGroupID: "my-message-group",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
        },
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
            Host: "localhost",
            Port: 9997,
        },
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
        },
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
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "localhost",
                    Port: 9997,
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
            QueueName: "my-queue",
            QueueType: operations.CreateOutputSystemByPackQueueTypeStandard,
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
        },
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
            Protocol: components.DestinationProtocolOptionsUDP,
            Host: "localhost",
            Port: 8125,
        },
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
            Protocol: components.DestinationProtocolOptionsUDP,
            Host: "localhost",
            Port: 8125,
        },
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
            URL: "https://endpoint1.collection.us2.sumologic.com",
        },
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
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](514),
        },
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
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](10090),
        },
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
            Domain: "longboard",
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
        },
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
            URL: criblcontrolplanesdkgo.Pointer("https://example.com/webhook"),
        },
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
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            WizConnectorID: "00000000-0000-0000-0000-000000000000",
            WizEnvironment: "test",
            DataCenter: "us1",
            WizSourcetype: "placeholder",
        },
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
        },
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
            ContainerName: "my-container",
            StagePath: "/tmp/staging",
        },
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
            ClusterURL: "https://mycluster.kusto.windows.net",
            Database: "mydatabase",
            Table: "mytable",
            IngestMode: components.IngestionModeStreaming.ToPointer(),
            OauthEndpoint: components.MicrosoftEntraIDAuthenticationEndpointOptionsSaslHTTPSLoginMicrosoftonlineCom,
            TenantID: "tenant-id",
            ClientID: "client-id",
            Scope: "https://mycluster.kusto.windows.net/.default",
            OauthType: components.OutputAzureDataExplorerAuthenticationMethodClientSecret,
            ClientSecret: criblcontrolplanesdkgo.Pointer("client-secret"),
            Format: components.DataFormatOptionsJSON.ToPointer(),
            Compress: components.CompressionOptions2Gzip,
        },
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
            Brokers: []string{
                "myeventhub.servicebus.windows.net:9093",
            },
            Topic: "logs",
        },
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
            LogType: "Cribl",
            AuthType: components.OutputAzureLogsAuthenticationMethodManual.ToPointer(),
            WorkspaceID: criblcontrolplanesdkgo.Pointer("workspace-id"),
            WorkspaceKey: criblcontrolplanesdkgo.Pointer("workspace-key"),
        },
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
            Region: "us",
            LogType: "UNKNOWN",
            GcpProjectID: "my-project",
            GcpInstance: "customer-id",
        },
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
            URL: "http://localhost:8123/",
            Database: "mydb",
            TableName: "mytable",
        },
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
            Endpoint: "https://account-id.r2.cloudflarestorage.com",
            Bucket: "my-bucket",
            StagePath: "/tmp/staging",
        },
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
            LogGroupName: "my-log-group",
            LogStreamName: "my-log-stream",
            Region: "us-east-1",
        },
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
            Brokers: []string{
                "pkc-xxxxx.us-east-1.aws.confluent.cloud:9092",
            },
            Topic: "logs",
        },
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
        },
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
        },
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
            SystemFields: []string{
                "cribl_pipe",
            },
            Streamtags: []string{},
            LoadBalanced: criblcontrolplanesdkgo.Pointer(false),
            TLS: &components.TLSSettingsClientSideTypeKafkaSchemaRegistry{
                Disabled: criblcontrolplanesdkgo.Pointer(true),
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
            },
            ResponseHonorRetryAfterHeader: criblcontrolplanesdkgo.Pointer(true),
            OnBackpressure: components.BackpressureBehaviorOptionsBlock.ToPointer(),
            UseRoundRobinDNS: criblcontrolplanesdkgo.Pointer(true),
            URL: criblcontrolplanesdkgo.Pointer("https://0.0.0.0:10200"),
        },
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
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](10090),
        },
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
            URL: "https://ingest.us.crowdstrike.com/api/ingest/hec/connection-id/v1/services/collector",
            Format: components.RequestFormatOptionsJSON,
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
        },
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
            WorkspaceID: "your-workspace-id",
            Scope: "my-scope",
            ClientID: "your-client-id",
            Catalog: "my-catalog",
            Schema: "my-schema",
            EventsVolumeName: "my-volume",
            ClientTextSecret: "your-client-secret",
        },
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
            APIKey: criblcontrolplanesdkgo.Pointer("your-api-key"),
        },
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
        },
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
        },
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
            Bucket: "my-bucket",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            StagePath: "/tmp/staging",
        },
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
            AuthType: components.OutputDynatraceHTTPAuthenticationTypeToken.ToPointer(),
            Format: components.OutputDynatraceHTTPFormatJSONArray,
            Endpoint: components.EndpointCloud,
            TelemetryType: components.TelemetryTypeLogs,
            Token: criblcontrolplanesdkgo.Pointer("your-api-key"),
        },
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
            Protocol: components.OutputDynatraceOtlpProtocolHTTP,
            Endpoint: "https://your-environment.live.dynatrace.com/api/v2/otlp",
            OtlpVersion: components.OtlpVersionOptions1OneDot3Dot1,
            EndpointType: components.EndpointTypeSaas,
            TokenSecret: "your-token-secret",
        },
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
            Index: "logs",
        },
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
            URL: "my-cloud-id",
            Index: "logs",
        },
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
            Bucket: "my-bucket",
            Region: "us-east1",
            StagePath: "/tmp/staging",
            Endpoint: "https://storage.googleapis.com",
            CollectorInstanceID: "11112222-3333-4444-5555-666677778888",
        },
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
            DestPath: "/var/log/output",
        },
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
            LogFormatType: components.SendEventsAsUnstructured,
            Region: criblcontrolplanesdkgo.Pointer("us"),
            CustomerID: criblcontrolplanesdkgo.Pointer("customer-id"),
        },
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
            LogLocationType: components.LogLocationTypeProject,
            LogNameExpression: "my-log",
            LogLocationExpression: "my-project",
        },
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
            Bucket: "my-bucket",
            Region: "us-east1",
            Endpoint: "https://storage.googleapis.com",
            StagePath: "/tmp/staging",
        },
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
            TopicName: "my-topic",
        },
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
                LokiURL: "https://logs-prod-us-central1.grafana.net",
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
            Protocol: components.DestinationProtocolOptionsTCP,
            Host: "localhost",
            Port: 2003,
        },
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
            Dataset: "my-dataset",
        },
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
            URL: "https://cloud.us.humio.com/api/v1/ingest/hec",
            Format: components.RequestFormatOptionsJSON,
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
        },
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
            URL: "http://localhost:8086",
            Database: criblcontrolplanesdkgo.Pointer("mydb"),
        },
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
            Brokers: []string{
                "localhost:9092",
            },
            Topic: "logs",
        },
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
            StreamName: "my-stream",
            Region: "us-east-1",
        },
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
            URL: "http://localhost:3100/loki/api/v1/push",
        },
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
            Topic: "logs",
            BootstrapServer: "myeventstream.servicebus.windows.net:9093",
        },
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
            Endpoint: "http://localhost:9000",
            Bucket: "my-bucket",
            StagePath: "/tmp/staging",
        },
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
            Brokers: []string{
                "b-1.example.xxxxx.c2.kafka.us-east-1.amazonaws.com:9092",
            },
            Topic: "logs",
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto,
            Region: "us-east-1",
        },
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
            Hosts: []components.OutputNetflowHost{
                components.OutputNetflowHost{
                    Host: "localhost",
                    Port: 2055,
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
            APIKey: criblcontrolplanesdkgo.Pointer("your-api-key"),
        },
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
            AccountID: "123456",
            EventType: "CriblEvent",
            APIKey: criblcontrolplanesdkgo.Pointer("your-api-key"),
        },
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
            Endpoint: "http://localhost:4317",
        },
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
            URL: "http://localhost:9091/api/v1/write",
        },
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
        },
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
            Rules: []components.OutputRouterRule{
                components.OutputRouterRule{
                    Filter: "true",
                    Output: "my-output",
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
            Bucket: "my-bucket",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
            StagePath: "/tmp/staging",
        },
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
            Bucket: "my-bucket",
            Region: "us-east-1",
            AssumeRoleArn: "arn:aws:iam::123456789012:role/my-role",
            StagePath: "/tmp/staging",
            AccountID: "123456789012",
            CustomSource: "my-custom-source",
        },
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
            LoginURL: "https://login.microsoftonline.com",
            Secret: "client-secret",
            ClientID: "client-id",
            EndpointURLConfiguration: components.EndpointConfigurationURL,
            URL: criblcontrolplanesdkgo.Pointer("https://your-workspace.ingest.monitor.azure.com"),
        },
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
            Region: components.RegionUs,
            Endpoint: components.AISIEMEndpointPathRootServicesCollectorEvent,
        },
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
            Endpoint: "ingest.lightstep.com:443",
            TokenSecret: "your-token-secret",
            OtlpVersion: components.OtlpVersionOptions1OneDot3Dot1,
            Protocol: components.ProtocolOptionsHTTP,
        },
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
            Realm: "us0",
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
        },
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
            Hosts: []components.OutputSnmpHost{
                components.OutputSnmpHost{
                    Host: "192.168.1.1",
                    Port: 161,
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
            TopicArn: "arn:aws:sns:us-east-1:123456789012:my-topic",
            MessageGroupID: "my-message-group",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
        },
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
            Host: "localhost",
            Port: 9997,
        },
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
        },
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
            Hosts: []components.ItemsTypeHosts{
                components.ItemsTypeHosts{
                    Host: "localhost",
                    Port: 9997,
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
            QueueName: "my-queue",
            QueueType: components.OutputSqsQueueTypeStandard,
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
        },
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
            Protocol: components.DestinationProtocolOptionsUDP,
            Host: "localhost",
            Port: 8125,
        },
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
            Protocol: components.DestinationProtocolOptionsUDP,
            Host: "localhost",
            Port: 8125,
        },
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
            URL: "https://endpoint1.collection.us2.sumologic.com",
        },
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
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](514),
        },
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
            Host: criblcontrolplanesdkgo.Pointer("localhost"),
            Port: criblcontrolplanesdkgo.Pointer[float64](10090),
        },
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
            Domain: "longboard",
            Token: criblcontrolplanesdkgo.Pointer("your-token"),
        },
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
            URL: criblcontrolplanesdkgo.Pointer("https://example.com/webhook"),
        },
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
            AuthType: components.AuthenticationMethodOptionsAuthTokensItemsManual.ToPointer(),
            WizConnectorID: "00000000-0000-0000-0000-000000000000",
            WizEnvironment: "test",
            DataCenter: "us1",
            WizSourcetype: "placeholder",
        },
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
        },
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