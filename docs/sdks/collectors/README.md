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

### Example Usage

<!-- UsageSnippet language="go" operationID="createSavedJob" method="post" path="/lib/jobs" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
            RemoveFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Run: &components.ScheduleTypeSavedJobCollectionRunSettings{
                    Type: components.ScheduleTypeSavedJobCollectionTypeCollection.ToPointer(),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](432.8),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](2023.34),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.TimeWarning{},
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Collector: components.CreateCollectorDatabase(
                components.CollectorDatabase{
                    Type: components.CollectorDatabaseTypeDatabase,
                    Conf: components.DatabaseCollectorConf{
                        ConnectionID: "<id>",
                        Query: "<value>",
                        DefaultBreakers: components.HiddenDefaultBreakersOptionsDatabaseCollectorConfCribl.ToPointer(),
                        Scheduling: &components.DatabaseCollectorConfScheduling{
                            StateTracking: &components.DatabaseCollectorConfStateTracking{
                                Enabled: criblcontrolplanesdkgo.Pointer(true),
                            },
                        },
                    },
                },
            ),
            Input: &components.InputTypeSavedJobCollection{
                BreakerRulesets: []string{
                    "<value 1>",
                },
                Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                    Command: criblcontrolplanesdkgo.Pointer("<value>"),
                    Args: []string{
                        "<value 1>",
                    },
                },
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

### Example Usage

<!-- UsageSnippet language="go" operationID="updateSavedJobById" method="patch" path="/lib/jobs/{id}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
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
            RemoveFields: []string{
                "<value 1>",
                "<value 2>",
            },
            Environment: criblcontrolplanesdkgo.Pointer("<value>"),
            Schedule: &components.ScheduleTypeSavedJobCollection{
                Enabled: criblcontrolplanesdkgo.Pointer(true),
                Run: &components.ScheduleTypeSavedJobCollectionRunSettings{
                    Type: components.ScheduleTypeSavedJobCollectionTypeCollection.ToPointer(),
                    Earliest: criblcontrolplanesdkgo.Pointer[float64](9142.96),
                    Latest: criblcontrolplanesdkgo.Pointer[float64](521.08),
                    TimestampTimezone: "<value>",
                    TimeWarning: &components.TimeWarning{},
                },
            },
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Collector: components.CreateCollectorDatabase(
                components.CollectorDatabase{
                    Type: components.CollectorDatabaseTypeDatabase,
                    Conf: components.DatabaseCollectorConf{
                        ConnectionID: "<id>",
                        Query: "<value>",
                        DefaultBreakers: components.HiddenDefaultBreakersOptionsDatabaseCollectorConfCribl.ToPointer(),
                        Scheduling: &components.DatabaseCollectorConfScheduling{
                            StateTracking: &components.DatabaseCollectorConfStateTracking{
                                Enabled: criblcontrolplanesdkgo.Pointer(false),
                            },
                        },
                    },
                },
            ),
            Input: &components.InputTypeSavedJobCollection{
                BreakerRulesets: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
                Preprocess: &components.PreprocessTypeSavedJobCollectionInput{
                    Command: criblcontrolplanesdkgo.Pointer("<value>"),
                    Args: []string{
                        "<value 1>",
                        "<value 2>",
                    },
                },
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