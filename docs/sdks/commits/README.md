# Versions.Commits

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

### Example Usage: VersionCommitExamplesCommitAll

<!-- UsageSnippet language="go" operationID="createVersionCommit" method="post" path="/version/commit" example="VersionCommitExamplesCommitAll" -->
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

    res, err := s.Versions.Commits.Create(ctx, components.GitCommitBody{
        Message: "Updated pipeline configuration for syslog parsing",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedGitCommitSummary != nil {
        // handle response
    }
}
```
### Example Usage: VersionCommitExamplesCommitSpecificFiles

<!-- UsageSnippet language="go" operationID="createVersionCommit" method="post" path="/version/commit" example="VersionCommitExamplesCommitSpecificFiles" -->
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

    res, err := s.Versions.Commits.Create(ctx, components.GitCommitBody{
        Effective: criblcontrolplanesdkgo.Pointer(true),
        Files: []string{
            "groups/default/local/cribl/pipelines/http_input/conf.yml",
            "groups/default/local/cribl/routes.yml",
        },
        Message: "Update Route and Pipeline for HTTP Sources",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedGitCommitSummary != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [components.GitCommitBody](../../models/components/gitcommitbody.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

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

    res, err := s.Versions.Commits.Diff(ctx, criblcontrolplanesdkgo.Pointer("<value>"), criblcontrolplanesdkgo.Pointer("example.file"), criblcontrolplanesdkgo.Pointer[int64](6362))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedGitDiffResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                 | Type                                                                                                                                      | Required                                                                                                                                  | Description                                                                                                                               |
| ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                                     | :heavy_check_mark:                                                                                                                        | The context to use for the request.                                                                                                       |
| `commit`                                                                                                                                  | **string*                                                                                                                                 | :heavy_minus_sign:                                                                                                                        | The Git commit hash to get the diff for.                                                                                                  |
| `filename`                                                                                                                                | **string*                                                                                                                                 | :heavy_minus_sign:                                                                                                                        | The relative path of the file to get the diff for.                                                                                        |
| `diffLineLimit`                                                                                                                           | **int64*                                                                                                                                  | :heavy_minus_sign:                                                                                                                        | Number of lines of the diff to return. Default is 1000. Set to <code>0</code> to return the full diff, regardless of the number of lines. |
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

    res, err := s.Versions.Commits.List(ctx, criblcontrolplanesdkgo.Pointer[int64](893.58))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedGitLogResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                             | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ctx`                                                                 | [context.Context](https://pkg.go.dev/context#Context)                 | :heavy_check_mark:                                                    | The context to use for the request.                                   |
| `count`                                                               | **int64*                                                              | :heavy_minus_sign:                                                    | Maximum number of commits to return in the response for this request. |
| `opts`                                                                | [][operations.Option](../../models/operations/option.md)              | :heavy_minus_sign:                                                    | The options for this request.                                         |

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

    res, err := s.Versions.Commits.Push(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedString != nil {
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

### Example Usage: VersionRevertExamplesForceRevertWithMessage

<!-- UsageSnippet language="go" operationID="createVersionRevert" method="post" path="/version/revert" example="VersionRevertExamplesForceRevertWithMessage" -->
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

    res, err := s.Versions.Commits.Revert(ctx, components.GitRevertParams{
        Commit: "a1b2c3d4e5f6",
        Force: criblcontrolplanesdkgo.Pointer(true),
        Message: criblcontrolplanesdkgo.Pointer("Revert commit due to misconfiguration in Pipeline settings"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedGitRevertResult != nil {
        // handle response
    }
}
```
### Example Usage: VersionRevertExamplesRevertCommit

<!-- UsageSnippet language="go" operationID="createVersionRevert" method="post" path="/version/revert" example="VersionRevertExamplesRevertCommit" -->
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

    res, err := s.Versions.Commits.Revert(ctx, components.GitRevertParams{
        Commit: "a1b2c3d4e5f6",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedGitRevertResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [components.GitRevertParams](../../models/components/gitrevertparams.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
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

    res, err := s.Versions.Commits.Get(ctx, criblcontrolplanesdkgo.Pointer("<value>"), criblcontrolplanesdkgo.Pointer("example.file"), criblcontrolplanesdkgo.Pointer[int64](7771.94))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedGitShowResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                 | Type                                                                                                                                      | Required                                                                                                                                  | Description                                                                                                                               |
| ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                                     | :heavy_check_mark:                                                                                                                        | The context to use for the request.                                                                                                       |
| `commit`                                                                                                                                  | **string*                                                                                                                                 | :heavy_minus_sign:                                                                                                                        | The Git commit hash to retrieve the diff and log message for.                                                                             |
| `filename`                                                                                                                                | **string*                                                                                                                                 | :heavy_minus_sign:                                                                                                                        | The relative path of the file to get the diff and log message for.                                                                        |
| `diffLineLimit`                                                                                                                           | **int64*                                                                                                                                  | :heavy_minus_sign:                                                                                                                        | Number of lines of the diff to return. Default is 1000. Set to <code>0</code> to return the full diff, regardless of the number of lines. |
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

    res, err := s.Versions.Commits.Undo(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedBoolean != nil {
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

**[*operations.CreateVersionUndoResponse](../../models/operations/createversionundoresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |