# Packs
(*Packs*)

## Overview

Actions related to Packs

### Available Operations

* [Install](#install) - Install a Pack
* [List](#list) - List all Packs
* [Upload](#upload) - Upload a Pack file
* [Delete](#delete) - Uninstall a Pack
* [Get](#get) - Get a Pack
* [Update](#update) - Upgrade a Pack

## Install

Install a Pack.<br><br>To install an uploaded Pack, provide the <code>source</code> value from the <code>PUT /packs</code> response as the <code>source</code> parameter in the request body.<br><br>To install a Pack by importing from a URL, provide the direct URL location of the <code>.crbl</code> file for the Pack as the <code>source</code> parameter in the request body.<br><br>To install a Pack by importing from a Git repository, provide <code>git+<repo-url></code> as the <code>source</code> parameter in the request body.<br><br>If you do not include the <code>source</code> parameter in the request body, an empty Pack is created.

### Example Usage

<!-- UsageSnippet language="go" operationID="createPacks" method="post" path="/packs" -->
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

    res, err := s.Packs.Install(ctx, components.CreatePackRequestBodyUnionPackRequestBody1(
        components.PackRequestBody1{
            ID: "<id>",
            Spec: criblcontrolplanesdkgo.Pointer("<value>"),
            Version: criblcontrolplanesdkgo.Pointer("<value>"),
            MinLogStreamVersion: criblcontrolplanesdkgo.Pointer("<value>"),
            DisplayName: criblcontrolplanesdkgo.Pointer("June30"),
            Author: criblcontrolplanesdkgo.Pointer("<value>"),
            Description: criblcontrolplanesdkgo.Pointer("and banish crossly abacus"),
            Source: criblcontrolplanesdkgo.Pointer("https://packs.cribl.io/dl/cribl-duo-rest-io/latest/cribl-duo-rest-io-latest.crbl"),
            Tags: &components.PackRequestBodyTags1{
                DataType: []string{
                    "double",
                    "boolean",
                },
                Domain: []string{
                    "delectable-transom.com",
                    "radiant-sightseeing.info",
                },
                Technology: []string{
                    "<value 1>",
                },
                Streamtags: []string{
                    "<value 1>",
                    "<value 2>",
                    "<value 3>",
                },
            },
            AllowCustomFunctions: criblcontrolplanesdkgo.Pointer(true),
            Force: criblcontrolplanesdkgo.Pointer(true),
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPackInstallInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.PackRequestBodyUnion](../../models/components/packrequestbodyunion.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

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

    res, err := s.Packs.List(ctx, criblcontrolplanesdkgo.Pointer("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPackInfo != nil {
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

## Upload

Upload a Pack file. Returns the <code>source</code> ID needed to install the Pack with <code>POST /packs source</code>, which you must call separately.

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePacks" method="put" path="/packs" -->
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

    example, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    res, err := s.Packs.Upload(ctx, "example.file", example)
    if err != nil {
        log.Fatal(err)
    }
    if res.UploadPackResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `filename`                                               | *string*                                                 | :heavy_check_mark:                                       | Filename of the Pack file to upload.                     |
| `requestBody`                                            | *any*                                                    | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UpdatePacksResponse](../../models/operations/updatepacksresponse.md), error**

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

    res, err := s.Packs.Delete(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPackInstallInfo != nil {
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

## Get

Get the specified Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="getPacksById" method="get" path="/packs/{id}" -->
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

    res, err := s.Packs.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPackInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The <code>id</code> of the Pack to get.                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetPacksByIDResponse](../../models/operations/getpacksbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Upgrade the specified Pack.</br></br>If the Pack includes any user–modified versions of default Cribl Knowledge resources such as lookups, copy the modified files locally for safekeeping before upgrading the Pack.Copy the modified files back to the upgraded Pack after you install it with <code>POST /packs</code> to overwrite the default versions in the Pack.</br></br>After you upgrade the Pack, update any Routes, Pipelines, Sources, and Destinations that use the previous Pack version so that they reference the upgraded Pack.

### Example Usage

<!-- UsageSnippet language="go" operationID="updatePacksById" method="patch" path="/packs/{id}" -->
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

    res, err := s.Packs.Update(ctx, "<id>", components.PackUpgradeRequest{
        AllowCustomFunctions: criblcontrolplanesdkgo.Pointer(true),
        Minor: criblcontrolplanesdkgo.Pointer("<value>"),
        Source: "https://github.com/criblpacks/cribl-palo-alto-networks/releases/download/1.1.4/cribl-palo-alto-networks-a3e5a19d-1.1.4.crbl",
        Spec: criblcontrolplanesdkgo.Pointer("<value>"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CountedPackInfo != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `id`                                                                           | *string*                                                                       | :heavy_check_mark:                                                             | The <code>id</code> of the Pack to upgrade.                                    |
| `packUpgradeRequest`                                                           | [components.PackUpgradeRequest](../../models/components/packupgraderequest.md) | :heavy_check_mark:                                                             | PackUpgradeRequest object                                                      |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.UpdatePacksByIDResponse](../../models/operations/updatepacksbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |