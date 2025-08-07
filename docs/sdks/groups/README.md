# Groups
(*Groups*)

## Overview

Actions related to Groups

### Available Operations

* [GetConfigVersion](#getconfigversion) - Retrieve the configuration version for a Worker Group or Edge Fleet
* [CreateByProduct](#createbyproduct) - Create a Worker Group or Edge Fleet for the specified Cribl product
* [GetByProduct](#getbyproduct) - List all Worker Groups or Edge Fleets for the specified Cribl product
* [Delete](#delete) - Delete a Worker Group or Edge Fleet
* [Get](#get) - Retrieve a Worker Group or Edge Fleet
* [Update](#update) - Update a Worker Group or Edge Fleet
* [DeployCommits](#deploycommits) - Deploy commits to a Worker Group or Edge Fleet
* [GetTeamAccessControlListByProduct](#getteamaccesscontrollistbyproduct) - Retrieve the Access Control List (ACL) for teams with permissions on a Worker Group or Edge Fleet for the specified Cribl product
* [GetAccessControlList](#getaccesscontrollist) - Retrieve the Access Control List (ACL) for a Worker Group or Edge Fleet

## GetConfigVersion

Get effective bundle version for given Group

### Example Usage

<!-- UsageSnippet language="go" operationID="getGroupsConfigVersionById" method="get" path="/master/groups/{id}/configVersion" -->
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

    res, err := s.Groups.GetConfigVersion(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Group ID                                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetGroupsConfigVersionByIDResponse](../../models/operations/getgroupsconfigversionbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateByProduct

Create a Fleet or Worker Group

### Example Usage

<!-- UsageSnippet language="go" operationID="createProductsGroupsByProduct" method="post" path="/products/{product}/groups" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
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

    res, err := s.Groups.CreateByProduct(ctx, operations.CreateProductsGroupsByProductProductStream, components.ConfigGroup{
        Cloud: &components.ConfigGroupCloud{
            Provider: components.CloudProviderAws.ToPointer(),
            Region: "<value>",
        },
        ConfigVersion: "<value>",
        DeployingWorkerCount: criblcontrolplanesdkgo.Float64(1848.32),
        Description: criblcontrolplanesdkgo.String("director um why forgery apud once er though off"),
        EstimatedIngestRate: criblcontrolplanesdkgo.Float64(6663.53),
        Git: &components.Git{
            Commit: criblcontrolplanesdkgo.String("<value>"),
            LocalChanges: criblcontrolplanesdkgo.Float64(2079.21),
            Log: []components.Commit{
                components.Commit{
                    AuthorEmail: criblcontrolplanesdkgo.String("<value>"),
                    AuthorName: criblcontrolplanesdkgo.String("<value>"),
                    Date: "2024-08-24",
                    Hash: "<value>",
                    Message: "<value>",
                    Short: "<value>",
                },
            },
        },
        ID: "<id>",
        IncompatibleWorkerCount: criblcontrolplanesdkgo.Float64(5487.26),
        Inherits: criblcontrolplanesdkgo.String("<value>"),
        IsFleet: criblcontrolplanesdkgo.Bool(false),
        IsSearch: criblcontrolplanesdkgo.Bool(false),
        LookupDeployments: []components.ConfigGroupLookups{
            components.ConfigGroupLookups{
                Context: "<value>",
                Lookups: []components.ConfigGroupLookupsLookup{
                    components.ConfigGroupLookupsLookup{
                        DeployedVersion: criblcontrolplanesdkgo.String("<value>"),
                        File: "<value>",
                        Version: criblcontrolplanesdkgo.String("<value>"),
                    },
                },
            },
        },
        MaxWorkerAge: criblcontrolplanesdkgo.String("<value>"),
        Name: criblcontrolplanesdkgo.String("<value>"),
        OnPrem: criblcontrolplanesdkgo.Bool(true),
        Provisioned: criblcontrolplanesdkgo.Bool(true),
        Streamtags: []string{
            "<value 1>",
            "<value 2>",
            "<value 3>",
        },
        Tags: criblcontrolplanesdkgo.String("<value>"),
        Type: components.ConfigGroupTypeLakeAccess.ToPointer(),
        UpgradeVersion: criblcontrolplanesdkgo.String("<value>"),
        WorkerCount: criblcontrolplanesdkgo.Float64(851.73),
        WorkerRemoteAccess: criblcontrolplanesdkgo.Bool(false),
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `product`                                                                                                          | [operations.CreateProductsGroupsByProductProduct](../../models/operations/createproductsgroupsbyproductproduct.md) | :heavy_check_mark:                                                                                                 | Cribl Product                                                                                                      |
| `configGroup`                                                                                                      | [components.ConfigGroup](../../models/components/configgroup.md)                                                   | :heavy_check_mark:                                                                                                 | ConfigGroup object                                                                                                 |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.CreateProductsGroupsByProductResponse](../../models/operations/createproductsgroupsbyproductresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetByProduct

Get a list of ConfigGroup objects

### Example Usage

<!-- UsageSnippet language="go" operationID="getProductsGroupsByProduct" method="get" path="/products/{product}/groups" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
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

    res, err := s.Groups.GetByProduct(ctx, operations.GetProductsGroupsByProductProductStream, criblcontrolplanesdkgo.String("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `product`                                                                                                    | [operations.GetProductsGroupsByProductProduct](../../models/operations/getproductsgroupsbyproductproduct.md) | :heavy_check_mark:                                                                                           | Cribl Product                                                                                                |
| `fields`                                                                                                     | **string*                                                                                                    | :heavy_minus_sign:                                                                                           | fields to add to results: git.commit, git.localChanges, git.log                                              |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.GetProductsGroupsByProductResponse](../../models/operations/getproductsgroupsbyproductresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a Fleet or Worker Group

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteGroupsById" method="delete" path="/master/groups/{id}" -->
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

    res, err := s.Groups.Delete(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | Group ID                                                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteGroupsByIDResponse](../../models/operations/deletegroupsbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Get a specific ConfigGroup object

### Example Usage

<!-- UsageSnippet language="go" operationID="getGroupsById" method="get" path="/master/groups/{id}" -->
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

    res, err := s.Groups.Get(ctx, "<id>", criblcontrolplanesdkgo.String("<value>"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                       | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `ctx`                                                           | [context.Context](https://pkg.go.dev/context#Context)           | :heavy_check_mark:                                              | The context to use for the request.                             |
| `id`                                                            | *string*                                                        | :heavy_check_mark:                                              | Group id                                                        |
| `fields`                                                        | **string*                                                       | :heavy_minus_sign:                                              | fields to add to results: git.commit, git.localChanges, git.log |
| `opts`                                                          | [][operations.Option](../../models/operations/option.md)        | :heavy_minus_sign:                                              | The options for this request.                                   |

### Response

**[*operations.GetGroupsByIDResponse](../../models/operations/getgroupsbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update a Fleet or Worker Group

### Example Usage

<!-- UsageSnippet language="go" operationID="updateGroupsById" method="patch" path="/master/groups/{id}" -->
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

    res, err := s.Groups.Update(ctx, "<id>", components.ConfigGroup{
        Cloud: &components.ConfigGroupCloud{
            Provider: components.CloudProviderAws.ToPointer(),
            Region: "<value>",
        },
        ConfigVersion: "<value>",
        DeployingWorkerCount: criblcontrolplanesdkgo.Float64(19.89),
        Description: criblcontrolplanesdkgo.String("jaywalk wrathful truly indeed definitive reflecting almost massive"),
        EstimatedIngestRate: criblcontrolplanesdkgo.Float64(7133.74),
        Git: &components.Git{
            Commit: criblcontrolplanesdkgo.String("<value>"),
            LocalChanges: criblcontrolplanesdkgo.Float64(370.43),
            Log: []components.Commit{
                components.Commit{
                    AuthorEmail: criblcontrolplanesdkgo.String("<value>"),
                    AuthorName: criblcontrolplanesdkgo.String("<value>"),
                    Date: "2024-08-29",
                    Hash: "<value>",
                    Message: "<value>",
                    Short: "<value>",
                },
            },
        },
        ID: "<id>",
        IncompatibleWorkerCount: criblcontrolplanesdkgo.Float64(7081.95),
        Inherits: criblcontrolplanesdkgo.String("<value>"),
        IsFleet: criblcontrolplanesdkgo.Bool(true),
        IsSearch: criblcontrolplanesdkgo.Bool(true),
        LookupDeployments: []components.ConfigGroupLookups{
            components.ConfigGroupLookups{
                Context: "<value>",
                Lookups: []components.ConfigGroupLookupsLookup{
                    components.ConfigGroupLookupsLookup{
                        DeployedVersion: criblcontrolplanesdkgo.String("<value>"),
                        File: "<value>",
                        Version: criblcontrolplanesdkgo.String("<value>"),
                    },
                },
            },
        },
        MaxWorkerAge: criblcontrolplanesdkgo.String("<value>"),
        Name: criblcontrolplanesdkgo.String("<value>"),
        OnPrem: criblcontrolplanesdkgo.Bool(true),
        Provisioned: criblcontrolplanesdkgo.Bool(true),
        Streamtags: []string{
            "<value 1>",
        },
        Tags: criblcontrolplanesdkgo.String("<value>"),
        Type: components.ConfigGroupTypeLakeAccess.ToPointer(),
        UpgradeVersion: criblcontrolplanesdkgo.String("<value>"),
        WorkerCount: criblcontrolplanesdkgo.Float64(9020.63),
        WorkerRemoteAccess: criblcontrolplanesdkgo.Bool(true),
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

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `id`                                                             | *string*                                                         | :heavy_check_mark:                                               | Group ID                                                         |
| `configGroup`                                                    | [components.ConfigGroup](../../models/components/configgroup.md) | :heavy_check_mark:                                               | ConfigGroup object                                               |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.UpdateGroupsByIDResponse](../../models/operations/updategroupsbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeployCommits

Deploy commits for a Fleet or Worker Group

### Example Usage

<!-- UsageSnippet language="go" operationID="updateGroupsDeployById" method="patch" path="/master/groups/{id}/deploy" -->
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

    res, err := s.Groups.DeployCommits(ctx, "<id>", components.DeployRequest{
        Lookups: []components.DeployRequestLookups{
            components.DeployRequestLookups{
                Context: "<value>",
                Lookups: []components.DeployRequestLookupsLookup{
                    components.DeployRequestLookupsLookup{
                        File: "<value>",
                        Version: "<value>",
                    },
                },
            },
        },
        Version: "<value>",
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

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `id`                                                                 | *string*                                                             | :heavy_check_mark:                                                   | Group ID                                                             |
| `deployRequest`                                                      | [components.DeployRequest](../../models/components/deployrequest.md) | :heavy_check_mark:                                                   | DeployRequest object                                                 |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.UpdateGroupsDeployByIDResponse](../../models/operations/updategroupsdeploybyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetTeamAccessControlListByProduct

ACL of team with permissions for resources in this Group

### Example Usage

<!-- UsageSnippet language="go" operationID="getProductsGroupsAclTeamsByProductAndId" method="get" path="/products/{product}/groups/{id}/acl/teams" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
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

    res, err := s.Groups.GetTeamAccessControlListByProduct(ctx, operations.GetProductsGroupsACLTeamsByProductAndIDProductStream, "<id>", operations.GetProductsGroupsACLTeamsByProductAndIDTypeDatasets.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `product`                                                                                                                              | [operations.GetProductsGroupsACLTeamsByProductAndIDProduct](../../models/operations/getproductsgroupsaclteamsbyproductandidproduct.md) | :heavy_check_mark:                                                                                                                     | Cribl Product                                                                                                                          |
| `id`                                                                                                                                   | *string*                                                                                                                               | :heavy_check_mark:                                                                                                                     | Group ID                                                                                                                               |
| `type_`                                                                                                                                | [*operations.GetProductsGroupsACLTeamsByProductAndIDType](../../models/operations/getproductsgroupsaclteamsbyproductandidtype.md)      | :heavy_minus_sign:                                                                                                                     | resource type by which to filter access levels                                                                                         |
| `opts`                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                               | :heavy_minus_sign:                                                                                                                     | The options for this request.                                                                                                          |

### Response

**[*operations.GetProductsGroupsACLTeamsByProductAndIDResponse](../../models/operations/getproductsgroupsaclteamsbyproductandidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetAccessControlList

ACL of members with permissions for resources in this Group

### Example Usage

<!-- UsageSnippet language="go" operationID="getGroupsAclById" method="get" path="/master/groups/{id}/acl" -->
```go
package main

import(
	"context"
	criblcontrolplanesdkgo "github.com/criblio/cribl-control-plane-sdk-go"
	"os"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"github.com/criblio/cribl-control-plane-sdk-go/models/operations"
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

    res, err := s.Groups.GetAccessControlList(ctx, "<id>", operations.GetGroupsACLByIDTypeInsights.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                           | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `ctx`                                                                               | [context.Context](https://pkg.go.dev/context#Context)                               | :heavy_check_mark:                                                                  | The context to use for the request.                                                 |
| `id`                                                                                | *string*                                                                            | :heavy_check_mark:                                                                  | Group id                                                                            |
| `type_`                                                                             | [*operations.GetGroupsACLByIDType](../../models/operations/getgroupsaclbyidtype.md) | :heavy_minus_sign:                                                                  | resource type by which to filter access levels                                      |
| `opts`                                                                              | [][operations.Option](../../models/operations/option.md)                            | :heavy_minus_sign:                                                                  | The options for this request.                                                       |

### Response

**[*operations.GetGroupsACLByIDResponse](../../models/operations/getgroupsaclbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.Error    | 500                | application/json   |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |