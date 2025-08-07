# LakeDatasets
(*LakeDatasets*)

## Overview

### Available Operations

* [Create](#create) - Create a Lake Dataset in the specified Lake
* [List](#list) - List all Lake Datasets in the specified Lake
* [Delete](#delete) - Delete a Lake Dataset in the specified Lake
* [Get](#get) - Retrieve a Lake Dataset in the specified Lake
* [Update](#update) - Update a Lake Dataset in the specified Lake

## Create

Create a Dataset in the specified Lake

### Example Usage

<!-- UsageSnippet language="go" operationID="createCriblLakeDatasetByLakeId" method="post" path="/products/lake/lakes/{lakeId}/datasets" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.LakeDatasets.Create(ctx, "<id>", components.CriblLakeDataset{
        AcceleratedFields: []string{
            "<value 1>",
            "<value 2>",
        },
        BucketName: criblcontrolplanesdkgo.String("<value>"),
        CacheConnection: &components.CacheConnection{
            AcceleratedFields: []string{
                "<value 1>",
                "<value 2>",
            },
            BackfillStatus: components.CacheConnectionBackfillStatusPending.ToPointer(),
            CacheRef: "<value>",
            CreatedAt: 7795.06,
            LakehouseConnectionType: components.LakehouseConnectionTypeCache.ToPointer(),
            MigrationQueryID: criblcontrolplanesdkgo.String("<id>"),
            RetentionInDays: 1466.58,
        },
        DeletionStartedAt: criblcontrolplanesdkgo.Float64(8310.58),
        Description: criblcontrolplanesdkgo.String("pleased toothbrush long brush smooth swiftly rightfully phooey chapel"),
        Format: components.CriblLakeDatasetFormatDdss.ToPointer(),
        HTTPDAUsed: criblcontrolplanesdkgo.Bool(true),
        ID: "<id>",
        RetentionPeriodInDays: criblcontrolplanesdkgo.Float64(456.37),
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
                    EarliestScannedTime: criblcontrolplanesdkgo.Float64(4334.7),
                    FinishedAt: criblcontrolplanesdkgo.Float64(6811.22),
                    LatestScannedTime: criblcontrolplanesdkgo.Float64(5303.3),
                    ObjectCount: criblcontrolplanesdkgo.Float64(9489.04),
                },
                ScanMode: components.ScanModeDetailed,
            },
        },
        StorageLocationID: criblcontrolplanesdkgo.String("<id>"),
        ViewName: criblcontrolplanesdkgo.String("<value>"),
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
| `lakeID`                                                                   | *string*                                                                   | :heavy_check_mark:                                                         | lake id that contains the Datasets                                         |
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

Get the list of Dataset contained in the specified Lake

### Example Usage

<!-- UsageSnippet language="go" operationID="getCriblLakeDatasetByLakeId" method="get" path="/products/lake/lakes/{lakeId}/datasets" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.LakeDatasets.List(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `lakeID`                                                 | *string*                                                 | :heavy_check_mark:                                       | lake id that contains the Datasets                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetCriblLakeDatasetByLakeIDResponse](../../models/operations/getcribllakedatasetbylakeidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a Dataset in the specified Lake

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteCriblLakeDatasetByLakeIdAndId" method="delete" path="/products/lake/lakes/{lakeId}/datasets/{id}" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `lakeID`                                                 | *string*                                                 | :heavy_check_mark:                                       | lake id that contains the Datasets                       |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | dataset id to delete                                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteCriblLakeDatasetByLakeIDAndIDResponse](../../models/operations/deletecribllakedatasetbylakeidandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get a Dataset in the specified Lake

### Example Usage

<!-- UsageSnippet language="go" operationID="getCriblLakeDatasetByLakeIdAndId" method="get" path="/products/lake/lakes/{lakeId}/datasets/{id}" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `lakeID`                                                 | *string*                                                 | :heavy_check_mark:                                       | lake id that contains the Datasets                       |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | dataset id to get                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetCriblLakeDatasetByLakeIDAndIDResponse](../../models/operations/getcribllakedatasetbylakeidandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update a Dataset in the specified Lake

### Example Usage

<!-- UsageSnippet language="go" operationID="updateCriblLakeDatasetByLakeIdAndId" method="patch" path="/products/lake/lakes/{lakeId}/datasets/{id}" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(components.Security{
            BearerAuth: criblcontrolplanesdkgo.String(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
        }),
    )

    res, err := s.LakeDatasets.Update(ctx, "<id>", "<id>", components.CriblLakeDataset{
        AcceleratedFields: []string{
            "<value 1>",
            "<value 2>",
        },
        BucketName: criblcontrolplanesdkgo.String("<value>"),
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
            MigrationQueryID: criblcontrolplanesdkgo.String("<id>"),
            RetentionInDays: 3769.62,
        },
        DeletionStartedAt: criblcontrolplanesdkgo.Float64(836.59),
        Description: criblcontrolplanesdkgo.String("highlight phew ponder but winding"),
        Format: components.CriblLakeDatasetFormatJSON.ToPointer(),
        HTTPDAUsed: criblcontrolplanesdkgo.Bool(true),
        ID: "<id>",
        RetentionPeriodInDays: criblcontrolplanesdkgo.Float64(602.09),
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
                    EarliestScannedTime: criblcontrolplanesdkgo.Float64(7659.78),
                    FinishedAt: criblcontrolplanesdkgo.Float64(6404.38),
                    LatestScannedTime: criblcontrolplanesdkgo.Float64(4426.77),
                    ObjectCount: criblcontrolplanesdkgo.Float64(8849.28),
                },
                ScanMode: components.ScanModeDetailed,
            },
        },
        StorageLocationID: criblcontrolplanesdkgo.String("<id>"),
        ViewName: criblcontrolplanesdkgo.String("<value>"),
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
| `lakeID`                                                                   | *string*                                                                   | :heavy_check_mark:                                                         | lake id that contains the Datasets                                         |
| `id`                                                                       | *string*                                                                   | :heavy_check_mark:                                                         | dataset id to update                                                       |
| `criblLakeDataset`                                                         | [components.CriblLakeDataset](../../models/components/cribllakedataset.md) | :heavy_check_mark:                                                         | CriblLakeDataset object                                                    |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.UpdateCriblLakeDatasetByLakeIDAndIDResponse](../../models/operations/updatecribllakedatasetbylakeidandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |