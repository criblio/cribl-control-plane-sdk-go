# Packs.Pipelines

## Overview

### Available Operations

* [Create](#create) - Create a Pipeline within a Pack
* [List](#list) - List all Pipelines within a Pack
* [Delete](#delete) - Delete a Pipeline within a Pack
* [Get](#get) - Get a Pipeline within a Pack
* [Update](#update) - Update a Pipeline within a Pack

## Create

Create a new Pipeline within the specified Pack.

### Example Usage: PipelineExamplesAggregateMetrics

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesAggregateMetrics" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("case afore bashfully smooth"),
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
                            FlushEventLimit: criblcontrolplanesdkgo.Pointer[float64](7595.29),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesAggregations" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("eek harvest horse ragged till ack"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
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
                            FlushEventLimit: criblcontrolplanesdkgo.Pointer[float64](7135.24),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesAutoTimestamp" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("garage whoa meanwhile blind er apropos"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesCEFSerializer" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("nice ick which brr um pro times mammoth"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesChain" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("quip excepting psst jungle configuration hence"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionChainConf{
                            Processor: "prometheus_metrics",
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesClone" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("pfft rout battle"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesComment" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("knavishly accentuate helplessly concrete factorize drat meaty quaff essence"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaComment{
                            Comment: criblcontrolplanesdkgo.Pointer("This function processes security events and enriches them with DNS lookups"),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesDNSLookup" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("near greatly midst kinase concerning gratefully whose"),
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
                                "<value 2>",
                                "<value 3>",
                            },
                            LookupFailLogLevel: components.LogLevelForFailedLookupsError.ToPointer(),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesDrop" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("cassava interestingly boss commandeer"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.FunctionConfSchemaDrop{},
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesDropDimensions" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("vacantly boohoo where gorgeous notarize"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesDynamicSampling" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("meanwhile starch cheerful what in whose soft"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesEmpty" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
        ID: "empty-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer(""),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{},
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesEval" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("repeatedly upon derby gently suffocate inquisitively when oh along whereas"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesEventBreaker" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("about gleaming beside"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionEventBreakerConf{
                            ExistingOrNew: components.ExistingOrNewNew,
                            ShouldMarkCriblBreaker: criblcontrolplanesdkgo.Pointer(true),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesFlatten" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("helpless mmm oh pace meanwhile below above"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesFoldKeys" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("concerning mothball worth during"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesGeoIP" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("until jaunty mid jazz tame trial"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesGrok" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("usually natural impeccable official louse custody"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesGuard" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("under morning worriedly"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
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
                            BackgroundDetection: criblcontrolplanesdkgo.Pointer(false),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesJSONUnroll" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("beneath flickering as"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.PipelineFunctionJSONUnrollConf{
                            Path: "allCars",
                            Name: criblcontrolplanesdkgo.Pointer("cars"),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesLookup" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("shampoo why internalize masquerade boohoo right talkative even until"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesMask" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("list quirkily yahoo cantaloupe perfectly outrageous carefully lucky"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesNumerify" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("somber furthermore crumble during spherical geez but"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesOTLPLogs" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("abaft who catalyze cuckoo past woeful woot"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaOtlpLogs{
                            DropNonLogEvents: criblcontrolplanesdkgo.Pointer(false),
                            BatchOTLPLogs: criblcontrolplanesdkgo.Pointer(true),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesOTLPMetrics" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("eek aching fort qua likewise brr"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesOTLPTraces" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("provided spirited whenever in like knowledgeably however"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesParser" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("madly aw fast sympathetically swear hence excluding"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesPublishMetrics" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("ack scorpion that sad than since backbone yippee"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesRedis" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("whale minor weekly qua writhing gee continually"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesRegexExtract" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("blah inasmuch pearl till than rust the sediment stir-fry give"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesRegexFilter" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("where sans surprised pace ah"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesRename" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("topsail harmful although telescope woot why pro insistent dimly blissfully"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesRollupMetrics" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("scorpion brand consequently aside unaccountably phew where quirkily know jubilantly"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesSNMPTrapSerialize" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("untrue frugal anenst inasmuch smug"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesSampling" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("by correctly forecast cumbersome an astride bran"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesSerialize" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("knowledgeably ligate optimal scotch wombat plus ugh print"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesSuppress" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("merit aw through inasmuch scholarship exempt jacket sinful wisely"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesTee" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("through furthermore woot designation"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesUnroll" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("interior indeed loftily contrast um stigmatize accomplished endow geez"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="createPipelinesByPack" method="post" path="/p/{pack}/pipelines" example="PipelineExamplesXMLUnroll" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Create(ctx, "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("hollow awful hmph"),
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
            Groups: map[string]components.PipelineGroups{

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
| `pack`                                                               | *string*                                                             | :heavy_check_mark:                                                   | The <code>id</code> of the Pack to create.                           |
| `pipeline`                                                           | [components.PipelineInput](../../models/components/pipelineinput.md) | :heavy_check_mark:                                                   | Pipeline object                                                      |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.CreatePipelinesByPackResponse](../../models/operations/createpipelinesbypackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Get a list of all Pipelines within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPipelinesByPack" method="get" path="/p/{pack}/pipelines" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.List(ctx, "<value>")
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
| `pack`                                                   | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Pack to list.                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetPipelinesByPackResponse](../../models/operations/getpipelinesbypackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete the specified Pipeline within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="deletePipelinesByPackAndId" method="delete" path="/p/{pack}/pipelines/{id}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Delete(ctx, "<id>", "<value>")
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
| `pack`                                                   | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Pack to delete.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeletePipelinesByPackAndIDResponse](../../models/operations/deletepipelinesbypackandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get the specified Pipeline within the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPipelinesByPackAndId" method="get" path="/p/{pack}/pipelines/{id}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Get(ctx, "<id>", "<value>")
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
| `pack`                                                   | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Pack to get.                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetPipelinesByPackAndIDResponse](../../models/operations/getpipelinesbypackandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the specified Pipeline within the specified Pack.</br></br>Provide a complete representation of the Pipeline that you want to update in the request body. This endpoint does not support partial updates. Cribl removes any omitted fields when updating the Pipeline.</br></br>Confirm that the configuration in your request body is correct before sending the request. If the configuration is incorrect, the updated Pipeline might not function as expected.

### Example Usage: PipelineExamplesAggregateMetrics

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesAggregateMetrics" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("yearn until think needy serialize"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
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
                            FlushEventLimit: criblcontrolplanesdkgo.Pointer[float64](2771.87),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesAggregations" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("yak handle internal unless conjecture ride trash thorough crossly failing"),
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
                            FlushEventLimit: criblcontrolplanesdkgo.Pointer[float64](1229.32),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesAutoTimestamp" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("lest however which doodle pushy capitalise holster chops respectful"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesCEFSerializer" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("hospitable although experienced carpool numeracy like stool"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesChain" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("whenever a carelessly doting"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.PipelineFunctionChainConf{
                            Processor: "prometheus_metrics",
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesClone" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("acclaimed celsius sideboard"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesComment" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("internationalize meanwhile earnest shipper"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaComment{
                            Comment: criblcontrolplanesdkgo.Pointer("This function processes security events and enriches them with DNS lookups"),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesDNSLookup" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("lumpy however cleverly cleaner knowledgeable defrag impact helplessly"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
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
                                "<value 2>",
                                "<value 3>",
                            },
                            LookupFailLogLevel: components.LogLevelForFailedLookupsError.ToPointer(),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesDrop" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("brightly mortise lamp sardonic because sure-footed narrowcast"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaDrop{},
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesDropDimensions" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("young times crocodile great masticate why ornate atop bah"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesDynamicSampling" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("whether lady where"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesEmpty" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
        ID: "empty-pipeline",
        Conf: components.ConfInput{
            AsyncFuncTimeout: criblcontrolplanesdkgo.Pointer[int64](1000),
            Output: criblcontrolplanesdkgo.Pointer("default"),
            Description: criblcontrolplanesdkgo.Pointer(""),
            Streamtags: []string{},
            Functions: []components.PipelineFunctionConfInput{},
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesEval" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("than nor following"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesEventBreaker" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("incidentally scornful past vivaciously aha genuine essence"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesFlatten" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("skyline beautifully of firm championship clamour impanel order"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesFoldKeys" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("amidst utterly loyally mould"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesGeoIP" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("catch nippy gah brightly gadzooks amused"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesGrok" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("afore sardonic glittering whether provided negligible pasta even opposite"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesGuard" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("circa corral crushing vision"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
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
                            BackgroundDetection: criblcontrolplanesdkgo.Pointer(false),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesJSONUnroll" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("revitalise yippee forgather although republican diligently near microchip gosh"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.PipelineFunctionJSONUnrollConf{
                            Path: "allCars",
                            Name: criblcontrolplanesdkgo.Pointer("cars"),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesLookup" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("slimy warmhearted meaningfully ack unabashedly pish likewise timely shore sprinkles"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesMask" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("anenst cop-out loftily illusion sugary passport astride"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesNumerify" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("blah cautious rudely for since"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesOTLPLogs" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("yummy bare absent as jagged owlishly pish upright"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
                        Conf: components.FunctionConfSchemaOtlpLogs{
                            DropNonLogEvents: criblcontrolplanesdkgo.Pointer(false),
                            BatchOTLPLogs: criblcontrolplanesdkgo.Pointer(true),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesOTLPMetrics" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("unlearn worthwhile tall questionable"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesOTLPTraces" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("aha interchange tabulate considering square"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesParser" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("schlep insolence gadzooks"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesPublishMetrics" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("coast if qua wear tuxedo likewise"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesRedis" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("though emphasise conservation selfish aside gently amendment internationalize submitter indeed"),
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
                            EnableClientSideCaching: criblcontrolplanesdkgo.Pointer(true),
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesRegexExtract" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("simplistic tankful consequently optimistically inferior bide as aha eek sharply"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(false),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesRegexFilter" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("rough voluminous woot sans"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesRename" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("closely lack vivaciously below needily draft norm sarong for robust"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesRollupMetrics" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("ignorance phooey offset insert"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(false),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesSNMPTrapSerialize" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("junior ah over sour"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.FunctionConfSchemaSnmpTrapSerialize{
                            Strict: criblcontrolplanesdkgo.Pointer(true),
                            DropFailedEvents: criblcontrolplanesdkgo.Pointer(true),
                            V3User: &components.FunctionConfSchemaSnmpTrapSerializeV3User{
                                Name: criblcontrolplanesdkgo.Pointer("<value>"),
                                AuthProtocol: components.AuthenticationProtocolOptionsV3UserSha384.ToPointer(),
                                AuthKey: "<value>",
                                PrivProtocol: criblcontrolplanesdkgo.Pointer("<value>"),
                            },
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesSampling" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("briefly summarise old mmm redress yowza"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesSerialize" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("language gulp pish bonnet keenly helpfully unit manner"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesSuppress" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("membership infatuated inquisitively save brr abaft puff before coarse nauseate"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesTee" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("horse punctually barring intently bootleg till"),
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
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesUnroll" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("unto confusion ha row premier intensely advertisement woot"),
                        Disabled: criblcontrolplanesdkgo.Pointer(false),
                        Final: criblcontrolplanesdkgo.Pointer(true),
                        Conf: components.PipelineFunctionUnrollConf{
                            SrcExpr: "_raw.split(/\\n/)",
                            DstField: "_raw",
                        },
                        GroupID: criblcontrolplanesdkgo.Pointer("<id>"),
                    },
                ),
            },
            Groups: map[string]components.PipelineGroups{

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

<!-- UsageSnippet language="go" operationID="updatePipelinesByPackAndId" method="patch" path="/p/{pack}/pipelines/{id}" example="PipelineExamplesXMLUnroll" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Packs.Pipelines.Update(ctx, "<id>", "<value>", components.PipelineInput{
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
                        Description: criblcontrolplanesdkgo.Pointer("yet misappropriate beyond conceptualize hawk tomorrow electric viability kettledrum"),
                        Disabled: criblcontrolplanesdkgo.Pointer(true),
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
            Groups: map[string]components.PipelineGroups{

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
| `pack`                                                               | *string*                                                             | :heavy_check_mark:                                                   | The <code>id</code> of the Pack to update.                           |
| `pipeline`                                                           | [components.PipelineInput](../../models/components/pipelineinput.md) | :heavy_check_mark:                                                   | Pipeline object                                                      |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.UpdatePipelinesByPackAndIDResponse](../../models/operations/updatepipelinesbypackandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |