# Commits
(*Versions.Commits*)

## Overview

### Available Operations

* [Create](#create) - Create a new commit for pending changes to the Cribl configuration
* [Diff](#diff) - Get the diff for a commit
* [List](#list) - List the commit history
* [Push](#push) - Push local commits to the remote repository
* [Revert](#revert) - Revert a commit in the local repository
* [Get](#get) - Get the diff and log message for a commit
* [Undo](#undo) - Discard uncommitted (staged) changes

## Create

Create a new commit for pending changes to the Cribl configuration. Any merge conflicts indicated in the response must be resolved using Git.</br></br>To commit only a subset of configuration changes, specify the files to include in the commit in the <code>files</code> array.

### Example Usage

<!-- UsageSnippet language="go" operationID="createVersionCommit" method="post" path="/version/commit" -->
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

    res, err := s.Versions.Commits.Create(ctx, components.GitCommitParams{
        Effective: criblcontrolplanesdkgo.Bool(false),
        Files: []string{
            "<value 1>",
        },
        Group: criblcontrolplanesdkgo.String("<value>"),
        Message: "<value>",
    }, criblcontrolplanesdkgo.String("<id>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |
| `gitCommitParams`                                                                 | [components.GitCommitParams](../../models/components/gitcommitparams.md)          | :heavy_check_mark:                                                                | GitCommitParams object                                                            |
| `groupID`                                                                         | **string*                                                                         | :heavy_minus_sign:                                                                | The <code>id</code> of the Worker Group or Edge Fleet to create a new commit for. |
| `opts`                                                                            | [][operations.Option](../../models/operations/option.md)                          | :heavy_minus_sign:                                                                | The options for this request.                                                     |

### Response

**[*operations.CreateVersionCommitResponse](../../models/operations/createversioncommitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Diff

Get the diff for a commit. Default is the latest commit (HEAD).

### Example Usage

<!-- UsageSnippet language="go" operationID="getVersionDiff" method="get" path="/version/diff" -->
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

    res, err := s.Versions.Commits.Diff(ctx, criblcontrolplanesdkgo.String("<value>"), criblcontrolplanesdkgo.String("<id>"), criblcontrolplanesdkgo.String("example.file"), criblcontrolplanesdkgo.Float64(6362))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                 | Type                                                                                                                                      | Required                                                                                                                                  | Description                                                                                                                               |
| ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                                     | :heavy_check_mark:                                                                                                                        | The context to use for the request.                                                                                                       |
| `commit`                                                                                                                                  | **string*                                                                                                                                 | :heavy_minus_sign:                                                                                                                        | The Git commit hash to get the diff for.                                                                                                  |
| `groupID`                                                                                                                                 | **string*                                                                                                                                 | :heavy_minus_sign:                                                                                                                        | The <code>id</code> of the Worker Group or Edge Fleet to get the diff for.                                                                |
| `filename`                                                                                                                                | **string*                                                                                                                                 | :heavy_minus_sign:                                                                                                                        | The relative path of the file to get the diff for.                                                                                        |
| `diffLineLimit`                                                                                                                           | **float64*                                                                                                                                | :heavy_minus_sign:                                                                                                                        | Number of lines of the diff to return. Default is 1000. Set to <code>0</code> to return the full diff, regardless of the number of lines. |
| `opts`                                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                                  | :heavy_minus_sign:                                                                                                                        | The options for this request.                                                                                                             |

### Response

**[*operations.GetVersionDiffResponse](../../models/operations/getversiondiffresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

List the commit history.</br></br>Analogous to <code>git log</code> for the Cribl configuration, allowing you to audit and review changes over time.

### Example Usage

<!-- UsageSnippet language="go" operationID="getVersion" method="get" path="/version" -->
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

    res, err := s.Versions.Commits.List(ctx, criblcontrolplanesdkgo.String("<id>"), criblcontrolplanesdkgo.Float64(893.58))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `groupID`                                                                            | **string*                                                                            | :heavy_minus_sign:                                                                   | The <code>id</code> of the Worker Group or Edge Fleet to get the commit history for. |
| `count`                                                                              | **float64*                                                                           | :heavy_minus_sign:                                                                   | Maximum number of commits to return in the response for this request.                |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetVersionResponse](../../models/operations/getversionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Push

Push all local commits from the local repository to the remote repository.

### Example Usage

<!-- UsageSnippet language="go" operationID="createVersionPush" method="post" path="/version/push" -->
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

    res, err := s.Versions.Commits.Push(ctx)
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CreateVersionPushResponse](../../models/operations/createversionpushresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Revert

Revert a commit in the local repository.

### Example Usage

<!-- UsageSnippet language="go" operationID="createVersionRevert" method="post" path="/version/revert" -->
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

    res, err := s.Versions.Commits.Revert(ctx, components.GitRevertParams{
        Commit: "<value>",
        Force: criblcontrolplanesdkgo.Bool(false),
        Message: criblcontrolplanesdkgo.String("<value>"),
    }, criblcontrolplanesdkgo.String("<id>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `gitRevertParams`                                                        | [components.GitRevertParams](../../models/components/gitrevertparams.md) | :heavy_check_mark:                                                       | GitRevertParams object                                                   |
| `groupID`                                                                | **string*                                                                | :heavy_minus_sign:                                                       | Group ID                                                                 |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.CreateVersionRevertResponse](../../models/operations/createversionrevertresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get the diff and log message for a commit. Default is the latest commit (HEAD).

### Example Usage

<!-- UsageSnippet language="go" operationID="getVersionShow" method="get" path="/version/show" -->
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

    res, err := s.Versions.Commits.Get(ctx, criblcontrolplanesdkgo.String("<value>"), criblcontrolplanesdkgo.String("<id>"), criblcontrolplanesdkgo.String("example.file"), criblcontrolplanesdkgo.Float64(7771.94))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                 | Type                                                                                                                                      | Required                                                                                                                                  | Description                                                                                                                               |
| ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                                     | :heavy_check_mark:                                                                                                                        | The context to use for the request.                                                                                                       |
| `commit`                                                                                                                                  | **string*                                                                                                                                 | :heavy_minus_sign:                                                                                                                        | The Git commit hash to retrieve the diff and log message for.                                                                             |
| `groupID`                                                                                                                                 | **string*                                                                                                                                 | :heavy_minus_sign:                                                                                                                        | The <code>id</code> of the Worker Group or Edge Fleet to get the diff and log message for.                                                |
| `filename`                                                                                                                                | **string*                                                                                                                                 | :heavy_minus_sign:                                                                                                                        | The relative path of the file to get the diff and log message for.                                                                        |
| `diffLineLimit`                                                                                                                           | **float64*                                                                                                                                | :heavy_minus_sign:                                                                                                                        | Number of lines of the diff to return. Default is 1000. Set to <code>0</code> to return the full diff, regardless of the number of lines. |
| `opts`                                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                                  | :heavy_minus_sign:                                                                                                                        | The options for this request.                                                                                                             |

### Response

**[*operations.GetVersionShowResponse](../../models/operations/getversionshowresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Undo

Discard all uncommitted (staged) configuration changes, resetting the working directory to the last committed state. Use only if you are certain that you do not need to preserve your local changes.

### Example Usage

<!-- UsageSnippet language="go" operationID="createVersionUndo" method="post" path="/version/undo" -->
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

    res, err := s.Versions.Commits.Undo(ctx, criblcontrolplanesdkgo.String("<id>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |
| `groupID`                                                                                 | **string*                                                                                 | :heavy_minus_sign:                                                                        | The <code>id</code> of the Worker Group or Edge Fleet to undo the uncommited changes for. |
| `opts`                                                                                    | [][operations.Option](../../models/operations/option.md)                                  | :heavy_minus_sign:                                                                        | The options for this request.                                                             |

### Response

**[*operations.CreateVersionUndoResponse](../../models/operations/createversionundoresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |