# LakeDatasets
(*LakeDatasets*)

## Overview

### Available Operations

* [Create](#create) - Create a Lake Dataset
* [List](#list) - List all Lake Datasets
* [Delete](#delete) - Delete a Lake Dataset
* [Get](#get) - Get a Lake Dataset
* [Update](#update) - Update a Lake Dataset

## Create

Create a new Lake Dataset in the specified Lake.

### Example Usage

<!-- UsageSnippet language="go" operationID="createCriblLakeDatasetByLakeId" method="post" path="/products/lake/lakes/{lakeId}/datasets" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.LakeDatasets.Create(ctx, "<id>", components.CriblLakeDataset{
        AcceleratedFields: []string{
            "<value 1>",
            "<value 2>",
        },
        BucketName: criblcontrolplanesdkgo.Pointer("<value>"),
        CacheConnection: &components.CacheConnection{
            AcceleratedFields: []string{
                "<value 1>",
                "<value 2>",
            },
            BackfillStatus: components.CacheConnectionBackfillStatusPending.ToPointer(),
            CacheRef: "<value>",
            CreatedAt: 7795.06,
            LakehouseConnectionType: components.LakehouseConnectionTypeCache.ToPointer(),
            MigrationQueryID: criblcontrolplanesdkgo.Pointer("<id>"),
            RetentionInDays: 1466.58,
        },
        DeletionStartedAt: criblcontrolplanesdkgo.Pointer[float64](8310.58),
        Description: criblcontrolplanesdkgo.Pointer("pleased toothbrush long brush smooth swiftly rightfully phooey chapel"),
        Format: components.CriblLakeDatasetFormatDdss.ToPointer(),
        HTTPDAUsed: criblcontrolplanesdkgo.Pointer(true),
        ID: "<id>",
        Metrics: &components.LakeDatasetMetrics{
            CurrentSizeBytes: 6170.04,
            MetricsDate: "<value>",
        },
        RetentionPeriodInDays: criblcontrolplanesdkgo.Pointer[float64](456.37),
        SearchConfig: &components.LakeDatasetSearchConfig{
            Datatypes: []string{
                "<value 1>",
            },
            Metadata: &components.DatasetMetadata{
                Earliest: "<value>",
                EnableAcceleration: true,
                FieldList: []string{
                    "<value 1>",
                    "<value 2>",
                },
                LatestRunInfo: &components.DatasetMetadataRunInfo{
                    EarliestScannedTime: criblcontrolplanesdkgo.Pointer[float64](4334.7),
                    FinishedAt: criblcontrolplanesdkgo.Pointer[float64](6811.22),
                    LatestScannedTime: criblcontrolplanesdkgo.Pointer[float64](5303.3),
                    ObjectCount: criblcontrolplanesdkgo.Pointer[float64](9489.04),
                },
                ScanMode: components.ScanModeDetailed,
            },
        },
        StorageLocationID: criblcontrolplanesdkgo.Pointer("<id>"),
        ViewName: criblcontrolplanesdkgo.Pointer("<value>"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `lakeID`                                                                   | *string*                                                                   | :heavy_check_mark:                                                         | The <code>id</code> of the Lake to create the Lake Dataset in.             |
| `criblLakeDataset`                                                         | [components.CriblLakeDataset](../../models/components/cribllakedataset.md) | :heavy_check_mark:                                                         | CriblLakeDataset object                                                    |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.CreateCriblLakeDatasetByLakeIDResponse](../../models/operations/createcribllakedatasetbylakeidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Get a list of all Lake Datasets in the specified Lake.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCriblLakeDatasetByLakeId" method="get" path="/products/lake/lakes/{lakeId}/datasets" -->
```go
package main

import(
	"context"
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

    res, err := s.LakeDatasets.List(ctx, operations.GetCriblLakeDatasetByLakeIDRequest{
        LakeID: "<id>",
        StorageLocationID: criblcontrolplanesdkgo.Pointer("<id>"),
        Format: criblcontrolplanesdkgo.Pointer("<value>"),
        ExcludeDDSS: criblcontrolplanesdkgo.Pointer(true),
        ExcludeDeleted: criblcontrolplanesdkgo.Pointer(true),
        ExcludeInternal: criblcontrolplanesdkgo.Pointer(false),
        ExcludeBYOS: criblcontrolplanesdkgo.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetCriblLakeDatasetByLakeIDRequest](../../models/operations/getcribllakedatasetbylakeidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.GetCriblLakeDatasetByLakeIDResponse](../../models/operations/getcribllakedatasetbylakeidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete the specified Lake Dataset in the specified Lake

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteCriblLakeDatasetByLakeIdAndId" method="delete" path="/products/lake/lakes/{lakeId}/datasets/{id}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.LakeDatasets.Delete(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                 | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ctx`                                                                     | [context.Context](https://pkg.go.dev/context#Context)                     | :heavy_check_mark:                                                        | The context to use for the request.                                       |
| `lakeID`                                                                  | *string*                                                                  | :heavy_check_mark:                                                        | The <code>id</code> of the Lake that contains the Lake Dataset to delete. |
| `id`                                                                      | *string*                                                                  | :heavy_check_mark:                                                        | The <code>id</code> of the Lake Dataset to delete.                        |
| `opts`                                                                    | [][operations.Option](../../models/operations/option.md)                  | :heavy_minus_sign:                                                        | The options for this request.                                             |

### Response

**[*operations.DeleteCriblLakeDatasetByLakeIDAndIDResponse](../../models/operations/deletecribllakedatasetbylakeidandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get the specified Lake Dataset in the specified Lake.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCriblLakeDatasetByLakeIdAndId" method="get" path="/products/lake/lakes/{lakeId}/datasets/{id}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.LakeDatasets.Get(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `lakeID`                                                               | *string*                                                               | :heavy_check_mark:                                                     | The <code>id</code> of the Lake that contains the Lake Dataset to get. |
| `id`                                                                   | *string*                                                               | :heavy_check_mark:                                                     | The <code>id</code> of the Lake Dataset to get.                        |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.GetCriblLakeDatasetByLakeIDAndIDResponse](../../models/operations/getcribllakedatasetbylakeidandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update the specified Lake Dataset in the specified Lake.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCriblLakeDatasetByLakeIdAndId" method="patch" path="/products/lake/lakes/{lakeId}/datasets/{id}" -->
```go
package main

import(
	"context"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.Pointer(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.LakeDatasets.Update(ctx, "<id>", "<id>", components.CriblLakeDatasetUpdate{
        AcceleratedFields: []string{
            "<value 1>",
            "<value 2>",
        },
        BucketName: criblcontrolplanesdkgo.Pointer("<value>"),
        CacheConnection: &components.CacheConnection{
            AcceleratedFields: []string{
                "<value 1>",
                "<value 2>",
                "<value 3>",
            },
            BackfillStatus: components.CacheConnectionBackfillStatusIncomplete.ToPointer(),
            CacheRef: "<value>",
            CreatedAt: 267.92,
            LakehouseConnectionType: components.LakehouseConnectionTypeZeroPoint.ToPointer(),
            MigrationQueryID: criblcontrolplanesdkgo.Pointer("<id>"),
            RetentionInDays: 3769.62,
        },
        DeletionStartedAt: criblcontrolplanesdkgo.Pointer[float64](836.59),
        Description: criblcontrolplanesdkgo.Pointer("highlight phew ponder but winding"),
        Format: components.CriblLakeDatasetUpdateFormatJSON.ToPointer(),
        HTTPDAUsed: criblcontrolplanesdkgo.Pointer(true),
        ID: criblcontrolplanesdkgo.Pointer("<id>"),
        Metrics: &components.LakeDatasetMetrics{
            CurrentSizeBytes: 6237.74,
            MetricsDate: "<value>",
        },
        RetentionPeriodInDays: criblcontrolplanesdkgo.Pointer[float64](602.09),
        SearchConfig: &components.LakeDatasetSearchConfig{
            Datatypes: []string{
                "<value 1>",
                "<value 2>",
            },
            Metadata: &components.DatasetMetadata{
                Earliest: "<value>",
                EnableAcceleration: false,
                FieldList: []string{
                    "<value 1>",
                },
                LatestRunInfo: &components.DatasetMetadataRunInfo{
                    EarliestScannedTime: criblcontrolplanesdkgo.Pointer[float64](7659.78),
                    FinishedAt: criblcontrolplanesdkgo.Pointer[float64](6404.38),
                    LatestScannedTime: criblcontrolplanesdkgo.Pointer[float64](4426.77),
                    ObjectCount: criblcontrolplanesdkgo.Pointer[float64](8849.28),
                },
                ScanMode: components.ScanModeDetailed,
            },
        },
        StorageLocationID: criblcontrolplanesdkgo.Pointer("<id>"),
        ViewName: criblcontrolplanesdkgo.Pointer("<value>"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `lakeID`                                                                               | *string*                                                                               | :heavy_check_mark:                                                                     | The <code>id</code> of the Lake that contains the Lake Dataset to update.              |
| `id`                                                                                   | *string*                                                                               | :heavy_check_mark:                                                                     | The <code>id</code> of the Lake Dataset to update.                                     |
| `criblLakeDatasetUpdate`                                                               | [components.CriblLakeDatasetUpdate](../../models/components/cribllakedatasetupdate.md) | :heavy_check_mark:                                                                     | CriblLakeDatasetUpdate object                                                          |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdateCriblLakeDatasetByLakeIDAndIDResponse](../../models/operations/updatecribllakedatasetbylakeidandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |