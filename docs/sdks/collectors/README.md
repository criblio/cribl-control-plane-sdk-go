# Collectors

## Overview

Actions related to Collectors

### Available Operations

* [Create](#create) - Create a Collector
* [List](#list) - List all Collectors
* [Delete](#delete) - Delete a Collector
* [Get](#get) - Get a Collector
* [Update](#update) - Update a Collector

## Create

Create a new Collector.

### Example Usage: CollectorExamplesAzureBlob

<!-- UsageSnippet language="go" operationID="createSavedJob" method="post" path="/lib/jobs" example="CollectorExamplesAzureBlob" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Collectors.Create(ctx, components.CreateSavedJobSavedJobExecutor(
        components.SavedJobExecutor{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("yowza than voluntarily phooey meanwhile"),
            Type: components.JobTypeOptionsSavedJobCollectionCollection,
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            RemoveFields: []string{
                "<value 1>",
                "<value 2>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(false),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(true),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](3006.78),
                Run: &components.ScheduleTypeSavedJobCollectionRunSettings{
                    Type: components.ScheduleTypeSavedJobCollectionTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1211.14),
                    LogLevel: components.ScheduleTypeSavedJobCollectionLogLevelDebug.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
                    Mode: "<value>",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](4847.66),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](3337.75),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.TimeWarning{},
                    Expression: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Executor: components.ExecutorTypeSavedJobExecutor{
                Type: "<value>",
                StoreTaskResults: criblcontrolplanesdkgo.Pointer(true),
                Conf: &components.ExecutorSpecificSettingsTypeSavedJobExecutorExecutor{},
            },
        },
    ), criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```
### Example Usage: CollectorExamplesCriblLake

<!-- UsageSnippet language="go" operationID="createSavedJob" method="post" path="/lib/jobs" example="CollectorExamplesCriblLake" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Collectors.Create(ctx, components.CreateSavedJobSavedJobScheduledSearch(
        components.SavedJobScheduledSearch{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("during disconnection where although airman"),
            Type: components.JobTypeOptionsSavedJobCollectionScheduledSearch,
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(true),
            RemoveFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(true),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](3006.78),
                Run: &components.ScheduleTypeSavedJobCollectionRunSettings{
                    Type: components.ScheduleTypeSavedJobCollectionTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1211.14),
                    LogLevel: components.ScheduleTypeSavedJobCollectionLogLevelDebug.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
                    Mode: "<value>",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](4847.66),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](3337.75),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.TimeWarning{},
                    Expression: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Streamtags: []string{
                "<value 1>",
            },
            SavedQueryID: "<id>",
        },
    ), criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```
### Example Usage: CollectorExamplesDatabase

<!-- UsageSnippet language="go" operationID="createSavedJob" method="post" path="/lib/jobs" example="CollectorExamplesDatabase" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Collectors.Create(ctx, components.CreateSavedJobSavedJobExecutor(
        components.SavedJobExecutor{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("yowza than voluntarily phooey meanwhile"),
            Type: components.JobTypeOptionsSavedJobCollectionCollection,
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            RemoveFields: []string{
                "<value 1>",
                "<value 2>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(false),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(true),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](3006.78),
                Run: &components.ScheduleTypeSavedJobCollectionRunSettings{
                    Type: components.ScheduleTypeSavedJobCollectionTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1211.14),
                    LogLevel: components.ScheduleTypeSavedJobCollectionLogLevelDebug.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
                    Mode: "<value>",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](4847.66),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](3337.75),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.TimeWarning{},
                    Expression: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Executor: components.ExecutorTypeSavedJobExecutor{
                Type: "<value>",
                StoreTaskResults: criblcontrolplanesdkgo.Pointer(true),
                Conf: &components.ExecutorSpecificSettingsTypeSavedJobExecutorExecutor{},
            },
        },
    ), criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```
### Example Usage: CollectorExamplesFilesystem

<!-- UsageSnippet language="go" operationID="createSavedJob" method="post" path="/lib/jobs" example="CollectorExamplesFilesystem" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Collectors.Create(ctx, components.CreateSavedJobSavedJobScheduledSearch(
        components.SavedJobScheduledSearch{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("during disconnection where although airman"),
            Type: components.JobTypeOptionsSavedJobCollectionScheduledSearch,
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(true),
            RemoveFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(true),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](3006.78),
                Run: &components.ScheduleTypeSavedJobCollectionRunSettings{
                    Type: components.ScheduleTypeSavedJobCollectionTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1211.14),
                    LogLevel: components.ScheduleTypeSavedJobCollectionLogLevelDebug.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
                    Mode: "<value>",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](4847.66),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](3337.75),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.TimeWarning{},
                    Expression: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Streamtags: []string{
                "<value 1>",
            },
            SavedQueryID: "<id>",
        },
    ), criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```
### Example Usage: CollectorExamplesGoogleCloudStorage

<!-- UsageSnippet language="go" operationID="createSavedJob" method="post" path="/lib/jobs" example="CollectorExamplesGoogleCloudStorage" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Collectors.Create(ctx, components.CreateSavedJobSavedJobScheduledSearch(
        components.SavedJobScheduledSearch{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("during disconnection where although airman"),
            Type: components.JobTypeOptionsSavedJobCollectionScheduledSearch,
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(true),
            RemoveFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(true),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](3006.78),
                Run: &components.ScheduleTypeSavedJobCollectionRunSettings{
                    Type: components.ScheduleTypeSavedJobCollectionTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1211.14),
                    LogLevel: components.ScheduleTypeSavedJobCollectionLogLevelDebug.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
                    Mode: "<value>",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](4847.66),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](3337.75),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.TimeWarning{},
                    Expression: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Streamtags: []string{
                "<value 1>",
            },
            SavedQueryID: "<id>",
        },
    ), criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```
### Example Usage: CollectorExamplesRest

<!-- UsageSnippet language="go" operationID="createSavedJob" method="post" path="/lib/jobs" example="CollectorExamplesRest" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Collectors.Create(ctx, components.CreateSavedJobSavedJobCollection(
        components.SavedJobCollection{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("pomelo outside offensively ew"),
            Type: components.JobTypeOptionsSavedJobCollectionExecutor,
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            RemoveFields: []string{
                "<value 1>",
                "<value 2>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(true),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](3006.78),
                Run: &components.ScheduleTypeSavedJobCollectionRunSettings{
                    Type: components.ScheduleTypeSavedJobCollectionTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1211.14),
                    LogLevel: components.ScheduleTypeSavedJobCollectionLogLevelDebug.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
                    Mode: "<value>",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](4847.66),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](3337.75),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.TimeWarning{},
                    Expression: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            WorkerAffinity: criblcontrolplanesdkgo.Pointer(true),
            Collector: components.CreateCollectorDatabase(
                components.CollectorDatabase{
                    Type: components.CollectorDatabaseTypeDatabase,
                    Conf: components.DatabaseCollectorConf{
                        ConnectionID: "<id>",
                        Query: "<value>",
                        QueryValidationEnabled: criblcontrolplanesdkgo.Pointer(true),
                        DefaultBreakers: components.HiddenDefaultBreakersOptionsDatabaseCollectorConfCribl.ToPointer(),
                        Scheduling: &components.DatabaseCollectorConfScheduling{
                            StateTracking: &components.DatabaseCollectorConfStateTracking{
                                Enabled: criblcontrolplanesdkgo.Pointer(false),
                            },
                        },
                    },
                    Destructive: criblcontrolplanesdkgo.Pointer(true),
                    Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            ),
            Input: &components.TypeCollectionWithBreakerRulesetsConstraint{
                Type: components.TypeCollectionWithBreakerRulesetsConstraintTypeCollection.ToPointer(),
                BreakerRulesets: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
                StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](9538.43),
                SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
                Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                    Disabled: true,
                    Command: criblcontrolplanesdkgo.Pointer("<value>"),
                    Args: []string{
                        "<value 1>",
                        "<value 2>",
                    },
                },
                ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
                Metadata: []components.ItemsTypeNotificationMetadata{
                    components.ItemsTypeNotificationMetadata{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                Output: criblcontrolplanesdkgo.Pointer("<value>"),
            },
        },
    ), criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```
### Example Usage: CollectorExamplesS3

<!-- UsageSnippet language="go" operationID="createSavedJob" method="post" path="/lib/jobs" example="CollectorExamplesS3" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Collectors.Create(ctx, components.CreateSavedJobSavedJobCollection(
        components.SavedJobCollection{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("pomelo outside offensively ew"),
            Type: components.JobTypeOptionsSavedJobCollectionExecutor,
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            RemoveFields: []string{
                "<value 1>",
                "<value 2>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(true),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](3006.78),
                Run: &components.ScheduleTypeSavedJobCollectionRunSettings{
                    Type: components.ScheduleTypeSavedJobCollectionTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1211.14),
                    LogLevel: components.ScheduleTypeSavedJobCollectionLogLevelDebug.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
                    Mode: "<value>",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](4847.66),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](3337.75),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.TimeWarning{},
                    Expression: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            WorkerAffinity: criblcontrolplanesdkgo.Pointer(true),
            Collector: components.CreateCollectorDatabase(
                components.CollectorDatabase{
                    Type: components.CollectorDatabaseTypeDatabase,
                    Conf: components.DatabaseCollectorConf{
                        ConnectionID: "<id>",
                        Query: "<value>",
                        QueryValidationEnabled: criblcontrolplanesdkgo.Pointer(true),
                        DefaultBreakers: components.HiddenDefaultBreakersOptionsDatabaseCollectorConfCribl.ToPointer(),
                        Scheduling: &components.DatabaseCollectorConfScheduling{
                            StateTracking: &components.DatabaseCollectorConfStateTracking{
                                Enabled: criblcontrolplanesdkgo.Pointer(false),
                            },
                        },
                    },
                    Destructive: criblcontrolplanesdkgo.Pointer(true),
                    Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            ),
            Input: &components.TypeCollectionWithBreakerRulesetsConstraint{
                Type: components.TypeCollectionWithBreakerRulesetsConstraintTypeCollection.ToPointer(),
                BreakerRulesets: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
                StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](9538.43),
                SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
                Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                    Disabled: true,
                    Command: criblcontrolplanesdkgo.Pointer("<value>"),
                    Args: []string{
                        "<value 1>",
                        "<value 2>",
                    },
                },
                ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
                Metadata: []components.ItemsTypeNotificationMetadata{
                    components.ItemsTypeNotificationMetadata{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                Output: criblcontrolplanesdkgo.Pointer("<value>"),
            },
        },
    ), criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```
### Example Usage: CollectorExamplesScript

<!-- UsageSnippet language="go" operationID="createSavedJob" method="post" path="/lib/jobs" example="CollectorExamplesScript" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Collectors.Create(ctx, components.CreateSavedJobSavedJobCollection(
        components.SavedJobCollection{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("pomelo outside offensively ew"),
            Type: components.JobTypeOptionsSavedJobCollectionExecutor,
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            RemoveFields: []string{
                "<value 1>",
                "<value 2>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(true),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(true),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](3006.78),
                Run: &components.ScheduleTypeSavedJobCollectionRunSettings{
                    Type: components.ScheduleTypeSavedJobCollectionTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1211.14),
                    LogLevel: components.ScheduleTypeSavedJobCollectionLogLevelDebug.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
                    Mode: "<value>",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](4847.66),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](3337.75),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.TimeWarning{},
                    Expression: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            WorkerAffinity: criblcontrolplanesdkgo.Pointer(true),
            Collector: components.CreateCollectorDatabase(
                components.CollectorDatabase{
                    Type: components.CollectorDatabaseTypeDatabase,
                    Conf: components.DatabaseCollectorConf{
                        ConnectionID: "<id>",
                        Query: "<value>",
                        QueryValidationEnabled: criblcontrolplanesdkgo.Pointer(true),
                        DefaultBreakers: components.HiddenDefaultBreakersOptionsDatabaseCollectorConfCribl.ToPointer(),
                        Scheduling: &components.DatabaseCollectorConfScheduling{
                            StateTracking: &components.DatabaseCollectorConfStateTracking{
                                Enabled: criblcontrolplanesdkgo.Pointer(false),
                            },
                        },
                    },
                    Destructive: criblcontrolplanesdkgo.Pointer(true),
                    Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            ),
            Input: &components.TypeCollectionWithBreakerRulesetsConstraint{
                Type: components.TypeCollectionWithBreakerRulesetsConstraintTypeCollection.ToPointer(),
                BreakerRulesets: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
                StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](9538.43),
                SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
                Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                    Disabled: true,
                    Command: criblcontrolplanesdkgo.Pointer("<value>"),
                    Args: []string{
                        "<value 1>",
                        "<value 2>",
                    },
                },
                ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
                Metadata: []components.ItemsTypeNotificationMetadata{
                    components.ItemsTypeNotificationMetadata{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                Output: criblcontrolplanesdkgo.Pointer("<value>"),
            },
        },
    ), criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```
### Example Usage: CollectorExamplesSplunk

<!-- UsageSnippet language="go" operationID="createSavedJob" method="post" path="/lib/jobs" example="CollectorExamplesSplunk" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Collectors.Create(ctx, components.CreateSavedJobSavedJobExecutor(
        components.SavedJobExecutor{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("yowza than voluntarily phooey meanwhile"),
            Type: components.JobTypeOptionsSavedJobCollectionCollection,
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            RemoveFields: []string{
                "<value 1>",
                "<value 2>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(false),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(true),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](3006.78),
                Run: &components.ScheduleTypeSavedJobCollectionRunSettings{
                    Type: components.ScheduleTypeSavedJobCollectionTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1211.14),
                    LogLevel: components.ScheduleTypeSavedJobCollectionLogLevelDebug.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
                    Mode: "<value>",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](4847.66),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](3337.75),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.TimeWarning{},
                    Expression: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Executor: components.ExecutorTypeSavedJobExecutor{
                Type: "<value>",
                StoreTaskResults: criblcontrolplanesdkgo.Pointer(true),
                Conf: &components.ExecutorSpecificSettingsTypeSavedJobExecutorExecutor{},
            },
        },
    ), criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |
| `savedJob`                                                  | [components.SavedJob](../../models/components/savedjob.md)  | :heavy_check_mark:                                          | SavedJob object                                             |
| `criblPack`                                                 | **string*                                                   | :heavy_minus_sign:                                          | The <code>id</code> of the Pack to create the Collector in. |
| `opts`                                                      | [][operations.Option](../../models/operations/option.md)    | :heavy_minus_sign:                                          | The options for this request.                               |

### Response

**[*operations.CreateSavedJobResponse](../../models/operations/createsavedjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Get a list of all Collectors.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSavedJob" method="get" path="/lib/jobs" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Collectors.List(ctx, criblcontrolplanesdkgo.Pointer("<value>"), criblcontrolplanesdkgo.Pointer("<value>"), criblcontrolplanesdkgo.Pointer("<id>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `collectorType`                                          | **string*                                                | :heavy_minus_sign:                                       | Filter by collector type                                 |
| `criblPack`                                              | **string*                                                | :heavy_minus_sign:                                       | Pack ID                                                  |
| `groupID`                                                | **string*                                                | :heavy_minus_sign:                                       | Worker group ID                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSavedJobResponse](../../models/operations/getsavedjobresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete the specified Collector.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteSavedJobById" method="delete" path="/lib/jobs/{id}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Collectors.Delete(ctx, "<id>", criblcontrolplanesdkgo.Pointer("<value>"), criblcontrolplanesdkgo.Pointer("<id>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `id`                                                                           | *string*                                                                       | :heavy_check_mark:                                                             | The <code>id</code> of the Collector to delete.                                |
| `criblPack`                                                                    | **string*                                                                      | :heavy_minus_sign:                                                             | The <code>id</code> of the Pack that includes the Collector to delete.         |
| `groupID`                                                                      | **string*                                                                      | :heavy_minus_sign:                                                             | The <code>id</code> of the Worker Group that includes the Collector to delete. |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.DeleteSavedJobByIDResponse](../../models/operations/deletesavedjobbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get the specified Collector.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSavedJobById" method="get" path="/lib/jobs/{id}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Collectors.Get(ctx, "<id>", criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                           | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ctx`                                                               | [context.Context](https://pkg.go.dev/context#Context)               | :heavy_check_mark:                                                  | The context to use for the request.                                 |
| `id`                                                                | *string*                                                            | :heavy_check_mark:                                                  | The <code>id</code> of the Collector to get.                        |
| `criblPack`                                                         | **string*                                                           | :heavy_minus_sign:                                                  | The <code>id</code> of the Pack that includes the Collector to get. |
| `opts`                                                              | [][operations.Option](../../models/operations/option.md)            | :heavy_minus_sign:                                                  | The options for this request.                                       |

### Response

**[*operations.GetSavedJobByIDResponse](../../models/operations/getsavedjobbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the specified Collector.<br><br>Provide a complete representation of the Collector that you want to update in the request body. This endpoint does not support partial updates. Cribl removes any omitted fields when updating the Collector.<br><br>Confirm that the configuration in your request body is correct before sending the request. If the configuration is incorrect, the updated Collector might not function as expected.

### Example Usage: CollectorExamplesAzureBlob

<!-- UsageSnippet language="go" operationID="updateSavedJobById" method="patch" path="/lib/jobs/{id}" example="CollectorExamplesAzureBlob" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Collectors.Update(ctx, "<id>", components.CreateSavedJobSavedJobScheduledSearch(
        components.SavedJobScheduledSearch{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("blaring spectate dark notwithstanding sparse obnoxiously editor"),
            Type: components.JobTypeOptionsSavedJobCollectionExecutor,
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(true),
            RemoveFields: []string{
                "<value 1>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(false),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(false),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1498.35),
                Run: &components.ScheduleTypeSavedJobCollectionRunSettings{
                    Type: components.ScheduleTypeSavedJobCollectionTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(false),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](9677.47),
                    LogLevel: components.ScheduleTypeSavedJobCollectionLogLevelError.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
                    Mode: "<value>",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](8882.78),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](6778.74),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.TimeWarning{},
                    Expression: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            SavedQueryID: "<id>",
        },
    ), criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```
### Example Usage: CollectorExamplesCriblLake

<!-- UsageSnippet language="go" operationID="updateSavedJobById" method="patch" path="/lib/jobs/{id}" example="CollectorExamplesCriblLake" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Collectors.Update(ctx, "<id>", components.CreateSavedJobSavedJobScheduledSearch(
        components.SavedJobScheduledSearch{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("blaring spectate dark notwithstanding sparse obnoxiously editor"),
            Type: components.JobTypeOptionsSavedJobCollectionExecutor,
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(true),
            RemoveFields: []string{
                "<value 1>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(false),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(false),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1498.35),
                Run: &components.ScheduleTypeSavedJobCollectionRunSettings{
                    Type: components.ScheduleTypeSavedJobCollectionTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(false),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](9677.47),
                    LogLevel: components.ScheduleTypeSavedJobCollectionLogLevelError.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
                    Mode: "<value>",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](8882.78),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](6778.74),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.TimeWarning{},
                    Expression: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            SavedQueryID: "<id>",
        },
    ), criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```
### Example Usage: CollectorExamplesDatabase

<!-- UsageSnippet language="go" operationID="updateSavedJobById" method="patch" path="/lib/jobs/{id}" example="CollectorExamplesDatabase" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Collectors.Update(ctx, "<id>", components.CreateSavedJobSavedJobScheduledSearch(
        components.SavedJobScheduledSearch{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("blaring spectate dark notwithstanding sparse obnoxiously editor"),
            Type: components.JobTypeOptionsSavedJobCollectionExecutor,
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(true),
            RemoveFields: []string{
                "<value 1>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(false),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(false),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1498.35),
                Run: &components.ScheduleTypeSavedJobCollectionRunSettings{
                    Type: components.ScheduleTypeSavedJobCollectionTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(false),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](9677.47),
                    LogLevel: components.ScheduleTypeSavedJobCollectionLogLevelError.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
                    Mode: "<value>",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](8882.78),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](6778.74),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.TimeWarning{},
                    Expression: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            SavedQueryID: "<id>",
        },
    ), criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```
### Example Usage: CollectorExamplesFilesystem

<!-- UsageSnippet language="go" operationID="updateSavedJobById" method="patch" path="/lib/jobs/{id}" example="CollectorExamplesFilesystem" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Collectors.Update(ctx, "<id>", components.CreateSavedJobSavedJobCollection(
        components.SavedJobCollection{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("unabashedly notwithstanding ugh digestive"),
            Type: components.JobTypeOptionsSavedJobCollectionScheduledSearch,
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            RemoveFields: []string{
                "<value 1>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(false),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(false),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1498.35),
                Run: &components.ScheduleTypeSavedJobCollectionRunSettings{
                    Type: components.ScheduleTypeSavedJobCollectionTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(false),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](9677.47),
                    LogLevel: components.ScheduleTypeSavedJobCollectionLogLevelError.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
                    Mode: "<value>",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](8882.78),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](6778.74),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.TimeWarning{},
                    Expression: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            WorkerAffinity: criblcontrolplanesdkgo.Pointer(false),
            Collector: components.CreateCollectorS3(
                components.CollectorS3{
                    Type: components.CollectorS3TypeS3,
                    Conf: components.CreateS3CollectorConfAuto(
                        components.S3AwsAuthenticationMethodAuto{
                            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
                            OutputName: criblcontrolplanesdkgo.Pointer("<value>"),
                            Bucket: "<value>",
                            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](2532.22),
                            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](6271.26),
                            Region: criblcontrolplanesdkgo.Pointer("<value>"),
                            Path: criblcontrolplanesdkgo.Pointer("/selinux"),
                            PartitioningScheme: components.S3AwsAuthenticationMethodAutoPartitioningSchemeNone.ToPointer(),
                            Extractors: []components.S3AwsAuthenticationMethodAutoExtractor{
                                components.S3AwsAuthenticationMethodAutoExtractor{
                                    Key: "<key>",
                                    Expression: "<value>",
                                },
                            },
                            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
                            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
                            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
                            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
                            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
                            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](2075.63),
                            MaxBatchSize: criblcontrolplanesdkgo.Pointer[float64](968.91),
                            Recurse: "<value>",
                            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
                            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
                            DisableTimeFilter: criblcontrolplanesdkgo.Pointer(true),
                        },
                    ),
                    Destructive: criblcontrolplanesdkgo.Pointer(false),
                    Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            ),
            Input: &components.TypeCollectionWithBreakerRulesetsConstraint{
                Type: components.TypeCollectionWithBreakerRulesetsConstraintTypeCollection.ToPointer(),
                BreakerRulesets: []string{
                    "<value 1>",
                    "<value 2>",
                },
                StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](2918.22),
                SendToRoutes: criblcontrolplanesdkgo.Pointer(false),
                Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                    Disabled: false,
                    Command: criblcontrolplanesdkgo.Pointer("<value>"),
                    Args: []string{
                        "<value 1>",
                    },
                },
                ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
                Metadata: []components.ItemsTypeNotificationMetadata{
                    components.ItemsTypeNotificationMetadata{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                Output: criblcontrolplanesdkgo.Pointer("<value>"),
            },
        },
    ), criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```
### Example Usage: CollectorExamplesGoogleCloudStorage

<!-- UsageSnippet language="go" operationID="updateSavedJobById" method="patch" path="/lib/jobs/{id}" example="CollectorExamplesGoogleCloudStorage" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Collectors.Update(ctx, "<id>", components.CreateSavedJobSavedJobCollection(
        components.SavedJobCollection{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("unabashedly notwithstanding ugh digestive"),
            Type: components.JobTypeOptionsSavedJobCollectionScheduledSearch,
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            RemoveFields: []string{
                "<value 1>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(false),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(false),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1498.35),
                Run: &components.ScheduleTypeSavedJobCollectionRunSettings{
                    Type: components.ScheduleTypeSavedJobCollectionTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(false),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](9677.47),
                    LogLevel: components.ScheduleTypeSavedJobCollectionLogLevelError.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
                    Mode: "<value>",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](8882.78),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](6778.74),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.TimeWarning{},
                    Expression: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            WorkerAffinity: criblcontrolplanesdkgo.Pointer(false),
            Collector: components.CreateCollectorS3(
                components.CollectorS3{
                    Type: components.CollectorS3TypeS3,
                    Conf: components.CreateS3CollectorConfAuto(
                        components.S3AwsAuthenticationMethodAuto{
                            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
                            OutputName: criblcontrolplanesdkgo.Pointer("<value>"),
                            Bucket: "<value>",
                            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](2532.22),
                            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](6271.26),
                            Region: criblcontrolplanesdkgo.Pointer("<value>"),
                            Path: criblcontrolplanesdkgo.Pointer("/selinux"),
                            PartitioningScheme: components.S3AwsAuthenticationMethodAutoPartitioningSchemeNone.ToPointer(),
                            Extractors: []components.S3AwsAuthenticationMethodAutoExtractor{
                                components.S3AwsAuthenticationMethodAutoExtractor{
                                    Key: "<key>",
                                    Expression: "<value>",
                                },
                            },
                            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
                            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
                            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
                            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
                            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
                            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](2075.63),
                            MaxBatchSize: criblcontrolplanesdkgo.Pointer[float64](968.91),
                            Recurse: "<value>",
                            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
                            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
                            DisableTimeFilter: criblcontrolplanesdkgo.Pointer(true),
                        },
                    ),
                    Destructive: criblcontrolplanesdkgo.Pointer(false),
                    Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            ),
            Input: &components.TypeCollectionWithBreakerRulesetsConstraint{
                Type: components.TypeCollectionWithBreakerRulesetsConstraintTypeCollection.ToPointer(),
                BreakerRulesets: []string{
                    "<value 1>",
                    "<value 2>",
                },
                StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](2918.22),
                SendToRoutes: criblcontrolplanesdkgo.Pointer(false),
                Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                    Disabled: false,
                    Command: criblcontrolplanesdkgo.Pointer("<value>"),
                    Args: []string{
                        "<value 1>",
                    },
                },
                ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
                Metadata: []components.ItemsTypeNotificationMetadata{
                    components.ItemsTypeNotificationMetadata{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                Output: criblcontrolplanesdkgo.Pointer("<value>"),
            },
        },
    ), criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```
### Example Usage: CollectorExamplesRest

<!-- UsageSnippet language="go" operationID="updateSavedJobById" method="patch" path="/lib/jobs/{id}" example="CollectorExamplesRest" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Collectors.Update(ctx, "<id>", components.CreateSavedJobSavedJobCollection(
        components.SavedJobCollection{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("unabashedly notwithstanding ugh digestive"),
            Type: components.JobTypeOptionsSavedJobCollectionScheduledSearch,
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            RemoveFields: []string{
                "<value 1>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(false),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(false),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1498.35),
                Run: &components.ScheduleTypeSavedJobCollectionRunSettings{
                    Type: components.ScheduleTypeSavedJobCollectionTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(false),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](9677.47),
                    LogLevel: components.ScheduleTypeSavedJobCollectionLogLevelError.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
                    Mode: "<value>",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](8882.78),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](6778.74),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.TimeWarning{},
                    Expression: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            WorkerAffinity: criblcontrolplanesdkgo.Pointer(false),
            Collector: components.CreateCollectorS3(
                components.CollectorS3{
                    Type: components.CollectorS3TypeS3,
                    Conf: components.CreateS3CollectorConfAuto(
                        components.S3AwsAuthenticationMethodAuto{
                            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
                            OutputName: criblcontrolplanesdkgo.Pointer("<value>"),
                            Bucket: "<value>",
                            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](2532.22),
                            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](6271.26),
                            Region: criblcontrolplanesdkgo.Pointer("<value>"),
                            Path: criblcontrolplanesdkgo.Pointer("/selinux"),
                            PartitioningScheme: components.S3AwsAuthenticationMethodAutoPartitioningSchemeNone.ToPointer(),
                            Extractors: []components.S3AwsAuthenticationMethodAutoExtractor{
                                components.S3AwsAuthenticationMethodAutoExtractor{
                                    Key: "<key>",
                                    Expression: "<value>",
                                },
                            },
                            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
                            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
                            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
                            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
                            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
                            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](2075.63),
                            MaxBatchSize: criblcontrolplanesdkgo.Pointer[float64](968.91),
                            Recurse: "<value>",
                            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
                            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
                            DisableTimeFilter: criblcontrolplanesdkgo.Pointer(true),
                        },
                    ),
                    Destructive: criblcontrolplanesdkgo.Pointer(true),
                    Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            ),
            Input: &components.TypeCollectionWithBreakerRulesetsConstraint{
                Type: components.TypeCollectionWithBreakerRulesetsConstraintTypeCollection.ToPointer(),
                BreakerRulesets: []string{
                    "<value 1>",
                    "<value 2>",
                },
                StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](2918.22),
                SendToRoutes: criblcontrolplanesdkgo.Pointer(false),
                Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                    Disabled: false,
                    Command: criblcontrolplanesdkgo.Pointer("<value>"),
                    Args: []string{
                        "<value 1>",
                    },
                },
                ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
                Metadata: []components.ItemsTypeNotificationMetadata{
                    components.ItemsTypeNotificationMetadata{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                Output: criblcontrolplanesdkgo.Pointer("<value>"),
            },
        },
    ), criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```
### Example Usage: CollectorExamplesS3

<!-- UsageSnippet language="go" operationID="updateSavedJobById" method="patch" path="/lib/jobs/{id}" example="CollectorExamplesS3" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Collectors.Update(ctx, "<id>", components.CreateSavedJobSavedJobCollection(
        components.SavedJobCollection{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("unabashedly notwithstanding ugh digestive"),
            Type: components.JobTypeOptionsSavedJobCollectionScheduledSearch,
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            RemoveFields: []string{
                "<value 1>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(false),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(false),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1498.35),
                Run: &components.ScheduleTypeSavedJobCollectionRunSettings{
                    Type: components.ScheduleTypeSavedJobCollectionTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(false),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](9677.47),
                    LogLevel: components.ScheduleTypeSavedJobCollectionLogLevelError.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
                    Mode: "<value>",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](8882.78),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](6778.74),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.TimeWarning{},
                    Expression: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            WorkerAffinity: criblcontrolplanesdkgo.Pointer(false),
            Collector: components.CreateCollectorS3(
                components.CollectorS3{
                    Type: components.CollectorS3TypeS3,
                    Conf: components.CreateS3CollectorConfAuto(
                        components.S3AwsAuthenticationMethodAuto{
                            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
                            OutputName: criblcontrolplanesdkgo.Pointer("<value>"),
                            Bucket: "<value>",
                            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](2532.22),
                            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](6271.26),
                            Region: criblcontrolplanesdkgo.Pointer("<value>"),
                            Path: criblcontrolplanesdkgo.Pointer("/selinux"),
                            PartitioningScheme: components.S3AwsAuthenticationMethodAutoPartitioningSchemeNone.ToPointer(),
                            Extractors: []components.S3AwsAuthenticationMethodAutoExtractor{
                                components.S3AwsAuthenticationMethodAutoExtractor{
                                    Key: "<key>",
                                    Expression: "<value>",
                                },
                            },
                            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
                            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
                            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
                            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
                            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
                            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](2075.63),
                            MaxBatchSize: criblcontrolplanesdkgo.Pointer[float64](968.91),
                            Recurse: "<value>",
                            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
                            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
                            DisableTimeFilter: criblcontrolplanesdkgo.Pointer(true),
                        },
                    ),
                    Destructive: criblcontrolplanesdkgo.Pointer(true),
                    Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            ),
            Input: &components.TypeCollectionWithBreakerRulesetsConstraint{
                Type: components.TypeCollectionWithBreakerRulesetsConstraintTypeCollection.ToPointer(),
                BreakerRulesets: []string{
                    "<value 1>",
                    "<value 2>",
                },
                StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](2918.22),
                SendToRoutes: criblcontrolplanesdkgo.Pointer(false),
                Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                    Disabled: false,
                    Command: criblcontrolplanesdkgo.Pointer("<value>"),
                    Args: []string{
                        "<value 1>",
                    },
                },
                ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
                Metadata: []components.ItemsTypeNotificationMetadata{
                    components.ItemsTypeNotificationMetadata{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                Output: criblcontrolplanesdkgo.Pointer("<value>"),
            },
        },
    ), criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```
### Example Usage: CollectorExamplesScript

<!-- UsageSnippet language="go" operationID="updateSavedJobById" method="patch" path="/lib/jobs/{id}" example="CollectorExamplesScript" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Collectors.Update(ctx, "<id>", components.CreateSavedJobSavedJobCollection(
        components.SavedJobCollection{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("unabashedly notwithstanding ugh digestive"),
            Type: components.JobTypeOptionsSavedJobCollectionScheduledSearch,
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            RemoveFields: []string{
                "<value 1>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(false),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(false),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1498.35),
                Run: &components.ScheduleTypeSavedJobCollectionRunSettings{
                    Type: components.ScheduleTypeSavedJobCollectionTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(false),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](9677.47),
                    LogLevel: components.ScheduleTypeSavedJobCollectionLogLevelError.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
                    Mode: "<value>",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](8882.78),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](6778.74),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.TimeWarning{},
                    Expression: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            WorkerAffinity: criblcontrolplanesdkgo.Pointer(false),
            Collector: components.CreateCollectorS3(
                components.CollectorS3{
                    Type: components.CollectorS3TypeS3,
                    Conf: components.CreateS3CollectorConfAuto(
                        components.S3AwsAuthenticationMethodAuto{
                            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
                            OutputName: criblcontrolplanesdkgo.Pointer("<value>"),
                            Bucket: "<value>",
                            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](2532.22),
                            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](6271.26),
                            Region: criblcontrolplanesdkgo.Pointer("<value>"),
                            Path: criblcontrolplanesdkgo.Pointer("/selinux"),
                            PartitioningScheme: components.S3AwsAuthenticationMethodAutoPartitioningSchemeNone.ToPointer(),
                            Extractors: []components.S3AwsAuthenticationMethodAutoExtractor{
                                components.S3AwsAuthenticationMethodAutoExtractor{
                                    Key: "<key>",
                                    Expression: "<value>",
                                },
                            },
                            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
                            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
                            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
                            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
                            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
                            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](2075.63),
                            MaxBatchSize: criblcontrolplanesdkgo.Pointer[float64](968.91),
                            Recurse: "<value>",
                            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
                            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
                            DisableTimeFilter: criblcontrolplanesdkgo.Pointer(true),
                        },
                    ),
                    Destructive: criblcontrolplanesdkgo.Pointer(false),
                    Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            ),
            Input: &components.TypeCollectionWithBreakerRulesetsConstraint{
                Type: components.TypeCollectionWithBreakerRulesetsConstraintTypeCollection.ToPointer(),
                BreakerRulesets: []string{
                    "<value 1>",
                    "<value 2>",
                },
                StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](2918.22),
                SendToRoutes: criblcontrolplanesdkgo.Pointer(false),
                Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                    Disabled: false,
                    Command: criblcontrolplanesdkgo.Pointer("<value>"),
                    Args: []string{
                        "<value 1>",
                    },
                },
                ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
                Metadata: []components.ItemsTypeNotificationMetadata{
                    components.ItemsTypeNotificationMetadata{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                Output: criblcontrolplanesdkgo.Pointer("<value>"),
            },
        },
    ), criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```
### Example Usage: CollectorExamplesSplunk

<!-- UsageSnippet language="go" operationID="updateSavedJobById" method="patch" path="/lib/jobs/{id}" example="CollectorExamplesSplunk" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.Collectors.Update(ctx, "<id>", components.CreateSavedJobSavedJobCollection(
        components.SavedJobCollection{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("unabashedly notwithstanding ugh digestive"),
            Type: components.JobTypeOptionsSavedJobCollectionScheduledSearch,
            TTL: criblcontrolplanesdkgo.Pointer("<value>"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            RemoveFields: []string{
                "<value 1>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(false),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(false),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("<value>"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1498.35),
                Run: &components.ScheduleTypeSavedJobCollectionRunSettings{
                    Type: components.ScheduleTypeSavedJobCollectionTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(false),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](9677.47),
                    LogLevel: components.ScheduleTypeSavedJobCollectionLogLevelError.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("<value>"),
                    Mode: "<value>",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("<value>"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](8882.78),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](6778.74),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.TimeWarning{},
                    Expression: criblcontrolplanesdkgo.Pointer("<value>"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            WorkerAffinity: criblcontrolplanesdkgo.Pointer(false),
            Collector: components.CreateCollectorS3(
                components.CollectorS3{
                    Type: components.CollectorS3TypeS3,
                    Conf: components.CreateS3CollectorConfAuto(
                        components.S3AwsAuthenticationMethodAuto{
                            AwsAuthenticationMethod: components.AuthenticationMethodOptionsS3CollectorConfAuto.ToPointer(),
                            OutputName: criblcontrolplanesdkgo.Pointer("<value>"),
                            Bucket: "<value>",
                            ParquetChunkSizeMB: criblcontrolplanesdkgo.Pointer[float64](2532.22),
                            ParquetChunkDownloadTimeout: criblcontrolplanesdkgo.Pointer[float64](6271.26),
                            Region: criblcontrolplanesdkgo.Pointer("<value>"),
                            Path: criblcontrolplanesdkgo.Pointer("/selinux"),
                            PartitioningScheme: components.S3AwsAuthenticationMethodAutoPartitioningSchemeNone.ToPointer(),
                            Extractors: []components.S3AwsAuthenticationMethodAutoExtractor{
                                components.S3AwsAuthenticationMethodAutoExtractor{
                                    Key: "<key>",
                                    Expression: "<value>",
                                },
                            },
                            Endpoint: criblcontrolplanesdkgo.Pointer("<value>"),
                            SignatureVersion: components.SignatureVersionOptionsS3CollectorConfV4.ToPointer(),
                            EnableAssumeRole: criblcontrolplanesdkgo.Pointer(true),
                            AssumeRoleArn: criblcontrolplanesdkgo.Pointer("<value>"),
                            AssumeRoleExternalID: criblcontrolplanesdkgo.Pointer("<id>"),
                            DurationSeconds: criblcontrolplanesdkgo.Pointer[float64](2075.63),
                            MaxBatchSize: criblcontrolplanesdkgo.Pointer[float64](968.91),
                            Recurse: "<value>",
                            ReuseConnections: criblcontrolplanesdkgo.Pointer(true),
                            RejectUnauthorized: criblcontrolplanesdkgo.Pointer(false),
                            VerifyPermissions: criblcontrolplanesdkgo.Pointer(true),
                            DisableTimeFilter: criblcontrolplanesdkgo.Pointer(true),
                        },
                    ),
                    Destructive: criblcontrolplanesdkgo.Pointer(false),
                    Encoding: criblcontrolplanesdkgo.Pointer("<value>"),
                },
            ),
            Input: &components.TypeCollectionWithBreakerRulesetsConstraint{
                Type: components.TypeCollectionWithBreakerRulesetsConstraintTypeCollection.ToPointer(),
                BreakerRulesets: []string{
                    "<value 1>",
                    "<value 2>",
                },
                StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](2918.22),
                SendToRoutes: criblcontrolplanesdkgo.Pointer(false),
                Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                    Disabled: false,
                    Command: criblcontrolplanesdkgo.Pointer("<value>"),
                    Args: []string{
                        "<value 1>",
                    },
                },
                ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("<value>"),
                Metadata: []components.ItemsTypeNotificationMetadata{
                    components.ItemsTypeNotificationMetadata{
                        Name: "<value>",
                        Value: "<value>",
                    },
                },
                Pipeline: criblcontrolplanesdkgo.Pointer("<value>"),
                Output: criblcontrolplanesdkgo.Pointer("<value>"),
            },
        },
    ), criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `id`                                                                   | *string*                                                               | :heavy_check_mark:                                                     | The <code>id</code> of the Collector to update.                        |
| `savedJob`                                                             | [components.SavedJob](../../models/components/savedjob.md)             | :heavy_check_mark:                                                     | SavedJob object                                                        |
| `criblPack`                                                            | **string*                                                              | :heavy_minus_sign:                                                     | The <code>id</code> of the Pack that includes the Collector to update. |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.UpdateSavedJobByIDResponse](../../models/operations/updatesavedjobbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |