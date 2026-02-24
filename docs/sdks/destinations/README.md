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
            ClusterURL: "https://mycluster.kusto.windows.net",
            Database: "mydatabase",
            Table: "mytable",
            IngestMode: operations.CreateOutputIngestionModeStreaming.ToPointer(),
            OauthEndpoint: components.MicrosoftEntraIDAuthenticationEndpointOptionsSaslHTTPSLoginMicrosoftonlineCom,
            TenantID: "tenant-id",
            ClientID: "client-id",
            Scope: "https://mycluster.kusto.windows.net/.default",
            OauthType: operations.CreateOutputOauthTypeAuthenticationMethodClientSecret,
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
            LogType: "Cribl",
            AuthType: operations.CreateOutputAuthenticationMethodAzureLogsManual.ToPointer(),
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
        },
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
        },
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
        },
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
        },
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
            AuthType: operations.CreateOutputAuthenticationTypeDynatraceHTTPToken.ToPointer(),
            Format: operations.CreateOutputFormatDynatraceHTTPJSONArray,
            Endpoint: operations.CreateOutputEndpointCloud,
            TelemetryType: operations.CreateOutputTelemetryTypeLogs,
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
            Protocol: operations.CreateOutputProtocolDynatraceOtlpHTTP,
            Endpoint: "https://your-environment.live.dynatrace.com/api/v2/otlp",
            OtlpVersion: components.OtlpVersionOptions1OneDot3Dot1,
            EndpointType: operations.CreateOutputEndpointTypeSaas,
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
            LogFormatType: operations.CreateOutputSendEventsAsUnstructured,
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
            LogLocationType: operations.CreateOutputLogLocationTypeProject,
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
### Example Usage: OutputCreateExamplesLocalSearchStorage

<!-- UsageSnippet language="go" operationID="createOutput" method="post" path="/system/outputs" example="OutputCreateExamplesLocalSearchStorage" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Create(ctx, operations.CreateCreateOutputRequestLocalSearchStorage(
        operations.CreateOutputOutputLocalSearchStorage{
            ID: "local-search-storage-output",
            Type: operations.CreateOutputTypeLocalSearchStorageLocalSearchStorage,
            URL: "http://localhost:8123/",
            Database: "default",
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
            Hosts: []operations.CreateOutputHostNetflow{
                operations.CreateOutputHostNetflow{
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
        },
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
            Rules: []operations.CreateOutputRule{
                operations.CreateOutputRule{
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
            LoginURL: "https://login.microsoftonline.com",
            Secret: "client-secret",
            ClientID: "client-id",
            EndpointURLConfiguration: operations.CreateOutputEndpointConfigurationURL,
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
            Region: operations.CreateOutputRegionUs,
            Endpoint: operations.CreateOutputAISIEMEndpointPathRootServicesCollectorEvent,
        },
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
            Hosts: []operations.CreateOutputHostSnmp{
                operations.CreateOutputHostSnmp{
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
        },
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
            QueueName: "my-queue",
            QueueType: operations.CreateOutputQueueTypeStandard,
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
        },
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
        },
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
        },
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
        },
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
        },
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
### Example Usage: OutputCreateExamplesLocalSearchStorage

<!-- UsageSnippet language="go" operationID="updateOutputById" method="patch" path="/system/outputs/{id}" example="OutputCreateExamplesLocalSearchStorage" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Destinations.Update(ctx, "<id>", components.CreateOutputLocalSearchStorage(
        components.OutputLocalSearchStorage{
            ID: criblcontrolplanesdkgo.Pointer("local-search-storage-output"),
            Type: components.OutputLocalSearchStorageTypeLocalSearchStorage,
            URL: "http://localhost:8123/",
            Database: "default",
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
        },
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
        },
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
        },
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