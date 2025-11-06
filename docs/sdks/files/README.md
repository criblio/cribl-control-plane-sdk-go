# Files
(*Versions.Commits.Files*)

## Overview

### Available Operations

* [Count](#count) - Get a count of files that changed since a commit
* [List](#list) - Get the names and statuses of files that changed since a commit

## Count

Get a count of the files that changed since a commit. Default is the latest commit (HEAD).

### Example Usage

<!-- UsageSnippet language="go" operationID="getVersionCount" method="get" path="/version/count" -->
```go
package main

import(
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

    res, err := s.Versions.Commits.Files.Count(ctx, criblcontrolplanesdkgo.Pointer("<id>"), criblcontrolplanesdkgo.Pointer("<id>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedListGitCountResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                   | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ctx`                                                                       | [context.Context](https://pkg.go.dev/context#Context)                       | :heavy_check_mark:                                                          | The context to use for the request.                                         |
| `groupID`                                                                   | **string*                                                                   | :heavy_minus_sign:                                                          | The <code>id</code> of the Worker Group or Edge Fleet to get the count for. |
| `id`                                                                        | **string*                                                                   | :heavy_minus_sign:                                                          | The Git commit hash to use as the starting point for the count.             |
| `opts`                                                                      | [][operations.Option](../../models/operations/option.md)                    | :heavy_minus_sign:                                                          | The options for this request.                                               |

### Response

**[*operations.GetVersionCountResponse](../../models/operations/getversioncountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Get the names and statuses of files that changed since a commit. Default is the latest commit (HEAD).

### Example Usage

<!-- UsageSnippet language="go" operationID="getVersionFiles" method="get" path="/version/files" -->
```go
package main

import(
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

    res, err := s.Versions.Commits.Files.List(ctx, criblcontrolplanesdkgo.Pointer("<id>"), criblcontrolplanesdkgo.Pointer("<id>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedListGitFilesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ctx`                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                   | :heavy_check_mark:                                                                      | The context to use for the request.                                                     |
| `groupID`                                                                               | **string*                                                                               | :heavy_minus_sign:                                                                      | The <code>id</code> of the Worker Group or Edge Fleet to get file names and status for. |
| `id`                                                                                    | **string*                                                                               | :heavy_minus_sign:                                                                      | The Git commit hash to use as the starting point for the request.                       |
| `opts`                                                                                  | [][operations.Option](../../models/operations/option.md)                                | :heavy_minus_sign:                                                                      | The options for this request.                                                           |

### Response

**[*operations.GetVersionFilesResponse](../../models/operations/getversionfilesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |