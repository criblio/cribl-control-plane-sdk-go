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
            Port: criblcontrolplanesdkgo.Pointer[float64](9109),
        },
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
            Port: 8088,
            HecAPI: "/services/collector",
        },
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
            Port: 10080,
        },
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
            Port: 10080,
        },
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
            Port: 10090,
        },
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
            Port: 8126,
        },
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
                    EventsPerSec: 10,
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
            Interval: 60,
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
            Port: 9200,
            ElasticAPI: "/",
        },
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
            Interval: criblcontrolplanesdkgo.Pointer[float64](60),
        },
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
            Port: 10080,
        },
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
                Port: 10080,
                PrometheusAPI: "/api/prom/push",
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
            Port: 10080,
        },
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
            Port: 10080,
        },
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
            Port: 10080,
            LokiAPI: "/loki/api/v1/push",
        },
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
            UDPPort: criblcontrolplanesdkgo.Pointer[float64](8125),
        },
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
            Port: 57000,
        },
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
            Port: 2055,
        },
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
            Port: 4317,
        },
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
            Interval: 60,
            LogLevel: operations.CreateInputSystemByPackLogLevelPrometheusInfo,
            TargetList: []string{
                "http://localhost:9090/metrics",
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
            Port: 10080,
            PrometheusAPI: "/write",
        },
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
            Port: 514,
        },
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
            Port: 161,
        },
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
            Port: 9997,
        },
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
            Port: 8088,
            SplunkHecAPI: "/services/collector",
        },
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
            CronSchedule: "0 * * * *",
            Endpoint: "/services/search/jobs/export",
            OutputMode: components.OutputModeOptionsSplunkCollectorConfJSON,
        },
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
                UDPPort: 514,
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
            Port: 10090,
        },
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
            Port: 10090,
        },
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
            Port: 5985,
            Subscriptions: []operations.CreateInputSystemByPackSubscription{
                operations.CreateInputSystemByPackSubscription{
                    SubscriptionName: "subscription-1",
                    ContentFormat: operations.CreateInputSystemByPackFormatRenderedText,
                    HeartbeatInterval: 60,
                    BatchTimeout: 5,
                    Targets: []string{},
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
            Port: 10080,
        },
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
            Port: 8088,
            HecAPI: "/services/collector",
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputAppscope(
        components.InputAppscope{
            ID: criblcontrolplanesdkgo.Pointer("appscope-source"),
            Type: components.InputAppscopeTypeAppscope,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: criblcontrolplanesdkgo.Pointer("0.0.0.0"),
            Port: criblcontrolplanesdkgo.Pointer[float64](9109),
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputAzureBlob(
        components.InputAzureBlob{
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputCloudflareHec(
        components.InputCloudflareHec{
            ID: criblcontrolplanesdkgo.Pointer("cloudflare-hec-source"),
            Type: components.InputCloudflareHecTypeCloudflareHec,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 8088,
            HecAPI: "/services/collector",
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputCollection(
        components.InputCollection{
            ID: criblcontrolplanesdkgo.Pointer("collection-source"),
            Type: components.InputCollectionTypeCollection,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputConfluentCloud(
        components.InputConfluentCloud{
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputCriblHTTP(
        components.InputCriblHTTP{
            ID: criblcontrolplanesdkgo.Pointer("cribl-http-source"),
            Type: components.InputCriblHTTPTypeCriblHTTP,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080,
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputCriblLakeHTTP(
        components.InputCriblLakeHTTP{
            ID: criblcontrolplanesdkgo.Pointer("cribl-lake-http-source"),
            Type: components.InputCriblLakeHTTPTypeCriblLakeHTTP,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080,
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputCriblTCP(
        components.InputCriblTCP{
            ID: criblcontrolplanesdkgo.Pointer("cribl-tcp-source"),
            Type: components.InputCriblTCPTypeCriblTCP,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10090,
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputCrowdstrike(
        components.InputCrowdstrike{
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputDatadogAgent(
        components.InputDatadogAgent{
            ID: criblcontrolplanesdkgo.Pointer("datadog-agent-source"),
            Type: components.InputDatadogAgentTypeDatadogAgent,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 8126,
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputDatagen(
        components.InputDatagen{
            ID: criblcontrolplanesdkgo.Pointer("datagen-source"),
            Type: components.InputDatagenTypeDatagen,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Samples: []components.Sample{
                components.Sample{
                    Sample: "sample.json",
                    EventsPerSec: 10,
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputEdgePrometheus(
        components.InputEdgePrometheus{
            ID: criblcontrolplanesdkgo.Pointer("edge-prometheus-source"),
            Type: components.InputEdgePrometheusTypeEdgePrometheus,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            DiscoveryType: components.InputEdgePrometheusDiscoveryTypeStatic,
            Interval: 60,
            Targets: []components.Target{
                components.Target{
                    Host: "localhost",
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputElastic(
        components.InputElastic{
            ID: criblcontrolplanesdkgo.Pointer("elastic-source"),
            Type: components.InputElasticTypeElastic,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "localhost",
            Port: 9200,
            ElasticAPI: "/",
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputEventhub(
        components.InputEventhub{
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputExec(
        components.InputExec{
            ID: criblcontrolplanesdkgo.Pointer("exec-source"),
            Type: components.InputExecTypeExec,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Command: "echo \"Hello World\"",
            Interval: criblcontrolplanesdkgo.Pointer[float64](60),
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputFile(
        components.InputFile{
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputFirehose(
        components.InputFirehose{
            ID: criblcontrolplanesdkgo.Pointer("firehose-source"),
            Type: components.InputFirehoseTypeFirehose,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080,
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputGooglePubsub(
        components.InputGooglePubsub{
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputGrafana(
        components.CreateInputGrafanaInputGrafanaGrafana1(
            components.InputGrafanaGrafana1{
                ID: criblcontrolplanesdkgo.Pointer("grafana-source"),
                Type: components.InputGrafanaType1Grafana,
                SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
                PqEnabled: criblcontrolplanesdkgo.Pointer(false),
                Host: "0.0.0.0",
                Port: 10080,
                PrometheusAPI: "/api/prom/push",
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputHTTP(
        components.InputHTTP{
            ID: criblcontrolplanesdkgo.Pointer("http-source"),
            Type: components.InputHTTPTypeHTTP,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080,
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputHTTPRaw(
        components.InputHTTPRaw{
            ID: criblcontrolplanesdkgo.Pointer("http-raw-source"),
            Type: components.InputHTTPRawTypeHTTPRaw,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080,
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputJournalFiles(
        components.InputJournalFiles{
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputKafka(
        components.InputKafka{
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputKinesis(
        components.InputKinesis{
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputKubeEvents(
        components.InputKubeEvents{
            ID: criblcontrolplanesdkgo.Pointer("kube-events-source"),
            Type: components.InputKubeEventsTypeKubeEvents,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputKubeLogs(
        components.InputKubeLogs{
            ID: criblcontrolplanesdkgo.Pointer("kube-logs-source"),
            Type: components.InputKubeLogsTypeKubeLogs,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputKubeMetrics(
        components.InputKubeMetrics{
            ID: criblcontrolplanesdkgo.Pointer("kube-metrics-source"),
            Type: components.InputKubeMetricsTypeKubeMetrics,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputLoki(
        components.InputLoki{
            ID: criblcontrolplanesdkgo.Pointer("loki-source"),
            Type: components.InputLokiTypeLoki,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080,
            LokiAPI: "/loki/api/v1/push",
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputMetrics(
        components.InputMetrics{
            ID: criblcontrolplanesdkgo.Pointer("metrics-source"),
            Type: components.InputMetricsTypeMetrics,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            UDPPort: criblcontrolplanesdkgo.Pointer[float64](8125),
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputModelDrivenTelemetry(
        components.InputModelDrivenTelemetry{
            ID: criblcontrolplanesdkgo.Pointer("mdt-source"),
            Type: components.InputModelDrivenTelemetryTypeModelDrivenTelemetry,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 57000,
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputMsk(
        components.InputMsk{
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputNetflow(
        components.InputNetflow{
            ID: criblcontrolplanesdkgo.Pointer("netflow-source"),
            Type: components.InputNetflowTypeNetflow,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 2055,
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputOffice365Mgmt(
        components.InputOffice365Mgmt{
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputOffice365MsgTrace(
        components.InputOffice365MsgTrace{
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputOffice365Service(
        components.InputOffice365Service{
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputOpenTelemetry(
        components.InputOpenTelemetry{
            ID: criblcontrolplanesdkgo.Pointer("otel-source"),
            Type: components.InputOpenTelemetryTypeOpenTelemetry,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 4317,
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputPrometheus(
        components.InputPrometheus{
            ID: criblcontrolplanesdkgo.Pointer("prometheus-source"),
            Type: components.InputPrometheusTypePrometheus,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            DiscoveryType: components.InputPrometheusDiscoveryTypeStatic.ToPointer(),
            Interval: 60,
            LogLevel: components.InputPrometheusLogLevelInfo,
            TargetList: []string{
                "http://localhost:9090/metrics",
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputPrometheusRw(
        components.InputPrometheusRw{
            ID: criblcontrolplanesdkgo.Pointer("prometheus-rw-source"),
            Type: components.InputPrometheusRwTypePrometheusRw,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080,
            PrometheusAPI: "/write",
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputRawUDP(
        components.InputRawUDP{
            ID: criblcontrolplanesdkgo.Pointer("raw-udp-source"),
            Type: components.InputRawUDPTypeRawUDP,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 514,
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputS3(
        components.InputS3{
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputS3Inventory(
        components.InputS3Inventory{
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputSecurityLake(
        components.InputSecurityLake{
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputSnmp(
        components.InputSnmp{
            ID: criblcontrolplanesdkgo.Pointer("snmp-source"),
            Type: components.InputSnmpTypeSnmp,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "192.168.1.1",
            Port: 161,
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputSplunk(
        components.InputSplunk{
            ID: criblcontrolplanesdkgo.Pointer("splunk-source"),
            Type: components.InputSplunkTypeSplunk,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 9997,
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputSplunkHec(
        components.InputSplunkHec{
            ID: criblcontrolplanesdkgo.Pointer("splunk-hec-source"),
            Type: components.InputSplunkHecTypeSplunkHec,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 8088,
            SplunkHecAPI: "/services/collector",
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputSplunkSearch(
        components.InputSplunkSearch{
            ID: criblcontrolplanesdkgo.Pointer("splunk-search-source"),
            Type: components.InputSplunkSearchTypeSplunkSearch,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            SearchHead: "https://localhost:8089",
            Search: "index=main",
            CronSchedule: "0 * * * *",
            Endpoint: "/services/search/jobs/export",
            OutputMode: components.OutputModeOptionsSplunkCollectorConfJSON,
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputSqs(
        components.InputSqs{
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputSyslog(
        components.CreateInputSyslogInputSyslogSyslog1(
            components.InputSyslogSyslog1{
                ID: criblcontrolplanesdkgo.Pointer("syslog-source"),
                Type: components.InputSyslogType1Syslog,
                SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
                PqEnabled: criblcontrolplanesdkgo.Pointer(false),
                Host: "0.0.0.0",
                UDPPort: 514,
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputSystemMetrics(
        components.InputSystemMetrics{
            ID: criblcontrolplanesdkgo.Pointer("system-metrics-source"),
            Type: components.InputSystemMetricsTypeSystemMetrics,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputSystemState(
        components.InputSystemState{
            ID: criblcontrolplanesdkgo.Pointer("system-state-source"),
            Type: components.InputSystemStateTypeSystemState,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputTCP(
        components.InputTCP{
            ID: criblcontrolplanesdkgo.Pointer("tcp-source"),
            Type: components.InputTCPTypeTCP,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10090,
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputTcpjson(
        components.InputTcpjson{
            ID: criblcontrolplanesdkgo.Pointer("tcpjson-source"),
            Type: components.InputTcpjsonTypeTcpjson,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10090,
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputWef(
        components.InputWef{
            ID: criblcontrolplanesdkgo.Pointer("wef-source"),
            Type: components.InputWefTypeWef,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 5985,
            Subscriptions: []components.Subscription{
                components.Subscription{
                    SubscriptionName: "subscription-1",
                    ContentFormat: components.InputWefFormatRenderedText,
                    HeartbeatInterval: 60,
                    BatchTimeout: 5,
                    Targets: []string{},
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputWinEventLogs(
        components.InputWinEventLogs{
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputWindowsMetrics(
        components.InputWindowsMetrics{
            ID: criblcontrolplanesdkgo.Pointer("windows-metrics-source"),
            Type: components.InputWindowsMetricsTypeWindowsMetrics,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputWiz(
        components.InputWiz{
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputWizWebhook(
        components.InputWizWebhook{
            ID: criblcontrolplanesdkgo.Pointer("wiz-webhook-source"),
            Type: components.InputWizWebhookTypeWizWebhook,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 10080,
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputZscalerHec(
        components.InputZscalerHec{
            ID: criblcontrolplanesdkgo.Pointer("zscaler-hec-source"),
            Type: components.InputZscalerHecTypeZscalerHec,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
            Host: "0.0.0.0",
            Port: 8088,
            HecAPI: "/services/collector",
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputCribl(
        components.InputCribl{
            ID: criblcontrolplanesdkgo.Pointer("cribl-source"),
            Type: components.InputCriblTypeCribl,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
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
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Sources.Update(ctx, "<id>", "<value>", components.CreateInputCriblmetrics(
        components.InputCriblmetrics{
            ID: criblcontrolplanesdkgo.Pointer("cribl-metrics-source"),
            Type: components.InputCriblmetricsTypeCriblmetrics,
            SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
            PqEnabled: criblcontrolplanesdkgo.Pointer(false),
        },
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