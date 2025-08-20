# Packs
(*Packs*)

## Overview

Actions related to Packs

### Available Operations

* [Install](#install) - Install a Pack
* [List](#list) - List all Packs
* [Delete](#delete) - Uninstall a Pack
* [Update](#update) - Upgrade a Pack

## Install

Install a Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPacks" method="post" path="/packs" -->
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

    res, err := s.Packs.Install(ctx, components.PackRequestBody{
        AllowCustomFunctions: criblcontrolplanesdkgo.Bool(false),
        Author: criblcontrolplanesdkgo.String("<value>"),
        Description: criblcontrolplanesdkgo.String("premeditation coincide although"),
        DisplayName: criblcontrolplanesdkgo.String("Myah14"),
        Exports: []string{
            "<value 1>",
        },
        Force: criblcontrolplanesdkgo.Bool(false),
        ID: "<id>",
        Inputs: criblcontrolplanesdkgo.Float64(4076.64),
        MinLogStreamVersion: criblcontrolplanesdkgo.String("<value>"),
        Outputs: criblcontrolplanesdkgo.Float64(2759.4),
        Source: "<value>",
        Spec: criblcontrolplanesdkgo.String("<value>"),
        Tags: &components.PackRequestBodyTags{
            DataType: []string{},
            Domain: []string{},
            Streamtags: []string{
                "<value 1>",
                "<value 2>",
            },
            Technology: []string{
                "<value 1>",
            },
        },
        Version: criblcontrolplanesdkgo.String("<value>"),
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
| `request`                                                                | [components.PackRequestBody](../../models/components/packrequestbody.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.CreatePacksResponse](../../models/operations/createpacksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Get a list of all Packs.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPacks" method="get" path="/packs" -->
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

    res, err := s.Packs.List(ctx, criblcontrolplanesdkgo.String("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                     | Type                                                                                                                                                                                                                          | Required                                                                                                                                                                                                                      | Description                                                                                                                                                                                                                   |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                                            | The context to use for the request.                                                                                                                                                                                           |
| `with`                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                            | Comma-separated list of additional properties to include in the response. When set, the response includes a count of the specified properties in the Pack. Available values are <code>inputs</code> and <code>outputs</code>. |
| `opts`                                                                                                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                            | The options for this request.                                                                                                                                                                                                 |

### Response

**[*operations.GetPacksResponse](../../models/operations/getpacksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Uninstall the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="deletePacksById" method="delete" path="/packs/{id}" -->
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

    res, err := s.Packs.Delete(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Pack to uninstall.            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeletePacksByIDResponse](../../models/operations/deletepacksbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Upgrade the specified Pack.</br></br>If the Pack includes any user–modified versions of default Cribl Knowledge resources such as lookups, copy the modified files locally for safekeeping before upgrading the Pack. Copy the modified files back to the upgraded Pack after you install it with <code>POST /packs</code> to overwrite the default versions in the Pack.</br></br>After you upgrade the Pack, update any Routes, Pipelines, Sources, and Destinations that use the previous Pack version so that they reference the upgraded Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePacksById" method="patch" path="/packs/{id}" -->
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

    res, err := s.Packs.Update(ctx, "<id>", criblcontrolplanesdkgo.String("<value>"), criblcontrolplanesdkgo.String("<value>"), criblcontrolplanesdkgo.String("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |
| `id`                                                        | *string*                                                    | :heavy_check_mark:                                          | The <code>id</code> of the Pack to upgrade.                 |
| `source`                                                    | **string*                                                   | :heavy_minus_sign:                                          | body string required Pack source                            |
| `minor`                                                     | **string*                                                   | :heavy_minus_sign:                                          | body boolean optional Only upgrade to minor/patch versions  |
| `spec`                                                      | **string*                                                   | :heavy_minus_sign:                                          | body string optional Specify a branch, tag or a semver spec |
| `opts`                                                      | [][operations.Option](../../models/operations/option.md)    | :heavy_minus_sign:                                          | The options for this request.                               |

### Response

**[*operations.UpdatePacksByIDResponse](../../models/operations/updatepacksbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |