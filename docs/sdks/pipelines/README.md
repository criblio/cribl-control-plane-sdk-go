# Pipelines

## Overview

Actions related to Pipelines

### Available Operations

* [List](#list) - List all Pipelines
* [Create](#create) - Create a Pipeline
* [Get](#get) - Get a Pipeline
* [Update](#update) - Update a Pipeline
* [Delete](#delete) - Delete a Pipeline

## List

Get a list of all Pipelines.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPipeline" method="get" path="/pipelines" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
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

**[*operations.ListPipelineResponse](../../models/operations/listpipelineresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a new Pipeline.

### Example Usage: PipelineExamplesAggregateMetrics

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesAggregateMetrics" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "aggregate-metrics-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that aggregates process metrics: CPU, memory, and bytes over time windows"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputAggregateMetrics(
                    components.PipelineFunctionAggregateMetrics{
                        Filter: criblcontrolplanesdkgo.Pointer("(_metric == 'proc.cpu_perc' || __criblMetrics[0].nameExpr.includes(\"'proc.cpu_perc'\")) || (_metric == 'proc.mem_perc' || __criblMetrics[0].nameExpr.includes(\"'proc.mem_perc'\")) || (_metric == 'proc.bytes_in' || __criblMetrics[0].nameExpr.includes(\"'proc.bytes_in'\"))"),
                        ID: components.PipelineFunctionAggregateMetricsIDAggregateMetrics,
                        Description: criblcontrolplanesdkgo.Pointer("modulo tall usefully jump platter scrape whoever till shy miskey"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionAggregateMetricsConf{
                            Passthrough: criblcontrolplanesdkgo.Pointer(false),
                            PreserveGroupBys: criblcontrolplanesdkgo.Pointer(false),
                            SufficientStatsOnly: criblcontrolplanesdkgo.Pointer(false),
                            Prefix: criblcontrolplanesdkgo.Pointer("<value>"),
                            TimeWindow: "10s",
                            Aggregations: []components.Aggregation{
                                components.Aggregation{
                                    MetricType: components.PipelineFunctionAggregateMetricsMetricTypeGauge,
                                    Agg: "avg(_value || proc.cpu_perc).as(proc.cpu_perc_avg)",
                                },
                                components.Aggregation{
                                    MetricType: components.PipelineFunctionAggregateMetricsMetricTypeGauge,
                                    Agg: "sum(_value || proc.mem_perc).as(proc.mem_perc_sum)",
                                },
                                components.Aggregation{
                                    MetricType: components.PipelineFunctionAggregateMetricsMetricTypeCounter,
                                    Agg: "count(_value || proc.bytes_in).as(proc.bytes_in_count)",
                                },
                            },
                            Groupbys: []string{
                                "proc",
                            },
                            FlushEventLimit: criblcontrolplanesdkgo.Pointer[float64](9795.75),
                            FlushMemLimit: criblcontrolplanesdkgo.Pointer("<value>"),
                            Cumulative: criblcontrolplanesdkgo.Pointer(false),
                            ShouldTreatDotsAsLiterals: criblcontrolplanesdkgo.Pointer(true),
                            Add: []components.PipelineFunctionAggregateMetricsAdd{
                                components.PipelineFunctionAggregateMetricsAdd{
                                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                                    Value: "<value>",
                                },
                            },
                            FlushOnInputClose: criblcontrolplanesdkgo.Pointer(true),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesAggregations

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesAggregations" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "aggregation-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that aggregates rejected bytes grouped by source address every 10 seconds"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputAggregation(
                    components.PipelineFunctionAggregation{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionAggregationIDAggregation,
                        Description: criblcontrolplanesdkgo.Pointer("at mad back within towards who save before redraw tooth"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionAggregationConf{
                            Passthrough: criblcontrolplanesdkgo.Pointer(false),
                            PreserveGroupBys: criblcontrolplanesdkgo.Pointer(false),
                            SufficientStatsOnly: criblcontrolplanesdkgo.Pointer(false),
                            MetricsMode: criblcontrolplanesdkgo.Pointer(false),
                            Prefix: criblcontrolplanesdkgo.Pointer("<value>"),
                            TimeWindow: "10s",
                            Aggregations: []string{
                                "sum(bytes).where(action==\"REJECT\").as(TotalBytes)",
                            },
                            Groupbys: []string{
                                "srcaddr",
                            },
                            FlushEventLimit: criblcontrolplanesdkgo.Pointer[float64](5847.83),
                            FlushMemLimit: criblcontrolplanesdkgo.Pointer("<value>"),
                            Cumulative: criblcontrolplanesdkgo.Pointer(false),
                            SearchAggMode: criblcontrolplanesdkgo.Pointer("<value>"),
                            Add: []components.ItemsTypeAdd{
                                components.ItemsTypeAdd{
                                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                                    Value: "<value>",
                                },
                            },
                            ShouldTreatDotsAsLiterals: criblcontrolplanesdkgo.Pointer(false),
                            FlushOnInputClose: criblcontrolplanesdkgo.Pointer(true),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesAutoTimestamp

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesAutoTimestamp" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "auto-timestamp-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that extracts timestamps from event data using auto timestamp function"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputAutoTimestamp(
                    components.PipelineFunctionAutoTimestamp{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionAutoTimestampIDAutoTimestamp,
                        Description: criblcontrolplanesdkgo.Pointer("drive pro shore pfft"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.FunctionConfSchemaAutoTimestamp{
                            SrcField: criblcontrolplanesdkgo.Pointer("_raw"),
                            DstField: criblcontrolplanesdkgo.Pointer("_time"),
                            DefaultTimezone: criblcontrolplanesdkgo.Pointer("local"),
                            TimeExpression: criblcontrolplanesdkgo.Pointer("time.getTime() / 1000"),
                            Offset: criblcontrolplanesdkgo.Pointer[float64](0),
                            MaxLen: criblcontrolplanesdkgo.Pointer[float64](150),
                            DefaultTime: components.DefaultTimeNow.ToPointer(),
                            LatestDateAllowed: criblcontrolplanesdkgo.Pointer("+1week"),
                            Spacer: criblcontrolplanesdkgo.Pointer("<value>"),
                            EarliestDateAllowed: criblcontrolplanesdkgo.Pointer("-420weeks"),
                            Timestamps: []components.Timestamp{
                                components.Timestamp{
                                    Regex: "/(\\d{1,2}\\/\\d{2}\\/\\d{4}\\s\\d{1,2}:\\d{2}:\\d{2}\\s\\w{2})/",
                                    Strptime: "%Y-%m-%d %H:%M:%S",
                                },
                            },
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesCEFSerializer

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesCEFSerializer" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "cef-serializer-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that formats events in CEF format with custom header and extension fields"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputCef(
                    components.PipelineFunctionCefInput{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionCefIDCef,
                        Description: criblcontrolplanesdkgo.Pointer("upright founder when yet"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaCefInput{
                            OutputField: criblcontrolplanesdkgo.Pointer("_raw"),
                            Header: []components.HeaderInput{
                                components.HeaderInput{
                                    Value: "'CEF:0'",
                                },
                                components.HeaderInput{
                                    Value: "'Cribl'",
                                },
                                components.HeaderInput{
                                    Value: "'Cribl'",
                                },
                                components.HeaderInput{
                                    Value: "C.version",
                                },
                                components.HeaderInput{
                                    Value: "420",
                                },
                                components.HeaderInput{
                                    Value: "'Cribl Event'",
                                },
                                components.HeaderInput{
                                    Value: "6",
                                },
                            },
                            Extension: []components.Extension{
                                components.Extension{
                                    Name: "c6a1Label",
                                    Value: "'Colorado_Ext_Bldg7'",
                                },
                            },
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesChain

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesChain" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "chain-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that chains to another pipeline for sequential data processing"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputChain(
                    components.PipelineFunctionChain{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionChainIDChain,
                        Description: criblcontrolplanesdkgo.Pointer("hm loyally whenever selfish whoever"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionChainConf{
                            Processor: "prometheus_metrics",
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesClone

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesClone" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "clone-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that creates cloned events with additional fields for comparison or routing"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputClone(
                    components.PipelineFunctionClone{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionCloneIDClone,
                        Description: criblcontrolplanesdkgo.Pointer("confiscate positively until plus bar wherever"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.FunctionConfSchemaClone{
                            Clones: []map[string]string{
                                map[string]string{
                                    "env": "staging",
                                },
                                map[string]string{
                                    "index": "clones",
                                },
                            },
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesComment

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesComment" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "comment-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline with comment function for documentation"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputComment(
                    components.PipelineFunctionComment{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionCommentIDComment,
                        Description: criblcontrolplanesdkgo.Pointer("magnificent ew scram tight"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.FunctionConfSchemaComment{
                            Comment: criblcontrolplanesdkgo.Pointer("This function processes security events and enriches them with DNS lookups"),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesDNSLookup

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesDNSLookup" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "dns-lookup-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that performs DNS lookups to resolve hostnames and IP addresses"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputDNSLookup(
                    components.PipelineFunctionDNSLookup{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionDNSLookupIDDNSLookup,
                        Description: criblcontrolplanesdkgo.Pointer("aha entice artistic meh oof"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.FunctionConfSchemaDNSLookup{
                            DNSLookupFields: []components.DNSLookupField{
                                components.DNSLookupField{
                                    InFieldName: criblcontrolplanesdkgo.Pointer("hostname"),
                                    ResourceRecordType: components.ResourceRecordTypeA.ToPointer(),
                                    OutFieldName: criblcontrolplanesdkgo.Pointer("hostname_ip"),
                                },
                            },
                            ReverseLookupFields: []components.ReverseLookupField{
                                components.ReverseLookupField{
                                    InFieldName: criblcontrolplanesdkgo.Pointer("src_ip"),
                                    OutFieldName: criblcontrolplanesdkgo.Pointer("src_hostname"),
                                },
                            },
                            DNSServers: []string{
                                "<value 1>",
                                "<value 2>",
                                "<value 3>",
                            },
                            CacheTTL: criblcontrolplanesdkgo.Pointer[float64](30),
                            MaxCacheSize: criblcontrolplanesdkgo.Pointer[float64](5000),
                            UseResolvConf: criblcontrolplanesdkgo.Pointer(false),
                            LookupFallback: criblcontrolplanesdkgo.Pointer(false),
                            DomainOverrides: []string{
                                "<value 1>",
                            },
                            LookupFailLogLevel: components.LogLevelForFailedLookupsError.ToPointer(),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesDrop

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesDrop" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "drop-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that drops events containing success messages"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputDrop(
                    components.PipelineFunctionDrop{
                        Filter: criblcontrolplanesdkgo.Pointer("_raw.search(/success/i)>=0"),
                        ID: components.PipelineFunctionDropIDDrop,
                        Description: criblcontrolplanesdkgo.Pointer("mid mockingly gah electronics sate"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaDrop{},
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesDropDimensions

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesDropDimensions" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "drop-dimensions-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that drops specified dimensions from metrics to reduce cardinality"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputDropDimensions(
                    components.PipelineFunctionDropDimensions{
                        Filter: criblcontrolplanesdkgo.Pointer("(_metric == 'proc.cpu_perc' || __criblMetrics[0].nameExpr.includes(\"'proc.cpu_perc'\")) && (__criblMetrics[0].dims.includes(\"proc\"))"),
                        ID: components.PipelineFunctionDropDimensionsIDDropDimensions,
                        Description: criblcontrolplanesdkgo.Pointer("sympathetically perfectly quarrelsomely excluding"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionDropDimensionsConf{
                            TimeWindow: "10s",
                            DropDimensions: []string{
                                "proc",
                                "pie",
                                "unit",
                            },
                            FlushOnInputClose: criblcontrolplanesdkgo.Pointer(true),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesDynamicSampling

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesDynamicSampling" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "dynamic-sampling-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that dynamically samples events based on volume using square root mode"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputDynamicSampling(
                    components.PipelineFunctionDynamicSampling{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionDynamicSamplingIDDynamicSampling,
                        Description: criblcontrolplanesdkgo.Pointer("likewise apropos by whoever unkempt ugh wherever"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.PipelineFunctionDynamicSamplingConf{
                            Mode: components.SampleModeSqrt,
                            KeyExpr: "`${domain}:${httpCode}`",
                            SamplePeriod: criblcontrolplanesdkgo.Pointer[float64](20),
                            MinEvents: criblcontrolplanesdkgo.Pointer[float64](3),
                            MaxSampleRate: criblcontrolplanesdkgo.Pointer[float64](3),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesEmpty

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesEmpty" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "empty-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer(""),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{},
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesEval

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesEval" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "eval-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that evaluates JavaScript expressions to add, modify, and remove fields"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputEval(
                    components.PipelineFunctionEval{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionEvalIDEval,
                        Description: criblcontrolplanesdkgo.Pointer("thread circa gah beside attest misappropriate humiliating"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaEval{
                            Add: []components.FunctionConfSchemaEvalAdd{
                                components.FunctionConfSchemaEvalAdd{
                                    Name: criblcontrolplanesdkgo.Pointer("action"),
                                    Value: "login == 'error' ? 'blocked' : action",
                                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                                },
                                components.FunctionConfSchemaEvalAdd{
                                    Name: criblcontrolplanesdkgo.Pointer("myTags"),
                                    Value: "login == 'error' ? [...myTags, 'error'] : myTags",
                                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                                },
                            },
                            Keep: []string{
                                "host",
                                "source",
                                "action",
                                "myTags",
                            },
                            Remove: []string{
                                "identification",
                            },
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesEventBreaker

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesEventBreaker" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "event-breaker-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that breaks large event streams into discrete events using regex"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputEventBreaker(
                    components.PipelineFunctionEventBreaker{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionEventBreakerIDEventBreaker,
                        Description: criblcontrolplanesdkgo.Pointer("upbeat powerfully camouflage upbeat spirit violently premier tricky bowler"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.PipelineFunctionEventBreakerConf{
                            ExistingOrNew: components.ExistingOrNewNew,
                            ShouldMarkCriblBreaker: criblcontrolplanesdkgo.Pointer(true),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesFlatten

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesFlatten" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "flatten-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that flattens nested JSON structures into top-level fields"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputFlatten(
                    components.PipelineFunctionFlatten{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionFlattenIDFlatten,
                        Description: criblcontrolplanesdkgo.Pointer("zebra busily woot alongside fess"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.FunctionConfSchemaFlatten{
                            Fields: []string{},
                            Prefix: criblcontrolplanesdkgo.Pointer(""),
                            Depth: criblcontrolplanesdkgo.Pointer[float64](5),
                            Delimiter: criblcontrolplanesdkgo.Pointer("_"),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesFoldKeys

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesFoldKeys" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "fold-keys-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that transforms flat field names with separators into nested structures"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputFoldkeys(
                    components.PipelineFunctionFoldkeys{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionFoldkeysIDFoldkeys,
                        Description: criblcontrolplanesdkgo.Pointer("rudely worthwhile cross everlasting ghost freezing majority duffel loftily abaft"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaFoldkeys{
                            DeleteOriginal: criblcontrolplanesdkgo.Pointer(true),
                            Separator: criblcontrolplanesdkgo.Pointer("_"),
                            SelectionRegExp: criblcontrolplanesdkgo.Pointer("^data"),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesGeoIP

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesGeoIP" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "geoip-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that enriches events with geolocation data from IP addresses"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputGeoip(
                    components.PipelineFunctionGeoip{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionGeoipIDGeoip,
                        Description: criblcontrolplanesdkgo.Pointer("decongestant substantiate soon er finally pfft override knottily"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.PipelineFunctionGeoipConf{
                            File: "GeoLite2-City.mmdb",
                            InField: criblcontrolplanesdkgo.Pointer("ip"),
                            OutField: criblcontrolplanesdkgo.Pointer("geoip"),
                            AdditionalFields: []components.AdditionalField{
                                components.AdditionalField{
                                    ExtraInField: "src_ip",
                                    ExtraOutField: "src_geoip",
                                },
                            },
                            OutFieldMappings: &components.OutputFieldMappings{},
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesGrok

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesGrok" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "grok-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that extracts structured fields from log data using Grok patterns"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputGrok(
                    components.PipelineFunctionGrok{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionGrokIDGrok,
                        Description: criblcontrolplanesdkgo.Pointer("hopelessly viability morning monthly decisive onto strange"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionGrokConf{
                            Pattern: "%{TIMESTAMP_ISO8601:event_time} %{LOGLEVEL:log_level} %{GREEDYDATA:log_message}",
                            PatternList: []components.PatternList{},
                            Source: criblcontrolplanesdkgo.Pointer("_raw"),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesGuard

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesGuard" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "guard-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that scans for sensitive data and applies mitigation expressions"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputSensitiveDataScanner(
                    components.PipelineFunctionSensitiveDataScanner{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionSensitiveDataScannerIDSensitiveDataScanner,
                        Description: criblcontrolplanesdkgo.Pointer("judgementally while why"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionSensitiveDataScannerConf{
                            Rules: []components.PipelineFunctionSensitiveDataScannerRule{
                                components.PipelineFunctionSensitiveDataScannerRule{
                                    RulesetID: "Finance_Global",
                                    ReplaceExpr: "'REDACTED'",
                                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                                },
                            },
                            Fields: []string{
                                "_raw",
                            },
                            ExcludeFields: []string{},
                            Flags: []components.Flag{
                                components.Flag{
                                    Name: criblcontrolplanesdkgo.Pointer("_sensitive"),
                                    Value: "true",
                                },
                            },
                            IncludeDetectedRules: criblcontrolplanesdkgo.Pointer(true),
                            BackgroundDetection: criblcontrolplanesdkgo.Pointer(true),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesJSONUnroll

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesJSONUnroll" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "json-unroll-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that unrolls JSON arrays into individual events while retaining parent fields"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputJSONUnroll(
                    components.PipelineFunctionJSONUnroll{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionJSONUnrollIDJSONUnroll,
                        Description: criblcontrolplanesdkgo.Pointer("unbearably boo regarding"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionJSONUnrollConf{
                            Path: "allCars",
                            Name: criblcontrolplanesdkgo.Pointer("cars"),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesLookup

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesLookup" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "lookup-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that enriches events with location data from IP address lookups"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputLookup(
                    components.PipelineFunctionLookup{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionLookupIDLookup,
                        Description: criblcontrolplanesdkgo.Pointer("ouch who phooey bookcase elegantly why now alive"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionLookupConf{
                            File: "ip_locations.csv",
                            DbLookup: criblcontrolplanesdkgo.Pointer(false),
                            MatchMode: "exact",
                            MatchType: "<value>",
                            ReloadPeriodSec: -1,
                            InFields: []components.InField{
                                components.InField{
                                    EventField: "destination_ip",
                                    LookupField: criblcontrolplanesdkgo.Pointer("ip"),
                                },
                            },
                            OutFields: []components.OutField{
                                components.OutField{
                                    LookupField: "location",
                                    EventField: criblcontrolplanesdkgo.Pointer("location"),
                                    DefaultValue: criblcontrolplanesdkgo.Pointer("Unknown"),
                                },
                            },
                            AddToEvent: criblcontrolplanesdkgo.Pointer(false),
                            IgnoreCase: false,
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesMask

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesMask" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "mask-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that masks Social Security numbers and other sensitive data"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputMask(
                    components.PipelineFunctionMask{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionMaskIDMask,
                        Description: criblcontrolplanesdkgo.Pointer("from handle for"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionMaskConf{
                            Rules: []components.PipelineFunctionMaskRule{
                                components.PipelineFunctionMaskRule{
                                    MatchRegex: "/(social=)(\\d+)/",
                                    ReplaceExpr: "`${g1}${C.Mask.md5(g2)}`",
                                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                                },
                            },
                            Fields: []string{
                                "_raw",
                            },
                            Depth: criblcontrolplanesdkgo.Pointer[int64](5),
                            Flags: []components.ItemsTypeAdd{
                                components.ItemsTypeAdd{
                                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                                    Value: "<value>",
                                },
                            },
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesNumerify

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesNumerify" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "numerify-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that converts string numbers to numeric type for mathematical operations"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputNumerify(
                    components.PipelineFunctionNumerify{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionNumerifyIDNumerify,
                        Description: criblcontrolplanesdkgo.Pointer("relieve whose fixed provided mechanically pfft powerfully embarrassment"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaNumerify{
                            Depth: criblcontrolplanesdkgo.Pointer[int64](5),
                            IgnoreFields: []string{},
                            FilterExpr: criblcontrolplanesdkgo.Pointer(""),
                            Format: components.FunctionConfSchemaNumerifyFormatNone.ToPointer(),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesOTLPLogs

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesOTLPLogs" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "otlp-logs-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that normalizes and batches OTLP log events from OpenTelemetry sources"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputOtlpLogs(
                    components.PipelineFunctionOtlpLogs{
                        Filter: criblcontrolplanesdkgo.Pointer("__inputId=='open_telemetry:open_telemetry'"),
                        ID: components.PipelineFunctionOtlpLogsIDOtlpLogs,
                        Description: criblcontrolplanesdkgo.Pointer("stuff catalyst close mortally down"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.FunctionConfSchemaOtlpLogs{
                            DropNonLogEvents: criblcontrolplanesdkgo.Pointer(false),
                            BatchOTLPLogs: criblcontrolplanesdkgo.Pointer(true),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesOTLPMetrics

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesOTLPMetrics" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "otlp-metrics-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that converts dimensional metrics to OTLP format and batches them by resource attributes"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputOtlpMetrics(
                    components.PipelineFunctionOtlpMetrics{
                        Filter: criblcontrolplanesdkgo.Pointer("__inputId=='prometheus_rw:prom_rw_in'"),
                        ID: components.PipelineFunctionOtlpMetricsIDOtlpMetrics,
                        Description: criblcontrolplanesdkgo.Pointer("interestingly towards bowler now keenly"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.FunctionConfSchemaOtlpMetrics{
                            ResourceAttributePrefixes: []string{
                                "service",
                                "system",
                                "telemetry",
                                "k8s",
                                "cloud",
                                "host",
                                "process",
                            },
                            DropNonMetricEvents: criblcontrolplanesdkgo.Pointer(false),
                            OtlpVersion: components.OtlpVersionOptionsZeroDot10Dot0.ToPointer(),
                            BatchOTLPMetrics: criblcontrolplanesdkgo.Pointer(true),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesOTLPTraces

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesOTLPTraces" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "otlp-traces-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that normalizes and batches OTLP trace events from OpenTelemetry sources"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputOtlpTraces(
                    components.PipelineFunctionOtlpTraces{
                        Filter: criblcontrolplanesdkgo.Pointer("__inputId=='open_telemetry:open_telemetry'"),
                        ID: components.PipelineFunctionOtlpTracesIDOtlpTraces,
                        Description: criblcontrolplanesdkgo.Pointer("rectangular lest great hollow"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaOtlpTraces{
                            DropNonTraceEvents: criblcontrolplanesdkgo.Pointer(false),
                            OtlpVersion: components.OtlpVersionOptionsZeroDot10Dot0.ToPointer(),
                            BatchOTLPTraces: criblcontrolplanesdkgo.Pointer(true),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesParser

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesParser" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "parser-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that extracts fields from key-value pair formatted data"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputSerde(
                    components.PipelineFunctionSerde{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionSerdeIDSerde,
                        Description: criblcontrolplanesdkgo.Pointer("pish rudely colorfully around overdue ouch uh-huh gadzooks motor if"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionSerdeConf{
                            Mode: components.OperationModeExtract,
                            Type: components.TypeOptionsKvp,
                            DelimChar: "<value>",
                            QuoteChar: "<value>",
                            EscapeChar: "<value>",
                            NullValue: "<value>",
                            SrcField: criblcontrolplanesdkgo.Pointer("_raw"),
                            DstField: criblcontrolplanesdkgo.Pointer("<value>"),
                            CleanFields: false,
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesPublishMetrics

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesPublishMetrics" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "publish-metrics-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that extracts metrics from events and formats them for metrics aggregation platforms"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputPublishMetrics(
                    components.PipelineFunctionPublishMetrics{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionPublishMetricsIDPublishMetrics,
                        Description: criblcontrolplanesdkgo.Pointer("transcend yahoo creative till thoughtfully upside-down cap amongst"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaPublishMetrics{
                            Fields: []components.FunctionConfSchemaPublishMetricsField{
                                components.FunctionConfSchemaPublishMetricsField{
                                    InFieldName: "bytes",
                                    OutFieldExpr: criblcontrolplanesdkgo.Pointer("'metric_name.bytes'"),
                                    MetricType: components.FunctionConfSchemaPublishMetricsMetricTypeGauge,
                                },
                                components.FunctionConfSchemaPublishMetricsField{
                                    InFieldName: "packets",
                                    OutFieldExpr: criblcontrolplanesdkgo.Pointer("'metric_name.packets'"),
                                    MetricType: components.FunctionConfSchemaPublishMetricsMetricTypeGauge,
                                },
                            },
                            Overwrite: criblcontrolplanesdkgo.Pointer(false),
                            Dimensions: []string{
                                "action",
                                "interface_id",
                                "dstaddr",
                            },
                            RemoveMetrics: []string{},
                            RemoveDimensions: []string{},
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesRedis

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesRedis" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "redis-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that retrieves values from Redis using GET command"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputRedis(
                    components.PipelineFunctionRedis{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionRedisIDRedis,
                        Description: criblcontrolplanesdkgo.Pointer("gut valiantly step-mother political outrun likewise"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionRedisConf{
                            Commands: []components.Command{
                                components.Command{
                                    OutField: criblcontrolplanesdkgo.Pointer("cached_value"),
                                    Command: "get",
                                    KeyExpr: "'user_session'",
                                    ArgsExpr: criblcontrolplanesdkgo.Pointer(""),
                                },
                            },
                            DeploymentType: components.DeploymentTypeStandalone.ToPointer(),
                            AuthType: components.PipelineFunctionRedisAuthenticationMethodNone.ToPointer(),
                            MaxBlockSecs: criblcontrolplanesdkgo.Pointer[float64](60),
                            EnableClientSideCaching: criblcontrolplanesdkgo.Pointer(false),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesRegexExtract

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesRegexExtract" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "regex-extract-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that extracts structured fields from log data using regex patterns with named capture groups"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputRegexExtract(
                    components.PipelineFunctionRegexExtract{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionRegexExtractIDRegexExtract,
                        Description: criblcontrolplanesdkgo.Pointer("against next mmm"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.PipelineFunctionRegexExtractConf{
                            Regex: "/metric1=(?<metric1>\\d+)/",
                            RegexList: []components.PipelineFunctionRegexExtractRegexList{
                                components.PipelineFunctionRegexExtractRegexList{
                                    Regex: "<value>",
                                },
                            },
                            Source: criblcontrolplanesdkgo.Pointer("_raw"),
                            Iterations: criblcontrolplanesdkgo.Pointer[float64](100),
                            FieldNameExpression: criblcontrolplanesdkgo.Pointer("<value>"),
                            Overwrite: criblcontrolplanesdkgo.Pointer(false),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesRegexFilter

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesRegexFilter" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "regex-filter-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that filters out events matching specific regex patterns"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputRegexFilter(
                    components.PipelineFunctionRegexFilter{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionRegexFilterIDRegexFilter,
                        Description: criblcontrolplanesdkgo.Pointer("extroverted certify fatally flawless"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaRegexFilter{
                            Regex: criblcontrolplanesdkgo.Pointer("/Opera/"),
                            RegexList: []components.FunctionConfSchemaRegexFilterRegexList{
                                components.FunctionConfSchemaRegexFilterRegexList{
                                    Regex: "<value>",
                                },
                            },
                            Field: criblcontrolplanesdkgo.Pointer("_raw"),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesRename

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesRename" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "rename-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that renames fields using key-value pairs and expressions"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputRename(
                    components.PipelineFunctionRename{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionRenameIDRename,
                        Description: criblcontrolplanesdkgo.Pointer("pish poorly but ham"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaRename{
                            BaseFields: []string{},
                            Rename: []components.Rename{
                                components.Rename{
                                    CurrentName: "level",
                                    NewName: "LEVEL",
                                },
                            },
                            RenameExpr: criblcontrolplanesdkgo.Pointer("name.startsWith('out') ? name.toUpperCase() : name"),
                            WildcardDepth: criblcontrolplanesdkgo.Pointer[int64](5),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesRollupMetrics

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesRollupMetrics" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "rollup-metrics-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that consolidates high-frequency metrics into manageable time windows"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputRollupMetrics(
                    components.PipelineFunctionRollupMetrics{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionRollupMetricsIDRollupMetrics,
                        Description: criblcontrolplanesdkgo.Pointer("neatly how apostrophize"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaRollupMetrics{
                            Dimensions: []string{
                                "*",
                            },
                            TimeWindow: criblcontrolplanesdkgo.Pointer("30s"),
                            GaugeRollup: components.GaugeUpdateLast.ToPointer(),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesSNMPTrapSerialize

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesSNMPTrapSerialize" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "snmp-trap-serialize-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that serializes events into SNMP trap format for SNMP trap destinations"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputSnmpTrapSerialize(
                    components.PipelineFunctionSnmpTrapSerialize{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionSnmpTrapSerializeIDSnmpTrapSerialize,
                        Description: criblcontrolplanesdkgo.Pointer("steel ack colorfully vivid chilly cook suspiciously rich equal"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaSnmpTrapSerialize{
                            Strict: criblcontrolplanesdkgo.Pointer(true),
                            DropFailedEvents: criblcontrolplanesdkgo.Pointer(true),
                            V3User: &components.FunctionConfSchemaSnmpTrapSerializeV3User{
                                Name: criblcontrolplanesdkgo.Pointer("<value>"),
                                AuthProtocol: components.AuthenticationProtocolOptionsV3UserSha256.ToPointer(),
                                AuthKey: "<value>",
                                PrivProtocol: criblcontrolplanesdkgo.Pointer("<value>"),
                            },
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesSampling

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesSampling" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "sampling-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that samples events at specified rates based on filter criteria"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputSampling(
                    components.PipelineFunctionSampling{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionSamplingIDSampling,
                        Description: criblcontrolplanesdkgo.Pointer("amongst rewrite into damp late"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaSampling{
                            Rules: []components.FunctionConfSchemaSamplingRule{
                                components.FunctionConfSchemaSamplingRule{
                                    Filter: "__status == 200",
                                    Rate: 5,
                                },
                            },
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesSerialize

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesSerialize" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "serialize-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that serializes event fields into JSON format"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputSerialize(
                    components.PipelineFunctionSerialize{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionSerializeIDSerialize,
                        Description: criblcontrolplanesdkgo.Pointer("peony masterpiece gee rigidly muted now entwine shrilly ouch"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.PipelineFunctionSerializeConf{
                            Type: components.PipelineFunctionSerializeTypeJSON,
                            DelimChar: "<value>",
                            QuoteChar: "<value>",
                            EscapeChar: "<value>",
                            NullValue: "<value>",
                            Fields: []string{
                                "city",
                                "state",
                            },
                            SrcField: criblcontrolplanesdkgo.Pointer(""),
                            DstField: criblcontrolplanesdkgo.Pointer("_raw"),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesSuppress

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesSuppress" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "suppress-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that suppresses duplicate events based on a key expression"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputSuppress(
                    components.PipelineFunctionSuppress{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionSuppressIDSuppress,
                        Description: criblcontrolplanesdkgo.Pointer("clean nucleotidase spectacles plus whenever blah whereas reiterate optimistically"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionSuppressConf{
                            KeyExpr: "`${ip}:${port}`",
                            Allow: 1,
                            SuppressPeriodSec: 30,
                            DropEventsMode: criblcontrolplanesdkgo.Pointer(true),
                            MaxCacheSize: criblcontrolplanesdkgo.Pointer[float64](50000),
                            CacheIdleTimeoutPeriods: criblcontrolplanesdkgo.Pointer[float64](2),
                            NumEventsIdleTimeoutTrigger: criblcontrolplanesdkgo.Pointer[float64](10000),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesTee

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesTee" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "tee-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that sends events to a command via stdin for debugging purposes"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputTee(
                    components.PipelineFunctionTee{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionTeeIDTee,
                        Description: criblcontrolplanesdkgo.Pointer("aw mid taxicab dandelion ah mundane hmph whispered"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.PipelineFunctionTeeConf{
                            Command: "tee",
                            Args: []string{
                                "/opt/cribl/foo.log",
                            },
                            RestartOnExit: criblcontrolplanesdkgo.Pointer(true),
                            Env: map[string]string{

                            },
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesUnroll

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesUnroll" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "unroll-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that unrolls array fields into separate events"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputUnroll(
                    components.PipelineFunctionUnroll{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionUnrollIDUnroll,
                        Description: criblcontrolplanesdkgo.Pointer("second possible above tag what"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.PipelineFunctionUnrollConf{
                            SrcExpr: "_raw.split(/\\n/)",
                            DstField: "_raw",
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesXMLUnroll

<!-- UsageSnippet language="go" operationID="createPipeline" method="post" path="/pipelines" example="PipelineExamplesXMLUnroll" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Create(ctx, components.PipelineInput{
        ID: "xml-unroll-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that unrolls XML elements into separate events"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputXMLUnroll(
                    components.PipelineFunctionXMLUnroll{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionXMLUnrollIDXMLUnroll,
                        Description: criblcontrolplanesdkgo.Pointer("within limply describe eek questioningly anesthetize who wear"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionXMLUnrollConf{
                            Unroll: "^Parent\\.Child$",
                            Inherit: criblcontrolplanesdkgo.Pointer("^Parent\\.(myID|branchLocation)$"),
                            UnrollIdxField: criblcontrolplanesdkgo.Pointer("unroll_idx"),
                            Pretty: criblcontrolplanesdkgo.Pointer(false),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [components.PipelineInput](../../models/components/pipelineinput.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.CreatePipelineResponse](../../models/operations/createpipelineresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get the specified Pipeline.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPipelineById" method="get" path="/pipelines/{id}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Pipeline to get.              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetPipelineByIDResponse](../../models/operations/getpipelinebyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the specified Pipeline.</br></br>Provide a complete representation of the Pipeline that you want to update in the request body. This endpoint does not support partial updates. Cribl removes any omitted fields when updating the Pipeline.</br></br>Confirm that the configuration in your request body is correct before sending the request. If the configuration is incorrect, the updated Pipeline might not function as expected.

### Example Usage: PipelineExamplesAggregateMetrics

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesAggregateMetrics" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "aggregate-metrics-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that aggregates process metrics: CPU, memory, and bytes over time windows"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputAggregateMetrics(
                    components.PipelineFunctionAggregateMetrics{
                        Filter: criblcontrolplanesdkgo.Pointer("(_metric == 'proc.cpu_perc' || __criblMetrics[0].nameExpr.includes(\"'proc.cpu_perc'\")) || (_metric == 'proc.mem_perc' || __criblMetrics[0].nameExpr.includes(\"'proc.mem_perc'\")) || (_metric == 'proc.bytes_in' || __criblMetrics[0].nameExpr.includes(\"'proc.bytes_in'\"))"),
                        ID: components.PipelineFunctionAggregateMetricsIDAggregateMetrics,
                        Description: criblcontrolplanesdkgo.Pointer("however unless tangible pine"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionAggregateMetricsConf{
                            Passthrough: criblcontrolplanesdkgo.Pointer(false),
                            PreserveGroupBys: criblcontrolplanesdkgo.Pointer(false),
                            SufficientStatsOnly: criblcontrolplanesdkgo.Pointer(false),
                            Prefix: criblcontrolplanesdkgo.Pointer("<value>"),
                            TimeWindow: "10s",
                            Aggregations: []components.Aggregation{
                                components.Aggregation{
                                    MetricType: components.PipelineFunctionAggregateMetricsMetricTypeGauge,
                                    Agg: "avg(_value || proc.cpu_perc).as(proc.cpu_perc_avg)",
                                },
                                components.Aggregation{
                                    MetricType: components.PipelineFunctionAggregateMetricsMetricTypeGauge,
                                    Agg: "sum(_value || proc.mem_perc).as(proc.mem_perc_sum)",
                                },
                                components.Aggregation{
                                    MetricType: components.PipelineFunctionAggregateMetricsMetricTypeCounter,
                                    Agg: "count(_value || proc.bytes_in).as(proc.bytes_in_count)",
                                },
                            },
                            Groupbys: []string{
                                "proc",
                            },
                            FlushEventLimit: criblcontrolplanesdkgo.Pointer[float64](6008.27),
                            FlushMemLimit: criblcontrolplanesdkgo.Pointer("<value>"),
                            Cumulative: criblcontrolplanesdkgo.Pointer(false),
                            ShouldTreatDotsAsLiterals: criblcontrolplanesdkgo.Pointer(true),
                            Add: []components.PipelineFunctionAggregateMetricsAdd{
                                components.PipelineFunctionAggregateMetricsAdd{
                                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                                    Value: "<value>",
                                },
                            },
                            FlushOnInputClose: criblcontrolplanesdkgo.Pointer(true),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesAggregations

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesAggregations" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "aggregation-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that aggregates rejected bytes grouped by source address every 10 seconds"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputAggregation(
                    components.PipelineFunctionAggregation{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionAggregationIDAggregation,
                        Description: criblcontrolplanesdkgo.Pointer("ick whenever pinstripe indeed scope into wide-eyed"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionAggregationConf{
                            Passthrough: criblcontrolplanesdkgo.Pointer(false),
                            PreserveGroupBys: criblcontrolplanesdkgo.Pointer(false),
                            SufficientStatsOnly: criblcontrolplanesdkgo.Pointer(false),
                            MetricsMode: criblcontrolplanesdkgo.Pointer(false),
                            Prefix: criblcontrolplanesdkgo.Pointer("<value>"),
                            TimeWindow: "10s",
                            Aggregations: []string{
                                "sum(bytes).where(action==\"REJECT\").as(TotalBytes)",
                            },
                            Groupbys: []string{
                                "srcaddr",
                            },
                            FlushEventLimit: criblcontrolplanesdkgo.Pointer[float64](6263.18),
                            FlushMemLimit: criblcontrolplanesdkgo.Pointer("<value>"),
                            Cumulative: criblcontrolplanesdkgo.Pointer(false),
                            SearchAggMode: criblcontrolplanesdkgo.Pointer("<value>"),
                            Add: []components.ItemsTypeAdd{
                                components.ItemsTypeAdd{
                                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                                    Value: "<value>",
                                },
                            },
                            ShouldTreatDotsAsLiterals: criblcontrolplanesdkgo.Pointer(false),
                            FlushOnInputClose: criblcontrolplanesdkgo.Pointer(true),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesAutoTimestamp

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesAutoTimestamp" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "auto-timestamp-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that extracts timestamps from event data using auto timestamp function"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputAutoTimestamp(
                    components.PipelineFunctionAutoTimestamp{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionAutoTimestampIDAutoTimestamp,
                        Description: criblcontrolplanesdkgo.Pointer("what defiantly switch aha where slake but lighthearted"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.FunctionConfSchemaAutoTimestamp{
                            SrcField: criblcontrolplanesdkgo.Pointer("_raw"),
                            DstField: criblcontrolplanesdkgo.Pointer("_time"),
                            DefaultTimezone: criblcontrolplanesdkgo.Pointer("local"),
                            TimeExpression: criblcontrolplanesdkgo.Pointer("time.getTime() / 1000"),
                            Offset: criblcontrolplanesdkgo.Pointer[float64](0),
                            MaxLen: criblcontrolplanesdkgo.Pointer[float64](150),
                            DefaultTime: components.DefaultTimeNow.ToPointer(),
                            LatestDateAllowed: criblcontrolplanesdkgo.Pointer("+1week"),
                            Spacer: criblcontrolplanesdkgo.Pointer("<value>"),
                            EarliestDateAllowed: criblcontrolplanesdkgo.Pointer("-420weeks"),
                            Timestamps: []components.Timestamp{
                                components.Timestamp{
                                    Regex: "/(\\d{1,2}\\/\\d{2}\\/\\d{4}\\s\\d{1,2}:\\d{2}:\\d{2}\\s\\w{2})/",
                                    Strptime: "%Y-%m-%d %H:%M:%S",
                                },
                            },
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesCEFSerializer

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesCEFSerializer" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "cef-serializer-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that formats events in CEF format with custom header and extension fields"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputCef(
                    components.PipelineFunctionCefInput{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionCefIDCef,
                        Description: criblcontrolplanesdkgo.Pointer("ah endow whoever notwithstanding now mmm homely scorn"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.FunctionConfSchemaCefInput{
                            OutputField: criblcontrolplanesdkgo.Pointer("_raw"),
                            Header: []components.HeaderInput{
                                components.HeaderInput{
                                    Value: "'CEF:0'",
                                },
                                components.HeaderInput{
                                    Value: "'Cribl'",
                                },
                                components.HeaderInput{
                                    Value: "'Cribl'",
                                },
                                components.HeaderInput{
                                    Value: "C.version",
                                },
                                components.HeaderInput{
                                    Value: "420",
                                },
                                components.HeaderInput{
                                    Value: "'Cribl Event'",
                                },
                                components.HeaderInput{
                                    Value: "6",
                                },
                            },
                            Extension: []components.Extension{
                                components.Extension{
                                    Name: "c6a1Label",
                                    Value: "'Colorado_Ext_Bldg7'",
                                },
                            },
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesChain

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesChain" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "chain-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that chains to another pipeline for sequential data processing"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputChain(
                    components.PipelineFunctionChain{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionChainIDChain,
                        Description: criblcontrolplanesdkgo.Pointer("into jaggedly truly"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionChainConf{
                            Processor: "prometheus_metrics",
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesClone

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesClone" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "clone-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that creates cloned events with additional fields for comparison or routing"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputClone(
                    components.PipelineFunctionClone{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionCloneIDClone,
                        Description: criblcontrolplanesdkgo.Pointer("self-assured meaningfully phooey consequently offset woot abnormally"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.FunctionConfSchemaClone{
                            Clones: []map[string]string{
                                map[string]string{
                                    "env": "staging",
                                },
                                map[string]string{
                                    "index": "clones",
                                },
                            },
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesComment

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesComment" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "comment-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline with comment function for documentation"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputComment(
                    components.PipelineFunctionComment{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionCommentIDComment,
                        Description: criblcontrolplanesdkgo.Pointer("lend however blindly sugary into except selfishly"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.FunctionConfSchemaComment{
                            Comment: criblcontrolplanesdkgo.Pointer("This function processes security events and enriches them with DNS lookups"),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesDNSLookup

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesDNSLookup" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "dns-lookup-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that performs DNS lookups to resolve hostnames and IP addresses"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputDNSLookup(
                    components.PipelineFunctionDNSLookup{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionDNSLookupIDDNSLookup,
                        Description: criblcontrolplanesdkgo.Pointer("psst penalise couch lest mid yowza versus chime until near"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaDNSLookup{
                            DNSLookupFields: []components.DNSLookupField{
                                components.DNSLookupField{
                                    InFieldName: criblcontrolplanesdkgo.Pointer("hostname"),
                                    ResourceRecordType: components.ResourceRecordTypeA.ToPointer(),
                                    OutFieldName: criblcontrolplanesdkgo.Pointer("hostname_ip"),
                                },
                            },
                            ReverseLookupFields: []components.ReverseLookupField{
                                components.ReverseLookupField{
                                    InFieldName: criblcontrolplanesdkgo.Pointer("src_ip"),
                                    OutFieldName: criblcontrolplanesdkgo.Pointer("src_hostname"),
                                },
                            },
                            DNSServers: []string{
                                "<value 1>",
                                "<value 2>",
                            },
                            CacheTTL: criblcontrolplanesdkgo.Pointer[float64](30),
                            MaxCacheSize: criblcontrolplanesdkgo.Pointer[float64](5000),
                            UseResolvConf: criblcontrolplanesdkgo.Pointer(false),
                            LookupFallback: criblcontrolplanesdkgo.Pointer(false),
                            DomainOverrides: []string{
                                "<value 1>",
                            },
                            LookupFailLogLevel: components.LogLevelForFailedLookupsError.ToPointer(),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesDrop

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesDrop" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "drop-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that drops events containing success messages"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputDrop(
                    components.PipelineFunctionDrop{
                        Filter: criblcontrolplanesdkgo.Pointer("_raw.search(/success/i)>=0"),
                        ID: components.PipelineFunctionDropIDDrop,
                        Description: criblcontrolplanesdkgo.Pointer("briefly equatorial ha"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaDrop{},
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesDropDimensions

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesDropDimensions" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "drop-dimensions-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that drops specified dimensions from metrics to reduce cardinality"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputDropDimensions(
                    components.PipelineFunctionDropDimensions{
                        Filter: criblcontrolplanesdkgo.Pointer("(_metric == 'proc.cpu_perc' || __criblMetrics[0].nameExpr.includes(\"'proc.cpu_perc'\")) && (__criblMetrics[0].dims.includes(\"proc\"))"),
                        ID: components.PipelineFunctionDropDimensionsIDDropDimensions,
                        Description: criblcontrolplanesdkgo.Pointer("overplay zowie so fat utilized bulky hence bashfully"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionDropDimensionsConf{
                            TimeWindow: "10s",
                            DropDimensions: []string{
                                "proc",
                                "pie",
                                "unit",
                            },
                            FlushOnInputClose: criblcontrolplanesdkgo.Pointer(true),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesDynamicSampling

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesDynamicSampling" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "dynamic-sampling-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that dynamically samples events based on volume using square root mode"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputDynamicSampling(
                    components.PipelineFunctionDynamicSampling{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionDynamicSamplingIDDynamicSampling,
                        Description: criblcontrolplanesdkgo.Pointer("fooey although circa oof veg nor till likewise"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.PipelineFunctionDynamicSamplingConf{
                            Mode: components.SampleModeSqrt,
                            KeyExpr: "`${domain}:${httpCode}`",
                            SamplePeriod: criblcontrolplanesdkgo.Pointer[float64](20),
                            MinEvents: criblcontrolplanesdkgo.Pointer[float64](3),
                            MaxSampleRate: criblcontrolplanesdkgo.Pointer[float64](3),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesEmpty

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesEmpty" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "empty-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer(""),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{},
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesEval

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesEval" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "eval-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that evaluates JavaScript expressions to add, modify, and remove fields"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputEval(
                    components.PipelineFunctionEval{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionEvalIDEval,
                        Description: criblcontrolplanesdkgo.Pointer("yahoo gah athwart inasmuch circa overcook cricket seriously yippee now"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaEval{
                            Add: []components.FunctionConfSchemaEvalAdd{
                                components.FunctionConfSchemaEvalAdd{
                                    Name: criblcontrolplanesdkgo.Pointer("action"),
                                    Value: "login == 'error' ? 'blocked' : action",
                                    Disabled: criblcontrolplanesdkgo.Pointer(true),
                                },
                                components.FunctionConfSchemaEvalAdd{
                                    Name: criblcontrolplanesdkgo.Pointer("myTags"),
                                    Value: "login == 'error' ? [...myTags, 'error'] : myTags",
                                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                                },
                            },
                            Keep: []string{
                                "host",
                                "source",
                                "action",
                                "myTags",
                            },
                            Remove: []string{
                                "identification",
                            },
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesEventBreaker

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesEventBreaker" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "event-breaker-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that breaks large event streams into discrete events using regex"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputEventBreaker(
                    components.PipelineFunctionEventBreaker{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionEventBreakerIDEventBreaker,
                        Description: criblcontrolplanesdkgo.Pointer("refer coaxingly dental"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionEventBreakerConf{
                            ExistingOrNew: components.ExistingOrNewNew,
                            ShouldMarkCriblBreaker: criblcontrolplanesdkgo.Pointer(true),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesFlatten

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesFlatten" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "flatten-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that flattens nested JSON structures into top-level fields"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputFlatten(
                    components.PipelineFunctionFlatten{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionFlattenIDFlatten,
                        Description: criblcontrolplanesdkgo.Pointer("privilege among how round freezing untidy"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaFlatten{
                            Fields: []string{},
                            Prefix: criblcontrolplanesdkgo.Pointer(""),
                            Depth: criblcontrolplanesdkgo.Pointer[float64](5),
                            Delimiter: criblcontrolplanesdkgo.Pointer("_"),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesFoldKeys

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesFoldKeys" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "fold-keys-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that transforms flat field names with separators into nested structures"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputFoldkeys(
                    components.PipelineFunctionFoldkeys{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionFoldkeysIDFoldkeys,
                        Description: criblcontrolplanesdkgo.Pointer("wherever atop modulo whenever incidentally wherever char upside-down numeracy gosh"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.FunctionConfSchemaFoldkeys{
                            DeleteOriginal: criblcontrolplanesdkgo.Pointer(true),
                            Separator: criblcontrolplanesdkgo.Pointer("_"),
                            SelectionRegExp: criblcontrolplanesdkgo.Pointer("^data"),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesGeoIP

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesGeoIP" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "geoip-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that enriches events with geolocation data from IP addresses"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputGeoip(
                    components.PipelineFunctionGeoip{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionGeoipIDGeoip,
                        Description: criblcontrolplanesdkgo.Pointer("hasty oh bah officially"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionGeoipConf{
                            File: "GeoLite2-City.mmdb",
                            InField: criblcontrolplanesdkgo.Pointer("ip"),
                            OutField: criblcontrolplanesdkgo.Pointer("geoip"),
                            AdditionalFields: []components.AdditionalField{
                                components.AdditionalField{
                                    ExtraInField: "src_ip",
                                    ExtraOutField: "src_geoip",
                                },
                            },
                            OutFieldMappings: &components.OutputFieldMappings{},
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesGrok

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesGrok" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "grok-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that extracts structured fields from log data using Grok patterns"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputGrok(
                    components.PipelineFunctionGrok{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionGrokIDGrok,
                        Description: criblcontrolplanesdkgo.Pointer("disposer gosh but chairperson foodstuffs excepting vanish as unlike planula"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.PipelineFunctionGrokConf{
                            Pattern: "%{TIMESTAMP_ISO8601:event_time} %{LOGLEVEL:log_level} %{GREEDYDATA:log_message}",
                            PatternList: []components.PatternList{},
                            Source: criblcontrolplanesdkgo.Pointer("_raw"),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesGuard

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesGuard" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "guard-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that scans for sensitive data and applies mitigation expressions"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputSensitiveDataScanner(
                    components.PipelineFunctionSensitiveDataScanner{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionSensitiveDataScannerIDSensitiveDataScanner,
                        Description: criblcontrolplanesdkgo.Pointer("outnumber lazily gah fondly crocodile off"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionSensitiveDataScannerConf{
                            Rules: []components.PipelineFunctionSensitiveDataScannerRule{
                                components.PipelineFunctionSensitiveDataScannerRule{
                                    RulesetID: "Finance_Global",
                                    ReplaceExpr: "'REDACTED'",
                                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                                },
                            },
                            Fields: []string{
                                "_raw",
                            },
                            ExcludeFields: []string{},
                            Flags: []components.Flag{
                                components.Flag{
                                    Name: criblcontrolplanesdkgo.Pointer("_sensitive"),
                                    Value: "true",
                                },
                            },
                            IncludeDetectedRules: criblcontrolplanesdkgo.Pointer(true),
                            BackgroundDetection: criblcontrolplanesdkgo.Pointer(true),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesJSONUnroll

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesJSONUnroll" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "json-unroll-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that unrolls JSON arrays into individual events while retaining parent fields"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputJSONUnroll(
                    components.PipelineFunctionJSONUnroll{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionJSONUnrollIDJSONUnroll,
                        Description: criblcontrolplanesdkgo.Pointer("sushi ack legitimize through transparency loftily"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.PipelineFunctionJSONUnrollConf{
                            Path: "allCars",
                            Name: criblcontrolplanesdkgo.Pointer("cars"),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesLookup

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesLookup" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "lookup-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that enriches events with location data from IP address lookups"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputLookup(
                    components.PipelineFunctionLookup{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionLookupIDLookup,
                        Description: criblcontrolplanesdkgo.Pointer("since communicate now whenever warped whose spice"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.PipelineFunctionLookupConf{
                            File: "ip_locations.csv",
                            DbLookup: criblcontrolplanesdkgo.Pointer(false),
                            MatchMode: "exact",
                            MatchType: "<value>",
                            ReloadPeriodSec: -1,
                            InFields: []components.InField{
                                components.InField{
                                    EventField: "destination_ip",
                                    LookupField: criblcontrolplanesdkgo.Pointer("ip"),
                                },
                            },
                            OutFields: []components.OutField{
                                components.OutField{
                                    LookupField: "location",
                                    EventField: criblcontrolplanesdkgo.Pointer("location"),
                                    DefaultValue: criblcontrolplanesdkgo.Pointer("Unknown"),
                                },
                            },
                            AddToEvent: criblcontrolplanesdkgo.Pointer(false),
                            IgnoreCase: false,
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesMask

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesMask" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "mask-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that masks Social Security numbers and other sensitive data"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputMask(
                    components.PipelineFunctionMask{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionMaskIDMask,
                        Description: criblcontrolplanesdkgo.Pointer("woot malfunction gnash whether athwart inasmuch"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.PipelineFunctionMaskConf{
                            Rules: []components.PipelineFunctionMaskRule{
                                components.PipelineFunctionMaskRule{
                                    MatchRegex: "/(social=)(\\d+)/",
                                    ReplaceExpr: "`${g1}${C.Mask.md5(g2)}`",
                                    Disabled: criblcontrolplanesdkgo.Pointer(false),
                                },
                            },
                            Fields: []string{
                                "_raw",
                            },
                            Depth: criblcontrolplanesdkgo.Pointer[int64](5),
                            Flags: []components.ItemsTypeAdd{
                                components.ItemsTypeAdd{
                                    Name: criblcontrolplanesdkgo.Pointer("<value>"),
                                    Value: "<value>",
                                },
                            },
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesNumerify

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesNumerify" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "numerify-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that converts string numbers to numeric type for mathematical operations"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputNumerify(
                    components.PipelineFunctionNumerify{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionNumerifyIDNumerify,
                        Description: criblcontrolplanesdkgo.Pointer("indeed greedy supposing"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.FunctionConfSchemaNumerify{
                            Depth: criblcontrolplanesdkgo.Pointer[int64](5),
                            IgnoreFields: []string{},
                            FilterExpr: criblcontrolplanesdkgo.Pointer(""),
                            Format: components.FunctionConfSchemaNumerifyFormatNone.ToPointer(),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesOTLPLogs

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesOTLPLogs" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "otlp-logs-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that normalizes and batches OTLP log events from OpenTelemetry sources"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputOtlpLogs(
                    components.PipelineFunctionOtlpLogs{
                        Filter: criblcontrolplanesdkgo.Pointer("__inputId=='open_telemetry:open_telemetry'"),
                        ID: components.PipelineFunctionOtlpLogsIDOtlpLogs,
                        Description: criblcontrolplanesdkgo.Pointer("skyscraper er thankfully uh-huh rotating lest shoulder"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaOtlpLogs{
                            DropNonLogEvents: criblcontrolplanesdkgo.Pointer(false),
                            BatchOTLPLogs: criblcontrolplanesdkgo.Pointer(true),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesOTLPMetrics

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesOTLPMetrics" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "otlp-metrics-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that converts dimensional metrics to OTLP format and batches them by resource attributes"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputOtlpMetrics(
                    components.PipelineFunctionOtlpMetrics{
                        Filter: criblcontrolplanesdkgo.Pointer("__inputId=='prometheus_rw:prom_rw_in'"),
                        ID: components.PipelineFunctionOtlpMetricsIDOtlpMetrics,
                        Description: criblcontrolplanesdkgo.Pointer("modulo mismatch enchanted"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaOtlpMetrics{
                            ResourceAttributePrefixes: []string{
                                "service",
                                "system",
                                "telemetry",
                                "k8s",
                                "cloud",
                                "host",
                                "process",
                            },
                            DropNonMetricEvents: criblcontrolplanesdkgo.Pointer(false),
                            OtlpVersion: components.OtlpVersionOptionsZeroDot10Dot0.ToPointer(),
                            BatchOTLPMetrics: criblcontrolplanesdkgo.Pointer(true),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesOTLPTraces

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesOTLPTraces" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "otlp-traces-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that normalizes and batches OTLP trace events from OpenTelemetry sources"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputOtlpTraces(
                    components.PipelineFunctionOtlpTraces{
                        Filter: criblcontrolplanesdkgo.Pointer("__inputId=='open_telemetry:open_telemetry'"),
                        ID: components.PipelineFunctionOtlpTracesIDOtlpTraces,
                        Description: criblcontrolplanesdkgo.Pointer("obstruct beyond suddenly"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.FunctionConfSchemaOtlpTraces{
                            DropNonTraceEvents: criblcontrolplanesdkgo.Pointer(false),
                            OtlpVersion: components.OtlpVersionOptionsZeroDot10Dot0.ToPointer(),
                            BatchOTLPTraces: criblcontrolplanesdkgo.Pointer(true),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesParser

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesParser" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "parser-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that extracts fields from key-value pair formatted data"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputSerde(
                    components.PipelineFunctionSerde{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionSerdeIDSerde,
                        Description: criblcontrolplanesdkgo.Pointer("acidly sympathetically accept calmly sophisticated provided through oily perspire"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionSerdeConf{
                            Mode: components.OperationModeExtract,
                            Type: components.TypeOptionsKvp,
                            DelimChar: "<value>",
                            QuoteChar: "<value>",
                            EscapeChar: "<value>",
                            NullValue: "<value>",
                            SrcField: criblcontrolplanesdkgo.Pointer("_raw"),
                            DstField: criblcontrolplanesdkgo.Pointer("<value>"),
                            CleanFields: false,
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesPublishMetrics

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesPublishMetrics" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "publish-metrics-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that extracts metrics from events and formats them for metrics aggregation platforms"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputPublishMetrics(
                    components.PipelineFunctionPublishMetrics{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionPublishMetricsIDPublishMetrics,
                        Description: criblcontrolplanesdkgo.Pointer("instead after mediocre tabulate cultivated"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.FunctionConfSchemaPublishMetrics{
                            Fields: []components.FunctionConfSchemaPublishMetricsField{
                                components.FunctionConfSchemaPublishMetricsField{
                                    InFieldName: "bytes",
                                    OutFieldExpr: criblcontrolplanesdkgo.Pointer("'metric_name.bytes'"),
                                    MetricType: components.FunctionConfSchemaPublishMetricsMetricTypeGauge,
                                },
                                components.FunctionConfSchemaPublishMetricsField{
                                    InFieldName: "packets",
                                    OutFieldExpr: criblcontrolplanesdkgo.Pointer("'metric_name.packets'"),
                                    MetricType: components.FunctionConfSchemaPublishMetricsMetricTypeGauge,
                                },
                            },
                            Overwrite: criblcontrolplanesdkgo.Pointer(false),
                            Dimensions: []string{
                                "action",
                                "interface_id",
                                "dstaddr",
                            },
                            RemoveMetrics: []string{},
                            RemoveDimensions: []string{},
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesRedis

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesRedis" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "redis-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that retrieves values from Redis using GET command"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputRedis(
                    components.PipelineFunctionRedis{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionRedisIDRedis,
                        Description: criblcontrolplanesdkgo.Pointer("mummify yum whose before however er unfit intermesh reasoning"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.PipelineFunctionRedisConf{
                            Commands: []components.Command{
                                components.Command{
                                    OutField: criblcontrolplanesdkgo.Pointer("cached_value"),
                                    Command: "get",
                                    KeyExpr: "'user_session'",
                                    ArgsExpr: criblcontrolplanesdkgo.Pointer(""),
                                },
                            },
                            DeploymentType: components.DeploymentTypeStandalone.ToPointer(),
                            AuthType: components.PipelineFunctionRedisAuthenticationMethodNone.ToPointer(),
                            MaxBlockSecs: criblcontrolplanesdkgo.Pointer[float64](60),
                            EnableClientSideCaching: criblcontrolplanesdkgo.Pointer(false),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesRegexExtract

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesRegexExtract" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "regex-extract-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that extracts structured fields from log data using regex patterns with named capture groups"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputRegexExtract(
                    components.PipelineFunctionRegexExtract{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionRegexExtractIDRegexExtract,
                        Description: criblcontrolplanesdkgo.Pointer("dependency huzzah gurn invite"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.PipelineFunctionRegexExtractConf{
                            Regex: "/metric1=(?<metric1>\\d+)/",
                            RegexList: []components.PipelineFunctionRegexExtractRegexList{
                                components.PipelineFunctionRegexExtractRegexList{
                                    Regex: "<value>",
                                },
                            },
                            Source: criblcontrolplanesdkgo.Pointer("_raw"),
                            Iterations: criblcontrolplanesdkgo.Pointer[float64](100),
                            FieldNameExpression: criblcontrolplanesdkgo.Pointer("<value>"),
                            Overwrite: criblcontrolplanesdkgo.Pointer(false),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesRegexFilter

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesRegexFilter" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "regex-filter-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that filters out events matching specific regex patterns"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputRegexFilter(
                    components.PipelineFunctionRegexFilter{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionRegexFilterIDRegexFilter,
                        Description: criblcontrolplanesdkgo.Pointer("mmm once duh respray"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.FunctionConfSchemaRegexFilter{
                            Regex: criblcontrolplanesdkgo.Pointer("/Opera/"),
                            RegexList: []components.FunctionConfSchemaRegexFilterRegexList{
                                components.FunctionConfSchemaRegexFilterRegexList{
                                    Regex: "<value>",
                                },
                            },
                            Field: criblcontrolplanesdkgo.Pointer("_raw"),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesRename

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesRename" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "rename-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that renames fields using key-value pairs and expressions"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputRename(
                    components.PipelineFunctionRename{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionRenameIDRename,
                        Description: criblcontrolplanesdkgo.Pointer("per chase suspiciously vary shimmering through unlearn gallery waft yet"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaRename{
                            BaseFields: []string{},
                            Rename: []components.Rename{
                                components.Rename{
                                    CurrentName: "level",
                                    NewName: "LEVEL",
                                },
                            },
                            RenameExpr: criblcontrolplanesdkgo.Pointer("name.startsWith('out') ? name.toUpperCase() : name"),
                            WildcardDepth: criblcontrolplanesdkgo.Pointer[int64](5),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesRollupMetrics

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesRollupMetrics" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "rollup-metrics-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that consolidates high-frequency metrics into manageable time windows"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputRollupMetrics(
                    components.PipelineFunctionRollupMetrics{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionRollupMetricsIDRollupMetrics,
                        Description: criblcontrolplanesdkgo.Pointer("bashfully whispered cheerfully stupendous dandelion hmph lest alongside perfection homely"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaRollupMetrics{
                            Dimensions: []string{
                                "*",
                            },
                            TimeWindow: criblcontrolplanesdkgo.Pointer("30s"),
                            GaugeRollup: components.GaugeUpdateLast.ToPointer(),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesSNMPTrapSerialize

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesSNMPTrapSerialize" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "snmp-trap-serialize-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that serializes events into SNMP trap format for SNMP trap destinations"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputSnmpTrapSerialize(
                    components.PipelineFunctionSnmpTrapSerialize{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionSnmpTrapSerializeIDSnmpTrapSerialize,
                        Description: criblcontrolplanesdkgo.Pointer("aha ew lender mosh"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaSnmpTrapSerialize{
                            Strict: criblcontrolplanesdkgo.Pointer(true),
                            DropFailedEvents: criblcontrolplanesdkgo.Pointer(true),
                            V3User: &components.FunctionConfSchemaSnmpTrapSerializeV3User{
                                Name: criblcontrolplanesdkgo.Pointer("<value>"),
                                AuthProtocol: components.AuthenticationProtocolOptionsV3UserMd5.ToPointer(),
                                AuthKey: "<value>",
                                PrivProtocol: criblcontrolplanesdkgo.Pointer("<value>"),
                            },
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesSampling

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesSampling" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "sampling-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that samples events at specified rates based on filter criteria"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputSampling(
                    components.PipelineFunctionSampling{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionSamplingIDSampling,
                        Description: criblcontrolplanesdkgo.Pointer("even dock modulo dreary whoever ew"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.FunctionConfSchemaSampling{
                            Rules: []components.FunctionConfSchemaSamplingRule{
                                components.FunctionConfSchemaSamplingRule{
                                    Filter: "__status == 200",
                                    Rate: 5,
                                },
                            },
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesSerialize

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesSerialize" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "serialize-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that serializes event fields into JSON format"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputSerialize(
                    components.PipelineFunctionSerialize{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionSerializeIDSerialize,
                        Description: criblcontrolplanesdkgo.Pointer("whoa ew grandiose meal that"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionSerializeConf{
                            Type: components.PipelineFunctionSerializeTypeJSON,
                            DelimChar: "<value>",
                            QuoteChar: "<value>",
                            EscapeChar: "<value>",
                            NullValue: "<value>",
                            Fields: []string{
                                "city",
                                "state",
                            },
                            SrcField: criblcontrolplanesdkgo.Pointer(""),
                            DstField: criblcontrolplanesdkgo.Pointer("_raw"),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesSuppress

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesSuppress" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "suppress-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that suppresses duplicate events based on a key expression"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputSuppress(
                    components.PipelineFunctionSuppress{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionSuppressIDSuppress,
                        Description: criblcontrolplanesdkgo.Pointer("ack alongside famously brr after punctuation crowded supposing swath"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.PipelineFunctionSuppressConf{
                            KeyExpr: "`${ip}:${port}`",
                            Allow: 1,
                            SuppressPeriodSec: 30,
                            DropEventsMode: criblcontrolplanesdkgo.Pointer(true),
                            MaxCacheSize: criblcontrolplanesdkgo.Pointer[float64](50000),
                            CacheIdleTimeoutPeriods: criblcontrolplanesdkgo.Pointer[float64](2),
                            NumEventsIdleTimeoutTrigger: criblcontrolplanesdkgo.Pointer[float64](10000),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesTee

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesTee" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "tee-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that sends events to a command via stdin for debugging purposes"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputTee(
                    components.PipelineFunctionTee{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionTeeIDTee,
                        Description: criblcontrolplanesdkgo.Pointer("delirious athletic remark chairperson order"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionTeeConf{
                            Command: "tee",
                            Args: []string{
                                "/opt/cribl/foo.log",
                            },
                            RestartOnExit: criblcontrolplanesdkgo.Pointer(true),
                            Env: map[string]string{

                            },
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesUnroll

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesUnroll" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "unroll-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that unrolls array fields into separate events"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputUnroll(
                    components.PipelineFunctionUnroll{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionUnrollIDUnroll,
                        Description: criblcontrolplanesdkgo.Pointer("eminent safely sonar verve sedately like yippee swing which rudely"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionUnrollConf{
                            SrcExpr: "_raw.split(/\\n/)",
                            DstField: "_raw",
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```
### Example Usage: PipelineExamplesXMLUnroll

<!-- UsageSnippet language="go" operationID="updatePipelineById" method="patch" path="/pipelines/{id}" example="PipelineExamplesXMLUnroll" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Update(ctx, "<id>", components.PipelineInput{
        ID: "xml-unroll-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer("Pipeline that unrolls XML elements into separate events"),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{
                components.CreatePipelineFunctionConfInputXMLUnroll(
                    components.PipelineFunctionXMLUnroll{
                        Filter: criblcontrolplanesdkgo.Pointer("true"),
                        ID: components.PipelineFunctionXMLUnrollIDXMLUnroll,
                        Description: criblcontrolplanesdkgo.Pointer("tousle jell preclude rationalize into"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.PipelineFunctionXMLUnrollConf{
                            Unroll: "^Parent\\.Child$",
                            Inherit: criblcontrolplanesdkgo.Pointer("^Parent\\.(myID|branchLocation)$"),
                            UnrollIdxField: criblcontrolplanesdkgo.Pointer("unroll_idx"),
                            Pretty: criblcontrolplanesdkgo.Pointer(false),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.AdditionalPropertiesTypePipelineConfGroups{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `id`                                                                 | *string*                                                             | :heavy_check_mark:                                                   | The <code>id</code> of the Pipeline to update.                       |
| `pipeline`                                                           | [components.PipelineInput](../../models/components/pipelineinput.md) | :heavy_check_mark:                                                   | Pipeline object                                                      |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.UpdatePipelineByIDResponse](../../models/operations/updatepipelinebyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete the specified Pipeline.

### Example Usage

<!-- UsageSnippet language="go" operationID="deletePipelineById" method="delete" path="/pipelines/{id}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Pipelines.Delete(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPipeline != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Pipeline to delete.           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeletePipelineByIDResponse](../../models/operations/deletepipelinebyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |