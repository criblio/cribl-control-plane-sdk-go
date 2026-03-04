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

    res, err := s.Collectors.Create(ctx, components.SavedJobCreateUpdate{
        ID: "azure-blob-collector",
        Schedule: &components.ScheduleOpts{
            CronSchedule: criblcontrolplanesdkgo.Pointer("0 */8 * * *"),
            Enabled: true,
            Run: components.RunSettings{
                AdditionalProperties: map[string]any{
                    "mode": "run",
                    "timeRangeType": "relative",
                    "earliest": -300,
                    "expression": "true",
                    "logLevel": "info",
                },
            },
            Tz: criblcontrolplanesdkgo.Pointer("UTC"),
        },
        Type: "collection",
    }, criblcontrolplanesdkgo.Pointer("<value>"))
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

    res, err := s.Collectors.Create(ctx, components.SavedJobCreateUpdate{
        ID: "cribl-lake-collector",
        Schedule: &components.ScheduleOpts{
            CronSchedule: criblcontrolplanesdkgo.Pointer("0 */2 * * *"),
            Enabled: true,
            Run: components.RunSettings{
                AdditionalProperties: map[string]any{
                    "mode": "run",
                    "timeRangeType": "relative",
                    "earliest": -300,
                    "expression": "true",
                    "logLevel": "info",
                },
            },
            Tz: criblcontrolplanesdkgo.Pointer("UTC"),
        },
        Type: "collection",
    }, criblcontrolplanesdkgo.Pointer("<value>"))
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

    res, err := s.Collectors.Create(ctx, components.SavedJobCreateUpdate{
        ID: "database-collector",
        Schedule: &components.ScheduleOpts{
            CronSchedule: criblcontrolplanesdkgo.Pointer("0 2 * * *"),
            Enabled: true,
            Run: components.RunSettings{
                AdditionalProperties: map[string]any{
                    "mode": "run",
                    "timeRangeType": "relative",
                    "earliest": -300,
                    "expression": "true",
                    "logLevel": "info",
                },
            },
            Tz: criblcontrolplanesdkgo.Pointer("UTC"),
        },
        Type: "collection",
    }, criblcontrolplanesdkgo.Pointer("<value>"))
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

    res, err := s.Collectors.Create(ctx, components.SavedJobCreateUpdate{
        ID: "filesystem-collector",
        Schedule: &components.ScheduleOpts{
            CronSchedule: criblcontrolplanesdkgo.Pointer("0 */2 * * *"),
            Enabled: true,
            Run: components.RunSettings{
                AdditionalProperties: map[string]any{
                    "mode": "run",
                    "timeRangeType": "relative",
                    "earliest": -300,
                    "expression": "true",
                    "logLevel": "info",
                },
            },
            Tz: criblcontrolplanesdkgo.Pointer("UTC"),
        },
        Type: "collection",
    }, criblcontrolplanesdkgo.Pointer("<value>"))
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

    res, err := s.Collectors.Create(ctx, components.SavedJobCreateUpdate{
        ID: "gcs-collector",
        Schedule: &components.ScheduleOpts{
            CronSchedule: criblcontrolplanesdkgo.Pointer("0 */12 * * *"),
            Enabled: true,
            Run: components.RunSettings{
                AdditionalProperties: map[string]any{
                    "mode": "run",
                    "timeRangeType": "relative",
                    "earliest": -300,
                    "expression": "true",
                    "logLevel": "info",
                },
            },
            Tz: criblcontrolplanesdkgo.Pointer("UTC"),
        },
        Type: "collection",
    }, criblcontrolplanesdkgo.Pointer("<value>"))
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

    res, err := s.Collectors.Create(ctx, components.SavedJobCreateUpdate{
        ID: "rest-collector",
        Schedule: &components.ScheduleOpts{
            CronSchedule: criblcontrolplanesdkgo.Pointer("0 */4 * * *"),
            Enabled: true,
            Run: components.RunSettings{
                AdditionalProperties: map[string]any{
                    "mode": "run",
                    "timeRangeType": "relative",
                    "earliest": -300,
                    "expression": "true",
                    "logLevel": "info",
                },
            },
            Tz: criblcontrolplanesdkgo.Pointer("UTC"),
        },
        Type: "collection",
    }, criblcontrolplanesdkgo.Pointer("<value>"))
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

    res, err := s.Collectors.Create(ctx, components.SavedJobCreateUpdate{
        ID: "s3-collector",
        Schedule: &components.ScheduleOpts{
            CronSchedule: criblcontrolplanesdkgo.Pointer("0 */6 * * *"),
            Enabled: true,
            Run: components.RunSettings{
                AdditionalProperties: map[string]any{
                    "mode": "run",
                    "timeRangeType": "relative",
                    "earliest": -300,
                    "expression": "true",
                    "logLevel": "info",
                },
            },
            Tz: criblcontrolplanesdkgo.Pointer("UTC"),
        },
        Type: "collection",
    }, criblcontrolplanesdkgo.Pointer("<value>"))
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

    res, err := s.Collectors.Create(ctx, components.SavedJobCreateUpdate{
        ID: "script-collector",
        Schedule: &components.ScheduleOpts{
            CronSchedule: criblcontrolplanesdkgo.Pointer("0 */3 * * *"),
            Enabled: true,
            Run: components.RunSettings{
                AdditionalProperties: map[string]any{
                    "mode": "run",
                    "timeRangeType": "relative",
                    "earliest": -300,
                    "expression": "true",
                    "logLevel": "info",
                },
            },
            Tz: criblcontrolplanesdkgo.Pointer("UTC"),
        },
        Type: "collection",
    }, criblcontrolplanesdkgo.Pointer("<value>"))
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

    res, err := s.Collectors.Create(ctx, components.SavedJobCreateUpdate{
        ID: "splunk-collector",
        Schedule: &components.ScheduleOpts{
            CronSchedule: criblcontrolplanesdkgo.Pointer("0 */1 * * *"),
            Enabled: true,
            Run: components.RunSettings{
                AdditionalProperties: map[string]any{
                    "mode": "run",
                    "timeRangeType": "relative",
                    "earliest": -300,
                    "expression": "true",
                    "logLevel": "info",
                },
            },
            Tz: criblcontrolplanesdkgo.Pointer("UTC"),
        },
        Type: "collection",
    }, criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `savedJobCreateUpdate`                                                             | [components.SavedJobCreateUpdate](../../models/components/savedjobcreateupdate.md) | :heavy_check_mark:                                                                 | SavedJobCreateUpdate object                                                        |
| `criblPack`                                                                        | **string*                                                                          | :heavy_minus_sign:                                                                 | The <code>id</code> of the Pack to create the Collector in.                        |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

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

    res, err := s.Collectors.Update(ctx, "<id>", components.SavedJobCreateUpdate{
        ID: "azure-blob-collector",
        Schedule: &components.ScheduleOpts{
            CronSchedule: criblcontrolplanesdkgo.Pointer("0 */8 * * *"),
            Enabled: true,
            Run: components.RunSettings{
                AdditionalProperties: map[string]any{
                    "mode": "run",
                    "timeRangeType": "relative",
                    "earliest": -300,
                    "expression": "true",
                    "logLevel": "info",
                },
            },
            Tz: criblcontrolplanesdkgo.Pointer("UTC"),
        },
        Type: "collection",
    }, criblcontrolplanesdkgo.Pointer("<value>"))
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

    res, err := s.Collectors.Update(ctx, "<id>", components.SavedJobCreateUpdate{
        ID: "cribl-lake-collector",
        Schedule: &components.ScheduleOpts{
            CronSchedule: criblcontrolplanesdkgo.Pointer("0 */2 * * *"),
            Enabled: true,
            Run: components.RunSettings{
                AdditionalProperties: map[string]any{
                    "mode": "run",
                    "timeRangeType": "relative",
                    "earliest": -300,
                    "expression": "true",
                    "logLevel": "info",
                },
            },
            Tz: criblcontrolplanesdkgo.Pointer("UTC"),
        },
        Type: "collection",
    }, criblcontrolplanesdkgo.Pointer("<value>"))
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

    res, err := s.Collectors.Update(ctx, "<id>", components.SavedJobCreateUpdate{
        ID: "database-collector",
        Schedule: &components.ScheduleOpts{
            CronSchedule: criblcontrolplanesdkgo.Pointer("0 2 * * *"),
            Enabled: true,
            Run: components.RunSettings{
                AdditionalProperties: map[string]any{
                    "mode": "run",
                    "timeRangeType": "relative",
                    "earliest": -300,
                    "expression": "true",
                    "logLevel": "info",
                },
            },
            Tz: criblcontrolplanesdkgo.Pointer("UTC"),
        },
        Type: "collection",
    }, criblcontrolplanesdkgo.Pointer("<value>"))
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

    res, err := s.Collectors.Update(ctx, "<id>", components.SavedJobCreateUpdate{
        ID: "filesystem-collector",
        Schedule: &components.ScheduleOpts{
            CronSchedule: criblcontrolplanesdkgo.Pointer("0 */2 * * *"),
            Enabled: true,
            Run: components.RunSettings{
                AdditionalProperties: map[string]any{
                    "mode": "run",
                    "timeRangeType": "relative",
                    "earliest": -300,
                    "expression": "true",
                    "logLevel": "info",
                },
            },
            Tz: criblcontrolplanesdkgo.Pointer("UTC"),
        },
        Type: "collection",
    }, criblcontrolplanesdkgo.Pointer("<value>"))
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

    res, err := s.Collectors.Update(ctx, "<id>", components.SavedJobCreateUpdate{
        ID: "gcs-collector",
        Schedule: &components.ScheduleOpts{
            CronSchedule: criblcontrolplanesdkgo.Pointer("0 */12 * * *"),
            Enabled: true,
            Run: components.RunSettings{
                AdditionalProperties: map[string]any{
                    "mode": "run",
                    "timeRangeType": "relative",
                    "earliest": -300,
                    "expression": "true",
                    "logLevel": "info",
                },
            },
            Tz: criblcontrolplanesdkgo.Pointer("UTC"),
        },
        Type: "collection",
    }, criblcontrolplanesdkgo.Pointer("<value>"))
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

    res, err := s.Collectors.Update(ctx, "<id>", components.SavedJobCreateUpdate{
        ID: "rest-collector",
        Schedule: &components.ScheduleOpts{
            CronSchedule: criblcontrolplanesdkgo.Pointer("0 */4 * * *"),
            Enabled: true,
            Run: components.RunSettings{
                AdditionalProperties: map[string]any{
                    "mode": "run",
                    "timeRangeType": "relative",
                    "earliest": -300,
                    "expression": "true",
                    "logLevel": "info",
                },
            },
            Tz: criblcontrolplanesdkgo.Pointer("UTC"),
        },
        Type: "collection",
    }, criblcontrolplanesdkgo.Pointer("<value>"))
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

    res, err := s.Collectors.Update(ctx, "<id>", components.SavedJobCreateUpdate{
        ID: "s3-collector",
        Schedule: &components.ScheduleOpts{
            CronSchedule: criblcontrolplanesdkgo.Pointer("0 */6 * * *"),
            Enabled: true,
            Run: components.RunSettings{
                AdditionalProperties: map[string]any{
                    "mode": "run",
                    "timeRangeType": "relative",
                    "earliest": -300,
                    "expression": "true",
                    "logLevel": "info",
                },
            },
            Tz: criblcontrolplanesdkgo.Pointer("UTC"),
        },
        Type: "collection",
    }, criblcontrolplanesdkgo.Pointer("<value>"))
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

    res, err := s.Collectors.Update(ctx, "<id>", components.SavedJobCreateUpdate{
        ID: "script-collector",
        Schedule: &components.ScheduleOpts{
            CronSchedule: criblcontrolplanesdkgo.Pointer("0 */3 * * *"),
            Enabled: true,
            Run: components.RunSettings{
                AdditionalProperties: map[string]any{
                    "mode": "run",
                    "timeRangeType": "relative",
                    "earliest": -300,
                    "expression": "true",
                    "logLevel": "info",
                },
            },
            Tz: criblcontrolplanesdkgo.Pointer("UTC"),
        },
        Type: "collection",
    }, criblcontrolplanesdkgo.Pointer("<value>"))
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

    res, err := s.Collectors.Update(ctx, "<id>", components.SavedJobCreateUpdate{
        ID: "splunk-collector",
        Schedule: &components.ScheduleOpts{
            CronSchedule: criblcontrolplanesdkgo.Pointer("0 */1 * * *"),
            Enabled: true,
            Run: components.RunSettings{
                AdditionalProperties: map[string]any{
                    "mode": "run",
                    "timeRangeType": "relative",
                    "earliest": -300,
                    "expression": "true",
                    "logLevel": "info",
                },
            },
            Tz: criblcontrolplanesdkgo.Pointer("UTC"),
        },
        Type: "collection",
    }, criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedSavedJob != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `id`                                                                               | *string*                                                                           | :heavy_check_mark:                                                                 | The <code>id</code> of the Collector to update.                                    |
| `savedJobCreateUpdate`                                                             | [components.SavedJobCreateUpdate](../../models/components/savedjobcreateupdate.md) | :heavy_check_mark:                                                                 | SavedJobCreateUpdate object                                                        |
| `criblPack`                                                                        | **string*                                                                          | :heavy_minus_sign:                                                                 | The <code>id</code> of the Pack that includes the Collector to update.             |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.UpdateSavedJobByIDResponse](../../models/operations/updatesavedjobbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |