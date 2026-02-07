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
            Description: criblcontrolplanesdkgo.Pointer("which um whenever pro chase video uh-huh atop baritone yum"),
            Type: components.JobTypeOptionsRunnableJobCollectionExecutor,
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
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
                CronSchedule: criblcontrolplanesdkgo.Pointer("*/5 * * * *"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1),
                Run: &components.RunSettingsTypeSavedJobCollectionSchedule{
                    Type: components.RunSettingsTypeSavedJobCollectionScheduleTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1),
                    LogLevel: components.LogLevelOptionsRunnableJobCollectionScheduleRunInfo.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
                    Mode: "list",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("relative"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](7324.19),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](6945.52),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.MetricsStore{},
                    Expression: criblcontrolplanesdkgo.Pointer("true"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("1MB"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("10MB"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            Executor: components.ExecutorTypeRunnableJobExecutor{
                Type: "<value>",
                StoreTaskResults: criblcontrolplanesdkgo.Pointer(true),
                Conf: &components.ExecutorSpecificSettingsTypeRunnableJobExecutorExecutor{},
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
            Description: criblcontrolplanesdkgo.Pointer("cautiously annual worth before cleaner parsnip nimble onset once burgeon"),
            Type: components.JobTypeOptionsRunnableJobCollectionCollection,
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            RemoveFields: []string{
                "<value 1>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(false),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(true),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("*/5 * * * *"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1),
                Run: &components.RunSettingsTypeSavedJobCollectionSchedule{
                    Type: components.RunSettingsTypeSavedJobCollectionScheduleTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1),
                    LogLevel: components.LogLevelOptionsRunnableJobCollectionScheduleRunInfo.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
                    Mode: "list",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("relative"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](2728.52),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](4162.9),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.MetricsStore{},
                    Expression: criblcontrolplanesdkgo.Pointer("true"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("1MB"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("10MB"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Description: criblcontrolplanesdkgo.Pointer("pomelo outside offensively ew"),
            Type: components.JobTypeOptionsRunnableJobCollectionExecutor,
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
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
                CronSchedule: criblcontrolplanesdkgo.Pointer("*/5 * * * *"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1),
                Run: &components.RunSettingsTypeSavedJobCollectionSchedule{
                    Type: components.RunSettingsTypeSavedJobCollectionScheduleTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1),
                    LogLevel: components.LogLevelOptionsRunnableJobCollectionScheduleRunInfo.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
                    Mode: "list",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("relative"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](432.8),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](2023.34),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.MetricsStore{},
                    Expression: criblcontrolplanesdkgo.Pointer("true"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("1MB"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("10MB"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Executor: components.ExecutorTypeRunnableJobExecutor{
                Type: "<value>",
                StoreTaskResults: criblcontrolplanesdkgo.Pointer(true),
                Conf: &components.ExecutorSpecificSettingsTypeRunnableJobExecutorExecutor{},
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

    res, err := s.Collectors.Create(ctx, components.CreateSavedJobSavedJobExecutor(
        components.SavedJobExecutor{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("pomelo outside offensively ew"),
            Type: components.JobTypeOptionsRunnableJobCollectionExecutor,
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
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
                CronSchedule: criblcontrolplanesdkgo.Pointer("*/5 * * * *"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1),
                Run: &components.RunSettingsTypeSavedJobCollectionSchedule{
                    Type: components.RunSettingsTypeSavedJobCollectionScheduleTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1),
                    LogLevel: components.LogLevelOptionsRunnableJobCollectionScheduleRunInfo.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
                    Mode: "list",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("relative"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](432.8),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](2023.34),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.MetricsStore{},
                    Expression: criblcontrolplanesdkgo.Pointer("true"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("1MB"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("10MB"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Executor: components.ExecutorTypeRunnableJobExecutor{
                Type: "<value>",
                StoreTaskResults: criblcontrolplanesdkgo.Pointer(true),
                Conf: &components.ExecutorSpecificSettingsTypeRunnableJobExecutorExecutor{},
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
            Description: criblcontrolplanesdkgo.Pointer("cautiously annual worth before cleaner parsnip nimble onset once burgeon"),
            Type: components.JobTypeOptionsRunnableJobCollectionCollection,
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            RemoveFields: []string{
                "<value 1>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(false),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(true),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("*/5 * * * *"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1),
                Run: &components.RunSettingsTypeSavedJobCollectionSchedule{
                    Type: components.RunSettingsTypeSavedJobCollectionScheduleTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1),
                    LogLevel: components.LogLevelOptionsRunnableJobCollectionScheduleRunInfo.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
                    Mode: "list",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("relative"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](2728.52),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](4162.9),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.MetricsStore{},
                    Expression: criblcontrolplanesdkgo.Pointer("true"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("1MB"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("10MB"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Type: components.JobTypeOptionsRunnableJobCollectionExecutor,
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
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
                CronSchedule: criblcontrolplanesdkgo.Pointer("*/5 * * * *"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1),
                Run: &components.RunSettingsTypeSavedJobCollectionSchedule{
                    Type: components.RunSettingsTypeSavedJobCollectionScheduleTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1),
                    LogLevel: components.LogLevelOptionsRunnableJobCollectionScheduleRunInfo.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
                    Mode: "list",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("relative"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](432.8),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](2023.34),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.MetricsStore{},
                    Expression: criblcontrolplanesdkgo.Pointer("true"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("1MB"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("10MB"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            WorkerAffinity: criblcontrolplanesdkgo.Pointer(false),
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
                                Enabled: criblcontrolplanesdkgo.Pointer(true),
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
                },
                StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
                SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
                Preprocess: &components.PreprocessType{
                    Disabled: true,
                    Command: criblcontrolplanesdkgo.Pointer("<value>"),
                    Args: []string{
                        "<value 1>",
                    },
                },
                ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
                Metadata: []components.ItemsTypeMetadata{
                    components.ItemsTypeMetadata{
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

    res, err := s.Collectors.Create(ctx, components.CreateSavedJobSavedJobScheduledSearch(
        components.SavedJobScheduledSearch{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("pomelo outside offensively ew"),
            Type: components.JobTypeOptionsRunnableJobCollectionExecutor,
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
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
                CronSchedule: criblcontrolplanesdkgo.Pointer("*/5 * * * *"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1),
                Run: &components.RunSettingsTypeSavedJobCollectionSchedule{
                    Type: components.RunSettingsTypeSavedJobCollectionScheduleTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1),
                    LogLevel: components.LogLevelOptionsRunnableJobCollectionScheduleRunInfo.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
                    Mode: "list",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("relative"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](432.8),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](2023.34),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.MetricsStore{},
                    Expression: criblcontrolplanesdkgo.Pointer("true"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("1MB"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("10MB"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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

    res, err := s.Collectors.Create(ctx, components.CreateSavedJobSavedJobScheduledSearch(
        components.SavedJobScheduledSearch{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("cautiously annual worth before cleaner parsnip nimble onset once burgeon"),
            Type: components.JobTypeOptionsRunnableJobCollectionCollection,
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            RemoveFields: []string{
                "<value 1>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(false),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(true),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("*/5 * * * *"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1),
                Run: &components.RunSettingsTypeSavedJobCollectionSchedule{
                    Type: components.RunSettingsTypeSavedJobCollectionScheduleTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1),
                    LogLevel: components.LogLevelOptionsRunnableJobCollectionScheduleRunInfo.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
                    Mode: "list",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("relative"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](2728.52),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](4162.9),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.MetricsStore{},
                    Expression: criblcontrolplanesdkgo.Pointer("true"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("1MB"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("10MB"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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

    res, err := s.Collectors.Create(ctx, components.CreateSavedJobSavedJobScheduledSearch(
        components.SavedJobScheduledSearch{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("cautiously annual worth before cleaner parsnip nimble onset once burgeon"),
            Type: components.JobTypeOptionsRunnableJobCollectionCollection,
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            RemoveFields: []string{
                "<value 1>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(false),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Skippable: criblcontrolplanesdkgo.Pointer(true),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("*/5 * * * *"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1),
                Run: &components.RunSettingsTypeSavedJobCollectionSchedule{
                    Type: components.RunSettingsTypeSavedJobCollectionScheduleTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1),
                    LogLevel: components.LogLevelOptionsRunnableJobCollectionScheduleRunInfo.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
                    Mode: "list",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("relative"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](2728.52),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](4162.9),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.MetricsStore{},
                    Expression: criblcontrolplanesdkgo.Pointer("true"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("1MB"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("10MB"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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

    res, err := s.Collectors.Update(ctx, "<id>", components.CreateSavedJobSavedJobExecutor(
        components.SavedJobExecutor{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("wee supposing blaring spectate dark"),
            Type: components.JobTypeOptionsRunnableJobCollectionExecutor,
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            RemoveFields: []string{
                "<value 1>",
                "<value 2>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(false),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(false),
                Skippable: criblcontrolplanesdkgo.Pointer(true),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("*/5 * * * *"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1),
                Run: &components.RunSettingsTypeSavedJobCollectionSchedule{
                    Type: components.RunSettingsTypeSavedJobCollectionScheduleTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1),
                    LogLevel: components.LogLevelOptionsRunnableJobCollectionScheduleRunInfo.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
                    Mode: "list",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("relative"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](4891.8),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](5936.75),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.MetricsStore{},
                    Expression: criblcontrolplanesdkgo.Pointer("true"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("1MB"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("10MB"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Executor: components.ExecutorTypeRunnableJobExecutor{
                Type: "<value>",
                StoreTaskResults: criblcontrolplanesdkgo.Pointer(true),
                Conf: &components.ExecutorSpecificSettingsTypeRunnableJobExecutorExecutor{},
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
            Description: criblcontrolplanesdkgo.Pointer("unabashedly notwithstanding ugh digestive"),
            Type: components.JobTypeOptionsRunnableJobCollectionScheduledSearch,
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
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
                CronSchedule: criblcontrolplanesdkgo.Pointer("*/5 * * * *"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1),
                Run: &components.RunSettingsTypeSavedJobCollectionSchedule{
                    Type: components.RunSettingsTypeSavedJobCollectionScheduleTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1),
                    LogLevel: components.LogLevelOptionsRunnableJobCollectionScheduleRunInfo.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
                    Mode: "list",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("relative"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](9142.96),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](521.08),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.MetricsStore{},
                    Expression: criblcontrolplanesdkgo.Pointer("true"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("1MB"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("10MB"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Description: criblcontrolplanesdkgo.Pointer("unabashedly notwithstanding ugh digestive"),
            Type: components.JobTypeOptionsRunnableJobCollectionScheduledSearch,
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
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
                CronSchedule: criblcontrolplanesdkgo.Pointer("*/5 * * * *"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1),
                Run: &components.RunSettingsTypeSavedJobCollectionSchedule{
                    Type: components.RunSettingsTypeSavedJobCollectionScheduleTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1),
                    LogLevel: components.LogLevelOptionsRunnableJobCollectionScheduleRunInfo.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
                    Mode: "list",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("relative"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](9142.96),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](521.08),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.MetricsStore{},
                    Expression: criblcontrolplanesdkgo.Pointer("true"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("1MB"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("10MB"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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

    res, err := s.Collectors.Update(ctx, "<id>", components.CreateSavedJobSavedJobExecutor(
        components.SavedJobExecutor{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("wee supposing blaring spectate dark"),
            Type: components.JobTypeOptionsRunnableJobCollectionExecutor,
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            RemoveFields: []string{
                "<value 1>",
                "<value 2>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(false),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(false),
                Skippable: criblcontrolplanesdkgo.Pointer(true),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("*/5 * * * *"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1),
                Run: &components.RunSettingsTypeSavedJobCollectionSchedule{
                    Type: components.RunSettingsTypeSavedJobCollectionScheduleTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1),
                    LogLevel: components.LogLevelOptionsRunnableJobCollectionScheduleRunInfo.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
                    Mode: "list",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("relative"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](4891.8),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](5936.75),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.MetricsStore{},
                    Expression: criblcontrolplanesdkgo.Pointer("true"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("1MB"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("10MB"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Executor: components.ExecutorTypeRunnableJobExecutor{
                Type: "<value>",
                StoreTaskResults: criblcontrolplanesdkgo.Pointer(true),
                Conf: &components.ExecutorSpecificSettingsTypeRunnableJobExecutorExecutor{},
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

    res, err := s.Collectors.Update(ctx, "<id>", components.CreateSavedJobSavedJobScheduledSearch(
        components.SavedJobScheduledSearch{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("horde before huzzah apud ha even murky favorite"),
            Type: components.JobTypeOptionsRunnableJobCollectionScheduledSearch,
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            RemoveFields: []string{
                "<value 1>",
                "<value 2>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(false),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(false),
                Skippable: criblcontrolplanesdkgo.Pointer(true),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("*/5 * * * *"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1),
                Run: &components.RunSettingsTypeSavedJobCollectionSchedule{
                    Type: components.RunSettingsTypeSavedJobCollectionScheduleTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1),
                    LogLevel: components.LogLevelOptionsRunnableJobCollectionScheduleRunInfo.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
                    Mode: "list",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("relative"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](6033.18),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](6852.51),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.MetricsStore{},
                    Expression: criblcontrolplanesdkgo.Pointer("true"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("1MB"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("10MB"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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
            Type: components.JobTypeOptionsRunnableJobCollectionScheduledSearch,
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
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
                CronSchedule: criblcontrolplanesdkgo.Pointer("*/5 * * * *"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1),
                Run: &components.RunSettingsTypeSavedJobCollectionSchedule{
                    Type: components.RunSettingsTypeSavedJobCollectionScheduleTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1),
                    LogLevel: components.LogLevelOptionsRunnableJobCollectionScheduleRunInfo.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
                    Mode: "list",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("relative"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](9142.96),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](521.08),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.MetricsStore{},
                    Expression: criblcontrolplanesdkgo.Pointer("true"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("1MB"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("10MB"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            WorkerAffinity: criblcontrolplanesdkgo.Pointer(false),
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
                StaleChannelFlushMs: criblcontrolplanesdkgo.Pointer[float64](10000),
                SendToRoutes: criblcontrolplanesdkgo.Pointer(true),
                Preprocess: &components.PreprocessType{
                    Disabled: true,
                    Command: criblcontrolplanesdkgo.Pointer("<value>"),
                    Args: []string{
                        "<value 1>",
                        "<value 2>",
                    },
                },
                ThrottleRatePerSec: criblcontrolplanesdkgo.Pointer("0"),
                Metadata: []components.ItemsTypeMetadata{
                    components.ItemsTypeMetadata{
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

    res, err := s.Collectors.Update(ctx, "<id>", components.CreateSavedJobSavedJobScheduledSearch(
        components.SavedJobScheduledSearch{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("unabashedly notwithstanding ugh digestive"),
            Type: components.JobTypeOptionsRunnableJobCollectionScheduledSearch,
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
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
                CronSchedule: criblcontrolplanesdkgo.Pointer("*/5 * * * *"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1),
                Run: &components.RunSettingsTypeSavedJobCollectionSchedule{
                    Type: components.RunSettingsTypeSavedJobCollectionScheduleTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1),
                    LogLevel: components.LogLevelOptionsRunnableJobCollectionScheduleRunInfo.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
                    Mode: "list",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("relative"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](9142.96),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](521.08),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.MetricsStore{},
                    Expression: criblcontrolplanesdkgo.Pointer("true"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("1MB"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("10MB"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
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

    res, err := s.Collectors.Update(ctx, "<id>", components.CreateSavedJobSavedJobExecutor(
        components.SavedJobExecutor{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("wee supposing blaring spectate dark"),
            Type: components.JobTypeOptionsRunnableJobCollectionExecutor,
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
            IgnoreGroupJobsLimit: criblcontrolplanesdkgo.Pointer(false),
            RemoveFields: []string{
                "<value 1>",
                "<value 2>",
            },
            ResumeOnBoot: criblcontrolplanesdkgo.Pointer(false),
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(false),
                Skippable: criblcontrolplanesdkgo.Pointer(true),
                ResumeMissed: criblcontrolplanesdkgo.Pointer(false),
                CronSchedule: criblcontrolplanesdkgo.Pointer("*/5 * * * *"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1),
                Run: &components.RunSettingsTypeSavedJobCollectionSchedule{
                    Type: components.RunSettingsTypeSavedJobCollectionScheduleTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1),
                    LogLevel: components.LogLevelOptionsRunnableJobCollectionScheduleRunInfo.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
                    Mode: "list",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("relative"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](4891.8),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](5936.75),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.MetricsStore{},
                    Expression: criblcontrolplanesdkgo.Pointer("true"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("1MB"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("10MB"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Executor: components.ExecutorTypeRunnableJobExecutor{
                Type: "<value>",
                StoreTaskResults: criblcontrolplanesdkgo.Pointer(true),
                Conf: &components.ExecutorSpecificSettingsTypeRunnableJobExecutorExecutor{},
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

    res, err := s.Collectors.Update(ctx, "<id>", components.CreateSavedJobSavedJobExecutor(
        components.SavedJobExecutor{
            ID: criblcontrolplanesdkgo.Pointer("<id>"),
            Description: criblcontrolplanesdkgo.Pointer("unabashedly notwithstanding ugh digestive"),
            Type: components.JobTypeOptionsRunnableJobCollectionScheduledSearch,
            TTL: criblcontrolplanesdkgo.Pointer("4h"),
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
                CronSchedule: criblcontrolplanesdkgo.Pointer("*/5 * * * *"),
                MaxConcurrentRuns: criblcontrolplanesdkgo.Pointer[float64](1),
                Run: &components.RunSettingsTypeSavedJobCollectionSchedule{
                    Type: components.RunSettingsTypeSavedJobCollectionScheduleTypeCollection.ToPointer(),
                    RescheduleDroppedTasks: criblcontrolplanesdkgo.Pointer(true),
                    MaxTaskReschedule: criblcontrolplanesdkgo.Pointer[float64](1),
                    LogLevel: components.LogLevelOptionsRunnableJobCollectionScheduleRunInfo.ToPointer(),
                    JobTimeout: criblcontrolplanesdkgo.Pointer("0"),
                    Mode: "list",
                    TimeRangeType: criblcontrolplanesdkgo.Pointer("relative"),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](9142.96),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](521.08),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.MetricsStore{},
                    Expression: criblcontrolplanesdkgo.Pointer("true"),
                    MinTaskSize: criblcontrolplanesdkgo.Pointer("1MB"),
                    MaxTaskSize: criblcontrolplanesdkgo.Pointer("10MB"),
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Executor: components.ExecutorTypeRunnableJobExecutor{
                Type: "<value>",
                StoreTaskResults: criblcontrolplanesdkgo.Pointer(true),
                Conf: &components.ExecutorSpecificSettingsTypeRunnableJobExecutorExecutor{},
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