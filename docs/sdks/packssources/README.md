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

    res, err := s.Packs.Sources.List(ctx, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                   | Type                                                                                                                                                        | Required                                                                                                                                                    | Description                                                                                                                                                 |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                                                                       | :heavy_check_mark:                                                                                                                                          | The context to use for the request.                                                                                                                         |
| `pack`                                                                                                                                                      | `string`                                                                                                                                                    | :heavy_check_mark:                                                                                                                                          | The <code>id</code> of the Pack.                                                                                                                            |
| `type_`                                                                                                                                                     | []`string`                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                          | Type of Source to include in the results. Each request can include only one <code>type</code> parameter; multiple parameters per request are not supported. |
| `opts`                                                                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                                                                    | :heavy_minus_sign:                                                                                                                                          | The options for this request.                                                                                                                               |

### Response

**[*operations.GetInputSystemByPackResponse](../../models/operations/getinputsystembypackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a new Source. The system-managed provenance field (JSON <code>criblSourceProvenance</code>) must be omitted from the request body within the specified Pack.

### Example Usage: InputCreateExamplesAnthropicCompliance

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesAnthropicCompliance" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyAnthropicCompliance(
        operations.CreateInputSystemByPackInputAnthropicCompliance{
            ID: "anthropic-compliance-source",
            Type: operations.CreateInputSystemByPackTypeAnthropicComplianceAnthropicCompliance,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            TextSecret: "anthropic-api-key-secret",
            ContentConfig: []operations.CreateInputSystemByPackContentConfigAnthropicCompliance{
                operations.CreateInputSystemByPackContentConfigAnthropicCompliance{
                    ContentType: "activities",
                    ContentDescription: criblcontrolplanesdkgo.Pointer("Compliance Activities"),
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    StateTracking: criblcontrolplanesdkgo.Pointer(true),
                    StateUpdateExpression: criblcontrolplanesdkgo.Pointer("__timestampExtracted !== false && {latestTime: (state.latestTime || 0) > _time ? state.latestTime : _time}"),
                    StateMergeExpression: criblcontrolplanesdkgo.Pointer("prevState.latestTime > newState.latestTime ? prevState : newState"),
                    CronSchedule: "*/5 * * * *",
                    Earliest: "-7d@d",
                    Latest: "now",
                    JobTimeout: criblcontrolplanesdkgo.Pointer("300"),
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: InputCreateExamplesAppleUnifiedLogs

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesAppleUnifiedLogs" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyAppleUnifiedLogs(
        operations.CreateInputSystemByPackInputAppleUnifiedLogs{
            ID: "apple-unified-logs-source",
            Type: operations.CreateInputSystemByPackTypeAppleUnifiedLogsAppleUnifiedLogs,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Predicate: "subsystem == \"com.apple.security\"",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: criblcontrolplanesdkgo.Pointer("0.0.0.0"),
            Port: criblcontrolplanesdkgo.Pointer[float64](9109.0),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            QueueName: "azure-blob-queue",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 8088.0,
            HecAPI: "/services/collector",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Brokers: []string{
                "pkc-xxxxx.us-east-1.aws.confluent.cloud:9092",
            },
            Topics: []string{
                "logs",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10090.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            QueueName: "crowdstrike-queue",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 8126.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Samples: []operations.CreateInputSystemByPackSample{
                operations.CreateInputSystemByPackSample{
                    Sample: "sample.json",
                    EventsPerSec: 10.0,
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            DiscoveryType: operations.CreateInputSystemByPackDiscoveryTypeEdgePrometheusStatic,
            Interval: 60.0,
            Targets: []operations.CreateInputSystemByPackTarget{
                operations.CreateInputSystemByPackTarget{
                    Host: "localhost",
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "localhost",
            Port: 9200.0,
            ElasticAPI: "/",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Brokers: []string{
                "myeventhub.servicebus.windows.net:9093",
            },
            Topics: []string{
                "logs",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: InputCreateExamplesEventhubAmqp

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesEventhubAmqp" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyEventhubAmqp(
        operations.CreateInputSystemByPackInputEventhubAmqp{
            ID: "eventhub-amqp-source",
            Type: operations.CreateInputSystemByPackTypeEventhubAmqpEventhubAmqp,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            EventHubName: criblcontrolplanesdkgo.Pointer("my-event-hub"),
            ConsumerGroup: "$Default",
            Checkpointing: operations.CreateInputSystemByPackCheckpointing{
                BlobStore: operations.CreateInputSystemByPackAzureBlobStorage{
                    ContainerName: "my-container",
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Command: "echo \"Hello World\"",
            Interval: criblcontrolplanesdkgo.Pointer[float64](60.0),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Mode: operations.CreateInputSystemByPackInputFileModeManual.ToPointer(),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            TopicName: "my-topic",
            SubscriptionName: "my-subscription",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
                SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
                PqEnabled: criblcontrolplanesdkgo.Pointer(false),
                Host: "0.0.0.0",
                Port: 10080.0,
                PrometheusAPI: "/api/prom/push",
            },
        ),
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Path: "/var/log/journal",
            Journals: []string{
                "system",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Brokers: []string{
                "localhost:9092",
            },
            Topics: []string{
                "logs",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            StreamName: "my-stream",
            Region: "us-east-1",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
            LokiAPI: "/loki/api/v1/push",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            UDPPort: criblcontrolplanesdkgo.Pointer[float64](8125.0),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: InputCreateExamplesMicrosoftGraph

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesMicrosoftGraph" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyMicrosoftGraph(
        operations.CreateInputSystemByPackInputMicrosoftGraph{
            ID: "microsoft-graph-source",
            Type: operations.CreateInputSystemByPackTypeMicrosoftGraphMicrosoftGraph,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            URL: "https://graph.microsoft.com/v1.0/admin/exchange/tracing/messageTraces",
            Interval: 15,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 57000.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Brokers: []string{
                "b-1.example.xxxxx.c2.kafka.us-east-1.amazonaws.com:9092",
            },
            Topics: []string{
                "logs",
            },
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto,
            Region: "us-east-1",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 2055.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            PlanType: components.SubscriptionPlanOptionsEnterpriseGcc,
            TenantID: "tenant-id",
            AppID: "app-id",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            URL: "https://reports.office365.com/ecp/reportingwebservice/reporting.svc/MessageTrace",
            Interval: 15,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            TenantID: "tenant-id",
            AppID: "app-id",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: InputCreateExamplesOkta

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesOkta" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyOkta(
        operations.CreateInputSystemByPackInputOkta{
            ID: "okta-source",
            Type: operations.CreateInputSystemByPackTypeOktaOkta,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            OktaDomain: "your-org",
            TextSecret: "okta-api-token-secret",
            CronSchedule: criblcontrolplanesdkgo.Pointer("*/5 * * * *"),
            Earliest: criblcontrolplanesdkgo.Pointer("-7d@d"),
            Latest: criblcontrolplanesdkgo.Pointer("now"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: InputCreateExamplesOpenAI

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesOpenAI" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyOpenai(
        operations.CreateInputSystemByPackInputOpenai{
            ID: "openai-source",
            Type: operations.CreateInputSystemByPackTypeOpenaiOpenai,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            ContentConfig: []operations.CreateInputSystemByPackContentConfigInput{
                operations.CreateInputSystemByPackContentConfigInput{
                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                    RequestParams: []components.HTTPDiscoveryHeaderConfInputPrometheus{
                        components.HTTPDiscoveryHeaderConfInputPrometheus{
                            Name: "effective_at[gt]",
                            Value: "`${Math.round(Date.now()/1000 - 3600)}`",
                        },
                        components.HTTPDiscoveryHeaderConfInputPrometheus{
                            Name: "limit",
                            Value: "100",
                        },
                    },
                    PaginationType: operations.CreateInputSystemByPackPaginationTypeResponseBody,
                    PaginationAttribute: []string{
                        "last_id",
                    },
                    PaginationLastPageExpr: criblcontrolplanesdkgo.Pointer("has_more === false"),
                    CronSchedule: "0 * * * *",
                    Earliest: "-1h",
                    Latest: "now",
                },
            },
            TextSecret: "openai-api-key-secret",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: InputCreateExamplesOpenAIComplianceLogs

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesOpenAIComplianceLogs" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyOpenaiComplianceLogs(
        operations.CreateInputSystemByPackInputOpenaiComplianceLogs{
            ID: "openai-compliance-logs-source",
            Type: operations.CreateInputSystemByPackTypeOpenaiComplianceLogsOpenaiComplianceLogs,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            TextSecret: "openai-api-key-secret",
            AccountType: operations.CreateInputSystemByPackAccountTypeWorkspace,
            CronSchedule: "*/15 * * * *",
            Earliest: criblcontrolplanesdkgo.Pointer("-1h"),
            Latest: criblcontrolplanesdkgo.Pointer("now"),
            WorkspaceID: criblcontrolplanesdkgo.Pointer("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
            WorkspaceEventTypes: []string{
                "AUDIT_LOG",
                "AUTH_LOG",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 4317.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            DiscoveryType: operations.CreateInputSystemByPackDiscoveryTypePrometheusStatic.ToPointer(),
            Interval: 60.0,
            LogLevel: components.LogLevelOptionsInfo,
            TargetList: []string{
                "http://localhost:9090/metrics",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
            PrometheusAPI: "/write",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 514.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            QueueName: "s3-notifications-queue",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            QueueName: "s3-inventory-queue",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            QueueName: "security-lake-queue",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: InputCreateExamplesServiceNowTable

<!-- UsageSnippet language="go" operationID="createInputSystemByPack" method="post" path="/p/{pack}/system/inputs" example="InputCreateExamplesServiceNowTable" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Create(ctx, "<value>", operations.CreateCreateInputSystemByPackRequestBodyServicenowTable(
        operations.CreateInputSystemByPackInputServicenowTable{
            ID: "servicenow-table-source",
            Type: operations.CreateInputSystemByPackTypeServicenowTableServicenowTable,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Instance: "https://example.service-now.com",
            TableName: "incident",
            Fields: []string{
                "sys_id",
                "number",
                "short_description",
            },
            PageSize: criblcontrolplanesdkgo.Pointer[int64](10000),
            CronSchedule: "0 * * * *",
            Earliest: "-1d",
            Latest: "now",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "192.168.1.1",
            Port: 161.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 9997.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 8088.0,
            SplunkHecAPI: "/services/collector",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            SearchHead: "https://localhost:8089",
            Search: "index=main",
            CronSchedule: "*/15 * * * *",
            Endpoint: "/services/search/v2/jobs/export",
            OutputMode: components.OutputModeOptionsSplunkCollectorConfJSON,
            AuthType: operations.CreateInputSystemByPackAuthenticationTypeSplunkSearchBasic,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            QueueName: "my-queue",
            QueueType: operations.CreateInputSystemByPackQueueTypeStandard,
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
                SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
                PqEnabled: criblcontrolplanesdkgo.Pointer(false),
                Host: "0.0.0.0",
                UDPPort: 514.0,
            },
        ),
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10090.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10090.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 5985.0,
            Subscriptions: []operations.CreateInputSystemByPackSubscription{
                operations.CreateInputSystemByPackSubscription{
                    SubscriptionName: "subscription-1",
                    ContentFormat: operations.CreateInputSystemByPackFormatRenderedText,
                    HeartbeatInterval: 60.0,
                    BatchTimeout: 5.0,
                    Targets: []string{},
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            LogNames: []string{
                "Application",
                "System",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Endpoint: "https://api.wiz.io",
            AuthURL: "https://auth.wiz.io/oauth/token",
            ClientID: "client-id",
            ContentConfig: []operations.CreateInputSystemByPackContentConfigWiz{},
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 8088.0,
            HecAPI: "/services/collector",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `pack`                                                                                                         | `string`                                                                                                       | :heavy_check_mark:                                                                                             | The <code>id</code> of the Pack.                                                                               |
| `requestBody`                                                                                                  | [operations.CreateInputSystemByPackRequestBody](../../models/operations/createinputsystembypackrequestbody.md) | :heavy_check_mark:                                                                                             | Input object.                                                                                                  |
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
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Source to get.                |
| `pack`                                                   | `string`                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Pack.                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetInputSystemByPackAndIDResponse](../../models/operations/getinputsystembypackandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the specified Source.<br/><br/>Provide a complete representation of the Source that you want to update in the request body. This endpoint does not support partial updates. Cribl removes any omitted fields when updating the Source.<br/><br/>Confirm that the configuration in your request body is correct before sending the request. If the configuration is incorrect, the updated Source might not function as expected.<br/><br/>Cribl preserves <code>criblSourceProvenance</code> when you omit it from the request body, and you cannot overwrite it through this endpoint within the specified Pack.

### Example Usage: InputCreateExamplesAnthropicCompliance

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesAnthropicCompliance" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputAnthropicCompliance(
        components.InputAnthropicComplianceInput{
            ID: criblcontrolplanesdkgo.Pointer("anthropic-compliance-source"),
            Type: components.InputAnthropicComplianceTypeAnthropicCompliance,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            TextSecret: "anthropic-api-key-secret",
            ContentConfig: []components.InputAnthropicComplianceContentConfig{
                components.InputAnthropicComplianceContentConfig{
                    ContentType: "activities",
                    ContentDescription: criblcontrolplanesdkgo.Pointer("Compliance Activities"),
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    StateTracking: criblcontrolplanesdkgo.Pointer(true),
                    StateUpdateExpression: criblcontrolplanesdkgo.Pointer("__timestampExtracted !== false && {latestTime: (state.latestTime || 0) > _time ? state.latestTime : _time}"),
                    StateMergeExpression: criblcontrolplanesdkgo.Pointer("prevState.latestTime > newState.latestTime ? prevState : newState"),
                    CronSchedule: "*/5 * * * *",
                    Earliest: "-7d@d",
                    Latest: "now",
                    JobTimeout: criblcontrolplanesdkgo.Pointer("300"),
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: InputCreateExamplesAppleUnifiedLogs

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesAppleUnifiedLogs" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputAppleUnifiedLogs(
        components.InputAppleUnifiedLogsInput{
            ID: criblcontrolplanesdkgo.Pointer("apple-unified-logs-source"),
            Type: components.InputAppleUnifiedLogsTypeAppleUnifiedLogs,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Predicate: "subsystem == \"com.apple.security\"",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
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
        components.InputAppscopeInput{
            ID: criblcontrolplanesdkgo.Pointer("appscope-source"),
            Type: components.InputAppscopeTypeAppscope,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: criblcontrolplanesdkgo.Pointer("0.0.0.0"),
            Port: criblcontrolplanesdkgo.Pointer[float64](9109.0),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputAzureBlobInput{
            ID: criblcontrolplanesdkgo.Pointer("azure-blob-source"),
            Type: components.InputAzureBlobTypeAzureBlob,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            QueueName: "azure-blob-queue",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputCloudflareHecInput{
            ID: criblcontrolplanesdkgo.Pointer("cloudflare-hec-source"),
            Type: components.InputCloudflareHecTypeCloudflareHec,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 8088.0,
            HecAPI: "/services/collector",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputCollectionInput{
            ID: criblcontrolplanesdkgo.Pointer("collection-source"),
            Type: components.InputCollectionTypeCollection,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputConfluentCloudInput{
            ID: criblcontrolplanesdkgo.Pointer("confluent-cloud-source"),
            Type: components.InputConfluentCloudTypeConfluentCloud,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Brokers: []string{
                "pkc-xxxxx.us-east-1.aws.confluent.cloud:9092",
            },
            Topics: []string{
                "logs",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputCriblHTTPInput{
            ID: criblcontrolplanesdkgo.Pointer("cribl-http-source"),
            Type: components.InputCriblHTTPTypeCriblHTTP,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputCriblLakeHTTPInput{
            ID: criblcontrolplanesdkgo.Pointer("cribl-lake-http-source"),
            Type: components.InputCriblLakeHTTPTypeCriblLakeHTTP,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputCriblTCPInput{
            ID: criblcontrolplanesdkgo.Pointer("cribl-tcp-source"),
            Type: components.InputCriblTCPTypeCriblTCP,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10090.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputCrowdstrikeInput{
            ID: criblcontrolplanesdkgo.Pointer("crowdstrike-source"),
            Type: components.InputCrowdstrikeTypeCrowdstrike,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            QueueName: "crowdstrike-queue",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputDatadogAgentInput{
            ID: criblcontrolplanesdkgo.Pointer("datadog-agent-source"),
            Type: components.InputDatadogAgentTypeDatadogAgent,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 8126.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputDatagenInput{
            ID: criblcontrolplanesdkgo.Pointer("datagen-source"),
            Type: components.InputDatagenTypeDatagen,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Samples: []components.InputDatagenSample{
                components.InputDatagenSample{
                    Sample: "sample.json",
                    EventsPerSec: 10.0,
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputEdgePrometheusInput{
            ID: criblcontrolplanesdkgo.Pointer("edge-prometheus-source"),
            Type: components.InputEdgePrometheusTypeEdgePrometheus,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            DiscoveryType: components.InputEdgePrometheusDiscoveryTypeStatic,
            Interval: 60.0,
            Targets: []components.InputEdgePrometheusTarget{
                components.InputEdgePrometheusTarget{
                    Host: "localhost",
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputElasticInput{
            ID: criblcontrolplanesdkgo.Pointer("elastic-source"),
            Type: components.InputElasticTypeElastic,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "localhost",
            Port: 9200.0,
            ElasticAPI: "/",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputEventhubInput{
            ID: criblcontrolplanesdkgo.Pointer("eventhub-source"),
            Type: components.InputEventhubTypeEventhub,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Brokers: []string{
                "myeventhub.servicebus.windows.net:9093",
            },
            Topics: []string{
                "logs",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: InputCreateExamplesEventhubAmqp

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesEventhubAmqp" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputEventhubAmqp(
        components.InputEventhubAmqpInput{
            ID: criblcontrolplanesdkgo.Pointer("eventhub-amqp-source"),
            Type: components.InputEventhubAmqpTypeEventhubAmqp,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            EventHubName: criblcontrolplanesdkgo.Pointer("my-event-hub"),
            ConsumerGroup: "$Default",
            Checkpointing: components.InputEventhubAmqpCheckpointing{
                BlobStore: components.InputEventhubAmqpAzureBlobStorage{
                    ContainerName: "my-container",
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputExecInput{
            ID: criblcontrolplanesdkgo.Pointer("exec-source"),
            Type: components.InputExecTypeExec,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Command: "echo \"Hello World\"",
            Interval: criblcontrolplanesdkgo.Pointer[float64](60.0),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputFileInput{
            ID: criblcontrolplanesdkgo.Pointer("file-source"),
            Type: components.InputFileTypeFile,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Mode: components.InputFileModeManual.ToPointer(),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputFirehoseInput{
            ID: criblcontrolplanesdkgo.Pointer("firehose-source"),
            Type: components.InputFirehoseTypeFirehose,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputGooglePubsubInput{
            ID: criblcontrolplanesdkgo.Pointer("google-pubsub-source"),
            Type: components.InputGooglePubsubTypeGooglePubsub,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            TopicName: "my-topic",
            SubscriptionName: "my-subscription",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.CreateInputGrafanaInputUnionInputGrafanaGrafanaInput1(
            components.InputGrafanaGrafanaInput1{
                ID: criblcontrolplanesdkgo.Pointer("grafana-source"),
                Type: components.InputGrafanaType1Grafana,
                SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
                PqEnabled: criblcontrolplanesdkgo.Pointer(false),
                Host: "0.0.0.0",
                Port: 10080.0,
                PrometheusAPI: "/api/prom/push",
            },
        ),
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputHTTPInput{
            ID: criblcontrolplanesdkgo.Pointer("http-source"),
            Type: components.InputHTTPTypeHTTP,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputHTTPRawInput{
            ID: criblcontrolplanesdkgo.Pointer("http-raw-source"),
            Type: components.InputHTTPRawTypeHTTPRaw,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputJournalFilesInput{
            ID: criblcontrolplanesdkgo.Pointer("journal-files-source"),
            Type: components.InputJournalFilesTypeJournalFiles,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Path: "/var/log/journal",
            Journals: []string{
                "system",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputKafkaInput{
            ID: criblcontrolplanesdkgo.Pointer("kafka-source"),
            Type: components.InputKafkaTypeKafka,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Brokers: []string{
                "localhost:9092",
            },
            Topics: []string{
                "logs",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputKinesisInput{
            ID: criblcontrolplanesdkgo.Pointer("kinesis-source"),
            Type: components.InputKinesisTypeKinesis,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            StreamName: "my-stream",
            Region: "us-east-1",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputKubeEventsInput{
            ID: criblcontrolplanesdkgo.Pointer("kube-events-source"),
            Type: components.InputKubeEventsTypeKubeEvents,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputKubeLogsInput{
            ID: criblcontrolplanesdkgo.Pointer("kube-logs-source"),
            Type: components.InputKubeLogsTypeKubeLogs,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputKubeMetricsInput{
            ID: criblcontrolplanesdkgo.Pointer("kube-metrics-source"),
            Type: components.InputKubeMetricsTypeKubeMetrics,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputLokiInput{
            ID: criblcontrolplanesdkgo.Pointer("loki-source"),
            Type: components.InputLokiTypeLoki,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
            LokiAPI: "/loki/api/v1/push",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputMetricsInput{
            ID: criblcontrolplanesdkgo.Pointer("metrics-source"),
            Type: components.InputMetricsTypeMetrics,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            UDPPort: criblcontrolplanesdkgo.Pointer[float64](8125.0),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: InputCreateExamplesMicrosoftGraph

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesMicrosoftGraph" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputMicrosoftGraph(
        components.InputMicrosoftGraphInput{
            ID: criblcontrolplanesdkgo.Pointer("microsoft-graph-source"),
            Type: components.InputMicrosoftGraphTypeMicrosoftGraph,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            URL: "https://graph.microsoft.com/v1.0/admin/exchange/tracing/messageTraces",
            Interval: 15,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputModelDrivenTelemetryInput{
            ID: criblcontrolplanesdkgo.Pointer("mdt-source"),
            Type: components.InputModelDrivenTelemetryTypeModelDrivenTelemetry,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 57000.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputMskInput{
            ID: criblcontrolplanesdkgo.Pointer("msk-source"),
            Type: components.InputMskTypeMsk,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Brokers: []string{
                "b-1.example.xxxxx.c2.kafka.us-east-1.amazonaws.com:9092",
            },
            Topics: []string{
                "logs",
            },
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto,
            Region: "us-east-1",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputNetflowInput{
            ID: criblcontrolplanesdkgo.Pointer("netflow-source"),
            Type: components.InputNetflowTypeNetflow,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 2055.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputOffice365MgmtInput{
            ID: criblcontrolplanesdkgo.Pointer("office365-mgmt-source"),
            Type: components.InputOffice365MgmtTypeOffice365Mgmt,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            PlanType: components.SubscriptionPlanOptionsEnterpriseGcc,
            TenantID: "tenant-id",
            AppID: "app-id",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputOffice365MsgTraceInput{
            ID: criblcontrolplanesdkgo.Pointer("office365-msg-trace-source"),
            Type: components.InputOffice365MsgTraceTypeOffice365MsgTrace,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            URL: "https://reports.office365.com/ecp/reportingwebservice/reporting.svc/MessageTrace",
            Interval: 15,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputOffice365ServiceInput{
            ID: criblcontrolplanesdkgo.Pointer("office365-service-source"),
            Type: components.InputOffice365ServiceTypeOffice365Service,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            TenantID: "tenant-id",
            AppID: "app-id",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: InputCreateExamplesOkta

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesOkta" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputOkta(
        components.InputOktaInput{
            ID: criblcontrolplanesdkgo.Pointer("okta-source"),
            Type: components.InputOktaTypeOkta,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            OktaDomain: "your-org",
            TextSecret: "okta-api-token-secret",
            CronSchedule: criblcontrolplanesdkgo.Pointer("*/5 * * * *"),
            Earliest: criblcontrolplanesdkgo.Pointer("-7d@d"),
            Latest: criblcontrolplanesdkgo.Pointer("now"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: InputCreateExamplesOpenAI

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesOpenAI" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputOpenai(
        components.InputOpenaiInput{
            ID: criblcontrolplanesdkgo.Pointer("openai-source"),
            Type: components.InputOpenaiTypeOpenai,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            ContentConfig: []components.InputOpenaiContentConfig{
                components.InputOpenaiContentConfig{
                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                    RequestParams: []components.HTTPDiscoveryHeaderConfInputPrometheus{
                        components.HTTPDiscoveryHeaderConfInputPrometheus{
                            Name: "effective_at[gt]",
                            Value: "`${Math.round(Date.now()/1000 - 3600)}`",
                        },
                        components.HTTPDiscoveryHeaderConfInputPrometheus{
                            Name: "limit",
                            Value: "100",
                        },
                    },
                    PaginationType: components.InputOpenaiPaginationTypeResponseBody,
                    PaginationAttribute: []string{
                        "last_id",
                    },
                    PaginationLastPageExpr: criblcontrolplanesdkgo.Pointer("has_more === false"),
                    CronSchedule: "0 * * * *",
                    Earliest: "-1h",
                    Latest: "now",
                },
            },
            TextSecret: "openai-api-key-secret",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: InputCreateExamplesOpenAIComplianceLogs

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesOpenAIComplianceLogs" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputOpenaiComplianceLogs(
        components.InputOpenaiComplianceLogsInput{
            ID: criblcontrolplanesdkgo.Pointer("openai-compliance-logs-source"),
            Type: components.InputOpenaiComplianceLogsTypeOpenaiComplianceLogs,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            TextSecret: "openai-api-key-secret",
            AccountType: components.InputOpenaiComplianceLogsAccountTypeWorkspace,
            CronSchedule: "*/15 * * * *",
            Earliest: criblcontrolplanesdkgo.Pointer("-1h"),
            Latest: criblcontrolplanesdkgo.Pointer("now"),
            WorkspaceID: criblcontrolplanesdkgo.Pointer("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
            WorkspaceEventTypes: []string{
                "AUDIT_LOG",
                "AUTH_LOG",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputOpenTelemetryInput{
            ID: criblcontrolplanesdkgo.Pointer("otel-source"),
            Type: components.InputOpenTelemetryTypeOpenTelemetry,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 4317.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputPrometheusInput{
            ID: criblcontrolplanesdkgo.Pointer("prometheus-source"),
            Type: components.InputPrometheusTypePrometheus,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            DiscoveryType: components.InputPrometheusDiscoveryTypeStatic.ToPointer(),
            Interval: 60.0,
            LogLevel: components.LogLevelOptionsInfo,
            TargetList: []string{
                "http://localhost:9090/metrics",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputPrometheusRwInput{
            ID: criblcontrolplanesdkgo.Pointer("prometheus-rw-source"),
            Type: components.InputPrometheusRwTypePrometheusRw,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
            PrometheusAPI: "/write",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputRawUDPInput{
            ID: criblcontrolplanesdkgo.Pointer("raw-udp-source"),
            Type: components.InputRawUDPTypeRawUDP,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 514.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputS3Input{
            ID: criblcontrolplanesdkgo.Pointer("s3-source"),
            Type: components.InputS3TypeS3,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            QueueName: "s3-notifications-queue",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputS3InventoryInput{
            ID: criblcontrolplanesdkgo.Pointer("s3-inventory-source"),
            Type: components.InputS3InventoryTypeS3Inventory,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            QueueName: "s3-inventory-queue",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputSecurityLakeInput{
            ID: criblcontrolplanesdkgo.Pointer("security-lake-source"),
            Type: components.InputSecurityLakeTypeSecurityLake,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            QueueName: "security-lake-queue",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: InputCreateExamplesServiceNowTable

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="InputCreateExamplesServiceNowTable" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputServicenowTable(
        components.InputServicenowTableInput{
            ID: criblcontrolplanesdkgo.Pointer("servicenow-table-source"),
            Type: components.InputServicenowTableTypeServicenowTable,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Instance: "https://example.service-now.com",
            TableName: "incident",
            Fields: []string{
                "sys_id",
                "number",
                "short_description",
            },
            PageSize: criblcontrolplanesdkgo.Pointer[int64](10000),
            CronSchedule: "0 * * * *",
            Earliest: "-1d",
            Latest: "now",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputSnmpInput{
            ID: criblcontrolplanesdkgo.Pointer("snmp-source"),
            Type: components.InputSnmpTypeSnmp,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "192.168.1.1",
            Port: 161.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputSplunkInput{
            ID: criblcontrolplanesdkgo.Pointer("splunk-source"),
            Type: components.InputSplunkTypeSplunk,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 9997.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputSplunkHecInput{
            ID: criblcontrolplanesdkgo.Pointer("splunk-hec-source"),
            Type: components.InputSplunkHecTypeSplunkHec,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 8088.0,
            SplunkHecAPI: "/services/collector",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputSplunkSearchInput{
            ID: criblcontrolplanesdkgo.Pointer("splunk-search-source"),
            Type: components.InputSplunkSearchTypeSplunkSearch,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            SearchHead: "https://localhost:8089",
            Search: "index=main",
            CronSchedule: "*/15 * * * *",
            Endpoint: "/services/search/v2/jobs/export",
            OutputMode: components.OutputModeOptionsSplunkCollectorConfJSON,
            AuthType: components.InputSplunkSearchAuthenticationTypeBasic,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputSqsInput{
            ID: criblcontrolplanesdkgo.Pointer("sqs-source"),
            Type: components.InputSqsTypeSqs,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            QueueName: "my-queue",
            QueueType: components.InputSqsQueueTypeStandard,
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.CreateInputSyslogInputUnionInputSyslogSyslogInput1(
            components.InputSyslogSyslogInput1{
                ID: criblcontrolplanesdkgo.Pointer("syslog-source"),
                Type: components.InputSyslogType1Syslog,
                SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
                PqEnabled: criblcontrolplanesdkgo.Pointer(false),
                Host: "0.0.0.0",
                UDPPort: 514.0,
            },
        ),
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputSystemMetricsInput{
            ID: criblcontrolplanesdkgo.Pointer("system-metrics-source"),
            Type: components.InputSystemMetricsTypeSystemMetrics,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputSystemStateInput{
            ID: criblcontrolplanesdkgo.Pointer("system-state-source"),
            Type: components.InputSystemStateTypeSystemState,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputTCPInput{
            ID: criblcontrolplanesdkgo.Pointer("tcp-source"),
            Type: components.InputTCPTypeTCP,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10090.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputTcpjsonInput{
            ID: criblcontrolplanesdkgo.Pointer("tcpjson-source"),
            Type: components.InputTcpjsonTypeTcpjson,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10090.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputWefInput{
            ID: criblcontrolplanesdkgo.Pointer("wef-source"),
            Type: components.InputWefTypeWef,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 5985.0,
            Subscriptions: []components.InputWefSubscription{
                components.InputWefSubscription{
                    SubscriptionName: "subscription-1",
                    ContentFormat: components.InputWefFormatRenderedText,
                    HeartbeatInterval: 60.0,
                    BatchTimeout: 5.0,
                    Targets: []string{},
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputWinEventLogsInput{
            ID: criblcontrolplanesdkgo.Pointer("win-event-logs-source"),
            Type: components.InputWinEventLogsTypeWinEventLogs,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            LogNames: []string{
                "Application",
                "System",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputWindowsMetricsInput{
            ID: criblcontrolplanesdkgo.Pointer("windows-metrics-source"),
            Type: components.InputWindowsMetricsTypeWindowsMetrics,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputWizInput{
            ID: criblcontrolplanesdkgo.Pointer("wiz-source"),
            Type: components.InputWizTypeWiz,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Endpoint: "https://api.wiz.io",
            AuthURL: "https://auth.wiz.io/oauth/token",
            ClientID: "client-id",
            ContentConfig: []components.InputWizContentConfig{},
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputWizWebhookInput{
            ID: criblcontrolplanesdkgo.Pointer("wiz-webhook-source"),
            Type: components.InputWizWebhookTypeWizWebhook,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputZscalerHecInput{
            ID: criblcontrolplanesdkgo.Pointer("zscaler-hec-source"),
            Type: components.InputZscalerHecTypeZscalerHec,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 8088.0,
            HecAPI: "/services/collector",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputCriblInput{
            ID: criblcontrolplanesdkgo.Pointer("cribl-source"),
            Type: components.InputCriblTypeCribl,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
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
        components.InputCriblmetricsInput{
            ID: criblcontrolplanesdkgo.Pointer("cribl-metrics-source"),
            Type: components.InputCriblmetricsTypeCriblmetrics,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesAnthropicCompliance

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesAnthropicCompliance" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputAnthropicCompliance(
        components.InputAnthropicComplianceInput{
            ID: criblcontrolplanesdkgo.Pointer("anthropic-compliance-source"),
            Type: components.InputAnthropicComplianceTypeAnthropicCompliance,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            TextSecret: "anthropic-api-key-secret",
            ContentConfig: []components.InputAnthropicComplianceContentConfig{
                components.InputAnthropicComplianceContentConfig{
                    ContentType: "activities",
                    ContentDescription: criblcontrolplanesdkgo.Pointer("Compliance Activities"),
                    Enabled: criblcontrolplanesdkgo.Pointer(true),
                    StateTracking: criblcontrolplanesdkgo.Pointer(true),
                    StateUpdateExpression: criblcontrolplanesdkgo.Pointer("__timestampExtracted !== false && {latestTime: (state.latestTime || 0) > _time ? state.latestTime : _time}"),
                    StateMergeExpression: criblcontrolplanesdkgo.Pointer("prevState.latestTime > newState.latestTime ? prevState : newState"),
                    CronSchedule: "*/5 * * * *",
                    Earliest: "-7d@d",
                    Latest: "now",
                    JobTimeout: criblcontrolplanesdkgo.Pointer("300"),
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesAppleUnifiedLogs

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesAppleUnifiedLogs" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputAppleUnifiedLogs(
        components.InputAppleUnifiedLogsInput{
            ID: criblcontrolplanesdkgo.Pointer("apple-unified-logs-source"),
            Type: components.InputAppleUnifiedLogsTypeAppleUnifiedLogs,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Predicate: "subsystem == \"com.apple.security\"",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesAppscope

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesAppscope" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputAppscopeInput{
            ID: criblcontrolplanesdkgo.Pointer("appscope-source"),
            Type: components.InputAppscopeTypeAppscope,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: criblcontrolplanesdkgo.Pointer("0.0.0.0"),
            Port: criblcontrolplanesdkgo.Pointer[float64](9109.0),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesAzureBlob

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesAzureBlob" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputAzureBlobInput{
            ID: criblcontrolplanesdkgo.Pointer("azure-blob-source"),
            Type: components.InputAzureBlobTypeAzureBlob,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            QueueName: "azure-blob-queue",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesCloudflareHec

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesCloudflareHec" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputCloudflareHecInput{
            ID: criblcontrolplanesdkgo.Pointer("cloudflare-hec-source"),
            Type: components.InputCloudflareHecTypeCloudflareHec,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 8088.0,
            HecAPI: "/services/collector",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesCollection

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesCollection" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputCollectionInput{
            ID: criblcontrolplanesdkgo.Pointer("collection-source"),
            Type: components.InputCollectionTypeCollection,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesConfluentCloud

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesConfluentCloud" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputConfluentCloudInput{
            ID: criblcontrolplanesdkgo.Pointer("confluent-cloud-source"),
            Type: components.InputConfluentCloudTypeConfluentCloud,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Brokers: []string{
                "pkc-xxxxx.us-east-1.aws.confluent.cloud:9092",
            },
            Topics: []string{
                "logs",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesCribl

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesCribl" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputCriblInput{
            ID: criblcontrolplanesdkgo.Pointer("cribl-source"),
            Type: components.InputCriblTypeCribl,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesCriblHttp

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesCriblHttp" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputCriblHTTPInput{
            ID: criblcontrolplanesdkgo.Pointer("cribl-http-source"),
            Type: components.InputCriblHTTPTypeCriblHTTP,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesCriblLakeHttp

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesCriblLakeHttp" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputCriblLakeHTTPInput{
            ID: criblcontrolplanesdkgo.Pointer("cribl-lake-http-source"),
            Type: components.InputCriblLakeHTTPTypeCriblLakeHTTP,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesCriblMetrics

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesCriblMetrics" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputCriblmetricsInput{
            ID: criblcontrolplanesdkgo.Pointer("cribl-metrics-source"),
            Type: components.InputCriblmetricsTypeCriblmetrics,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesCriblTcp

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesCriblTcp" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputCriblTCPInput{
            ID: criblcontrolplanesdkgo.Pointer("cribl-tcp-source"),
            Type: components.InputCriblTCPTypeCriblTCP,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10090.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesCrowdstrike

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesCrowdstrike" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputCrowdstrikeInput{
            ID: criblcontrolplanesdkgo.Pointer("crowdstrike-source"),
            Type: components.InputCrowdstrikeTypeCrowdstrike,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            QueueName: "crowdstrike-queue",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesDatadogAgent

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesDatadogAgent" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputDatadogAgentInput{
            ID: criblcontrolplanesdkgo.Pointer("datadog-agent-source"),
            Type: components.InputDatadogAgentTypeDatadogAgent,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 8126.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesDatagen

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesDatagen" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputDatagenInput{
            ID: criblcontrolplanesdkgo.Pointer("datagen-source"),
            Type: components.InputDatagenTypeDatagen,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Samples: []components.InputDatagenSample{
                components.InputDatagenSample{
                    Sample: "sample.json",
                    EventsPerSec: 10.0,
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesEdgePrometheus

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesEdgePrometheus" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputEdgePrometheusInput{
            ID: criblcontrolplanesdkgo.Pointer("edge-prometheus-source"),
            Type: components.InputEdgePrometheusTypeEdgePrometheus,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            DiscoveryType: components.InputEdgePrometheusDiscoveryTypeStatic,
            Interval: 60.0,
            Targets: []components.InputEdgePrometheusTarget{
                components.InputEdgePrometheusTarget{
                    Host: "localhost",
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesElastic

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesElastic" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputElasticInput{
            ID: criblcontrolplanesdkgo.Pointer("elastic-source"),
            Type: components.InputElasticTypeElastic,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "localhost",
            Port: 9200.0,
            ElasticAPI: "/",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesEventhub

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesEventhub" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputEventhubInput{
            ID: criblcontrolplanesdkgo.Pointer("eventhub-source"),
            Type: components.InputEventhubTypeEventhub,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Brokers: []string{
                "myeventhub.servicebus.windows.net:9093",
            },
            Topics: []string{
                "logs",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesEventhubAmqp

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesEventhubAmqp" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputEventhubAmqp(
        components.InputEventhubAmqpInput{
            ID: criblcontrolplanesdkgo.Pointer("eventhub-amqp-source"),
            Type: components.InputEventhubAmqpTypeEventhubAmqp,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            EventHubName: criblcontrolplanesdkgo.Pointer("my-event-hub"),
            ConsumerGroup: "$Default",
            Checkpointing: components.InputEventhubAmqpCheckpointing{
                BlobStore: components.InputEventhubAmqpAzureBlobStorage{
                    ContainerName: "my-container",
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesExec

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesExec" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputExecInput{
            ID: criblcontrolplanesdkgo.Pointer("exec-source"),
            Type: components.InputExecTypeExec,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Command: "echo \"Hello World\"",
            Interval: criblcontrolplanesdkgo.Pointer[float64](60.0),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesFile

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesFile" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputFileInput{
            ID: criblcontrolplanesdkgo.Pointer("file-source"),
            Type: components.InputFileTypeFile,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Mode: components.InputFileModeManual.ToPointer(),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesFirehose

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesFirehose" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputFirehoseInput{
            ID: criblcontrolplanesdkgo.Pointer("firehose-source"),
            Type: components.InputFirehoseTypeFirehose,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesGooglePubsub

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesGooglePubsub" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputGooglePubsubInput{
            ID: criblcontrolplanesdkgo.Pointer("google-pubsub-source"),
            Type: components.InputGooglePubsubTypeGooglePubsub,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            TopicName: "my-topic",
            SubscriptionName: "my-subscription",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesGrafana

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesGrafana" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.CreateInputGrafanaInputUnionInputGrafanaGrafanaInput1(
            components.InputGrafanaGrafanaInput1{
                ID: criblcontrolplanesdkgo.Pointer("grafana-source"),
                Type: components.InputGrafanaType1Grafana,
                SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
                PqEnabled: criblcontrolplanesdkgo.Pointer(false),
                Host: "0.0.0.0",
                Port: 10080.0,
                PrometheusAPI: "/api/prom/push",
            },
        ),
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesHttp

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesHttp" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputHTTPInput{
            ID: criblcontrolplanesdkgo.Pointer("http-source"),
            Type: components.InputHTTPTypeHTTP,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesHttpRaw

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesHttpRaw" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputHTTPRawInput{
            ID: criblcontrolplanesdkgo.Pointer("http-raw-source"),
            Type: components.InputHTTPRawTypeHTTPRaw,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesJournalFiles

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesJournalFiles" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputJournalFilesInput{
            ID: criblcontrolplanesdkgo.Pointer("journal-files-source"),
            Type: components.InputJournalFilesTypeJournalFiles,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Path: "/var/log/journal",
            Journals: []string{
                "system",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesKafka

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesKafka" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputKafkaInput{
            ID: criblcontrolplanesdkgo.Pointer("kafka-source"),
            Type: components.InputKafkaTypeKafka,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Brokers: []string{
                "localhost:9092",
            },
            Topics: []string{
                "logs",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesKinesis

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesKinesis" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputKinesisInput{
            ID: criblcontrolplanesdkgo.Pointer("kinesis-source"),
            Type: components.InputKinesisTypeKinesis,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            StreamName: "my-stream",
            Region: "us-east-1",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesKubeEvents

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesKubeEvents" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputKubeEventsInput{
            ID: criblcontrolplanesdkgo.Pointer("kube-events-source"),
            Type: components.InputKubeEventsTypeKubeEvents,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesKubeLogs

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesKubeLogs" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputKubeLogsInput{
            ID: criblcontrolplanesdkgo.Pointer("kube-logs-source"),
            Type: components.InputKubeLogsTypeKubeLogs,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesKubeMetrics

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesKubeMetrics" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputKubeMetricsInput{
            ID: criblcontrolplanesdkgo.Pointer("kube-metrics-source"),
            Type: components.InputKubeMetricsTypeKubeMetrics,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesLoki

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesLoki" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputLokiInput{
            ID: criblcontrolplanesdkgo.Pointer("loki-source"),
            Type: components.InputLokiTypeLoki,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
            LokiAPI: "/loki/api/v1/push",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesMetrics

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesMetrics" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputMetricsInput{
            ID: criblcontrolplanesdkgo.Pointer("metrics-source"),
            Type: components.InputMetricsTypeMetrics,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            UDPPort: criblcontrolplanesdkgo.Pointer[float64](8125.0),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesMicrosoftGraph

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesMicrosoftGraph" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputMicrosoftGraph(
        components.InputMicrosoftGraphInput{
            ID: criblcontrolplanesdkgo.Pointer("microsoft-graph-source"),
            Type: components.InputMicrosoftGraphTypeMicrosoftGraph,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            URL: "https://graph.microsoft.com/v1.0/admin/exchange/tracing/messageTraces",
            Interval: 15,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesModelDrivenTelemetry

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesModelDrivenTelemetry" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputModelDrivenTelemetryInput{
            ID: criblcontrolplanesdkgo.Pointer("mdt-source"),
            Type: components.InputModelDrivenTelemetryTypeModelDrivenTelemetry,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 57000.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesMsk

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesMsk" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputMskInput{
            ID: criblcontrolplanesdkgo.Pointer("msk-source"),
            Type: components.InputMskTypeMsk,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Brokers: []string{
                "b-1.example.xxxxx.c2.kafka.us-east-1.amazonaws.com:9092",
            },
            Topics: []string{
                "logs",
            },
            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto,
            Region: "us-east-1",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesNetflow

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesNetflow" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputNetflowInput{
            ID: criblcontrolplanesdkgo.Pointer("netflow-source"),
            Type: components.InputNetflowTypeNetflow,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 2055.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesOffice365Mgmt

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesOffice365Mgmt" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputOffice365MgmtInput{
            ID: criblcontrolplanesdkgo.Pointer("office365-mgmt-source"),
            Type: components.InputOffice365MgmtTypeOffice365Mgmt,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            PlanType: components.SubscriptionPlanOptionsEnterpriseGcc,
            TenantID: "tenant-id",
            AppID: "app-id",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesOffice365MsgTrace

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesOffice365MsgTrace" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputOffice365MsgTraceInput{
            ID: criblcontrolplanesdkgo.Pointer("office365-msg-trace-source"),
            Type: components.InputOffice365MsgTraceTypeOffice365MsgTrace,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            URL: "https://reports.office365.com/ecp/reportingwebservice/reporting.svc/MessageTrace",
            Interval: 15,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesOffice365Service

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesOffice365Service" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputOffice365ServiceInput{
            ID: criblcontrolplanesdkgo.Pointer("office365-service-source"),
            Type: components.InputOffice365ServiceTypeOffice365Service,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            TenantID: "tenant-id",
            AppID: "app-id",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesOkta

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesOkta" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputOkta(
        components.InputOktaInput{
            ID: criblcontrolplanesdkgo.Pointer("okta-source"),
            Type: components.InputOktaTypeOkta,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            OktaDomain: "your-org",
            TextSecret: "okta-api-token-secret",
            CronSchedule: criblcontrolplanesdkgo.Pointer("*/5 * * * *"),
            Earliest: criblcontrolplanesdkgo.Pointer("-7d@d"),
            Latest: criblcontrolplanesdkgo.Pointer("now"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesOpenAI

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesOpenAI" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputOpenai(
        components.InputOpenaiInput{
            ID: criblcontrolplanesdkgo.Pointer("openai-source"),
            Type: components.InputOpenaiTypeOpenai,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            ContentConfig: []components.InputOpenaiContentConfig{
                components.InputOpenaiContentConfig{
                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                    RequestParams: []components.HTTPDiscoveryHeaderConfInputPrometheus{
                        components.HTTPDiscoveryHeaderConfInputPrometheus{
                            Name: "effective_at[gt]",
                            Value: "`${Math.round(Date.now()/1000 - 3600)}`",
                        },
                        components.HTTPDiscoveryHeaderConfInputPrometheus{
                            Name: "limit",
                            Value: "100",
                        },
                    },
                    PaginationType: components.InputOpenaiPaginationTypeResponseBody,
                    PaginationAttribute: []string{
                        "last_id",
                    },
                    PaginationLastPageExpr: criblcontrolplanesdkgo.Pointer("has_more === false"),
                    CronSchedule: "0 * * * *",
                    Earliest: "-1h",
                    Latest: "now",
                },
            },
            TextSecret: "openai-api-key-secret",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesOpenAIComplianceLogs

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesOpenAIComplianceLogs" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputOpenaiComplianceLogs(
        components.InputOpenaiComplianceLogsInput{
            ID: criblcontrolplanesdkgo.Pointer("openai-compliance-logs-source"),
            Type: components.InputOpenaiComplianceLogsTypeOpenaiComplianceLogs,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            TextSecret: "openai-api-key-secret",
            AccountType: components.InputOpenaiComplianceLogsAccountTypeWorkspace,
            CronSchedule: "*/15 * * * *",
            Earliest: criblcontrolplanesdkgo.Pointer("-1h"),
            Latest: criblcontrolplanesdkgo.Pointer("now"),
            WorkspaceID: criblcontrolplanesdkgo.Pointer("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"),
            WorkspaceEventTypes: []string{
                "AUDIT_LOG",
                "AUTH_LOG",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesOpenTelemetry

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesOpenTelemetry" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputOpenTelemetryInput{
            ID: criblcontrolplanesdkgo.Pointer("otel-source"),
            Type: components.InputOpenTelemetryTypeOpenTelemetry,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 4317.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesPrometheus

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesPrometheus" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputPrometheusInput{
            ID: criblcontrolplanesdkgo.Pointer("prometheus-source"),
            Type: components.InputPrometheusTypePrometheus,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            DiscoveryType: components.InputPrometheusDiscoveryTypeStatic.ToPointer(),
            Interval: 60.0,
            LogLevel: components.LogLevelOptionsInfo,
            TargetList: []string{
                "http://localhost:9090/metrics",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesPrometheusRw

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesPrometheusRw" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputPrometheusRwInput{
            ID: criblcontrolplanesdkgo.Pointer("prometheus-rw-source"),
            Type: components.InputPrometheusRwTypePrometheusRw,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
            PrometheusAPI: "/write",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesRawUdp

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesRawUdp" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputRawUDPInput{
            ID: criblcontrolplanesdkgo.Pointer("raw-udp-source"),
            Type: components.InputRawUDPTypeRawUDP,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 514.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesS3

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesS3" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputS3Input{
            ID: criblcontrolplanesdkgo.Pointer("s3-source"),
            Type: components.InputS3TypeS3,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            QueueName: "s3-notifications-queue",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesS3Inventory

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesS3Inventory" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputS3InventoryInput{
            ID: criblcontrolplanesdkgo.Pointer("s3-inventory-source"),
            Type: components.InputS3InventoryTypeS3Inventory,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            QueueName: "s3-inventory-queue",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesSecurityLake

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesSecurityLake" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputSecurityLakeInput{
            ID: criblcontrolplanesdkgo.Pointer("security-lake-source"),
            Type: components.InputSecurityLakeTypeSecurityLake,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            QueueName: "security-lake-queue",
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesServiceNowTable

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesServiceNowTable" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputServicenowTable(
        components.InputServicenowTableInput{
            ID: criblcontrolplanesdkgo.Pointer("servicenow-table-source"),
            Type: components.InputServicenowTableTypeServicenowTable,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Instance: "https://example.service-now.com",
            TableName: "incident",
            Fields: []string{
                "sys_id",
                "number",
                "short_description",
            },
            PageSize: criblcontrolplanesdkgo.Pointer[int64](10000),
            CronSchedule: "0 * * * *",
            Earliest: "-1d",
            Latest: "now",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesSnmp

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesSnmp" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputSnmpInput{
            ID: criblcontrolplanesdkgo.Pointer("snmp-source"),
            Type: components.InputSnmpTypeSnmp,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "192.168.1.1",
            Port: 161.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesSplunk

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesSplunk" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputSplunkInput{
            ID: criblcontrolplanesdkgo.Pointer("splunk-source"),
            Type: components.InputSplunkTypeSplunk,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 9997.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesSplunkHec

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesSplunkHec" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputSplunkHecInput{
            ID: criblcontrolplanesdkgo.Pointer("splunk-hec-source"),
            Type: components.InputSplunkHecTypeSplunkHec,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 8088.0,
            SplunkHecAPI: "/services/collector",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesSplunkSearch

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesSplunkSearch" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputSplunkSearchInput{
            ID: criblcontrolplanesdkgo.Pointer("splunk-search-source"),
            Type: components.InputSplunkSearchTypeSplunkSearch,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            SearchHead: "https://localhost:8089",
            Search: "index=main",
            CronSchedule: "*/15 * * * *",
            Endpoint: "/services/search/v2/jobs/export",
            OutputMode: components.OutputModeOptionsSplunkCollectorConfJSON,
            AuthType: components.InputSplunkSearchAuthenticationTypeBasic,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesSqs

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesSqs" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputSqsInput{
            ID: criblcontrolplanesdkgo.Pointer("sqs-source"),
            Type: components.InputSqsTypeSqs,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            QueueName: "my-queue",
            QueueType: components.InputSqsQueueTypeStandard,
            Region: criblcontrolplanesdkgo.Pointer("us-east-1"),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesSyslog

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesSyslog" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.CreateInputSyslogInputUnionInputSyslogSyslogInput1(
            components.InputSyslogSyslogInput1{
                ID: criblcontrolplanesdkgo.Pointer("syslog-source"),
                Type: components.InputSyslogType1Syslog,
                SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
                PqEnabled: criblcontrolplanesdkgo.Pointer(false),
                Host: "0.0.0.0",
                UDPPort: 514.0,
            },
        ),
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesSystemMetrics

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesSystemMetrics" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputSystemMetricsInput{
            ID: criblcontrolplanesdkgo.Pointer("system-metrics-source"),
            Type: components.InputSystemMetricsTypeSystemMetrics,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesSystemState

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesSystemState" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputSystemStateInput{
            ID: criblcontrolplanesdkgo.Pointer("system-state-source"),
            Type: components.InputSystemStateTypeSystemState,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesTcp

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesTcp" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputTCPInput{
            ID: criblcontrolplanesdkgo.Pointer("tcp-source"),
            Type: components.InputTCPTypeTCP,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10090.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesTcpjson

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesTcpjson" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputTcpjsonInput{
            ID: criblcontrolplanesdkgo.Pointer("tcpjson-source"),
            Type: components.InputTcpjsonTypeTcpjson,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10090.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesWef

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesWef" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputWefInput{
            ID: criblcontrolplanesdkgo.Pointer("wef-source"),
            Type: components.InputWefTypeWef,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 5985.0,
            Subscriptions: []components.InputWefSubscription{
                components.InputWefSubscription{
                    SubscriptionName: "subscription-1",
                    ContentFormat: components.InputWefFormatRenderedText,
                    HeartbeatInterval: 60.0,
                    BatchTimeout: 5.0,
                    Targets: []string{},
                },
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesWinEventLogs

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesWinEventLogs" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputWinEventLogsInput{
            ID: criblcontrolplanesdkgo.Pointer("win-event-logs-source"),
            Type: components.InputWinEventLogsTypeWinEventLogs,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            LogNames: []string{
                "Application",
                "System",
            },
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesWindowsMetrics

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesWindowsMetrics" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputWindowsMetricsInput{
            ID: criblcontrolplanesdkgo.Pointer("windows-metrics-source"),
            Type: components.InputWindowsMetricsTypeWindowsMetrics,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesWiz

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesWiz" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputWizInput{
            ID: criblcontrolplanesdkgo.Pointer("wiz-source"),
            Type: components.InputWizTypeWiz,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Endpoint: "https://api.wiz.io",
            AuthURL: "https://auth.wiz.io/oauth/token",
            ClientID: "client-id",
            ContentConfig: []components.InputWizContentConfig{},
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesWizWebhook

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesWizWebhook" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputWizWebhookInput{
            ID: criblcontrolplanesdkgo.Pointer("wiz-webhook-source"),
            Type: components.InputWizWebhookTypeWizWebhook,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080.0,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```
### Example Usage: UpdateInputExamplesZscalerHec

<!-- UsageSnippet language="go" operationID="updateInputSystemByPackAndId" method="patch" path="/p/{pack}/system/inputs/{id}" example="UpdateInputExamplesZscalerHec" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
        components.InputZscalerHecInput{
            ID: criblcontrolplanesdkgo.Pointer("zscaler-hec-source"),
            Type: components.InputZscalerHecTypeZscalerHec,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 8088.0,
            HecAPI: "/services/collector",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Source to update.             |
| `pack`                                                   | `string`                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Pack.                         |
| `input`                                                  | [components.Input](../../models/components/input.md)     | :heavy_check_mark:                                       | Input object.                                            |
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
    if res.CountedInputResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | `string`                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Source to delete.             |
| `pack`                                                   | `string`                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Pack.                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteInputSystemByPackAndIDResponse](../../models/operations/deleteinputsystembypackandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |