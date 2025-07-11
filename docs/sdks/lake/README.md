# Lake
(*Lake*)

## Overview

Actions related to Lake

### Available Operations

* [CreateCriblLakeDatasetByLakeID](#createcribllakedatasetbylakeid) - Create a Dataset in the specified Lake
* [GetCriblLakeDatasetByLakeID](#getcribllakedatasetbylakeid) - Get the list of Dataset contained in the specified Lake
* [DeleteCriblLakeDatasetByLakeIDAndID](#deletecribllakedatasetbylakeidandid) - Delete a Dataset in the specified Lake
* [GetCriblLakeDatasetByLakeIDAndID](#getcribllakedatasetbylakeidandid) - Get a Dataset in the specified Lake
* [UpdateCriblLakeDatasetByLakeIDAndID](#updatecribllakedatasetbylakeidandid) - Update a Dataset in the specified Lake
* [CreateCriblLakeStorageLocationByLakeID](#createcribllakestoragelocationbylakeid) - Create a Storage Location in the specified Lake
* [GetCriblLakeStorageLocationByLakeID](#getcribllakestoragelocationbylakeid) - Get the list of Storage Locations contained in the specified Lake
* [DeleteCriblLakeStorageLocationByLakeIDAndID](#deletecribllakestoragelocationbylakeidandid) - Delete a Storage Location in the specified Lake
* [GetCriblLakeStorageLocationByLakeIDAndID](#getcribllakestoragelocationbylakeidandid) - Get a Storage Location in the specified Lake
* [UpdateCriblLakeStorageLocationByLakeIDAndID](#updatecribllakestoragelocationbylakeidandid) - Update a Storage Location in the specified Lake

## CreateCriblLakeDatasetByLakeID

Create a Dataset in the specified Lake

### Example Usage

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
        criblcontrolplanesdkgo.WithSecurity(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
    )

    res, err := s.Lake.CreateCriblLakeDatasetByLakeID(ctx, "<id>", components.CriblLakeDataset{
        BucketName: "<value>",
        ID: "<id>",
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

## GetCriblLakeDatasetByLakeID

Get the list of Dataset contained in the specified Lake

### Example Usage

```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
    )

    res, err := s.Lake.GetCriblLakeDatasetByLakeID(ctx, "<id>")
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

## DeleteCriblLakeDatasetByLakeIDAndID

Delete a Dataset in the specified Lake

### Example Usage

```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
    )

    res, err := s.Lake.DeleteCriblLakeDatasetByLakeIDAndID(ctx, "<id>", "<id>")
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

## GetCriblLakeDatasetByLakeIDAndID

Get a Dataset in the specified Lake

### Example Usage

```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
    )

    res, err := s.Lake.GetCriblLakeDatasetByLakeIDAndID(ctx, "<id>", "<id>")
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

## UpdateCriblLakeDatasetByLakeIDAndID

Update a Dataset in the specified Lake

### Example Usage

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
        criblcontrolplanesdkgo.WithSecurity(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
    )

    res, err := s.Lake.UpdateCriblLakeDatasetByLakeIDAndID(ctx, "<id>", "<id>", components.CriblLakeDataset{
        BucketName: "<value>",
        ID: "<id>",
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

## CreateCriblLakeStorageLocationByLakeID

Create a Storage Location in the specified Lake

### Example Usage

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
        criblcontrolplanesdkgo.WithSecurity(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
    )

    res, err := s.Lake.CreateCriblLakeStorageLocationByLakeID(ctx, "<id>", components.CriblLakeStorageLocation{
        Config: components.CriblLakeStorageLocationConfig{
            BucketName: "<value>",
            Region: "<value>",
        },
        Credentials: components.Credentials{
            Method: components.CredentialsMethodAuto,
        },
        ID: "<id>",
        Provider: components.ProviderAwsS3,
        Status: components.CreateStatusCriblLakeStorageLocationStatusEnum(
            components.CriblLakeStorageLocationStatusEnumDelayed,
        ),
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `lakeID`                                                                                   | *string*                                                                                   | :heavy_check_mark:                                                                         | lake id that contains the Storage Locations                                                |
| `criblLakeStorageLocation`                                                                 | [components.CriblLakeStorageLocation](../../models/components/cribllakestoragelocation.md) | :heavy_check_mark:                                                                         | CriblLakeStorageLocation object                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateCriblLakeStorageLocationByLakeIDResponse](../../models/operations/createcribllakestoragelocationbylakeidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetCriblLakeStorageLocationByLakeID

Get the list of Storage Locations contained in the specified Lake

### Example Usage

```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
    )

    res, err := s.Lake.GetCriblLakeStorageLocationByLakeID(ctx, "<id>")
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
| `lakeID`                                                 | *string*                                                 | :heavy_check_mark:                                       | lake id that contains the Storage Locations              |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetCriblLakeStorageLocationByLakeIDResponse](../../models/operations/getcribllakestoragelocationbylakeidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteCriblLakeStorageLocationByLakeIDAndID

Delete a Storage Location in the specified Lake

### Example Usage

```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
    )

    res, err := s.Lake.DeleteCriblLakeStorageLocationByLakeIDAndID(ctx, "<id>", "<id>")
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
| `lakeID`                                                 | *string*                                                 | :heavy_check_mark:                                       | lake id that contains the Storage Locations              |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | storage location id to delete                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteCriblLakeStorageLocationByLakeIDAndIDResponse](../../models/operations/deletecribllakestoragelocationbylakeidandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetCriblLakeStorageLocationByLakeIDAndID

Get a Storage Location in the specified Lake

### Example Usage

```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"log"
)

func main() {
    ctx := context.Background()

    s := criblcontrolplanesdkgo.New(
        "https://api.example.com",
        criblcontrolplanesdkgo.WithSecurity(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
    )

    res, err := s.Lake.GetCriblLakeStorageLocationByLakeIDAndID(ctx, "<id>", "<id>")
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
| `lakeID`                                                 | *string*                                                 | :heavy_check_mark:                                       | lake id that contains the Storage Locations              |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | storage location id to get                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetCriblLakeStorageLocationByLakeIDAndIDResponse](../../models/operations/getcribllakestoragelocationbylakeidandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateCriblLakeStorageLocationByLakeIDAndID

Update a Storage Location in the specified Lake

### Example Usage

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
        criblcontrolplanesdkgo.WithSecurity(os.Getenv("CRIBLCONTROLPLANE_BEARER_AUTH")),
    )

    res, err := s.Lake.UpdateCriblLakeStorageLocationByLakeIDAndID(ctx, "<id>", "<id>", components.CriblLakeStorageLocation{
        Config: components.CriblLakeStorageLocationConfig{
            BucketName: "<value>",
            Region: "<value>",
        },
        Credentials: components.Credentials{
            Method: components.CredentialsMethodAuto,
        },
        ID: "<id>",
        Provider: components.ProviderAwsS3,
        Status: components.CreateStatusStatusBlocked(
            components.StatusBlockedBlocked,
        ),
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `lakeID`                                                                                   | *string*                                                                                   | :heavy_check_mark:                                                                         | lake id that contains the Storage Locations                                                |
| `id`                                                                                       | *string*                                                                                   | :heavy_check_mark:                                                                         | storage location id to update                                                              |
| `criblLakeStorageLocation`                                                                 | [components.CriblLakeStorageLocation](../../models/components/cribllakestoragelocation.md) | :heavy_check_mark:                                                                         | CriblLakeStorageLocation object                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.UpdateCriblLakeStorageLocationByLakeIDAndIDResponse](../../models/operations/updatecribllakestoragelocationbylakeidandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |