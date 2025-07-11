# Versioning
(*Versioning*)

## Overview

Actions related to Versioning

### Available Operations

* [GetSystemProjectsVersionCountByGroupIDAndProjectID](#getsystemprojectsversioncountbygroupidandprojectid) - Get the count of files of changed
* [GetSystemProjectsVersionDiffByGroupIDAndProjectID](#getsystemprojectsversiondiffbygroupidandprojectid) - Get the textual diff for given commit
* [GetSystemProjectsVersionFilesByGroupIDAndProjectID](#getsystemprojectsversionfilesbygroupidandprojectid) - Get the files changed
* [CreateSystemProjectsVersionCommitByGroupIDAndProjectID](#createsystemprojectsversioncommitbygroupidandprojectid) - Commit project changes
* [CreateSystemProjectsVersionCommitByProjectID](#createsystemprojectsversioncommitbyprojectid) - Commit project changes.
* [GetSystemProjectsVersionCountByProjectID](#getsystemprojectsversioncountbyprojectid) - get the count of files of changed
* [GetSystemProjectsVersionDiffByProjectID](#getsystemprojectsversiondiffbyprojectid) - get the textual diff for given commit
* [GetSystemProjectsVersionFilesByProjectID](#getsystemprojectsversionfilesbyprojectid) - get the files changed
* [CreateSystemProjectsVersionRevertByProjectID](#createsystemprojectsversionrevertbyprojectid) - Revert project changes.
* [GetSystemProjectsVersionShowByProjectID](#getsystemprojectsversionshowbyprojectid) - Show project changes.
* [UndoLastCommit](#undolastcommit) - undo the last commit
* [GetVersionBranch](#getversionbranch) - get the list of branches
* [CreateVersionCommit](#createversioncommit) - create a new commit containing the current configs the given log message describing the changes.
* [GetVersionCount](#getversioncount) - get the count of files of changed
* [GetVersionCurrentBranch](#getversioncurrentbranch) - returns git branch that the config is checked out to, if any
* [GetVersionDiff](#getversiondiff) - get the textual diff for given commit
* [GetVersionFiles](#getversionfiles) - get the files changed
* [GetVersionInfo](#getversioninfo) - Get info about versioning availability
* [CreateVersionPush](#createversionpush) - push the current configs to the remote repository.
* [CreateVersionRevert](#createversionrevert) - revert a commit
* [GetVersionShow](#getversionshow) - get the log message and textual diff for given commit
* [GetVersionStatus](#getversionstatus) - get the the working tree status
* [CreateVersionSync](#createversionsync) - syncs with remote repo via POST requests

## GetSystemProjectsVersionCountByGroupIDAndProjectID

Get the count of files of changed

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

    res, err := s.Versioning.GetSystemProjectsVersionCountByGroupIDAndProjectID(ctx, "<id>", "<id>")
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
| `groupID`                                                | *string*                                                 | :heavy_check_mark:                                       | Group Id                                                 |
| `projectID`                                              | *string*                                                 | :heavy_check_mark:                                       | Project Id                                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSystemProjectsVersionCountByGroupIDAndProjectIDResponse](../../models/operations/getsystemprojectsversioncountbygroupidandprojectidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetSystemProjectsVersionDiffByGroupIDAndProjectID

Get the textual diff for given commit

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

    res, err := s.Versioning.GetSystemProjectsVersionDiffByGroupIDAndProjectID(ctx, "<id>", "<id>", nil)
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
| `groupID`                                                | *string*                                                 | :heavy_check_mark:                                       | Group Id                                                 |
| `projectID`                                              | *string*                                                 | :heavy_check_mark:                                       | Project Id                                               |
| `commit`                                                 | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSystemProjectsVersionDiffByGroupIDAndProjectIDResponse](../../models/operations/getsystemprojectsversiondiffbygroupidandprojectidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetSystemProjectsVersionFilesByGroupIDAndProjectID

Get the files changed

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

    res, err := s.Versioning.GetSystemProjectsVersionFilesByGroupIDAndProjectID(ctx, "<id>", "<id>")
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
| `groupID`                                                | *string*                                                 | :heavy_check_mark:                                       | Group Id                                                 |
| `projectID`                                              | *string*                                                 | :heavy_check_mark:                                       | Project Id                                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSystemProjectsVersionFilesByGroupIDAndProjectIDResponse](../../models/operations/getsystemprojectsversionfilesbygroupidandprojectidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateSystemProjectsVersionCommitByGroupIDAndProjectID

Commit project changes

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

    res, err := s.Versioning.CreateSystemProjectsVersionCommitByGroupIDAndProjectID(ctx, "<id>", "<id>", components.ProjectGitCommitParams{
        Message: "<value>",
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
| `groupID`                                                                              | *string*                                                                               | :heavy_check_mark:                                                                     | Group Id                                                                               |
| `projectID`                                                                            | *string*                                                                               | :heavy_check_mark:                                                                     | Project Id                                                                             |
| `projectGitCommitParams`                                                               | [components.ProjectGitCommitParams](../../models/components/projectgitcommitparams.md) | :heavy_check_mark:                                                                     | ProjectGitCommitParams object                                                          |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateSystemProjectsVersionCommitByGroupIDAndProjectIDResponse](../../models/operations/createsystemprojectsversioncommitbygroupidandprojectidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateSystemProjectsVersionCommitByProjectID

Commit project changes.

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

    res, err := s.Versioning.CreateSystemProjectsVersionCommitByProjectID(ctx, "<id>", components.ProjectGitCommitParams{
        Message: "<value>",
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
| `projectID`                                                                            | *string*                                                                               | :heavy_check_mark:                                                                     | Project Id                                                                             |
| `projectGitCommitParams`                                                               | [components.ProjectGitCommitParams](../../models/components/projectgitcommitparams.md) | :heavy_check_mark:                                                                     | ProjectGitCommitParams object                                                          |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateSystemProjectsVersionCommitByProjectIDResponse](../../models/operations/createsystemprojectsversioncommitbyprojectidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetSystemProjectsVersionCountByProjectID

get the count of files of changed

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

    res, err := s.Versioning.GetSystemProjectsVersionCountByProjectID(ctx, "<id>", nil)
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
| `projectID`                                              | *string*                                                 | :heavy_check_mark:                                       | Project ID                                               |
| `id`                                                     | **string*                                                | :heavy_minus_sign:                                       | Commit ID                                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSystemProjectsVersionCountByProjectIDResponse](../../models/operations/getsystemprojectsversioncountbyprojectidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetSystemProjectsVersionDiffByProjectID

get the textual diff for given commit

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

    res, err := s.Versioning.GetSystemProjectsVersionDiffByProjectID(ctx, "<id>", nil, nil, nil)
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
| `projectID`                                              | *string*                                                 | :heavy_check_mark:                                       | Project Id                                               |
| `commit`                                                 | **string*                                                | :heavy_minus_sign:                                       | Commit hash (default is HEAD)                            |
| `filename`                                               | **string*                                                | :heavy_minus_sign:                                       | Filename                                                 |
| `diffLineLimit`                                          | **float64*                                               | :heavy_minus_sign:                                       | Limit maximum lines in the diff                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSystemProjectsVersionDiffByProjectIDResponse](../../models/operations/getsystemprojectsversiondiffbyprojectidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetSystemProjectsVersionFilesByProjectID

get the files changed

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

    res, err := s.Versioning.GetSystemProjectsVersionFilesByProjectID(ctx, "<id>", nil)
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
| `projectID`                                              | *string*                                                 | :heavy_check_mark:                                       | Project ID                                               |
| `id`                                                     | **string*                                                | :heavy_minus_sign:                                       | Commit ID                                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSystemProjectsVersionFilesByProjectIDResponse](../../models/operations/getsystemprojectsversionfilesbyprojectidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateSystemProjectsVersionRevertByProjectID

Revert project changes.

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

    res, err := s.Versioning.CreateSystemProjectsVersionRevertByProjectID(ctx, "<id>", components.ProjectGitCommitParams{
        Message: "<value>",
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
| `projectID`                                                                            | *string*                                                                               | :heavy_check_mark:                                                                     | Project Id                                                                             |
| `projectGitCommitParams`                                                               | [components.ProjectGitCommitParams](../../models/components/projectgitcommitparams.md) | :heavy_check_mark:                                                                     | ProjectGitCommitParams object                                                          |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreateSystemProjectsVersionRevertByProjectIDResponse](../../models/operations/createsystemprojectsversionrevertbyprojectidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetSystemProjectsVersionShowByProjectID

Show project changes.

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

    res, err := s.Versioning.GetSystemProjectsVersionShowByProjectID(ctx, "<id>", nil, nil, nil)
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
| `projectID`                                              | *string*                                                 | :heavy_check_mark:                                       | Project Id                                               |
| `id`                                                     | **string*                                                | :heavy_minus_sign:                                       | Commit ID                                                |
| `filename`                                               | **string*                                                | :heavy_minus_sign:                                       | Filename                                                 |
| `diffLineLimit`                                          | **float64*                                               | :heavy_minus_sign:                                       | Limit maximum lines in the diff                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetSystemProjectsVersionShowByProjectIDResponse](../../models/operations/getsystemprojectsversionshowbyprojectidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UndoLastCommit

undo the last commit

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

    res, err := s.Versioning.UndoLastCommit(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `group`                                                  | **string*                                                | :heavy_minus_sign:                                       | Group ID                                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UndoLastCommitResponse](../../models/operations/undolastcommitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetVersionBranch

get the list of branches

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

    res, err := s.Versioning.GetVersionBranch(ctx)
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

**[*operations.GetVersionBranchResponse](../../models/operations/getversionbranchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateVersionCommit

create a new commit containing the current configs the given log message describing the changes.

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

    res, err := s.Versioning.CreateVersionCommit(ctx, components.GitCommitParams{
        Message: "<value>",
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [components.GitCommitParams](../../models/components/gitcommitparams.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.CreateVersionCommitResponse](../../models/operations/createversioncommitresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetVersionCount

get the count of files of changed

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

    res, err := s.Versioning.GetVersionCount(ctx, nil, nil)
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
| `group`                                                  | **string*                                                | :heavy_minus_sign:                                       | Group ID                                                 |
| `id`                                                     | **string*                                                | :heavy_minus_sign:                                       | Commit ID                                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetVersionCountResponse](../../models/operations/getversioncountresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetVersionCurrentBranch

returns git branch that the config is checked out to, if any

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

    res, err := s.Versioning.GetVersionCurrentBranch(ctx)
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

**[*operations.GetVersionCurrentBranchResponse](../../models/operations/getversioncurrentbranchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetVersionDiff

get the textual diff for given commit

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

    res, err := s.Versioning.GetVersionDiff(ctx, nil, nil, nil, nil)
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
| `commit`                                                 | **string*                                                | :heavy_minus_sign:                                       | Commit hash (default is HEAD)                            |
| `group`                                                  | **string*                                                | :heavy_minus_sign:                                       | Group ID                                                 |
| `filename`                                               | **string*                                                | :heavy_minus_sign:                                       | Filename                                                 |
| `diffLineLimit`                                          | **float64*                                               | :heavy_minus_sign:                                       | Limit maximum lines in the diff                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetVersionDiffResponse](../../models/operations/getversiondiffresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetVersionFiles

get the files changed

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

    res, err := s.Versioning.GetVersionFiles(ctx, nil, nil)
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
| `group`                                                  | **string*                                                | :heavy_minus_sign:                                       | Group ID                                                 |
| `id`                                                     | **string*                                                | :heavy_minus_sign:                                       | Commit ID                                                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetVersionFilesResponse](../../models/operations/getversionfilesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetVersionInfo

Get info about versioning availability

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

    res, err := s.Versioning.GetVersionInfo(ctx)
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

**[*operations.GetVersionInfoResponse](../../models/operations/getversioninforesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateVersionPush

push the current configs to the remote repository.

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

    res, err := s.Versioning.CreateVersionPush(ctx)
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

## CreateVersionRevert

revert a commit

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

    res, err := s.Versioning.CreateVersionRevert(ctx, components.GitRevertParams{
        Commit: "<value>",
        Message: "<value>",
    }, nil)
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
| `group`                                                                  | **string*                                                                | :heavy_minus_sign:                                                       | Group ID                                                                 |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.CreateVersionRevertResponse](../../models/operations/createversionrevertresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetVersionShow

get the log message and textual diff for given commit

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

    res, err := s.Versioning.GetVersionShow(ctx, nil, nil, nil, nil)
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
| `commit`                                                 | **string*                                                | :heavy_minus_sign:                                       | Commit hash (default is HEAD)                            |
| `group`                                                  | **string*                                                | :heavy_minus_sign:                                       | Group ID                                                 |
| `filename`                                               | **string*                                                | :heavy_minus_sign:                                       | Filename                                                 |
| `diffLineLimit`                                          | **float64*                                               | :heavy_minus_sign:                                       | Limit maximum lines in the diff                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetVersionShowResponse](../../models/operations/getversionshowresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetVersionStatus

get the the working tree status

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

    res, err := s.Versioning.GetVersionStatus(ctx, nil)
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
| `group`                                                  | **string*                                                | :heavy_minus_sign:                                       | Group ID                                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetVersionStatusResponse](../../models/operations/getversionstatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateVersionSync

syncs with remote repo via POST requests

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

    res, err := s.Versioning.CreateVersionSync(ctx)
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

**[*operations.CreateVersionSyncResponse](../../models/operations/createversionsyncresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |